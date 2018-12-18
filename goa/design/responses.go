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
