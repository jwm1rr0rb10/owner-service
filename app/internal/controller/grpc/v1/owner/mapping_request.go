package owner

import (
	gRPCOwnerService "github.com/ogamor69wm1rr0rb/contracts/gen/go/owner_service/v1"
	policyOwner "github.com/ogamor69wm1rr0rb/owner-service/app/internal/controller/grpc/v1/policy/owner"
)

func convertCreateownerToStruct(*gRPCOwnerService.CreateOwnerRequest) *policyOwner.CreateOwnerRequest {

}