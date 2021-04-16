package test

import (
	"os"
	"testing"

	"github.com/davidandradeduarte/uif-testing-in-golang-udemy/src/api/app"
	"github.com/mercadolibre/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	go app.StartApp()
	os.Exit(m.Run())
}
