package service

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type FileService struct{}

func (fs FileService) PostFile(db *gorm.DB, c *gin.Context) (File, error) {
	var f File
	file, err := c.FormFile("file")
	if err != nil {
		return f, err
	}

	name := c.PostForm("name")
	filename := "./data/" + filepath.Base(file.Filename)

	f = File{
		Name: name,
		Path: filename,
	}

	if err := db.Create(&f).Error; err != nil {
		return f, err
	}

	fmt.Println(name)
	fmt.Println(filename)
	fmt.Printf("%#v\n", f)

	if err := c.SaveUploadedFile(file, filename); err != nil {
		return f, err
	}

	return f, nil
}