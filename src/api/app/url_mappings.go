package app

import "github.com/davidandradeduarte/uif-testing-in-golang-udemy/src/api/controllers"

func mapUrls() {
	router.GET("/locations/countries/:country_id", controllers.GetCountry)
}
