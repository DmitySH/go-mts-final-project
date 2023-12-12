// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/driver.proto

package driver

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

type TripStatus int32

const (
	TripStatus_UNKNOWN       TripStatus = 0
	TripStatus_DRIVER_SEARCH TripStatus = 1
	TripStatus_DRIVER_FOUND  TripStatus = 2
	TripStatus_ON_POSITION   TripStatus = 3
	TripStatus_STARTED       TripStatus = 4
	TripStatus_ENDED         TripStatus = 5
	TripStatus_CANCELED      TripStatus = 6
)

// Enum value maps for TripStatus.
var (
	TripStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "DRIVER_SEARCH",
		2: "DRIVER_FOUND",
		3: "ON_POSITION",
		4: "STARTED",
		5: "ENDED",
		6: "CANCELED",
	}
	TripStatus_value = map[string]int32{
		"UNKNOWN":       0,
		"DRIVER_SEARCH": 1,
		"DRIVER_FOUND":  2,
		"ON_POSITION":   3,
		"STARTED":       4,
		"ENDED":         5,
		"CANCELED":      6,
	}
)

func (x TripStatus) Enum() *TripStatus {
	p := new(TripStatus)
	*p = x
	return p
}

func (x TripStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TripStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_driver_proto_enumTypes[0].Descriptor()
}

func (TripStatus) Type() protoreflect.EnumType {
	return &file_api_driver_proto_enumTypes[0]
}

func (x TripStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TripStatus.Descriptor instead.
func (TripStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_driver_proto_rawDescGZIP(), []int{0}
}

type GetTripsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTripsRequest) Reset() {
	*x = GetTripsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_driver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTripsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripsRequest) ProtoMessage() {}

func (x *GetTripsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_driver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripsRequest.ProtoReflect.Descriptor instead.
func (*GetTripsRequest) Descriptor() ([]byte, []int) {
	return file_api_driver_proto_rawDescGZIP(), []int{0}
}

type GetTripsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trips []*Trip `protobuf:"bytes,1,rep,name=trips,proto3" json:"trips,omitempty"`
}

func (x *GetTripsResponse) Reset() {
	*x = GetTripsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_driver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTripsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripsResponse) ProtoMessage() {}

func (x *GetTripsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_driver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripsResponse.ProtoReflect.Descriptor instead.
func (*GetTripsResponse) Descriptor() ([]byte, []int) {
	return file_api_driver_proto_rawDescGZIP(), []int{1}
}

func (x *GetTripsResponse) GetTrips() []*Trip {
	if x != nil {
		return x.Trips
	}
	return nil
}

type GetTripByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TripId string `protobuf:"bytes,1,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`
}

func (x *GetTripByIDRequest) Reset() {
	*x = GetTripByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_driver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTripByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripByIDRequest) ProtoMessage() {}

func (x *GetTripByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_driver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripByIDRequest.ProtoReflect.Descriptor instead.
func (*GetTripByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_driver_proto_rawDescGZIP(), []int{2}
}

func (x *GetTripByIDRequest) GetTripId() string {
	if x != nil {
		return x.TripId
	}
	return ""
}

type GetTripByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trip *Trip `protobuf:"bytes,1,opt,name=trip,proto3" json:"trip,omitempty"`
}

func (x *GetTripByIDResponse) Reset() {
	*x = GetTripByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_driver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTripByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripByIDResponse) ProtoMessage() {}

func (x *GetTripByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_driver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripByIDResponse.ProtoReflect.Descriptor instead.
func (*GetTripByIDResponse) Descriptor() ([]byte, []int) {
	return file_api_driver_proto_rawDescGZIP(), []int{3}
}

func (x *GetTripByIDResponse) GetTrip() *Trip {
	if x != nil {
		return x.Trip
	}
	return nil
}

type CancelTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TripId string `protobuf:"bytes,1,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *CancelTripRequest) Reset() {
	*x = CancelTripRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_driver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelTripRequest) ProtoMessage() {}

func (x *CancelTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_driver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelTripRequest.ProtoReflect.Descriptor instead.
func (*CancelTripRequest) Descriptor() ([]byte, []int) {
	return file_api_driver_proto_rawDescGZIP(), []int{4}
}

func (x *CancelTripRequest) GetTripId() string {
	if x != nil {
		return x.TripId
	}
	return ""
}

func (x *CancelTripRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type CancelTripResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelTripResponse) Reset() {
	*x = CancelTripResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_driver_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelTripResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelTripResponse) ProtoMessage() {}

func (x *CancelTripResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_driver_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelTripResponse.ProtoReflect.Descriptor instead.
func (*CancelTripResponse) Descriptor() ([]byte, []int) {
	return file_api_driver_proto_rawDescGZIP(), []int{5}
}

type AcceptTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TripId string `protobuf:"bytes,1,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`
}

func (x *AcceptTripRequest) Reset() {
	*x = AcceptTripRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_driver_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptTripRequest) ProtoMessage() {}

func (x *AcceptTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_driver_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptTripRequest.ProtoReflect.Descriptor instead.
func (*AcceptTripRequest) Descriptor() ([]byte, []int) {
	return file_api_driver_proto_rawDescGZIP(), []int{6}
}

func (x *AcceptTripRequest) GetTripId() string {
	if x != nil {
		return x.TripId
	}
	return ""
}

type AcceptTripResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AcceptTripResponse) Reset() {
	*x = AcceptTripResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_driver_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptTripResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptTripResponse) ProtoMessage() {}

func (x *AcceptTripResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_driver_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptTripResponse.ProtoReflect.Descriptor instead.
func (*AcceptTripResponse) Descriptor() ([]byte, []int) {
	return file_api_driver_proto_rawDescGZIP(), []int{7}
}

type StartTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TripId string `protobuf:"bytes,1,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`
}

func (x *StartTripRequest) Reset() {
	*x = StartTripRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_driver_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTripRequest) ProtoMessage() {}

func (x *StartTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_driver_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTripRequest.ProtoReflect.Descriptor instead.
func (*StartTripRequest) Descriptor() ([]byte, []int) {
	return file_api_driver_proto_rawDescGZIP(), []int{8}
}

func (x *StartTripRequest) GetTripId() string {
	if x != nil {
		return x.TripId
	}
	return ""
}

type StartTripResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartTripResponse) Reset() {
	*x = StartTripResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_driver_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTripResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTripResponse) ProtoMessage() {}

func (x *StartTripResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_driver_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTripResponse.ProtoReflect.Descriptor instead.
func (*StartTripResponse) Descriptor() ([]byte, []int) {
	return file_api_driver_proto_rawDescGZIP(), []int{9}
}

type Trip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DriverId string         `protobuf:"bytes,2,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	From     *LatLngLiteral `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To       *LatLngLiteral `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Price    *Money         `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	Status   TripStatus     `protobuf:"varint,6,opt,name=status,proto3,enum=taxi.driver.TripStatus" json:"status,omitempty"`
}

func (x *Trip) Reset() {
	*x = Trip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_driver_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trip) ProtoMessage() {}

func (x *Trip) ProtoReflect() protoreflect.Message {
	mi := &file_api_driver_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trip.ProtoReflect.Descriptor instead.
func (*Trip) Descriptor() ([]byte, []int) {
	return file_api_driver_proto_rawDescGZIP(), []int{10}
}

func (x *Trip) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Trip) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *Trip) GetFrom() *LatLngLiteral {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Trip) GetTo() *LatLngLiteral {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *Trip) GetPrice() *Money {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *Trip) GetStatus() TripStatus {
	if x != nil {
		return x.Status
	}
	return TripStatus_UNKNOWN
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
		mi := &file_api_driver_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatLngLiteral) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatLngLiteral) ProtoMessage() {}

func (x *LatLngLiteral) ProtoReflect() protoreflect.Message {
	mi := &file_api_driver_proto_msgTypes[11]
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
	return file_api_driver_proto_rawDescGZIP(), []int{11}
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

type Money struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount   float64 `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency string  `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *Money) Reset() {
	*x = Money{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_driver_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Money) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Money) ProtoMessage() {}

func (x *Money) ProtoReflect() protoreflect.Message {
	mi := &file_api_driver_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Money.ProtoReflect.Descriptor instead.
func (*Money) Descriptor() ([]byte, []int) {
	return file_api_driver_proto_rawDescGZIP(), []int{12}
}

func (x *Money) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Money) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

var File_api_driver_proto protoreflect.FileDescriptor

var file_api_driver_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x74, 0x61, 0x78, 0x69, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x3b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x74, 0x72, 0x69, 0x70, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x61, 0x78, 0x69, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x52, 0x05, 0x74, 0x72, 0x69, 0x70, 0x73, 0x22, 0x2d, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x69, 0x70, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x72, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x74, 0x61, 0x78, 0x69, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e,
	0x54, 0x72, 0x69, 0x70, 0x52, 0x04, 0x74, 0x72, 0x69, 0x70, 0x22, 0x44, 0x0a, 0x11, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x72, 0x69, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x22, 0x14, 0x0a, 0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x72, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72,
	0x69, 0x70, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x54, 0x72,
	0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x0a, 0x10, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x72, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x72, 0x69, 0x70, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xea, 0x01, 0x0a,
	0x04, 0x54, 0x72, 0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x74, 0x61, 0x78, 0x69, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x4c,
	0x61, 0x74, 0x4c, 0x6e, 0x67, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x74, 0x61, 0x78, 0x69, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x74,
	0x4c, 0x6e, 0x67, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x28,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x74, 0x61, 0x78, 0x69, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x74, 0x61, 0x78, 0x69, 0x2e,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x33, 0x0a, 0x0d, 0x4c, 0x61, 0x74,
	0x4c, 0x6e, 0x67, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x22, 0x3b,
	0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2a, 0x75, 0x0a, 0x0a, 0x54,
	0x72, 0x69, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x52, 0x49, 0x56, 0x45, 0x52,
	0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x52, 0x49,
	0x56, 0x45, 0x52, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4f,
	0x4e, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4e, 0x44,
	0x45, 0x44, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44,
	0x10, 0x06, 0x32, 0xa8, 0x04, 0x0a, 0x06, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x5e, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x73, 0x12, 0x1c, 0x2e, 0x74, 0x61, 0x78, 0x69,
	0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x61, 0x78, 0x69, 0x2e, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x62, 0x05,
	0x74, 0x72, 0x69, 0x70, 0x73, 0x12, 0x06, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x73, 0x12, 0x70, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1f, 0x2e, 0x74,
	0x61, 0x78, 0x69, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x69, 0x70, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x74, 0x61, 0x78, 0x69, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x69, 0x70, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x62, 0x04, 0x74, 0x72, 0x69, 0x70, 0x12, 0x10, 0x2f,
	0x74, 0x72, 0x69, 0x70, 0x73, 0x2f, 0x7b, 0x74, 0x72, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x6e, 0x0a, 0x0a, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x69, 0x70, 0x12, 0x1e, 0x2e,
	0x74, 0x61, 0x78, 0x69, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x74, 0x61, 0x78, 0x69, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x17, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2f, 0x7b,
	0x74, 0x72, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12,
	0x6e, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x54, 0x72, 0x69, 0x70, 0x12, 0x1e, 0x2e,
	0x74, 0x61, 0x78, 0x69, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x74, 0x61, 0x78, 0x69, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x17, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2f, 0x7b,
	0x74, 0x72, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12,
	0x6c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x69, 0x70, 0x12, 0x1e, 0x2e, 0x74,
	0x61, 0x78, 0x69, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74,
	0x61, 0x78, 0x69, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x16, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2f, 0x7b, 0x74,
	0x72, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x40, 0x5a,
	0x3e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x73, 0x65, 0x2d,
	0x6d, 0x74, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x64, 0x61, 0x73, 0x68, 0x61, 0x67, 0x61, 0x72, 0x6f,
	0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x61, 0x78, 0x69, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_driver_proto_rawDescOnce sync.Once
	file_api_driver_proto_rawDescData = file_api_driver_proto_rawDesc
)

func file_api_driver_proto_rawDescGZIP() []byte {
	file_api_driver_proto_rawDescOnce.Do(func() {
		file_api_driver_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_driver_proto_rawDescData)
	})
	return file_api_driver_proto_rawDescData
}

var file_api_driver_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_driver_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_driver_proto_goTypes = []interface{}{
	(TripStatus)(0),             // 0: taxi.driver.TripStatus
	(*GetTripsRequest)(nil),     // 1: taxi.driver.GetTripsRequest
	(*GetTripsResponse)(nil),    // 2: taxi.driver.GetTripsResponse
	(*GetTripByIDRequest)(nil),  // 3: taxi.driver.GetTripByIDRequest
	(*GetTripByIDResponse)(nil), // 4: taxi.driver.GetTripByIDResponse
	(*CancelTripRequest)(nil),   // 5: taxi.driver.CancelTripRequest
	(*CancelTripResponse)(nil),  // 6: taxi.driver.CancelTripResponse
	(*AcceptTripRequest)(nil),   // 7: taxi.driver.AcceptTripRequest
	(*AcceptTripResponse)(nil),  // 8: taxi.driver.AcceptTripResponse
	(*StartTripRequest)(nil),    // 9: taxi.driver.StartTripRequest
	(*StartTripResponse)(nil),   // 10: taxi.driver.StartTripResponse
	(*Trip)(nil),                // 11: taxi.driver.Trip
	(*LatLngLiteral)(nil),       // 12: taxi.driver.LatLngLiteral
	(*Money)(nil),               // 13: taxi.driver.Money
}
var file_api_driver_proto_depIdxs = []int32{
	11, // 0: taxi.driver.GetTripsResponse.trips:type_name -> taxi.driver.Trip
	11, // 1: taxi.driver.GetTripByIDResponse.trip:type_name -> taxi.driver.Trip
	12, // 2: taxi.driver.Trip.from:type_name -> taxi.driver.LatLngLiteral
	12, // 3: taxi.driver.Trip.to:type_name -> taxi.driver.LatLngLiteral
	13, // 4: taxi.driver.Trip.price:type_name -> taxi.driver.Money
	0,  // 5: taxi.driver.Trip.status:type_name -> taxi.driver.TripStatus
	1,  // 6: taxi.driver.Driver.GetTrips:input_type -> taxi.driver.GetTripsRequest
	3,  // 7: taxi.driver.Driver.GetTripByID:input_type -> taxi.driver.GetTripByIDRequest
	5,  // 8: taxi.driver.Driver.CancelTrip:input_type -> taxi.driver.CancelTripRequest
	7,  // 9: taxi.driver.Driver.AcceptTrip:input_type -> taxi.driver.AcceptTripRequest
	7,  // 10: taxi.driver.Driver.StartTrip:input_type -> taxi.driver.AcceptTripRequest
	2,  // 11: taxi.driver.Driver.GetTrips:output_type -> taxi.driver.GetTripsResponse
	4,  // 12: taxi.driver.Driver.GetTripByID:output_type -> taxi.driver.GetTripByIDResponse
	6,  // 13: taxi.driver.Driver.CancelTrip:output_type -> taxi.driver.CancelTripResponse
	8,  // 14: taxi.driver.Driver.AcceptTrip:output_type -> taxi.driver.AcceptTripResponse
	8,  // 15: taxi.driver.Driver.StartTrip:output_type -> taxi.driver.AcceptTripResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_driver_proto_init() }
func file_api_driver_proto_init() {
	if File_api_driver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_driver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTripsRequest); i {
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
		file_api_driver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTripsResponse); i {
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
		file_api_driver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTripByIDRequest); i {
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
		file_api_driver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTripByIDResponse); i {
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
		file_api_driver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelTripRequest); i {
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
		file_api_driver_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelTripResponse); i {
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
		file_api_driver_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptTripRequest); i {
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
		file_api_driver_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptTripResponse); i {
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
		file_api_driver_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartTripRequest); i {
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
		file_api_driver_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartTripResponse); i {
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
		file_api_driver_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trip); i {
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
		file_api_driver_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_driver_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Money); i {
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
			RawDescriptor: file_api_driver_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_driver_proto_goTypes,
		DependencyIndexes: file_api_driver_proto_depIdxs,
		EnumInfos:         file_api_driver_proto_enumTypes,
		MessageInfos:      file_api_driver_proto_msgTypes,
	}.Build()
	File_api_driver_proto = out.File
	file_api_driver_proto_rawDesc = nil
	file_api_driver_proto_goTypes = nil
	file_api_driver_proto_depIdxs = nil
}