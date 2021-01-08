// Code generated by smithy-go-codegen DO NOT EDIT.

package emrcontainers

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emrcontainers/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a virtual cluster. Virtual cluster is a managed entity on Amazon EMR on
// EKS. You can create, describe, list and delete virtual clusters. They do not
// consume any additional resource in your system. A single virtual cluster maps to
// a single Kubernetes namespace. Given this relationship, you can model virtual
// clusters the same way you model Kubernetes namespaces to meet your requirements.
func (c *Client) CreateVirtualCluster(ctx context.Context, params *CreateVirtualClusterInput, optFns ...func(*Options)) (*CreateVirtualClusterOutput, error) {
	if params == nil {
		params = &CreateVirtualClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVirtualCluster", params, optFns, addOperationCreateVirtualClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVirtualClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVirtualClusterInput struct {

	// The client token of the virtual cluster.
	//
	// This member is required.
	ClientToken *string

	// The container provider of the virtual cluster.
	//
	// This member is required.
	ContainerProvider *types.ContainerProvider

	// The specified name of the virtual cluster.
	//
	// This member is required.
	Name *string

	// The tags assigned to the virtual cluster.
	Tags map[string]string
}

type CreateVirtualClusterOutput struct {

	// This output contains the ARN of virtual cluster.
	Arn *string

	// This output contains the virtual cluster ID.
	Id *string

	// This output contains the name of the virtual cluster.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateVirtualClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateVirtualCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVirtualCluster{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = addIdempotencyToken_opCreateVirtualClusterMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateVirtualClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVirtualCluster(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateVirtualCluster struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateVirtualCluster) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateVirtualCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateVirtualClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateVirtualClusterInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateVirtualClusterMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateVirtualCluster{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateVirtualCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "emr-containers",
		OperationName: "CreateVirtualCluster",
	}
}