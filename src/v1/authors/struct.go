package authors

type Author struct {
	ID        string        `bson:"_id,omitempty" json:"id,omitempty"`
	USERID    string        `bson:"userId,omitempty" json:"userId,omitempty"`
	ARTICLES  []interface{} `bson:"articles,omitempty" json:"articles,omitempty"`
	CREATEDAT string        `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}
