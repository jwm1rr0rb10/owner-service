package owner

import (
	// standard libraries
	"context"

	"github.com/ogamor69wm1rr0rb/libraries/backend/golang/sfqb"
	gRPCOwnerService "github.com/ogamor69wm1rr0rb/contracts/gen/go/owner_service/v1"

	// internal libraries
	policyOwner "github.com/ogamor69wm1rr0rb/owner-service/app/internal/controller/grpc/v1/policy/owner"
	domainOwnermodel "github.com/ogamor69wm1rr0rb/owner-service/app/internal/domain/owner/model"

)

type policy interface {
	CreateOwner(context.Context, *policyOwner.CreateOwnerRequest) (*domainOwnermodel.Owner, error)
	UpdateOwner(context.Context, *policyOwner.UpdateOwnerRequest) (*domainOwnermodel.Owner, error)
	SwitchPrivateOwnerStatus(context.Context, *policyOwner.SwitchPrivateOwnerStatusRequest) (*domainOwnermodel.Owner, error)
	GetOwnerByID(context.Context, id string) (*domainOwnermodel.Owner, error)
	// Why here we using *policyOwner.OwnerResponse when we using search and can find a lot of
	// because we creating struct where will be *[]domainOwnerMode.Owner
	// wrong because we should get not only owners but and counts owners... for pages.
	// we have the same things how offset and ... example
	// we need 10 elements with 0 value: 10 offset: 1, next page: value: 10, offset 10, and we satarting with 11...
	SearchOwner(context.Context, *policyOwner.SearchOwnerRequest) (*policyOwner.OwnerResponse, error) 
}


type Controller struct {
	*gRPCOwnerService.UnimplementedOwnerServiceServer
	policy policy
}

func NewConstructorController(policy policy) *Controller{
	return &Controller{
		policy: policy,
	}
}
