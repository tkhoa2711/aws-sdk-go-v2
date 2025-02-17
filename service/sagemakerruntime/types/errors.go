// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// An internal failure occurred.
type InternalFailure struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InternalFailure) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalFailure) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalFailure) ErrorCode() string             { return "InternalFailure" }
func (e *InternalFailure) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Model (owned by the customer in the container) returned 4xx or 5xx error code.
type ModelError struct {
	Message *string

	OriginalStatusCode *int32
	OriginalMessage    *string
	LogStreamArn       *string

	noSmithyDocumentSerde
}

func (e *ModelError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ModelError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ModelError) ErrorCode() string             { return "ModelError" }
func (e *ModelError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service is unavailable. Try your call again.
type ServiceUnavailable struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ServiceUnavailable) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailable) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailable) ErrorCode() string             { return "ServiceUnavailable" }
func (e *ServiceUnavailable) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Inspect your request and try again.
type ValidationError struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationError) ErrorCode() string             { return "ValidationError" }
func (e *ValidationError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
