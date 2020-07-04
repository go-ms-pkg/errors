//Package errors provides custom implementation of errors for
//basic error logging and stack tracing.
package errors

//Function is a custom alias of type string to define which function
//and operation encountered an error.
type Layer string

//Method is a custom alias of type string which identified which method
//along the stack caused the error.
type Method string

//Type is a custom alias of type int to define what type of error
//occurred.  Pre-defined constants of usual errors typically
//encountered are defined in this package.
type Type int

//Pre-defined constants of usual errors encountered.
//
//NotFound - For Resource Not Found Errors - HTTP 404
//
//Unauthorized - For Unauthorized Access of Resources - HTTP 403
//
//InternalServerError - For Generic Internal Server Errors - HTTP 500
//
//BadRequest - For Erroneous API Requests - HTTP 400
//
//DatabaseOperationFailed - For slightly opinionated database level
//related errors
//
//ServiceOperationFailed - For slightly opinionated service operation
//level errors
//
//HandlerOperationFailed - For slightly opinionated handler level
//operation errors
//
//EncodingFailed - For encoding errors resulting from marshal processes
//
//DecodingFailed - For decoding errors resulting from unmarshal processes
const (
	NotFound Type = iota
	Unauthorized
	InternalServerError
	BadRequest
	DatabaseOperationFailed
	ServiceOperationFailed
	HandlerOperationFailed
	EncodingFailed
	DecodingFailed
)

//Error is a custom struct definition with generic/basic information
//for errors encountered by workloads designed to wrap source errors
//for stack tracing purposes when bubbled up to caller.
//L - layer of service in which the error is encountered
//M - method call in which the error is encountered
//T - type of error encountered (one of the pre-defined types)
//E - wrapped error
type Error struct {
	L Layer
	M Method
	T Type
	E error
}

func (e *Error) Error() string {
	return ""
}
