## 


### Version


Inputs: `null`

Response: `{"Version":"string value","APIVersion":768,"BlockDelay":42}`

## Auth


### AuthNew


Inputs: `[null]`

Response: `"Ynl0ZSBhcnJheQ=="`

### AuthVerify


Inputs: `["string value"]`

Response: `null`

## Chain
The Chain method group contains methods for interacting with the
blockchain, but that do not require any form of state computation


### ChainExport
There are not yet any comments for this method.

Inputs: `[[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `"Ynl0ZSBhcnJheQ=="`

### ChainGetBlock
ChainGetBlock returns the block specified by the given CID


Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `{"Miner":"t01234","Ticket":{"VRFProof":"Ynl0ZSBhcnJheQ=="},"ElectionProof":{"VRFProof":"Ynl0ZSBhcnJheQ=="},"BeaconEntries":null,"WinPoStProof":null,"Parents":null,"ParentWeight":"0","Height":10101,"ParentStateRoot":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"ParentMessageReceipts":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"Messages":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"BLSAggregate":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="},"Timestamp":42,"BlockSig":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="},"ForkSignaling":42}`

### ChainGetBlockMessages
There are not yet any comments for this method.

Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `{"BlsMessages":null,"SecpkMessages":null,"Cids":null}`

### ChainGetGenesis
There are not yet any comments for this method.

Inputs: `null`

Response: `{"Cids":null,"Blocks":null,"Height":0}`

### ChainGetMessage
There are not yet any comments for this method.

Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `{"Version":9,"To":"t01234","From":"t01234","Nonce":42,"Value":"0","GasPrice":"0","GasLimit":9,"Method":1,"Params":"Ynl0ZSBhcnJheQ=="}`

### ChainGetNode
There are not yet any comments for this method.

Inputs: `["string value"]`

Response: `{"Cid":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"Obj":{}}`

### ChainGetParentMessages
There are not yet any comments for this method.

Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `null`

### ChainGetParentReceipts
There are not yet any comments for this method.

Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `null`

### ChainGetPath
There are not yet any comments for this method.

Inputs: `[[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}],[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `null`

### ChainGetRandomness
ChainGetRandomness is used to sample the chain for randomness


Inputs: `[[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}],2,10101,"Ynl0ZSBhcnJheQ=="]`

Response: `null`

### ChainGetTipSet
There are not yet any comments for this method.

Inputs: `[[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"Cids":null,"Blocks":null,"Height":0}`

### ChainGetTipSetByHeight
There are not yet any comments for this method.

Inputs: `[10101,[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"Cids":null,"Blocks":null,"Height":0}`

### ChainHasObj
There are not yet any comments for this method.

Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `true`

### ChainHead
ChainHead returns the current head of the chain


Inputs: `null`

Response: `{"Cids":null,"Blocks":null,"Height":0}`

### ChainNotify
ChainNotify returns channel with chain head updates
First message is guaranteed to be of len == 1, and type == 'current'


Inputs: `null`

Response: `null`

### ChainReadObj
There are not yet any comments for this method.

Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `"Ynl0ZSBhcnJheQ=="`

### ChainSetHead
There are not yet any comments for this method.

Inputs: `[[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{}`

### ChainStatObj
There are not yet any comments for this method.

Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `{"Size":42,"Links":42}`

### ChainTipSetWeight
There are not yet any comments for this method.

Inputs: `[[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `"0"`

## Client
The Client methods all have to do with interacting with the storage and
retrieval markets as a client


### ClientCalcCommP
There are not yet any comments for this method.

Inputs: `["string value","t01234"]`

Response: `{"Root":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"Size":1024}`

### ClientFindData
There are not yet any comments for this method.

Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `null`

### ClientGenCar
There are not yet any comments for this method.

Inputs: `[{"Path":"string value","IsCAR":true},"string value"]`

Response: `{}`

### ClientGetDealInfo
ClientGetDeal info returns the latest information about a given deal


Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `{"ProposalCid":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"State":42,"Message":"string value","Provider":"t01234","PieceCID":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"Size":42,"PricePerEpoch":"0","Duration":42,"DealID":5432}`

### ClientHasLocal
There are not yet any comments for this method.

Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `true`

### ClientImport
ClientImport imports file under the specified path into filestore


Inputs: `[{"Path":"string value","IsCAR":true}]`

Response: `{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}`

### ClientListDeals
There are not yet any comments for this method.

Inputs: `null`

Response: `null`

### ClientListImports
ClientListImports lists imported files and their root CIDs


Inputs: `null`

Response: `null`

### ClientQueryAsk
There are not yet any comments for this method.

Inputs: `["12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf","t01234"]`

Response: `{"Ask":{"Price":"0","MinPieceSize":1032,"MaxPieceSize":1032,"Miner":"t01234","Timestamp":10101,"Expiry":10101,"SeqNo":42},"Signature":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="}}`

### ClientRetrieve
There are not yet any comments for this method.

Inputs: `[{"Root":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"Size":42,"Total":"0","PaymentInterval":42,"PaymentIntervalIncrease":42,"Client":"t01234","Miner":"t01234","MinerPeerID":"12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"},{"Path":"string value","IsCAR":true}]`

Response: `{}`

### ClientStartDeal
ClientStartDeal proposes a deal with a miner


Inputs: `[{"Data":{"TransferType":"string value","Root":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"PieceCid":null,"PieceSize":1024},"Wallet":"t01234","Miner":"t01234","EpochPrice":"0","MinBlocksDuration":42,"DealStartEpoch":10101}]`

Response: `null`

## I


### ID


Inputs: `null`

Response: `"12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"`

## Log


### LogList


Inputs: `null`

Response: `null`

### LogSetLevel


Inputs: `["string value","string value"]`

Response: `{}`

## Market


### MarketEnsureAvailable
MarketFreeBalance


Inputs: `["t01234","t01234","0"]`

Response: `{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}`

## Miner


### MinerCreateBlock
There are not yet any comments for this method.

Inputs: `[{"Miner":"t01234","Parents":[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}],"Ticket":{"VRFProof":"Ynl0ZSBhcnJheQ=="},"Eproof":{"VRFProof":"Ynl0ZSBhcnJheQ=="},"BeaconValues":null,"Messages":null,"Epoch":10101,"Timestamp":42,"WinningPoStProof":null}]`

Response: `{"Header":{"Miner":"t01234","Ticket":{"VRFProof":"Ynl0ZSBhcnJheQ=="},"ElectionProof":{"VRFProof":"Ynl0ZSBhcnJheQ=="},"BeaconEntries":null,"WinPoStProof":null,"Parents":null,"ParentWeight":"0","Height":10101,"ParentStateRoot":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"ParentMessageReceipts":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"Messages":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"BLSAggregate":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="},"Timestamp":42,"BlockSig":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="},"ForkSignaling":42},"BlsMessages":null,"SecpkMessages":null}`

### MinerGetBaseInfo
There are not yet any comments for this method.

Inputs: `["t01234",10101,[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"MinerPower":"0","NetworkPower":"0","Sectors":null,"WorkerKey":"t01234","SectorSize":34359738368,"PrevBeaconEntry":{"Round":42,"Data":"Ynl0ZSBhcnJheQ=="},"BeaconEntries":null}`

## Mpool
The Mpool methods are for interacting with the message pool. The message pool
manages all incoming and outgoing 'messages' going over the network.


### MpoolEstimateGasPrice
There are not yet any comments for this method.

Inputs: `[42,"t01234",9,[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `"0"`

### MpoolGetNonce
There are not yet any comments for this method.

Inputs: `["t01234"]`

Response: `42`

### MpoolPending
There are not yet any comments for this method.

Inputs: `[[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `null`

### MpoolPush
There are not yet any comments for this method.

Inputs: `[{"Message":{"Version":9,"To":"t01234","From":"t01234","Nonce":42,"Value":"0","GasPrice":"0","GasLimit":9,"Method":1,"Params":"Ynl0ZSBhcnJheQ=="},"Signature":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="}}]`

Response: `{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}`

### MpoolPushMessage
get nonce, sign, push


Inputs: `[{"Version":9,"To":"t01234","From":"t01234","Nonce":42,"Value":"0","GasPrice":"0","GasLimit":9,"Method":1,"Params":"Ynl0ZSBhcnJheQ=="}]`

Response: `{"Message":{"Version":9,"To":"t01234","From":"t01234","Nonce":42,"Value":"0","GasPrice":"0","GasLimit":9,"Method":1,"Params":"Ynl0ZSBhcnJheQ=="},"Signature":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="}}`

### MpoolSub
There are not yet any comments for this method.

Inputs: `null`

Response: `{"Type":0,"Message":{"Message":{"Version":9,"To":"t01234","From":"t01234","Nonce":42,"Value":"0","GasPrice":"0","GasLimit":9,"Method":1,"Params":"Ynl0ZSBhcnJheQ=="},"Signature":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="}}}`

## Msig
The Msig methods are used to interact with multisig wallets on the
filecoin network


### MsigApprove
There are not yet any comments for this method.

Inputs: `["t01234",42,"t01234","t01234","0","t01234",42,"Ynl0ZSBhcnJheQ=="]`

Response: `{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}`

### MsigCancel
There are not yet any comments for this method.

Inputs: `["t01234",42,"t01234","t01234","0","t01234",42,"Ynl0ZSBhcnJheQ=="]`

Response: `{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}`

### MsigCreate
There are not yet any comments for this method.

Inputs: `[9,null,"0","t01234","0"]`

Response: `{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}`

### MsigGetAvailableBalance
There are not yet any comments for this method.

Inputs: `["t01234",[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `"0"`

### MsigPropose
There are not yet any comments for this method.

Inputs: `["t01234","t01234","0","t01234",42,"Ynl0ZSBhcnJheQ=="]`

Response: `{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}`

## Net


### NetAddrsListen


Inputs: `null`

Response: `{"Addrs":null,"ID":"12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"}`

### NetConnect


Inputs: `[{"Addrs":null,"ID":"12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"}]`

Response: `{}`

### NetConnectedness


Inputs: `["12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"]`

Response: `1`

### NetDisconnect


Inputs: `["12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"]`

Response: `{}`

### NetFindPeer


Inputs: `["12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"]`

Response: `{"Addrs":null,"ID":"12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"}`

### NetPeers


Inputs: `null`

Response: `null`

## Paych
The Paych methods are for interacting with and managing payment channels


### PaychAllocateLane
There are not yet any comments for this method.

Inputs: `["t01234"]`

Response: `42`

### PaychClose
There are not yet any comments for this method.

Inputs: `["t01234"]`

Response: `{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}`

### PaychGet
There are not yet any comments for this method.

Inputs: `["t01234","t01234","0"]`

Response: `{"Channel":"t01234","ChannelMessage":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}}`

### PaychList
There are not yet any comments for this method.

Inputs: `null`

Response: `null`

### PaychNewPayment
There are not yet any comments for this method.

Inputs: `["t01234","t01234",null]`

Response: `{"Channel":"t01234","ChannelMessage":null,"Vouchers":null}`

### PaychStatus
There are not yet any comments for this method.

Inputs: `["t01234"]`

Response: `{"ControlAddr":"t01234","Direction":1}`

### PaychVoucherAdd
There are not yet any comments for this method.

Inputs: `["t01234",{"TimeLockMin":10101,"TimeLockMax":10101,"SecretPreimage":"Ynl0ZSBhcnJheQ==","Extra":{"Actor":"t01234","Method":1,"Data":"Ynl0ZSBhcnJheQ=="},"Lane":42,"Nonce":42,"Amount":"0","MinSettleHeight":10101,"Merges":null,"Signature":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="}},"Ynl0ZSBhcnJheQ==","0"]`

Response: `"0"`

### PaychVoucherCheckSpendable
There are not yet any comments for this method.

Inputs: `["t01234",{"TimeLockMin":10101,"TimeLockMax":10101,"SecretPreimage":"Ynl0ZSBhcnJheQ==","Extra":{"Actor":"t01234","Method":1,"Data":"Ynl0ZSBhcnJheQ=="},"Lane":42,"Nonce":42,"Amount":"0","MinSettleHeight":10101,"Merges":null,"Signature":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="}},"Ynl0ZSBhcnJheQ==","Ynl0ZSBhcnJheQ=="]`

Response: `true`

### PaychVoucherCheckValid
There are not yet any comments for this method.

Inputs: `["t01234",{"TimeLockMin":10101,"TimeLockMax":10101,"SecretPreimage":"Ynl0ZSBhcnJheQ==","Extra":{"Actor":"t01234","Method":1,"Data":"Ynl0ZSBhcnJheQ=="},"Lane":42,"Nonce":42,"Amount":"0","MinSettleHeight":10101,"Merges":null,"Signature":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="}}]`

Response: `{}`

### PaychVoucherCreate
There are not yet any comments for this method.

Inputs: `["t01234","0",42]`

Response: `{"TimeLockMin":10101,"TimeLockMax":10101,"SecretPreimage":"Ynl0ZSBhcnJheQ==","Extra":{"Actor":"t01234","Method":1,"Data":"Ynl0ZSBhcnJheQ=="},"Lane":42,"Nonce":42,"Amount":"0","MinSettleHeight":10101,"Merges":null,"Signature":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="}}`

### PaychVoucherList
There are not yet any comments for this method.

Inputs: `["t01234"]`

Response: `null`

### PaychVoucherSubmit
There are not yet any comments for this method.

Inputs: `["t01234",{"TimeLockMin":10101,"TimeLockMax":10101,"SecretPreimage":"Ynl0ZSBhcnJheQ==","Extra":{"Actor":"t01234","Method":1,"Data":"Ynl0ZSBhcnJheQ=="},"Lane":42,"Nonce":42,"Amount":"0","MinSettleHeight":10101,"Merges":null,"Signature":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="}}]`

Response: `{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}`

## State
The State methods are used to query, inspect, and interact with chain state


### StateAccountKey
There are not yet any comments for this method.

Inputs: `["t01234",[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `"t01234"`

### StateCall
if tipset is nil, we'll use heaviest


Inputs: `[{"Version":9,"To":"t01234","From":"t01234","Nonce":42,"Value":"0","GasPrice":"0","GasLimit":9,"Method":1,"Params":"Ynl0ZSBhcnJheQ=="},[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"Msg":{"Version":9,"To":"t01234","From":"t01234","Nonce":42,"Value":"0","GasPrice":"0","GasLimit":9,"Method":1,"Params":"Ynl0ZSBhcnJheQ=="},"MsgRct":{"ExitCode":0,"Return":"Ynl0ZSBhcnJheQ==","GasUsed":9},"InternalExecutions":null,"Error":"string value","Duration":60000000000}`

### StateChangedActors
There are not yet any comments for this method.

Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `{"t01236":{"Code":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"Head":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"Nonce":42,"Balance":"0"}}`

### StateCompute
There are not yet any comments for this method.

Inputs: `[10101,null,[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"Root":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"Trace":null}`

### StateGetActor
There are not yet any comments for this method.

Inputs: `["t01234",[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"Code":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"Head":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"Nonce":42,"Balance":"0"}`

### StateGetReceipt
There are not yet any comments for this method.

Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"ExitCode":0,"Return":"Ynl0ZSBhcnJheQ==","GasUsed":9}`

### StateListActors
There are not yet any comments for this method.

Inputs: `[[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `null`

### StateListMessages
There are not yet any comments for this method.

Inputs: `[{"Version":9,"To":"t01234","From":"t01234","Nonce":42,"Value":"0","GasPrice":"0","GasLimit":9,"Method":1,"Params":"Ynl0ZSBhcnJheQ=="},[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}],10101]`

Response: `null`

### StateListMiners
There are not yet any comments for this method.

Inputs: `[[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `null`

### StateLookupID
There are not yet any comments for this method.

Inputs: `["t01234",[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `"t01234"`

### StateMarketBalance
There are not yet any comments for this method.

Inputs: `["t01234",[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"Escrow":"0","Locked":"0"}`

### StateMarketDeals
There are not yet any comments for this method.

Inputs: `[[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"t026363":{"Proposal":{"PieceCID":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"PieceSize":1032,"VerifiedDeal":true,"Client":"t01234","Provider":"t01234","StartEpoch":10101,"EndEpoch":10101,"StoragePricePerEpoch":"0","ProviderCollateral":"0","ClientCollateral":"0"},"State":{"SectorStartEpoch":10101,"LastUpdatedEpoch":10101,"SlashEpoch":10101}}}`

### StateMarketParticipants
There are not yet any comments for this method.

Inputs: `[[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"t026363":{"Escrow":"0","Locked":"0"}}`

### StateMarketStorageDeal
There are not yet any comments for this method.

Inputs: `[5432,[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"Proposal":{"PieceCID":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"PieceSize":1032,"VerifiedDeal":true,"Client":"t01234","Provider":"t01234","StartEpoch":10101,"EndEpoch":10101,"StoragePricePerEpoch":"0","ProviderCollateral":"0","ClientCollateral":"0"},"State":{"SectorStartEpoch":10101,"LastUpdatedEpoch":10101,"SlashEpoch":10101}}`

### StateMinerAvailableBalance
There are not yet any comments for this method.

Inputs: `["t01234",[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `"0"`

### StateMinerDeadlines
There are not yet any comments for this method.

Inputs: `["t01234",[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"Due":["sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI=","sAI="]}`

### StateMinerFaults
There are not yet any comments for this method.

Inputs: `["t01234",[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `"sAI="`

### StateMinerInfo
There are not yet any comments for this method.

Inputs: `["t01234",[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"Owner":"t01234","Worker":"t01234","PendingWorkerKey":{"NewWorker":"t01234","EffectiveAt":10101},"PeerId":"12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf","SealProofType":2,"SectorSize":34359738368,"WindowPoStPartitionSectors":42}`

### StateMinerInitialPledgeCollateral
There are not yet any comments for this method.

Inputs: `["t01234",9,[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `"0"`

### StateMinerPower
There are not yet any comments for this method.

Inputs: `["t01234",[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"MinerPower":{"RawBytePower":"0","QualityAdjPower":"0"},"TotalPower":{"RawBytePower":"0","QualityAdjPower":"0"}}`

### StateMinerProvingDeadline
There are not yet any comments for this method.

Inputs: `["t01234",[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"CurrentEpoch":10101,"PeriodStart":10101,"Index":42,"Open":10101,"Close":10101,"Challenge":10101,"FaultCutoff":10101}`

### StateMinerProvingSet
There are not yet any comments for this method.

Inputs: `["t01234",[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `null`

### StateMinerRecoveries
There are not yet any comments for this method.

Inputs: `["t01234",[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `"sAI="`

### StateMinerSectorCount
There are not yet any comments for this method.

Inputs: `["t01234",[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"Sset":42,"Pset":42}`

### StateMinerSectors
There are not yet any comments for this method.

Inputs: `["t01234","sAI=",true,[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `null`

### StateNetworkName
There are not yet any comments for this method.

Inputs: `null`

Response: `"lotus"`

### StatePledgeCollateral
There are not yet any comments for this method.

Inputs: `[[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `"0"`

### StateReadState
There are not yet any comments for this method.

Inputs: `[{"Code":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"Head":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"Nonce":42,"Balance":"0"},[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"Balance":"0","State":{}}`

### StateReplay
There are not yet any comments for this method.

Inputs: `[[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}],{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `{"Msg":{"Version":9,"To":"t01234","From":"t01234","Nonce":42,"Value":"0","GasPrice":"0","GasLimit":9,"Method":1,"Params":"Ynl0ZSBhcnJheQ=="},"MsgRct":{"ExitCode":0,"Return":"Ynl0ZSBhcnJheQ==","GasUsed":9},"InternalExecutions":null,"Error":"string value","Duration":60000000000}`

### StateSearchMsg
There are not yet any comments for this method.

Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `{"Receipt":{"ExitCode":0,"Return":"Ynl0ZSBhcnJheQ==","GasUsed":9},"TipSet":{"Cids":null,"Blocks":null,"Height":0}}`

### StateSectorPreCommitInfo
There are not yet any comments for this method.

Inputs: `["t01234",9,[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},{"/":"bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"}]]`

Response: `{"Info":{"RegisteredProof":2,"SectorNumber":9,"SealedCID":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"SealRandEpoch":10101,"DealIDs":null,"Expiration":10101},"PreCommitDeposit":"0","PreCommitEpoch":10101}`

### StateWaitMsg
There are not yet any comments for this method.

Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `{"Receipt":{"ExitCode":0,"Return":"Ynl0ZSBhcnJheQ==","GasUsed":9},"TipSet":{"Cids":null,"Blocks":null,"Height":0}}`

## Sync
The Sync method group contains methods for interacting with and
observing the lotus sync service


### SyncCheckBad
There are not yet any comments for this method.

Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `"string value"`

### SyncIncomingBlocks
There are not yet any comments for this method.

Inputs: `null`

Response: `{"Miner":"t01234","Ticket":{"VRFProof":"Ynl0ZSBhcnJheQ=="},"ElectionProof":{"VRFProof":"Ynl0ZSBhcnJheQ=="},"BeaconEntries":null,"WinPoStProof":null,"Parents":null,"ParentWeight":"0","Height":10101,"ParentStateRoot":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"ParentMessageReceipts":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"Messages":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"BLSAggregate":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="},"Timestamp":42,"BlockSig":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="},"ForkSignaling":42}`

### SyncMarkBad
There are not yet any comments for this method.

Inputs: `[{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"}]`

Response: `{}`

### SyncState
SyncState returns the current status of the lotus sync system


Inputs: `null`

Response: `{"ActiveSyncs":null}`

### SyncSubmitBlock
SyncSubmitBlock can be used to submit a newly created block to the
network through this node


Inputs: `[{"Header":{"Miner":"t01234","Ticket":{"VRFProof":"Ynl0ZSBhcnJheQ=="},"ElectionProof":{"VRFProof":"Ynl0ZSBhcnJheQ=="},"BeaconEntries":null,"WinPoStProof":null,"Parents":null,"ParentWeight":"0","Height":10101,"ParentStateRoot":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"ParentMessageReceipts":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"Messages":{"/":"bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"},"BLSAggregate":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="},"Timestamp":42,"BlockSig":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="},"ForkSignaling":42},"BlsMessages":null,"SecpkMessages":null}]`

Response: `{}`

## Wallet


### WalletBalance
There are not yet any comments for this method.

Inputs: `["t01234"]`

Response: `"0"`

### WalletDefaultAddress
There are not yet any comments for this method.

Inputs: `null`

Response: `"t01234"`

### WalletExport
There are not yet any comments for this method.

Inputs: `["t01234"]`

Response: `{"Type":"string value","PrivateKey":"Ynl0ZSBhcnJheQ=="}`

### WalletHas
There are not yet any comments for this method.

Inputs: `["t01234"]`

Response: `true`

### WalletImport
There are not yet any comments for this method.

Inputs: `[{"Type":"string value","PrivateKey":"Ynl0ZSBhcnJheQ=="}]`

Response: `"t01234"`

### WalletList
There are not yet any comments for this method.

Inputs: `null`

Response: `null`

### WalletNew
There are not yet any comments for this method.

Inputs: `[2]`

Response: `"t01234"`

### WalletSetDefault
There are not yet any comments for this method.

Inputs: `["t01234"]`

Response: `{}`

### WalletSign
There are not yet any comments for this method.

Inputs: `["t01234","Ynl0ZSBhcnJheQ=="]`

Response: `{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="}`

### WalletSignMessage
There are not yet any comments for this method.

Inputs: `["t01234",{"Version":9,"To":"t01234","From":"t01234","Nonce":42,"Value":"0","GasPrice":"0","GasLimit":9,"Method":1,"Params":"Ynl0ZSBhcnJheQ=="}]`

Response: `{"Message":{"Version":9,"To":"t01234","From":"t01234","Nonce":42,"Value":"0","GasPrice":"0","GasLimit":9,"Method":1,"Params":"Ynl0ZSBhcnJheQ=="},"Signature":{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="}}`

### WalletVerify
There are not yet any comments for this method.

Inputs: `["t01234","Ynl0ZSBhcnJheQ==",{"Type":2,"Data":"Ynl0ZSBhcnJheQ=="}]`

Response: `true`

