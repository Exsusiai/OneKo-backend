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

	router := gin.Default()
	router.GET("/v1/listings", getApartments)
	router.GET("/v1/listings/:id", getApartment)
	router.POST("/v1/listings", createApartment)
	router.DELETE("/v1/listings/:id", deleteApartment)
	router.Run(":8000")
}

func getApartments(c *gin.Context) {
	var apartments []Apartment
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
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

func generateID(ctx context.Context) (string, error) {
	cursor, err := collection.Find(ctx, bson.D{}, options.Find().SetSort(bson.D{{"id", -1}}).SetLimit(1))
	if err != nil {
		return "", err
	}
	var lastApartment Apartment
	if cursor.Next(ctx) {
		cursor.Decode(&lastApartment)
		lastID, _ := strconv.Atoi(lastApartment.ID)
		return strconv.Itoa(lastID + 1), nil
	} else {
		// If there are no documents in the collection, start from 1
		return "1", nil
	}
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
























m, 
