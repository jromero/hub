// Code generated by goa v3.2.2, DO NOT EDIT.
//
// auth service
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package auth

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// The auth service exposes endpoint to authenticate User against GitHub OAuth
type Service interface {
	// Authenticates users against GitHub OAuth
	Authenticate(context.Context, *AuthenticatePayload) (res *AuthenticateResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "auth"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"Authenticate"}

// AuthenticatePayload is the payload type of the auth service Authenticate
// method.
type AuthenticatePayload struct {
	// OAuth Authorization code of User
	Code string
}

// AuthenticateResult is the result type of the auth service Authenticate
// method.
type AuthenticateResult struct {
	// JWT
	Token string
}

// MakeInvalidCode builds a goa.ServiceError from an error.
func MakeInvalidCode(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "invalid-code",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "internal-error",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}