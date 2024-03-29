## [4.2.0](https://github.com/filecoin-shipyard/js-lotus-client-schema/compare/v4.1.1...v4.2.0) (2022-03-15)


### Features

* update to lotus@1.15.0 ([25f04ef](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/25f04effc1c6973f56f1afc56dfa7adc79b66cc5))

### [4.1.1](https://github.com/filecoin-shipyard/js-lotus-client-schema/compare/v4.1.0...v4.1.1) (2022-03-15)


### Trivial Changes

* **deps:** bump actions/checkout from 2 to 3 ([#38](https://github.com/filecoin-shipyard/js-lotus-client-schema/issues/38)) ([766a4e3](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/766a4e3da4aae5dbf86e121101a82a29627a7c37))
* **no-release:** bump actions/checkout from 2 to 3 ([#36](https://github.com/filecoin-shipyard/js-lotus-client-schema/issues/36)) ([3f2ce11](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/3f2ce112dc5dd48a406fc0ce45284764ff6bee0b))
* **no-release:** bump actions/setup-node from 2.5.1 to 3.0.0 ([#35](https://github.com/filecoin-shipyard/js-lotus-client-schema/issues/35)) ([02509f2](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/02509f2d732340d31acfb611791c5b427a084ffc))

## [4.1.0](https://github.com/filecoin-shipyard/js-lotus-client-schema/compare/v4.0.7...v4.1.0) (2022-02-14)


### Features

* lotus@v1.13.2, and fix schema gen CJS export ([864168d](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/864168d7e07ca7faef61258499f06346b6382364))

### [4.0.7](https://github.com/filecoin-shipyard/js-lotus-client-schema/compare/v4.0.6...v4.0.7) (2022-01-20)


### Bug Fixes

* use 16.x for release script, as required by dual-publish ([23e7b74](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/23e7b740ab83098ff14095d49a83c3838a9096d2))


### Trivial Changes

* **deps-dev:** bump dual-publish from 2.0.2 to 3.0.0 ([6b7297c](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/6b7297caba22c9ef9623b94cd222cbe0968bc840))
* **no-release:** bump actions/setup-node from 2.4.1 to 2.5.0 ([#28](https://github.com/filecoin-shipyard/js-lotus-client-schema/issues/28)) ([afdf4c5](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/afdf4c543ac5f89d6b74a4e59507c4db367681dd))
* **no-release:** bump actions/setup-node from 2.5.0 to 2.5.1 ([#30](https://github.com/filecoin-shipyard/js-lotus-client-schema/issues/30)) ([2a00ddd](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/2a00ddd6f044601f0dc821c6e23457118159265a))

### [4.0.6](https://github.com/filecoin-shipyard/js-lotus-client-schema/compare/v4.0.5...v4.0.6) (2021-11-18)


### Bug Fixes

* **ci:** exclude main files from build on windows ([6d11bee](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/6d11bee6ec5dfd3014faab784a066f0d0c12c13a))

### [4.0.5](https://github.com/filecoin-shipyard/js-lotus-client-schema/compare/v4.0.4...v4.0.5) (2021-11-04)


### Trivial Changes

* **deps:** bump actions/checkout from 2.3.5 to 2.4.0 ([1e8f0a3](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/1e8f0a3598fc5531941d5230370731f8e4bee710))

### [4.0.4](https://github.com/filecoin-shipyard/js-lotus-client-schema/compare/v4.0.3...v4.0.4) (2021-10-18)


### Trivial Changes

* **deps:** bump actions/checkout from 2.3.4 to 2.3.5 ([052bf1e](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/052bf1ea5a3c9a157165c1ab890f1569ee19f5e2))

### [4.0.3](https://github.com/filecoin-shipyard/js-lotus-client-schema/compare/v4.0.2...v4.0.3) (2021-09-28)


### Trivial Changes

* **deps:** bump actions/setup-node from 2.4.0 to 2.4.1 ([6725260](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/6725260ea23fa7838c678698d6b9e169c39e45ed))

### [4.0.2](https://github.com/filecoin-shipyard/js-lotus-client-schema/compare/v4.0.1...v4.0.2) (2021-09-27)


### Trivial Changes

* **deps-dev:** bump dual-publish from 1.0.9 to 2.0.0 ([3395fd6](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/3395fd63248718245a591a100e5a2aa8fb2e0097))

### [4.0.1](https://github.com/filecoin-shipyard/js-lotus-client-schema/compare/v4.0.0...v4.0.1) (2021-08-23)


### Trivial Changes

* make allowance for web3bot 'sync' PRs ([c5b00c8](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/c5b00c83239860521e5eac4093310c6f9fbf9d56)), closes [#12](https://github.com/filecoin-shipyard/js-lotus-client-schema/issues/12)

## [4.0.0](https://github.com/filecoin-shipyard/js-lotus-client-schema/compare/v3.0.2...v4.0.0) (2021-08-20)


### ⚠ BREAKING CHANGES

* update tsgen for lotus@1.11.1
* update schemagen for lotus@1.11.1, re-gen schema, make gen script do more

### Features

* update schemagen for lotus@1.11.1, re-gen schema, make gen script do more ([a9b76d8](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/a9b76d86c67713b6fca9997855f9f2e441a68b16))
* update tsgen for lotus@1.11.1 ([5961f8a](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/5961f8ac176bff1fa7139d4dd9e74de4200bc359))


### Bug Fixes

* add "generated by" note for schemagen'd code ([85d47c5](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/85d47c5a2c086bc9448fb330d1b88624f53bb533))
* make dual-publish work again with generated code ([f63928b](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/f63928b1e343855e9a0dc99182df28af7d677d50))


### Trivial Changes

* move generated API methods to separate module ([d936879](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/d936879bebd146aca96e4366b0fdb86cf16e78b4))

### [3.0.2](https://github.com/filecoin-shipyard/js-lotus-client-schema/compare/v3.0.1...v3.0.2) (2021-08-20)


### Trivial Changes

* add releasing section to readme ([#9](https://github.com/filecoin-shipyard/js-lotus-client-schema/issues/9)) ([302ab28](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/302ab289de22d7e5f7344d1214095d67b54d0fac))
* semantic-release for auto-releases, and dependabot ([bcc77e6](https://github.com/filecoin-shipyard/js-lotus-client-schema/commit/bcc77e64f4b306879b67b2b5b5b383f38f603370))

# js-lotus-api-schema changelog


## [1.0.0] - 2020-10-19

### Changed
- Updated for mainnet. The export name has changed from `testnet` to `mainnet`:

    ```js
    const { mainnet } = require('@filecoin-shipyard/lotus-client-schema')
    ```
    or 
    ```js
    import { mainnet } from '@filecoin-shipyard/lotus-client-schema'
    ```

## [0.0.11] - 2020-07-10

### Changed
- Update `testnet` schema with methods from Lotus `next` branch.
