package owner

import "time"

type CreateOwnerRequest struct {
	ID string // UUIDv7 we crete UUID in Policy layer
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Adress string `json:"adress"`
	Country string `json:"country"`
	SocialNetwork []byte `json:"social_network"`
	FollowerCount uint64 `json:"follower_count"`
	VisabilityStatus bool `json:"visability_status"`
	EnableStatus bool `json:"enable_status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewConstructorCreateOwnerrequest(
	id, firstName, lastName, email, phoneNumber, adress, country string,
	socialNetwork []byte,
	followerCounter uint64,
	visabilityStatus, EnableStatus bool,
	createAt, updateAt time.Time,
) *CreateOwnerRequest {
	return &CreateOwnerRequest{
		ID: id,
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		PhoneNumber: phoneNumber,
		Adress: adress,
		Country: country,
		SocialNetwork: socialNetwork,
		FollowerCount: followerCount,
		VisabilityStatus: visabilityStatus,
		EnableStatus: EnableStatus,
		CreatedAt: createAt,
		UpdatedAt: updateAt,
	}
}