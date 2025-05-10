package controllers

import (
	"net/http"
	"os"
	"time"
	"v2/initializers"
	models "v2/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	
	var body struct{
		Name string
		Email string
		Password string
	}

	c.Bind(&body)

	var existingEmail models.Users

	if err := initializers.DB.Where("email = ?", body.Email).First(&existingEmail).Error; err == nil {
		c.JSON(400 , gin.H{
		"error": "Email already exists",
	})
	}



 hash , err:=	bcrypt.GenerateFromPassword([]byte(body.Password) , 10)
 if err !=nil{
	c.JSON(500 , gin.H{
		"error": "Cannot Generate Password",
	})
 }


	user := models.Users{Name: body.Name,Email: body.Email , Password:string(hash) }

	result := initializers.DB.Create(&user)

	if result.Error != nil{
		c.Status(400)
		return
	}


	c.JSON(200, gin.H{
		"user": user,
	})
}


func Login(c *gin.Context) {
	
	var body struct{
		Email string
		Password string
	}
	c.Bind(&body)

	var user models.Users

	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest , gin.H{
			"error" : "Invalid Email or password",
		})
	}

	

	err:= bcrypt.CompareHashAndPassword([]byte(user.Password) , []byte(body.Password))

	if err != nil{
		c.JSON(400 , gin.H{
		"error": "Incorrect Password",
	})
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"sub": user.ID,
	"exp": time.Now().Add(time.Hour * 24 ).Unix(),
})

// Sign and get the complete encoded token as a string using the secret
tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

if err != nil{
	c.JSON(http.StatusBadRequest , gin.H{
		"error" : "Failed To generate token",
	})
	return
}


c.JSON(200, gin.H{
		"token": tokenString,
		"userEmail" : user.Email,
	})

}