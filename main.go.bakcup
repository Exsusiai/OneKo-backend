package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	UserId   string `bson:"user_id"`
	UserName string `bson:"user_name"`
	Password string `bson:"password"`
}

var collection *mongo.Collection

func init() {
	// connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	collection = client.Database("test").Collection("users")
}

func showUsers(c *gin.Context) {
	var users []User
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting users"})
			return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
			var user User
			cursor.Decode(&user)
			users = append(users, user)
	}

	var html string
	for _, user := range users {
			html += "<p>User ID: " + user.UserId + ", User Name: " + user.UserName + ", Password: " + user.Password + "</p>"
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}

func search(c *gin.Context) {
	users := []User{}
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var user User
		err := cur.Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
			return
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, users)
}

func login(c *gin.Context) {
	var loginUser  User
	if err := c.ShouldBindJSON(&loginUser ); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// change password to md5
	hasher := md5.New()
	hasher.Write([]byte(loginUser .Password))
	loginUser.Password = hex.EncodeToString(hasher.Sum(nil))

	var foundUser User
	err := collection.FindOne(context.Background(), bson.M{"user_name": loginUser.UserName}).Decode(&foundUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid username or password"})
		return
	}

	isValid := (foundUser.Password == loginUser.Password)
	if isValid {
			fmt.Println("login succeed!")
		} else {
			fmt.Println("login failed! Wrong password")
		}
	c.JSON(http.StatusOK, gin.H{"login": isValid})
}

func signup(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser User
	err := collection.FindOne(context.Background(), bson.M{"user_name": newUser.UserName}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking username"})
			return
		}
	} else {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	// change password to md5
	hasher := md5.New()
	hasher.Write([]byte(newUser.Password))
	newUser.Password = hex.EncodeToString(hasher.Sum(nil))

	count, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user count"})
		return
	}

	newUser.UserId = strconv.Itoa(int(count) + 1)
	_, err = collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func main() {
	router := gin.Default()

	router.GET("/users", showUsers)

	router.GET("/search", search)

	router.POST("/login", login)

	router.POST("/signup", signup)

	router.Run(":8080")
}
