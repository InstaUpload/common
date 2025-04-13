package types

import "errors"

var ErrDataFound = errors.New("Data found, but not expected")
var ErrDataNotFound = errors.New("Data not found, but expected")
var ErrIncorrectDataReceived = errors.New("Incorrect data received")
var ErrUnauthorized = errors.New("Unauthorized")
