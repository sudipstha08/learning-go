package controller

import (
	"learning-go/models"
	"learning-go/services"
	"learning-go/validators"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []models.Video
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []models.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video models.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *controller) Update(ctx *gin.Context) error {
	var video models.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	} 
	video.ID = id

	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Update(video)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var video models.Video
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)

	if err != nil {
		return err
	}
	video.ID = id
	c.service.Delete(video)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	//PASSING DATA TO TEMPLATE
	ctx.HTML(http.StatusOK, "index.html", data)
}
