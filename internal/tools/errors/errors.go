package errors

import "errors"

var (
UserNotFound = errors.New("UserNotFound")
UserAlreadyExists = errors.New("UserAlreadyExists")
ConflictOnUsers = errors.New("ConflictOnUsers")

RoomAlreadyExists = errors.New("RoomWithThisNameAlreadyExists")
RoomNotFound = errors.New("RoomNotFound")

MemberAlreadyExists = errors.New("MemberAlreadyExists")
MemberNotFound = errors.New("MemberNotFound")
)

