package authors

import (
	"go-simple-rest/src/v1/authors/model"
	"go-simple-rest/src/v1/authors/repo"
	"go-simple-rest/src/v1/authors/service"
	"go-simple-rest/src/v1/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
	authorId, err := utils.ParseParamToPrimitiveObjectId(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	repository, err := repo.New(database)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	service, err := service.New(repository)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	res, err := service.ShowAuthor(authorId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Retrieved single author", "data": res})
}

// AuthorUpdate godoc
// @Tags Authors
// @Summary Update a specific author
// @Description Updates an author
// @Param id path string true "Author ID"
// @Produce json
// Accept application/json
// @Success 200 {string} message "Response"
// @Success 422 {string} message "Response"
// @Failure 500 {object} string "Error"
// @Router /authors/{id} [put]
func Update(c *gin.Context) {
	authorId, err := utils.ParseParamToPrimitiveObjectId(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	var author model.Author
	if err := c.BindJSON(&author); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Unable to process entities"})
		return
	}

	repository, err := repo.New(database)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	service, err := service.New(repository)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	res, err := service.UpdateAuthor(authorId, author)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Author updated successfully", "data": res})
}

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
	authorId, _ := utils.ParseParamToPrimitiveObjectId(c.Param("id"))

	repository, err := repo.New(database)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	service, err := service.New(repository)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	res, err := service.DeleteAuthor(authorId)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Author deleted successfully", "data": res})
}
