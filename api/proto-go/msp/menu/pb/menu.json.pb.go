// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: menu.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*GetMenuRequest)(nil)
var _ json.Unmarshaler = (*GetMenuRequest)(nil)
var _ json.Marshaler = (*GetMenuResponse)(nil)
var _ json.Unmarshaler = (*GetMenuResponse)(nil)
var _ json.Marshaler = (*MenuItem)(nil)
var _ json.Unmarshaler = (*MenuItem)(nil)
var _ json.Marshaler = (*GetSettingRequest)(nil)
var _ json.Unmarshaler = (*GetSettingRequest)(nil)
var _ json.Marshaler = (*GetSettingResponse)(nil)
var _ json.Unmarshaler = (*GetSettingResponse)(nil)
var _ json.Marshaler = (*EngineSetting)(nil)
var _ json.Unmarshaler = (*EngineSetting)(nil)

// GetMenuRequest implement json.Marshaler.
func (m *GetMenuRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetMenuRequest implement json.Marshaler.
func (m *GetMenuRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetMenuResponse implement json.Marshaler.
func (m *GetMenuResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetMenuResponse implement json.Marshaler.
func (m *GetMenuResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// MenuItem implement json.Marshaler.
func (m *MenuItem) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// MenuItem implement json.Marshaler.
func (m *MenuItem) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetSettingRequest implement json.Marshaler.
func (m *GetSettingRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetSettingRequest implement json.Marshaler.
func (m *GetSettingRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetSettingResponse implement json.Marshaler.
func (m *GetSettingResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetSettingResponse implement json.Marshaler.
func (m *GetSettingResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// EngineSetting implement json.Marshaler.
func (m *EngineSetting) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// EngineSetting implement json.Marshaler.
func (m *EngineSetting) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
