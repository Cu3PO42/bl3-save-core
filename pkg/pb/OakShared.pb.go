// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.11.3
// source: OakShared.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type Vec3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y"`
	Z float32 `protobuf:"fixed32,3,opt,name=z,proto3" json:"z"`
}

func (x *Vec3) Reset() {
	*x = Vec3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OakShared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vec3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vec3) ProtoMessage() {}

func (x *Vec3) ProtoReflect() protoreflect.Message {
	mi := &file_OakShared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vec3.ProtoReflect.Descriptor instead.
func (*Vec3) Descriptor() ([]byte, []int) {
	return file_OakShared_proto_rawDescGZIP(), []int{0}
}

func (x *Vec3) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vec3) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Vec3) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type GameStatSaveGameData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatValue int32  `protobuf:"varint,1,opt,name=stat_value,json=statValue,proto3" json:"stat_value"`
	StatPath  string `protobuf:"bytes,2,opt,name=stat_path,json=statPath,proto3" json:"stat_path"`
}

func (x *GameStatSaveGameData) Reset() {
	*x = GameStatSaveGameData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OakShared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStatSaveGameData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStatSaveGameData) ProtoMessage() {}

func (x *GameStatSaveGameData) ProtoReflect() protoreflect.Message {
	mi := &file_OakShared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStatSaveGameData.ProtoReflect.Descriptor instead.
func (*GameStatSaveGameData) Descriptor() ([]byte, []int) {
	return file_OakShared_proto_rawDescGZIP(), []int{1}
}

func (x *GameStatSaveGameData) GetStatValue() int32 {
	if x != nil {
		return x.StatValue
	}
	return 0
}

func (x *GameStatSaveGameData) GetStatPath() string {
	if x != nil {
		return x.StatPath
	}
	return ""
}

type InventoryCategorySaveData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseCategoryDefinitionHash uint32 `protobuf:"varint,1,opt,name=base_category_definition_hash,json=baseCategoryDefinitionHash,proto3" json:"base_category_definition_hash"`
	Quantity                   int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity"`
}

func (x *InventoryCategorySaveData) Reset() {
	*x = InventoryCategorySaveData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OakShared_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryCategorySaveData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryCategorySaveData) ProtoMessage() {}

func (x *InventoryCategorySaveData) ProtoReflect() protoreflect.Message {
	mi := &file_OakShared_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryCategorySaveData.ProtoReflect.Descriptor instead.
func (*InventoryCategorySaveData) Descriptor() ([]byte, []int) {
	return file_OakShared_proto_rawDescGZIP(), []int{2}
}

func (x *InventoryCategorySaveData) GetBaseCategoryDefinitionHash() uint32 {
	if x != nil {
		return x.BaseCategoryDefinitionHash
	}
	return 0
}

func (x *InventoryCategorySaveData) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type OakSDUSaveGameData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SduLevel    int32  `protobuf:"varint,1,opt,name=sdu_level,json=sduLevel,proto3" json:"sdu_level"`
	SduDataPath string `protobuf:"bytes,2,opt,name=sdu_data_path,json=sduDataPath,proto3" json:"sdu_data_path"`
}

func (x *OakSDUSaveGameData) Reset() {
	*x = OakSDUSaveGameData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OakShared_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OakSDUSaveGameData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OakSDUSaveGameData) ProtoMessage() {}

func (x *OakSDUSaveGameData) ProtoReflect() protoreflect.Message {
	mi := &file_OakShared_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OakSDUSaveGameData.ProtoReflect.Descriptor instead.
func (*OakSDUSaveGameData) Descriptor() ([]byte, []int) {
	return file_OakShared_proto_rawDescGZIP(), []int{3}
}

func (x *OakSDUSaveGameData) GetSduLevel() int32 {
	if x != nil {
		return x.SduLevel
	}
	return 0
}

func (x *OakSDUSaveGameData) GetSduDataPath() string {
	if x != nil {
		return x.SduDataPath
	}
	return ""
}

type RegisteredDownloadableEntitlement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Consumed   uint32 `protobuf:"varint,2,opt,name=consumed,proto3" json:"consumed"`
	Registered bool   `protobuf:"varint,3,opt,name=registered,proto3" json:"registered"`
	Seen       bool   `protobuf:"varint,4,opt,name=seen,proto3" json:"seen"`
}

func (x *RegisteredDownloadableEntitlement) Reset() {
	*x = RegisteredDownloadableEntitlement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OakShared_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisteredDownloadableEntitlement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisteredDownloadableEntitlement) ProtoMessage() {}

func (x *RegisteredDownloadableEntitlement) ProtoReflect() protoreflect.Message {
	mi := &file_OakShared_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisteredDownloadableEntitlement.ProtoReflect.Descriptor instead.
func (*RegisteredDownloadableEntitlement) Descriptor() ([]byte, []int) {
	return file_OakShared_proto_rawDescGZIP(), []int{4}
}

func (x *RegisteredDownloadableEntitlement) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RegisteredDownloadableEntitlement) GetConsumed() uint32 {
	if x != nil {
		return x.Consumed
	}
	return 0
}

func (x *RegisteredDownloadableEntitlement) GetRegistered() bool {
	if x != nil {
		return x.Registered
	}
	return false
}

func (x *RegisteredDownloadableEntitlement) GetSeen() bool {
	if x != nil {
		return x.Seen
	}
	return false
}

type RegisteredDownloadableEntitlements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntitlementSourceAssetPath string                               `protobuf:"bytes,1,opt,name=entitlement_source_asset_path,json=entitlementSourceAssetPath,proto3" json:"entitlement_source_asset_path"`
	EntitlementIds             []int64                              `protobuf:"varint,2,rep,packed,name=entitlement_ids,json=entitlementIds,proto3" json:"entitlement_ids"`
	Entitlements               []*RegisteredDownloadableEntitlement `protobuf:"bytes,3,rep,name=entitlements,proto3" json:"entitlements"`
}

func (x *RegisteredDownloadableEntitlements) Reset() {
	*x = RegisteredDownloadableEntitlements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OakShared_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisteredDownloadableEntitlements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisteredDownloadableEntitlements) ProtoMessage() {}

func (x *RegisteredDownloadableEntitlements) ProtoReflect() protoreflect.Message {
	mi := &file_OakShared_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisteredDownloadableEntitlements.ProtoReflect.Descriptor instead.
func (*RegisteredDownloadableEntitlements) Descriptor() ([]byte, []int) {
	return file_OakShared_proto_rawDescGZIP(), []int{5}
}

func (x *RegisteredDownloadableEntitlements) GetEntitlementSourceAssetPath() string {
	if x != nil {
		return x.EntitlementSourceAssetPath
	}
	return ""
}

func (x *RegisteredDownloadableEntitlements) GetEntitlementIds() []int64 {
	if x != nil {
		return x.EntitlementIds
	}
	return nil
}

func (x *RegisteredDownloadableEntitlements) GetEntitlements() []*RegisteredDownloadableEntitlement {
	if x != nil {
		return x.Entitlements
	}
	return nil
}

type ChallengeStatSaveGameData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentStatValue  int32  `protobuf:"varint,1,opt,name=current_stat_value,json=currentStatValue,proto3" json:"current_stat_value"`
	ChallengeStatPath string `protobuf:"bytes,2,opt,name=challenge_stat_path,json=challengeStatPath,proto3" json:"challenge_stat_path"`
}

func (x *ChallengeStatSaveGameData) Reset() {
	*x = ChallengeStatSaveGameData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OakShared_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeStatSaveGameData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeStatSaveGameData) ProtoMessage() {}

func (x *ChallengeStatSaveGameData) ProtoReflect() protoreflect.Message {
	mi := &file_OakShared_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeStatSaveGameData.ProtoReflect.Descriptor instead.
func (*ChallengeStatSaveGameData) Descriptor() ([]byte, []int) {
	return file_OakShared_proto_rawDescGZIP(), []int{6}
}

func (x *ChallengeStatSaveGameData) GetCurrentStatValue() int32 {
	if x != nil {
		return x.CurrentStatValue
	}
	return 0
}

func (x *ChallengeStatSaveGameData) GetChallengeStatPath() string {
	if x != nil {
		return x.ChallengeStatPath
	}
	return ""
}

type OakChallengeRewardSaveGameData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeRewardClaimed bool `protobuf:"varint,1,opt,name=challenge_reward_claimed,json=challengeRewardClaimed,proto3" json:"challenge_reward_claimed"`
}

func (x *OakChallengeRewardSaveGameData) Reset() {
	*x = OakChallengeRewardSaveGameData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OakShared_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OakChallengeRewardSaveGameData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OakChallengeRewardSaveGameData) ProtoMessage() {}

func (x *OakChallengeRewardSaveGameData) ProtoReflect() protoreflect.Message {
	mi := &file_OakShared_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OakChallengeRewardSaveGameData.ProtoReflect.Descriptor instead.
func (*OakChallengeRewardSaveGameData) Descriptor() ([]byte, []int) {
	return file_OakShared_proto_rawDescGZIP(), []int{7}
}

func (x *OakChallengeRewardSaveGameData) GetChallengeRewardClaimed() bool {
	if x != nil {
		return x.ChallengeRewardClaimed
	}
	return false
}

type ChallengeSaveGameData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompletedCount         int32                             `protobuf:"varint,1,opt,name=completed_count,json=completedCount,proto3" json:"completed_count"`
	IsActive               bool                              `protobuf:"varint,2,opt,name=is_active,json=isActive,proto3" json:"is_active"`
	CurrentlyCompleted     bool                              `protobuf:"varint,3,opt,name=currently_completed,json=currentlyCompleted,proto3" json:"currently_completed"`
	CompletedProgressLevel int32                             `protobuf:"varint,4,opt,name=completed_progress_level,json=completedProgressLevel,proto3" json:"completed_progress_level"`
	ProgressCounter        int32                             `protobuf:"varint,5,opt,name=progress_counter,json=progressCounter,proto3" json:"progress_counter"`
	StatInstanceState      []*ChallengeStatSaveGameData      `protobuf:"bytes,6,rep,name=stat_instance_state,json=statInstanceState,proto3" json:"stat_instance_state"`
	ChallengeClassPath     string                            `protobuf:"bytes,7,opt,name=challenge_class_path,json=challengeClassPath,proto3" json:"challenge_class_path"`
	ChallengeRewardInfo    []*OakChallengeRewardSaveGameData `protobuf:"bytes,8,rep,name=challenge_reward_info,json=challengeRewardInfo,proto3" json:"challenge_reward_info"`
}

func (x *ChallengeSaveGameData) Reset() {
	*x = ChallengeSaveGameData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OakShared_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeSaveGameData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeSaveGameData) ProtoMessage() {}

func (x *ChallengeSaveGameData) ProtoReflect() protoreflect.Message {
	mi := &file_OakShared_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeSaveGameData.ProtoReflect.Descriptor instead.
func (*ChallengeSaveGameData) Descriptor() ([]byte, []int) {
	return file_OakShared_proto_rawDescGZIP(), []int{8}
}

func (x *ChallengeSaveGameData) GetCompletedCount() int32 {
	if x != nil {
		return x.CompletedCount
	}
	return 0
}

func (x *ChallengeSaveGameData) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *ChallengeSaveGameData) GetCurrentlyCompleted() bool {
	if x != nil {
		return x.CurrentlyCompleted
	}
	return false
}

func (x *ChallengeSaveGameData) GetCompletedProgressLevel() int32 {
	if x != nil {
		return x.CompletedProgressLevel
	}
	return 0
}

func (x *ChallengeSaveGameData) GetProgressCounter() int32 {
	if x != nil {
		return x.ProgressCounter
	}
	return 0
}

func (x *ChallengeSaveGameData) GetStatInstanceState() []*ChallengeStatSaveGameData {
	if x != nil {
		return x.StatInstanceState
	}
	return nil
}

func (x *ChallengeSaveGameData) GetChallengeClassPath() string {
	if x != nil {
		return x.ChallengeClassPath
	}
	return ""
}

func (x *ChallengeSaveGameData) GetChallengeRewardInfo() []*OakChallengeRewardSaveGameData {
	if x != nil {
		return x.ChallengeRewardInfo
	}
	return nil
}

var File_OakShared_proto protoreflect.FileDescriptor

var file_OakShared_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x4f, 0x61, 0x6b, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x4f, 0x61, 0x6b, 0x53, 0x61, 0x76, 0x65, 0x22, 0x30, 0x0a, 0x04, 0x56, 0x65,
	0x63, 0x33, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78,
	0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x0c,
	0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x7a, 0x22, 0x52, 0x0a, 0x14,
	0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x53, 0x61, 0x76, 0x65, 0x47, 0x61, 0x6d, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x22, 0x7a, 0x0a, 0x19, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x61, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a,
	0x1d, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x1a, 0x62, 0x61, 0x73, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x55, 0x0a, 0x12,
	0x4f, 0x61, 0x6b, 0x53, 0x44, 0x55, 0x53, 0x61, 0x76, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x64, 0x75, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x64, 0x75, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x22, 0x0a, 0x0d, 0x73, 0x64, 0x75, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x64, 0x75, 0x44, 0x61, 0x74, 0x61, 0x50,
	0x61, 0x74, 0x68, 0x22, 0x83, 0x01, 0x0a, 0x21, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x65, 0x65, 0x6e, 0x22, 0xe0, 0x01, 0x0a, 0x22, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x41, 0x0a, 0x1d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x12, 0x4e, 0x0a, 0x0c,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x4f, 0x61, 0x6b, 0x53, 0x61, 0x76, 0x65, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x61,
	0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x79, 0x0a, 0x19,
	0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x53, 0x61, 0x76,
	0x65, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x5a, 0x0a, 0x1e, 0x4f, 0x61, 0x6b, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x61, 0x76,
	0x65, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x18, 0x63, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x63, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x43, 0x6c, 0x61, 0x69,
	0x6d, 0x65, 0x64, 0x22, 0xd6, 0x03, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x53, 0x61, 0x76, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a,
	0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x6c, 0x79,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x29,
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x13, 0x73, 0x74, 0x61,
	0x74, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x4f, 0x61, 0x6b, 0x53, 0x61, 0x76, 0x65,
	0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x53, 0x61,
	0x76, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x11, 0x73, 0x74, 0x61, 0x74,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a,
	0x14, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x5b, 0x0a, 0x15, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x4f, 0x61, 0x6b, 0x53, 0x61, 0x76, 0x65, 0x2e, 0x4f, 0x61, 0x6b, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x61, 0x76, 0x65, 0x47,
	0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x13, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_OakShared_proto_rawDescOnce sync.Once
	file_OakShared_proto_rawDescData = file_OakShared_proto_rawDesc
)

func file_OakShared_proto_rawDescGZIP() []byte {
	file_OakShared_proto_rawDescOnce.Do(func() {
		file_OakShared_proto_rawDescData = protoimpl.X.CompressGZIP(file_OakShared_proto_rawDescData)
	})
	return file_OakShared_proto_rawDescData
}

var file_OakShared_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_OakShared_proto_goTypes = []interface{}{
	(*Vec3)(nil),                               // 0: OakSave.Vec3
	(*GameStatSaveGameData)(nil),               // 1: OakSave.GameStatSaveGameData
	(*InventoryCategorySaveData)(nil),          // 2: OakSave.InventoryCategorySaveData
	(*OakSDUSaveGameData)(nil),                 // 3: OakSave.OakSDUSaveGameData
	(*RegisteredDownloadableEntitlement)(nil),  // 4: OakSave.RegisteredDownloadableEntitlement
	(*RegisteredDownloadableEntitlements)(nil), // 5: OakSave.RegisteredDownloadableEntitlements
	(*ChallengeStatSaveGameData)(nil),          // 6: OakSave.ChallengeStatSaveGameData
	(*OakChallengeRewardSaveGameData)(nil),     // 7: OakSave.OakChallengeRewardSaveGameData
	(*ChallengeSaveGameData)(nil),              // 8: OakSave.ChallengeSaveGameData
}
var file_OakShared_proto_depIdxs = []int32{
	4, // 0: OakSave.RegisteredDownloadableEntitlements.entitlements:type_name -> OakSave.RegisteredDownloadableEntitlement
	6, // 1: OakSave.ChallengeSaveGameData.stat_instance_state:type_name -> OakSave.ChallengeStatSaveGameData
	7, // 2: OakSave.ChallengeSaveGameData.challenge_reward_info:type_name -> OakSave.OakChallengeRewardSaveGameData
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_OakShared_proto_init() }
func file_OakShared_proto_init() {
	if File_OakShared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_OakShared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vec3); i {
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
		file_OakShared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameStatSaveGameData); i {
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
		file_OakShared_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryCategorySaveData); i {
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
		file_OakShared_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OakSDUSaveGameData); i {
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
		file_OakShared_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisteredDownloadableEntitlement); i {
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
		file_OakShared_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisteredDownloadableEntitlements); i {
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
		file_OakShared_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeStatSaveGameData); i {
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
		file_OakShared_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OakChallengeRewardSaveGameData); i {
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
		file_OakShared_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeSaveGameData); i {
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
			RawDescriptor: file_OakShared_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OakShared_proto_goTypes,
		DependencyIndexes: file_OakShared_proto_depIdxs,
		MessageInfos:      file_OakShared_proto_msgTypes,
	}.Build()
	File_OakShared_proto = out.File
	file_OakShared_proto_rawDesc = nil
	file_OakShared_proto_goTypes = nil
	file_OakShared_proto_depIdxs = nil
}
