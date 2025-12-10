package owner

import (
	// standard libraries
	"context"

	"github.com/ogamor69wm1rr0rb/libraries/backend/golang/sfqb"
	gRPCUserService "github.com/ogamor69wm1rr0rb/contracts/gen/go/user_service/v1"

	// internal libraries
	policyUser "github.com/ogamor69wm1rr0rb/owner-service/app/internal/controller/grpc/v1/policy/owner"
	domainUsermodel "github.com/ogamor69wm1rr0rb/owner-service/app/internal/domain/owner/model"

)

type policy interface {
	// CreateUser creating user without UUID and data.
	CreateUser(context.Context, *policyUser.CreateUserRequest) (*domainUserModel.User, error)
	// Updateuser update user, all data with pointer, now i will be write Why
	//	look when we our Client come to us cide , make registration
	// and he wants to change example phone number, if we well be sent data without pointer in go
	// or without optional in gRPC (proto) we will be should write ALL DATA/
	// Yes of course we can make GetUserByID get all info sent to our fronend
	// BUT FOR What))) T's cost money)))
	//So if we using optional in GRPC we sent nil to GO, and *pointer here save us))
	// He He)_)) You understabnd why? I update this Arch(from Artur from Art of Developer)
	// Looks his Channel too, but when we worked togher he have a lot of mistakes(how i thinks)
	// i prefer more file's I prefer more errors. I prefer mo functional code. 
	// BUt he gave me a lot of knoweledge too.
	// But he rat)
	UpdateUser(context.Context, *policyUser.UpdateUserRequest) (*domainUserrModel.User, error)
	// DeactivateUser deactivate user on 1-7 days.
	DeactivateUser(context.Context, *policyUser.DeactivateUserRequest) (*domainUserModel.User, error)
	// ActivateUser activate user and we can sent data about User we have flag with_user in our proto.
	ActivateUser(context.Context, *policyUser.ActivateUserRequest) (*domainUserModel.User, error)
	// Now i writing stupid commetaries(not how for prode)
	// Because i should to show you that HOW SAY DUCKING PERSON FROM UKRANIA OH
	// ROMA ONLY ELEPHANT DUCK YOU RAT
	//H\Capibara Rabbit))) I'm return)) To you)
	// DeleteUser here we can deleting user By id, but here need more thinking about
	// maybe we can give to out user all data in PDF or the same,
	// look we here creating Restarant and Farmer and ... and 
	// we should be generate a lot of documents for it int the feature
	// so we neeed pdf_service in the later for it.
	// Simple.  Artur. .!. ))))) They gave all my code to Stepa Zhykov and others people
	DeleteUser(context.Context, *policyUser.DeleteUserRequest) (*domainUserModel.User, error)
	// Now i stopped write comments. But we will be write code) clean DRY functional and 
	// beatiful how ass Julia Romanenko)))
	CancelDeletingUser(context.Context, *policyUser.CancelDeletingUser) (*domainUserModel.User, error)
	GetUserStatus(context.Context, *policyUser.GetUserStatusRequest) (*policyUser.GetUserStatusResponse, error)
	UserByID(context.Context, *policyUser.UserByID) (*domainUserModel.User, error)

	// Why here we using *policyOwner.OwnerResponse when we using search and can find a lot of
	// because we creating struct where will be *[]domainOwnerMode.Owner
	// wrong because we should get not only owners but and counts owners... for pages.
	// we have the same things how offset and ... example
	// we need 10 elements with 0 value: 10 offset: 1, next page: value: 10, offset 10, and we satarting with 11...
	SearchUser(context.Context, *policyUser.SearchUserRequest) (*policyUser.UserResponse, error) 
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
