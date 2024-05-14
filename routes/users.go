package routes

import (
	"crud_basic/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func updateUser(context *gin.Context) {
	userName := context.Param("name")

	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the id"})
	// 	return
	// }

	var userUpdated models.User
	err := context.ShouldBindJSON(&userUpdated)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Something went Wrong with update"})
		return
	}

	err = userUpdated.Update(userName)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "succesfully"})

}

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cannot get the users"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"User": users, "message": "succesfully"})
}

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Errors"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong went create db"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "succesfuly"})

}
