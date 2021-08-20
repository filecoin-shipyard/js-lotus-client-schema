const {
  CommonMethods,
  FullNodeMethods,
  StorageMinerMethods,
  GatewayMethods,
  WalletMethods,
  WorkerMethods
} = require('./gen.js')

const common = {
  methods: CommonMethods
}

const fullNode = {
  methods: {
    ...CommonMethods,
    ...FullNodeMethods
  }
}

const storageMiner = {
  methods: {
    ...CommonMethods,
    ...StorageMinerMethods
  }
}

const gatewayApi = {
  methods: GatewayMethods
}

const walletApi = {
  methods: WalletMethods
}

const workerApi = {
  methods: WorkerMethods
}

module.exports = {
  common,
  fullNode,
  storageMiner,
  gatewayApi,
  walletApi,
  workerApi
}
