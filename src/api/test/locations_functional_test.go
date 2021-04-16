package test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/davidandradeduarte/uif-testing-in-golang-udemy/src/api/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestGetCountriesNotFound(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"Country not found","error":"not_found","status":404,"cause":[]}`,
	})

	response, err := http.Get("localhost:8080/locations/countries/AR")
	assert.Nil(t, err)
	assert.NotNil(t, response)

	var apiErr errors.ApiError
	bytes, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(bytes, &apiErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "not_found", apiErr.Error)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}
