package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// LoginPayload payload for login
var LoginPayload = Type("LoginPayload", func() {
	Attribute("email", String, func() {
		MinLength(6)
		MaxLength(400)
		Format("email")
		Example("jamesbond@gmail.com")
	})

	Attribute("password", String, func() {
		MinLength(8)
		MaxLength(100)
		Example("abcd1234")
	})
	Required("email", "password")
})

// RegisterPayload payload for register
var RegisterPayload = Type("RegisterPayload", func() {
	Attribute("email", String, func() {
		MinLength(6)
		MaxLength(150)
		Format("email")
		Example("jamesbond@gmail.com")
	})

	Attribute("contact_name", String, func() {
		MinLength(1)
		MaxLength(200)
		Example("Doe")
	})

	Attribute("phone", String, func() {
		MaxLength(50)
	})

	Attribute("password", String, func() {
		MinLength(8)
		MaxLength(100)
		Example("abcd1234")
	})

	Attribute("company_address", String, func() {
		MinLength(1)
		MaxLength(500)
		Example("445 Mount Eden Road, Mount Eden, Auckland")
	})

	Attribute("company_name", String, func() {
		MinLength(1)
		MaxLength(500)
		Example("Express company")
	})

	Required("email", "password", "contact_name", "phone", "company_address", "company_name")
})
