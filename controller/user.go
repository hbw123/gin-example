package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserIns() UserHandler {
	return UserHandler{}
}

// @Title CreateUser
// @Description create users
// @Param	name	query 	string	true	"name"
// @Success 200 {string} string
// @Failure 403 query is empty
// @router /user [GET]
func (u *UserHandler) Get(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, name)
}

func (u *UserHandler) Post(c *gin.Context) {
	name := c.PostForm("name")
	c.JSON(http.StatusOK, name)
}

func (u *UserHandler) UpLoad(c *gin.Context) {
	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.Error(err)
	}
	c.String(http.StatusOK, file.Filename)
}

func (u *UserHandler) GetV2(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, fmt.Sprintf("%sV2", name))
}
