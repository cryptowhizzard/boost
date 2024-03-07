// Code generated by github.com/filecoin-project/boost/gen/api. DO NOT EDIT.

package api

import (
	"context"
	"errors"
	"time"

	smtypes "github.com/filecoin-project/boost/storagemarket/types"
	"github.com/filecoin-project/boost/storagemarket/types/legacytypes"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/go-state-types/crypto"
	lotus_api "github.com/filecoin-project/lotus/api"
	apitypes "github.com/filecoin-project/lotus/api/types"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	metrics "github.com/libp2p/go-libp2p/core/metrics"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multihash"
)

var ErrNotSupported = errors.New("method not supported")

type BoostStruct struct {
	CommonStruct

	NetStruct

	Internal struct {
		BlockstoreGet func(p0 context.Context, p1 cid.Cid) ([]byte, error) `perm:"read"`

		BlockstoreGetSize func(p0 context.Context, p1 cid.Cid) (int, error) `perm:"read"`

		BlockstoreHas func(p0 context.Context, p1 cid.Cid) (bool, error) `perm:"read"`

		BoostDeal func(p0 context.Context, p1 uuid.UUID) (*smtypes.ProviderDealState, error) `perm:"admin"`

		BoostDealBySignedProposalCid func(p0 context.Context, p1 cid.Cid) (*smtypes.ProviderDealState, error) `perm:"admin"`

		BoostDirectDeal func(p0 context.Context, p1 smtypes.DirectDealParams) (*ProviderDealRejectionInfo, error) `perm:"admin"`

		BoostDummyDeal func(p0 context.Context, p1 smtypes.DealParams) (*ProviderDealRejectionInfo, error) `perm:"admin"`

		BoostIndexerAnnounceAllDeals func(p0 context.Context) error `perm:"admin"`

		BoostIndexerAnnounceDeal func(p0 context.Context, p1 *smtypes.ProviderDealState) (cid.Cid, error) `perm:"admin"`

		BoostIndexerAnnounceDealRemoved func(p0 context.Context, p1 cid.Cid) (cid.Cid, error) `perm:"admin"`

		BoostIndexerAnnounceLatest func(p0 context.Context) (cid.Cid, error) `perm:"admin"`

		BoostIndexerAnnounceLatestHttp func(p0 context.Context, p1 []string) (cid.Cid, error) `perm:"admin"`

		BoostIndexerAnnounceLegacyDeal func(p0 context.Context, p1 cid.Cid) (cid.Cid, error) `perm:"admin"`

		BoostIndexerListMultihashes func(p0 context.Context, p1 []byte) ([]multihash.Multihash, error) `perm:"admin"`

		BoostLegacyDealByProposalCid func(p0 context.Context, p1 cid.Cid) (legacytypes.MinerDeal, error) `perm:"admin"`

		BoostOfflineDealWithData func(p0 context.Context, p1 uuid.UUID, p2 string, p3 bool) (*ProviderDealRejectionInfo, error) `perm:"admin"`

		OnlineBackup func(p0 context.Context, p1 string) error `perm:"admin"`

		PdBuildIndexForPieceCid func(p0 context.Context, p1 cid.Cid) error `perm:"admin"`

		PdCleanup func(p0 context.Context) error `perm:"admin"`

		PdRemoveDealForPiece func(p0 context.Context, p1 cid.Cid, p2 string) error `perm:"admin"`
	}
}

type BoostStub struct {
	CommonStub

	NetStub
}

type ChainIOStruct struct {
	Internal struct {
		ChainHasObj func(p0 context.Context, p1 cid.Cid) (bool, error) ``

		ChainReadObj func(p0 context.Context, p1 cid.Cid) ([]byte, error) ``
	}
}

type ChainIOStub struct {
}

type CommonStruct struct {
	Internal struct {
		AuthNew func(p0 context.Context, p1 []auth.Permission) ([]byte, error) `perm:"admin"`

		AuthVerify func(p0 context.Context, p1 string) ([]auth.Permission, error) `perm:"read"`

		Discover func(p0 context.Context) (apitypes.OpenRPCDocument, error) `perm:"read"`

		LogList func(p0 context.Context) ([]string, error) `perm:"write"`

		LogSetLevel func(p0 context.Context, p1 string, p2 string) error `perm:"write"`
	}
}

type CommonStub struct {
}

type CommonNetStruct struct {
	CommonStruct

	NetStruct

	Internal struct {
	}
}

type CommonNetStub struct {
	CommonStub

	NetStub
}

type NetStruct struct {
	Internal struct {
		ID func(p0 context.Context) (peer.ID, error) `perm:"read"`

		NetAddrsListen func(p0 context.Context) (peer.AddrInfo, error) `perm:"read"`

		NetAgentVersion func(p0 context.Context, p1 peer.ID) (string, error) `perm:"read"`

		NetAutoNatStatus func(p0 context.Context) (lotus_api.NatInfo, error) `perm:"read"`

		NetBandwidthStats func(p0 context.Context) (metrics.Stats, error) `perm:"read"`

		NetBandwidthStatsByPeer func(p0 context.Context) (map[string]metrics.Stats, error) `perm:"read"`

		NetBandwidthStatsByProtocol func(p0 context.Context) (map[protocol.ID]metrics.Stats, error) `perm:"read"`

		NetBlockAdd func(p0 context.Context, p1 lotus_api.NetBlockList) error `perm:"admin"`

		NetBlockList func(p0 context.Context) (lotus_api.NetBlockList, error) `perm:"read"`

		NetBlockRemove func(p0 context.Context, p1 lotus_api.NetBlockList) error `perm:"admin"`

		NetConnect func(p0 context.Context, p1 peer.AddrInfo) error `perm:"write"`

		NetConnectedness func(p0 context.Context, p1 peer.ID) (network.Connectedness, error) `perm:"read"`

		NetDisconnect func(p0 context.Context, p1 peer.ID) error `perm:"write"`

		NetFindPeer func(p0 context.Context, p1 peer.ID) (peer.AddrInfo, error) `perm:"read"`

		NetLimit func(p0 context.Context, p1 string) (lotus_api.NetLimit, error) `perm:"read"`

		NetPeerInfo func(p0 context.Context, p1 peer.ID) (*lotus_api.ExtendedPeerInfo, error) `perm:"read"`

		NetPeers func(p0 context.Context) ([]peer.AddrInfo, error) `perm:"read"`

		NetPing func(p0 context.Context, p1 peer.ID) (time.Duration, error) `perm:"read"`

		NetProtectAdd func(p0 context.Context, p1 []peer.ID) error `perm:"admin"`

		NetProtectList func(p0 context.Context) ([]peer.ID, error) `perm:"read"`

		NetProtectRemove func(p0 context.Context, p1 []peer.ID) error `perm:"admin"`

		NetPubsubScores func(p0 context.Context) ([]lotus_api.PubsubScore, error) `perm:"read"`

		NetSetLimit func(p0 context.Context, p1 string, p2 lotus_api.NetLimit) error `perm:"admin"`

		NetStat func(p0 context.Context, p1 string) (lotus_api.NetStat, error) `perm:"read"`
	}
}

type NetStub struct {
}

type WalletStruct struct {
	Internal struct {
		WalletDelete func(p0 context.Context, p1 address.Address) error `perm:"admin"`

		WalletExport func(p0 context.Context, p1 address.Address) (*types.KeyInfo, error) `perm:"admin"`

		WalletHas func(p0 context.Context, p1 address.Address) (bool, error) `perm:"admin"`

		WalletImport func(p0 context.Context, p1 *types.KeyInfo) (address.Address, error) `perm:"admin"`

		WalletList func(p0 context.Context) ([]address.Address, error) `perm:"admin"`

		WalletNew func(p0 context.Context, p1 types.KeyType) (address.Address, error) `perm:"admin"`

		WalletSign func(p0 context.Context, p1 address.Address, p2 []byte) (*crypto.Signature, error) `perm:"admin"`
	}
}

type WalletStub struct {
}

func (s *BoostStruct) BlockstoreGet(p0 context.Context, p1 cid.Cid) ([]byte, error) {
	if s.Internal.BlockstoreGet == nil {
		return *new([]byte), ErrNotSupported
	}
	return s.Internal.BlockstoreGet(p0, p1)
}

func (s *BoostStub) BlockstoreGet(p0 context.Context, p1 cid.Cid) ([]byte, error) {
	return *new([]byte), ErrNotSupported
}

func (s *BoostStruct) BlockstoreGetSize(p0 context.Context, p1 cid.Cid) (int, error) {
	if s.Internal.BlockstoreGetSize == nil {
		return 0, ErrNotSupported
	}
	return s.Internal.BlockstoreGetSize(p0, p1)
}

func (s *BoostStub) BlockstoreGetSize(p0 context.Context, p1 cid.Cid) (int, error) {
	return 0, ErrNotSupported
}

func (s *BoostStruct) BlockstoreHas(p0 context.Context, p1 cid.Cid) (bool, error) {
	if s.Internal.BlockstoreHas == nil {
		return false, ErrNotSupported
	}
	return s.Internal.BlockstoreHas(p0, p1)
}

func (s *BoostStub) BlockstoreHas(p0 context.Context, p1 cid.Cid) (bool, error) {
	return false, ErrNotSupported
}

func (s *BoostStruct) BoostDeal(p0 context.Context, p1 uuid.UUID) (*smtypes.ProviderDealState, error) {
	if s.Internal.BoostDeal == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.BoostDeal(p0, p1)
}

func (s *BoostStub) BoostDeal(p0 context.Context, p1 uuid.UUID) (*smtypes.ProviderDealState, error) {
	return nil, ErrNotSupported
}

func (s *BoostStruct) BoostDealBySignedProposalCid(p0 context.Context, p1 cid.Cid) (*smtypes.ProviderDealState, error) {
	if s.Internal.BoostDealBySignedProposalCid == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.BoostDealBySignedProposalCid(p0, p1)
}

func (s *BoostStub) BoostDealBySignedProposalCid(p0 context.Context, p1 cid.Cid) (*smtypes.ProviderDealState, error) {
	return nil, ErrNotSupported
}

func (s *BoostStruct) BoostDirectDeal(p0 context.Context, p1 smtypes.DirectDealParams) (*ProviderDealRejectionInfo, error) {
	if s.Internal.BoostDirectDeal == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.BoostDirectDeal(p0, p1)
}

func (s *BoostStub) BoostDirectDeal(p0 context.Context, p1 smtypes.DirectDealParams) (*ProviderDealRejectionInfo, error) {
	return nil, ErrNotSupported
}

func (s *BoostStruct) BoostDummyDeal(p0 context.Context, p1 smtypes.DealParams) (*ProviderDealRejectionInfo, error) {
	if s.Internal.BoostDummyDeal == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.BoostDummyDeal(p0, p1)
}

func (s *BoostStub) BoostDummyDeal(p0 context.Context, p1 smtypes.DealParams) (*ProviderDealRejectionInfo, error) {
	return nil, ErrNotSupported
}

func (s *BoostStruct) BoostIndexerAnnounceAllDeals(p0 context.Context) error {
	if s.Internal.BoostIndexerAnnounceAllDeals == nil {
		return ErrNotSupported
	}
	return s.Internal.BoostIndexerAnnounceAllDeals(p0)
}

func (s *BoostStub) BoostIndexerAnnounceAllDeals(p0 context.Context) error {
	return ErrNotSupported
}

func (s *BoostStruct) BoostIndexerAnnounceDeal(p0 context.Context, p1 *smtypes.ProviderDealState) (cid.Cid, error) {
	if s.Internal.BoostIndexerAnnounceDeal == nil {
		return *new(cid.Cid), ErrNotSupported
	}
	return s.Internal.BoostIndexerAnnounceDeal(p0, p1)
}

func (s *BoostStub) BoostIndexerAnnounceDeal(p0 context.Context, p1 *smtypes.ProviderDealState) (cid.Cid, error) {
	return *new(cid.Cid), ErrNotSupported
}

func (s *BoostStruct) BoostIndexerAnnounceDealRemoved(p0 context.Context, p1 cid.Cid) (cid.Cid, error) {
	if s.Internal.BoostIndexerAnnounceDealRemoved == nil {
		return *new(cid.Cid), ErrNotSupported
	}
	return s.Internal.BoostIndexerAnnounceDealRemoved(p0, p1)
}

func (s *BoostStub) BoostIndexerAnnounceDealRemoved(p0 context.Context, p1 cid.Cid) (cid.Cid, error) {
	return *new(cid.Cid), ErrNotSupported
}

func (s *BoostStruct) BoostIndexerAnnounceLatest(p0 context.Context) (cid.Cid, error) {
	if s.Internal.BoostIndexerAnnounceLatest == nil {
		return *new(cid.Cid), ErrNotSupported
	}
	return s.Internal.BoostIndexerAnnounceLatest(p0)
}

func (s *BoostStub) BoostIndexerAnnounceLatest(p0 context.Context) (cid.Cid, error) {
	return *new(cid.Cid), ErrNotSupported
}

func (s *BoostStruct) BoostIndexerAnnounceLatestHttp(p0 context.Context, p1 []string) (cid.Cid, error) {
	if s.Internal.BoostIndexerAnnounceLatestHttp == nil {
		return *new(cid.Cid), ErrNotSupported
	}
	return s.Internal.BoostIndexerAnnounceLatestHttp(p0, p1)
}

func (s *BoostStub) BoostIndexerAnnounceLatestHttp(p0 context.Context, p1 []string) (cid.Cid, error) {
	return *new(cid.Cid), ErrNotSupported
}

func (s *BoostStruct) BoostIndexerAnnounceLegacyDeal(p0 context.Context, p1 cid.Cid) (cid.Cid, error) {
	if s.Internal.BoostIndexerAnnounceLegacyDeal == nil {
		return *new(cid.Cid), ErrNotSupported
	}
	return s.Internal.BoostIndexerAnnounceLegacyDeal(p0, p1)
}

func (s *BoostStub) BoostIndexerAnnounceLegacyDeal(p0 context.Context, p1 cid.Cid) (cid.Cid, error) {
	return *new(cid.Cid), ErrNotSupported
}

func (s *BoostStruct) BoostIndexerListMultihashes(p0 context.Context, p1 []byte) ([]multihash.Multihash, error) {
	if s.Internal.BoostIndexerListMultihashes == nil {
		return *new([]multihash.Multihash), ErrNotSupported
	}
	return s.Internal.BoostIndexerListMultihashes(p0, p1)
}

func (s *BoostStub) BoostIndexerListMultihashes(p0 context.Context, p1 []byte) ([]multihash.Multihash, error) {
	return *new([]multihash.Multihash), ErrNotSupported
}

func (s *BoostStruct) BoostLegacyDealByProposalCid(p0 context.Context, p1 cid.Cid) (legacytypes.MinerDeal, error) {
	if s.Internal.BoostLegacyDealByProposalCid == nil {
		return *new(legacytypes.MinerDeal), ErrNotSupported
	}
	return s.Internal.BoostLegacyDealByProposalCid(p0, p1)
}

func (s *BoostStub) BoostLegacyDealByProposalCid(p0 context.Context, p1 cid.Cid) (legacytypes.MinerDeal, error) {
	return *new(legacytypes.MinerDeal), ErrNotSupported
}

func (s *BoostStruct) BoostOfflineDealWithData(p0 context.Context, p1 uuid.UUID, p2 string, p3 bool) (*ProviderDealRejectionInfo, error) {
	if s.Internal.BoostOfflineDealWithData == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.BoostOfflineDealWithData(p0, p1, p2, p3)
}

func (s *BoostStub) BoostOfflineDealWithData(p0 context.Context, p1 uuid.UUID, p2 string, p3 bool) (*ProviderDealRejectionInfo, error) {
	return nil, ErrNotSupported
}

func (s *BoostStruct) OnlineBackup(p0 context.Context, p1 string) error {
	if s.Internal.OnlineBackup == nil {
		return ErrNotSupported
	}
	return s.Internal.OnlineBackup(p0, p1)
}

func (s *BoostStub) OnlineBackup(p0 context.Context, p1 string) error {
	return ErrNotSupported
}

func (s *BoostStruct) PdBuildIndexForPieceCid(p0 context.Context, p1 cid.Cid) error {
	if s.Internal.PdBuildIndexForPieceCid == nil {
		return ErrNotSupported
	}
	return s.Internal.PdBuildIndexForPieceCid(p0, p1)
}

func (s *BoostStub) PdBuildIndexForPieceCid(p0 context.Context, p1 cid.Cid) error {
	return ErrNotSupported
}

func (s *BoostStruct) PdCleanup(p0 context.Context) error {
	if s.Internal.PdCleanup == nil {
		return ErrNotSupported
	}
	return s.Internal.PdCleanup(p0)
}

func (s *BoostStub) PdCleanup(p0 context.Context) error {
	return ErrNotSupported
}

func (s *BoostStruct) PdRemoveDealForPiece(p0 context.Context, p1 cid.Cid, p2 string) error {
	if s.Internal.PdRemoveDealForPiece == nil {
		return ErrNotSupported
	}
	return s.Internal.PdRemoveDealForPiece(p0, p1, p2)
}

func (s *BoostStub) PdRemoveDealForPiece(p0 context.Context, p1 cid.Cid, p2 string) error {
	return ErrNotSupported
}

func (s *ChainIOStruct) ChainHasObj(p0 context.Context, p1 cid.Cid) (bool, error) {
	if s.Internal.ChainHasObj == nil {
		return false, ErrNotSupported
	}
	return s.Internal.ChainHasObj(p0, p1)
}

func (s *ChainIOStub) ChainHasObj(p0 context.Context, p1 cid.Cid) (bool, error) {
	return false, ErrNotSupported
}

func (s *ChainIOStruct) ChainReadObj(p0 context.Context, p1 cid.Cid) ([]byte, error) {
	if s.Internal.ChainReadObj == nil {
		return *new([]byte), ErrNotSupported
	}
	return s.Internal.ChainReadObj(p0, p1)
}

func (s *ChainIOStub) ChainReadObj(p0 context.Context, p1 cid.Cid) ([]byte, error) {
	return *new([]byte), ErrNotSupported
}

func (s *CommonStruct) AuthNew(p0 context.Context, p1 []auth.Permission) ([]byte, error) {
	if s.Internal.AuthNew == nil {
		return *new([]byte), ErrNotSupported
	}
	return s.Internal.AuthNew(p0, p1)
}

func (s *CommonStub) AuthNew(p0 context.Context, p1 []auth.Permission) ([]byte, error) {
	return *new([]byte), ErrNotSupported
}

func (s *CommonStruct) AuthVerify(p0 context.Context, p1 string) ([]auth.Permission, error) {
	if s.Internal.AuthVerify == nil {
		return *new([]auth.Permission), ErrNotSupported
	}
	return s.Internal.AuthVerify(p0, p1)
}

func (s *CommonStub) AuthVerify(p0 context.Context, p1 string) ([]auth.Permission, error) {
	return *new([]auth.Permission), ErrNotSupported
}

func (s *CommonStruct) Discover(p0 context.Context) (apitypes.OpenRPCDocument, error) {
	if s.Internal.Discover == nil {
		return *new(apitypes.OpenRPCDocument), ErrNotSupported
	}
	return s.Internal.Discover(p0)
}

func (s *CommonStub) Discover(p0 context.Context) (apitypes.OpenRPCDocument, error) {
	return *new(apitypes.OpenRPCDocument), ErrNotSupported
}

func (s *CommonStruct) LogList(p0 context.Context) ([]string, error) {
	if s.Internal.LogList == nil {
		return *new([]string), ErrNotSupported
	}
	return s.Internal.LogList(p0)
}

func (s *CommonStub) LogList(p0 context.Context) ([]string, error) {
	return *new([]string), ErrNotSupported
}

func (s *CommonStruct) LogSetLevel(p0 context.Context, p1 string, p2 string) error {
	if s.Internal.LogSetLevel == nil {
		return ErrNotSupported
	}
	return s.Internal.LogSetLevel(p0, p1, p2)
}

func (s *CommonStub) LogSetLevel(p0 context.Context, p1 string, p2 string) error {
	return ErrNotSupported
}

func (s *NetStruct) ID(p0 context.Context) (peer.ID, error) {
	if s.Internal.ID == nil {
		return *new(peer.ID), ErrNotSupported
	}
	return s.Internal.ID(p0)
}

func (s *NetStub) ID(p0 context.Context) (peer.ID, error) {
	return *new(peer.ID), ErrNotSupported
}

func (s *NetStruct) NetAddrsListen(p0 context.Context) (peer.AddrInfo, error) {
	if s.Internal.NetAddrsListen == nil {
		return *new(peer.AddrInfo), ErrNotSupported
	}
	return s.Internal.NetAddrsListen(p0)
}

func (s *NetStub) NetAddrsListen(p0 context.Context) (peer.AddrInfo, error) {
	return *new(peer.AddrInfo), ErrNotSupported
}

func (s *NetStruct) NetAgentVersion(p0 context.Context, p1 peer.ID) (string, error) {
	if s.Internal.NetAgentVersion == nil {
		return "", ErrNotSupported
	}
	return s.Internal.NetAgentVersion(p0, p1)
}

func (s *NetStub) NetAgentVersion(p0 context.Context, p1 peer.ID) (string, error) {
	return "", ErrNotSupported
}

func (s *NetStruct) NetAutoNatStatus(p0 context.Context) (lotus_api.NatInfo, error) {
	if s.Internal.NetAutoNatStatus == nil {
		return *new(lotus_api.NatInfo), ErrNotSupported
	}
	return s.Internal.NetAutoNatStatus(p0)
}

func (s *NetStub) NetAutoNatStatus(p0 context.Context) (lotus_api.NatInfo, error) {
	return *new(lotus_api.NatInfo), ErrNotSupported
}

func (s *NetStruct) NetBandwidthStats(p0 context.Context) (metrics.Stats, error) {
	if s.Internal.NetBandwidthStats == nil {
		return *new(metrics.Stats), ErrNotSupported
	}
	return s.Internal.NetBandwidthStats(p0)
}

func (s *NetStub) NetBandwidthStats(p0 context.Context) (metrics.Stats, error) {
	return *new(metrics.Stats), ErrNotSupported
}

func (s *NetStruct) NetBandwidthStatsByPeer(p0 context.Context) (map[string]metrics.Stats, error) {
	if s.Internal.NetBandwidthStatsByPeer == nil {
		return *new(map[string]metrics.Stats), ErrNotSupported
	}
	return s.Internal.NetBandwidthStatsByPeer(p0)
}

func (s *NetStub) NetBandwidthStatsByPeer(p0 context.Context) (map[string]metrics.Stats, error) {
	return *new(map[string]metrics.Stats), ErrNotSupported
}

func (s *NetStruct) NetBandwidthStatsByProtocol(p0 context.Context) (map[protocol.ID]metrics.Stats, error) {
	if s.Internal.NetBandwidthStatsByProtocol == nil {
		return *new(map[protocol.ID]metrics.Stats), ErrNotSupported
	}
	return s.Internal.NetBandwidthStatsByProtocol(p0)
}

func (s *NetStub) NetBandwidthStatsByProtocol(p0 context.Context) (map[protocol.ID]metrics.Stats, error) {
	return *new(map[protocol.ID]metrics.Stats), ErrNotSupported
}

func (s *NetStruct) NetBlockAdd(p0 context.Context, p1 lotus_api.NetBlockList) error {
	if s.Internal.NetBlockAdd == nil {
		return ErrNotSupported
	}
	return s.Internal.NetBlockAdd(p0, p1)
}

func (s *NetStub) NetBlockAdd(p0 context.Context, p1 lotus_api.NetBlockList) error {
	return ErrNotSupported
}

func (s *NetStruct) NetBlockList(p0 context.Context) (lotus_api.NetBlockList, error) {
	if s.Internal.NetBlockList == nil {
		return *new(lotus_api.NetBlockList), ErrNotSupported
	}
	return s.Internal.NetBlockList(p0)
}

func (s *NetStub) NetBlockList(p0 context.Context) (lotus_api.NetBlockList, error) {
	return *new(lotus_api.NetBlockList), ErrNotSupported
}

func (s *NetStruct) NetBlockRemove(p0 context.Context, p1 lotus_api.NetBlockList) error {
	if s.Internal.NetBlockRemove == nil {
		return ErrNotSupported
	}
	return s.Internal.NetBlockRemove(p0, p1)
}

func (s *NetStub) NetBlockRemove(p0 context.Context, p1 lotus_api.NetBlockList) error {
	return ErrNotSupported
}

func (s *NetStruct) NetConnect(p0 context.Context, p1 peer.AddrInfo) error {
	if s.Internal.NetConnect == nil {
		return ErrNotSupported
	}
	return s.Internal.NetConnect(p0, p1)
}

func (s *NetStub) NetConnect(p0 context.Context, p1 peer.AddrInfo) error {
	return ErrNotSupported
}

func (s *NetStruct) NetConnectedness(p0 context.Context, p1 peer.ID) (network.Connectedness, error) {
	if s.Internal.NetConnectedness == nil {
		return *new(network.Connectedness), ErrNotSupported
	}
	return s.Internal.NetConnectedness(p0, p1)
}

func (s *NetStub) NetConnectedness(p0 context.Context, p1 peer.ID) (network.Connectedness, error) {
	return *new(network.Connectedness), ErrNotSupported
}

func (s *NetStruct) NetDisconnect(p0 context.Context, p1 peer.ID) error {
	if s.Internal.NetDisconnect == nil {
		return ErrNotSupported
	}
	return s.Internal.NetDisconnect(p0, p1)
}

func (s *NetStub) NetDisconnect(p0 context.Context, p1 peer.ID) error {
	return ErrNotSupported
}

func (s *NetStruct) NetFindPeer(p0 context.Context, p1 peer.ID) (peer.AddrInfo, error) {
	if s.Internal.NetFindPeer == nil {
		return *new(peer.AddrInfo), ErrNotSupported
	}
	return s.Internal.NetFindPeer(p0, p1)
}

func (s *NetStub) NetFindPeer(p0 context.Context, p1 peer.ID) (peer.AddrInfo, error) {
	return *new(peer.AddrInfo), ErrNotSupported
}

func (s *NetStruct) NetLimit(p0 context.Context, p1 string) (lotus_api.NetLimit, error) {
	if s.Internal.NetLimit == nil {
		return *new(lotus_api.NetLimit), ErrNotSupported
	}
	return s.Internal.NetLimit(p0, p1)
}

func (s *NetStub) NetLimit(p0 context.Context, p1 string) (lotus_api.NetLimit, error) {
	return *new(lotus_api.NetLimit), ErrNotSupported
}

func (s *NetStruct) NetPeerInfo(p0 context.Context, p1 peer.ID) (*lotus_api.ExtendedPeerInfo, error) {
	if s.Internal.NetPeerInfo == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.NetPeerInfo(p0, p1)
}

func (s *NetStub) NetPeerInfo(p0 context.Context, p1 peer.ID) (*lotus_api.ExtendedPeerInfo, error) {
	return nil, ErrNotSupported
}

func (s *NetStruct) NetPeers(p0 context.Context) ([]peer.AddrInfo, error) {
	if s.Internal.NetPeers == nil {
		return *new([]peer.AddrInfo), ErrNotSupported
	}
	return s.Internal.NetPeers(p0)
}

func (s *NetStub) NetPeers(p0 context.Context) ([]peer.AddrInfo, error) {
	return *new([]peer.AddrInfo), ErrNotSupported
}

func (s *NetStruct) NetPing(p0 context.Context, p1 peer.ID) (time.Duration, error) {
	if s.Internal.NetPing == nil {
		return *new(time.Duration), ErrNotSupported
	}
	return s.Internal.NetPing(p0, p1)
}

func (s *NetStub) NetPing(p0 context.Context, p1 peer.ID) (time.Duration, error) {
	return *new(time.Duration), ErrNotSupported
}

func (s *NetStruct) NetProtectAdd(p0 context.Context, p1 []peer.ID) error {
	if s.Internal.NetProtectAdd == nil {
		return ErrNotSupported
	}
	return s.Internal.NetProtectAdd(p0, p1)
}

func (s *NetStub) NetProtectAdd(p0 context.Context, p1 []peer.ID) error {
	return ErrNotSupported
}

func (s *NetStruct) NetProtectList(p0 context.Context) ([]peer.ID, error) {
	if s.Internal.NetProtectList == nil {
		return *new([]peer.ID), ErrNotSupported
	}
	return s.Internal.NetProtectList(p0)
}

func (s *NetStub) NetProtectList(p0 context.Context) ([]peer.ID, error) {
	return *new([]peer.ID), ErrNotSupported
}

func (s *NetStruct) NetProtectRemove(p0 context.Context, p1 []peer.ID) error {
	if s.Internal.NetProtectRemove == nil {
		return ErrNotSupported
	}
	return s.Internal.NetProtectRemove(p0, p1)
}

func (s *NetStub) NetProtectRemove(p0 context.Context, p1 []peer.ID) error {
	return ErrNotSupported
}

func (s *NetStruct) NetPubsubScores(p0 context.Context) ([]lotus_api.PubsubScore, error) {
	if s.Internal.NetPubsubScores == nil {
		return *new([]lotus_api.PubsubScore), ErrNotSupported
	}
	return s.Internal.NetPubsubScores(p0)
}

func (s *NetStub) NetPubsubScores(p0 context.Context) ([]lotus_api.PubsubScore, error) {
	return *new([]lotus_api.PubsubScore), ErrNotSupported
}

func (s *NetStruct) NetSetLimit(p0 context.Context, p1 string, p2 lotus_api.NetLimit) error {
	if s.Internal.NetSetLimit == nil {
		return ErrNotSupported
	}
	return s.Internal.NetSetLimit(p0, p1, p2)
}

func (s *NetStub) NetSetLimit(p0 context.Context, p1 string, p2 lotus_api.NetLimit) error {
	return ErrNotSupported
}

func (s *NetStruct) NetStat(p0 context.Context, p1 string) (lotus_api.NetStat, error) {
	if s.Internal.NetStat == nil {
		return *new(lotus_api.NetStat), ErrNotSupported
	}
	return s.Internal.NetStat(p0, p1)
}

func (s *NetStub) NetStat(p0 context.Context, p1 string) (lotus_api.NetStat, error) {
	return *new(lotus_api.NetStat), ErrNotSupported
}

func (s *WalletStruct) WalletDelete(p0 context.Context, p1 address.Address) error {
	if s.Internal.WalletDelete == nil {
		return ErrNotSupported
	}
	return s.Internal.WalletDelete(p0, p1)
}

func (s *WalletStub) WalletDelete(p0 context.Context, p1 address.Address) error {
	return ErrNotSupported
}

func (s *WalletStruct) WalletExport(p0 context.Context, p1 address.Address) (*types.KeyInfo, error) {
	if s.Internal.WalletExport == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.WalletExport(p0, p1)
}

func (s *WalletStub) WalletExport(p0 context.Context, p1 address.Address) (*types.KeyInfo, error) {
	return nil, ErrNotSupported
}

func (s *WalletStruct) WalletHas(p0 context.Context, p1 address.Address) (bool, error) {
	if s.Internal.WalletHas == nil {
		return false, ErrNotSupported
	}
	return s.Internal.WalletHas(p0, p1)
}

func (s *WalletStub) WalletHas(p0 context.Context, p1 address.Address) (bool, error) {
	return false, ErrNotSupported
}

func (s *WalletStruct) WalletImport(p0 context.Context, p1 *types.KeyInfo) (address.Address, error) {
	if s.Internal.WalletImport == nil {
		return *new(address.Address), ErrNotSupported
	}
	return s.Internal.WalletImport(p0, p1)
}

func (s *WalletStub) WalletImport(p0 context.Context, p1 *types.KeyInfo) (address.Address, error) {
	return *new(address.Address), ErrNotSupported
}

func (s *WalletStruct) WalletList(p0 context.Context) ([]address.Address, error) {
	if s.Internal.WalletList == nil {
		return *new([]address.Address), ErrNotSupported
	}
	return s.Internal.WalletList(p0)
}

func (s *WalletStub) WalletList(p0 context.Context) ([]address.Address, error) {
	return *new([]address.Address), ErrNotSupported
}

func (s *WalletStruct) WalletNew(p0 context.Context, p1 types.KeyType) (address.Address, error) {
	if s.Internal.WalletNew == nil {
		return *new(address.Address), ErrNotSupported
	}
	return s.Internal.WalletNew(p0, p1)
}

func (s *WalletStub) WalletNew(p0 context.Context, p1 types.KeyType) (address.Address, error) {
	return *new(address.Address), ErrNotSupported
}

func (s *WalletStruct) WalletSign(p0 context.Context, p1 address.Address, p2 []byte) (*crypto.Signature, error) {
	if s.Internal.WalletSign == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.WalletSign(p0, p1, p2)
}

func (s *WalletStub) WalletSign(p0 context.Context, p1 address.Address, p2 []byte) (*crypto.Signature, error) {
	return nil, ErrNotSupported
}

var _ Boost = new(BoostStruct)
var _ ChainIO = new(ChainIOStruct)
var _ Common = new(CommonStruct)
var _ CommonNet = new(CommonNetStruct)
var _ Net = new(NetStruct)
var _ Wallet = new(WalletStruct)
