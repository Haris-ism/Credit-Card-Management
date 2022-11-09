# Credit Card Management System

## Techs: Golang (Gin), Postgres SQL

### Summary

This is a backend project for credit card management system.

To access the system you will need JWT Authentication, to get the JWT please use the sign-in and sign-up API.

After signing in you can register,update,delete a bank account.

You can attach a credit card to an existing bank account.

You can also update and delete an existing credit card from an existing bank account.

You cannot do the credit card operation if the bank account is not exist yet.

For more detail please see the API Endpoints below.

### Note:

    - Bankend Port: 8888

    - Postgresql port: 5433


### if you are using docker:

Required:
1. Docker

How to run this project:
1. Change your directory to root directory ./Credit-Card-Management
1. Build the images and containers by entering the following command: "./build.sh"
2. To run the program, just simply run: "./run.sh"

Note: if you are using windows and cannot run the ".sh" file, then right click on the ".sh" file and choose "open with git bash"


### if you are not using docker:

Required:
1. Go version 1.19
2. Postgresql version 14

How to run this project:
1. Change your directory to root directory ./Credit-Card-Management
2. Run the following command: "go run main.go"

### API Endpoints

1. Sign Up:
    1. Endpoint:"/signup"
    2. Method:POST
    3. Body:    {
                    "email":"test@test.com",
                    "password":"test123"
                }

2. Sign In:  method 
    1. Endpoint:"/signin"
    2. Method:POST
    3. Body:{
                "email"     :"test@test.com",
                "password"  :"test123"
            }

3. Get all data (Sign In Required):  method 
    1. Endpoint:"/"
    2. Method:GET

4. Register Bank Account: "/" method 
    1. Endpoint:"/"
    2. Method:POST
    3. Body:{
                "name"  :"john",
                "job"   :"developer"
            }

5. Update Bank Account:
    1. Endpoint:"/:id" 
    2. Method:PUT
    3. Body:{
                "name"  :"john",
                "job"   :"analyst"
            }

6. Delete Bank Account:
    1. Endpoint:"/:id"
    2. Method:DELETE

7. Register Credit Card:
    1. Endpoint:"/creditcards"
    2. Method:POST
    3. Body:{
                "users_id":1,
                "bank":"BCA",
                "limit":10000000
            }

8. Update Credit Card:
    1. Endpoint:"/creditcards"
    2. Method:PUT
    3. Body:{
                "users_id":1,
                "bank":"Mandiri",
                "limit":20000000,
                "ammount":5000000
            }

9. Delete Credit Card:
    1. Endpoint:"/creditcards/:id"
    2. Method:DELETE