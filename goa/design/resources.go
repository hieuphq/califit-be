package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("authentication", func() {
	BasePath("/auth")
	Action("login", func() {
		NoSecurity()
		Routing(
			POST("/login"),
		)
		Description("Sign a company user in")
		Payload(LoginPayload)
		Response(OK, LoginRes)
		Response(InternalServerError)
		Response(BadRequest, ErrorMedia)
	})

	Action("logout", func() {
		NoSecurity()
		Routing(
			POST("/logout"),
		)
		Description("Sign a company user in")
		Payload(LoginPayload)
		Response(OK, LoginRes)
		Response(InternalServerError)
		Response(BadRequest, ErrorMedia)
	})

	Action("register", func() {
		NoSecurity()
		Routing(
			POST("/register"),
		)
		Description("Create a new user")
		Payload(RegisterPayload)
		Response(OK, Token)
		Response(InternalServerError)
		Response(BadRequest, ErrorMedia)
	})
})
