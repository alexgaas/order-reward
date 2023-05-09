# Technical specification

## Accrual Loyalty Service

The "Accrual Loyalty Service" (hereafter referred to as "the Service") is a service that is described in the specification below. 
The Service must call the "Loyalty Points Calculation Service" (hereafter referred to as "the Points"), 
which is a third-party API used for collecting rewards in the processing of accrual.

---

### Functional requirements

The Service consists of a set of HTTP API endpoints that meet the following requirements:

* Registration, authentication, and authorization of users.
* Acceptance of order numbers from registered users.
* Accounting and retrieval of order numbers per registered user.
* Accounting and retrieval of loyalty reward account per registered user.
* Validation of accepted order numbers.
* Accrual of a defined reward to the loyal user account for each valid order.

With these endpoints, users can register, authenticate, and authorize themselves to use the service. 
Once registered, they can submit order numbers, view their order history, and check their loyalty reward account balance. 
The service also ensures that submitted order numbers are valid and accrues a reward to the user's loyalty account for each valid order.

### Common logic of communication user and Service

The following steps describe the common logic of communication between the user and the loyalty service:

1. The user registers for the loyalty service.
2. The user purchases something from an online store (which is not implemented in this service).
3. The online store sends the order information to the loyalty service (which is not implemented in this service).
4. The user provides the order number to the loyalty service.
5. The loyalty service links the order number with the user and validates the order.
6. If the user has a correct/positive account, the loyalty service accrues points to the user's account.
7. The user can then withdraw the available points to purchase any other order, either fully or partially.

By following these steps, users can participate in the loyalty program and earn points that can be used to purchase items in the future. 
The loyalty service ensures that the points are accurately accrued and available for the user to withdraw when needed.

### HTTP API

API Endpoints:

* `POST /api/user/register` — User registration.
* `POST /api/user/login` — User authentication.
* `POST /api/orders` — Add an order number for accrual operations.
* `GET /api/orders` — Retrieve a list of user order numbers, including their processing status and accrual information.
* `POST /api/orders/withdraw` — Request to withdraw points from the loyal reward account as part of a new order process.
* `GET /api/orders/withdrawals` — Retrieve information about loyal reward account withdrawal operations.
* `GET /api/balance` — Get the reward balance for the user.

Non-functional (technical) requirements:

* Database storage should use SQLite.
* The table structure should be based on the functional requirements and HTTP API definition.
* The type and format of stored authorization information (password or any other sensitive information) is not defined and can be decided based on best practices.
* The client is allowed to use/support HTTP requests/responses with compression.
* Although the client is not obligated to follow HTTP specification, it is highly recommended to validate HTTP requests/responses.
* The user authorization/authentication algorithm is not defined and can be implemented based on the chosen approach.
* Order numbers should be unique and never repeat.
* One order number can be added to processing one time per one user only.
* An order may not have any point accrual.
* The reward is calculated in virtual points, assuming 1 point equals 1 $USD.

These API endpoints and non-functional requirements ensure that users can easily register and authenticate themselves, 
add order numbers for accrual operations, and retrieve information about their orders and reward balances. 
The non-functional requirements also ensure that the service uses best practices for data storage and HTTP requests and responses, 
and that user authentication and authorization are secure.

#### **User Registration**

Handler: `POST /api/user/register`.

This API handler is used to register a user with a unique login and password. 
Upon successful registration, the service returns authentication information as a response. 
The following request specification should be followed:

Request Specification:

```
POST /api/user/register HTTP/1.1
Content-Type: application/json
...

{
	"login": "<login>",
	"password": "<password>"
}
```

Response Codes:

- `200` — The user has been successfully registered and authenticated.
- `400` — The request is incorrect or malformed.
- `409` — The specified login already exists in the system.
- `500` — An internal server error has occurred.

#### **User Login**

Handler: `POST /api/user/login`.

This API handler is used to authenticate a user with a unique login and password.
Upon successful authentication, the service returns authentication information as a response.
The following request specification should be followed:

Request Specification:

```
POST /api/user/login HTTP/1.1
Content-Type: application/json
...

{
	"login": "<login>",
	"password": "<password>"
}
```

Response Codes:

- `200` — The user has been successfully authenticated.
- `400` — The request is incorrect or malformed.
- `500` — An internal server error has occurred.

#### **Add Order Number**

Handler: `POST /api/orders`.

This endpoint is only available to authenticated users. 
The order number is a number sequence of random length that can be validated for correctness using the Luhn algorithm. 
The following request specification should be followed:

Request Specification:

```
POST /api/orders HTTP/1.1
Content-Type: text/plain
...

12345678903
```

Response Codes:

- `200` — The specified order number has already been added by this user.
- `202` — The specified order number has been added to the processing queue successfully.
- `400` — The request is incorrect or malformed.
- `401` — The user has not been authenticated.
- `409` — The specified order number has already been added by another user.
- `422` — The specified order number has an invalid format that does not comply with the Luhn algorithm.
- `500` — An internal server error has occurred.

#### **Get List Of Orders**

Handler: `GET /api/orders`.

This endpoint is only available to authenticated users.
The response must include the order number sorted in descending order, with the newest on top. 
The date format should be in RFC3339.

The following are the available statuses to manage accrual of rewards:

- `NEW` — order added to the system, but the accrual process for that order has not been started yet.
- `PROCESSING` — accrual process started.
- `INVALID` — during the process of accrual, the service encountered an error.
- `PROCESSED` — accrual process has been successful, and reward has been added successfully.

Request Specification:

```
GET /api/orders HTTP/1.1
Content-Length: 0
```

Response Codes:

- `200` — request was processed successfully. The response specification is as follows:

    ```
    200 OK HTTP/1.1
    Content-Type: application/json
    ...
    
    [
    	{
            "number": "9278923470",
            "status": "PROCESSED",
            "accrual": 500,
            "uploaded_at": "2020-12-10T15:15:45+03:00"
        },
        {
            "number": "12345678903",
            "status": "PROCESSING",
            "uploaded_at": "2020-12-10T15:12:01+03:00"
        },
        {
            "number": "346436439",
            "status": "INVALID",
            "uploaded_at": "2020-12-09T16:09:53+03:00"
        }
    ]
    ```

- `204` — Data not found.
- `401` — The user has not been authenticated.
- `500` — An internal server error has occurred.

#### **Get User Reward Balance**

Handler: `GET /api/balance`.

This endpoint is only accessible to authenticated users. 
The response will include the current reward balance and the total sum of rewards used over the entire period of reward activity.

Request Specification:

```
GET /api/balance HTTP/1.1
Content-Length: 0
```

Response codes:

- `200` — request been processed successfully.

  Response specification:

    ```
    200 OK HTTP/1.1
    Content-Type: application/json
    ...
    
    {
    	"current": 500.5,
    	"withdrawn": 42
    }
    ```

- `401` — The user has not been authenticated.
- `500` — An internal server error has occurred.

#### **Request Balance Withdrawal**

Handler: `POST /api/orders/withdraw`

This handler is available only to authenticated users. Users can specify the order number and the sum of points 
they want to withdraw in the request body, where order is an abstract order number used to partially or fully pay using rewards, 
and sum is the number of points to withdraw.

Please note that there is no need to use any external API or make any external calls to successfully withdraw.

Request Specification:

```
POST /api/orders/withdraw HTTP/1.1
Content-Type: application/json

{
    "order": "2377225624",
    "sum": 751
}
```

Response Codes:

- `200` — The request has been processed successfully.
- `401` — The user has not been authenticated.
- `402` — There are not enough points on the user's account.
- `422` — The order number is invalid.
- `500` — An internal server error has occurred.

#### **Get Withdraw Information**

Handler: `GET /api/orders/withdrawals`.

This handler is available only to authenticated users. 
Withdrawals in response must be sorted out from newest to oldest by date. Date format - RFC3339.

Request Specification:

```
GET /api/orders/withdrawals HTTP/1.1
Content-Length: 0
```

Response Codes:

- `200` — The request has been processed successfully.

  Response Body:

    ```
    200 OK HTTP/1.1
    Content-Type: application/json
    ...
    
    [
        {
            "order": "2377225624",
            "sum": 500,
            "processed_at": "2020-12-09T16:09:57+03:00"
        }
    ]
    ```

- `204` - There is no any withdrawal yet from user's balance.
- `401` — The user has not been authenticated.
- `500` — An internal server error has occurred.

### Communication with Loyalty Points Calculation Service

Handler: `GET /orders/{number}`. 

This handler retrieves data about reward processing of accrual.

Request Specification:

```
GET /orders/{number} HTTP/1.1
Content-Length: 0
```

Response Codes:

- `200` — The request has been processed successfully.

  Response Body:

    ```
    200 OK HTTP/1.1
    Content-Type: application/json
    ...
    
    {
        "order": "<number>",
        "status": "PROCESSED",
        "accrual": 500
    }
    ```

  Field definitions in response body:

    - `order` — the order number.
    - `status` — the processing status of the accrual:

        - `REGISTERED` — the order is registered, but the accrual process has not yet started.
        - `INVALID` — the order has not been approved for the accrual process, and the reward will not be added to the user's account.
        - `PROCESSING` — the accrual process is in progress.
        - `PROCESSED` — the accrual process has been successfully completed.

    - `accrual` — the reward points to be added to the user's account. If there are no points to reward, this field should not be included in the payload.

- `429` — the number of requests has been exceeded.

  Response body:

    ```
    429 Too Many Requests HTTP/1.1
    Content-Type: text/plain
    Retry-After: 60
    
    No more than N requests per minute allowed
    ```

- `500` — An internal server error has occurred.

Orders can be added to the accrual process at any time after being added to the system. 
There are no specifications or restrictions on the timing of accrual. 
The `INVALID` and `PROCESSED` statuses are final and cannot be recalculated.

The general number of requests is not restricted.

### Configuration of Oder Loyalty System

The Order Loyalty System service supports the following configuration approaches:

- The address and port to start the service can be specified using the environment variable `RUN_ADDRESS` or the flag `-a`.
- The connection string to the database can be specified using the environment variable `DATABASE_URI` or the flag `-d`.
- The address of the accrual system can be specified using the environment variable `ACCRUAL_SYSTEM_ADDRESS` or the flag `-r`.