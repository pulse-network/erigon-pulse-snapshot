package snapshothashes

import (
	_ "embed"
)

//go:embed pulsechain-mainnet.toml
var PulseChainMainnet []byte

//go:embed pulsechain-testnet-v4.toml
var PulseChainTestnetV4 []byte

//go:embed history/pulsechain-mainnet.toml
var PulseChainMainnetHistory []byte

//go:embed history/pulsechain-testnet-v4.toml
var PulseChainTestnetV4History []byte
