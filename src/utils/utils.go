package utils

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
	"net/http"
)

type HttpError struct {
	StatusCode int
	Message    string
}

func (e *HttpError) Error() string {
	return e.Message
}

func ShowHttpError(error *HttpError, w http.ResponseWriter) {
	log.Printf("Error %d: %s", error.StatusCode, error.Message)

	errObject := map[string]interface{}{
		"status_code": error.StatusCode,
		"message":     error.Message,
	}
	msg, _ := json.Marshal(errObject)

	w.WriteHeader(error.StatusCode)
	w.Write(msg)
}

func WeiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
}
