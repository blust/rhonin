package router

import (
	"bw/app/controllers"

	eagle "github.com/bridge-wall/bw-eagle"
)

func Register()  {
	images()
}

func images()  {

	eagle.HttpEngine.Any("/api/image/list", controllers.List)
	//ctrl := controllers.NewImageCtrl()
	//eagle.HttpEngine.POST("/api/image/list", ctrl.List)
	//g := eagle.HttpEngine.Group("/api/image")
	//{
	//
	//}
}