package articles

type Article struct {
	ID     string `bson:"_id,omitempty" json:"id,omitempty"`
	TITLE  string `bson:"title" json:"title"`
	AUTHOR string `bson:"author" json:"author"`
}
