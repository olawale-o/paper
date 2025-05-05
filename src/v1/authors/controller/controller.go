package controller

import (
	"github.com/gin-gonic/gin"
)

type AuthorController interface {
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
	Show(c *gin.Context)

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
	Update(c *gin.Context)

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
	Delete(c *gin.Context)
}
