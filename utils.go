package etherscan

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type BigInt struct {
	big.Int
}

func (bigint *BigInt) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	bigint.SetString(s, 10)

	return nil
}

func urlByNetworkID(id NetworkID) (string, error) {
	switch id {
	case Mainnet:
		return baseUrlMainnet, nil
	case Rinkeby:
		return baseUrlRinkeby, nil
	}

	return "", fmt.Errorf("unknown network id %d", id)
}
