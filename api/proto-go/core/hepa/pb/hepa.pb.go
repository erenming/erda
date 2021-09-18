// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: hepa.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize   int64 `protobuf:"varint,1,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	CurPage    int64 `protobuf:"varint,2,opt,name=curPage,proto3" json:"curPage,omitempty"`
	TotalNum   int64 `protobuf:"varint,3,opt,name=totalNum,proto3" json:"totalNum,omitempty"`
	StartIndex int64 `protobuf:"varint,4,opt,name=startIndex,proto3" json:"startIndex,omitempty"`
	EndIndex   int64 `protobuf:"varint,5,opt,name=endIndex,proto3" json:"endIndex,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hepa_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_hepa_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_hepa_proto_rawDescGZIP(), []int{0}
}

func (x *Page) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Page) GetCurPage() int64 {
	if x != nil {
		return x.CurPage
	}
	return 0
}

func (x *Page) GetTotalNum() int64 {
	if x != nil {
		return x.TotalNum
	}
	return 0
}

func (x *Page) GetStartIndex() int64 {
	if x != nil {
		return x.StartIndex
	}
	return 0
}

func (x *Page) GetEndIndex() int64 {
	if x != nil {
		return x.EndIndex
	}
	return 0
}

type PageResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *structpb.Value `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Page   *Page           `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *PageResult) Reset() {
	*x = PageResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hepa_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageResult) ProtoMessage() {}

func (x *PageResult) ProtoReflect() protoreflect.Message {
	mi := &file_hepa_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageResult.ProtoReflect.Descriptor instead.
func (*PageResult) Descriptor() ([]byte, []int) {
	return file_hepa_proto_rawDescGZIP(), []int{1}
}

func (x *PageResult) GetResult() *structpb.Value {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *PageResult) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type NewPageResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  *structpb.Value `protobuf:"bytes,1,opt,name=list,proto3" json:"list,omitempty"`
	Total int64           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *NewPageResult) Reset() {
	*x = NewPageResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hepa_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPageResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPageResult) ProtoMessage() {}

func (x *NewPageResult) ProtoReflect() protoreflect.Message {
	mi := &file_hepa_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPageResult.ProtoReflect.Descriptor instead.
func (*NewPageResult) Descriptor() ([]byte, []int) {
	return file_hepa_proto_rawDescGZIP(), []int{2}
}

func (x *NewPageResult) GetList() *structpb.Value {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *NewPageResult) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_hepa_proto protoreflect.FileDescriptor

var file_hepa_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x72,
	0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x04, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x22, 0x66, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x2e, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x28, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x51, 0x0a, 0x0d, 0x4e, 0x65, 0x77,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x34, 0x5a, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x65, 0x70, 0x61, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hepa_proto_rawDescOnce sync.Once
	file_hepa_proto_rawDescData = file_hepa_proto_rawDesc
)

func file_hepa_proto_rawDescGZIP() []byte {
	file_hepa_proto_rawDescOnce.Do(func() {
		file_hepa_proto_rawDescData = protoimpl.X.CompressGZIP(file_hepa_proto_rawDescData)
	})
	return file_hepa_proto_rawDescData
}

var file_hepa_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_hepa_proto_goTypes = []interface{}{
	(*Page)(nil),           // 0: erda.core.hepa.Page
	(*PageResult)(nil),     // 1: erda.core.hepa.PageResult
	(*NewPageResult)(nil),  // 2: erda.core.hepa.NewPageResult
	(*structpb.Value)(nil), // 3: google.protobuf.Value
}
var file_hepa_proto_depIdxs = []int32{
	3, // 0: erda.core.hepa.PageResult.result:type_name -> google.protobuf.Value
	0, // 1: erda.core.hepa.PageResult.page:type_name -> erda.core.hepa.Page
	3, // 2: erda.core.hepa.NewPageResult.list:type_name -> google.protobuf.Value
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_hepa_proto_init() }
func file_hepa_proto_init() {
	if File_hepa_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hepa_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_hepa_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageResult); i {
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
		file_hepa_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPageResult); i {
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
			RawDescriptor: file_hepa_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hepa_proto_goTypes,
		DependencyIndexes: file_hepa_proto_depIdxs,
		MessageInfos:      file_hepa_proto_msgTypes,
	}.Build()
	File_hepa_proto = out.File
	file_hepa_proto_rawDesc = nil
	file_hepa_proto_goTypes = nil
	file_hepa_proto_depIdxs = nil
}
