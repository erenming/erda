// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: notify_group.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*CreateNotifyGroupRequest)(nil)
var _ json.Unmarshaler = (*CreateNotifyGroupRequest)(nil)
var _ json.Marshaler = (*NotifyTarget)(nil)
var _ json.Unmarshaler = (*NotifyTarget)(nil)
var _ json.Marshaler = (*Target)(nil)
var _ json.Unmarshaler = (*Target)(nil)
var _ json.Marshaler = (*CreateNotifyGroupResponse)(nil)
var _ json.Unmarshaler = (*CreateNotifyGroupResponse)(nil)
var _ json.Marshaler = (*NotifyGroup)(nil)
var _ json.Unmarshaler = (*NotifyGroup)(nil)
var _ json.Marshaler = (*QueryNotifyGroupRequest)(nil)
var _ json.Unmarshaler = (*QueryNotifyGroupRequest)(nil)
var _ json.Marshaler = (*QueryNotifyGroupResponse)(nil)
var _ json.Unmarshaler = (*QueryNotifyGroupResponse)(nil)
var _ json.Marshaler = (*QueryNotifyGroupData)(nil)
var _ json.Unmarshaler = (*QueryNotifyGroupData)(nil)
var _ json.Marshaler = (*GetNotifyGroupRequest)(nil)
var _ json.Unmarshaler = (*GetNotifyGroupRequest)(nil)
var _ json.Marshaler = (*GetNotifyGroupResponse)(nil)
var _ json.Unmarshaler = (*GetNotifyGroupResponse)(nil)
var _ json.Marshaler = (*UpdateNotifyGroupRequest)(nil)
var _ json.Unmarshaler = (*UpdateNotifyGroupRequest)(nil)
var _ json.Marshaler = (*UpdateNotifyGroupResponse)(nil)
var _ json.Unmarshaler = (*UpdateNotifyGroupResponse)(nil)
var _ json.Marshaler = (*GetNotifyGroupDetailRequest)(nil)
var _ json.Unmarshaler = (*GetNotifyGroupDetailRequest)(nil)
var _ json.Marshaler = (*GetNotifyGroupDetailResponse)(nil)
var _ json.Unmarshaler = (*GetNotifyGroupDetailResponse)(nil)
var _ json.Marshaler = (*NotifyGroupDetail)(nil)
var _ json.Unmarshaler = (*NotifyGroupDetail)(nil)
var _ json.Marshaler = (*NotifyUser)(nil)
var _ json.Unmarshaler = (*NotifyUser)(nil)
var _ json.Marshaler = (*DeleteNotifyGroupRequest)(nil)
var _ json.Unmarshaler = (*DeleteNotifyGroupRequest)(nil)
var _ json.Marshaler = (*DeleteNotifyGroupResponse)(nil)
var _ json.Unmarshaler = (*DeleteNotifyGroupResponse)(nil)
var _ json.Marshaler = (*BatchGetNotifyGroupRequest)(nil)
var _ json.Unmarshaler = (*BatchGetNotifyGroupRequest)(nil)
var _ json.Marshaler = (*BatchGetNotifyGroupResponse)(nil)
var _ json.Unmarshaler = (*BatchGetNotifyGroupResponse)(nil)

// CreateNotifyGroupRequest implement json.Marshaler.
func (m *CreateNotifyGroupRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateNotifyGroupRequest implement json.Marshaler.
func (m *CreateNotifyGroupRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NotifyTarget implement json.Marshaler.
func (m *NotifyTarget) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NotifyTarget implement json.Marshaler.
func (m *NotifyTarget) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Target implement json.Marshaler.
func (m *Target) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Target implement json.Marshaler.
func (m *Target) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateNotifyGroupResponse implement json.Marshaler.
func (m *CreateNotifyGroupResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateNotifyGroupResponse implement json.Marshaler.
func (m *CreateNotifyGroupResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NotifyGroup implement json.Marshaler.
func (m *NotifyGroup) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NotifyGroup implement json.Marshaler.
func (m *NotifyGroup) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryNotifyGroupRequest implement json.Marshaler.
func (m *QueryNotifyGroupRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryNotifyGroupRequest implement json.Marshaler.
func (m *QueryNotifyGroupRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryNotifyGroupResponse implement json.Marshaler.
func (m *QueryNotifyGroupResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryNotifyGroupResponse implement json.Marshaler.
func (m *QueryNotifyGroupResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryNotifyGroupData implement json.Marshaler.
func (m *QueryNotifyGroupData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryNotifyGroupData implement json.Marshaler.
func (m *QueryNotifyGroupData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNotifyGroupRequest implement json.Marshaler.
func (m *GetNotifyGroupRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNotifyGroupRequest implement json.Marshaler.
func (m *GetNotifyGroupRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNotifyGroupResponse implement json.Marshaler.
func (m *GetNotifyGroupResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNotifyGroupResponse implement json.Marshaler.
func (m *GetNotifyGroupResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateNotifyGroupRequest implement json.Marshaler.
func (m *UpdateNotifyGroupRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateNotifyGroupRequest implement json.Marshaler.
func (m *UpdateNotifyGroupRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateNotifyGroupResponse implement json.Marshaler.
func (m *UpdateNotifyGroupResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateNotifyGroupResponse implement json.Marshaler.
func (m *UpdateNotifyGroupResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNotifyGroupDetailRequest implement json.Marshaler.
func (m *GetNotifyGroupDetailRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNotifyGroupDetailRequest implement json.Marshaler.
func (m *GetNotifyGroupDetailRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNotifyGroupDetailResponse implement json.Marshaler.
func (m *GetNotifyGroupDetailResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNotifyGroupDetailResponse implement json.Marshaler.
func (m *GetNotifyGroupDetailResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NotifyGroupDetail implement json.Marshaler.
func (m *NotifyGroupDetail) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NotifyGroupDetail implement json.Marshaler.
func (m *NotifyGroupDetail) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NotifyUser implement json.Marshaler.
func (m *NotifyUser) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NotifyUser implement json.Marshaler.
func (m *NotifyUser) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteNotifyGroupRequest implement json.Marshaler.
func (m *DeleteNotifyGroupRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteNotifyGroupRequest implement json.Marshaler.
func (m *DeleteNotifyGroupRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteNotifyGroupResponse implement json.Marshaler.
func (m *DeleteNotifyGroupResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteNotifyGroupResponse implement json.Marshaler.
func (m *DeleteNotifyGroupResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// BatchGetNotifyGroupRequest implement json.Marshaler.
func (m *BatchGetNotifyGroupRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// BatchGetNotifyGroupRequest implement json.Marshaler.
func (m *BatchGetNotifyGroupRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// BatchGetNotifyGroupResponse implement json.Marshaler.
func (m *BatchGetNotifyGroupResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// BatchGetNotifyGroupResponse implement json.Marshaler.
func (m *BatchGetNotifyGroupResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
