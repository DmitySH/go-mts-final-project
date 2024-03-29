// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/location.proto

package location

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetDriversRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat    float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng    float64 `protobuf:"fixed64,2,opt,name=lng,proto3" json:"lng,omitempty"`
	Radius float64 `protobuf:"fixed64,3,opt,name=radius,proto3" json:"radius,omitempty"`
}

func (x *GetDriversRequest) Reset() {
	*x = GetDriversRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDriversRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDriversRequest) ProtoMessage() {}

func (x *GetDriversRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDriversRequest.ProtoReflect.Descriptor instead.
func (*GetDriversRequest) Descriptor() ([]byte, []int) {
	return file_api_location_proto_rawDescGZIP(), []int{0}
}

func (x *GetDriversRequest) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *GetDriversRequest) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

func (x *GetDriversRequest) GetRadius() float64 {
	if x != nil {
		return x.Radius
	}
	return 0
}

type GetDriversResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Drivers []*Driver `protobuf:"bytes,1,rep,name=drivers,proto3" json:"drivers,omitempty"`
}

func (x *GetDriversResponse) Reset() {
	*x = GetDriversResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDriversResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDriversResponse) ProtoMessage() {}

func (x *GetDriversResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDriversResponse.ProtoReflect.Descriptor instead.
func (*GetDriversResponse) Descriptor() ([]byte, []int) {
	return file_api_location_proto_rawDescGZIP(), []int{1}
}

func (x *GetDriversResponse) GetDrivers() []*Driver {
	if x != nil {
		return x.Drivers
	}
	return nil
}

type UpdateDriverLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId string         `protobuf:"bytes,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	Location *LatLngLiteral `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *UpdateDriverLocationRequest) Reset() {
	*x = UpdateDriverLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_location_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDriverLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDriverLocationRequest) ProtoMessage() {}

func (x *UpdateDriverLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_location_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDriverLocationRequest.ProtoReflect.Descriptor instead.
func (*UpdateDriverLocationRequest) Descriptor() ([]byte, []int) {
	return file_api_location_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateDriverLocationRequest) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *UpdateDriverLocationRequest) GetLocation() *LatLngLiteral {
	if x != nil {
		return x.Location
	}
	return nil
}

type UpdateDriverLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateDriverLocationResponse) Reset() {
	*x = UpdateDriverLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_location_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDriverLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDriverLocationResponse) ProtoMessage() {}

func (x *UpdateDriverLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_location_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDriverLocationResponse.ProtoReflect.Descriptor instead.
func (*UpdateDriverLocationResponse) Descriptor() ([]byte, []int) {
	return file_api_location_proto_rawDescGZIP(), []int{3}
}

type Driver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Auto string  `protobuf:"bytes,3,opt,name=auto,proto3" json:"auto,omitempty"`
	Lat  float64 `protobuf:"fixed64,4,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng  float64 `protobuf:"fixed64,5,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *Driver) Reset() {
	*x = Driver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_location_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Driver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Driver) ProtoMessage() {}

func (x *Driver) ProtoReflect() protoreflect.Message {
	mi := &file_api_location_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Driver.ProtoReflect.Descriptor instead.
func (*Driver) Descriptor() ([]byte, []int) {
	return file_api_location_proto_rawDescGZIP(), []int{4}
}

func (x *Driver) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Driver) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Driver) GetAuto() string {
	if x != nil {
		return x.Auto
	}
	return ""
}

func (x *Driver) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Driver) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

type LatLngLiteral struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng float64 `protobuf:"fixed64,2,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *LatLngLiteral) Reset() {
	*x = LatLngLiteral{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_location_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatLngLiteral) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatLngLiteral) ProtoMessage() {}

func (x *LatLngLiteral) ProtoReflect() protoreflect.Message {
	mi := &file_api_location_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatLngLiteral.ProtoReflect.Descriptor instead.
func (*LatLngLiteral) Descriptor() ([]byte, []int) {
	return file_api_location_proto_rawDescGZIP(), []int{5}
}

func (x *LatLngLiteral) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *LatLngLiteral) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

var File_api_location_proto protoreflect.FileDescriptor

var file_api_location_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x61, 0x78, 0x69, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61,
	0x64, 0x69, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x64, 0x69,
	0x75, 0x73, 0x22, 0x45, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x61, 0x78, 0x69,
	0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x52, 0x07, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x22, 0x74, 0x0a, 0x1b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x61, 0x78, 0x69, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x4c, 0x69,
	0x74, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x1e, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x64, 0x0a, 0x06, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x75, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x75, 0x74,
	0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x6c, 0x6e, 0x67, 0x22, 0x33, 0x0a, 0x0d, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x4c,
	0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x32, 0x93, 0x02, 0x0a, 0x08, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x12, 0x20, 0x2e, 0x74, 0x61, 0x78, 0x69, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x61, 0x78, 0x69, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x62, 0x07, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x12, 0x08, 0x2f, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x12, 0x98, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a,
	0x2e, 0x74, 0x61, 0x78, 0x69, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x74, 0x61, 0x78,
	0x69, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a,
	0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x7b, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x73, 0x65, 0x2d, 0x6d, 0x74, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x64, 0x61, 0x73, 0x68, 0x61, 0x67,
	0x61, 0x72, 0x6f, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x61, 0x78, 0x69, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_location_proto_rawDescOnce sync.Once
	file_api_location_proto_rawDescData = file_api_location_proto_rawDesc
)

func file_api_location_proto_rawDescGZIP() []byte {
	file_api_location_proto_rawDescOnce.Do(func() {
		file_api_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_location_proto_rawDescData)
	})
	return file_api_location_proto_rawDescData
}

var file_api_location_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_location_proto_goTypes = []interface{}{
	(*GetDriversRequest)(nil),            // 0: taxi.location.GetDriversRequest
	(*GetDriversResponse)(nil),           // 1: taxi.location.GetDriversResponse
	(*UpdateDriverLocationRequest)(nil),  // 2: taxi.location.UpdateDriverLocationRequest
	(*UpdateDriverLocationResponse)(nil), // 3: taxi.location.UpdateDriverLocationResponse
	(*Driver)(nil),                       // 4: taxi.location.Driver
	(*LatLngLiteral)(nil),                // 5: taxi.location.LatLngLiteral
}
var file_api_location_proto_depIdxs = []int32{
	4, // 0: taxi.location.GetDriversResponse.drivers:type_name -> taxi.location.Driver
	5, // 1: taxi.location.UpdateDriverLocationRequest.location:type_name -> taxi.location.LatLngLiteral
	0, // 2: taxi.location.Location.GetDrivers:input_type -> taxi.location.GetDriversRequest
	2, // 3: taxi.location.Location.UpdateDriverLocation:input_type -> taxi.location.UpdateDriverLocationRequest
	1, // 4: taxi.location.Location.GetDrivers:output_type -> taxi.location.GetDriversResponse
	3, // 5: taxi.location.Location.UpdateDriverLocation:output_type -> taxi.location.UpdateDriverLocationResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_location_proto_init() }
func file_api_location_proto_init() {
	if File_api_location_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDriversRequest); i {
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
		file_api_location_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDriversResponse); i {
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
		file_api_location_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDriverLocationRequest); i {
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
		file_api_location_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDriverLocationResponse); i {
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
		file_api_location_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Driver); i {
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
		file_api_location_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatLngLiteral); i {
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
			RawDescriptor: file_api_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_location_proto_goTypes,
		DependencyIndexes: file_api_location_proto_depIdxs,
		MessageInfos:      file_api_location_proto_msgTypes,
	}.Build()
	File_api_location_proto = out.File
	file_api_location_proto_rawDesc = nil
	file_api_location_proto_goTypes = nil
	file_api_location_proto_depIdxs = nil
}
