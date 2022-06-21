go 1.16

module github.com/ameyasalagre/go-rest

require controllers v1.0.0 // indirect

replace controllers => ./controllers

require modules v1.0.0 // indirect

replace modules => ./modules

require routes v1.0.0 // indirect

replace routes => ./routes

require server v1.0.0

replace server => ./server

require utils v1.0.0 // indirect

replace utils => ./utils

require (
	github.com/aws/aws-sdk-go v1.34.28 // indirect
	github.com/gobuffalo/genny v0.1.1 // indirect
	github.com/gobuffalo/gogen v0.1.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/karrick/godirwalk v1.10.3 // indirect
	github.com/pelletier/go-toml v1.7.0 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	go.mongodb.org/mongo-driver v1.9.1 // indirect
	go.uber.org/zap v1.21.0 // indirect
)
