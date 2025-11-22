package owner

import (
	"context"

	domainOwnermodel "github.com/ogamor69wm1rr0rb/owner-service/app/internal/domain/owner/model"
)


type service interface {
	CreateOwner(context.Context, *domainOwnermodel.CreateOwnerInput) (*domainOwnermodel.Owner, error)
	UpdateOwner(context.Context, *domainOwnermodel.UpdateOwnerInput) (*domainOwnermodel.Owner, error)
	SwitchPrivateSOwnerStatus(context.Context, *domainOwnermodel.SwitchPrivateSOwnerStatusInput) (*)
}