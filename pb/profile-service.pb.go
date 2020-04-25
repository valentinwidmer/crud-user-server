// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: profile-service.proto

package profile

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Date struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Day   int32 `protobuf:"varint,1,opt,name=day,proto3" json:"day,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Year  int32 `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *Date) Reset() {
	*x = Date{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Date) ProtoMessage() {}

func (x *Date) ProtoReflect() protoreflect.Message {
	mi := &file_profile_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Date.ProtoReflect.Descriptor instead.
func (*Date) Descriptor() ([]byte, []int) {
	return file_profile_service_proto_rawDescGZIP(), []int{0}
}

func (x *Date) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *Date) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *Date) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

type ImageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ImageData) Reset() {
	*x = ImageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageData) ProtoMessage() {}

func (x *ImageData) ProtoReflect() protoreflect.Message {
	mi := &file_profile_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageData.ProtoReflect.Descriptor instead.
func (*ImageData) Descriptor() ([]byte, []int) {
	return file_profile_service_proto_rawDescGZIP(), []int{1}
}

func (x *ImageData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProfileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GivenName string `protobuf:"bytes,2,opt,name=givenName,proto3" json:"givenName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Birthday  *Date  `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ProfileData) Reset() {
	*x = ProfileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileData) ProtoMessage() {}

func (x *ProfileData) ProtoReflect() protoreflect.Message {
	mi := &file_profile_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileData.ProtoReflect.Descriptor instead.
func (*ProfileData) Descriptor() ([]byte, []int) {
	return file_profile_service_proto_rawDescGZIP(), []int{2}
}

func (x *ProfileData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProfileData) GetGivenName() string {
	if x != nil {
		return x.GivenName
	}
	return ""
}

func (x *ProfileData) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *ProfileData) GetBirthday() *Date {
	if x != nil {
		return x.Birthday
	}
	return nil
}

func (x *ProfileData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ProfileOneof:
	//	*Profile_ProfileData
	//	*Profile_ImageData
	ProfileOneof isProfile_ProfileOneof `protobuf_oneof:"profile_oneof"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_profile_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_profile_service_proto_rawDescGZIP(), []int{3}
}

func (m *Profile) GetProfileOneof() isProfile_ProfileOneof {
	if m != nil {
		return m.ProfileOneof
	}
	return nil
}

func (x *Profile) GetProfileData() *ProfileData {
	if x, ok := x.GetProfileOneof().(*Profile_ProfileData); ok {
		return x.ProfileData
	}
	return nil
}

func (x *Profile) GetImageData() *ImageData {
	if x, ok := x.GetProfileOneof().(*Profile_ImageData); ok {
		return x.ImageData
	}
	return nil
}

type isProfile_ProfileOneof interface {
	isProfile_ProfileOneof()
}

type Profile_ProfileData struct {
	ProfileData *ProfileData `protobuf:"bytes,1,opt,name=profileData,proto3,oneof"`
}

type Profile_ImageData struct {
	ImageData *ImageData `protobuf:"bytes,2,opt,name=imageData,proto3,oneof"`
}

func (*Profile_ProfileData) isProfile_ProfileOneof() {}

func (*Profile_ImageData) isProfile_ProfileOneof() {}

type ProfileId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProfileId) Reset() {
	*x = ProfileId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileId) ProtoMessage() {}

func (x *ProfileId) ProtoReflect() protoreflect.Message {
	mi := &file_profile_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileId.ProtoReflect.Descriptor instead.
func (*ProfileId) Descriptor() ([]byte, []int) {
	return file_profile_service_proto_rawDescGZIP(), []int{4}
}

func (x *ProfileId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_profile_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_profile_service_proto_rawDescGZIP(), []int{5}
}

func (x *Status) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_profile_service_proto protoreflect.FileDescriptor

var file_profile_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x22, 0x42, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x22, 0x1f, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x98, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x29, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65,
	0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0x88, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52,
	0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0f, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0x1b, 0x0a, 0x09, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xef, 0x01, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a,
	0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x28, 0x01, 0x12, 0x35, 0x0a, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x30, 0x01, 0x12, 0x37, 0x0a,
	0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x28, 0x01, 0x12, 0x34, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profile_service_proto_rawDescOnce sync.Once
	file_profile_service_proto_rawDescData = file_profile_service_proto_rawDesc
)

func file_profile_service_proto_rawDescGZIP() []byte {
	file_profile_service_proto_rawDescOnce.Do(func() {
		file_profile_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_profile_service_proto_rawDescData)
	})
	return file_profile_service_proto_rawDescData
}

var file_profile_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_profile_service_proto_goTypes = []interface{}{
	(*Date)(nil),        // 0: profile.Date
	(*ImageData)(nil),   // 1: profile.ImageData
	(*ProfileData)(nil), // 2: profile.ProfileData
	(*Profile)(nil),     // 3: profile.Profile
	(*ProfileId)(nil),   // 4: profile.ProfileId
	(*Status)(nil),      // 5: profile.Status
}
var file_profile_service_proto_depIdxs = []int32{
	0, // 0: profile.ProfileData.birthday:type_name -> profile.Date
	2, // 1: profile.Profile.profileData:type_name -> profile.ProfileData
	1, // 2: profile.Profile.imageData:type_name -> profile.ImageData
	3, // 3: profile.ProfileService.createProfile:input_type -> profile.Profile
	4, // 4: profile.ProfileService.readProfile:input_type -> profile.ProfileId
	3, // 5: profile.ProfileService.updateProfile:input_type -> profile.Profile
	4, // 6: profile.ProfileService.deleteProfile:input_type -> profile.ProfileId
	4, // 7: profile.ProfileService.createProfile:output_type -> profile.ProfileId
	3, // 8: profile.ProfileService.readProfile:output_type -> profile.Profile
	4, // 9: profile.ProfileService.updateProfile:output_type -> profile.ProfileId
	5, // 10: profile.ProfileService.deleteProfile:output_type -> profile.Status
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_profile_service_proto_init() }
func file_profile_service_proto_init() {
	if File_profile_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profile_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Date); i {
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
		file_profile_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageData); i {
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
		file_profile_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileData); i {
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
		file_profile_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_profile_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileId); i {
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
		file_profile_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
	file_profile_service_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Profile_ProfileData)(nil),
		(*Profile_ImageData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_profile_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profile_service_proto_goTypes,
		DependencyIndexes: file_profile_service_proto_depIdxs,
		MessageInfos:      file_profile_service_proto_msgTypes,
	}.Build()
	File_profile_service_proto = out.File
	file_profile_service_proto_rawDesc = nil
	file_profile_service_proto_goTypes = nil
	file_profile_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileServiceClient interface {
	CreateProfile(ctx context.Context, opts ...grpc.CallOption) (ProfileService_CreateProfileClient, error)
	ReadProfile(ctx context.Context, in *ProfileId, opts ...grpc.CallOption) (ProfileService_ReadProfileClient, error)
	UpdateProfile(ctx context.Context, opts ...grpc.CallOption) (ProfileService_UpdateProfileClient, error)
	DeleteProfile(ctx context.Context, in *ProfileId, opts ...grpc.CallOption) (*Status, error)
}

type profileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileServiceClient(cc grpc.ClientConnInterface) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) CreateProfile(ctx context.Context, opts ...grpc.CallOption) (ProfileService_CreateProfileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProfileService_serviceDesc.Streams[0], "/profile.ProfileService/createProfile", opts...)
	if err != nil {
		return nil, err
	}
	x := &profileServiceCreateProfileClient{stream}
	return x, nil
}

type ProfileService_CreateProfileClient interface {
	Send(*Profile) error
	CloseAndRecv() (*ProfileId, error)
	grpc.ClientStream
}

type profileServiceCreateProfileClient struct {
	grpc.ClientStream
}

func (x *profileServiceCreateProfileClient) Send(m *Profile) error {
	return x.ClientStream.SendMsg(m)
}

func (x *profileServiceCreateProfileClient) CloseAndRecv() (*ProfileId, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ProfileId)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *profileServiceClient) ReadProfile(ctx context.Context, in *ProfileId, opts ...grpc.CallOption) (ProfileService_ReadProfileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProfileService_serviceDesc.Streams[1], "/profile.ProfileService/readProfile", opts...)
	if err != nil {
		return nil, err
	}
	x := &profileServiceReadProfileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProfileService_ReadProfileClient interface {
	Recv() (*Profile, error)
	grpc.ClientStream
}

type profileServiceReadProfileClient struct {
	grpc.ClientStream
}

func (x *profileServiceReadProfileClient) Recv() (*Profile, error) {
	m := new(Profile)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *profileServiceClient) UpdateProfile(ctx context.Context, opts ...grpc.CallOption) (ProfileService_UpdateProfileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProfileService_serviceDesc.Streams[2], "/profile.ProfileService/updateProfile", opts...)
	if err != nil {
		return nil, err
	}
	x := &profileServiceUpdateProfileClient{stream}
	return x, nil
}

type ProfileService_UpdateProfileClient interface {
	Send(*Profile) error
	CloseAndRecv() (*ProfileId, error)
	grpc.ClientStream
}

type profileServiceUpdateProfileClient struct {
	grpc.ClientStream
}

func (x *profileServiceUpdateProfileClient) Send(m *Profile) error {
	return x.ClientStream.SendMsg(m)
}

func (x *profileServiceUpdateProfileClient) CloseAndRecv() (*ProfileId, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ProfileId)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *profileServiceClient) DeleteProfile(ctx context.Context, in *ProfileId, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/deleteProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
type ProfileServiceServer interface {
	CreateProfile(ProfileService_CreateProfileServer) error
	ReadProfile(*ProfileId, ProfileService_ReadProfileServer) error
	UpdateProfile(ProfileService_UpdateProfileServer) error
	DeleteProfile(context.Context, *ProfileId) (*Status, error)
}

// UnimplementedProfileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (*UnimplementedProfileServiceServer) CreateProfile(ProfileService_CreateProfileServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}
func (*UnimplementedProfileServiceServer) ReadProfile(*ProfileId, ProfileService_ReadProfileServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadProfile not implemented")
}
func (*UnimplementedProfileServiceServer) UpdateProfile(ProfileService_UpdateProfileServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (*UnimplementedProfileServiceServer) DeleteProfile(context.Context, *ProfileId) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProfile not implemented")
}

func RegisterProfileServiceServer(s *grpc.Server, srv ProfileServiceServer) {
	s.RegisterService(&_ProfileService_serviceDesc, srv)
}

func _ProfileService_CreateProfile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProfileServiceServer).CreateProfile(&profileServiceCreateProfileServer{stream})
}

type ProfileService_CreateProfileServer interface {
	SendAndClose(*ProfileId) error
	Recv() (*Profile, error)
	grpc.ServerStream
}

type profileServiceCreateProfileServer struct {
	grpc.ServerStream
}

func (x *profileServiceCreateProfileServer) SendAndClose(m *ProfileId) error {
	return x.ServerStream.SendMsg(m)
}

func (x *profileServiceCreateProfileServer) Recv() (*Profile, error) {
	m := new(Profile)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProfileService_ReadProfile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProfileId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProfileServiceServer).ReadProfile(m, &profileServiceReadProfileServer{stream})
}

type ProfileService_ReadProfileServer interface {
	Send(*Profile) error
	grpc.ServerStream
}

type profileServiceReadProfileServer struct {
	grpc.ServerStream
}

func (x *profileServiceReadProfileServer) Send(m *Profile) error {
	return x.ServerStream.SendMsg(m)
}

func _ProfileService_UpdateProfile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProfileServiceServer).UpdateProfile(&profileServiceUpdateProfileServer{stream})
}

type ProfileService_UpdateProfileServer interface {
	SendAndClose(*ProfileId) error
	Recv() (*Profile, error)
	grpc.ServerStream
}

type profileServiceUpdateProfileServer struct {
	grpc.ServerStream
}

func (x *profileServiceUpdateProfileServer) SendAndClose(m *ProfileId) error {
	return x.ServerStream.SendMsg(m)
}

func (x *profileServiceUpdateProfileServer) Recv() (*Profile, error) {
	m := new(Profile)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProfileService_DeleteProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).DeleteProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/DeleteProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).DeleteProfile(ctx, req.(*ProfileId))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profile.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "deleteProfile",
			Handler:    _ProfileService_DeleteProfile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "createProfile",
			Handler:       _ProfileService_CreateProfile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "readProfile",
			Handler:       _ProfileService_ReadProfile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "updateProfile",
			Handler:       _ProfileService_UpdateProfile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "profile-service.proto",
}