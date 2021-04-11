package main

import (
	"fmt"

	"github.com/davidandradeduarte/uif-testing-in-golang-udemy/src/api/providers/locations_provider"
)

func main() {
	country, err := locations_provider.GetCountry("AR")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(country)
}
