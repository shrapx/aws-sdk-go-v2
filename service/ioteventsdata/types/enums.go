// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeResourceNotFoundException   ErrorCode = "ResourceNotFoundException"
	ErrorCodeInvalidRequestException     ErrorCode = "InvalidRequestException"
	ErrorCodeInternalFailureException    ErrorCode = "InternalFailureException"
	ErrorCodeServiceUnavailableException ErrorCode = "ServiceUnavailableException"
	ErrorCodeThrottlingException         ErrorCode = "ThrottlingException"
)

// Values returns all known values for ErrorCode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"ResourceNotFoundException",
		"InvalidRequestException",
		"InternalFailureException",
		"ServiceUnavailableException",
		"ThrottlingException",
	}
}
