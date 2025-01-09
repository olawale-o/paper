package auth

type LoginAuth struct {
	USERNAME string `bson:"username" json:"username" validate:"required,min=1"`
	PASSWORD string `bson:"username" json:"password" validate:"required,min=4"`
}

type RegisterAuth struct {
	USERNAME  string `bson:"username" json:"username" validate:"required"`
	PASSWORD  string `bson:"username" json:"password" validate:"required"`
	FIRSTNAME string `bson:"firstname" json:"firstname" validate:"required"`
	LASTNAME  string `bson:"lastname" json:"lastname" validate:"required"`
}
