package models

import (
	"jmind-test/src/utils"
	"math/big"
	"strings"
)

type Transaction struct {
	Value string
}

type Block struct {
	BlockNumber int
	Result      *struct {
		Transactions []*Transaction
	}
}

func (block *Block) GetBlockTotals() *BlockTotals {
	totalAmount := big.NewFloat(0)
	transactionsCount := len(block.Result.Transactions)

	for _, transaction := range block.Result.Transactions {
		numberStr := strings.Replace(transaction.Value, "0x", "", -1)

		num := new(big.Int)
		num.SetString(numberStr, 16)

		totalAmount.Add(totalAmount, utils.WeiToEther(num))
	}

	totalAmountFloat64, _ := totalAmount.Float64()

	return &BlockTotals{
		BlockNumber:  block.BlockNumber,
		Transactions: transactionsCount,
		Amount:       totalAmountFloat64,
	}
}
