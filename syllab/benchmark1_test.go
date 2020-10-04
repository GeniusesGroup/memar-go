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
Benchmark1SyllabUnSafeDecode-8   	120155781	      9.89 ns/op	       0 B/op	       0 allocs/op
Benchmark1SyllabSafeDecode-8     	 1499546	       804 ns/op	    2688 B/op	       1 allocs/op
Benchmark1protoBufDecode-8       	 1000000	      1222 ns/op	    2768 B/op	       2 allocs/op

Benchmark1SyllabUnSafeEncode-8     	 1538323	       780 ns/op	    2688 B/op	       1 allocs/op
Benchmark1SyllabSafeEncode-8	   	 1491774	       795 ns/op	    2688 B/op	       1 allocs/op
Benchmark1protoBufEncode-8      	 1000000	      1039 ns/op	    2688 B/op	       1 allocs/op

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

	syllabMarshaledTest1 = make([]byte, unMarshaledTest1.syllabLen())
	unMarshaledTest1.syllabSafeEncoder(syllabMarshaledTest1)

	file_test_proto_init()
	var te = Test1{
		CaptchaID: unMarshaledTest1.CaptchaID,
		Image:     unMarshaledTest1.Image,
	}
	protoBufMarshaledTest1, _ = proto.Marshal(&te)

	// fmt.Print("Syllab test initialized!!", "\n")
}

/*
	Decode && Encode Benchmarks
*/

func Benchmark1SyllabUnSafeDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test1
		t.syllabUnSafeDecoder(syllabMarshaledTest1)
	}
}

func Benchmark1SyllabSafeDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test1
		t.syllabSafeDecoder(syllabMarshaledTest1)
	}
}

func Benchmark1protoBufDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t Test1
		t.protoBufDecoder(protoBufMarshaledTest1)
	}
}

func Benchmark1SyllabUnSafeEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf = make([]byte, unMarshaledTest1.syllabLen())
		unMarshaledTest1.syllabUnSafeEncoder(buf)
	}
}

func Benchmark1SyllabSafeEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf = make([]byte, unMarshaledTest1.syllabLen())
		unMarshaledTest1.syllabSafeEncoder(buf)
	}
}

func Benchmark1protoBufEncode(b *testing.B) {
	var te = Test1{
		CaptchaID: unMarshaledTest1.CaptchaID,
		Image:     unMarshaledTest1.Image,
	}
	for n := 0; n < b.N; n++ {
		te.protoBufEncoder()
	}
}

/*
	Decode && Encode Tests
*/

func Test1SyllabUnsafeDecode(b *testing.T) {
	var t test1
	var err = t.syllabUnSafeDecoder(syllabMarshaledTest1)
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

func Test1SyllabSafeDecode(b *testing.T) {
	var t test1
	var err = t.syllabSafeDecoder(syllabMarshaledTest1)
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

func Test1ProtoBufDecode(b *testing.T) {
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

func Test1SyllabSafeEncode(b *testing.T) {
	var buf = make([]byte, unMarshaledTest1.syllabLen())
	unMarshaledTest1.syllabSafeEncoder(buf)
	if len(buf) != len(syllabMarshaledTest1) {
		b.Fail()
	}
}

func Test1SyllabUnSafeEncode(b *testing.T) {
	var buf = make([]byte, unMarshaledTest1.syllabLen())
	unMarshaledTest1.syllabUnSafeEncoder(buf)
	if len(buf) != len(syllabMarshaledTest1) {
		b.Fail()
	}
}

/*
	Syllab Encoder and Decoder
*/

func (t *test1) syllabSafeDecoder(buf []byte) (err error) {
	t.CaptchaID = GetUInt64(buf, 0)
	t.Image = GetByteArray(buf, 8)
	return
}

func (t *test1) syllabUnSafeDecoder(buf []byte) (err error) {
	t.CaptchaID = GetUInt64(buf, 0)
	// t.Image = UnsafeGetByteArray(buf, 8)
	var add uint32 = GetUInt32(buf, 8)
	var len uint32 = GetUInt32(buf, 8+4)
	t.Image = buf[add : add+len]
	return
}

func (t *test1) syllabSafeEncoder(buf []byte) {
	var heapAddr uint32 = t.syllabStackLen()
	var ln = uint32(len(t.Image))

	SetUInt64(buf, 0, t.CaptchaID)
	SetUInt32(buf, 8, heapAddr)
	SetUInt32(buf, 12, ln)
	copy(buf[heapAddr:], t.Image)
}

func (t *test1) syllabUnSafeEncoder(buf []byte) {
	var heapAddr uint32 = uint32(t.syllabStackLen())

	SetUInt64(buf, 0, t.CaptchaID)
	SetByteArray(buf, t.Image, 8, heapAddr)
}

func (t *test1) syllabStackLen() (ln uint32) {
	return 16 // 16 >> 8+(1*8)
}

func (t *test1) syllabHeapLen() (ln uint32) {
	ln += uint32(len(t.Image))
	return
}

func (t *test1) syllabLen() int {
	return int(t.syllabStackLen() + t.syllabHeapLen())
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
