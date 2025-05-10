package main

import (
	"v2/initializers"
	models "v2/model"
)

func init() {

	initializers.LoadEnvVar()
	initializers.ConnectToDB()

}


func main(){
	initializers.DB.AutoMigrate(&models.Users{})
}