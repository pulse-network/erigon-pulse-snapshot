package snapshothashes

import (
	_ "embed"
)

//go:embed pulsechain-mainnet.toml
var PulseChainMainnet []byte

//go:embed pulsechain-testnet-v3.toml
var PulseChainTestnetV3 []byte

//go:embed history/pulsechain-mainnet.toml
var PulseChainMainnetHistory []byte

//go:embed history/pulsechain-testnet-v3.toml
var PulseChainTestnetV3History []byte
