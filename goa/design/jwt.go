package design

import (
	"github.com/goadesign/goa/design/apidsl"
)

// JWT config
var JWT = apidsl.JWTSecurity("jwt", func() {
	apidsl.Description("Use JWT to authenticate")
	apidsl.Header("Authorization")
	apidsl.Scope("api:access", "API access") // Define "api:access" scope
})
