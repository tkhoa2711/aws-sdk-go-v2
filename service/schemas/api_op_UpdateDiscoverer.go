// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/schemas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the discoverer
func (c *Client) UpdateDiscoverer(ctx context.Context, params *UpdateDiscovererInput, optFns ...func(*Options)) (*UpdateDiscovererOutput, error) {
	if params == nil {
		params = &UpdateDiscovererInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDiscoverer", params, optFns, c.addOperationUpdateDiscovererMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDiscovererOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDiscovererInput struct {

	// The ID of the discoverer.
	//
	// This member is required.
	DiscovererId *string

	// Support discovery of schemas in events sent to the bus from another account.
	// (default: true)
	CrossAccount bool

	// The description of the discoverer to update.
	Description *string

	noSmithyDocumentSerde
}

type UpdateDiscovererOutput struct {

	// The Status if the discoverer will discover schemas from events sent from another
	// account.
	CrossAccount bool

	// The description of the discoverer.
	Description *string

	// The ARN of the discoverer.
	DiscovererArn *string

	// The ID of the discoverer.
	DiscovererId *string

	// The ARN of the event bus.
	SourceArn *string

	// The state of the discoverer.
	State types.DiscovererState

	// Tags associated with the resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDiscovererMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDiscoverer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDiscoverer{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateDiscovererValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDiscoverer(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateDiscoverer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "UpdateDiscoverer",
	}
}
