// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: workload_service/protobuff/wl.proto

package protobuff

import (
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
		mi := &file_workload_service_protobuff_wl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStationWorkloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStationWorkloadRequest) ProtoMessage() {}

func (x *GetStationWorkloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workload_service_protobuff_wl_proto_msgTypes[0]
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
	return file_workload_service_protobuff_wl_proto_rawDescGZIP(), []int{0}
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
		mi := &file_workload_service_protobuff_wl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationData) ProtoMessage() {}

func (x *StationData) ProtoReflect() protoreflect.Message {
	mi := &file_workload_service_protobuff_wl_proto_msgTypes[1]
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
	return file_workload_service_protobuff_wl_proto_rawDescGZIP(), []int{1}
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

type DayWork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DayWorkload map[int32]string `protobuf:"bytes,1,rep,name=DayWorkload,proto3" json:"DayWorkload,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DayWork) Reset() {
	*x = DayWork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workload_service_protobuff_wl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DayWork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DayWork) ProtoMessage() {}

func (x *DayWork) ProtoReflect() protoreflect.Message {
	mi := &file_workload_service_protobuff_wl_proto_msgTypes[2]
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
	return file_workload_service_protobuff_wl_proto_rawDescGZIP(), []int{2}
}

func (x *DayWork) GetDayWorkload() map[int32]string {
	if x != nil {
		return x.DayWorkload
	}
	return nil
}

var File_workload_service_protobuff_wl_proto protoreflect.FileDescriptor

var file_workload_service_protobuff_wl_proto_rawDesc = []byte{
	0x0a, 0x23, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x2f, 0x77, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x42, 0x22, 0xdc, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x52,
	0x65, 0x73, 0x70, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x42,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x57, 0x6f, 0x72, 0x6b, 0x4c, 0x6f, 0x61, 0x64, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x57, 0x6f, 0x72, 0x6b, 0x4c, 0x6f, 0x61, 0x64, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x57, 0x6f, 0x72, 0x6b, 0x4c, 0x6f,
	0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x49, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x70,
	0x57, 0x6f, 0x72, 0x6b, 0x4c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x1e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x44, 0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x86, 0x01, 0x0a, 0x07, 0x44, 0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x12,
	0x3b, 0x0a, 0x0b, 0x44, 0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x44, 0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x2e, 0x44,
	0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x44, 0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x3e, 0x0a, 0x10,
	0x44, 0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x53, 0x0a, 0x0f,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x40, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x30,
	0x01, 0x42, 0x1e, 0x5a, 0x1c, 0x2e, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workload_service_protobuff_wl_proto_rawDescOnce sync.Once
	file_workload_service_protobuff_wl_proto_rawDescData = file_workload_service_protobuff_wl_proto_rawDesc
)

func file_workload_service_protobuff_wl_proto_rawDescGZIP() []byte {
	file_workload_service_protobuff_wl_proto_rawDescOnce.Do(func() {
		file_workload_service_protobuff_wl_proto_rawDescData = protoimpl.X.CompressGZIP(file_workload_service_protobuff_wl_proto_rawDescData)
	})
	return file_workload_service_protobuff_wl_proto_rawDescData
}

var file_workload_service_protobuff_wl_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_workload_service_protobuff_wl_proto_goTypes = []interface{}{
	(*GetStationWorkloadRequest)(nil), // 0: GetStationWorkloadRequest
	(*StationData)(nil),               // 1: StationData
	(*DayWork)(nil),                   // 2: DayWork
	nil,                               // 3: StationData.RespWorkLoadEntry
	nil,                               // 4: DayWork.DayWorkloadEntry
}
var file_workload_service_protobuff_wl_proto_depIdxs = []int32{
	3, // 0: StationData.RespWorkLoad:type_name -> StationData.RespWorkLoadEntry
	4, // 1: DayWork.DayWorkload:type_name -> DayWork.DayWorkloadEntry
	2, // 2: StationData.RespWorkLoadEntry.value:type_name -> DayWork
	0, // 3: WorkloadService.GetStationWorkload:input_type -> GetStationWorkloadRequest
	1, // 4: WorkloadService.GetStationWorkload:output_type -> StationData
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_workload_service_protobuff_wl_proto_init() }
func file_workload_service_protobuff_wl_proto_init() {
	if File_workload_service_protobuff_wl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_workload_service_protobuff_wl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_workload_service_protobuff_wl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_workload_service_protobuff_wl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_workload_service_protobuff_wl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workload_service_protobuff_wl_proto_goTypes,
		DependencyIndexes: file_workload_service_protobuff_wl_proto_depIdxs,
		MessageInfos:      file_workload_service_protobuff_wl_proto_msgTypes,
	}.Build()
	File_workload_service_protobuff_wl_proto = out.File
	file_workload_service_protobuff_wl_proto_rawDesc = nil
	file_workload_service_protobuff_wl_proto_goTypes = nil
	file_workload_service_protobuff_wl_proto_depIdxs = nil
}
