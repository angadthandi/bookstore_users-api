package users

import (
	"net/http"
	"strconv"

	"github.com/angadthandi/bookstore_users-api/domain/users"
	"github.com/angadthandi/bookstore_users-api/services"
	"github.com/angadthandi/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	// This commented code is replaced by gin code below...
	//
	// b, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// TODO handle err
	// 	return
	// }

	// err = json.Unmarshal(b, &user)
	// if err != nil {
	// 	// TODO handle json err
	// 	return
	// }

	// Gin code replaces above code...
	err := c.ShouldBindJSON(&user)
	if err != nil {
		// bad request error...
		restErr := errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}

	ret, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// user create err
		c.JSON(saveErr.Status, saveErr)
		return
	}

	// c.String(http.StatusNotImplemented, "TODO")
	c.JSON(http.StatusCreated, ret)
}

func GetUser(c *gin.Context) {
	userID, uErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if uErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
	}

	ret, getErr := services.GetUser(userID)
	if getErr != nil {
		// user create err
		c.JSON(getErr.Status, getErr)
		return
	}

	// c.String(http.StatusNotImplemented, "TODO")
	c.JSON(http.StatusOK, ret)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "TODO")
}
