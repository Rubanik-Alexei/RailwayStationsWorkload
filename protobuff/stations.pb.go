// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: protobuff/stations.proto

package protobuff

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

type GetStationWorkloadFromDBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StationName string `protobuf:"bytes,1,opt,name=stationName,proto3" json:"stationName,omitempty"`
}

func (x *GetStationWorkloadFromDBRequest) Reset() {
	*x = GetStationWorkloadFromDBRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuff_stations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStationWorkloadFromDBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStationWorkloadFromDBRequest) ProtoMessage() {}

func (x *GetStationWorkloadFromDBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuff_stations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStationWorkloadFromDBRequest.ProtoReflect.Descriptor instead.
func (*GetStationWorkloadFromDBRequest) Descriptor() ([]byte, []int) {
	return file_protobuff_stations_proto_rawDescGZIP(), []int{0}
}

func (x *GetStationWorkloadFromDBRequest) GetStationName() string {
	if x != nil {
		return x.StationName
	}
	return ""
}

type StationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RespstationName string              `protobuf:"bytes,1,opt,name=RespstationName,proto3" json:"RespstationName,omitempty"`
	RespWorkLoad    map[string]*DayWork `protobuf:"bytes,2,rep,name=RespWorkLoad,proto3" json:"RespWorkLoad,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Error           string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *StationData) Reset() {
	*x = StationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuff_stations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationData) ProtoMessage() {}

func (x *StationData) ProtoReflect() protoreflect.Message {
	mi := &file_protobuff_stations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationData.ProtoReflect.Descriptor instead.
func (*StationData) Descriptor() ([]byte, []int) {
	return file_protobuff_stations_proto_rawDescGZIP(), []int{1}
}

func (x *StationData) GetRespstationName() string {
	if x != nil {
		return x.RespstationName
	}
	return ""
}

func (x *StationData) GetRespWorkLoad() map[string]*DayWork {
	if x != nil {
		return x.RespWorkLoad
	}
	return nil
}

func (x *StationData) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetStationWorkloadFromDBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StationWorkloads []*StationData `protobuf:"bytes,1,rep,name=StationWorkloads,proto3" json:"StationWorkloads,omitempty"`
}

func (x *GetStationWorkloadFromDBResponse) Reset() {
	*x = GetStationWorkloadFromDBResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuff_stations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStationWorkloadFromDBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStationWorkloadFromDBResponse) ProtoMessage() {}

func (x *GetStationWorkloadFromDBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuff_stations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStationWorkloadFromDBResponse.ProtoReflect.Descriptor instead.
func (*GetStationWorkloadFromDBResponse) Descriptor() ([]byte, []int) {
	return file_protobuff_stations_proto_rawDescGZIP(), []int{2}
}

func (x *GetStationWorkloadFromDBResponse) GetStationWorkloads() []*StationData {
	if x != nil {
		return x.StationWorkloads
	}
	return nil
}

type GetStationWorkloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StationName string `protobuf:"bytes,1,opt,name=stationName,proto3" json:"stationName,omitempty"`
	IsUpdateDB  bool   `protobuf:"varint,2,opt,name=isUpdateDB,proto3" json:"isUpdateDB,omitempty"`
}

func (x *GetStationWorkloadRequest) Reset() {
	*x = GetStationWorkloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuff_stations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStationWorkloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStationWorkloadRequest) ProtoMessage() {}

func (x *GetStationWorkloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuff_stations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStationWorkloadRequest.ProtoReflect.Descriptor instead.
func (*GetStationWorkloadRequest) Descriptor() ([]byte, []int) {
	return file_protobuff_stations_proto_rawDescGZIP(), []int{3}
}

func (x *GetStationWorkloadRequest) GetStationName() string {
	if x != nil {
		return x.StationName
	}
	return ""
}

func (x *GetStationWorkloadRequest) GetIsUpdateDB() bool {
	if x != nil {
		return x.IsUpdateDB
	}
	return false
}

type DayWork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DayWorkload map[int32]string `protobuf:"bytes,1,rep,name=DayWorkload,proto3" json:"DayWorkload,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DayWork) Reset() {
	*x = DayWork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuff_stations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DayWork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DayWork) ProtoMessage() {}

func (x *DayWork) ProtoReflect() protoreflect.Message {
	mi := &file_protobuff_stations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DayWork.ProtoReflect.Descriptor instead.
func (*DayWork) Descriptor() ([]byte, []int) {
	return file_protobuff_stations_proto_rawDescGZIP(), []int{4}
}

func (x *DayWork) GetDayWorkload() map[int32]string {
	if x != nil {
		return x.DayWorkload
	}
	return nil
}

type GetStationWorkloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkLoad map[string]*DayWork `protobuf:"bytes,1,rep,name=WorkLoad,proto3" json:"WorkLoad,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Error    string              `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *GetStationWorkloadResponse) Reset() {
	*x = GetStationWorkloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuff_stations_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStationWorkloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStationWorkloadResponse) ProtoMessage() {}

func (x *GetStationWorkloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuff_stations_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStationWorkloadResponse.ProtoReflect.Descriptor instead.
func (*GetStationWorkloadResponse) Descriptor() ([]byte, []int) {
	return file_protobuff_stations_proto_rawDescGZIP(), []int{5}
}

func (x *GetStationWorkloadResponse) GetWorkLoad() map[string]*DayWork {
	if x != nil {
		return x.WorkLoad
	}
	return nil
}

func (x *GetStationWorkloadResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_protobuff_stations_proto protoreflect.FileDescriptor

var file_protobuff_stations_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x72,
	0x6f, 0x6d, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xdc, 0x01,
	0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a,
	0x0f, 0x52, 0x65, 0x73, 0x70, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x57,
	0x6f, 0x72, 0x6b, 0x4c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x57, 0x6f, 0x72, 0x6b, 0x4c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x52,
	0x65, 0x73, 0x70, 0x57, 0x6f, 0x72, 0x6b, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x1a, 0x49, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x70, 0x57, 0x6f, 0x72, 0x6b, 0x4c, 0x6f, 0x61,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x44, 0x61, 0x79, 0x57, 0x6f, 0x72,
	0x6b, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5c, 0x0a, 0x20,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x10, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x22, 0x5d, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69,
	0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x42, 0x22, 0x86, 0x01, 0x0a, 0x07, 0x44, 0x61,
	0x79, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x3b, 0x0a, 0x0b, 0x44, 0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x44, 0x61, 0x79,
	0x57, 0x6f, 0x72, 0x6b, 0x2e, 0x44, 0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x44, 0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x1a, 0x3e, 0x0a, 0x10, 0x44, 0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xc0, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x45, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x4c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x4c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x57, 0x6f, 0x72, 0x6b, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x45,
	0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x4c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x1e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x44, 0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xe1, 0x02, 0x0a, 0x09, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x20, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x6d, 0x61, 0x6e, 0x79, 0x2f, 0x7b, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x30, 0x01, 0x12, 0x72, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x7b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x7d, 0x2f, 0x7b, 0x69, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x42, 0x7d,
	0x12, 0x77, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x42, 0x12, 0x20, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x7b, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuff_stations_proto_rawDescOnce sync.Once
	file_protobuff_stations_proto_rawDescData = file_protobuff_stations_proto_rawDesc
)

func file_protobuff_stations_proto_rawDescGZIP() []byte {
	file_protobuff_stations_proto_rawDescOnce.Do(func() {
		file_protobuff_stations_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuff_stations_proto_rawDescData)
	})
	return file_protobuff_stations_proto_rawDescData
}

var file_protobuff_stations_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protobuff_stations_proto_goTypes = []interface{}{
	(*GetStationWorkloadFromDBRequest)(nil),  // 0: GetStationWorkloadFromDBRequest
	(*StationData)(nil),                      // 1: StationData
	(*GetStationWorkloadFromDBResponse)(nil), // 2: GetStationWorkloadFromDBResponse
	(*GetStationWorkloadRequest)(nil),        // 3: GetStationWorkloadRequest
	(*DayWork)(nil),                          // 4: DayWork
	(*GetStationWorkloadResponse)(nil),       // 5: GetStationWorkloadResponse
	nil,                                      // 6: StationData.RespWorkLoadEntry
	nil,                                      // 7: DayWork.DayWorkloadEntry
	nil,                                      // 8: GetStationWorkloadResponse.WorkLoadEntry
}
var file_protobuff_stations_proto_depIdxs = []int32{
	6, // 0: StationData.RespWorkLoad:type_name -> StationData.RespWorkLoadEntry
	1, // 1: GetStationWorkloadFromDBResponse.StationWorkloads:type_name -> StationData
	7, // 2: DayWork.DayWorkload:type_name -> DayWork.DayWorkloadEntry
	8, // 3: GetStationWorkloadResponse.WorkLoad:type_name -> GetStationWorkloadResponse.WorkLoadEntry
	4, // 4: StationData.RespWorkLoadEntry.value:type_name -> DayWork
	4, // 5: GetStationWorkloadResponse.WorkLoadEntry.value:type_name -> DayWork
	0, // 6: MyService.GetManyStationWorkload:input_type -> GetStationWorkloadFromDBRequest
	3, // 7: MyService.GetStationWorkload:input_type -> GetStationWorkloadRequest
	0, // 8: MyService.GetStationWorkloadFromDB:input_type -> GetStationWorkloadFromDBRequest
	1, // 9: MyService.GetManyStationWorkload:output_type -> StationData
	5, // 10: MyService.GetStationWorkload:output_type -> GetStationWorkloadResponse
	2, // 11: MyService.GetStationWorkloadFromDB:output_type -> GetStationWorkloadFromDBResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_protobuff_stations_proto_init() }
func file_protobuff_stations_proto_init() {
	if File_protobuff_stations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuff_stations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStationWorkloadFromDBRequest); i {
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
		file_protobuff_stations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationData); i {
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
		file_protobuff_stations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStationWorkloadFromDBResponse); i {
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
		file_protobuff_stations_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStationWorkloadRequest); i {
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
		file_protobuff_stations_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DayWork); i {
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
		file_protobuff_stations_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStationWorkloadResponse); i {
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
			RawDescriptor: file_protobuff_stations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuff_stations_proto_goTypes,
		DependencyIndexes: file_protobuff_stations_proto_depIdxs,
		MessageInfos:      file_protobuff_stations_proto_msgTypes,
	}.Build()
	File_protobuff_stations_proto = out.File
	file_protobuff_stations_proto_rawDesc = nil
	file_protobuff_stations_proto_goTypes = nil
	file_protobuff_stations_proto_depIdxs = nil
}
