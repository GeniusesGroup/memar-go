/* For license and copyright information please see LEGAL file in repository */

package syllab

import (
	"bytes"
	"fmt"
	"reflect"
	"sync"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

/*
Benchmark1SyllabPlainDecode-8   	120155781	      9.89 ns/op	       0 B/op	       0 allocs/op
Benchmark1SyllabLibDecode-8     	100000000	      11.2 ns/op	       0 B/op	       0 allocs/op
Benchmark1protoBufLibDecode-8   	 1045512	      1158 ns/op	    2768 B/op	       2 allocs/op

Benchmark1SyllabPlainEncode-8   	 1478232	       780 ns/op	    2688 B/op	       1 allocs/op
Benchmark1SyllabLibEncode-8     	 1456250	       803 ns/op	    2688 B/op	       1 allocs/op
Benchmark1protoBufEncode-8      	 1000000	      1028 ns/op	    2688 B/op	       1 allocs/op

note1: Syllab decode 500X faster than Libgo json decode and 4000X faster than standard GO json!
note2: Syllab encode 7X faster than Libgo json encode and 9X faster than standard GO json!
note3: ProtoBuf decode 5X faster than Libgo json decode and 37X faster than standard GO json!
note4: ProtoBuf encode 5X faster than Libgo json encode and 7X faster than standard GO json!
*/

/*
	Test data
*/

type test1 struct {
	CaptchaID uint64
	Image     []byte
}

var unMarshaledTest1 = test1{
	CaptchaID: 1824074400375950161,
}
var syllabMarshaledTest1 []byte
var protoBufMarshaledTest1 []byte

func init() {
	const sliceLen = 2400
	unMarshaledTest1.Image = make([]byte, sliceLen)
	var j uint8
	for i := 0; i < sliceLen; i++ {
		unMarshaledTest1.Image[i] = j
		j++
	}

	syllabMarshaledTest1 = unMarshaledTest1.syllabPlainEncoder()

	file_test_proto_init()
	var te = Test1{}
	te.CaptchaID = unMarshaledTest1.CaptchaID
	te.Image = unMarshaledTest1.Image
	protoBufMarshaledTest1, _ = proto.Marshal(&te)

	// fmt.Print("Syllab test initialized!!", "\n")
}

/*
	Decode && Encode Benchmarks
*/

func Benchmark1SyllabPlainDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test1
		t.syllabPlainDecoder(syllabMarshaledTest1)
	}
}

func Benchmark1SyllabLibDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test1
		t.syllabLibDecoder(syllabMarshaledTest1)
	}
}

func Benchmark1protoBufLibDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t Test1
		t.protoBufDecoder(protoBufMarshaledTest1)
	}
}

func Benchmark1SyllabPlainEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unMarshaledTest1.syllabPlainEncoder()
	}
}

func Benchmark1SyllabLibEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unMarshaledTest1.syllabLibEncoder()
	}
}

func Benchmark1protoBufEncode(b *testing.B) {
	var te = Test1{}
	te.CaptchaID = unMarshaledTest1.CaptchaID
	te.Image = unMarshaledTest1.Image
	for n := 0; n < b.N; n++ {
		te.protoBufEncoder()
	}
}

/*
	Decode && Encode Tests
*/

func Test1SyllabPlainDecode(b *testing.T) {
	var t test1
	var err = t.syllabPlainDecoder(syllabMarshaledTest1)
	if err != nil {
		fmt.Print(err, "\n")
		b.Fail()
	} else if t.CaptchaID != unMarshaledTest1.CaptchaID {
		fmt.Print("Decoded CaptchaID not same\n")
		b.Fail()
	} else if !bytes.Equal(t.Image, unMarshaledTest1.Image) {
		fmt.Print("Decoded Image not same\n")
		b.Fail()
	}
}

func Test1SyllabLibDecode(b *testing.T) {
	var t test1
	var err = t.syllabLibDecoder(syllabMarshaledTest1)
	if err != nil {
		fmt.Print(err, "\n")
		b.Fail()
	} else if t.CaptchaID != unMarshaledTest1.CaptchaID {
		fmt.Print("Decoded CaptchaID not same\n")
		b.Fail()
	} else if !bytes.Equal(t.Image, unMarshaledTest1.Image) {
		fmt.Print("Decoded Image not same\n")
		b.Fail()
	}
}

func Test1ProtoBufLibDecode(b *testing.T) {
	var t Test1
	var err = t.protoBufDecoder(protoBufMarshaledTest1)
	if err != nil {
		fmt.Print(err, "\n")
		b.Fail()
	} else if t.CaptchaID != unMarshaledTest1.CaptchaID {
		fmt.Print("Decoded CaptchaID not same\n")
		b.Fail()
	} else if !bytes.Equal(t.Image, unMarshaledTest1.Image) {
		fmt.Print("Decoded Image not same\n")
		b.Fail()
	}
}

func Test1SyllabPlainEncode(b *testing.T) {
	var buf = unMarshaledTest1.syllabPlainEncoder()
	if len(buf) != len(syllabMarshaledTest1) {
		b.Fail()
	}
}

func Test1SyllabLibEncode(b *testing.T) {
	var buf = unMarshaledTest1.syllabLibEncoder()
	if len(buf) != len(syllabMarshaledTest1) {
		b.Fail()
	}
}

/*
	Syllab Encoder and Decoder
*/

func (t *test1) syllabPlainDecoder(buf []byte) (err error) {
	t.CaptchaID = uint64(buf[0]) | uint64(buf[1])<<8 | uint64(buf[2])<<16 | uint64(buf[3])<<24 | uint64(buf[4])<<32 | uint64(buf[5])<<40 | uint64(buf[6])<<48 | uint64(buf[7])<<56

	var imageAdd = uint32(buf[8]) | uint32(buf[9])<<8 | uint32(buf[10])<<16 | uint32(buf[11])<<24
	var imageLen = uint32(buf[12]) | uint32(buf[13])<<8 | uint32(buf[14])<<16 | uint32(buf[15])<<24
	t.Image = buf[imageAdd : imageAdd+imageLen]
	return
}

func (t *test1) syllabLibDecoder(buf []byte) (err error) {
	t.CaptchaID = GetUInt64(buf[0:])
	t.Image = GetByteArray(buf, 8)
	return
}

func (t *test1) syllabPlainEncoder() (buf []byte) {
	buf = make([]byte, t.syllabLen())
	var heapAddr = 16

	buf[0] = byte(t.CaptchaID)
	buf[1] = byte(t.CaptchaID >> 8)
	buf[2] = byte(t.CaptchaID >> 16)
	buf[3] = byte(t.CaptchaID >> 24)
	buf[4] = byte(t.CaptchaID >> 32)
	buf[5] = byte(t.CaptchaID >> 40)
	buf[6] = byte(t.CaptchaID >> 48)
	buf[7] = byte(t.CaptchaID >> 56)

	buf[8] = byte(heapAddr) // Heap start index
	buf[9] = byte(heapAddr >> 8)
	buf[10] = byte(heapAddr >> 16)
	buf[11] = byte(heapAddr >> 24)
	var ln = len(t.Image)
	buf[12] = byte(ln)
	buf[13] = byte(ln >> 8)
	buf[14] = byte(ln >> 16)
	buf[15] = byte(ln >> 24)
	copy(buf[heapAddr:], t.Image)

	return
}

func (t *test1) syllabLibEncoder() (buf []byte) {
	buf = make([]byte, t.syllabLen())
	var heapAddr uint32 = 16

	SetUInt64(buf[0:], t.CaptchaID)
	SetByteArray(buf, t.Image, 8, heapAddr)
	return
}

func (t *test1) syllabLen() uint64 {
	return 16 + uint64(len(t.Image)) // 16 >> 8+(1*8)
}

/*
	FlatBuffers Encoder and decoder

```fbs
namespace syllab;
struct test1 {
	CaptchaID:uint64;
	Image:bytes;
}
```
*/

// TODO:::

/*
	ProtoBuf Encoder and decoder

>> protoc --go_out=./ test.proto
```proto
syntax = "proto3";
package syllab;
message test1 {
	fixed64 CaptchaID = 1;
	bytes Image = 2;
}
```
*/

func (x *Test1) protoBufDecoder(buf []byte) (err error) {
	err = proto.Unmarshal(buf, x)
	return
}

func (x *Test1) protoBufEncoder() (buf []byte, err error) {
	buf, err = proto.Marshal(x)
	return
}

type Test1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaptchaID uint64 `protobuf:"fixed64,1,opt,name=CaptchaID,proto3" json:"CaptchaID,omitempty"`
	Image     []byte `protobuf:"bytes,2,opt,name=Image,proto3" json:"Image,omitempty"`
}

func (x *Test1) Reset() {
	*x = Test1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test1) ProtoMessage() {}

func (x *Test1) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test1.ProtoReflect.Descriptor instead.
func (*Test1) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *Test1) GetCaptchaID() uint64 {
	if x != nil {
		return x.CaptchaID
	}
	return 0
}

func (x *Test1) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x79,
	0x6c, 0x6c, 0x61, 0x62, 0x22, 0x3b, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x31, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06,
	0x52, 0x09, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_test_proto_goTypes = []interface{}{
	(*Test1)(nil), // 0: syllab.test1
}
var file_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
