package helper

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func AdminValidate(c *gin.Context, id string) error {
	userId := c.GetString("uId")
	log.Println(userId)
	roleUser := c.GetString("role")
	log.Println(roleUser)

	if id == "" {
		if roleUser != "ADMIN" {
			err := errors.New("unathorizedto access this resource check")
			return err
		}
	} else {
		if roleUser != "ADMIN" || userId != id {
			err := errors.New("unathorizedto access this resource check")
			return err
		}
	}
	return nil
}
