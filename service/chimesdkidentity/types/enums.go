// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeBadRequest                           ErrorCode = "BadRequest"
	ErrorCodeConflict                             ErrorCode = "Conflict"
	ErrorCodeForbidden                            ErrorCode = "Forbidden"
	ErrorCodeNotFound                             ErrorCode = "NotFound"
	ErrorCodePreconditionFailed                   ErrorCode = "PreconditionFailed"
	ErrorCodeResourceLimitExceeded                ErrorCode = "ResourceLimitExceeded"
	ErrorCodeServiceFailure                       ErrorCode = "ServiceFailure"
	ErrorCodeAccessDenied                         ErrorCode = "AccessDenied"
	ErrorCodeServiceUnavailable                   ErrorCode = "ServiceUnavailable"
	ErrorCodeThrottled                            ErrorCode = "Throttled"
	ErrorCodeThrottling                           ErrorCode = "Throttling"
	ErrorCodeUnauthorized                         ErrorCode = "Unauthorized"
	ErrorCodeUnprocessable                        ErrorCode = "Unprocessable"
	ErrorCodeVoiceConnectorGroupAssociationsExist ErrorCode = "VoiceConnectorGroupAssociationsExist"
	ErrorCodePhoneNumberAssociationsExist         ErrorCode = "PhoneNumberAssociationsExist"
)

// Values returns all known values for ErrorCode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"BadRequest",
		"Conflict",
		"Forbidden",
		"NotFound",
		"PreconditionFailed",
		"ResourceLimitExceeded",
		"ServiceFailure",
		"AccessDenied",
		"ServiceUnavailable",
		"Throttled",
		"Throttling",
		"Unauthorized",
		"Unprocessable",
		"VoiceConnectorGroupAssociationsExist",
		"PhoneNumberAssociationsExist",
	}
}
