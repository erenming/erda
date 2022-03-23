// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: guide.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*ListGuideRequest)(nil)
var _ json.Unmarshaler = (*ListGuideRequest)(nil)
var _ json.Marshaler = (*ListGuideResponse)(nil)
var _ json.Unmarshaler = (*ListGuideResponse)(nil)
var _ json.Marshaler = (*Guide)(nil)
var _ json.Unmarshaler = (*Guide)(nil)
var _ json.Marshaler = (*GittarPushPayloadEvent)(nil)
var _ json.Unmarshaler = (*GittarPushPayloadEvent)(nil)
var _ json.Marshaler = (*Content)(nil)
var _ json.Unmarshaler = (*Content)(nil)
var _ json.Marshaler = (*Pusher)(nil)
var _ json.Unmarshaler = (*Pusher)(nil)
var _ json.Marshaler = (*CreateGuideResponse)(nil)
var _ json.Unmarshaler = (*CreateGuideResponse)(nil)
var _ json.Marshaler = (*ProcessGuideRequest)(nil)
var _ json.Unmarshaler = (*ProcessGuideRequest)(nil)
var _ json.Marshaler = (*ProcessGuideResponse)(nil)
var _ json.Unmarshaler = (*ProcessGuideResponse)(nil)

// ListGuideRequest implement json.Marshaler.
func (m *ListGuideRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListGuideRequest implement json.Marshaler.
func (m *ListGuideRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListGuideResponse implement json.Marshaler.
func (m *ListGuideResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListGuideResponse implement json.Marshaler.
func (m *ListGuideResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Guide implement json.Marshaler.
func (m *Guide) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Guide implement json.Marshaler.
func (m *Guide) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GittarPushPayloadEvent implement json.Marshaler.
func (m *GittarPushPayloadEvent) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GittarPushPayloadEvent implement json.Marshaler.
func (m *GittarPushPayloadEvent) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Content implement json.Marshaler.
func (m *Content) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Content implement json.Marshaler.
func (m *Content) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Pusher implement json.Marshaler.
func (m *Pusher) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Pusher implement json.Marshaler.
func (m *Pusher) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateGuideResponse implement json.Marshaler.
func (m *CreateGuideResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateGuideResponse implement json.Marshaler.
func (m *CreateGuideResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ProcessGuideRequest implement json.Marshaler.
func (m *ProcessGuideRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ProcessGuideRequest implement json.Marshaler.
func (m *ProcessGuideRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ProcessGuideResponse implement json.Marshaler.
func (m *ProcessGuideResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ProcessGuideResponse implement json.Marshaler.
func (m *ProcessGuideResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
