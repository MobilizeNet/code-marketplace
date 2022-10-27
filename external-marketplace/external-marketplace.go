package externalmarketplace

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/coder/code-marketplace/src/api/httpapi"
	"github.com/coder/code-marketplace/src/api/httpmw"
)

type ExternalMarketplaceInterface interface {
	HttpRequest(httpMethod string, relativeUrl string, bodyReader io.Reader) (int, []byte)
}

type ExternalMarketplace struct {
	BaseUrl string
}

func (extmarket *ExternalMarketplace) HttpRequest(httpMethod string, relativeUrl string, bodyReader io.Reader) (int, []byte) {

	url := extmarket.BaseUrl + relativeUrl

	// Gets the extensions from the Open Source marketplace (OpenVsx)
	req, reqErr := http.NewRequest(httpMethod, url, bodyReader)
	if reqErr != nil {
		bytes, _ := json.Marshal(httpapi.ErrorResponse{
			Message:   "Unable to create the new request",
			Detail:    "Check that the request boyd is valid JSON or the http method (" + httpMethod + ") is correct",
			RequestID: httpmw.RequestID(req),
		})
		return http.StatusBadRequest, bytes
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		bytes, _ := json.Marshal(httpapi.ErrorResponse{
			Message: "Unable to call " + url,
			Detail:  "Check the Url or the connection to internet",
		})
		return http.StatusInternalServerError, bytes
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return res.StatusCode, body
}
