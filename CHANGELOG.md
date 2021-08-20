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
