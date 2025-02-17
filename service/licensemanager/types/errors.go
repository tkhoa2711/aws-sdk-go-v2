// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Access to resource denied.
type AccessDeniedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The AWS user account does not have permission to perform the action. Check the
// IAM policy associated with this account.
type AuthorizationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *AuthorizationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AuthorizationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AuthorizationException) ErrorCode() string             { return "AuthorizationException" }
func (e *AuthorizationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There was a conflict processing the request. Try your request again.
type ConflictException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string             { return "ConflictException" }
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The entitlement is not allowed.
type EntitlementNotAllowedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *EntitlementNotAllowedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EntitlementNotAllowedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EntitlementNotAllowedException) ErrorCode() string             { return "EntitlementNotAllowedException" }
func (e *EntitlementNotAllowedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A dependency required to run the API is missing.
type FailedDependencyException struct {
	Message *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *FailedDependencyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FailedDependencyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FailedDependencyException) ErrorCode() string             { return "FailedDependencyException" }
func (e *FailedDependencyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request uses too many filters or too many filter values.
type FilterLimitExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *FilterLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FilterLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FilterLimitExceededException) ErrorCode() string             { return "FilterLimitExceededException" }
func (e *FilterLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more parameter values are not valid.
type InvalidParameterValueException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterValueException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterValueException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterValueException) ErrorCode() string             { return "InvalidParameterValueException" }
func (e *InvalidParameterValueException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// License Manager cannot allocate a license to a resource because of its state.
// For example, you cannot allocate a license to an instance in the process of
// shutting down.
type InvalidResourceStateException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidResourceStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidResourceStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidResourceStateException) ErrorCode() string             { return "InvalidResourceStateException" }
func (e *InvalidResourceStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You do not have enough licenses available to support a new resource launch.
type LicenseUsageException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *LicenseUsageException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LicenseUsageException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LicenseUsageException) ErrorCode() string             { return "LicenseUsageException" }
func (e *LicenseUsageException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There are no entitlements found for this license, or the entitlement maximum
// count is reached.
type NoEntitlementsAllowedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *NoEntitlementsAllowedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoEntitlementsAllowedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoEntitlementsAllowedException) ErrorCode() string             { return "NoEntitlementsAllowedException" }
func (e *NoEntitlementsAllowedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Too many requests have been submitted. Try again after a brief wait.
type RateLimitExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *RateLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RateLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RateLimitExceededException) ErrorCode() string             { return "RateLimitExceededException" }
func (e *RateLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This is not the correct Region for the resource. Try again.
type RedirectException struct {
	Message *string

	Location *string

	noSmithyDocumentSerde
}

func (e *RedirectException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RedirectException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RedirectException) ErrorCode() string             { return "RedirectException" }
func (e *RedirectException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Your resource limits have been exceeded.
type ResourceLimitExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceLimitExceededException) ErrorCode() string             { return "ResourceLimitExceededException" }
func (e *ResourceLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource cannot be found.
type ResourceNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The server experienced an internal error. Try again.
type ServerInternalException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ServerInternalException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServerInternalException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServerInternalException) ErrorCode() string             { return "ServerInternalException" }
func (e *ServerInternalException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The digital signature method is unsupported. Try your request again.
type UnsupportedDigitalSignatureMethodException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UnsupportedDigitalSignatureMethodException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedDigitalSignatureMethodException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedDigitalSignatureMethodException) ErrorCode() string {
	return "UnsupportedDigitalSignatureMethodException"
}
func (e *UnsupportedDigitalSignatureMethodException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The provided input is not valid. Try your request again.
type ValidationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string             { return "ValidationException" }
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
