package controllers

import (
	"encoding/json"
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"jmind-test/src/components"
	"jmind-test/src/repositories"
	"jmind-test/src/utils"
	"log"
	"net/http"
	"strconv"
)

type BlockController struct {
	Controller
}

var Block BlockController

func (self *BlockController) Total(w http.ResponseWriter, r *http.Request) *utils.HttpError {
	blockNumber, err := strconv.Atoi(mux.Vars(r)["blockNumber"])
	if err != nil {
		return &utils.HttpError{StatusCode: http.StatusBadRequest, Message: "Invalid block number."}
	}

	log.Printf("Searching for block %d in DB...", blockNumber)

	totals, err := repositories.BlockTotalsRep.FindTotalsByBlockNumber(blockNumber)
	if err != nil && err != pg.ErrNoRows {
		return &utils.HttpError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	} else if err == pg.ErrNoRows {
		log.Printf("Requesting for block %d...", blockNumber)

		block, err := components.EtherscanApi.GetBlockByNumberRequest(blockNumber)
		if err != nil {
			return &utils.HttpError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		} else if block.Result == nil {
			return &utils.HttpError{StatusCode: http.StatusNotFound, Message: "Requested block doesn't exist."}
		}

		log.Printf("Block %d has been received.", blockNumber)

		totals, err = block.GetBlockTotals()
		if err != nil {
			return &utils.HttpError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		}

		_, err = repositories.BlockTotalsRep.InsertNewBlockTotals(totals)
		if err != nil {
			return &utils.HttpError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		}
	}

	log.Printf("Block %d totals: 1) Transactions count: %d; 2) Amount: %f.",
		blockNumber, totals.Transactions, totals.Amount)

	responseData, err := json.Marshal(totals)
	return self.SendResponse(w, responseData)
}
