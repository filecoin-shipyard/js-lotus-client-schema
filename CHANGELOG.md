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
