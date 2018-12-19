package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Token token response
var Token = MediaType("application/vnd.token+json", func() {
	Description("A token")
	Attributes(func() {
		Attribute("token", String, "A JWT token")
	})

	View("default", func() {
		Attribute("token")
	})
})

// LoginRes login response
var LoginRes = MediaType("application/vnd.login_response+json", func() {
	Description("Login response")

	Attribute("token", String)
	Attribute("user_info", UserInfo)

	View("default", func() {
		Attribute("token")
		Attribute("user_info")
	})
})

// UserInfo user response
var UserInfo = MediaType("application/vnd.user_info+json", func() {
	Description("User info")

	Attribute("contact_name", String)
	Attribute("email", String)
	Attribute("phone", String)
	Attribute("company_name", String)
	Attribute("company_address", String)

	View("default", func() {
		Attribute("contact_name")
		Attribute("email")
		Attribute("phone")
		Attribute("company_name")
		Attribute("company_address")
	})
})

// City response data
var City = MediaType("application/vnd.city+json", func() {
	Description("a city detail")

	Attribute("city_id", String)
	Attribute("city_name", String)
	Attribute("created_at", DateTime)

	View("default", func() {
		Attribute("city_id", String)
		Attribute("city_name", String)
		Attribute("created_at", DateTime)
	})

	View("general", func() {
		Attribute("city_id", String)
		Attribute("city_name", String)
		Attribute("created_at", DateTime)
	})
})

// CityList enum of city
var CityList = MediaType("application/vnd.cities+json", func() {
	Description("A list of City")

	Attribute("data", ArrayOf(City, func() {
		View("general")
	}))
	Attribute("paginate", Paginate)

	View("default", func() {
		Attribute("data")
		Attribute("paginate")
	})
})

// Paginate response
var Paginate = MediaType("application/vnd.paginate+json", func() {
	Description("pagination of a resources")
	Attribute("current_page", Integer)
	Attribute("total_page", Integer)
	Attribute("total_item", Integer)

	View("default", func() {
		Attribute("current_page")
		Attribute("total_page")
		Attribute("total_item")
	})
})
