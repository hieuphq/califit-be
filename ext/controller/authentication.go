package controller

import (
	"github.com/goadesign/goa"
	"github.com/hieuphq/califit-be/goa/app"
)

// AuthenticationController implements the authentication resource.
type AuthenticationController struct {
	*goa.Controller
}

// NewAuthenticationController creates a authentication controller.
func NewAuthenticationController(service *goa.Service) *AuthenticationController {
	return &AuthenticationController{Controller: service.NewController("AuthenticationController")}
}

// Login runs the login action.
func (c *AuthenticationController) Login(ctx *app.LoginAuthenticationContext) error {
	// AuthenticationController_Login: start_implement

	// Put your logic here

	res := &app.LoginResponse{}
	return ctx.OK(res)
	// AuthenticationController_Login: end_implement
}

// Logout runs the logout action.
func (c *AuthenticationController) Logout(ctx *app.LogoutAuthenticationContext) error {
	// AuthenticationController_Logout: start_implement

	// Put your logic here

	res := &app.LoginResponse{}
	return ctx.OK(res)
	// AuthenticationController_Logout: end_implement
}

// Register runs the register action.
func (c *AuthenticationController) Register(ctx *app.RegisterAuthenticationContext) error {
	// AuthenticationController_Register: start_implement

	// Put your logic here

	res := &app.Token{}
	return ctx.OK(res)
	// AuthenticationController_Register: end_implement
}
