// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/home/graph/v1/device.proto

package graph

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// Third-party partner's device definition.
type Device struct {
	// Third-party partner's device ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Hardware type of the device (e.g. light, outlet, etc).
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Traits supported by the device.
	Traits []string `protobuf:"bytes,3,rep,name=traits,proto3" json:"traits,omitempty"`
	// Name of the device given by the third party. This includes names given to
	// the device via third party device manufacturer's app, model names for the
	// device, etc.
	Name *DeviceNames `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Indicates whether the state of this device is being reported to Google
	// through ReportStateAndNotification call.
	WillReportState bool `protobuf:"varint,5,opt,name=will_report_state,json=willReportState,proto3" json:"will_report_state,omitempty"`
	// If the third-party partner's cloud configuration includes placing devices
	// in rooms, the name of the room can be provided here.
	RoomHint string `protobuf:"bytes,6,opt,name=room_hint,json=roomHint,proto3" json:"room_hint,omitempty"`
	// As in roomHint, for structures that users set up in the partner's system.
	StructureHint string `protobuf:"bytes,7,opt,name=structure_hint,json=structureHint,proto3" json:"structure_hint,omitempty"`
	// Device manufacturer, model, hardware version, and software version.
	DeviceInfo *DeviceInfo `protobuf:"bytes,8,opt,name=device_info,json=deviceInfo,proto3" json:"device_info,omitempty"`
	// Attributes for the traits supported by the device.
	Attributes *_struct.Struct `protobuf:"bytes,9,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// Custom JSON data provided by the manufacturer and attached to QUERY and
	// EXECUTE requests in AoG.
	CustomData string `protobuf:"bytes,10,opt,name=custom_data,json=customData,proto3" json:"custom_data,omitempty"`
	// IDs of other devices associated with this device. This is used to
	// represent a device group (e.g. bonded zone) or "facets" synced
	// through different flows (e.g. Google Nest Hub Max with a Nest Camera).
	//
	// This may also be used to pass in alternate IDs used to identify a cloud
	// synced device for local execution (i.e. local verification). If used for
	// local verification, this field is synced from the cloud.
	OtherDeviceIds []*AgentOtherDeviceId `protobuf:"bytes,11,rep,name=other_device_ids,json=otherDeviceIds,proto3" json:"other_device_ids,omitempty"`
	// Indicates whether the device is capable of sending notifications. This
	// field will be set by the agent (partner) on an incoming SYNC. If a device
	// is not capable of generating notifications, the partner should set this
	// flag to false. If a partner is not capable of calling
	// ReportStateAndNotification to send notifications to Google, the partner
	// should set this flag to false. If there is a user setting in the partner
	// app to enable notifications and it is turned off, the partner should set
	// this flag to false.
	NotificationSupportedByAgent bool     `protobuf:"varint,12,opt,name=notification_supported_by_agent,json=notificationSupportedByAgent,proto3" json:"notification_supported_by_agent,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_1729f8e53993f499, []int{0}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Device) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Device) GetTraits() []string {
	if m != nil {
		return m.Traits
	}
	return nil
}

func (m *Device) GetName() *DeviceNames {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Device) GetWillReportState() bool {
	if m != nil {
		return m.WillReportState
	}
	return false
}

func (m *Device) GetRoomHint() string {
	if m != nil {
		return m.RoomHint
	}
	return ""
}

func (m *Device) GetStructureHint() string {
	if m != nil {
		return m.StructureHint
	}
	return ""
}

func (m *Device) GetDeviceInfo() *DeviceInfo {
	if m != nil {
		return m.DeviceInfo
	}
	return nil
}

func (m *Device) GetAttributes() *_struct.Struct {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Device) GetCustomData() string {
	if m != nil {
		return m.CustomData
	}
	return ""
}

func (m *Device) GetOtherDeviceIds() []*AgentOtherDeviceId {
	if m != nil {
		return m.OtherDeviceIds
	}
	return nil
}

func (m *Device) GetNotificationSupportedByAgent() bool {
	if m != nil {
		return m.NotificationSupportedByAgent
	}
	return false
}

// Different names for the device.
type DeviceNames struct {
	// Primary name of the device, generally provided by the user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Additional names provided by the user for the device.
	Nicknames []string `protobuf:"bytes,2,rep,name=nicknames,proto3" json:"nicknames,omitempty"`
	// List of names provided by the partner rather than the user, often
	// manufacturer names, SKUs, etc.
	DefaultNames         []string `protobuf:"bytes,3,rep,name=default_names,json=defaultNames,proto3" json:"default_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceNames) Reset()         { *m = DeviceNames{} }
func (m *DeviceNames) String() string { return proto.CompactTextString(m) }
func (*DeviceNames) ProtoMessage()    {}
func (*DeviceNames) Descriptor() ([]byte, []int) {
	return fileDescriptor_1729f8e53993f499, []int{1}
}

func (m *DeviceNames) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceNames.Unmarshal(m, b)
}
func (m *DeviceNames) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceNames.Marshal(b, m, deterministic)
}
func (m *DeviceNames) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceNames.Merge(m, src)
}
func (m *DeviceNames) XXX_Size() int {
	return xxx_messageInfo_DeviceNames.Size(m)
}
func (m *DeviceNames) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceNames.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceNames proto.InternalMessageInfo

func (m *DeviceNames) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceNames) GetNicknames() []string {
	if m != nil {
		return m.Nicknames
	}
	return nil
}

func (m *DeviceNames) GetDefaultNames() []string {
	if m != nil {
		return m.DefaultNames
	}
	return nil
}

// Device information.
type DeviceInfo struct {
	// Device manufacturer.
	Manufacturer string `protobuf:"bytes,1,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	// Device model.
	Model string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	// Device hardware version.
	HwVersion string `protobuf:"bytes,3,opt,name=hw_version,json=hwVersion,proto3" json:"hw_version,omitempty"`
	// Device software version.
	SwVersion            string   `protobuf:"bytes,4,opt,name=sw_version,json=swVersion,proto3" json:"sw_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceInfo) Reset()         { *m = DeviceInfo{} }
func (m *DeviceInfo) String() string { return proto.CompactTextString(m) }
func (*DeviceInfo) ProtoMessage()    {}
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1729f8e53993f499, []int{2}
}

func (m *DeviceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceInfo.Unmarshal(m, b)
}
func (m *DeviceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceInfo.Marshal(b, m, deterministic)
}
func (m *DeviceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceInfo.Merge(m, src)
}
func (m *DeviceInfo) XXX_Size() int {
	return xxx_messageInfo_DeviceInfo.Size(m)
}
func (m *DeviceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceInfo proto.InternalMessageInfo

func (m *DeviceInfo) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *DeviceInfo) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *DeviceInfo) GetHwVersion() string {
	if m != nil {
		return m.HwVersion
	}
	return ""
}

func (m *DeviceInfo) GetSwVersion() string {
	if m != nil {
		return m.SwVersion
	}
	return ""
}

// Identifies a device in the third party or first party system.
type AgentOtherDeviceId struct {
	// The agent's ID. Generally it is the agent's AoG project id.
	AgentId string `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	// Device ID defined by the agent. The device_id must be unique.
	DeviceId             string   `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentOtherDeviceId) Reset()         { *m = AgentOtherDeviceId{} }
func (m *AgentOtherDeviceId) String() string { return proto.CompactTextString(m) }
func (*AgentOtherDeviceId) ProtoMessage()    {}
func (*AgentOtherDeviceId) Descriptor() ([]byte, []int) {
	return fileDescriptor_1729f8e53993f499, []int{3}
}

func (m *AgentOtherDeviceId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentOtherDeviceId.Unmarshal(m, b)
}
func (m *AgentOtherDeviceId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentOtherDeviceId.Marshal(b, m, deterministic)
}
func (m *AgentOtherDeviceId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentOtherDeviceId.Merge(m, src)
}
func (m *AgentOtherDeviceId) XXX_Size() int {
	return xxx_messageInfo_AgentOtherDeviceId.Size(m)
}
func (m *AgentOtherDeviceId) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentOtherDeviceId.DiscardUnknown(m)
}

var xxx_messageInfo_AgentOtherDeviceId proto.InternalMessageInfo

func (m *AgentOtherDeviceId) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *AgentOtherDeviceId) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func init() {
	proto.RegisterType((*Device)(nil), "google.home.graph.v1.Device")
	proto.RegisterType((*DeviceNames)(nil), "google.home.graph.v1.DeviceNames")
	proto.RegisterType((*DeviceInfo)(nil), "google.home.graph.v1.DeviceInfo")
	proto.RegisterType((*AgentOtherDeviceId)(nil), "google.home.graph.v1.AgentOtherDeviceId")
}

func init() { proto.RegisterFile("google/home/graph/v1/device.proto", fileDescriptor_1729f8e53993f499) }

var fileDescriptor_1729f8e53993f499 = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x5f, 0x6b, 0x14, 0x3f,
	0x14, 0x65, 0xff, 0x74, 0xbb, 0x7b, 0xb7, 0xed, 0xef, 0x67, 0x28, 0x1a, 0xb5, 0xd2, 0xed, 0x8a,
	0xb0, 0xf8, 0x30, 0x43, 0x2b, 0x22, 0xe2, 0x53, 0x4b, 0x05, 0x0b, 0xa2, 0x32, 0x05, 0x1f, 0x7c,
	0x09, 0xd9, 0x4d, 0x76, 0x26, 0x38, 0x93, 0x0c, 0x49, 0x66, 0x4b, 0xdf, 0xfd, 0x26, 0x7e, 0x51,
	0xc9, 0xcd, 0xec, 0x76, 0x8b, 0xc5, 0xb7, 0xe4, 0xdc, 0x93, 0x7b, 0xcf, 0x9c, 0x7b, 0x06, 0x4e,
	0x72, 0x63, 0xf2, 0x52, 0xa6, 0x85, 0xa9, 0x64, 0x9a, 0x5b, 0x5e, 0x17, 0xe9, 0xea, 0x34, 0x15,
	0x72, 0xa5, 0x16, 0x32, 0xa9, 0xad, 0xf1, 0x86, 0x1c, 0x46, 0x4a, 0x12, 0x28, 0x09, 0x52, 0x92,
	0xd5, 0xe9, 0xb3, 0xa3, 0xf6, 0x21, 0x72, 0xe6, 0xcd, 0x32, 0x75, 0xde, 0x36, 0x0b, 0x1f, 0xdf,
	0x4c, 0x7f, 0xf7, 0x61, 0x70, 0x89, 0x4d, 0xc8, 0x01, 0x74, 0x95, 0xa0, 0x9d, 0x49, 0x67, 0x36,
	0xca, 0xba, 0x4a, 0x10, 0x02, 0x7d, 0x7f, 0x5b, 0x4b, 0xda, 0x45, 0x04, 0xcf, 0xe4, 0x31, 0x0c,
	0xbc, 0xe5, 0xca, 0x3b, 0xda, 0x9b, 0xf4, 0x66, 0xa3, 0xac, 0xbd, 0x91, 0xb7, 0xd0, 0xd7, 0xbc,
	0x92, 0xb4, 0x3f, 0xe9, 0xcc, 0xc6, 0x67, 0x27, 0xc9, 0x43, 0x4a, 0x92, 0x38, 0xe7, 0x0b, 0xaf,
	0xa4, 0xcb, 0x90, 0x4e, 0x5e, 0xc3, 0xa3, 0x1b, 0x55, 0x96, 0xcc, 0xca, 0xda, 0x58, 0xcf, 0x9c,
	0xe7, 0x5e, 0xd2, 0x9d, 0x49, 0x67, 0x36, 0xcc, 0xfe, 0x0b, 0x85, 0x0c, 0xf1, 0xeb, 0x00, 0x93,
	0xe7, 0x30, 0xb2, 0xc6, 0x54, 0xac, 0x50, 0xda, 0xd3, 0x01, 0x6a, 0x1a, 0x06, 0xe0, 0x93, 0xd2,
	0x9e, 0xbc, 0x82, 0x83, 0xf8, 0x59, 0x8d, 0x95, 0x91, 0xb1, 0x8b, 0x8c, 0xfd, 0x0d, 0x8a, 0xb4,
	0x73, 0x18, 0x47, 0xc7, 0x98, 0xd2, 0x4b, 0x43, 0x87, 0xa8, 0x76, 0xf2, 0x2f, 0xb5, 0x57, 0x7a,
	0x69, 0x32, 0x10, 0x9b, 0x33, 0x79, 0x07, 0xc0, 0xbd, 0xb7, 0x6a, 0xde, 0x78, 0xe9, 0xe8, 0x08,
	0x3b, 0x3c, 0x59, 0x77, 0x58, 0x7b, 0x9c, 0x5c, 0xe3, 0xd8, 0x6c, 0x8b, 0x4a, 0x8e, 0x61, 0xbc,
	0x68, 0x9c, 0x37, 0x15, 0x13, 0xdc, 0x73, 0x0a, 0xa8, 0x0f, 0x22, 0x74, 0xc9, 0x3d, 0x27, 0x19,
	0xfc, 0x6f, 0x7c, 0x21, 0x2d, 0x5b, 0x4b, 0x14, 0x8e, 0x8e, 0x27, 0xbd, 0xd9, 0xf8, 0x6c, 0xf6,
	0xb0, 0xc2, 0xf3, 0x5c, 0x6a, 0xff, 0x35, 0x3c, 0x69, 0xb5, 0x8a, 0xec, 0xc0, 0x6c, 0x5f, 0x1d,
	0xf9, 0x08, 0xc7, 0xda, 0x78, 0xb5, 0x54, 0x0b, 0xee, 0x95, 0xd1, 0xcc, 0x35, 0x75, 0x70, 0x54,
	0x0a, 0x36, 0xbf, 0x65, 0x3c, 0xbc, 0xa7, 0x7b, 0x68, 0xf7, 0xd1, 0x36, 0xed, 0x7a, 0xcd, 0xba,
	0xb8, 0xc5, 0x19, 0x53, 0x01, 0xe3, 0xad, 0xe5, 0x85, 0x64, 0xe0, 0xb6, 0x63, 0x56, 0xe2, 0x2a,
	0x8f, 0x60, 0xa4, 0xd5, 0xe2, 0x67, 0x38, 0x3b, 0xda, 0xc5, 0x70, 0xdc, 0x01, 0xe4, 0x25, 0xec,
	0x0b, 0xb9, 0xe4, 0x4d, 0xe9, 0x59, 0x64, 0xc4, 0xf8, 0xec, 0xb5, 0x20, 0xb6, 0x9d, 0xfe, 0xea,
	0x00, 0xdc, 0xb9, 0x4e, 0xa6, 0xb0, 0x57, 0x71, 0xdd, 0x2c, 0x39, 0xee, 0xcf, 0xb6, 0xd3, 0xee,
	0x61, 0xe4, 0x10, 0x76, 0x2a, 0x23, 0x64, 0xd9, 0x86, 0x34, 0x5e, 0xc8, 0x0b, 0x80, 0xe2, 0x86,
	0xad, 0xa4, 0x75, 0xca, 0x68, 0xda, 0xc3, 0xd2, 0xa8, 0xb8, 0xf9, 0x1e, 0x81, 0x50, 0x76, 0x77,
	0xe5, 0x7e, 0x2c, 0xbb, 0x75, 0x79, 0xfa, 0x19, 0xc8, 0xdf, 0xce, 0x92, 0xa7, 0x30, 0x44, 0xbf,
	0xd8, 0xe6, 0x1f, 0xd9, 0xc5, 0xfb, 0x95, 0x08, 0xc9, 0xdc, 0xac, 0xac, 0x15, 0x32, 0x6c, 0x13,
	0x23, 0x2e, 0xe6, 0x40, 0x17, 0xa6, 0x7a, 0x70, 0x81, 0x17, 0xad, 0xa9, 0xdf, 0x42, 0x6a, 0x7e,
	0xbc, 0x6f, 0x29, 0xb9, 0x29, 0xb9, 0xce, 0x13, 0x63, 0xf3, 0x34, 0x97, 0x1a, 0x13, 0x95, 0xc6,
	0x12, 0xaf, 0x95, 0xbb, 0xff, 0xff, 0x7f, 0xc0, 0xc3, 0x7c, 0x80, 0xac, 0x37, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x78, 0x5e, 0xb4, 0x17, 0x24, 0x04, 0x00, 0x00,
}
