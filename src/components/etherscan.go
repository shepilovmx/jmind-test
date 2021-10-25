package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"jmind-test/src/utils"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Etherscan struct{}

var EtherscanApi Etherscan

type Transaction struct {
	Value string
}

type Block struct {
	Result struct {
		Transactions []Transaction
	}
}

type BlockTotals struct {
	Transactions int        `json:"transactions"`
	Amount       *big.Float `json:"amount"`
}

func (_ *Etherscan) GetBlockByNumberRequest(blockNumber int) (*Block, error) {
	var blockData Block

	payload := url.Values{
		"module":  []string{"proxy"},
		"action":  []string{"eth_getBlockByNumber"},
		"tag":     []string{fmt.Sprintf("%x", blockNumber)},
		"boolean": []string{"true"},
		"apikey":  []string{os.Getenv("ETHERSCAN_API_KEY")},
	}

	response, err := http.Get(os.Getenv("ETHERSCAN_API_URL") + "?" + payload.Encode())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		return nil, errors.New(fmt.Sprintf("Failed to get block. Response: %s", string(body)))
	}

	dec := json.NewDecoder(response.Body)
	err = dec.Decode(&blockData)
	if err != nil {
		return nil, err
	}

	return &blockData, nil
}

func (_ *Etherscan) GetBlockTotals(block *Block) (*BlockTotals, error) {
	totalAmount := big.NewFloat(0)

	for _, transaction := range block.Result.Transactions {
		numberStr := strings.Replace(transaction.Value, "0x", "", -1)

		num := new(big.Int)
		num.SetString(numberStr, 16)

		totalAmount.Add(totalAmount, utils.WeiToEther(num))
	}

	return &BlockTotals{
		Transactions: len(block.Result.Transactions),
		Amount:       totalAmount,
	}, nil
}
