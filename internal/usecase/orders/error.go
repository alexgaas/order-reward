package orders_usecase

import "errors"

var ErrOrderNumberIsNotValid = errors.New("order number is not valid")
var ErrNegativeSum = errors.New("sum must be > 0")
