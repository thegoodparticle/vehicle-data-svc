// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: vehicledata/vehicle-data.proto

package vehicledata

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegistrationNumber string `protobuf:"bytes,1,opt,name=RegistrationNumber,proto3" json:"RegistrationNumber,omitempty"`
}

func (x *RegistrationRequest) Reset() {
	*x = RegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicledata_vehicle_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationRequest) ProtoMessage() {}

func (x *RegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vehicledata_vehicle_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationRequest.ProtoReflect.Descriptor instead.
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return file_vehicledata_vehicle_data_proto_rawDescGZIP(), []int{0}
}

func (x *RegistrationRequest) GetRegistrationNumber() string {
	if x != nil {
		return x.RegistrationNumber
	}
	return ""
}

type ChassisEngineNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChassisNumber string `protobuf:"bytes,1,opt,name=ChassisNumber,proto3" json:"ChassisNumber,omitempty"`
	EngineNumber  string `protobuf:"bytes,2,opt,name=EngineNumber,proto3" json:"EngineNumber,omitempty"`
}

func (x *ChassisEngineNumberRequest) Reset() {
	*x = ChassisEngineNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicledata_vehicle_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChassisEngineNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChassisEngineNumberRequest) ProtoMessage() {}

func (x *ChassisEngineNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vehicledata_vehicle_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChassisEngineNumberRequest.ProtoReflect.Descriptor instead.
func (*ChassisEngineNumberRequest) Descriptor() ([]byte, []int) {
	return file_vehicledata_vehicle_data_proto_rawDescGZIP(), []int{1}
}

func (x *ChassisEngineNumberRequest) GetChassisNumber() string {
	if x != nil {
		return x.ChassisNumber
	}
	return ""
}

func (x *ChassisEngineNumberRequest) GetEngineNumber() string {
	if x != nil {
		return x.EngineNumber
	}
	return ""
}

type VehicleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegistrationNumber string                 `protobuf:"bytes,1,opt,name=RegistrationNumber,proto3" json:"RegistrationNumber,omitempty"`
	VehicleModel       string                 `protobuf:"bytes,2,opt,name=VehicleModel,proto3" json:"VehicleModel,omitempty"`
	Company            string                 `protobuf:"bytes,3,opt,name=Company,proto3" json:"Company,omitempty"`
	RegistrationDate   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=RegistrationDate,proto3" json:"RegistrationDate,omitempty"`
}

func (x *VehicleInfo) Reset() {
	*x = VehicleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicledata_vehicle_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleInfo) ProtoMessage() {}

func (x *VehicleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_vehicledata_vehicle_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleInfo.ProtoReflect.Descriptor instead.
func (*VehicleInfo) Descriptor() ([]byte, []int) {
	return file_vehicledata_vehicle_data_proto_rawDescGZIP(), []int{2}
}

func (x *VehicleInfo) GetRegistrationNumber() string {
	if x != nil {
		return x.RegistrationNumber
	}
	return ""
}

func (x *VehicleInfo) GetVehicleModel() string {
	if x != nil {
		return x.VehicleModel
	}
	return ""
}

func (x *VehicleInfo) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *VehicleInfo) GetRegistrationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.RegistrationDate
	}
	return nil
}

type LicenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LicenseNumber string `protobuf:"bytes,1,opt,name=LicenseNumber,proto3" json:"LicenseNumber,omitempty"`
}

func (x *LicenseRequest) Reset() {
	*x = LicenseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicledata_vehicle_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseRequest) ProtoMessage() {}

func (x *LicenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vehicledata_vehicle_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseRequest.ProtoReflect.Descriptor instead.
func (*LicenseRequest) Descriptor() ([]byte, []int) {
	return file_vehicledata_vehicle_data_proto_rawDescGZIP(), []int{3}
}

func (x *LicenseRequest) GetLicenseNumber() string {
	if x != nil {
		return x.LicenseNumber
	}
	return ""
}

type DriverInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LicenseNumber string                 `protobuf:"bytes,1,opt,name=LicenseNumber,proto3" json:"LicenseNumber,omitempty"`
	DriverName    string                 `protobuf:"bytes,2,opt,name=DriverName,proto3" json:"DriverName,omitempty"`
	LicenseDate   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=LicenseDate,proto3" json:"LicenseDate,omitempty"`
}

func (x *DriverInfo) Reset() {
	*x = DriverInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicledata_vehicle_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverInfo) ProtoMessage() {}

func (x *DriverInfo) ProtoReflect() protoreflect.Message {
	mi := &file_vehicledata_vehicle_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverInfo.ProtoReflect.Descriptor instead.
func (*DriverInfo) Descriptor() ([]byte, []int) {
	return file_vehicledata_vehicle_data_proto_rawDescGZIP(), []int{4}
}

func (x *DriverInfo) GetLicenseNumber() string {
	if x != nil {
		return x.LicenseNumber
	}
	return ""
}

func (x *DriverInfo) GetDriverName() string {
	if x != nil {
		return x.DriverName
	}
	return ""
}

func (x *DriverInfo) GetLicenseDate() *timestamppb.Timestamp {
	if x != nil {
		return x.LicenseDate
	}
	return nil
}

var File_vehicledata_vehicle_data_proto protoreflect.FileDescriptor

var file_vehicledata_vehicle_data_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45,
	0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x66, 0x0a, 0x1a, 0x43, 0x68, 0x61, 0x73, 0x73, 0x69, 0x73,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x73, 0x73, 0x69, 0x73, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x68, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xc3, 0x01,
	0x0a, 0x0b, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a,
	0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a,
	0x0c, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x46, 0x0a, 0x10, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x22, 0x36, 0x0a, 0x0e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x4c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x90, 0x01, 0x0a, 0x0a,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0d, 0x4c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x3c, 0x0a, 0x0b, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x32, 0xa9,
	0x02, 0x0a, 0x0b, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x5c,
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x1d,
	0x47, 0x65, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79,
	0x43, 0x68, 0x61, 0x73, 0x73, 0x69, 0x73, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x27, 0x2e,
	0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x68, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x00, 0x12, 0x56, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x79, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1b, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x67, 0x6f, 0x6f, 0x64,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x2d, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x73, 0x76, 0x63, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vehicledata_vehicle_data_proto_rawDescOnce sync.Once
	file_vehicledata_vehicle_data_proto_rawDescData = file_vehicledata_vehicle_data_proto_rawDesc
)

func file_vehicledata_vehicle_data_proto_rawDescGZIP() []byte {
	file_vehicledata_vehicle_data_proto_rawDescOnce.Do(func() {
		file_vehicledata_vehicle_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_vehicledata_vehicle_data_proto_rawDescData)
	})
	return file_vehicledata_vehicle_data_proto_rawDescData
}

var file_vehicledata_vehicle_data_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_vehicledata_vehicle_data_proto_goTypes = []interface{}{
	(*RegistrationRequest)(nil),        // 0: vehicledata.RegistrationRequest
	(*ChassisEngineNumberRequest)(nil), // 1: vehicledata.ChassisEngineNumberRequest
	(*VehicleInfo)(nil),                // 2: vehicledata.VehicleInfo
	(*LicenseRequest)(nil),             // 3: vehicledata.LicenseRequest
	(*DriverInfo)(nil),                 // 4: vehicledata.DriverInfo
	(*timestamppb.Timestamp)(nil),      // 5: google.protobuf.Timestamp
}
var file_vehicledata_vehicle_data_proto_depIdxs = []int32{
	5, // 0: vehicledata.VehicleInfo.RegistrationDate:type_name -> google.protobuf.Timestamp
	5, // 1: vehicledata.DriverInfo.LicenseDate:type_name -> google.protobuf.Timestamp
	0, // 2: vehicledata.VehicleData.GetVehicleDataByRegistration:input_type -> vehicledata.RegistrationRequest
	1, // 3: vehicledata.VehicleData.GetVehicleDataByChassisEngine:input_type -> vehicledata.ChassisEngineNumberRequest
	3, // 4: vehicledata.VehicleData.GetDriverDataByLicenseNumber:input_type -> vehicledata.LicenseRequest
	2, // 5: vehicledata.VehicleData.GetVehicleDataByRegistration:output_type -> vehicledata.VehicleInfo
	2, // 6: vehicledata.VehicleData.GetVehicleDataByChassisEngine:output_type -> vehicledata.VehicleInfo
	4, // 7: vehicledata.VehicleData.GetDriverDataByLicenseNumber:output_type -> vehicledata.DriverInfo
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_vehicledata_vehicle_data_proto_init() }
func file_vehicledata_vehicle_data_proto_init() {
	if File_vehicledata_vehicle_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vehicledata_vehicle_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationRequest); i {
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
		file_vehicledata_vehicle_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChassisEngineNumberRequest); i {
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
		file_vehicledata_vehicle_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleInfo); i {
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
		file_vehicledata_vehicle_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseRequest); i {
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
		file_vehicledata_vehicle_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverInfo); i {
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
			RawDescriptor: file_vehicledata_vehicle_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vehicledata_vehicle_data_proto_goTypes,
		DependencyIndexes: file_vehicledata_vehicle_data_proto_depIdxs,
		MessageInfos:      file_vehicledata_vehicle_data_proto_msgTypes,
	}.Build()
	File_vehicledata_vehicle_data_proto = out.File
	file_vehicledata_vehicle_data_proto_rawDesc = nil
	file_vehicledata_vehicle_data_proto_goTypes = nil
	file_vehicledata_vehicle_data_proto_depIdxs = nil
}
