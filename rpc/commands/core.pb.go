// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commands/core.proto

package commands

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PlatformInstallReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	PlatformPackage      string    `protobuf:"bytes,2,opt,name=platform_package,json=platformPackage,proto3" json:"platform_package,omitempty"`
	Architecture         string    `protobuf:"bytes,3,opt,name=architecture,proto3" json:"architecture,omitempty"`
	Version              string    `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PlatformInstallReq) Reset()         { *m = PlatformInstallReq{} }
func (m *PlatformInstallReq) String() string { return proto.CompactTextString(m) }
func (*PlatformInstallReq) ProtoMessage()    {}
func (*PlatformInstallReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed02318f567db566, []int{0}
}

func (m *PlatformInstallReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformInstallReq.Unmarshal(m, b)
}
func (m *PlatformInstallReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformInstallReq.Marshal(b, m, deterministic)
}
func (m *PlatformInstallReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformInstallReq.Merge(m, src)
}
func (m *PlatformInstallReq) XXX_Size() int {
	return xxx_messageInfo_PlatformInstallReq.Size(m)
}
func (m *PlatformInstallReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformInstallReq.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformInstallReq proto.InternalMessageInfo

func (m *PlatformInstallReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *PlatformInstallReq) GetPlatformPackage() string {
	if m != nil {
		return m.PlatformPackage
	}
	return ""
}

func (m *PlatformInstallReq) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

func (m *PlatformInstallReq) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PlatformInstallResp struct {
	Progress             *DownloadProgress `protobuf:"bytes,1,opt,name=progress,proto3" json:"progress,omitempty"`
	TaskProgress         *TaskProgress     `protobuf:"bytes,2,opt,name=task_progress,json=taskProgress,proto3" json:"task_progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PlatformInstallResp) Reset()         { *m = PlatformInstallResp{} }
func (m *PlatformInstallResp) String() string { return proto.CompactTextString(m) }
func (*PlatformInstallResp) ProtoMessage()    {}
func (*PlatformInstallResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed02318f567db566, []int{1}
}

func (m *PlatformInstallResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformInstallResp.Unmarshal(m, b)
}
func (m *PlatformInstallResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformInstallResp.Marshal(b, m, deterministic)
}
func (m *PlatformInstallResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformInstallResp.Merge(m, src)
}
func (m *PlatformInstallResp) XXX_Size() int {
	return xxx_messageInfo_PlatformInstallResp.Size(m)
}
func (m *PlatformInstallResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformInstallResp.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformInstallResp proto.InternalMessageInfo

func (m *PlatformInstallResp) GetProgress() *DownloadProgress {
	if m != nil {
		return m.Progress
	}
	return nil
}

func (m *PlatformInstallResp) GetTaskProgress() *TaskProgress {
	if m != nil {
		return m.TaskProgress
	}
	return nil
}

type PlatformDownloadReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	PlatformPackage      string    `protobuf:"bytes,2,opt,name=platform_package,json=platformPackage,proto3" json:"platform_package,omitempty"`
	Architecture         string    `protobuf:"bytes,3,opt,name=architecture,proto3" json:"architecture,omitempty"`
	Version              string    `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PlatformDownloadReq) Reset()         { *m = PlatformDownloadReq{} }
func (m *PlatformDownloadReq) String() string { return proto.CompactTextString(m) }
func (*PlatformDownloadReq) ProtoMessage()    {}
func (*PlatformDownloadReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed02318f567db566, []int{2}
}

func (m *PlatformDownloadReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformDownloadReq.Unmarshal(m, b)
}
func (m *PlatformDownloadReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformDownloadReq.Marshal(b, m, deterministic)
}
func (m *PlatformDownloadReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformDownloadReq.Merge(m, src)
}
func (m *PlatformDownloadReq) XXX_Size() int {
	return xxx_messageInfo_PlatformDownloadReq.Size(m)
}
func (m *PlatformDownloadReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformDownloadReq.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformDownloadReq proto.InternalMessageInfo

func (m *PlatformDownloadReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *PlatformDownloadReq) GetPlatformPackage() string {
	if m != nil {
		return m.PlatformPackage
	}
	return ""
}

func (m *PlatformDownloadReq) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

func (m *PlatformDownloadReq) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PlatformDownloadResp struct {
	Progress             *DownloadProgress `protobuf:"bytes,1,opt,name=progress,proto3" json:"progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PlatformDownloadResp) Reset()         { *m = PlatformDownloadResp{} }
func (m *PlatformDownloadResp) String() string { return proto.CompactTextString(m) }
func (*PlatformDownloadResp) ProtoMessage()    {}
func (*PlatformDownloadResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed02318f567db566, []int{3}
}

func (m *PlatformDownloadResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformDownloadResp.Unmarshal(m, b)
}
func (m *PlatformDownloadResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformDownloadResp.Marshal(b, m, deterministic)
}
func (m *PlatformDownloadResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformDownloadResp.Merge(m, src)
}
func (m *PlatformDownloadResp) XXX_Size() int {
	return xxx_messageInfo_PlatformDownloadResp.Size(m)
}
func (m *PlatformDownloadResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformDownloadResp.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformDownloadResp proto.InternalMessageInfo

func (m *PlatformDownloadResp) GetProgress() *DownloadProgress {
	if m != nil {
		return m.Progress
	}
	return nil
}

type PlatformUninstallReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	PlatformPackage      string    `protobuf:"bytes,2,opt,name=platform_package,json=platformPackage,proto3" json:"platform_package,omitempty"`
	Architecture         string    `protobuf:"bytes,3,opt,name=architecture,proto3" json:"architecture,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PlatformUninstallReq) Reset()         { *m = PlatformUninstallReq{} }
func (m *PlatformUninstallReq) String() string { return proto.CompactTextString(m) }
func (*PlatformUninstallReq) ProtoMessage()    {}
func (*PlatformUninstallReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed02318f567db566, []int{4}
}

func (m *PlatformUninstallReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformUninstallReq.Unmarshal(m, b)
}
func (m *PlatformUninstallReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformUninstallReq.Marshal(b, m, deterministic)
}
func (m *PlatformUninstallReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformUninstallReq.Merge(m, src)
}
func (m *PlatformUninstallReq) XXX_Size() int {
	return xxx_messageInfo_PlatformUninstallReq.Size(m)
}
func (m *PlatformUninstallReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformUninstallReq.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformUninstallReq proto.InternalMessageInfo

func (m *PlatformUninstallReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *PlatformUninstallReq) GetPlatformPackage() string {
	if m != nil {
		return m.PlatformPackage
	}
	return ""
}

func (m *PlatformUninstallReq) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

type PlatformUninstallResp struct {
	TaskProgress         *TaskProgress `protobuf:"bytes,1,opt,name=task_progress,json=taskProgress,proto3" json:"task_progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PlatformUninstallResp) Reset()         { *m = PlatformUninstallResp{} }
func (m *PlatformUninstallResp) String() string { return proto.CompactTextString(m) }
func (*PlatformUninstallResp) ProtoMessage()    {}
func (*PlatformUninstallResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed02318f567db566, []int{5}
}

func (m *PlatformUninstallResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformUninstallResp.Unmarshal(m, b)
}
func (m *PlatformUninstallResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformUninstallResp.Marshal(b, m, deterministic)
}
func (m *PlatformUninstallResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformUninstallResp.Merge(m, src)
}
func (m *PlatformUninstallResp) XXX_Size() int {
	return xxx_messageInfo_PlatformUninstallResp.Size(m)
}
func (m *PlatformUninstallResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformUninstallResp.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformUninstallResp proto.InternalMessageInfo

func (m *PlatformUninstallResp) GetTaskProgress() *TaskProgress {
	if m != nil {
		return m.TaskProgress
	}
	return nil
}

type PlatformUpgradeReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	PlatformPackage      string    `protobuf:"bytes,2,opt,name=platform_package,json=platformPackage,proto3" json:"platform_package,omitempty"`
	Architecture         string    `protobuf:"bytes,3,opt,name=architecture,proto3" json:"architecture,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PlatformUpgradeReq) Reset()         { *m = PlatformUpgradeReq{} }
func (m *PlatformUpgradeReq) String() string { return proto.CompactTextString(m) }
func (*PlatformUpgradeReq) ProtoMessage()    {}
func (*PlatformUpgradeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed02318f567db566, []int{6}
}

func (m *PlatformUpgradeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformUpgradeReq.Unmarshal(m, b)
}
func (m *PlatformUpgradeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformUpgradeReq.Marshal(b, m, deterministic)
}
func (m *PlatformUpgradeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformUpgradeReq.Merge(m, src)
}
func (m *PlatformUpgradeReq) XXX_Size() int {
	return xxx_messageInfo_PlatformUpgradeReq.Size(m)
}
func (m *PlatformUpgradeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformUpgradeReq.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformUpgradeReq proto.InternalMessageInfo

func (m *PlatformUpgradeReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *PlatformUpgradeReq) GetPlatformPackage() string {
	if m != nil {
		return m.PlatformPackage
	}
	return ""
}

func (m *PlatformUpgradeReq) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

type PlatformUpgradeResp struct {
	Progress             *DownloadProgress `protobuf:"bytes,1,opt,name=progress,proto3" json:"progress,omitempty"`
	TaskProgress         *TaskProgress     `protobuf:"bytes,2,opt,name=task_progress,json=taskProgress,proto3" json:"task_progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PlatformUpgradeResp) Reset()         { *m = PlatformUpgradeResp{} }
func (m *PlatformUpgradeResp) String() string { return proto.CompactTextString(m) }
func (*PlatformUpgradeResp) ProtoMessage()    {}
func (*PlatformUpgradeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed02318f567db566, []int{7}
}

func (m *PlatformUpgradeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformUpgradeResp.Unmarshal(m, b)
}
func (m *PlatformUpgradeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformUpgradeResp.Marshal(b, m, deterministic)
}
func (m *PlatformUpgradeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformUpgradeResp.Merge(m, src)
}
func (m *PlatformUpgradeResp) XXX_Size() int {
	return xxx_messageInfo_PlatformUpgradeResp.Size(m)
}
func (m *PlatformUpgradeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformUpgradeResp.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformUpgradeResp proto.InternalMessageInfo

func (m *PlatformUpgradeResp) GetProgress() *DownloadProgress {
	if m != nil {
		return m.Progress
	}
	return nil
}

func (m *PlatformUpgradeResp) GetTaskProgress() *TaskProgress {
	if m != nil {
		return m.TaskProgress
	}
	return nil
}

type PlatformSearchReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	SearchArgs           string    `protobuf:"bytes,2,opt,name=search_args,json=searchArgs,proto3" json:"search_args,omitempty"`
	AllVersions          bool      `protobuf:"varint,3,opt,name=all_versions,json=allVersions,proto3" json:"all_versions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PlatformSearchReq) Reset()         { *m = PlatformSearchReq{} }
func (m *PlatformSearchReq) String() string { return proto.CompactTextString(m) }
func (*PlatformSearchReq) ProtoMessage()    {}
func (*PlatformSearchReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed02318f567db566, []int{8}
}

func (m *PlatformSearchReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformSearchReq.Unmarshal(m, b)
}
func (m *PlatformSearchReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformSearchReq.Marshal(b, m, deterministic)
}
func (m *PlatformSearchReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformSearchReq.Merge(m, src)
}
func (m *PlatformSearchReq) XXX_Size() int {
	return xxx_messageInfo_PlatformSearchReq.Size(m)
}
func (m *PlatformSearchReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformSearchReq.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformSearchReq proto.InternalMessageInfo

func (m *PlatformSearchReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *PlatformSearchReq) GetSearchArgs() string {
	if m != nil {
		return m.SearchArgs
	}
	return ""
}

func (m *PlatformSearchReq) GetAllVersions() bool {
	if m != nil {
		return m.AllVersions
	}
	return false
}

type PlatformSearchResp struct {
	SearchOutput         []*Platform `protobuf:"bytes,1,rep,name=search_output,json=searchOutput,proto3" json:"search_output,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PlatformSearchResp) Reset()         { *m = PlatformSearchResp{} }
func (m *PlatformSearchResp) String() string { return proto.CompactTextString(m) }
func (*PlatformSearchResp) ProtoMessage()    {}
func (*PlatformSearchResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed02318f567db566, []int{9}
}

func (m *PlatformSearchResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformSearchResp.Unmarshal(m, b)
}
func (m *PlatformSearchResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformSearchResp.Marshal(b, m, deterministic)
}
func (m *PlatformSearchResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformSearchResp.Merge(m, src)
}
func (m *PlatformSearchResp) XXX_Size() int {
	return xxx_messageInfo_PlatformSearchResp.Size(m)
}
func (m *PlatformSearchResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformSearchResp.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformSearchResp proto.InternalMessageInfo

func (m *PlatformSearchResp) GetSearchOutput() []*Platform {
	if m != nil {
		return m.SearchOutput
	}
	return nil
}

type PlatformListReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	UpdatableOnly        bool      `protobuf:"varint,2,opt,name=updatable_only,json=updatableOnly,proto3" json:"updatable_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PlatformListReq) Reset()         { *m = PlatformListReq{} }
func (m *PlatformListReq) String() string { return proto.CompactTextString(m) }
func (*PlatformListReq) ProtoMessage()    {}
func (*PlatformListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed02318f567db566, []int{10}
}

func (m *PlatformListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformListReq.Unmarshal(m, b)
}
func (m *PlatformListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformListReq.Marshal(b, m, deterministic)
}
func (m *PlatformListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformListReq.Merge(m, src)
}
func (m *PlatformListReq) XXX_Size() int {
	return xxx_messageInfo_PlatformListReq.Size(m)
}
func (m *PlatformListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformListReq.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformListReq proto.InternalMessageInfo

func (m *PlatformListReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *PlatformListReq) GetUpdatableOnly() bool {
	if m != nil {
		return m.UpdatableOnly
	}
	return false
}

type PlatformListResp struct {
	InstalledPlatform    []*Platform `protobuf:"bytes,1,rep,name=installed_platform,json=installedPlatform,proto3" json:"installed_platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PlatformListResp) Reset()         { *m = PlatformListResp{} }
func (m *PlatformListResp) String() string { return proto.CompactTextString(m) }
func (*PlatformListResp) ProtoMessage()    {}
func (*PlatformListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed02318f567db566, []int{11}
}

func (m *PlatformListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformListResp.Unmarshal(m, b)
}
func (m *PlatformListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformListResp.Marshal(b, m, deterministic)
}
func (m *PlatformListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformListResp.Merge(m, src)
}
func (m *PlatformListResp) XXX_Size() int {
	return xxx_messageInfo_PlatformListResp.Size(m)
}
func (m *PlatformListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformListResp.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformListResp proto.InternalMessageInfo

func (m *PlatformListResp) GetInstalledPlatform() []*Platform {
	if m != nil {
		return m.InstalledPlatform
	}
	return nil
}

type Platform struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Installed            string   `protobuf:"bytes,2,opt,name=Installed,proto3" json:"Installed,omitempty"`
	Latest               string   `protobuf:"bytes,3,opt,name=Latest,proto3" json:"Latest,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Maintainer           string   `protobuf:"bytes,5,opt,name=Maintainer,proto3" json:"Maintainer,omitempty"`
	Website              string   `protobuf:"bytes,6,opt,name=Website,proto3" json:"Website,omitempty"`
	Email                string   `protobuf:"bytes,7,opt,name=Email,proto3" json:"Email,omitempty"`
	Boards               []*Board `protobuf:"bytes,8,rep,name=Boards,proto3" json:"Boards,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Platform) Reset()         { *m = Platform{} }
func (m *Platform) String() string { return proto.CompactTextString(m) }
func (*Platform) ProtoMessage()    {}
func (*Platform) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed02318f567db566, []int{12}
}

func (m *Platform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platform.Unmarshal(m, b)
}
func (m *Platform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platform.Marshal(b, m, deterministic)
}
func (m *Platform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform.Merge(m, src)
}
func (m *Platform) XXX_Size() int {
	return xxx_messageInfo_Platform.Size(m)
}
func (m *Platform) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform.DiscardUnknown(m)
}

var xxx_messageInfo_Platform proto.InternalMessageInfo

func (m *Platform) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Platform) GetInstalled() string {
	if m != nil {
		return m.Installed
	}
	return ""
}

func (m *Platform) GetLatest() string {
	if m != nil {
		return m.Latest
	}
	return ""
}

func (m *Platform) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Platform) GetMaintainer() string {
	if m != nil {
		return m.Maintainer
	}
	return ""
}

func (m *Platform) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Platform) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Platform) GetBoards() []*Board {
	if m != nil {
		return m.Boards
	}
	return nil
}

type Board struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Fqbn                 string   `protobuf:"bytes,2,opt,name=fqbn,proto3" json:"fqbn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Board) Reset()         { *m = Board{} }
func (m *Board) String() string { return proto.CompactTextString(m) }
func (*Board) ProtoMessage()    {}
func (*Board) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed02318f567db566, []int{13}
}

func (m *Board) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Board.Unmarshal(m, b)
}
func (m *Board) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Board.Marshal(b, m, deterministic)
}
func (m *Board) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Board.Merge(m, src)
}
func (m *Board) XXX_Size() int {
	return xxx_messageInfo_Board.Size(m)
}
func (m *Board) XXX_DiscardUnknown() {
	xxx_messageInfo_Board.DiscardUnknown(m)
}

var xxx_messageInfo_Board proto.InternalMessageInfo

func (m *Board) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Board) GetFqbn() string {
	if m != nil {
		return m.Fqbn
	}
	return ""
}

func init() {
	proto.RegisterType((*PlatformInstallReq)(nil), "cc.arduino.cli.commands.PlatformInstallReq")
	proto.RegisterType((*PlatformInstallResp)(nil), "cc.arduino.cli.commands.PlatformInstallResp")
	proto.RegisterType((*PlatformDownloadReq)(nil), "cc.arduino.cli.commands.PlatformDownloadReq")
	proto.RegisterType((*PlatformDownloadResp)(nil), "cc.arduino.cli.commands.PlatformDownloadResp")
	proto.RegisterType((*PlatformUninstallReq)(nil), "cc.arduino.cli.commands.PlatformUninstallReq")
	proto.RegisterType((*PlatformUninstallResp)(nil), "cc.arduino.cli.commands.PlatformUninstallResp")
	proto.RegisterType((*PlatformUpgradeReq)(nil), "cc.arduino.cli.commands.PlatformUpgradeReq")
	proto.RegisterType((*PlatformUpgradeResp)(nil), "cc.arduino.cli.commands.PlatformUpgradeResp")
	proto.RegisterType((*PlatformSearchReq)(nil), "cc.arduino.cli.commands.PlatformSearchReq")
	proto.RegisterType((*PlatformSearchResp)(nil), "cc.arduino.cli.commands.PlatformSearchResp")
	proto.RegisterType((*PlatformListReq)(nil), "cc.arduino.cli.commands.PlatformListReq")
	proto.RegisterType((*PlatformListResp)(nil), "cc.arduino.cli.commands.PlatformListResp")
	proto.RegisterType((*Platform)(nil), "cc.arduino.cli.commands.Platform")
	proto.RegisterType((*Board)(nil), "cc.arduino.cli.commands.Board")
}

func init() { proto.RegisterFile("commands/core.proto", fileDescriptor_ed02318f567db566) }

var fileDescriptor_ed02318f567db566 = []byte{
	// 628 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xdd, 0x6a, 0x14, 0x31,
	0x14, 0x66, 0xb6, 0xed, 0x76, 0x7b, 0xba, 0xfd, 0x4b, 0x5b, 0x1d, 0x44, 0x6a, 0x3b, 0x50, 0x68,
	0x91, 0xee, 0x82, 0x82, 0x77, 0x5e, 0x58, 0x5a, 0x61, 0xa5, 0xda, 0x65, 0xb4, 0x0a, 0xa2, 0x2c,
	0xd9, 0x99, 0x74, 0x1b, 0x9a, 0x49, 0xd2, 0x24, 0x63, 0xe9, 0x8b, 0xf8, 0x00, 0xe2, 0x85, 0x0f,
	0xa1, 0x6f, 0xe4, 0x43, 0xc8, 0x64, 0x92, 0xd9, 0x2d, 0x75, 0x45, 0xa4, 0x17, 0xeb, 0xd5, 0xe6,
	0x7c, 0x73, 0xce, 0x97, 0xef, 0x3b, 0xc9, 0x09, 0x0b, 0xab, 0x89, 0xc8, 0x32, 0xcc, 0x53, 0xdd,
	0x4e, 0x84, 0x22, 0x2d, 0xa9, 0x84, 0x11, 0xe8, 0x6e, 0x92, 0xb4, 0xb0, 0x4a, 0x73, 0xca, 0x45,
	0x2b, 0x61, 0xb4, 0xe5, 0x73, 0xee, 0xad, 0x8f, 0x64, 0x67, 0x99, 0xe0, 0x65, 0x7e, 0xf4, 0x3d,
	0x00, 0xd4, 0x65, 0xd8, 0x9c, 0x0a, 0x95, 0x75, 0xb8, 0x36, 0x98, 0xb1, 0x98, 0x5c, 0xa0, 0xa7,
	0xd0, 0xa0, 0x45, 0xc4, 0x13, 0x12, 0x06, 0x9b, 0xc1, 0xce, 0xfc, 0xa3, 0xad, 0xd6, 0x18, 0xe6,
	0x56, 0xc7, 0x25, 0xc6, 0x55, 0x09, 0xda, 0x85, 0x65, 0xe9, 0x48, 0x7b, 0x12, 0x27, 0xe7, 0x78,
	0x40, 0xc2, 0xda, 0x66, 0xb0, 0x33, 0x17, 0x2f, 0x79, 0xbc, 0x5b, 0xc2, 0x28, 0x82, 0x26, 0x56,
	0xc9, 0x19, 0x35, 0x24, 0x31, 0xb9, 0x22, 0xe1, 0x94, 0x4d, 0xbb, 0x86, 0xa1, 0x10, 0x66, 0x3f,
	0x11, 0xa5, 0xa9, 0xe0, 0xe1, 0xb4, 0xfd, 0xec, 0xc3, 0xe8, 0x5b, 0x00, 0xab, 0x37, 0xe4, 0x6b,
	0x89, 0x0e, 0xa1, 0x21, 0x95, 0x18, 0x28, 0xa2, 0xb5, 0xd3, 0xbf, 0x3b, 0x56, 0xff, 0x81, 0xb8,
	0xe4, 0x4c, 0xe0, 0xb4, 0xeb, 0x0a, 0xe2, 0xaa, 0x14, 0xbd, 0x80, 0x05, 0x83, 0xf5, 0x79, 0xaf,
	0xe2, 0xaa, 0x59, 0xae, 0xed, 0xb1, 0x5c, 0x6f, 0xb0, 0x3e, 0xaf, 0x78, 0x9a, 0x66, 0x24, 0x8a,
	0x7e, 0x8c, 0x48, 0xf5, 0x5b, 0xfe, 0x4f, 0xad, 0xfe, 0x08, 0x6b, 0x37, 0xe5, 0xdf, 0x5a, 0xab,
	0xa3, 0xaf, 0xc1, 0x90, 0xff, 0x84, 0xd3, 0x09, 0xbd, 0x8a, 0x51, 0x02, 0xeb, 0xbf, 0x51, 0xa9,
	0xe5, 0xcd, 0xab, 0x12, 0xfc, 0xfb, 0x55, 0xf9, 0x32, 0x32, 0x94, 0x27, 0x72, 0xa0, 0x70, 0x4a,
	0x26, 0xaf, 0x13, 0xa3, 0xa3, 0x57, 0x89, 0x9c, 0xcc, 0xd1, 0xfb, 0x1c, 0xc0, 0x8a, 0x97, 0xfa,
	0x9a, 0x14, 0x2e, 0x6e, 0xa1, 0x9d, 0x0f, 0x60, 0x5e, 0x5b, 0xae, 0x1e, 0x56, 0x03, 0xed, 0x3a,
	0x09, 0x25, 0xf4, 0x4c, 0x0d, 0x34, 0xda, 0x82, 0x26, 0x66, 0xac, 0xe7, 0xe6, 0x47, 0xdb, 0x26,
	0x36, 0xe2, 0x79, 0xcc, 0xd8, 0x5b, 0x07, 0x45, 0x1f, 0x86, 0xe7, 0xec, 0x75, 0x69, 0x89, 0x9e,
	0xc3, 0x82, 0x63, 0x16, 0xb9, 0x91, 0xb9, 0x09, 0x83, 0xcd, 0xa9, 0x3f, 0xaa, 0xf3, 0x1c, 0x71,
	0xb3, 0xac, 0x3b, 0xb6, 0x65, 0xd1, 0x25, 0x2c, 0xf9, 0x2f, 0x47, 0x54, 0x9b, 0x5b, 0xf0, 0xbc,
	0x0d, 0x8b, 0xb9, 0x4c, 0xb1, 0xc1, 0x7d, 0x46, 0x7a, 0x82, 0xb3, 0x2b, 0x6b, 0xbb, 0x11, 0x2f,
	0x54, 0xe8, 0x31, 0x67, 0x57, 0x51, 0x0a, 0xcb, 0xd7, 0x37, 0xd6, 0x12, 0x75, 0x01, 0xb9, 0x71,
	0x21, 0x69, 0xcf, 0xdf, 0xb7, 0xbf, 0x77, 0xb6, 0x52, 0x15, 0x7b, 0x28, 0xfa, 0x19, 0x40, 0xc3,
	0x07, 0x68, 0x11, 0x6a, 0x9d, 0x03, 0x6b, 0x69, 0x2e, 0xae, 0x75, 0x0e, 0xd0, 0x7d, 0x98, 0xeb,
	0xf8, 0x0a, 0x77, 0x36, 0x43, 0x00, 0xdd, 0x81, 0xfa, 0x11, 0x36, 0x44, 0x1b, 0x77, 0xb3, 0x5d,
	0x84, 0x10, 0x4c, 0xbf, 0xc2, 0x19, 0x71, 0x4f, 0x9f, 0x5d, 0xa3, 0x0d, 0x80, 0x97, 0x98, 0x72,
	0x83, 0x29, 0x27, 0x2a, 0x9c, 0x29, 0x8f, 0x79, 0x88, 0x14, 0x2f, 0xe6, 0x3b, 0xd2, 0xd7, 0xd4,
	0x90, 0xb0, 0x5e, 0xbe, 0x98, 0x2e, 0x44, 0x6b, 0x30, 0x73, 0x98, 0x61, 0xca, 0xc2, 0x59, 0x8b,
	0x97, 0x01, 0x7a, 0x02, 0xf5, 0x7d, 0x81, 0x55, 0xaa, 0xc3, 0x86, 0x35, 0xbf, 0x31, 0xd6, 0xbc,
	0x4d, 0x8b, 0x5d, 0x76, 0xd4, 0x86, 0x19, 0xbb, 0x2a, 0x44, 0xf2, 0x42, 0x64, 0x69, 0xd6, 0xae,
	0x0b, 0xec, 0xf4, 0xa2, 0xcf, 0x9d, 0x53, 0xbb, 0xde, 0xdf, 0x7b, 0xff, 0x70, 0x40, 0xcd, 0x59,
	0xde, 0x2f, 0x18, 0xdb, 0x6e, 0x07, 0xff, 0xbb, 0x97, 0x30, 0xda, 0x56, 0x32, 0x69, 0xfb, 0xdd,
	0xfa, 0x75, 0xfb, 0x87, 0xe0, 0xf1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x05, 0xe5, 0x68,
	0x57, 0x08, 0x00, 0x00,
}