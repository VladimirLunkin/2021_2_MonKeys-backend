package responses

import (
	"dripapp/internal/dripapp/models"
	"dripapp/internal/pkg/logger"
	"encoding/json"
	"log"
	"net/http"
)

type JSON struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}

func SendOKResp(resp JSON, w http.ResponseWriter) {
	byteResp, err := json.Marshal(resp)
	if err != nil {
		SendErrorResponse(w, models.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Error encoding json",
		},
			logger.DripLogger.ErrorLogging,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(byteResp)
	if err != nil {
		SendErrorResponse(w, models.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Error write byte",
		},
			logger.DripLogger.ErrorLogging,
		)
		return
	}
	log.Printf("CODE %d", resp.Status)

}

func SendErrorResponse(w http.ResponseWriter, httpErr models.HTTPError, logging func(int, string)) {
	var resp JSON
	resp.Status = httpErr.Code

	body, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logging(httpErr.Code, httpErr.Message)
}
