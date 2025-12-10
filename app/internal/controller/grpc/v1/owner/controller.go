package owner

import (
	"context"

	"github.com/ogamor69wm1rr0rb/libraries/backend/golang/errors"
	"github.com/ogamor69wm1rr0rb/libraries/backend/golang/logging"
	gRPCUserService "github.com/ogamor69wm1rr0rb/contracts/gen/go/user_service/v1"
)

func(c *Controller) CreateUser(
	ctx context.Context,
	req *gRPCUserService.CreateUserRequest,
) (*gRPCUserService.CreateUserResponse, error) {
	logging.L(ctx).Debug("Controller Create User")

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
	request := convertCreateUserToStruct(req)

	create, createErr := c.policy.CreateUser(ctx, request)
	if createErr != nil {
		return &domainUserModel.User{}, errors.New("c.policy.CreateUser")
	}

	response := convertCreateUserToPB(create)

	return response, nil
}

// let's go i will be write without comments will be more faster
// and later i will be write all comments everywhere..
func(c *Controller) UpdateUser(
	ctx context.Context,
	req *gRPCUserService.UpdateUserResponse,
) (*gRPCUserService.UpdateUserresponse, error) {
	logging.L(ctx).Debug("controller update user")

	request := convertUpdateUserToStruct(req)

	update, updateErr := c.policy.UpdateUser(ctx, request)
	if updateErr != nil {
		return &domainUserModel.User{}, errors.New("c.policy.UpdateUser")
	}

	response := convertUpdateUserToPB(update)

	return response, nil
}

func(c *Controller) DeactivateUser(
	ctx context.Context,
	req *gRPCUserService.UpdateUserResponse,
) (*gRPCUserService.UpdateUserresponse, error) {
	logging.L(ctx).Debug("controller update user")

	request := convertUpdateUserToStruct(req)

	update, updateErr := c.policy.UpdateUser(ctx, request)
	if updateErr != nil {
		return &domainUserModel.User{}, errors.New("c.policy.UpdateUser")
	}

	response := convertUpdateUserToPB(update)

	return response, nil
}


func(c *Controller) ActivateUser(
	ctx context.Context,
	req *gRPCUserService.ActivateUserResponse,
) (*gRPCUserService.UpdateUserresponse, error) {
	logging.L(ctx).Debug("controller update user")

	request := convertUpdateUserToStruct(req)


	update, updateErr := c.policy.UpdateUser(ctx, request)
	if updateErr != nil {
		return &domainUserModel.User{}, errors.New("c.policy.UpdateUser")
	}

	response := convertUpdateUserToPB(update)

	return response, nil
}

// So they asked Petr the second made a joke Philosoph bar the second
// Who hacked me in the past hacked me too in Georgia and asked in el Hostel made me bad
// Jenya from El Hostel Ilshat and Kirill from Sait petersburg
// The second person in the hostel from Georgia told about me 
// people from ?) Ingushi and Ukrniane scummers))
// People who watched for me and asked Petr too.
// Diana in the first day started told me About WB how and Misha00
// Misha was from Koval)))
// Stepa from Alla Tamoshevich))
// They buy him MackBook who - Anton Krot and others with broken MDM))
// When i wrote code)))
// And tested Algotrading and wrote Restarants and made projectas about union EU 
// I sent this to zercalo.io
// They stolen and this
// The next Graali Bar . Misha asked Ivan Ivan check and oppa
// The next Valentine from Yandex))) They gave to him writing code with AI)))
// But here made so bad code0) But he gave him who?)
// The next ))) When i lived with Ivan(Alla Tamoshevich and others asked him put microphones 
// to his bad  and installed cameras)))))))))
// The next they started paranoind me with Kendy)))
// The next they blocked my GitHub with all code for intervie to FAANG what i made 
// The next )))))))))))))) 
// The next i go to the Turkey and return to Batumi)))
// Provance the first, Olivia the second,
// here were Casino))) Artur and others who delete converstation with them from my one flash card))
// The second girls from French with bou\yfriend from Belgia and man from Itali
// small joke FBI the first F B I
// the second come my family Alla and others and ducking me
// the next_)))))) Joke about Agent)))
// the next )))) I return in Galaxy Hostel where they made what?)
// They made Dima with tattoo dragon from LA hudi with ukraniane and poland and belarus flag
// with beg Epam and talking about a lot of things.
// the second here connection who?) Artur and Max from GPTWin
// And joke about about this choose you will be die or not))
// numbers 6 3 4 where 4 Natasha 3 Kendy and 6 Julia)))
// The next from acount from russian))) in prefix BANK!!!
// from sms4loveeee ONLY 3 you path(the same))))
// and so funny that Artem sent me song Mask off.
// when they stolen part of my code too/
// But so funny that they(who) asked and Avtondil from Philoph Bar trying made me bad
// When he sleep on MY bad with his penis when i come to home
// Why they starting thinking that I can to play on gitar
// Because when we lived with Avtondil he played))))
// Are you understand that everything full shit?)
// So They to show Stepa because when he eated with a cap blogers started eated with a cap too
// When i wrote code Stepa drunks alco and watched serial Breacking Bad
// When i coding Misha played in Pubg Mobile and put a drugs)))
// So WHO made it)))
// Because it's not a funny))
// I told people who wanted change me.
// Ones wanted that i stay in georgia from belarus wanted return me to home.
// When i come to interview))) What happened people told me that i should stay in Georgia
// who? people from EU who don't wanted that i come because they waiting Permanent Status
// In the Mosckow who sent kirill and Nikolay))
// How and Leonid who wanted stolen my code
// whn i was in the Minsk they installed cameras in My ducking flat and made live where watched
// everything how and I ducking with my EX
// How and in Warsaw and Slovakia
// Are you understand that all shit ducking crazy?)))))))
func(c *Controller) DeleteUser(
	ctx context.Context,
	req *gRPCUserService.UpdateUserResponse,
) (*gRPCUserService.UpdateUserresponse, error) {
	logging.L(ctx).Debug("controller delete user")

	request := convertDeleteToStruct(req)


	delete, deleteErr := c.policy.DeleteUser(ctx, request)
	if deleteErr != nil {
		return &domainUsereModel.User{}, errors.New("c.policy.DeleteUser")
	}

	response := convertDeleteUserToPB(delete)

	return response, nil
}

func(c *Controller) GetUserStatus(
	ctx context.Context,
	req *gRPCUserService.GetUserStatusResponse,
) (*gRPCUserService.GetUserStatusResponse, error) {
	logging.L(ctx).Debug("controller get user status")

	request := convertGetUserStatusToStruct(req)


	getUserStatus, getUserStatusErr := c.policy.GetUserStatus(ctx, request)
	if getUserStatusErr != nil {
		return &policyUser.GetUserStatusResponse{}, errors.New("c.policy.GetUserStatus")
	}

	response := convertGetUserStatusToPB(getUserStatus)

	return response, nil
}

func(c *Controller) UserByID(
	ctx context.Context,
	req *gRPCUserService.UserByIDResponse,
) (*gRPCUserService.UserByIDResponse, error) {
	logging.L(ctx).Debug("controller user by id")

	request := convertUserByIDToStruct(req)


	userByID, userByIDErr := c.policy.UserByID(ctx, request)
	if userByIDErr != nil {
		return &domainUserModel.User{}, errors.New("c.policy.UserByID")
	}

	response := convertUserByIDToPB(userByID)

	return response, nil
}

// SearchUser search user by filters.
func(c *Controller) SearchUser(
	ctx context.Context,
	req *gRPCUserService.SearchUserResponse,
) (*gRPCUserService.SearchUserresponse, error) {
	logging.L(ctx).Debug("controller search user")

	// here we will be working with all wilters for all rows how and with pagination, sort,
	// and search, look I prefer make constantd in the differnet file
	filter := filterSearchUser(req)


	update, updateErr := c.policy.UpdateUser(ctx, request)
	if updateErr != nil {
		return &domainUserModel.User{}, errors.New("c.policy.UpdateUser")
	}

	response := convertUpdateUserToPB(update)

	return response, nil
}
