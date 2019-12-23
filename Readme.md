# Book Repository

Book Repository performs simple CRUD operations in Golang. 
Book records are stored in MYSQL database.

### Tech and Installation
* Book Repository is a go-modules project.
* API Endpoints are in bookstore-routes.go file.
* DB in MYSQL is connected to perform CRUD operations.
* [Golang] 
* [MYSQL] 
### Start the Application 
 * go run main.go
    Application will start on http port 8084   

### API Endpoints in Application 
* GET /book/{id}
* GET /book/
* POST /book/
  Input Body : {
	"name" : "harry potter",
	"author" : "j.k.rowlings",
	"publication" : "test"
}
* DELETE /book/{id}
* PUT /book/{id}
  Input Body : {
	"name" : "harry potter",
	"author" : "j.k.rowlings",
	"publication" : "test"
}