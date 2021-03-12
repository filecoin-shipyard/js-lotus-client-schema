const commonMethods = {
  AuthNew: {},
  AuthVerify: {},
  Closing: {
    subscription: true
  },
  ID: {},
  LogList: {},
  LogSetLevel: {},
  NetAddrsListen: {},
  NetAgentVersion: {},
  NetAutoNatStatus: {},
  NetBandwidthStats: {},
  NetBandwidthStatsByPeer: {},
  NetBandwidthStatsByProtocol: {},
  NetBlockAdd: {},
  NetBlockList: {},
  NetBlockRemove: {},
  NetConnect: {},
  NetConnectedness: {},
  NetDisconnect: {},
  NetFindPeer: {},
  NetPeerInfo: {},
  NetPeers: {},
  NetPubsubScores: {},
  Session: {},
  Shutdown: {},
  Version: {}
}

const fullNodeMethods = {
  AuthNew: {},
  AuthVerify: {},
  BeaconGetEntry: {},
  ChainDeleteObj: {},
  ChainExport: {
    subscription: true
  },
  ChainGetBlock: {},
  ChainGetBlockMessages: {},
  ChainGetGenesis: {},
  ChainGetMessage: {},
  ChainGetNode: {},
  ChainGetParentMessages: {},
  ChainGetParentReceipts: {},
  ChainGetPath: {},
  ChainGetRandomnessFromBeacon: {},
  ChainGetRandomnessFromTickets: {},
  ChainGetTipSet: {},
  ChainGetTipSetByHeight: {},
  ChainHasObj: {},
  ChainHead: {},
  ChainNotify: {
    subscription: true
  },
  ChainReadObj: {},
  ChainSetHead: {},
  ChainStatObj: {},
  ChainTipSetWeight: {},
  ClientCalcCommP: {},
  ClientCancelDataTransfer: {},
  ClientDataTransferUpdates: {
    subscription: true
  },
  ClientDealPieceCID: {},
  ClientDealSize: {},
  ClientFindData: {},
  ClientGenCar: {},
  ClientGetDealInfo: {},
  ClientGetDealStatus: {},
  ClientGetDealUpdates: {
    subscription: true
  },
  ClientHasLocal: {},
  ClientImport: {},
  ClientListDataTransfers: {},
  ClientListDeals: {},
  ClientListImports: {},
  ClientMinerQueryOffer: {},
  ClientQueryAsk: {},
  ClientRemoveImport: {},
  ClientRestartDataTransfer: {},
  ClientRetrieve: {},
  ClientRetrieveTryRestartInsufficientFunds: {},
  ClientRetrieveWithEvents: {
    subscription: true
  },
  ClientStartDeal: {},
  Closing: {
    subscription: true
  },
  CreateBackup: {},
  GasEstimateFeeCap: {},
  GasEstimateGasLimit: {},
  GasEstimateGasPremium: {},
  GasEstimateMessageGas: {},
  ID: {},
  LogList: {},
  LogSetLevel: {},
  MarketAddBalance: {},
  MarketGetReserved: {},
  MarketReleaseFunds: {},
  MarketReserveFunds: {},
  MarketWithdraw: {},
  MinerCreateBlock: {},
  MinerGetBaseInfo: {},
  MpoolBatchPush: {},
  MpoolBatchPushMessage: {},
  MpoolBatchPushUntrusted: {},
  MpoolClear: {},
  MpoolGetConfig: {},
  MpoolGetNonce: {},
  MpoolPending: {},
  MpoolPush: {},
  MpoolPushMessage: {},
  MpoolPushUntrusted: {},
  MpoolSelect: {},
  MpoolSetConfig: {},
  MpoolSub: {
    subscription: true
  },
  MsigAddApprove: {},
  MsigAddCancel: {},
  MsigAddPropose: {},
  MsigApprove: {},
  MsigApproveTxnHash: {},
  MsigCancel: {},
  MsigCreate: {},
  MsigGetAvailableBalance: {},
  MsigGetPending: {},
  MsigGetVested: {},
  MsigGetVestingSchedule: {},
  MsigPropose: {},
  MsigRemoveSigner: {},
  MsigSwapApprove: {},
  MsigSwapCancel: {},
  MsigSwapPropose: {},
  NetAddrsListen: {},
  NetAgentVersion: {},
  NetAutoNatStatus: {},
  NetBandwidthStats: {},
  NetBandwidthStatsByPeer: {},
  NetBandwidthStatsByProtocol: {},
  NetBlockAdd: {},
  NetBlockList: {},
  NetBlockRemove: {},
  NetConnect: {},
  NetConnectedness: {},
  NetDisconnect: {},
  NetFindPeer: {},
  NetPeerInfo: {},
  NetPeers: {},
  NetPubsubScores: {},
  PaychAllocateLane: {},
  PaychAvailableFunds: {},
  PaychAvailableFundsByFromTo: {},
  PaychCollect: {},
  PaychGet: {},
  PaychGetWaitReady: {},
  PaychList: {},
  PaychNewPayment: {},
  PaychSettle: {},
  PaychStatus: {},
  PaychVoucherAdd: {},
  PaychVoucherCheckSpendable: {},
  PaychVoucherCheckValid: {},
  PaychVoucherCreate: {},
  PaychVoucherList: {},
  PaychVoucherSubmit: {},
  Session: {},
  Shutdown: {},
  StateAccountKey: {},
  StateAllMinerFaults: {},
  StateCall: {},
  StateChangedActors: {},
  StateCirculatingSupply: {},
  StateCompute: {},
  StateDealProviderCollateralBounds: {},
  StateDecodeParams: {},
  StateGetActor: {},
  StateGetReceipt: {},
  StateListActors: {},
  StateListMessages: {},
  StateListMiners: {},
  StateLookupID: {},
  StateMarketBalance: {},
  StateMarketDeals: {},
  StateMarketParticipants: {},
  StateMarketStorageDeal: {},
  StateMinerActiveSectors: {},
  StateMinerAvailableBalance: {},
  StateMinerDeadlines: {},
  StateMinerFaults: {},
  StateMinerInfo: {},
  StateMinerInitialPledgeCollateral: {},
  StateMinerPartitions: {},
  StateMinerPower: {},
  StateMinerPreCommitDepositForPower: {},
  StateMinerProvingDeadline: {},
  StateMinerRecoveries: {},
  StateMinerSectorAllocated: {},
  StateMinerSectorCount: {},
  StateMinerSectors: {},
  StateNetworkName: {},
  StateNetworkVersion: {},
  StateReadState: {},
  StateReplay: {},
  StateSearchMsg: {},
  StateSearchMsgLimited: {},
  StateSectorExpiration: {},
  StateSectorGetInfo: {},
  StateSectorPartition: {},
  StateSectorPreCommitInfo: {},
  StateVMCirculatingSupplyInternal: {},
  StateVerifiedClientStatus: {},
  StateVerifiedRegistryRootKey: {},
  StateVerifierStatus: {},
  StateWaitMsg: {},
  StateWaitMsgLimited: {},
  SyncCheckBad: {},
  SyncCheckpoint: {},
  SyncIncomingBlocks: {
    subscription: true
  },
  SyncMarkBad: {},
  SyncState: {},
  SyncSubmitBlock: {},
  SyncUnmarkAllBad: {},
  SyncUnmarkBad: {},
  SyncValidateTipset: {},
  Version: {},
  WalletBalance: {},
  WalletDefaultAddress: {},
  WalletDelete: {},
  WalletExport: {},
  WalletHas: {},
  WalletImport: {},
  WalletList: {},
  WalletNew: {},
  WalletSetDefault: {},
  WalletSign: {},
  WalletSignMessage: {},
  WalletValidateAddress: {},
  WalletVerify: {}
}

const storageMinerMethods = {
  ActorAddress: {},
  ActorAddressConfig: {},
  ActorSectorSize: {},
  AuthNew: {},
  AuthVerify: {},
  CheckProvable: {},
  Closing: {
    subscription: true
  },
  CreateBackup: {},
  DealsConsiderOfflineRetrievalDeals: {},
  DealsConsiderOfflineStorageDeals: {},
  DealsConsiderOnlineRetrievalDeals: {},
  DealsConsiderOnlineStorageDeals: {},
  DealsConsiderUnverifiedStorageDeals: {},
  DealsConsiderVerifiedStorageDeals: {},
  DealsImportData: {},
  DealsList: {},
  DealsPieceCidBlocklist: {},
  DealsSetConsiderOfflineRetrievalDeals: {},
  DealsSetConsiderOfflineStorageDeals: {},
  DealsSetConsiderOnlineRetrievalDeals: {},
  DealsSetConsiderOnlineStorageDeals: {},
  DealsSetConsiderUnverifiedStorageDeals: {},
  DealsSetConsiderVerifiedStorageDeals: {},
  DealsSetPieceCidBlocklist: {},
  ID: {},
  LogList: {},
  LogSetLevel: {},
  MarketCancelDataTransfer: {},
  MarketDataTransferUpdates: {
    subscription: true
  },
  MarketGetAsk: {},
  MarketGetDealUpdates: {
    subscription: true
  },
  MarketGetRetrievalAsk: {},
  MarketImportDealData: {},
  MarketListDataTransfers: {},
  MarketListDeals: {},
  MarketListIncompleteDeals: {},
  MarketListRetrievalDeals: {},
  MarketPendingDeals: {},
  MarketPublishPendingDeals: {},
  MarketRestartDataTransfer: {},
  MarketSetAsk: {},
  MarketSetRetrievalAsk: {},
  MiningBase: {},
  NetAddrsListen: {},
  NetAgentVersion: {},
  NetAutoNatStatus: {},
  NetBandwidthStats: {},
  NetBandwidthStatsByPeer: {},
  NetBandwidthStatsByProtocol: {},
  NetBlockAdd: {},
  NetBlockList: {},
  NetBlockRemove: {},
  NetConnect: {},
  NetConnectedness: {},
  NetDisconnect: {},
  NetFindPeer: {},
  NetPeerInfo: {},
  NetPeers: {},
  NetPubsubScores: {},
  PiecesGetCIDInfo: {},
  PiecesGetPieceInfo: {},
  PiecesListCidInfos: {},
  PiecesListPieces: {},
  PledgeSector: {},
  ReturnAddPiece: {},
  ReturnFetch: {},
  ReturnFinalizeSector: {},
  ReturnMoveStorage: {},
  ReturnReadPiece: {},
  ReturnReleaseUnsealed: {},
  ReturnSealCommit1: {},
  ReturnSealCommit2: {},
  ReturnSealPreCommit1: {},
  ReturnSealPreCommit2: {},
  ReturnUnsealPiece: {},
  SealingAbort: {},
  SealingSchedDiag: {},
  SectorGetExpectedSealDuration: {},
  SectorGetSealDelay: {},
  SectorMarkForUpgrade: {},
  SectorRemove: {},
  SectorSetExpectedSealDuration: {},
  SectorSetSealDelay: {},
  SectorStartSealing: {},
  SectorTerminate: {},
  SectorTerminateFlush: {},
  SectorTerminatePending: {},
  SectorsList: {},
  SectorsListInStates: {},
  SectorsRefs: {},
  SectorsStatus: {},
  SectorsSummary: {},
  SectorsUpdate: {},
  Session: {},
  Shutdown: {},
  StorageAddLocal: {},
  StorageAttach: {},
  StorageBestAlloc: {},
  StorageDeclareSector: {},
  StorageDropSector: {},
  StorageFindSector: {},
  StorageInfo: {},
  StorageList: {},
  StorageLocal: {},
  StorageLock: {},
  StorageReportHealth: {},
  StorageStat: {},
  StorageTryLock: {},
  Version: {},
  WorkerConnect: {},
  WorkerJobs: {},
  WorkerStats: {}
}

const gatewayApiMethods = {
  ChainGetBlockMessages: {},
  ChainGetMessage: {},
  ChainGetTipSet: {},
  ChainGetTipSetByHeight: {},
  ChainHasObj: {},
  ChainHead: {},
  ChainNotify: {
    subscription: true
  },
  ChainReadObj: {},
  GasEstimateMessageGas: {},
  MpoolPush: {},
  MsigGetAvailableBalance: {},
  MsigGetPending: {},
  MsigGetVested: {},
  StateAccountKey: {},
  StateDealProviderCollateralBounds: {},
  StateGetActor: {},
  StateGetReceipt: {},
  StateListMiners: {},
  StateLookupID: {},
  StateMarketBalance: {},
  StateMarketStorageDeal: {},
  StateMinerInfo: {},
  StateMinerPower: {},
  StateMinerProvingDeadline: {},
  StateNetworkVersion: {},
  StateSearchMsg: {},
  StateSectorGetInfo: {},
  StateVerifiedClientStatus: {},
  StateWaitMsg: {}
}

const walletApiMethods = {
  WalletDelete: {},
  WalletExport: {},
  WalletHas: {},
  WalletImport: {},
  WalletList: {},
  WalletNew: {},
  WalletSign: {}
}

const workerApiMethods = {
  AddPiece: {},
  Enabled: {},
  Fetch: {},
  FinalizeSector: {},
  Info: {},
  MoveStorage: {},
  Paths: {},
  ProcessSession: {},
  ReadPiece: {},
  ReleaseUnsealed: {},
  Remove: {},
  SealCommit1: {},
  SealCommit2: {},
  SealPreCommit1: {},
  SealPreCommit2: {},
  Session: {},
  SetEnabled: {},
  StorageAddLocal: {},
  TaskDisable: {},
  TaskEnable: {},
  TaskTypes: {},
  UnsealPiece: {},
  Version: {},
  WaitQuiet: {}
}

const common = {
  methods: commonMethods
}

const fullNode = {
  methods: {
    ...commonMethods,
    ...fullNodeMethods
  }
}

const storageMiner = {
  methods: {
    ...commonMethods,
    ...storageMinerMethods
  }
}

const gatewayApi = {
  methods: gatewayApiMethods
}

const walletApi = {
  methods: walletApiMethods
}

const workerApi = {
  methods: workerApiMethods
}

module.exports = {
  common,
  fullNode,
  storageMiner,
  gatewayApi,
  walletApi,
  workerApi
}
