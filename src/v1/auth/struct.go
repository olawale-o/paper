package auth

type LoginAuth struct {
	USERNAME string `bson:"username" json:"username"`
	PASSWORD string `bson:"username" json:"password"`
}

type RegisterAuth struct {
	USERNAME  string `bson:"username" json:"username"`
	PASSWORD  string `bson:"username" json:"password"`
	FIRSTNAME string `bson:"firstname" json:"firstname"`
	LASTNAME  string `bson:"lastname" json:"lastname"`
}
