// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: flow/legacy/entities/account.proto

package entities

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

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance uint64        `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Code    []byte        `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Keys    []*AccountKey `protobuf:"bytes,4,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_legacy_entities_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_flow_legacy_entities_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_flow_legacy_entities_account_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Account) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Account) GetCode() []byte {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *Account) GetKeys() []*AccountKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

type AccountKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index          uint32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	PublicKey      []byte `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	SignAlgo       uint32 `protobuf:"varint,3,opt,name=sign_algo,json=signAlgo,proto3" json:"sign_algo,omitempty"`
	HashAlgo       uint32 `protobuf:"varint,4,opt,name=hash_algo,json=hashAlgo,proto3" json:"hash_algo,omitempty"`
	Weight         uint32 `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	SequenceNumber uint32 `protobuf:"varint,6,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
}

func (x *AccountKey) Reset() {
	*x = AccountKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_legacy_entities_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountKey) ProtoMessage() {}

func (x *AccountKey) ProtoReflect() protoreflect.Message {
	mi := &file_flow_legacy_entities_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountKey.ProtoReflect.Descriptor instead.
func (*AccountKey) Descriptor() ([]byte, []int) {
	return file_flow_legacy_entities_account_proto_rawDescGZIP(), []int{1}
}

func (x *AccountKey) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *AccountKey) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *AccountKey) GetSignAlgo() uint32 {
	if x != nil {
		return x.SignAlgo
	}
	return 0
}

func (x *AccountKey) GetHashAlgo() uint32 {
	if x != nil {
		return x.HashAlgo
	}
	return 0
}

func (x *AccountKey) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *AccountKey) GetSequenceNumber() uint32 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}

var File_flow_legacy_entities_account_proto protoreflect.FileDescriptor

var file_flow_legacy_entities_account_proto_rawDesc = []byte{
	0x0a, 0x22, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x7b,
	0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x28, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x0a,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x12, 0x1b, 0x0a, 0x09,
	0x68, 0x61, 0x73, 0x68, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x68, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x5e, 0x0a, 0x23, 0x6f, 0x72,
	0x67, 0x2e, 0x6f, 0x6e, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6e,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63,
	0x79, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_flow_legacy_entities_account_proto_rawDescOnce sync.Once
	file_flow_legacy_entities_account_proto_rawDescData = file_flow_legacy_entities_account_proto_rawDesc
)

func file_flow_legacy_entities_account_proto_rawDescGZIP() []byte {
	file_flow_legacy_entities_account_proto_rawDescOnce.Do(func() {
		file_flow_legacy_entities_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_flow_legacy_entities_account_proto_rawDescData)
	})
	return file_flow_legacy_entities_account_proto_rawDescData
}

var file_flow_legacy_entities_account_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_flow_legacy_entities_account_proto_goTypes = []interface{}{
	(*Account)(nil),    // 0: entities.Account
	(*AccountKey)(nil), // 1: entities.AccountKey
}
var file_flow_legacy_entities_account_proto_depIdxs = []int32{
	1, // 0: entities.Account.keys:type_name -> entities.AccountKey
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_flow_legacy_entities_account_proto_init() }
func file_flow_legacy_entities_account_proto_init() {
	if File_flow_legacy_entities_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flow_legacy_entities_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_flow_legacy_entities_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountKey); i {
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
			RawDescriptor: file_flow_legacy_entities_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flow_legacy_entities_account_proto_goTypes,
		DependencyIndexes: file_flow_legacy_entities_account_proto_depIdxs,
		MessageInfos:      file_flow_legacy_entities_account_proto_msgTypes,
	}.Build()
	File_flow_legacy_entities_account_proto = out.File
	file_flow_legacy_entities_account_proto_rawDesc = nil
	file_flow_legacy_entities_account_proto_goTypes = nil
	file_flow_legacy_entities_account_proto_depIdxs = nil
}
