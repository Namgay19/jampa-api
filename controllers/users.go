package controllers

import (
	"namgay/jampa/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)	

func RegisterUser(c *gin.Context) {
	 var input registerUser

	 if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	 }

	 if len(input.Password) < 5 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Password should be longer than 5 characters"})
		return
	 }

	 password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	 user := models.User{
		 FirstName: input.FirstName,
		 LastName: input.LastName,
		 Email: input.Email,
		 Phone: input.Phone,
		 Password: string(password),
	 }

	 result := models.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": user.ID})
}

func Login(c *gin.Context) {
	var input loginUser

	if err:= c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err:= bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	mySigningKey := []byte(os.Getenv("JWT_SECRET"))

	// Create the Claims
	claims := MyCustomClaims{
	int(user.ID),
	jwt.StandardClaims{
		ExpiresAt: jwt.At(time.Now().Add(time.Hour)),
		Issuer: "namgay",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)

	c.Header("Authorization", "Bearer " + ss)
	c.JSON(http.StatusOK, gin.H{"message": "successfully logged in"})
}


type MyCustomClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

type registerUser struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone string `json:"phone"`
}

type loginUser struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}