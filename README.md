# RestAPI-Pinjaman-Online
[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

## Dependency in this GO program
```sh
gorm.io/gorm
gorm.io/driver/postgres
github.com/gin-gonic/gin
github.com/joho/godotenv
github.com/dgrijalva/jwt-go
github.com/golang/crypto
github.com/mashingan/smapping
```

## Installation

* git clone https://github.com/kelvinnich/RestAPI-Pinjaman-Online \n
* cd RestAPI-Pinjaman-Online
* Edit .env
* Set your database connection details
* go run main.go
 
# API Endpoints

| Method | Endpoints |    Description     |
| ------ | ------ | ------- |
| POST | pinjol/auth/register | Register account must given name,email,no ktp,password to body request
| POST | pinjol/auth/login | login account must given email & password to body request
| GET | pinjol/nasabah/profile | Get Customer
| PUT |  pinjol/nasabah/update | Edit Customer
| POST | pinjol/document/uploadDocs | Add document customers
| GET | pinjol/document/:id | Get document customers by id
| PUT | pinjol/document/:id | Update document by id
| DELETE | pinjol/document/:id | Delete document by id
| POST | pinjol/pekerjaan/addJobs | Add jobs customers
| GET | pinjol/pekerjaan/:id | Get jobs customers by id
| PUT | pinjol/pekerjaan/:id | Update jobs customers by id
| DELETE | pinjol/pekerjaan/:id | Delete jobs customers by id
| POST | pinjol/pinjaman/uang | the customer applies for a loan
| GET |  pinjol/pinjaman/customerID | verify loans by customer id
| PUT | pinjol/pinjaman/update | customer loan updates by id
| DELETE | pinjol/pinjaman/update | delete lending data by id
| GET | pinjol/pembayaran/:status | retrieve a list of payment data based on status
| GET | pinjol/pembayaran/:id | View monthly loan payments by id
| GET | pinjol/pembayaran/:id | see the total loan payments by id
| POST | pinjol/pembayaran/perbulan | make monthly payments
| PUT  |  pinjol/pembayaran/:id | update payment by id 
| DELETE | pinjol/pembayaran/:id  | Delete payment loan by id
| GET | history/pembayaran/:id | view payment history by id
| GET | history/pembayaran | view all payment history
