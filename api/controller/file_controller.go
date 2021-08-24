package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nnnk7777/mysql-file-server/service"
)

func (ctrl Controller) HandlePostFile(c *gin.Context)  {
	var s service.FileService
	db := c.MustGet("db").(*gorm.DB)
	p, err := s.PostFile(db, c)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}