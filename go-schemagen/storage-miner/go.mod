module github.com/filecoin-shipyard/js-lotus-client-schema/go-schemagen/storage-miner

go 1.14

require (
	github.com/filecoin-project/go-address v0.0.2-0.20200504173055-8b6f2fb2b3ef
	github.com/filecoin-project/go-bitfield v0.0.4-0.20200703174658-f4a5758051a1
	github.com/filecoin-project/go-jsonrpc v0.1.1-0.20200602181149-522144ab4e24
	github.com/filecoin-project/lotus v0.3.1-0.20200521170719-0d3f602d58eb
	github.com/filecoin-project/specs-actors v0.7.1
	github.com/ipfs/go-cid v0.0.6
	github.com/ipfs/go-filestore v1.0.0
	github.com/libp2p/go-libp2p-core v0.6.0
	github.com/libp2p/go-libp2p-peer v0.2.0
	github.com/multiformats/go-multiaddr v0.2.2
)

replace github.com/filecoin-project/lotus => ../../lotus
