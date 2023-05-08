Models:
    User (login, password)
    Order (User, number, status, accrual, uploaded_at)
    OrderLog (User, orderNumber, sum, processed_at)
    Account (User, balance)