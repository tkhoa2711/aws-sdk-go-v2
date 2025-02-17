// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the resources that you added to a resource shares or the resources that
// are shared with you.
func (c *Client) ListResources(ctx context.Context, params *ListResourcesInput, optFns ...func(*Options)) (*ListResourcesOutput, error) {
	if params == nil {
		params = &ListResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResources", params, optFns, c.addOperationListResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourcesInput struct {

	// The type of owner.
	//
	// This member is required.
	ResourceOwner types.ResourceOwner

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The principal.
	Principal *string

	// The Amazon Resource Names (ARNs) of the resources.
	ResourceArns []string

	// The Amazon Resource Names (ARN) of the resource shares.
	ResourceShareArns []string

	// The resource type. Valid values: acm-pca:CertificateAuthority | appmesh:Mesh |
	// codebuild:Project | codebuild:ReportGroup | ec2:CapacityReservation |
	// ec2:DedicatedHost | ec2:LocalGatewayRouteTable | ec2:PrefixList | ec2:Subnet |
	// ec2:TrafficMirrorTarget | ec2:TransitGateway | imagebuilder:Component |
	// imagebuilder:Image | imagebuilder:ImageRecipe | imagebuilder:ContainerRecipe |
	// glue:Catalog | glue:Database | glue:Table | license-manager:LicenseConfiguration
	// I network-firewall:FirewallPolicy | network-firewall:StatefulRuleGroup |
	// network-firewall:StatelessRuleGroup | outposts:Outpost | resource-groups:Group |
	// rds:Cluster | route53resolver:FirewallRuleGroup
	// |route53resolver:ResolverQueryLogConfig | route53resolver:ResolverRule |
	// s3-outposts:Outpost | ssm-contacts:Contact | ssm-incidents:ResponsePlan
	ResourceType *string

	noSmithyDocumentSerde
}

type ListResourcesOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the resources.
	Resources []types.Resource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListResources{}, middleware.After)
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
	if err = addOpListResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResources(options.Region), middleware.Before); err != nil {
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

// ListResourcesAPIClient is a client that implements the ListResources operation.
type ListResourcesAPIClient interface {
	ListResources(context.Context, *ListResourcesInput, ...func(*Options)) (*ListResourcesOutput, error)
}

var _ ListResourcesAPIClient = (*Client)(nil)

// ListResourcesPaginatorOptions is the paginator options for ListResources
type ListResourcesPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourcesPaginator is a paginator for ListResources
type ListResourcesPaginator struct {
	options   ListResourcesPaginatorOptions
	client    ListResourcesAPIClient
	params    *ListResourcesInput
	nextToken *string
	firstPage bool
}

// NewListResourcesPaginator returns a new ListResourcesPaginator
func NewListResourcesPaginator(client ListResourcesAPIClient, params *ListResourcesInput, optFns ...func(*ListResourcesPaginatorOptions)) *ListResourcesPaginator {
	if params == nil {
		params = &ListResourcesInput{}
	}

	options := ListResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourcesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListResources page.
func (p *ListResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListResources(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "ListResources",
	}
}
