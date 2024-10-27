# go-lang-gin-postgresql-example
Deploying a Golang RESTful API with Go using Gin, Gorm and PostgreSQL - Example use Gin and Postgre Sql

In this tutorial, we’ll demonstrate how to build a bookstore REST API that provides book data and performs CRUD operations.

Before we get begin, I’ll assume that you:

- Have Go installed on your machine
- Understand the basics of Go language
- Have a general understanding of RESTful API

## Installation

Let’s start by initializing a new Go module to manage our project’s dependencies. 
Make sure you run this command inside your Go environment folder:

``go run main.go``

Now let’s install [Gin](https://github.com/gin-gonic/gin) and [Gorm](https://github.com/jinzhu/gorm) packages:

``go get github.com/gin-gonic/gin gorm.io/gorm gorm.io/driver/postgres``
``go get github.com/joho/godotenv``

## Setting up the server

Let’s start by creating a Hello World server inside the main.go file:
```
package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })

  r.Run()
}
```

To test it out, we’ll start our server by running the command below:
``go run main.go``

## Setting up the database

The next thing we need to do is to build our database models.

Model is a class (or structs in Go) that allows us to communicate with a specific table in our database. In Gorm, we can create our models by defining a Go struct. This model will contain the properties that represent fields in our database table. Since we’re trying to build a bookstore API, let’s create a Book model:

```
\\models\contact.go
package models

type Contact struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email" gorm:"unique"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	ZipCode     string `json:"zip_code"`
	Country     string `json:"country"`
	Tags        string `json:"tags"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateContactInput struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email" binding:"required"`
	Address     string `json:"address"`
	City        string `json:"City"`
	State       string `json:"state"`
	ZipCode     string `json:"zip_code"`
	Country     string `json:"country"`
	Tags        string `json:"tags"`
}

type struct UpdateContactInput {
	FirstName  string `json:"fist_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number"`
	Address string `json:"address"`
	City string `json:"City"`
	State string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
	Tags string `json:"tags"`
}

```

Next, we need to create a utility function called ConnectDatabase that allows us to create a connection to the database and migrate our model’s schema. 
We can put this inside the setup.go file in our models module:

```
// models/setup.go

package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "host=localhost user=root password=simplecrm2024 dbname=crm_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Contact{})
	if err != nil {
		return
	}

	DB = database
}

```

In main.go, we need to call the following function before we run our app:
```
package main

import (
  "github.com/gin-gonic/gin"

  "github.com/landofcoder/go-lang-gin-postgresql-example/models" // new
)

func main() {
  r := gin.Default()

  models.ConnectDatabase() // new

  r.Run()
}
```

## Setting up the RESTful routes

We’re almost there!

The last thing we need to do is to implement our controllers. In the previous section, we learned how to create a route handler (i.e., controller) inside our main.go file. 
However, this approach makes our code much harder to maintain. 
Instead of doing that, we can put our controllers inside a separate module called controllers.
file: ``controllers\contact.go`` with there functions:

- The Read handler function
- The Create handler function
- The Create handler function for single data
- The Update handler function
- The Delete handler function

## Demo Data

- run endpoint: ``/contacts``
Return data:
```
{
  "data": []
}
```
- send a POST request to ``/contacts`` endpoint with this request body:

```
{
  "first_name": "Thuan",
  "last_name": "Luu",
  "phone_number": "+84 035 653 5598",
  "email": "info@landofcoder.com",
  "company": "landofcoder",
  "job_title": "CEO",
  "address": "54 My Dinh Street",
  "city": "Ha Noi",
  "state": "Hanoi",
  "zip_code": "100000",
  "country": "Vietnam",
  "tags": "An Entrepreneur,Software Engineer,Saas,Golang,Reactjs,Magento,Shopify,Woocommerce,Wordpress"
}
```

The response should look like this:
```
{
  "data": {
    "id": 1,
    "first_name": "Thuan",
    "last_name": "Luu",
    ...
  }
}
```

- Let’s run the server and fetch ``/contacts/1`` to get the book we just created:

```
{
  "data": {
    "id": 1,
    "first_name": "Thuan",
    "last_name": "Luu",
    ...
  }
}
```

- Let’s test it out! Fire a PATCH request to the /contacts/:id endpoint to update the book title:

```
{
  "company": "Bavaan"
}
```

The result should be as follows:
```
{
  "data": {
    "id": 1,
    "first_name": "Thuan",
    "last_name": "Luu",
    "company": "Bavaan",
    ...
  }
}
```

- Let’s test it out by sending a DELETE request to the /contacts/1 endpoint:

```
{
  "data": true
}
```

If we fetch all books in /contacts, we’ll see an empty array again:
```
{
  "data": []
}
```