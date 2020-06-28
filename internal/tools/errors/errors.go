package errors

import "errors"


var UserNotFound = errors.New("UserNotFound")
var UserAlreadyExists = errors.New("UserAlreadyExists")
var ConflictOnUsers = errors.New("ConflictOnUsers")
