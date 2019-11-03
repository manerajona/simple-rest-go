# simple-rest-go
Simple CRUD for a RESTful go application.

docker pull mongo

docker run -d -p 27017-27019:27017-27019 --name mongodb mongo

go get -u github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/mongo

go build && ./rest

