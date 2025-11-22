package owner

import (
	gRPCOwnerService "github.com/ogamor69wm1rr0rb/contracts/gen/go/owner_service/v1"

	policyOwner "github.com/ogamor69wm1rr0rb/owner-service/app/internal/controller/grpc/v1/policy/owner"
	domainOwnermodel "github.com/ogamor69wm1rr0rb/owner-service/app/internal/domain/owner/model"
)

func convertCreateOwnerToPB(*policyOwner.CreateOwnerRequest) *gRPCOwnerService.CreateOwnerResponse {
	
}
