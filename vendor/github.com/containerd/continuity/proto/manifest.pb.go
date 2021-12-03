// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: manifest.proto

package proto

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

// Manifest specifies the entries in a container bundle, keyed and sorted by
// path.
type Manifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource []*Resource `protobuf:"bytes,1,rep,name=resource,proto3" json:"resource,omitempty"`
}

func (x *Manifest) Reset() {
	*x = Manifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manifest) ProtoMessage() {}

func (x *Manifest) ProtoReflect() protoreflect.Message {
	mi := &file_manifest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manifest.ProtoReflect.Descriptor instead.
func (*Manifest) Descriptor() ([]byte, []int) {
	return file_manifest_proto_rawDescGZIP(), []int{0}
}

func (x *Manifest) GetResource() []*Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path specifies the path from the bundle root. If more than one
	// path is present, the entry may represent a hardlink, rather than using
	// a link target. The path format is operating system specific.
	Path []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	// Uid specifies the user id for the resource.
	Uid int64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	// Gid specifies the group id for the resource.
	Gid int64 `protobuf:"varint,3,opt,name=gid,proto3" json:"gid,omitempty"`
	// user and group are not currently used but their field numbers have been
	// reserved for future use. As such, they are marked as deprecated.
	//
	// Deprecated: Do not use.
	User string `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"` // "deprecated" stands for "reserved" here
	// Deprecated: Do not use.
	Group string `protobuf:"bytes,5,opt,name=group,proto3" json:"group,omitempty"` // "deprecated" stands for "reserved" here
	// Mode defines the file mode and permissions. We've used the same
	// bit-packing from Go's os package,
	// http://golang.org/pkg/os/#FileMode, since they've done the work of
	// creating a cross-platform layout.
	Mode uint32 `protobuf:"varint,6,opt,name=mode,proto3" json:"mode,omitempty"`
	// Size specifies the size in bytes of the resource. This is only valid
	// for regular files.
	Size uint64 `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	// Digest specifies the content digest of the target file. Only valid for
	// regular files. The strings are formatted in OCI style, i.e. <alg>:<encoded>.
	// For detailed information about the format, please refer to OCI Image Spec:
	// https://github.com/opencontainers/image-spec/blob/master/descriptor.md#digests-and-verification
	// The digests are sorted in lexical order and implementations may choose
	// which algorithms they prefer.
	Digest []string `protobuf:"bytes,8,rep,name=digest,proto3" json:"digest,omitempty"`
	// Target defines the target of a hard or soft link. Absolute links start
	// with a slash and specify the resource relative to the bundle root.
	// Relative links do not start with a slash and are relative to the
	// resource path.
	Target string `protobuf:"bytes,9,opt,name=target,proto3" json:"target,omitempty"`
	// Major specifies the major device number for character and block devices.
	Major uint64 `protobuf:"varint,10,opt,name=major,proto3" json:"major,omitempty"`
	// Minor specifies the minor device number for character and block devices.
	Minor uint64 `protobuf:"varint,11,opt,name=minor,proto3" json:"minor,omitempty"`
	// Xattr provides storage for extended attributes for the target resource.
	Xattr []*XAttr `protobuf:"bytes,12,rep,name=xattr,proto3" json:"xattr,omitempty"`
	// Ads stores one or more alternate data streams for the target resource.
	Ads []*ADSEntry `protobuf:"bytes,13,rep,name=ads,proto3" json:"ads,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_manifest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_manifest_proto_rawDescGZIP(), []int{1}
}

func (x *Resource) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *Resource) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Resource) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

// Deprecated: Do not use.
func (x *Resource) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

// Deprecated: Do not use.
func (x *Resource) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Resource) GetMode() uint32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *Resource) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Resource) GetDigest() []string {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *Resource) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Resource) GetMajor() uint64 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *Resource) GetMinor() uint64 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *Resource) GetXattr() []*XAttr {
	if x != nil {
		return x.Xattr
	}
	return nil
}

func (x *Resource) GetAds() []*ADSEntry {
	if x != nil {
		return x.Ads
	}
	return nil
}

// XAttr encodes extended attributes for a resource.
type XAttr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name specifies the attribute name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Data specifies the associated data for the attribute.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *XAttr) Reset() {
	*x = XAttr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XAttr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XAttr) ProtoMessage() {}

func (x *XAttr) ProtoReflect() protoreflect.Message {
	mi := &file_manifest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XAttr.ProtoReflect.Descriptor instead.
func (*XAttr) Descriptor() ([]byte, []int) {
	return file_manifest_proto_rawDescGZIP(), []int{2}
}

func (x *XAttr) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *XAttr) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// ADSEntry encodes information for a Windows Alternate Data Stream.
type ADSEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name specifices the stream name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Data specifies the stream data.
	// See also the description about the digest below.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Digest is a CAS representation of the stream data.
	//
	// At least one of data or digest MUST be specified, and either one of them
	// SHOULD be specified.
	//
	// How to access the actual data using the digest is implementation-specific,
	// and implementations can choose not to implement digest.
	// So, digest SHOULD be used only when the stream data is large.
	Digest string `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (x *ADSEntry) Reset() {
	*x = ADSEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADSEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADSEntry) ProtoMessage() {}

func (x *ADSEntry) ProtoReflect() protoreflect.Message {
	mi := &file_manifest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADSEntry.ProtoReflect.Descriptor instead.
func (*ADSEntry) Descriptor() ([]byte, []int) {
	return file_manifest_proto_rawDescGZIP(), []int{3}
}

func (x *ADSEntry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ADSEntry) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ADSEntry) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

var File_manifest_proto protoreflect.FileDescriptor

var file_manifest_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x08, 0x4d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0xbf, 0x02, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x05, 0x78,
	0x61, 0x74, 0x74, 0x72, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x58, 0x41, 0x74, 0x74, 0x72, 0x52, 0x05, 0x78, 0x61, 0x74, 0x74, 0x72, 0x12,
	0x21, 0x0a, 0x03, 0x61, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x44, 0x53, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x61,
	0x64, 0x73, 0x22, 0x2f, 0x0a, 0x05, 0x58, 0x41, 0x74, 0x74, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x4a, 0x0a, 0x08, 0x41, 0x44, 0x53, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75,
	0x69, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manifest_proto_rawDescOnce sync.Once
	file_manifest_proto_rawDescData = file_manifest_proto_rawDesc
)

func file_manifest_proto_rawDescGZIP() []byte {
	file_manifest_proto_rawDescOnce.Do(func() {
		file_manifest_proto_rawDescData = protoimpl.X.CompressGZIP(file_manifest_proto_rawDescData)
	})
	return file_manifest_proto_rawDescData
}

var file_manifest_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_manifest_proto_goTypes = []interface{}{
	(*Manifest)(nil), // 0: proto.Manifest
	(*Resource)(nil), // 1: proto.Resource
	(*XAttr)(nil),    // 2: proto.XAttr
	(*ADSEntry)(nil), // 3: proto.ADSEntry
}
var file_manifest_proto_depIdxs = []int32{
	1, // 0: proto.Manifest.resource:type_name -> proto.Resource
	2, // 1: proto.Resource.xattr:type_name -> proto.XAttr
	3, // 2: proto.Resource.ads:type_name -> proto.ADSEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_manifest_proto_init() }
func file_manifest_proto_init() {
	if File_manifest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manifest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manifest); i {
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
		file_manifest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_manifest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XAttr); i {
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
		file_manifest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ADSEntry); i {
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
			RawDescriptor: file_manifest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_manifest_proto_goTypes,
		DependencyIndexes: file_manifest_proto_depIdxs,
		MessageInfos:      file_manifest_proto_msgTypes,
	}.Build()
	File_manifest_proto = out.File
	file_manifest_proto_rawDesc = nil
	file_manifest_proto_goTypes = nil
	file_manifest_proto_depIdxs = nil
}
