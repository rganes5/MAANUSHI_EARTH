# MAANUSHI_EARTH_E-COMMERCE_GO_GIN_CLEAN_ARCH
Maanushi for Earth is an initiative to support rural women from Kumily to generate livelihood through sustainable and eco-friendly businesses. As part of reaching out to more women through our initiative, this platform offers a space for sustainable and women-owned businesses to list their products as well.

This project is an ecommerce application built using Go programming language and the Gin framework. It follows the clean code architecture, which promotes separation of concerns and maintainability.

## Project Overview
The ecommerce-gin-clean-arch project is designed to provide a performant and feature-rich ecommerce solution. It includes functionalities such as user authentication, product management, shopping cart, order processing, and payment integration.

## Used Packages
The project utilizes the following packages:
1. [GIN](github.com/gin-gonic/gin): A web framework written in Go that combines high performance with an API similar to Martini.
2. [JWT](github.com/golang-jwt/jwt): A Go implementation of JSON Web Tokens for secure authentication and authorization.
3. [GORM](https://gorm.io/index.html) with [PostgreSQL](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL): A powerful ORM library for Go that simplifies database operations.
4. [Wire](https://github.com/google/wire): A code generation tool for dependency injection, making it easier to connect components.
5. [Viper](https://github.com/spf13/viper): A configuration solution for Go applications, supporting various formats and 12-Factor app principles.
6. [swag](https://github.com/swaggo/swag) with [gin-swagger](https://github.com/swaggo/gin-swagger) and [swaggo files](github.com/swaggo/files): Converts Go annotations to Swagger Documentation 2.0 for API documentation.
7. [Twilio](https://github.com/twilio/twilio-go): A Go client library for the Twilio API, enabling communication via SMS, voice, and other channels.
8. [Razorpay](https://github.com/razorpay/razorpay-go): A Go client library for the Razorpay API, facilitating payment processing and management.

Please refer to the respective package documentation for more information on how to use and integrate these packages into your Go application.

# Setup Instructions
To use and test the MAANUSHI_EARTH_E-COMMERCE_GO_GIN_CLEAN_ARCH project, please follow these steps:

### Prerequisites
Make sure you have the following installed on your system:
- Go (https://golang.org/doc/install)
- PostgreSQL (https://www.postgresql.org/download/)

### 1. Clone the Repository
Clone the MAANUSHI_EARTH_E-COMMERCE_GO_GIN_CLEAN_ARCH repository to your local system:
```bash
git clone https://github.com/rganes5/MAANUSHI_EARTH_E-COMMERCE.git
cd MAANUSHI_EARTH_E-COMMERCE_GO_GIN_CLEAN_ARCH
```
### 2. Install Dependencies
Install the required dependencies using either the provided Makefile command or Go's built-in module management:
```bash
# Using Makefile
make deps
# OR using Go
go mod tidy
```
### 3. Configure Environment Variables
details provided at the end of file
### 4. Make Swagger Files (For Swagger API Documentation)
```bash
make swag
```
# To Run The Application
```bash
make run
```
### To See The API Documentation
1. visit [swagger] ***http://localhost:3000/swagger/index.html***

# To Test The Application
### 1. Generate Mock Files
```bash
make generate
```
### 2. Run The Test Functions
```bash
make test
```
### 3. Mock Repository layers
```bash
go install github.com/golang/mock/mockgen@v1.6.0
```
```bash
go get github.com/golang/mock/gomock
```
```bash
mockgen -source=pkg/repository/interface/admin.go -destination=pkg/repository/mockRepo/mock_admin.go -packag
e=mockRepo
```

# Set up Environment Variables
Set up the necessary environment variables in a .env file at the project's root directory. Below are the variables required:
### PostgreSQL database details
1. DB_HOST="```your database host name```"
2. DB_NAME="```your database name```"
3. DB_USER="```your database user name```"
4. DB_PASSWORD="```your database owner password```"
5. DB_PORT="```your database running port number```"
### JWT
1. ADMIN_AUTH_KEY="```secret code for signing admin JWT token```"
2. USER_AUTH_KEY="```secret code for signing user JWT token```"
### Twilio
1. AUTH_TOKEN="your Twilio authentication token"
2. ACCOUNT_SID="```your Twilio account SID```"
3. SERVICE_SID="```your Twilio messaging service SID```"
### Razorpay
1. RAZOR_PAY_KEY="```your Razorpay API test key```"
2. RAZOR_PAY_SECRET="```your Razorpay API test secret key```"

