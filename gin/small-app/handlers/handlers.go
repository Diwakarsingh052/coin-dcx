package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"small-app/data/user"
	"strconv"
)

// /user?user_id=2

// GetUser is entry point for /user endpoint
// think how would you handle the request when someone hit this endpoint

func GetUserGin(c *gin.Context) {
	userIdString := c.Query("user_id")
	//converting it to make sure it is a valid uint64
	userId, err := strconv.ParseUint(userIdString, 10, 64)
	if err != nil {
		log.Println(err)
		appErr := map[string]string{"Message": http.StatusText(http.StatusBadRequest)}

		// setting error status code and responding in json
		c.AbortWithStatusJSON(http.StatusBadRequest, appErr)

		return

	}

	//fetching the user with the userId provided
	u, err := user.FetchUser(userId)
	if err != nil {

		log.Println(err)
		//aborting with json
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "user id not found",
		})
		return

	}
	c.JSON(http.StatusOK, u)
}
