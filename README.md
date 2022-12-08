# Credit Card Management System

## Techs: Golang (Gin), Postgres SQL

### Summary

This is a backend project for credit card management system.

To access the system you will need JWT Authentication, to get the JWT please use the sign-in and sign-up API.

After signing in you can register,update,delete a bank account.

You can attach a credit card to an existing bank account.

You can also update and delete an existing credit card from an existing bank account.

You cannot do the credit card operation if the bank account is not exist yet.

For more detail please see the API Endpoints below or you can import the postman requests [here](https://api.postman.com/collections/23612342-5836f2d5-0cbf-4f85-9943-e5467584795c?access_key=PMAT-01GKQYFFFS1ZPTFP0PHMN59TCP).

### Note:
- Bankend Port: 8888
- Postgresql port: 5432

### if you are using docker:

Required:
1. Docker

How to run this project:
1. Change your directory to root directory ./Credit-Card-Management
2. Build and Run the images and containers by entering the following command: "docker-compose up"


### if you are not using docker:

Required:
1. Go version 1.19
2. Postgresql version 14

How to run this project:
1. Change your directory to root directory ./Credit-Card-Management
2. Run the following command: "go run main.go"

### API Endpoints

1. Sign Up:
    - Endpoint:"/signup"
    - Method:POST
    - Body:    {
                    "email":"test@test.com",
                    "password":"test123"
                }

2. Sign In:  method 
    - Endpoint:"/signin"
    - Method:POST
    - Body:{
                "email"     :"test@test.com",
                "password"  :"test123"
            }

3. Get all data (Sign In Required):  method 
    - Endpoint:"/"
    - Method:GET

4. Register Bank Account: "/" method 
    - Endpoint:"/"
    - Method:POST
    - Body:{
                "name"  :"john",
                "job"   :"developer"
            }

5. Update Bank Account:
    - Endpoint:"/:id" 
    - Method:PUT
    - Body:{
                "name"  :"john",
                "job"   :"analyst"
            }

6. Delete Bank Account:
    - Endpoint:"/:id"
    - Method:DELETE

7. Register Credit Card:
    - Endpoint:"/creditcards"
    - Method:POST
    - Body:{
                "users_id":1,
                "bank":"BCA",
                "limit":10000000
            }

8. Update Credit Card:
    - Endpoint:"/creditcards"
    - Method:PUT
    - Body:{
                "users_id":1,
                "bank":"Mandiri",
                "limit":20000000,
                "ammount":5000000
            }

9. Delete Credit Card:
    - Endpoint:"/creditcards/:id"
    - Method:DELETE