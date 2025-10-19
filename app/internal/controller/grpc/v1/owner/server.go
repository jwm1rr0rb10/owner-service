package owner

import (
	// standard libraries
	"context"

	//other libraries

	// internal libraries
	policyOwner "github.com/ogamor69wm1rr0rb/owner-service/app/internal/controller/grpc/v1/policy/owner"
)

type policy interface {
	CreateOwner(context.Context, *policyOwner.CreateOwnerRequest) (*policyOwner.CreateOwnerResponse, error)
}
