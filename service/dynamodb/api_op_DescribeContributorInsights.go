// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about contributor insights, for a given table or global
// secondary index.
func (c *Client) DescribeContributorInsights(ctx context.Context, params *DescribeContributorInsightsInput, optFns ...func(*Options)) (*DescribeContributorInsightsOutput, error) {
	stack := middleware.NewStack("DescribeContributorInsights", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpDescribeContributorInsightsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDescribeContributorInsightsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeContributorInsights(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "DescribeContributorInsights",
			Err:           err,
		}
	}
	out := result.(*DescribeContributorInsightsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeContributorInsightsInput struct {
	// The name of the global secondary index to describe, if applicable.
	IndexName *string
	// The name of the table to describe.
	TableName *string
}

type DescribeContributorInsightsOutput struct {
	// List of names of the associated Alpine rules.
	ContributorInsightsRuleList []*string
	// Current Status contributor insights.
	ContributorInsightsStatus types.ContributorInsightsStatus
	// Returns information about the last failure that encountered. The most common
	// exceptions for a FAILED status are:
	//
	//     * LimitExceededException - Per-account
	// Amazon CloudWatch Contributor Insights rule limit reached. Please disable
	// Contributor Insights for other tables/indexes OR disable Contributor Insights
	// rules before retrying.
	//
	//     * AccessDeniedException - Amazon CloudWatch
	// Contributor Insights rules cannot be modified due to insufficient permissions.
	//
	//
	// * AccessDeniedException - Failed to create service-linked role for Contributor
	// Insights due to insufficient permissions.
	//
	//     * InternalServerError - Failed to
	// create Amazon CloudWatch Contributor Insights rules. Please retry request.
	FailureException *types.FailureException
	// The name of the global secondary index being described.
	IndexName *string
	// Timestamp of the last time the status was changed.
	LastUpdateDateTime *time.Time
	// The name of the table being described.
	TableName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpDescribeContributorInsightsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeContributorInsights{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeContributorInsights{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeContributorInsights(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "DynamoDB GoV2",
		ServiceID:      "dynamodbgov2",
		EndpointPrefix: "dynamodbgov2",
		SigningName:    "dynamodb",
		OperationName:  "DescribeContributorInsights",
	}
}
