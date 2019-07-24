// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bitbucket-repository-management-service/dm/proto-files/domain/todo.proto

package domain

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Taks we have to do
type ToDo struct {
	// Unique integer identifier of the todo task
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Title of the task
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Detail description of the todo task
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Date and time to remind the todo task
	Reminder             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=reminder,proto3" json:"reminder,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ToDo) Reset()         { *m = ToDo{} }
func (m *ToDo) String() string { return proto.CompactTextString(m) }
func (*ToDo) ProtoMessage()    {}
func (*ToDo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2852dedc1961400, []int{0}
}

func (m *ToDo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToDo.Unmarshal(m, b)
}
func (m *ToDo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToDo.Marshal(b, m, deterministic)
}
func (m *ToDo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToDo.Merge(m, src)
}
func (m *ToDo) XXX_Size() int {
	return xxx_messageInfo_ToDo.Size(m)
}
func (m *ToDo) XXX_DiscardUnknown() {
	xxx_messageInfo_ToDo.DiscardUnknown(m)
}

var xxx_messageInfo_ToDo proto.InternalMessageInfo

func (m *ToDo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ToDo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ToDo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ToDo) GetReminder() *timestamp.Timestamp {
	if m != nil {
		return m.Reminder
	}
	return nil
}

// Request data to create new todo task
type CreateRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Task entity to add
	ToDo                 *ToDo    `protobuf:"bytes,2,opt,name=toDo,proto3" json:"toDo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2852dedc1961400, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetToDo() *ToDo {
	if m != nil {
		return m.ToDo
	}
	return nil
}

// Contains data of created todo task
type CreateResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// ID of created task
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2852dedc1961400, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Request data to read todo task
type ReadRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Unique integer identifier of the todo task
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2852dedc1961400, []int{3}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Contains todo task data specified in by ID request
type ReadResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Task entity read by ID
	ToDo                 *ToDo    `protobuf:"bytes,2,opt,name=toDo,proto3" json:"toDo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2852dedc1961400, []int{4}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadResponse) GetToDo() *ToDo {
	if m != nil {
		return m.ToDo
	}
	return nil
}

// Request data to update todo task
type UpdateRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Task entity to update
	ToDo                 *ToDo    `protobuf:"bytes,2,opt,name=toDo,proto3" json:"toDo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2852dedc1961400, []int{5}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateRequest) GetToDo() *ToDo {
	if m != nil {
		return m.ToDo
	}
	return nil
}

// Contains status of update operation
type UpdateResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Contains number of entities have beed updated
	// Equals 1 in case of succesfull update
	Updated              int64    `protobuf:"varint,2,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2852dedc1961400, []int{6}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateResponse) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

// Request data to delete todo task
type DeleteRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Unique integer identifier of the todo task to delete
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2852dedc1961400, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Contains status of delete operation
type DeleteResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Contains number of entities have beed deleted
	// Equals 1 in case of succesfull delete
	Deleted              int64    `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2852dedc1961400, []int{8}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

// Request data to read all todo task
type ReadAllRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllRequest) Reset()         { *m = ReadAllRequest{} }
func (m *ReadAllRequest) String() string { return proto.CompactTextString(m) }
func (*ReadAllRequest) ProtoMessage()    {}
func (*ReadAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2852dedc1961400, []int{9}
}

func (m *ReadAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllRequest.Unmarshal(m, b)
}
func (m *ReadAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllRequest.Marshal(b, m, deterministic)
}
func (m *ReadAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllRequest.Merge(m, src)
}
func (m *ReadAllRequest) XXX_Size() int {
	return xxx_messageInfo_ReadAllRequest.Size(m)
}
func (m *ReadAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllRequest proto.InternalMessageInfo

func (m *ReadAllRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

// Contains list of all todo tasks
type ReadAllResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// List of all todo tasks
	ToDos                []*ToDo  `protobuf:"bytes,2,rep,name=toDos,proto3" json:"toDos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllResponse) Reset()         { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAllResponse) ProtoMessage()    {}
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2852dedc1961400, []int{10}
}

func (m *ReadAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllResponse.Unmarshal(m, b)
}
func (m *ReadAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllResponse.Marshal(b, m, deterministic)
}
func (m *ReadAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllResponse.Merge(m, src)
}
func (m *ReadAllResponse) XXX_Size() int {
	return xxx_messageInfo_ReadAllResponse.Size(m)
}
func (m *ReadAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllResponse proto.InternalMessageInfo

func (m *ReadAllResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadAllResponse) GetToDos() []*ToDo {
	if m != nil {
		return m.ToDos
	}
	return nil
}

func init() {
	proto.RegisterType((*ToDo)(nil), "domain.ToDo")
	proto.RegisterType((*CreateRequest)(nil), "domain.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "domain.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "domain.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "domain.ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "domain.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "domain.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "domain.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "domain.DeleteResponse")
	proto.RegisterType((*ReadAllRequest)(nil), "domain.ReadAllRequest")
	proto.RegisterType((*ReadAllResponse)(nil), "domain.ReadAllResponse")
}

func init() {
	proto.RegisterFile("bitbucket-repository-management-service/dm/proto-files/domain/todo.proto", fileDescriptor_f2852dedc1961400)
}

var fileDescriptor_f2852dedc1961400 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x41, 0x8f, 0xd3, 0x30,
	0x10, 0x85, 0x95, 0xa4, 0xbb, 0xb0, 0x93, 0x6d, 0x40, 0x11, 0x87, 0x68, 0x2f, 0x44, 0x3e, 0xf5,
	0x92, 0x44, 0x14, 0xa9, 0xe2, 0xd0, 0x0b, 0x6d, 0x25, 0x38, 0x47, 0xe5, 0xc2, 0x2d, 0x89, 0xa7,
	0x91, 0x45, 0x12, 0x1b, 0xdb, 0x41, 0xe2, 0x0f, 0xf0, 0xbb, 0x51, 0xec, 0xba, 0x42, 0x85, 0x14,
	0x21, 0x6e, 0xf1, 0xcc, 0x9b, 0x79, 0xdf, 0xb3, 0x62, 0xf8, 0x58, 0x33, 0x5d, 0x8f, 0xcd, 0x17,
	0xd4, 0x99, 0x44, 0xc1, 0x15, 0xd3, 0x5c, 0x7e, 0xcf, 0xfa, 0x6a, 0xa8, 0x5a, 0xec, 0x71, 0xd0,
	0x99, 0x42, 0xf9, 0x8d, 0x35, 0x58, 0xd0, 0xbe, 0x10, 0x92, 0x6b, 0x9e, 0x9d, 0x58, 0x87, 0xaa,
	0xa0, 0xbc, 0xaf, 0xd8, 0x50, 0x68, 0x4e, 0x79, 0x6e, 0xea, 0xf1, 0xbd, 0x2d, 0x3d, 0xbd, 0x6e,
	0x39, 0x6f, 0x3b, 0xb4, 0xea, 0x7a, 0x3c, 0x15, 0x9a, 0xf5, 0xa8, 0x74, 0xd5, 0x0b, 0x2b, 0x24,
	0x3f, 0x3c, 0x58, 0x1c, 0xf9, 0x81, 0xc7, 0x11, 0xf8, 0x8c, 0x26, 0x5e, 0xea, 0xad, 0x82, 0xd2,
	0x67, 0x34, 0x7e, 0x05, 0x77, 0x9a, 0xe9, 0x0e, 0x13, 0x3f, 0xf5, 0x56, 0x0f, 0xa5, 0x3d, 0xc4,
	0x29, 0x84, 0x14, 0x55, 0x23, 0x99, 0xd0, 0x8c, 0x0f, 0x49, 0x60, 0x7a, 0xbf, 0x96, 0xe2, 0x0d,
	0x3c, 0x97, 0xd8, 0xb3, 0x81, 0xa2, 0x4c, 0x16, 0xa9, 0xb7, 0x0a, 0xd7, 0x4f, 0xb9, 0x85, 0xc8,
	0x1d, 0x44, 0x7e, 0x74, 0x10, 0xe5, 0x45, 0x4b, 0xf6, 0xb0, 0xdc, 0x4b, 0xac, 0x34, 0x96, 0xf8,
	0x75, 0x44, 0xa5, 0xe3, 0x97, 0x10, 0x54, 0x82, 0x19, 0xa2, 0x87, 0x72, 0xfa, 0x8c, 0x53, 0x58,
	0x68, 0x7e, 0xe0, 0x86, 0x28, 0x5c, 0x3f, 0xe6, 0x36, 0x63, 0x3e, 0xe1, 0x97, 0xa6, 0x43, 0xd6,
	0x10, 0xb9, 0x25, 0x4a, 0xf0, 0x41, 0xe1, 0x1f, 0xb6, 0xd8, 0xa0, 0xbe, 0x0b, 0x4a, 0x0a, 0x08,
	0x4b, 0xac, 0xe8, 0xbc, 0xed, 0xf5, 0xc0, 0x0e, 0x1e, 0xed, 0xc0, 0xac, 0xc5, 0xdf, 0x41, 0xf7,
	0xb0, 0xfc, 0x24, 0xe8, 0x7f, 0xa6, 0xdd, 0x42, 0xe4, 0x96, 0xcc, 0xa2, 0x24, 0xf0, 0x6c, 0x34,
	0x1a, 0x97, 0xc0, 0x1d, 0xc9, 0x1b, 0x58, 0x1e, 0xb0, 0xc3, 0x5b, 0x08, 0xd7, 0xc9, 0xb7, 0x10,
	0xb9, 0x91, 0x5b, 0x86, 0xd4, 0x68, 0x2e, 0x86, 0xe7, 0x23, 0x21, 0x10, 0x4d, 0xf7, 0xf6, 0xbe,
	0xeb, 0x66, 0x1d, 0xc9, 0x07, 0x78, 0x71, 0xd1, 0xcc, 0x5a, 0x10, 0xb8, 0x9b, 0xf2, 0xab, 0xc4,
	0x4f, 0x83, 0xdf, 0xae, 0xc6, 0xb6, 0x76, 0xef, 0x3e, 0x6f, 0xfe, 0xe1, 0x31, 0xb5, 0x52, 0x34,
	0xe7, 0x57, 0x54, 0xdf, 0x9b, 0xdf, 0xf4, 0xed, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xc2,
	0xc8, 0x4e, 0x8d, 0x03, 0x00, 0x00,
}
