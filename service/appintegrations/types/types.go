// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// The event filter.
type EventFilter struct {

	// The source of the events.
	//
	// This member is required.
	Source *string

	noSmithyDocumentSerde
}

// The event integration.
type EventIntegration struct {

	// The event integration description.
	Description *string

	// The Amazon EventBridge bus for the event integration.
	EventBridgeBus *string

	// The event integration filter.
	EventFilter *EventFilter

	// The Amazon Resource Name (ARN) of the event integration.
	EventIntegrationArn *string

	// The name of the event integration.
	Name *string

	// The tags.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The event integration association.
type EventIntegrationAssociation struct {

	// The metadata associated with the client.
	ClientAssociationMetadata map[string]string

	// The identifier for the client that is associated with the event integration.
	ClientId *string

	// The name of the EventBridge rule.
	EventBridgeRuleName *string

	// The Amazon Resource Name (ARN) for the event integration association.
	EventIntegrationAssociationArn *string

	// The identifier for the event integration association.
	EventIntegrationAssociationId *string

	// The name of the event integration.
	EventIntegrationName *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
