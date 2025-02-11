package config

import (
	"errors"
	"fmt"
)

type ConsumerChain struct {
	Name                string     `toml:"name"`
	LCDEndpoint         string     `toml:"lcd-endpoint"`
	BaseDenom           string     `toml:"base-denom"`
	Denoms              DenomInfos `toml:"denoms"`
	ChainID             string     `toml:"chain-id"`
	BechWalletPrefix    string     `toml:"bech-wallet-prefix"`
	BechValidatorPrefix string     `toml:"bech-validator-prefix"`
	BechConsensusPrefix string     `toml:"bech-consensus-prefix"`
	Queries             Queries    `toml:"queries"`
}

func (c *ConsumerChain) GetQueries() Queries {
	return c.Queries
}

func (c *ConsumerChain) GetHost() string {
	return c.LCDEndpoint
}

func (c *ConsumerChain) GetName() string {
	return c.Name
}

func (c *ConsumerChain) Validate() error {
	if c.Name == "" {
		return errors.New("empty chain name")
	}

	if c.LCDEndpoint == "" {
		return errors.New("no LCD endpoint provided")
	}

	if c.ChainID == "" {
		return errors.New("no chain-id provided")
	}

	for index, denomInfo := range c.Denoms {
		if err := denomInfo.Validate(); err != nil {
			return fmt.Errorf("error in denom #%d: %s", index, err)
		}
	}

	return nil
}
