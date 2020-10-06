package users

import (
	"net/http"
	"strconv"

	"github.com/angadthandi/bookstore_users-api/domain/users"
	"github.com/angadthandi/bookstore_users-api/services"
	"github.com/angadthandi/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userID, uErr := strconv.ParseInt(userIDParam, 10, 64)
	if uErr != nil {
		return 0, errors.NewBadRequestError("invalid user id")
	}

	return userID, nil
}

func Create(c *gin.Context) {
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

func Get(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
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

func Update(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		// bad request error...
		restErr := errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	ret, restErr := services.UpdateUser(isPartial, user)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, ret)
}

func Delete(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	restErr := services.DeleteUser(userID)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users)
}
