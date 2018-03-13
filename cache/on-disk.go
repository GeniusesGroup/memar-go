//Copyright 2018 SabzCity
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Use binary cache to store data for often access and response to requests
// https://github.com/syndtr/goleveldb

package cache

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/google/btree"
)

const (
	defaultBasePath             = "DiskCache"
	defaultFilePerm os.FileMode = 0666
	defaultPathPerm os.FileMode = 0777
)

var (
	defaultTransform = func(s string) []string { return []string{} }
	errCanceled      = errors.New("canceled")
	errEmptyKey      = errors.New("empty key")
	errBadKey        = errors.New("bad key")
)

// DiskCache implements the DiskCache interface. You shouldn't construct DiskCache
// structures directly; instead, use the New constructor.
type DiskCache struct {
	DiskCacheOptions
	mu        sync.RWMutex
	cache     map[string][]byte
	cacheSize uint64
}

// DiskCacheOptions define a set of properties that dictate DiskCache behavior.
// All values are optional.
type DiskCacheOptions struct {
	BasePath     string
	Transform    TransformFunction
	CacheSizeMax uint64 // bytes
	PathPerm     os.FileMode
	FilePerm     os.FileMode

	Index     Index
	IndexLess LessFunction
}

// TransformFunction transforms a key into a slice of strings, with each
// element in the slice representing a directory in the file path where the
// key's entry will eventually be stored.
//
// For example, if TransformFunc transforms "abcdef" to ["ab", "cde", "f"],
// the final location of the data file will be <basedir>/ab/cde/f/abcdef
type TransformFunction func(s string) []string

// NewDiskCache returns an initialized DiskCache structure, ready to use.
// If the path identified by baseDir already contains data,
// it will be accessible, but not yet cached.
func NewDiskCache(bco DiskCacheOptions) *DiskCache {
	if bco.BasePath == "" {
		bco.BasePath = defaultBasePath
	}
	if bco.Transform == nil {
		bco.Transform = defaultTransform
	}
	if bco.PathPerm == 0 {
		bco.PathPerm = defaultPathPerm
	}
	if bco.FilePerm == 0 {
		bco.FilePerm = defaultFilePerm
	}

	bc := &DiskCache{
		DiskCacheOptions: bco,
		cache:            map[string][]byte{},
		cacheSize:        0,
	}

	if bc.Index != nil && bc.IndexLess != nil {
		bc.Index.Initialize(bc.IndexLess, bc.Keys(nil))
	}

	return bc
}

// Write synchronously writes the key-value pair to disk, making it immediately
// available for reads. Write relies on the filesystem to perform an eventual
// sync to physical media. If you need stronger guarantees, see WriteStream.
func (bc *DiskCache) Write(key string, val []byte) error {
	return bc.WriteStream(key, bytes.NewBuffer(val), false)
}

// WriteStream writes the data represented by the io.Reader to the disk, under
// the provided key. If sync is true, WriteStream performs an explicit sync on
// the file as soon as it's written.
//
// bytes.Buffer provides io.Reader semantics for basic data types.
func (bc *DiskCache) WriteStream(key string, r io.Reader, sync bool) error {
	if len(key) <= 0 {
		return errEmptyKey
	}

	bc.mu.Lock()
	defer bc.mu.Unlock()

	return bc.writeStreamWithLock(key, r, sync)
}

// writeStream does no input validation checking.
// TODO: use atomic FS ops.
func (bc *DiskCache) writeStreamWithLock(key string, r io.Reader, sync bool) error {
	if err := bc.ensurePathWithLock(key); err != nil {
		return fmt.Errorf("ensure path: %s", err)
	}

	mode := os.O_WRONLY | os.O_CREATE | os.O_TRUNC // overwrite if exists
	f, err := os.OpenFile(bc.completeFilename(key), mode, bc.FilePerm)
	if err != nil {
		return fmt.Errorf("open file: %s", err)
	}

	wc := io.WriteCloser(&nopWriteCloser{f})
	if _, err := io.Copy(wc, r); err != nil {
		f.Close() // error deliberately ignored
		return fmt.Errorf("i/o copy: %s", err)
	}

	if sync {
		if err := f.Sync(); err != nil {
			f.Close() // error deliberately ignored
			return fmt.Errorf("file sync: %s", err)
		}
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("file close: %s", err)
	}

	if bc.Index != nil {
		bc.Index.Insert(key)
	}

	bc.bustCacheWithLock(key) // cache only on read

	return nil
}

// Read reads the key and returns the value.
// If the key is available in the cache, Read won't touch the disk.
// If the key is not in the cache, Read will have the side-effect of
// lazily caching the value.
func (bc *DiskCache) Read(key string) ([]byte, error) {
	rc, err := bc.ReadStream(key, false)
	if err != nil {
		return []byte{}, err
	}
	defer rc.Close()
	return ioutil.ReadAll(rc)
}

// ReadStream reads the key and returns the value (data) as an io.ReadCloser.
// If the value is cached from a previous read, and direct is false,
// ReadStream will use the cached value. Otherwise, it will return a handle to
// the file on disk, and cache the data on read.
//
// If direct is true, ReadStream will lazily delete any cached value for the
// key, and return a direct handle to the file on disk.
func (bc *DiskCache) ReadStream(key string, direct bool) (io.ReadCloser, error) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	if val, ok := bc.cache[key]; ok {
		if !direct {
			buf := bytes.NewBuffer(val)
			return ioutil.NopCloser(buf), nil
		}

		go func() {
			bc.mu.Lock()
			defer bc.mu.Unlock()
			bc.uncacheWithLock(key, uint64(len(val)))
		}()
	}

	return bc.readWithRLock(key)
}

// read ignores the cache, and returns an io.ReadCloser representing the
// decompressed data for the given key, streamed from the disk. Clients should
// acquire a read lock on the DiskCache and check the cache themselves before
// calling read.
func (bc *DiskCache) readWithRLock(key string) (io.ReadCloser, error) {
	filename := bc.completeFilename(key)

	fi, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	if fi.IsDir() {
		return nil, os.ErrNotExist
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var r io.Reader
	if bc.CacheSizeMax > 0 {
		r = newSiphon(f, bc, key)
	} else {
		r = &closingReader{f}
	}

	var rc = io.ReadCloser(ioutil.NopCloser(r))

	return rc, nil
}

// closingReader provides a Reader that automatically closes the
// embedded ReadCloser when it reaches EOF
type closingReader struct {
	rc io.ReadCloser
}

func (cr closingReader) Read(p []byte) (int, error) {
	n, err := cr.rc.Read(p)
	if err == io.EOF {
		if closeErr := cr.rc.Close(); closeErr != nil {
			return n, closeErr // close must succeed for Read to succeed
		}
	}
	return n, err
}

// siphon is like a TeeReader: it copies all data read through it to an
// internal buffer, and moves that buffer to the cache at EOF.
type siphon struct {
	f   *os.File
	d   *DiskCache
	key string
	buf *bytes.Buffer
}

// newSiphon constructs a siphoning reader that represents the passed file.
// When a successful series of reads ends in an EOF, the siphon will write
// the buffered data to DiskCache's cache under the given key.
func newSiphon(f *os.File, d *DiskCache, key string) io.Reader {
	return &siphon{
		f:   f,
		d:   d,
		key: key,
		buf: &bytes.Buffer{},
	}
}

// Read implements the io.Reader interface for siphon.
func (s *siphon) Read(p []byte) (int, error) {
	n, err := s.f.Read(p)

	if err == nil {
		return s.buf.Write(p[0:n]) // Write must succeed for Read to succeed
	}

	if err == io.EOF {
		s.d.cacheWithoutLock(s.key, s.buf.Bytes()) // cache may fail
		if closeErr := s.f.Close(); closeErr != nil {
			return n, closeErr // close must succeed for Read to succeed
		}
		return n, err
	}

	return n, err
}

// Erase synchronously erases the given key from the disk and the cache.
func (bc *DiskCache) Erase(key string) error {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	bc.bustCacheWithLock(key)

	// erase from index
	if bc.Index != nil {
		bc.Index.Delete(key)
	}

	// erase from disk
	filename := bc.completeFilename(key)
	if s, err := os.Stat(filename); err == nil {
		if s.IsDir() {
			return errBadKey
		}
		if err = os.Remove(filename); err != nil {
			return err
		}
	} else {
		// Return err as-is so caller can do os.IsNotExist(err).
		return err
	}

	// clean up and return
	bc.pruneDirsWithLock(key)
	return nil
}

// EraseAll will delete all of the data from the store, both in the cache and on
// the disk. Note that EraseAll doesn't distinguish DiskCache-related data from non-
// DiskCache-related data. Care should be taken to always specify a DiskCache base
// directory that is exclusively for DiskCache data.
func (bc *DiskCache) EraseAll() error {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.cache = make(map[string][]byte)
	bc.cacheSize = 0
	return os.RemoveAll(bc.BasePath)
}

// Has returns true if the given key exists.
func (bc *DiskCache) Has(key string) bool {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	if _, ok := bc.cache[key]; ok {
		return true
	}

	filename := bc.completeFilename(key)
	s, err := os.Stat(filename)
	if err != nil {
		return false
	}
	if s.IsDir() {
		return false
	}

	return true
}

// Keys returns a channel that will yield every key accessible by the store,
// in undefined order. If a cancel channel is provided, closing it will
// terminate and close the keys channel.
func (bc *DiskCache) Keys(cancel <-chan struct{}) <-chan string {
	return bc.KeysPrefix("", cancel)
}

// KeysPrefix returns a channel that will yield every key accessible by the
// store with the given prefix, in undefined order. If a cancel channel is
// provided, closing it will terminate and close the keys channel. If the
// provided prefix is the empty string, all keys will be yielded.
func (bc *DiskCache) KeysPrefix(prefix string, cancel <-chan struct{}) <-chan string {
	var prepath string
	if prefix == "" {
		prepath = bc.BasePath
	} else {
		prepath = bc.pathFor(prefix)
	}
	c := make(chan string)
	go func() {
		filepath.Walk(prepath, walker(c, prefix, cancel))
		close(c)
	}()
	return c
}

// walker returns a function which satisfies the filepath.WalkFunc interface.
// It sends every non-directory file entry down the channel c.
func walker(c chan<- string, prefix string, cancel <-chan struct{}) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasPrefix(info.Name(), prefix) {
			return nil // "pass"
		}

		select {
		case c <- info.Name():
		case <-cancel:
			return errCanceled
		}

		return nil
	}
}

// pathFor returns the absolute path for location on the filesystem where the
// data for the given key will be stored.
func (bc *DiskCache) pathFor(key string) string {
	return filepath.Join(bc.BasePath, filepath.Join(bc.Transform(key)...))
}

// ensurePathWithLock is a helper function that generates all necessary
// directories on the filesystem for the given key.
func (bc *DiskCache) ensurePathWithLock(key string) error {
	return os.MkdirAll(bc.pathFor(key), bc.PathPerm)
}

// completeFilename returns the absolute path to the file for the given key.
func (bc *DiskCache) completeFilename(key string) string {
	return filepath.Join(bc.pathFor(key), key)
}

// cacheWithLock attempts to cache the given key-value pair in the store's
// cache. It can fail if the value is larger than the cache's maximum size.
func (bc *DiskCache) cacheWithLock(key string, val []byte) error {
	valueSize := uint64(len(val))
	if err := bc.ensureCacheSpaceWithLock(valueSize); err != nil {
		return fmt.Errorf("%s; not caching", err)
	}

	// be very strict about memory guarantees
	if (bc.cacheSize + valueSize) > bc.CacheSizeMax {
		panic(fmt.Sprintf("failed to make room for value (%d/%d)", valueSize, bc.CacheSizeMax))
	}

	bc.cache[key] = val
	bc.cacheSize += valueSize
	return nil
}

// cacheWithoutLock acquires the store's (write) mutex and calls cacheWithLock.
func (bc *DiskCache) cacheWithoutLock(key string, val []byte) error {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	return bc.cacheWithLock(key, val)
}

func (bc *DiskCache) bustCacheWithLock(key string) {
	if val, ok := bc.cache[key]; ok {
		bc.uncacheWithLock(key, uint64(len(val)))
	}
}

func (bc *DiskCache) uncacheWithLock(key string, sz uint64) {
	bc.cacheSize -= sz
	delete(bc.cache, key)
}

// pruneDirsWithLock deletes empty directories in the path walk leading to the
// key k. Typically this function is called after an Erase is made.
func (bc *DiskCache) pruneDirsWithLock(key string) error {
	pathlist := bc.Transform(key)
	for i := range pathlist {
		dir := filepath.Join(bc.BasePath, filepath.Join(pathlist[:len(pathlist)-i]...))

		// thanks to Steven Blenkinsop for this snippet
		switch fi, err := os.Stat(dir); true {
		case err != nil:
			return err
		case !fi.IsDir():
			panic(fmt.Sprintf("corrupt dirstate at %s", dir))
		}

		nlinks, err := filepath.Glob(filepath.Join(dir, "*"))
		if err != nil {
			return err
		} else if len(nlinks) > 0 {
			return nil // has subdirs -- do not prune
		}
		if err = os.Remove(dir); err != nil {
			return err
		}
	}

	return nil
}

// ensureCacheSpaceWithLock deletes entries from the cache in arbitrary order
// until the cache has at least valueSize bytes available.
func (bc *DiskCache) ensureCacheSpaceWithLock(valueSize uint64) error {
	if valueSize > bc.CacheSizeMax {
		return fmt.Errorf("value size (%d bytes) too large for cache (%d bytes)", valueSize, bc.CacheSizeMax)
	}

	safe := func() bool { return (bc.cacheSize + valueSize) <= bc.CacheSizeMax }

	for key, val := range bc.cache {
		if safe() {
			break
		}

		bc.uncacheWithLock(key, uint64(len(val)))
	}

	if !safe() {
		panic(fmt.Sprintf("%d bytes still won't fit in the cache! (max %d bytes)", valueSize, bc.CacheSizeMax))
	}

	return nil
}

// nopWriteCloser wraps an io.Writer and provides a no-op Close method to
// satisfy the io.WriteCloser interface.
type nopWriteCloser struct {
	io.Writer
}

func (wc *nopWriteCloser) Write(p []byte) (int, error) { return wc.Writer.Write(p) }
func (wc *nopWriteCloser) Close() error                { return nil }

// Index is a generic interface for things that can
// provide an ordered list of keys.
type Index interface {
	Initialize(less LessFunction, keys <-chan string)
	Insert(key string)
	Delete(key string)
	Keys(from string, n int) []string
}

// LessFunction is used to initialize an Index of keys in a specific order.
type LessFunction func(string, string) bool

// btreeString is a custom data type that satisfies the BTree Less interface,
// making the strings it wraps sortable by the BTree package.
type btreeString struct {
	s string
	l LessFunction
}

// Less satisfies the BTree.Less interface using the btreeString's LessFunction.
func (s btreeString) Less(i btree.Item) bool {
	return s.l(s.s, i.(btreeString).s)
}

// BTreeIndex is an implementation of the Index interface using google/btree.
type BTreeIndex struct {
	sync.RWMutex
	LessFunction
	*btree.BTree
}

// Initialize populates the BTree tree with data from the keys channel,
// according to the passed less function. It's destructive to the BTreeIndex.
func (i *BTreeIndex) Initialize(less LessFunction, keys <-chan string) {
	i.Lock()
	defer i.Unlock()
	i.LessFunction = less
	i.BTree = rebuild(less, keys)
}

// Insert inserts the given key (only) into the BTree tree.
func (i *BTreeIndex) Insert(key string) {
	i.Lock()
	defer i.Unlock()
	if i.BTree == nil || i.LessFunction == nil {
		panic("uninitialized index")
	}
	i.BTree.ReplaceOrInsert(btreeString{s: key, l: i.LessFunction})
}

// Delete removes the given key (only) from the BTree tree.
func (i *BTreeIndex) Delete(key string) {
	i.Lock()
	defer i.Unlock()
	if i.BTree == nil || i.LessFunction == nil {
		panic("uninitialized index")
	}
	i.BTree.Delete(btreeString{s: key, l: i.LessFunction})
}

// Keys yields a maximum of n keys in order. If the passed 'from' key is empty,
// Keys will return the first n keys. If the passed 'from' key is non-empty, the
// first key in the returned slice will be the key that immediately follows the
// passed key, in key order.
func (i *BTreeIndex) Keys(from string, n int) []string {
	i.RLock()
	defer i.RUnlock()

	if i.BTree == nil || i.LessFunction == nil {
		panic("uninitialized index")
	}

	if i.BTree.Len() <= 0 {
		return []string{}
	}

	btreeFrom := btreeString{s: from, l: i.LessFunction}
	skipFirst := true
	if len(from) <= 0 || !i.BTree.Has(btreeFrom) {
		// no such key, so fabricate an always-smallest item
		btreeFrom = btreeString{s: "", l: func(string, string) bool { return true }}
		skipFirst = false
	}

	keys := []string{}
	iterator := func(i btree.Item) bool {
		keys = append(keys, i.(btreeString).s)
		return len(keys) < n
	}
	i.BTree.AscendGreaterOrEqual(btreeFrom, iterator)

	if skipFirst && len(keys) > 0 {
		keys = keys[1:]
	}

	return keys
}

// rebuildIndex does the work of regenerating the index
// with the given keys.
func rebuild(less LessFunction, keys <-chan string) *btree.BTree {
	tree := btree.New(2)
	for key := range keys {
		tree.ReplaceOrInsert(btreeString{s: key, l: less})
	}
	return tree
}
