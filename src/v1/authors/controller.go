package authors

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var client, ctx, err = db.Connect()

// var userCollection = client.Database("go").Collection("users")

// func Index(c *gin.Context) {
// 	res, err := ShowAuthors()
// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, gin.H{"message": "Retrieved all authors", "data": res})
// }
//

// Author godoc
// @Tags Authors
// @Summary Get author by id
// @Description Retrieves a specific author by ID.
// @Param id path string true "Author ID"
// @Produce json
// @Success 200 {object} model.Author "Response"
// @Failure 400 {object} string "Error"
// @Failure 500 {object} string "Error"
// @Router /authors/{id} [get]
func Show(c *gin.Context) {
	authorId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	res, err := ShowAuthor(authorId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Retrieved single author", "data": res})
}

// func Update(c *gin.Context) {
// 	authorId, _ := primitive.ObjectIDFromHex(c.Param("id"))

// 	var author Author
// 	if err := c.BindJSON(&author); err != nil {
// 		log.Println(err)
// 		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Unable to process entities"})
// 		return
// 	}

// 	res, err := UpdateAuthor(authorId, author)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, gin.H{"message": "Author updated successfully", "data": res})
// }

// AuthorDelete godoc
// @Tags Authors
// @Summary Delete a specific author
// @Description Deletes an author
// @Param id path string true "Author ID"
// @Produce json
// Accept application/json
// @Success 200 {string} message "Response"
// @Failure 500 {object} string "Error"
// @Router /authors/{id} [delete]
func Delete(c *gin.Context) {
	authorId, _ := primitive.ObjectIDFromHex(c.Param("id"))

	res, err := DeleteAuthor(authorId)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Author deleted successfully", "data": res})
}
