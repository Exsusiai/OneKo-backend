# Readme.md

## OneKo-backend v0.1.0
### About OneKo
[Oneko](https://oneko.app) is a platform that provides rental information integration for Chinese living in Berlin. We are currently developing a web version and a WeChat Mini Program version of the service.

### About OneKo-backend
OneKo-backend is the backend module of [OneKo project](https://oneko.app). It uses Go and MongoDB to build a backend server. The server provides RESTful API for frontend and other modules. 

And it is worth mentioning that the author of this project has no previous experience in backend development, and this project was developed with the help of GPT-4.

The purpose of this project is not only to complete the corresponding back-end engineering, but also to explore how efficiently and how good a developer can complete code engineering with the help of the current LLM(large language model) when they have no development experience in a certain field.

## New features

...

## API Endpoints

- ‘GET /v1/listings‘: Get a listing of all apartments.
- ‘GET /v1/listings/:id’: Get specific information about an apartment by ID.
- ‘POST /v1/listings‘: Create a new apartment. A JSON object matching the Apartment structure needs to be provided.
- ‘DELETE /v1/listings/:id‘: Delete an apartment. The ID of the apartment to be deleted is required.


## Data structure

We use two main data structures: Apartment and Address.

Apartment:
```
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

```

Address:
```
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

```
## Versioning and changelog

| Release Version | Date | Publisher | Berief changelog |
| --- | --- | --- | --- |
| 0.1.0 | 23.05.2023 | Jingsheng Chen | first push & add readme|

## Authors

- [OneKo app](https://oneko.app/)

- [Jingsheng Chen](https://jingsheng.dev/)

## Acknowledgments

Thanks to the authors of all external libraries used in this project and the contributors to these libraries.