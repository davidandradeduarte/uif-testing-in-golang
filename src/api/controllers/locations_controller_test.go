package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/davidandradeduarte/uif-testing-in-golang-udemy/src/api/domain/locations"
	"github.com/davidandradeduarte/uif-testing-in-golang-udemy/src/api/services"
	"github.com/davidandradeduarte/uif-testing-in-golang-udemy/src/api/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

type locationServiceMock struct{}

var (
	getCountryFunc func(countryId string) (*locations.Country, *errors.ApiError)
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func (s *locationServiceMock) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return getCountryFunc(countryId)
}

func TestGetCountryNotFound(t *testing.T) {
	// Mock locations service
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return nil, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: "Country not found",
		}
	}
	services.LocationsService = &locationServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}
	if err != nil {
		panic(err)
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusNotFound, response.Code)

	var apiErr errors.ApiError
	err = json.Unmarshal(response.Body.Bytes(), &apiErr)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}
