package repository

import "errors"

var ErrUserAlreadyExists = errors.New("user with such credentials already exist")
var ErrInvalidLoginPassword = errors.New("invalid login/password")
var ErrNoOrders = errors.New("orders not found")
var ErrOrderExists = errors.New("order early uploaded")
var ErrOrderExistsAnother = errors.New("order early uploaded another user")
