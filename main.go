package main

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func main() {
	//Create a new MongoDB client, the function returns a pointer to the mongo.Client structure and an error object
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	collection = client.Database("oneko").Collection("apartments")

	// Create a new Gin router
	router := gin.Default()
	router.GET("/v1/listings", getApartments)
	router.GET("/v1/listing/:id", getApartment)
	router.POST("/v1/listing", createApartment)
	router.DELETE("/v1/listing/:id", deleteApartment)
	//Starts the HTTP server, enters an internal listening loop, and listens for incoming HTTP requests on the given address and port
	router.Run(":8000")
}

func getApartments(c *gin.Context) {
	var apartments []Apartment
	//context.Background() is an empty context, which is generally used when you don't care about cancellation, deadlines, or incoming values. 
	//bson.D{} is an empty BSON document, which means there is no filter condition, 
	//so this method will return all documents in the collection.
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//defer the execution of a function to the function that contains it (getApartments) will be executed when it is about to return. 
	//cursor.Close(context.Background()) is an operation to close the database cursor, which can prevent resource leaks.
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var apartment Apartment
		cursor.Decode(&apartment)
		apartments = append(apartments, apartment)
	}

	c.JSON(http.StatusOK, apartments)
}

func getApartment(c *gin.Context) {
	id := c.Param("id")
	var apartment Apartment
	err := collection.FindOne(context.Background(), bson.D{{"id", id}}).Decode(&apartment)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No apartment found with given ID"})
		return
	}

	c.JSON(http.StatusOK, apartment)
}

func createApartment(c *gin.Context) {
	var apartment Apartment
// Check if BindJSON encountered an error while performing this parsing and binding operation. 
// Possible errors include:
// 1.The body of the request is not valid JSON.
// 2.The requested JSON data could not be matched to the Apartment structure.
	if err := c.BindJSON(&apartment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

		// Generate ID if not provided
		if apartment.ID == "" {
			apartment.ID = strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(10000))
		} else {
		// Check if ID already exists
		var existingApartment Apartment
			err := collection.FindOne(context.Background(), bson.D{{"id", apartment.ID}}).Decode(&existingApartment)
			if err != nil {
				if err != mongo.ErrNoDocuments {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
			} else {
				// if ID already exists, return error
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID already exists"})
				return
			}
		}

	_, err := collection.InsertOne(context.Background(), apartment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, apartment.ID)
}

func deleteApartment(c *gin.Context) {
	id := c.Param("id")
	_, err := collection.DeleteOne(context.Background(), bson.D{{"id", id}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Apartment deleted successfully"})
}