package snapshothashes

import (
	_ "embed"
)

//go:embed pulsechain-mainnet.toml
var PulseChainMainnet []byte

//go:embed pulsechain-testnet.toml
var PulseChainTestnet []byte

//go:embed history/pulsechain-mainnet.toml
var PulseChainMainnetHistory []byte

//go:embed history/pulsechain-testnet.toml
var PulseChainTestnetHistory []byte
