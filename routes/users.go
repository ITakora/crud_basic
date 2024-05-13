package routes

import (
	"crud_basic/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Errors"})
		return
	}

	err = user.Save()

	if err != nil {

		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong went create db"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "succesfuly"})

}
