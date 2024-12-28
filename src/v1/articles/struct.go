package articles

type Article struct {
	ID        string `bson:"_id,omitempty" json:"id,omitempty"`
	TITLE     string `bson:"title" json:"title"`
	CONTENT   string `bson:"content" json:"content"`
	AUTHORID  string `bson:"authorId,omitempty" json:"authorId,omitempty"`
	LIKES     int    `bson:"likes,omitempty" json:"likes,omitempty"`
	VIEWS     int    `bson:"views,omitempty" json:"views,omitempty"`
	CREATEDAT string `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UPDATEDAT string `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
