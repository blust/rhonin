package controllers

import (
	eagle "github.com/bridge-wall/bw-eagle"
	"github.com/gin-gonic/gin"
	"sync"
)

type ImageCtrl struct {
}

var (
	imageCtrl     *ImageCtrl
	imageCtrlOnce sync.Once
)

func NewImageCtrl() *ImageCtrl {
	imageCtrlOnce.Do(func() {
		imageCtrl = &ImageCtrl{}
	})
	return imageCtrl
}

func List(c *gin.Context) {//
	//fmt.Println("hello world, my name is bw")
	eagle.RenderJson(c, eagle.NewCodeResponse(eagle.Success, nil))
}
