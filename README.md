# Lotus JS Client Schema

This package contains some .js files that describe methods exported by the
[Lotus](https://github.com/filecoin-project/lotus) JSON-RPC API.

See [filecoin-shipyard/js-lotus-client](https://github.com/filecoin-shipyard/js-lotus-client)
for examples that show how to use these schemas.

## Usage

```js
const { LotusRPC } = require('@filecoin-shipyard/lotus-client-rpc')
const { NodejsProvider } = require('@filecoin-shipyard/lotus-client-provider-nodejs')
const { mainnet } = require('@filecoin-shipyard/lotus-client-schema')

const url = 'ws://127.0.0.1:7777/rpc/v0'
const provider = new NodejsProvider(url)
const client = new LotusRPC(provider, { schema: mainnet.fullNode })
```

## Schemas

The following schema objects are exported:

* `mainnet.fullNode` - used to talk to a Lotus "Full Node" (`lotus`)
* `mainnet.storageMiner` - used to talk to a Lotus storage miner (`lotus-miner`)
* `mainnet.gatewayApi` - used to talk to public Lotus API Gateway (`lotus-gateway`)
* `mainnet.walletApi` - a subset of methods for wallets
* `mainnet.workerApi` - used to talk to `lotus-worker` (for mining)
* `mainnet.common` - a subset of methods used by a number of different daemons

## Schema Format

Right now, the schema is extremely simple, as the JS library only needs to know if a method supports subscriptions or not. The JS library is very light, so all method parameters are just passed through to the JSON-RPC request.

In the future, the schema could be extended to provide more parameter and type information to make it easier to validate and marshal/unmarshal requests and responses.

For example, the entry for the `ChainHead` method is currently just:

```
  ChainHead: {}
```

The `ChainNotify` entry supports JSON-RPC subscriptions, so it looks like this:

```
  ChainNotify: {
    subscription: true
  }
```

Each API is represented by an Object with a single `methods` property which
is an Object with all the methods supported.

If you are trying to use a JSON-RPC API method that isn't in the schema yet,
don't be afraid to "monkey-patch" the schemas distributed via this package
to add any missing methods ... eg. `mainnet.fullNode.methods.MyMissingMethod = {}`

## Updating the schemas

Lotus is constantly being updated. There is a simple tool (written in go) in the `go-schemagen` directory that links against a Lotus source code tree (modify the replace directory in go.mod or make a symlink) and outputs the current set of method signatures as JSON.

## License

Dual-licensed under [MIT](https://github.com/filecoin-project/lotus/blob/master/LICENSE-MIT) +
[Apache 2.0](https://github.com/filecoin-project/lotus/blob/master/LICENSE-APACHE)
