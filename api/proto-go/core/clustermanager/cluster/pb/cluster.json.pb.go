// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: cluster.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*ErrResponse)(nil)
var _ json.Unmarshaler = (*ErrResponse)(nil)
var _ json.Marshaler = (*ListClusterRequest)(nil)
var _ json.Unmarshaler = (*ListClusterRequest)(nil)
var _ json.Marshaler = (*ListClusterResponse)(nil)
var _ json.Unmarshaler = (*ListClusterResponse)(nil)
var _ json.Marshaler = (*GetClusterRequest)(nil)
var _ json.Unmarshaler = (*GetClusterRequest)(nil)
var _ json.Marshaler = (*GetClusterResponse)(nil)
var _ json.Unmarshaler = (*GetClusterResponse)(nil)
var _ json.Marshaler = (*UpdateClusterRequest)(nil)
var _ json.Unmarshaler = (*UpdateClusterRequest)(nil)
var _ json.Marshaler = (*UpdateClusterResponse)(nil)
var _ json.Unmarshaler = (*UpdateClusterResponse)(nil)
var _ json.Marshaler = (*CreateClusterRequest)(nil)
var _ json.Unmarshaler = (*CreateClusterRequest)(nil)
var _ json.Marshaler = (*CreateClusterResponse)(nil)
var _ json.Unmarshaler = (*CreateClusterResponse)(nil)
var _ json.Marshaler = (*DeleteClusterRequest)(nil)
var _ json.Unmarshaler = (*DeleteClusterRequest)(nil)
var _ json.Marshaler = (*DeleteClusterResponse)(nil)
var _ json.Unmarshaler = (*DeleteClusterResponse)(nil)
var _ json.Marshaler = (*PatchClusterRequest)(nil)
var _ json.Unmarshaler = (*PatchClusterRequest)(nil)
var _ json.Marshaler = (*PatchClusterResponse)(nil)
var _ json.Unmarshaler = (*PatchClusterResponse)(nil)
var _ json.Marshaler = (*ClusterInfo)(nil)
var _ json.Unmarshaler = (*ClusterInfo)(nil)
var _ json.Marshaler = (*ClusterSchedConfig)(nil)
var _ json.Unmarshaler = (*ClusterSchedConfig)(nil)
var _ json.Marshaler = (*OpsConfig)(nil)
var _ json.Unmarshaler = (*OpsConfig)(nil)
var _ json.Marshaler = (*SysConf)(nil)
var _ json.Unmarshaler = (*SysConf)(nil)
var _ json.Marshaler = (*Cluster)(nil)
var _ json.Unmarshaler = (*Cluster)(nil)
var _ json.Marshaler = (*SSH)(nil)
var _ json.Unmarshaler = (*SSH)(nil)
var _ json.Marshaler = (*FPS)(nil)
var _ json.Unmarshaler = (*FPS)(nil)
var _ json.Marshaler = (*Storage)(nil)
var _ json.Unmarshaler = (*Storage)(nil)
var _ json.Marshaler = (*Gluster)(nil)
var _ json.Unmarshaler = (*Gluster)(nil)
var _ json.Marshaler = (*Docker)(nil)
var _ json.Unmarshaler = (*Docker)(nil)
var _ json.Marshaler = (*Node)(nil)
var _ json.Unmarshaler = (*Node)(nil)
var _ json.Marshaler = (*MySQL)(nil)
var _ json.Unmarshaler = (*MySQL)(nil)
var _ json.Marshaler = (*OpenVPN)(nil)
var _ json.Unmarshaler = (*OpenVPN)(nil)
var _ json.Marshaler = (*Platform)(nil)
var _ json.Unmarshaler = (*Platform)(nil)
var _ json.Marshaler = (*ManageConfig)(nil)
var _ json.Unmarshaler = (*ManageConfig)(nil)

// ErrResponse implement json.Marshaler.
func (m *ErrResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ErrResponse implement json.Marshaler.
func (m *ErrResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListClusterRequest implement json.Marshaler.
func (m *ListClusterRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListClusterRequest implement json.Marshaler.
func (m *ListClusterRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListClusterResponse implement json.Marshaler.
func (m *ListClusterResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListClusterResponse implement json.Marshaler.
func (m *ListClusterResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetClusterRequest implement json.Marshaler.
func (m *GetClusterRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetClusterRequest implement json.Marshaler.
func (m *GetClusterRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetClusterResponse implement json.Marshaler.
func (m *GetClusterResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetClusterResponse implement json.Marshaler.
func (m *GetClusterResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateClusterRequest implement json.Marshaler.
func (m *UpdateClusterRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateClusterRequest implement json.Marshaler.
func (m *UpdateClusterRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateClusterResponse implement json.Marshaler.
func (m *UpdateClusterResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateClusterResponse implement json.Marshaler.
func (m *UpdateClusterResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateClusterRequest implement json.Marshaler.
func (m *CreateClusterRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateClusterRequest implement json.Marshaler.
func (m *CreateClusterRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateClusterResponse implement json.Marshaler.
func (m *CreateClusterResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateClusterResponse implement json.Marshaler.
func (m *CreateClusterResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteClusterRequest implement json.Marshaler.
func (m *DeleteClusterRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteClusterRequest implement json.Marshaler.
func (m *DeleteClusterRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteClusterResponse implement json.Marshaler.
func (m *DeleteClusterResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteClusterResponse implement json.Marshaler.
func (m *DeleteClusterResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// PatchClusterRequest implement json.Marshaler.
func (m *PatchClusterRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// PatchClusterRequest implement json.Marshaler.
func (m *PatchClusterRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// PatchClusterResponse implement json.Marshaler.
func (m *PatchClusterResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// PatchClusterResponse implement json.Marshaler.
func (m *PatchClusterResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ClusterInfo implement json.Marshaler.
func (m *ClusterInfo) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ClusterInfo implement json.Marshaler.
func (m *ClusterInfo) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ClusterSchedConfig implement json.Marshaler.
func (m *ClusterSchedConfig) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ClusterSchedConfig implement json.Marshaler.
func (m *ClusterSchedConfig) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OpsConfig implement json.Marshaler.
func (m *OpsConfig) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OpsConfig implement json.Marshaler.
func (m *OpsConfig) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// SysConf implement json.Marshaler.
func (m *SysConf) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// SysConf implement json.Marshaler.
func (m *SysConf) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Cluster implement json.Marshaler.
func (m *Cluster) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Cluster implement json.Marshaler.
func (m *Cluster) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// SSH implement json.Marshaler.
func (m *SSH) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// SSH implement json.Marshaler.
func (m *SSH) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// FPS implement json.Marshaler.
func (m *FPS) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// FPS implement json.Marshaler.
func (m *FPS) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Storage implement json.Marshaler.
func (m *Storage) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Storage implement json.Marshaler.
func (m *Storage) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Gluster implement json.Marshaler.
func (m *Gluster) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Gluster implement json.Marshaler.
func (m *Gluster) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Docker implement json.Marshaler.
func (m *Docker) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Docker implement json.Marshaler.
func (m *Docker) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Node implement json.Marshaler.
func (m *Node) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Node implement json.Marshaler.
func (m *Node) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// MySQL implement json.Marshaler.
func (m *MySQL) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// MySQL implement json.Marshaler.
func (m *MySQL) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OpenVPN implement json.Marshaler.
func (m *OpenVPN) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OpenVPN implement json.Marshaler.
func (m *OpenVPN) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Platform implement json.Marshaler.
func (m *Platform) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Platform implement json.Marshaler.
func (m *Platform) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ManageConfig implement json.Marshaler.
func (m *ManageConfig) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ManageConfig implement json.Marshaler.
func (m *ManageConfig) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
