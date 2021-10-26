package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"jmind-test/src/models"
	"jmind-test/src/utils"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Etherscan struct{}

var EtherscanApi *Etherscan

type Transaction struct {
	Value string
}

type Block struct {
	BlockNumber int
	Result      *struct {
		Transactions []*Transaction
	}
}

func (_ *Etherscan) GetBlockByNumberRequest(blockNumber int) (*Block, error) {
	blockData := Block{
		BlockNumber: blockNumber,
	}

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

func (block *Block) GetBlockTotals() (*models.BlockTotals, error) {
	totalAmount := big.NewFloat(0)
	transactionsCount := len(block.Result.Transactions)

	for _, transaction := range block.Result.Transactions {
		numberStr := strings.Replace(transaction.Value, "0x", "", -1)

		num := new(big.Int)
		num.SetString(numberStr, 16)

		totalAmount.Add(totalAmount, utils.WeiToEther(num))
	}

	totalAmountFloat64, _ := totalAmount.Float64()

	return &models.BlockTotals{
		BlockNumber:  block.BlockNumber,
		Transactions: transactionsCount,
		Amount:       totalAmountFloat64,
	}, nil
}
