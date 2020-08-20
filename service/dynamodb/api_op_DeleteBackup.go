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
)

// Deletes an existing backup of a table. You can call DeleteBackup at a maximum
// rate of 10 times per second.
func (c *Client) DeleteBackup(ctx context.Context, params *DeleteBackupInput, optFns ...func(*Options)) (*DeleteBackupOutput, error) {
	stack := middleware.NewStack("DeleteBackup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpDeleteBackupMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteBackupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBackup(options.Region), middleware.Before)

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
			OperationName: "DeleteBackup",
			Err:           err,
		}
	}
	out := result.(*DeleteBackupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBackupInput struct {
	// The ARN associated with the backup.
	BackupArn *string
}

type DeleteBackupOutput struct {
	// Contains the description of the backup created for the table.
	BackupDescription *types.BackupDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpDeleteBackupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteBackup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteBackup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteBackup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "DynamoDB GoV2",
		ServiceID:      "dynamodbgov2",
		EndpointPrefix: "dynamodbgov2",
		SigningName:    "dynamodb",
		OperationName:  "DeleteBackup",
	}
}
