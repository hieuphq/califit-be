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

var _ = Resource("city", func() {
	BasePath("/city")

	Action("list", func() {
		Routing(GET(""))
		Params(func() {
			Param("limit", Integer, "limit for paginate", func() {
				Default(10)
			})
			Param("offset", Integer, "offset for paginate", func() {
				Default(0)
			})
			Param("name", String, "query for find cities by name", func() {
				Default("")
			})
		})
		Description("List cities")
		Response(OK, CityList)
		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("show", func() {
		Routing(GET("/:cityID"))
		Params(func() {
			Param("cityID", String, "City ID")
		})
		Description("Get a city by ID")
		Response(OK, City)
		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})

var _ = Resource("center", func() {
	BasePath("/center")

	Action("list", func() {
		Routing(GET(""))
		Params(func() {
			Param("limit", Integer, "limit for paginate", func() {
				Default(10)
			})
			Param("offset", Integer, "offset for paginate", func() {
				Default(0)
			})
			Param("name", String, "query for find centers by name", func() {
				Default("")
			})
			Param("cityID", String, "query for find centers in city", func() {
				Default("")
			})
		})
		Description("List centers")
		Response(OK, CenterList)
		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("show", func() {
		Routing(GET("/:centerID"))
		Params(func() {
			Param("centerID", String, "Center ID")
		})
		Description("Get a center by ID")
		Response(OK, Center)
		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})
