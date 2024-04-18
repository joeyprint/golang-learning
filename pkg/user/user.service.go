package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, mockUsers)
}

func getUsersById(context *gin.Context) {
	var id string = context.Param("id")

	for _, user := range mockUsers {
		if user.Id == id {
			context.IndentedJSON(http.StatusOK, user)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found user"})
}
