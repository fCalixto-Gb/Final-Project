package errors

import (
	"errors"
	"fmt"

	"github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/user/proto"
)

type ErrUserNotFound struct {
	Err error
}

func (e *ErrUserNotFound) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewErrUserNotFound() *ErrUserNotFound {
	return &ErrUserNotFound{Err: errors.New("user not found")}
}

type ErrAlreadyExists struct {
	Err error
}

func (e *ErrAlreadyExists) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewErrAlreadyExists() *ErrAlreadyExists {
	return &ErrAlreadyExists{Err: errors.New("user already exists")}
}

type ErrUnkown struct {
	Err error
}

func (e *ErrUnkown) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewErrUnkown() *ErrUnkown {
	return &ErrUnkown{Err: errors.New("unkown error")}
}

type ErrInvalidArguments struct {
	Err error
}

func (e *ErrInvalidArguments) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewErrInvalidArguments() *ErrInvalidArguments {
	return &ErrInvalidArguments{Err: errors.New("invalid argument(s)")}
}

type ErrInvalidCredentials struct {
	Err error
}

func (e *ErrInvalidCredentials) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewErrInvalidCredentials() *ErrInvalidCredentials {
	return &ErrInvalidCredentials{Err: errors.New("invalid credential(s)")}
}

// Receives a custom error, returns the corresponding proto status struct filled
// Used in grpc service - transport layer
func ErrToGRPCcode(e error) *proto.Status {
	var status proto.Status
	switch e.(type) {
	case *ErrUserNotFound:
		status.Code = 5
		status.Message = e.Error()
	case *ErrAlreadyExists:
		status.Code = 6
		status.Message = e.Error()
	case *ErrInvalidArguments:
		status.Code = 3
		status.Message = e.Error()
	case *ErrInvalidCredentials:
		status.Code = 7
		status.Message = e.Error()
	default:
		status.Code = 2
		status.Message = "unkown error"
	}
	return &status
}
