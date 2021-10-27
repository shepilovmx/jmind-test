package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"jmind-test/src/models"
	"net/http"
	"net/url"
	"os"
)

type Etherscan struct{}

var EtherscanApi *Etherscan

func (_ *Etherscan) GetBlockByNumberRequest(blockNumber int) (*models.Block, error) {
	blockData := models.Block{
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
