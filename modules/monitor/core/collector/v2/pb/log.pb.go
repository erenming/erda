// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: log.proto

package pb

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Log struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Source               string            `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Stream               string            `protobuf:"bytes,3,opt,name=stream,proto3" json:"stream,omitempty"`
	Offset               int64             `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Timestamp            int64             `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Content              string            `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Tags                 map[string]string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels               map[string]string `protobuf:"bytes,8,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{0}
}
func (m *Log) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Log.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return m.Size()
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Log) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Log) GetStream() string {
	if m != nil {
		return m.Stream
	}
	return ""
}

func (m *Log) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Log) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Log) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Log) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Log) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type LogBatch struct {
	Logs                 []*Log   `protobuf:"bytes,1,rep,name=Logs,proto3" json:"Logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogBatch) Reset()         { *m = LogBatch{} }
func (m *LogBatch) String() string { return proto.CompactTextString(m) }
func (*LogBatch) ProtoMessage()    {}
func (*LogBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{1}
}
func (m *LogBatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LogBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LogBatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LogBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogBatch.Merge(m, src)
}
func (m *LogBatch) XXX_Size() int {
	return m.Size()
}
func (m *LogBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_LogBatch.DiscardUnknown(m)
}

var xxx_messageInfo_LogBatch proto.InternalMessageInfo

func (m *LogBatch) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func init() {
	proto.RegisterType((*Log)(nil), "pb.Log")
	proto.RegisterMapType((map[string]string)(nil), "pb.Log.LabelsEntry")
	proto.RegisterMapType((map[string]string)(nil), "pb.Log.TagsEntry")
	proto.RegisterType((*LogBatch)(nil), "pb.LogBatch")
}

func init() { proto.RegisterFile("log.proto", fileDescriptor_a153da538f858886) }

var fileDescriptor_a153da538f858886 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0xd9, 0xa4, 0x4d, 0x9b, 0xa9, 0x88, 0xae, 0x22, 0x4b, 0x95, 0x52, 0x0a, 0x62, 0x41,
	0x4c, 0x41, 0x0f, 0xfe, 0x39, 0x16, 0xbc, 0xe5, 0x14, 0x7c, 0x81, 0x4d, 0xba, 0xdd, 0x06, 0x93,
	0x4e, 0xc8, 0x6e, 0x84, 0xbe, 0xa1, 0x47, 0x5f, 0x40, 0x90, 0x3e, 0x89, 0x74, 0xba, 0x51, 0xaf,
	0xde, 0xf6, 0xfb, 0xe6, 0x37, 0xc3, 0x0c, 0x0b, 0x61, 0x81, 0x3a, 0xaa, 0x6a, 0xb4, 0xc8, 0xbd,
	0x2a, 0x1d, 0xde, 0xe8, 0xdc, 0xae, 0x9a, 0x34, 0xca, 0xb0, 0x9c, 0x69, 0xd4, 0x38, 0xa3, 0x52,
	0xda, 0x2c, 0x89, 0x08, 0xe8, 0xb5, 0x6f, 0x99, 0x7c, 0x7a, 0xe0, 0xc7, 0xa8, 0xf9, 0x21, 0x78,
	0xf9, 0x42, 0xb0, 0x31, 0x9b, 0x86, 0x89, 0x97, 0x2f, 0xf8, 0x19, 0x04, 0x06, 0x9b, 0x3a, 0x53,
	0xc2, 0x23, 0xe7, 0x88, 0xbc, 0xad, 0x95, 0x2c, 0x85, 0xef, 0x3c, 0xd1, 0xce, 0xe3, 0x72, 0x69,
	0x94, 0x15, 0x9d, 0x31, 0x9b, 0xfa, 0x89, 0x23, 0x7e, 0x01, 0xa1, 0xcd, 0x4b, 0x65, 0xac, 0x2c,
	0x2b, 0xd1, 0xa5, 0xd2, 0xaf, 0xe0, 0x02, 0x7a, 0x19, 0xae, 0xad, 0x5a, 0x5b, 0x11, 0xd0, 0xb8,
	0x16, 0xf9, 0x25, 0x74, 0xac, 0xd4, 0x46, 0xf4, 0xc6, 0xfe, 0x74, 0x70, 0x7b, 0x1c, 0x55, 0x69,
	0x14, 0xa3, 0x8e, 0x5e, 0xa4, 0x36, 0xcf, 0x6b, 0x5b, 0x6f, 0x12, 0x2a, 0xf3, 0x6b, 0x08, 0x0a,
	0x99, 0xaa, 0xc2, 0x88, 0x3e, 0x05, 0x4f, 0xda, 0x60, 0x4c, 0x76, 0x1f, 0x75, 0x91, 0xe1, 0x3d,
	0x84, 0x3f, 0xfd, 0xfc, 0x08, 0xfc, 0x57, 0xb5, 0x71, 0x17, 0xef, 0x9e, 0xfc, 0x14, 0xba, 0x6f,
	0xb2, 0x68, 0xda, 0x8b, 0xf7, 0xf0, 0xe4, 0x3d, 0xb0, 0xe1, 0x23, 0x0c, 0xfe, 0xcc, 0xfb, 0x4f,
	0xeb, 0xe4, 0x0a, 0xfa, 0x31, 0xea, 0xb9, 0xb4, 0xd9, 0x8a, 0x9f, 0x43, 0x27, 0x46, 0x6d, 0x04,
	0xa3, 0x55, 0x7b, 0x6e, 0xd5, 0x84, 0xe4, 0xfc, 0xe0, 0x7d, 0x3b, 0x62, 0x1f, 0xdb, 0x11, 0xfb,
	0xda, 0x8e, 0x58, 0x1a, 0xd0, 0xef, 0xdc, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x93, 0x7e, 0xf1,
	0xb7, 0xdd, 0x01, 0x00, 0x00,
}

func (m *Log) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Log) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Log) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Labels) > 0 {
		for k := range m.Labels {
			v := m.Labels[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintLog(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintLog(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintLog(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Tags) > 0 {
		for k := range m.Tags {
			v := m.Tags[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintLog(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintLog(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintLog(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintLog(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x32
	}
	if m.Timestamp != 0 {
		i = encodeVarintLog(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x28
	}
	if m.Offset != 0 {
		i = encodeVarintLog(dAtA, i, uint64(m.Offset))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Stream) > 0 {
		i -= len(m.Stream)
		copy(dAtA[i:], m.Stream)
		i = encodeVarintLog(dAtA, i, uint64(len(m.Stream)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Source) > 0 {
		i -= len(m.Source)
		copy(dAtA[i:], m.Source)
		i = encodeVarintLog(dAtA, i, uint64(len(m.Source)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintLog(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LogBatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogBatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LogBatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Logs) > 0 {
		for iNdEx := len(m.Logs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Logs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLog(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintLog(dAtA []byte, offset int, v uint64) int {
	offset -= sovLog(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Log) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	l = len(m.Source)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	l = len(m.Stream)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	if m.Offset != 0 {
		n += 1 + sovLog(uint64(m.Offset))
	}
	if m.Timestamp != 0 {
		n += 1 + sovLog(uint64(m.Timestamp))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	if len(m.Tags) > 0 {
		for k, v := range m.Tags {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovLog(uint64(len(k))) + 1 + len(v) + sovLog(uint64(len(v)))
			n += mapEntrySize + 1 + sovLog(uint64(mapEntrySize))
		}
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovLog(uint64(len(k))) + 1 + len(v) + sovLog(uint64(len(v)))
			n += mapEntrySize + 1 + sovLog(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LogBatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Logs) > 0 {
		for _, e := range m.Logs {
			l = e.Size()
			n += 1 + l + sovLog(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLog(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLog(x uint64) (n int) {
	return sovLog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Log) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Log: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Log: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stream", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stream = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tags == nil {
				m.Tags = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLog
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLog
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthLog
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthLog
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLog
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthLog
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthLog
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipLog(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthLog
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Tags[mapkey] = mapvalue
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLog
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLog
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthLog
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthLog
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLog
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthLog
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthLog
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipLog(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthLog
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LogBatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogBatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogBatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Logs = append(m.Logs, &Log{})
			if err := m.Logs[len(m.Logs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLog(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLog
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLog
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLog
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLog
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLog
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLog
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLog        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLog          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLog = fmt.Errorf("proto: unexpected end of group")
)
