# DIALOG's BookStore-API

v1.0

## Description
This API is a test for Dialog's Backend Engineer role. It consists in a basic Golang CRUD API
interacting with a 2 table POSTGRES DB called **bookstore**:
- Authors
- Books


### Recommended tools:

 - DBeaver Community(22.3.0)
 - IDE JetBrains GoLand(2023.1.3)
 - Docker Desktop(v20.10.21)
 - Postman(10.15.0)
 - Git for Windows(2.38.1)

****
## Local Setup (for Windows 10)

### **Install Go >= 1.18 (https://go.dev/doc/effective_go)**

In your terminal go to the project root folder and run the commands to install the required libraries:

**Viper Lib**
>> **go get github.com/spf13/viper**

**Chi Lib** 
>> **go get github.com/go-chi/chi/v5**

### **Docker Compose**
To build the Database defined in the docker compose file run the line below in the terminal:
>> **docker compose up**   *or* **docker compose up -d**

To stop the containers defined in docker compose file:

> **docker compose down**
>
To stop and remove containers, networks, and volumes defined in a Docker Compose file (for troubleshooting):
> 
> **docker compose down --remove-orphans**

### **Populating the Authors Table**
To transfer the authors.csv file data to the database run in the **terminal** (in **/DialogBookStoreAPI folder**):
>> **go run populate/populate_authors.go**

PS.: In this version it may take 60 to 100 minutes. You can leave it running while you do other activities.

### Running the API Locally
In the **main.go** file folder run:
>> **go run main.go**

****
## Testing

To test, go to the folder *collection* and import in ***Postman*** or ***Insomnia*** to build the fields as below:
### To test the GET ALL Authors Route

> **[POST] http://localhost:9000/Author/**

### To test the GET Authors by Id Route

> **[POST] http://localhost:9000/Author/{id}**

### To test the POST (Insert) Book Route

> **[POST] http://localhost:9000/books/**

### To test the PUT (Update) Book Route
> **[PUT] http://localhost:9000/books/{id}**


### To test the GET ALL Books Route
> **[GET] http://localhost:9000/books/**


### To test the GET Book by Id Route
> **[GET] http://localhost:9000/books/{id}**


### To test the DELETE Book by Id Route
> **[GET] http://localhost:9000/books/{id}**

## Thank you for reading! Enjoy the CRUD! :) 