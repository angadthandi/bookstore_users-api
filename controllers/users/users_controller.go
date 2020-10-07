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

	userService := services.NewUserService()
	ret, saveErr := userService.CreateUser(user)
	if saveErr != nil {
		// user create err
		c.JSON(saveErr.Status, saveErr)
		return
	}

	out, err := ret.Marshal(c.GetHeader("X-Public") == "true")
	if err != nil {
		// marshal err
		restErr := errors.NewInternalServerError("user marshal error")
		c.JSON(restErr.Status, restErr)
		return
	}

	// c.String(http.StatusNotImplemented, "TODO")
	c.JSON(http.StatusCreated, out)
}

func Get(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	userService := services.NewUserService()
	ret, getErr := userService.GetUser(userID)
	if getErr != nil {
		// user create err
		c.JSON(getErr.Status, getErr)
		return
	}

	out, err := ret.Marshal(c.GetHeader("X-Public") == "true")
	if err != nil {
		// marshal err
		restErr := errors.NewInternalServerError("user marshal error")
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, out)
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

	userService := services.NewUserService()
	ret, restErr := userService.UpdateUser(isPartial, user)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	out, err := ret.Marshal(c.GetHeader("X-Public") == "true")
	if err != nil {
		// marshal err
		restErr = errors.NewInternalServerError("user marshal error")
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, out)
}

func Delete(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	userService := services.NewUserService()
	restErr := userService.DeleteUser(userID)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	userService := services.NewUserService()
	users, searchErr := userService.Search(status)
	if searchErr != nil {
		c.JSON(searchErr.Status, searchErr)
		return
	}

	out, err := users.Marshal(c.GetHeader("X-Public") == "true")
	if err != nil {
		// marshal err
		restErr := errors.NewInternalServerError("user marshal error")
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, out)
}
