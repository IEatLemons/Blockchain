package ganache

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (G *Ganache) Balance(Address string, Block *big.Int) (balance *big.Int, err error) {
	account := common.HexToAddress(Address)
	balance, err = G.Cli.BalanceAt(context.Background(), account, Block)
	return
}
