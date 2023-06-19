// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchain_security/ccv/consumer/v1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cometbft/cometbft/abci/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	_07_tendermint "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	types1 "github.com/cosmos/interchain-security/v2/x/ccv/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the CCV consumer chain genesis state
type GenesisState struct {
	Params            Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	ProviderClientId  string `protobuf:"bytes,2,opt,name=provider_client_id,json=providerClientId,proto3" json:"provider_client_id,omitempty"`
	ProviderChannelId string `protobuf:"bytes,3,opt,name=provider_channel_id,json=providerChannelId,proto3" json:"provider_channel_id,omitempty"`
	NewChain          bool   `protobuf:"varint,4,opt,name=new_chain,json=newChain,proto3" json:"new_chain,omitempty"`
	// ProviderClientState filled in on new chain, nil on restart.
	ProviderClientState *_07_tendermint.ClientState `protobuf:"bytes,5,opt,name=provider_client_state,json=providerClientState,proto3" json:"provider_client_state,omitempty"`
	// ProviderConsensusState filled in on new chain, nil on restart.
	ProviderConsensusState *_07_tendermint.ConsensusState `protobuf:"bytes,6,opt,name=provider_consensus_state,json=providerConsensusState,proto3" json:"provider_consensus_state,omitempty"`
	// MaturingPackets nil on new chain, filled in on restart.
	MaturingPackets []MaturingVSCPacket `protobuf:"bytes,7,rep,name=maturing_packets,json=maturingPackets,proto3" json:"maturing_packets"`
	// InitialValset filled in on new chain and on restart.
	InitialValSet []types.ValidatorUpdate `protobuf:"bytes,8,rep,name=initial_val_set,json=initialValSet,proto3" json:"initial_val_set"`
	// HeightToValsetUpdateId nil on new chain, filled in on restart.
	HeightToValsetUpdateId []HeightToValsetUpdateID `protobuf:"bytes,9,rep,name=height_to_valset_update_id,json=heightToValsetUpdateId,proto3" json:"height_to_valset_update_id"`
	// OutstandingDowntimes nil on new chain, filled  in on restart.
	OutstandingDowntimeSlashing []OutstandingDowntime `protobuf:"bytes,10,rep,name=outstanding_downtime_slashing,json=outstandingDowntimeSlashing,proto3" json:"outstanding_downtime_slashing"`
	// PendingConsumerPackets nil on new chain, filled in on restart.
	PendingConsumerPackets types1.ConsumerPacketDataList `protobuf:"bytes,11,opt,name=pending_consumer_packets,json=pendingConsumerPackets,proto3" json:"pending_consumer_packets"`
	// LastTransmissionBlockHeight nil on new chain, filled in on restart.
	LastTransmissionBlockHeight LastTransmissionBlockHeight `protobuf:"bytes,12,opt,name=last_transmission_block_height,json=lastTransmissionBlockHeight,proto3" json:"last_transmission_block_height"`
	PreCCV                      bool                        `protobuf:"varint,13,opt,name=preCCV,proto3" json:"preCCV,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2db73a6057a27482, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetProviderClientId() string {
	if m != nil {
		return m.ProviderClientId
	}
	return ""
}

func (m *GenesisState) GetProviderChannelId() string {
	if m != nil {
		return m.ProviderChannelId
	}
	return ""
}

func (m *GenesisState) GetNewChain() bool {
	if m != nil {
		return m.NewChain
	}
	return false
}

func (m *GenesisState) GetProviderClientState() *_07_tendermint.ClientState {
	if m != nil {
		return m.ProviderClientState
	}
	return nil
}

func (m *GenesisState) GetProviderConsensusState() *_07_tendermint.ConsensusState {
	if m != nil {
		return m.ProviderConsensusState
	}
	return nil
}

func (m *GenesisState) GetMaturingPackets() []MaturingVSCPacket {
	if m != nil {
		return m.MaturingPackets
	}
	return nil
}

func (m *GenesisState) GetInitialValSet() []types.ValidatorUpdate {
	if m != nil {
		return m.InitialValSet
	}
	return nil
}

func (m *GenesisState) GetHeightToValsetUpdateId() []HeightToValsetUpdateID {
	if m != nil {
		return m.HeightToValsetUpdateId
	}
	return nil
}

func (m *GenesisState) GetOutstandingDowntimeSlashing() []OutstandingDowntime {
	if m != nil {
		return m.OutstandingDowntimeSlashing
	}
	return nil
}

func (m *GenesisState) GetPendingConsumerPackets() types1.ConsumerPacketDataList {
	if m != nil {
		return m.PendingConsumerPackets
	}
	return types1.ConsumerPacketDataList{}
}

func (m *GenesisState) GetLastTransmissionBlockHeight() LastTransmissionBlockHeight {
	if m != nil {
		return m.LastTransmissionBlockHeight
	}
	return LastTransmissionBlockHeight{}
}

func (m *GenesisState) GetPreCCV() bool {
	if m != nil {
		return m.PreCCV
	}
	return false
}

// HeightValsetUpdateID defines the genesis information for the mapping
// of each block height to a valset update id
type HeightToValsetUpdateID struct {
	Height         uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	ValsetUpdateId uint64 `protobuf:"varint,2,opt,name=valset_update_id,json=valsetUpdateId,proto3" json:"valset_update_id,omitempty"`
}

func (m *HeightToValsetUpdateID) Reset()         { *m = HeightToValsetUpdateID{} }
func (m *HeightToValsetUpdateID) String() string { return proto.CompactTextString(m) }
func (*HeightToValsetUpdateID) ProtoMessage()    {}
func (*HeightToValsetUpdateID) Descriptor() ([]byte, []int) {
	return fileDescriptor_2db73a6057a27482, []int{1}
}
func (m *HeightToValsetUpdateID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeightToValsetUpdateID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeightToValsetUpdateID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeightToValsetUpdateID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeightToValsetUpdateID.Merge(m, src)
}
func (m *HeightToValsetUpdateID) XXX_Size() int {
	return m.Size()
}
func (m *HeightToValsetUpdateID) XXX_DiscardUnknown() {
	xxx_messageInfo_HeightToValsetUpdateID.DiscardUnknown(m)
}

var xxx_messageInfo_HeightToValsetUpdateID proto.InternalMessageInfo

func (m *HeightToValsetUpdateID) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *HeightToValsetUpdateID) GetValsetUpdateId() uint64 {
	if m != nil {
		return m.ValsetUpdateId
	}
	return 0
}

// OutstandingDowntime defines the genesis information for each validator
// flagged with an outstanding downtime slashing.
type OutstandingDowntime struct {
	ValidatorConsensusAddress string `protobuf:"bytes,1,opt,name=validator_consensus_address,json=validatorConsensusAddress,proto3" json:"validator_consensus_address,omitempty"`
}

func (m *OutstandingDowntime) Reset()         { *m = OutstandingDowntime{} }
func (m *OutstandingDowntime) String() string { return proto.CompactTextString(m) }
func (*OutstandingDowntime) ProtoMessage()    {}
func (*OutstandingDowntime) Descriptor() ([]byte, []int) {
	return fileDescriptor_2db73a6057a27482, []int{2}
}
func (m *OutstandingDowntime) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutstandingDowntime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutstandingDowntime.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutstandingDowntime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutstandingDowntime.Merge(m, src)
}
func (m *OutstandingDowntime) XXX_Size() int {
	return m.Size()
}
func (m *OutstandingDowntime) XXX_DiscardUnknown() {
	xxx_messageInfo_OutstandingDowntime.DiscardUnknown(m)
}

var xxx_messageInfo_OutstandingDowntime proto.InternalMessageInfo

func (m *OutstandingDowntime) GetValidatorConsensusAddress() string {
	if m != nil {
		return m.ValidatorConsensusAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "interchain_security.ccv.consumer.v1.GenesisState")
	proto.RegisterType((*HeightToValsetUpdateID)(nil), "interchain_security.ccv.consumer.v1.HeightToValsetUpdateID")
	proto.RegisterType((*OutstandingDowntime)(nil), "interchain_security.ccv.consumer.v1.OutstandingDowntime")
}

func init() {
	proto.RegisterFile("interchain_security/ccv/consumer/v1/genesis.proto", fileDescriptor_2db73a6057a27482)
}

var fileDescriptor_2db73a6057a27482 = []byte{
	// 786 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4b, 0x8f, 0xe3, 0x34,
	0x1c, 0x6f, 0x76, 0x87, 0xd2, 0x7a, 0x76, 0xd9, 0xc1, 0x03, 0x55, 0x68, 0x45, 0x28, 0x03, 0x87,
	0x4a, 0x40, 0xac, 0x16, 0x09, 0x21, 0x21, 0x10, 0x4c, 0x47, 0x82, 0x4a, 0x0b, 0xac, 0xda, 0xdd,
	0x22, 0xed, 0x25, 0x72, 0x1d, 0x93, 0x58, 0x9b, 0xd8, 0x91, 0xed, 0x64, 0xd8, 0x03, 0x17, 0xae,
	0x5c, 0xf8, 0x58, 0x7b, 0x9c, 0x23, 0x27, 0x84, 0x66, 0xbe, 0x03, 0x67, 0x14, 0xdb, 0xe9, 0x83,
	0xe9, 0x88, 0x9e, 0x12, 0xe7, 0xff, 0x7b, 0xfc, 0x1f, 0x8e, 0x0d, 0xc6, 0x8c, 0x6b, 0x2a, 0x49,
	0x8a, 0x19, 0x8f, 0x14, 0x25, 0xa5, 0x64, 0xfa, 0x25, 0x22, 0xa4, 0x42, 0x44, 0x70, 0x55, 0xe6,
	0x54, 0xa2, 0x6a, 0x8c, 0x12, 0xca, 0xa9, 0x62, 0x2a, 0x2c, 0xa4, 0xd0, 0x02, 0x7e, 0xb0, 0x87,
	0x12, 0x12, 0x52, 0x85, 0x0d, 0x25, 0xac, 0xc6, 0xfd, 0x0f, 0xef, 0xd2, 0xad, 0xc6, 0xf5, 0xc3,
	0x4a, 0xf5, 0x27, 0x87, 0xb8, 0xaf, 0x65, 0x2d, 0x67, 0xa0, 0x29, 0x8f, 0xa9, 0xcc, 0x19, 0xd7,
	0x08, 0xaf, 0x08, 0x43, 0xfa, 0x65, 0x41, 0x5d, 0x6e, 0x7d, 0xc4, 0x56, 0x04, 0x65, 0x2c, 0x49,
	0x35, 0xc9, 0x18, 0xe5, 0x5a, 0xa1, 0x2d, 0x74, 0x35, 0xde, 0x5a, 0x39, 0xc2, 0xfb, 0x35, 0x81,
	0x08, 0x49, 0x11, 0x49, 0x31, 0xe7, 0x34, 0x33, 0x8e, 0xf6, 0xd5, 0x41, 0x82, 0x44, 0x88, 0x24,
	0xa3, 0xc8, 0xac, 0x56, 0xe5, 0xcf, 0x28, 0x2e, 0x25, 0xd6, 0x4c, 0x70, 0x17, 0x7f, 0x2b, 0x11,
	0x89, 0x30, 0xaf, 0xa8, 0x7e, 0xb3, 0x5f, 0xcf, 0xfe, 0xe9, 0x80, 0x07, 0xdf, 0xda, 0xbe, 0x2d,
	0x34, 0xd6, 0x14, 0xce, 0x40, 0xbb, 0xc0, 0x12, 0xe7, 0xca, 0xf7, 0x86, 0xde, 0xe8, 0x78, 0xf2,
	0x51, 0x78, 0x40, 0x1f, 0xc3, 0x27, 0x86, 0x72, 0x7e, 0xf4, 0xea, 0xaf, 0xf7, 0x5a, 0x73, 0x27,
	0x00, 0x3f, 0x06, 0xb0, 0x90, 0xa2, 0x62, 0x31, 0x95, 0x91, 0xad, 0x33, 0x62, 0xb1, 0x7f, 0x6f,
	0xe8, 0x8d, 0xba, 0xf3, 0x93, 0x26, 0x32, 0x35, 0x81, 0x59, 0x0c, 0x43, 0x70, 0xba, 0x41, 0xdb,
	0xca, 0x6a, 0xf8, 0x7d, 0x03, 0x7f, 0x73, 0x0d, 0xb7, 0x91, 0x59, 0x0c, 0x07, 0xa0, 0xcb, 0xe9,
	0x65, 0x64, 0x12, 0xf3, 0x8f, 0x86, 0xde, 0xa8, 0x33, 0xef, 0x70, 0x7a, 0x39, 0xad, 0xd7, 0x30,
	0x02, 0x6f, 0xff, 0xd7, 0x5a, 0xd5, 0xe5, 0xf9, 0xaf, 0x35, 0x45, 0xad, 0x48, 0xb8, 0x3d, 0x80,
	0x70, 0xab, 0xe5, 0xd5, 0x38, 0xb4, 0x59, 0x99, 0x8e, 0xcc, 0x4f, 0x77, 0x53, 0xb5, 0x6d, 0x4a,
	0x81, 0xbf, 0x31, 0x10, 0x5c, 0x51, 0xae, 0x4a, 0xe5, 0x3c, 0xda, 0xc6, 0x23, 0xfc, 0x5f, 0x8f,
	0x86, 0x66, 0x6d, 0x7a, 0x6b, 0x9b, 0x9d, 0xef, 0x30, 0x01, 0x27, 0x39, 0xd6, 0xa5, 0x64, 0x3c,
	0x89, 0x0a, 0x4c, 0x5e, 0x50, 0xad, 0xfc, 0xd7, 0x87, 0xf7, 0x47, 0xc7, 0x93, 0xcf, 0x0e, 0x1a,
	0xcd, 0xf7, 0x8e, 0xbc, 0x5c, 0x4c, 0x9f, 0x18, 0xba, 0x9b, 0xd2, 0xa3, 0x46, 0xd5, 0x7e, 0x55,
	0xf0, 0x07, 0xf0, 0x88, 0x71, 0xa6, 0x19, 0xce, 0xa2, 0x0a, 0x67, 0x91, 0xa2, 0xda, 0xef, 0x18,
	0x9f, 0xe1, 0x76, 0xe2, 0xf5, 0x5e, 0x0e, 0x97, 0x38, 0x63, 0x31, 0xd6, 0x42, 0x3e, 0x2b, 0x62,
	0xac, 0xa9, 0x53, 0x7c, 0xe8, 0xe8, 0x4b, 0x9c, 0x2d, 0xa8, 0x86, 0xbf, 0x82, 0x7e, 0x4a, 0xeb,
	0xf2, 0x23, 0x2d, 0x6a, 0x45, 0x45, 0x75, 0x54, 0x1a, 0x7c, 0x3d, 0xd7, 0xae, 0x91, 0xfe, 0xe2,
	0xa0, 0x12, 0xbe, 0x33, 0x32, 0x4f, 0xc5, 0xd2, 0x88, 0x58, 0xcf, 0xd9, 0x85, 0x73, 0xed, 0xa5,
	0xfb, 0xa2, 0x31, 0xfc, 0xcd, 0x03, 0xef, 0x8a, 0x52, 0x2b, 0x8d, 0x79, 0x5c, 0xf7, 0x2e, 0x16,
	0x97, 0x5c, 0xb3, 0x9c, 0x46, 0x2a, 0xc3, 0x2a, 0x65, 0x3c, 0xf1, 0x81, 0x49, 0xe1, 0xf3, 0x83,
	0x52, 0xf8, 0x71, 0xa3, 0x74, 0xe1, 0x84, 0x9c, 0xff, 0x40, 0xdc, 0x0e, 0x2d, 0x9c, 0x05, 0x94,
	0xc0, 0x2f, 0xa8, 0xf5, 0x6f, 0xd4, 0xd6, 0x43, 0x3c, 0x36, 0xdb, 0x64, 0x72, 0xa7, 0xbd, 0xdb,
	0x22, 0x35, 0xc7, 0x8e, 0xe8, 0x02, 0x6b, 0xfc, 0x98, 0xa9, 0x66, 0x80, 0x3d, 0xa7, 0xbc, 0x0b,
	0x52, 0xf0, 0x77, 0x0f, 0x04, 0x19, 0x56, 0x3a, 0xd2, 0x12, 0x73, 0x95, 0x33, 0xa5, 0x98, 0xe0,
	0xd1, 0x2a, 0x13, 0xe4, 0x45, 0x64, 0x7b, 0xe5, 0x3f, 0x30, 0xd6, 0x5f, 0x1f, 0x54, 0xf9, 0x63,
	0xac, 0xf4, 0xd3, 0x2d, 0xa5, 0xf3, 0x5a, 0xc8, 0x4e, 0xa4, 0xe9, 0x40, 0x76, 0x37, 0x04, 0xf6,
	0x40, 0xbb, 0x90, 0x74, 0x3a, 0x5d, 0xfa, 0x0f, 0xcd, 0x3f, 0xea, 0x56, 0x67, 0xcf, 0x41, 0x6f,
	0xff, 0x58, 0x6b, 0x86, 0x4b, 0xb3, 0x3e, 0x81, 0x8e, 0xe6, 0x6e, 0x05, 0x47, 0xe0, 0xe4, 0xd6,
	0x2e, 0xba, 0x67, 0x10, 0x6f, 0x54, 0x3b, 0xa3, 0x3f, 0x7b, 0x06, 0x4e, 0xf7, 0xcc, 0x0b, 0x7e,
	0x05, 0x06, 0x55, 0xb3, 0x71, 0xb7, 0x7e, 0x5a, 0x1c, 0xc7, 0x92, 0x2a, 0x7b, 0xde, 0x75, 0xe7,
	0xef, 0xac, 0x21, 0xeb, 0xff, 0xf0, 0x1b, 0x0b, 0x38, 0xff, 0xe9, 0xd5, 0x75, 0xe0, 0x5d, 0x5d,
	0x07, 0xde, 0xdf, 0xd7, 0x81, 0xf7, 0xc7, 0x4d, 0xd0, 0xba, 0xba, 0x09, 0x5a, 0x7f, 0xde, 0x04,
	0xad, 0xe7, 0x5f, 0x26, 0x4c, 0xa7, 0xe5, 0x2a, 0x24, 0x22, 0x47, 0x44, 0xa8, 0x5c, 0x28, 0xb4,
	0x69, 0xed, 0x27, 0xeb, 0x2b, 0xa3, 0x9a, 0xa0, 0x5f, 0x76, 0xef, 0x0d, 0x73, 0x29, 0xac, 0xda,
	0xe6, 0x2c, 0xfe, 0xf4, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0x37, 0xab, 0xfd, 0xe6, 0x06,
	0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PreCCV {
		i--
		if m.PreCCV {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x68
	}
	{
		size, err := m.LastTransmissionBlockHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size, err := m.PendingConsumerPackets.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	if len(m.OutstandingDowntimeSlashing) > 0 {
		for iNdEx := len(m.OutstandingDowntimeSlashing) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OutstandingDowntimeSlashing[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.HeightToValsetUpdateId) > 0 {
		for iNdEx := len(m.HeightToValsetUpdateId) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HeightToValsetUpdateId[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.InitialValSet) > 0 {
		for iNdEx := len(m.InitialValSet) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialValSet[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.MaturingPackets) > 0 {
		for iNdEx := len(m.MaturingPackets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaturingPackets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.ProviderConsensusState != nil {
		{
			size, err := m.ProviderConsensusState.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.ProviderClientState != nil {
		{
			size, err := m.ProviderClientState.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.NewChain {
		i--
		if m.NewChain {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.ProviderChannelId) > 0 {
		i -= len(m.ProviderChannelId)
		copy(dAtA[i:], m.ProviderChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ProviderChannelId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ProviderClientId) > 0 {
		i -= len(m.ProviderClientId)
		copy(dAtA[i:], m.ProviderClientId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ProviderClientId)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *HeightToValsetUpdateID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeightToValsetUpdateID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeightToValsetUpdateID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ValsetUpdateId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ValsetUpdateId))
		i--
		dAtA[i] = 0x10
	}
	if m.Height != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OutstandingDowntime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutstandingDowntime) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutstandingDowntime) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorConsensusAddress) > 0 {
		i -= len(m.ValidatorConsensusAddress)
		copy(dAtA[i:], m.ValidatorConsensusAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorConsensusAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.ProviderClientId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ProviderChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.NewChain {
		n += 2
	}
	if m.ProviderClientState != nil {
		l = m.ProviderClientState.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.ProviderConsensusState != nil {
		l = m.ProviderConsensusState.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.MaturingPackets) > 0 {
		for _, e := range m.MaturingPackets {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.InitialValSet) > 0 {
		for _, e := range m.InitialValSet {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.HeightToValsetUpdateId) > 0 {
		for _, e := range m.HeightToValsetUpdateId {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.OutstandingDowntimeSlashing) > 0 {
		for _, e := range m.OutstandingDowntimeSlashing {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.PendingConsumerPackets.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.LastTransmissionBlockHeight.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.PreCCV {
		n += 2
	}
	return n
}

func (m *HeightToValsetUpdateID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovGenesis(uint64(m.Height))
	}
	if m.ValsetUpdateId != 0 {
		n += 1 + sovGenesis(uint64(m.ValsetUpdateId))
	}
	return n
}

func (m *OutstandingDowntime) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorConsensusAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewChain", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NewChain = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderClientState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ProviderClientState == nil {
				m.ProviderClientState = &_07_tendermint.ClientState{}
			}
			if err := m.ProviderClientState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderConsensusState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ProviderConsensusState == nil {
				m.ProviderConsensusState = &_07_tendermint.ConsensusState{}
			}
			if err := m.ProviderConsensusState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaturingPackets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaturingPackets = append(m.MaturingPackets, MaturingVSCPacket{})
			if err := m.MaturingPackets[len(m.MaturingPackets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialValSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialValSet = append(m.InitialValSet, types.ValidatorUpdate{})
			if err := m.InitialValSet[len(m.InitialValSet)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeightToValsetUpdateId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeightToValsetUpdateId = append(m.HeightToValsetUpdateId, HeightToValsetUpdateID{})
			if err := m.HeightToValsetUpdateId[len(m.HeightToValsetUpdateId)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutstandingDowntimeSlashing", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutstandingDowntimeSlashing = append(m.OutstandingDowntimeSlashing, OutstandingDowntime{})
			if err := m.OutstandingDowntimeSlashing[len(m.OutstandingDowntimeSlashing)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingConsumerPackets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PendingConsumerPackets.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTransmissionBlockHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastTransmissionBlockHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreCCV", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PreCCV = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HeightToValsetUpdateID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HeightToValsetUpdateID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeightToValsetUpdateID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValsetUpdateId", wireType)
			}
			m.ValsetUpdateId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValsetUpdateId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OutstandingDowntime) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OutstandingDowntime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutstandingDowntime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorConsensusAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorConsensusAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
