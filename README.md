# Accrual Loyalty Service

Accrual Loyalty Service - Golang microservice example based on concept of Clean Architecture.

For detailed technical specifications, please refer to the [SPECIFICATION.md](https://github.com/alexgaas/order-reward/blob/main/SPECIFICATION.md) file.

### API Endpoints:

* `POST /api/user/register` — User registration.
* `POST /api/user/login` — User authentication.
* `POST /api/orders` — Add an order number for accrual operations.
* `GET /api/orders` — Retrieve a list of user order numbers, including their processing status and accrual information.
* `GET /api/balance` — Get the reward balance for the user.
* `POST /api/orders/withdraw` — Request to withdraw points from the loyal reward account as part of a new order process.
* `GET /api/orders/withdrawals` — Retrieve information about loyal reward account withdrawal operations.

### Build and Run

```BASH
go build -o app cmd/service/main.go
```

You can get help on arguments to start the application by using the `-h` flag:

```BASH
./app -h
Usage of ./app:
  -a string
        Address of application, for example: http://localhost:8000
  -d string
        Database connection string (sqlite), for example: order_reward.db
  -r string
        Accrual system address, for example: http://localhost:8500
```

Once you start the application, it will create all the necessary tables in the database to operate.

### Tags:
* Golang
* Clean Architecture
* Test driven design
* Microservices
* REST API
* [Hoverfly](https://docs.hoverfly.io/en/latest/) (for integration tests)