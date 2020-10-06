// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a specified reusable delegation set, including the
// four name servers that are assigned to the delegation set.
func (c *Client) GetReusableDelegationSet(ctx context.Context, params *GetReusableDelegationSetInput, optFns ...func(*Options)) (*GetReusableDelegationSetOutput, error) {
	stack := middleware.NewStack("GetReusableDelegationSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetReusableDelegationSetMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetReusableDelegationSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetReusableDelegationSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addSanitizeURLMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "GetReusableDelegationSet",
			Err:           err,
		}
	}
	out := result.(*GetReusableDelegationSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get information about a specified reusable delegation set.
type GetReusableDelegationSetInput struct {

	// The ID of the reusable delegation set that you want to get a list of name
	// servers for.
	//
	// This member is required.
	Id *string
}

// A complex type that contains the response to the GetReusableDelegationSet
// request.
type GetReusableDelegationSetOutput struct {

	// A complex type that contains information about the reusable delegation set.
	//
	// This member is required.
	DelegationSet *types.DelegationSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetReusableDelegationSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetReusableDelegationSet{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetReusableDelegationSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetReusableDelegationSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetReusableDelegationSet",
	}
}
