package owner

import (
	"context"

	"github.com/ogamor69wm1rr0rb/libraries/backend/golang/errors"
	"github.com/ogamor69wm1rr0rb/libraries/backend/golang/logging"
	gRPCOwnerService "github.com/ogamor69wm1rr0rb/contracts/gen/go/owner_service/v1"
)

func(c *Controller) CreateOwner(
	ctx context.Context,
	req *gRPCOwnerService.CreateOwnerRequest,
) (*gRPCOwnerServiceCreateOwnerResponse, error) {
	logging.L(ctx).Debug("Controller Create Owner")

	// now need build our contracts and make stcrut with all rows for convert data to Struct and to PB
	// we can return error, but here not matter. But sometimes we should, example,
	// i love decimal(lib) in Golang and when i using NUMERIC(how for price and big int)
	// I prefer using Decimal + NUMERIC(16, 8) Why? we know about size in DB and thinking about cost memory.
	// everythng about optimization our system and when i convert a lot of data type I can return mistakes.
	// but anytime when I working with error i make 
	// amount, amountErr := ... and lib for convert Float to Decimal(but)
	// if you wants working with big numb float not good choose.
	// prefer string.))) 
	// and else if you working with JS on FrontEnd you should know that JS calculate bignumbers how shit)))
	// and better working with big numbers on backends
	request := convertCreateownerToStruct(req)

	create, createErr := c.policy.CreateOwner(ctx, request)
	if createErr != nil {
		return &policyOwer.CreateOwnerRequest{}, errors.New("c.policy.CreateOwner")
	}

	response := convertCreateownerToPB(create)

	return response, nil
}