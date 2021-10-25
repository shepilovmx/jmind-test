package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"jmind-test/src/components"
	"jmind-test/src/utils"
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

	block, err := components.EtherscanApi.GetBlockByNumberRequest(blockNumber)
	if err != nil {
		return &utils.HttpError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	totals, err := components.EtherscanApi.GetBlockTotals(block)
	if err != nil {
		return &utils.HttpError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	responseData, err := json.Marshal(totals)

	return self.SendResponse(w, responseData)
}
