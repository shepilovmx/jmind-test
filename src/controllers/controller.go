package controllers

import (
	"jmind-test/src/utils"
	"net/http"
)

type Action func(rw http.ResponseWriter, r *http.Request) *utils.HttpError

type Controller struct{}

func (c *Controller) Perform(a Action) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := a(w, r); err != nil {
			utils.ShowHttpError(err, w)
		}
	})
}

func (c *Controller) SendResponse(w http.ResponseWriter, responseData []byte) *utils.HttpError {
	_, err := w.Write(responseData)
	if err != nil {
		return &utils.HttpError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}
	return nil
}
