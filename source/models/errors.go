package models

type ErrorCode int

const (
	ErrorRequestDataInvalid ErrorCode = 1
	ErrorDatabase           ErrorCode = 2
	ErrorJsonMarshal        ErrorCode = 3

	ErrorOther ErrorCode = 4
)
