/* For license and copyright information please see the LEGAL file in the code repository */

package duration

// A Second duration represents the elapsed time between two instants as an int64 second count.
// https://en.wikipedia.org/wiki/Second
type Second int64

func (d *Second) ToNanoSecond() (nsec NanoSecond) {
	// TODO::: check overflow??
	nsec = NanoSecond(*d) * OneSecond
	return
}
