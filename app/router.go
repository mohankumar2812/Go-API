package app

import (
	"example/controller"
)

func Router() {
	r.POST("/createUser", controller.CreateUser)
	r.GET("/getUser", controller.GetUser)
	r.PATCH("/updateUser/:mail/*name",controller.UpdateUser)
	r.DELETE("/deleteUser/:mail", controller.DeleteUser)
	r.POST("/login",controller.Login)
}
