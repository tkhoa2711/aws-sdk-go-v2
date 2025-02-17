// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new version of a model and begins training. Models are managed as part
// of an Amazon Rekognition Custom Labels project. You can specify one training
// dataset and one testing dataset. The response from CreateProjectVersion is an
// Amazon Resource Name (ARN) for the version of the model. Training takes a while
// to complete. You can get the current status by calling DescribeProjectVersions.
// Once training has successfully completed, call DescribeProjectVersions to get
// the training results and evaluate the model. After evaluating the model, you
// start the model by calling StartProjectVersion. This operation requires
// permissions to perform the rekognition:CreateProjectVersion action.
func (c *Client) CreateProjectVersion(ctx context.Context, params *CreateProjectVersionInput, optFns ...func(*Options)) (*CreateProjectVersionOutput, error) {
	if params == nil {
		params = &CreateProjectVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProjectVersion", params, optFns, c.addOperationCreateProjectVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProjectVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProjectVersionInput struct {

	// The Amazon S3 bucket location to store the results of training. The S3 bucket
	// can be in any AWS account as long as the caller has s3:PutObject permissions on
	// the S3 bucket.
	//
	// This member is required.
	OutputConfig *types.OutputConfig

	// The ARN of the Amazon Rekognition Custom Labels project that manages the model
	// that you want to train.
	//
	// This member is required.
	ProjectArn *string

	// The dataset to use for testing.
	//
	// This member is required.
	TestingData *types.TestingData

	// The dataset to use for training.
	//
	// This member is required.
	TrainingData *types.TrainingData

	// A name for the version of the model. This value must be unique.
	//
	// This member is required.
	VersionName *string

	// The identifier for your AWS Key Management Service (AWS KMS) customer master key
	// (CMK). You can supply the Amazon Resource Name (ARN) of your CMK, the ID of your
	// CMK, an alias for your CMK, or an alias ARN. The key is used to encrypt training
	// and test images copied into the service for model training. Your source images
	// are unaffected. The key is also used to encrypt training results and manifest
	// files written to the output Amazon S3 bucket (OutputConfig). If you choose to
	// use your own CMK, you need the following permissions on the CMK.
	//
	// *
	// kms:CreateGrant
	//
	// * kms:DescribeKey
	//
	// * kms:GenerateDataKey
	//
	// * kms:Decrypt
	//
	// If you
	// don't specify a value for KmsKeyId, images copied into the service are encrypted
	// using a key that AWS owns and manages.
	KmsKeyId *string

	// A set of tags (key-value pairs) that you want to attach to the model.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateProjectVersionOutput struct {

	// The ARN of the model version that was created. Use DescribeProjectVersion to get
	// the current status of the training operation.
	ProjectVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProjectVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProjectVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProjectVersion{}, middleware.After)
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
	if err = addOpCreateProjectVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProjectVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateProjectVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "CreateProjectVersion",
	}
}
