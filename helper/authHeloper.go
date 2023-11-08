package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func AdminValidate(c *gin.Context, id string) error {
	userType := c.GetString("role")
	userId := c.GetString("uid")

	if id == "" {
		if userType != "ADMIN" {
			err := errors.New("unathorizedto access this resource check")
			return err
		}
	} else {
		if userType != "ADMIN" || userId != id {
			err := errors.New("unathorizedto access this resource check")
			return err
		}
	}
	return nil
}
