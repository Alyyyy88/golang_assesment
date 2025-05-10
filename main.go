package main

import (
	"v2/controllers"
	"v2/initializers"

	"github.com/gin-gonic/gin"
)


func init(){
	initializers.LoadEnvVar()
	initializers.ConnectToDB()
}


func main(){
	  router := gin.Default()
  router.POST("/users/register",controllers.Register)
  router.POST("/users/login",controllers.Login)
  router.Run()
}