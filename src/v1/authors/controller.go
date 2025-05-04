package authors

import (
	"go-simple-rest/src/v1/authors/dao"
	"go-simple-rest/src/v1/authors/model"
	"go-simple-rest/src/v1/authors/repo"
	"go-simple-rest/src/v1/authors/service"
	"go-simple-rest/src/v1/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var client, ctx, err = db.Connect()

// var userCollection = client.Database("go").Collection("users")

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
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusBadRequest, Success: false, Message: "Invalid ID", Data: nil})
		return
	}
	authorDAO, _ := dao.New(database)
	repository := repo.NewRepository(authorDAO)

	service, _ := service.New(repository)

	res, err := service.ShowAuthor(authorId)
	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusInternalServerError, Success: false, Message: err.Error(), Data: nil})
		return
	}
	utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusOK, Success: true, Message: "Author retrieved successfully", Data: res})
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
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusBadRequest, Success: false, Message: "Invalid ID", Data: nil})

		return
	}

	var author model.Author
	if err := c.BindJSON(&author); err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusUnprocessableEntity, Success: false, Message: "Unable to process entities", Data: nil})
		return
	}

	authorDAO, _ := dao.New(database)
	repository := repo.NewRepository(authorDAO)

	service, _ := service.New(repository)

	res, err := service.UpdateAuthor(authorId, author)

	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusInternalServerError, Success: false, Message: err.Error(), Data: nil})

		return
	}

	utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusOK, Success: true, Message: "Author updated successfully", Data: res})
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
	authorDAO, _ := dao.New(database)
	repository := repo.NewRepository(authorDAO)

	service, _ := service.New(repository)

	res, err := service.DeleteAuthor(authorId)

	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusInternalServerError, Success: false, Message: err.Error(), Data: nil})
		return
	}

	utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusOK, Success: true, Message: "Author deleted successfully", Data: res})
}
