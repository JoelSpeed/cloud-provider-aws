// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allocates a Dedicated Host to your account. At a minimum, specify the supported
// instance type or instance family, the Availability Zone in which to allocate the
// host, and the number of hosts to allocate.
func (c *Client) AllocateHosts(ctx context.Context, params *AllocateHostsInput, optFns ...func(*Options)) (*AllocateHostsOutput, error) {
	if params == nil {
		params = &AllocateHostsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AllocateHosts", params, optFns, c.addOperationAllocateHostsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AllocateHostsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AllocateHostsInput struct {

	// The Availability Zone in which to allocate the Dedicated Host.
	//
	// This member is required.
	AvailabilityZone *string

	// The IDs of the Outpost hardware assets on which to allocate the Dedicated
	// Hosts. Targeting specific hardware assets on an Outpost can help to minimize
	// latency between your workloads. This parameter is supported only if you specify
	// OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this
	// parameter.
	//
	//   - If you specify this parameter, you can omit Quantity. In this case, Amazon
	//   EC2 allocates a Dedicated Host on each specified hardware asset.
	//
	//   - If you specify both AssetIds and Quantity, then the value for Quantity must
	//   be equal to the number of asset IDs specified.
	AssetIds []string

	// Indicates whether the host accepts any untargeted instance launches that match
	// its instance type configuration, or if it only accepts Host tenancy instance
	// launches that specify its unique host ID. For more information, see [Understanding auto-placement and affinity]in the
	// Amazon EC2 User Guide.
	//
	// Default: off
	//
	// [Understanding auto-placement and affinity]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/how-dedicated-hosts-work.html#dedicated-hosts-understanding
	AutoPlacement types.AutoPlacement

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see [Ensuring Idempotency].
	//
	// [Ensuring Idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	ClientToken *string

	// Indicates whether to enable or disable host maintenance for the Dedicated Host.
	// For more information, see [Host maintenance]in the Amazon EC2 User Guide.
	//
	// [Host maintenance]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-maintenance.html
	HostMaintenance types.HostMaintenance

	// Indicates whether to enable or disable host recovery for the Dedicated Host.
	// Host recovery is disabled by default. For more information, see [Host recovery]in the Amazon
	// EC2 User Guide.
	//
	// Default: off
	//
	// [Host recovery]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-recovery.html
	HostRecovery types.HostRecovery

	// Specifies the instance family to be supported by the Dedicated Hosts. If you
	// specify an instance family, the Dedicated Hosts support multiple instance types
	// within that instance family.
	//
	// If you want the Dedicated Hosts to support a specific instance type only, omit
	// this parameter and specify InstanceType instead. You cannot specify
	// InstanceFamily and InstanceType in the same request.
	InstanceFamily *string

	// Specifies the instance type to be supported by the Dedicated Hosts. If you
	// specify an instance type, the Dedicated Hosts support instances of the specified
	// instance type only.
	//
	// If you want the Dedicated Hosts to support multiple instance types in a
	// specific instance family, omit this parameter and specify InstanceFamily
	// instead. You cannot specify InstanceType and InstanceFamily in the same request.
	InstanceType *string

	// The Amazon Resource Name (ARN) of the Amazon Web Services Outpost on which to
	// allocate the Dedicated Host. If you specify OutpostArn, you can optionally
	// specify AssetIds.
	//
	// If you are allocating the Dedicated Host in a Region, omit this parameter.
	OutpostArn *string

	// The number of Dedicated Hosts to allocate to your account with these
	// parameters. If you are allocating the Dedicated Hosts on an Outpost, and you
	// specify AssetIds, you can omit this parameter. In this case, Amazon EC2
	// allocates a Dedicated Host on each specified hardware asset. If you specify both
	// AssetIds and Quantity, then the value that you specify for Quantity must be
	// equal to the number of asset IDs specified.
	Quantity *int32

	// The tags to apply to the Dedicated Host during creation.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

// Contains the output of AllocateHosts.
type AllocateHostsOutput struct {

	// The ID of the allocated Dedicated Host. This is used to launch an instance onto
	// a specific host.
	HostIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAllocateHostsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpAllocateHosts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAllocateHosts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AllocateHosts"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpAllocateHostsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAllocateHosts(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAllocateHosts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AllocateHosts",
	}
}
