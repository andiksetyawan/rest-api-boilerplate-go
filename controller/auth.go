package controller

import (
	"context"
	"github.com/andiksetyawan/rest-api-boilerplate-go/db"
	"github.com/andiksetyawan/rest-api-boilerplate-go/model"
	"github.com/andiksetyawan/rest-api-boilerplate-go/storage"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

type JWTClaims struct {
	jwt.StandardClaims
	Guid,
	Email,
	Group string
}

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var login login
	err := c.BindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	var user model.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = db.Mongo.Collection("user").FindOne(ctx, bson.M{"email": login.Email}).Decode(&user)
	if err != nil {
		log.Printf("email %v not found", login.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "invalid username/password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		log.Printf("email %v, password is wrong", login.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "invalid username/password"})
		return
	}

	signedToken, err := generateToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "token": signedToken})
}

func SignUp(c *gin.Context) {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	file, _, err := c.Request.FormFile("picture")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer file.Close()

	fileName := uuid.New().String()
	_, err = storage.Upload("user", fileName, file)
	if err != nil {
		log.Println(err)
	}

	user.GUID = uuid.New().String()
	user.Picture = fileName
	user.CreatedAt = time.Now().Unix()
	user.IsSuspend = false
	user.Group = "user"

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}
	user.Password = string(hash)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := db.Mongo.Collection("user").InsertOne(ctx, user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	token, _ := generateToken(&user)
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "successfully signup", "token": token})
}

func generateToken(user *model.User) (string, error) {
	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    os.Getenv("APPLICATION_NAME"),
			ExpiresAt: time.Now().Add(time.Duration(24) * time.Hour).Unix(),
		},
		Guid:  user.GUID,
		Email: user.Email,
		Group: user.Group,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SIGNATURE_KEY")))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
