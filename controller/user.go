package controller

import (
	"context"
	"github.com/andiksetyawan/rest-api-boilerplate-go/db"
	"github.com/andiksetyawan/rest-api-boilerplate-go/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

func GetUser(c *gin.Context) {
	guid, _ := c.Get("guid")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user model.User
	opt := options.FindOne().SetProjection(bson.M{"password": 0})
	err := db.Mongo.Collection("user").FindOne(ctx, bson.M{"guid": guid}, opt).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": true, "message": "", "data": user})
}
