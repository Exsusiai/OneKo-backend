package main

type Apartment struct {
	ID            string   `json:"id" bson:"id"`
	Title         string   `json:"title" bson:"title"`
	Description   string   `json:"description" bson:"description"`
	Price         int      `json:"price" bson:"price"`
	Address       Address  `json:"address" bson:"address"`
	Pictures      []string `json:"pictures" bson:"pictures"`
	Rooms         int      `json:"rooms" bson:"rooms"`
	Area          int      `json:"area" bson:"area"`
	MaxTenants    int      `json:"maxTenants" bson:"maxTenants"`
	Type          string   `json:"type" bson:"type"`
	AvailableFrom string   `json:"availableFrom" bson:"availableFrom"`
	AvailableTo   string   `json:"availableTo" bson:"availableTo"`
	IsSaved       bool     `json:"isSaved" bson:"isSaved"`
	CreatedAt     string   `json:"createdAt" bson:"createdAt"`
	UpdatedAt     string   `json:"updatedAt" bson:"updatedAt"`
}

type Address struct {
	Street    string  `json:"street" bson:"street"`
	District  string  `json:"district" bson:"district"`
	ZipCode   string  `json:"zipCode" bson:"zipCode"`
	City      string  `json:"city" bson:"city"`
	State     string  `json:"state" bson:"state"`
	Country   string  `json:"country" bson:"country"`
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}
