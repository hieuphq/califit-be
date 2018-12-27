package design

import (
	"github.com/goadesign/goa/design/apidsl"
)

var _ = apidsl.API("califit-be Backend API", func() {
	apidsl.Title("The califit-be Backend API")
	apidsl.Description("An API for califit-be")
	apidsl.Contact(func() {
		apidsl.Name("Hieu Phan")
		apidsl.Email("hieu.phq@gmail.com")
	})
	apidsl.Host("localhost:9000")
	apidsl.Scheme("http")
	apidsl.BasePath("/api/")
	apidsl.Origin("*", func() {
		apidsl.Methods("GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS")
		apidsl.Headers("Accept", "Content-Type", "Authorization")
		apidsl.Expose("Content-Type", "Origin")
		apidsl.MaxAge(600)
		apidsl.Credentials()
	})
	apidsl.Consumes("application/json")
	apidsl.Produces("application/json")
})
