package controllers

import (
	"net/http"

	"github.com/LainForge/Neura-Launch-Dashboard/helpers"
	"github.com/LainForge/Neura-Launch-Dashboard/initializers"
	"github.com/LainForge/Neura-Launch-Dashboard/models"
	"github.com/gin-gonic/gin"
)

func CreateNewProject(ctx *gin.Context) {

	user, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	// Get the project name
	var body struct {
		Name string
	}

	err := ctx.Bind(&body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// if the project name is taken
	var project_exists models.Project

	result := initializers.DB.Where("name = ?", body.Name).First(&project_exists)

	if result.Error == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Project name already taken"})
		return
	}

	// generate a new token
	token, err := helpers.GenerateProjectToken(32)

	// Create the project
	project := models.Project{Name: body.Name, Token: token, UserID: user.(models.User).ID}

	result = initializers.DB.Create(&project)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	// Respond
	ctx.JSON(http.StatusOK, gin.H{"message": "Project Created Successfully", "token": token})
}

// Get all projects for a user
func GetProjects(ctx *gin.Context) {

	user, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	var projects []models.Project

	result := initializers.DB.Where("user_id = ?", user.(models.User).ID).Find(&projects)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"projects": projects})
}

// Get a project by token
func GetProject(ctx *gin.Context) {

	user, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	var project models.Project

	result := initializers.DB.Where("token = ? AND user_id = ?", ctx.Param("token"), user.(models.User).ID).First(&project)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch project"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"project": project})
}

// Delete a project by token
func DeleteProject(ctx *gin.Context) {

	user, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	var project models.Project

	result := initializers.DB.Where("token = ? AND user_id = ?", ctx.Param("token"), user.(models.User).ID).Delete(&project)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}
