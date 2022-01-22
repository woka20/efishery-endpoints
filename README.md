# Efishery-Task #
This repo for efishery backend engineer skill test -Woka-
# Table of Contents
* [Auth-App](#auth-app)
	* [Setup](#auth-app-setup)
	* [Build & Run](#auth-app-build-and-run)
	* [Description](#auth-app-description)
    * [Structure Design](#auth-app-structure-design)

* [Fetch-App](#fetch-app)
	* [Setup](#fetch-app-setup)
	* [Build & Run](#fetch-app-build-and-run)
	* [Description](#fetch-app-description)
    * [Structure Design](#fetch-app-structure-design)


## Auth-App
`auth-app` will manage new user, password, and JWT generation process using a file based database.

### Auth-App Setup
To set up the application, you need to have `docker` installed in your machine or `go1.16` if you want to run the app directly in your local machine. If you already have them, set these values in `auth-app/.env` for the app configuration.
* `SERVER_PORT` to define which port that the application will listen to 
	* The default port is `8080`
* `FILE_PATH` to define the path to where the file to store user data
* `SECRET` to define which secret to use to generate/parse JWT

### Auth-App Build and Run
To run the application directly in local machine by using :
```
$ cd auth-app
$ go mod tidy
// run 'go get <name_package>' if needed to pull the needed library/package
$ go run main.go
```
### Auth-App Description
Auth-App has 4 endpoints that can be accessed. 
* GET `/hello` 
-- This endpoint will return a `"Hello World!"` string and `Status Code` `200`. This endpoint can be used to check whether the app is running or not.
* POST `/adduser`
-- This endpoint will add user data that is received from the request body into the database  which is file based  located in `data/alldata.csv`

 and will return the 4 characters password for that user. This endpoint requires request body, other than that will return error
```
//HTTP Request Body (JSON)
{
	"name": "mr.x",
	"phone": "022",
	"role": "admin"
}

//HTTP Response (Application/JSON)
{
	"status_code": 200,
	"data": "<4 characters string>"
}
```
* POST `/token`
-- This endpoint will receive a `phone` and `password` and return a generated JWT using the `SECRET` in `auth-app/.env` to generate string of token that has the correct/matching `phone` and `password`.
```
//HTTP Request Body (JSON)
{
	"phone": "000",
	"password": "<4 characters string>"
}

//HTTP Response (Application/JSON)
{
	"status_code": 200,
	"data": "<JWT token string>"
}
```
* GET `/claims`
-- This endpoint will check the `Authorization` header of the request, verify the JWT inside the header, and return the claims of user information field except password.
```
//HTTP Request Header (Bearer Token)
Authorization: Bearer <JWT>

//HTTP Response (Application/JSON)
{
	"status_code": 200,
	"data": {
		"name": "Mr.x",
		"phone": "000",
		"role": "admin",
		"timestamp": "22 Jan 22 09:07 UTC"
	}
}
```

### Auth-App Structure Design
``` 
Using kataras/iris Golang Web Framework

1. Request Body is parsed and assigned into a context variable which will be handled by Handler functions.
    a. Handler will parse the request body and header and assign those values to variables/struct which is taken from the 'model' package that is shared through all levels/layers of functions.
    b. Handler will also return an HTTP Response after processing the received/parsed data. Whether it's error or successful process.
    c. Handler will only handle request parsing, response, auth, and user/client communication related task.
2. Handler will calls Logic functions that will do data processing with the parsed values from request in variables/struct from 'model' package as the functions parameters.
    a. Logic functions will only do data processing. All the data needed for the process are either given by Handler as parameters or fetched by Repo functions from Database
    b. Logic functions can also be a private functions that can only be called internally in the package, consumed by another function to execute its goal.
3. Repo will fetch data from all external sources, in this case a database, which is using a csv file for the database. No SQL query, only read and write function.                       
```

## Fetch-App
`fetch-app` will fetch and process data/resources
### Fetch-App Setup
Fetch-App environment set its values in `fetch-app/.env` for the app configuration.
* `PORT` to define which port that the application will listen to 
	* The default port is `5000`
* `SECRET` to define which secret to use to generate/parse JWT
* `KEY` to be used when getting currency conversion rate from https://free.currencyconverterapi.com.
    * You need to register and verify your email first on the free plan to get the API Key
* `CACHE_DURATION` to define how long before the cache need to be updated, its sets into 30 minutes as default.


### Fetch-App Build and Run
T to run the application directly in local machine by using :
```
$ cd fetch-app
$ go mod tidy
// run 'go get <name_package>' if needed to pull the needed library/package
$ go run main.go
```

### Fetch-App Summary
Fetch-App has 3 endpoints that can be accessed. 
All endpoints below have middlewares in `fetch-app/routing`  that will verify the JWT secret and/or the role inside JWT is 'admin'. Therefore these endpoints require a valid JWT to work correctly.

//HTTP Request Header (Bearer Token)
Authorization: Bearer <JWT>
```
* GET `/claims`
-- This endpoint will check the `Authorization` header of the request, verify the JWT inside the header, and return the `Private Claims` of the JWT.
```
```
//HTTP Request Header (Bearer Token)
Authorization: Bearer <JWT>

//HTTP Response (Application/JSON)
{
	"status_code": 200,
	"data": {
		"name": "Mr.x",
		"phone": "000",
		"role": "admin",
		"timestamp": "22 Jan 22 09:07 UTC"
	}
}
```
* GET `/commodityList`
-- This endpoint will fetch commodities data from https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list, clean the data from `nil`, `null`, and/or empty values, then returns the cleaned data with additional field of the commodity price in USD currency
```
//HTTP Request Header (Bearer Token)
Authorization: Bearer <JWT>

//HTTP Response (Application/JSON)
{
	"status_code": 200,
    "data": [
       {
			"UUID": "51effbd1-11e6-4b42-8b1c-e3bb39dd7a94",
			"Komoditas": "Ikan Lele",
			"Area_Provinsi": "SULAWESI BARAT",
			"Area_Kota": "BANDUNG",
			"Size": "50",
			"Price": "15000000",
			"Tgl_parsed": "",
			"Timestamp": "",
			"USD_Price": "1046.77332"
		},
		{

			"UUID": "51effbd1-11e6-4b42-8b1c-e3bb39dd7a94",
			"Komoditas": "Ikan Lele",
			"Area_Provinsi": "SULAWESI BARAT",
			"Area_Kota": "BANDUNG",
			"Size": "50",
			"Price": "15000000",
			"Tgl_parsed": "",
			"Timestamp": "",
			"USD_Price": "1046.77332"

		},
        ...
    ]
}
```
* GET `/aggregateData`
-- This endpoint will fetch commodities data from https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list, clean the data from `nil`, `null`, and/or empty values, then returns the aggregated data by the `area_provinsi` value, weekly, and returns the max, min, avg, and median profit (assuming profit is `price` * `size`). This endpoint requires 'admin' role inside the valid JWT in the header of the request.
```
//HTTP Request Header (Bearer Token)
Authorization: Bearer <JWT>

//HTTP Response (Application/JSON)
{
	"status_code": 200,
    "data": [
		        
		  {

			"Province": "BALI",
			"Profit": {
				"2022": {
						"2": 266664,
						"3": 78375975680
		                }
		              },
			"Max": 78375975680,
			"Min": 266664,
			"Average": 39188121172,
			"Median": 39188121172
		},
		{
			"Province": "BANTEN",
			"Profit": {
				"2022": {
						"2": 2222
						 }
					  },
			"Max": 2222,
			"Min": 2222,
			"Average": 2222,
			"Median": 2222
		},
		{
			"Province": "DKI JAKARTA",
			"Profit": {
				"2022": {
					"2": 60000
						}
					},
			"Max": 60000,
			"Min": 60000,
			"Average": 60000,
			"Median": 60000
		},
        ...
    ]
}
```
### Fetch-App System Design
``` 
Using iris Golang Web Framework

1. HTTP Request will be validated and authorized by the Middleware before be forwarded to Handler. Especially the JWT inside the header request.
    a. Failure on validating or authorizing HTTP Request will make Middleware to return an HTTP Response to user/client instead of Handler.
2. HTTP Request is parsed and assigned into a context variable which will be handled by Handler functions.
    a. Handler will parse the request body and header and assign those values to variables/struct which is taken from the 'model' package that is shared through all levels/layers of functions.
    b. Handler will also return an HTTP Response after processing the received/parsed data. Whether it's an error/failed or successful process.
    c. Handler will only handle request parsing, response, auth, and user/client communication related task.
3. Handler will calls Logic functions that will do data processing with the parsed values from request in variables/struct from 'model' package as the functions parameters.
    a. Logic functions will only do data processing. All the data needed for the process are either given by Handler as parameters or fetched by Repo functions from Database.
    b. Logic functions can also be a private functions that can only be called internally in the package. These functions function to help processing the data and will not be called nor return values to Handler.
4. Logic will calls Repo functions to fetch more data that is needed for the process.
    a. Logic will try to get the data from the cache first before calling Repo functions.
    b. There is a worker that will call Repo functions to get all the needed data and store it to the cache at some interval ticker time.
5. Repo will fetch data from all external sources, in this case an external HTTP server that hosts the data. It will fetch the data and parse them and return them before returning the parsed data to Logic.
    a. Fetch-App is getting data from an external HTTP server. Therefore, most of the Repo functions are built to send HTTP requests to the server and parse the response for the data before returning the data to the function that calls it.                       
```
> Written with [StackEdit](https://stackedit.io/).
