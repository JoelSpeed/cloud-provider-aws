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

// Imports a disk into an EBS snapshot.
//
// For more information, see [Importing a disk as a snapshot using VM Import/Export] in the VM Import/Export User Guide.
//
// [Importing a disk as a snapshot using VM Import/Export]: https://docs.aws.amazon.com/vm-import/latest/userguide/vmimport-import-snapshot.html
func (c *Client) ImportSnapshot(ctx context.Context, params *ImportSnapshotInput, optFns ...func(*Options)) (*ImportSnapshotOutput, error) {
	if params == nil {
		params = &ImportSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportSnapshot", params, optFns, c.addOperationImportSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportSnapshotInput struct {

	// The client-specific data.
	ClientData *types.ClientData

	// Token to enable idempotency for VM import requests.
	ClientToken *string

	// The description string for the import snapshot task.
	Description *string

	// Information about the disk container.
	DiskContainer *types.SnapshotDiskContainer

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Specifies whether the destination snapshot of the imported image should be
	// encrypted. The default KMS key for EBS is used unless you specify a non-default
	// KMS key using KmsKeyId . For more information, see [Amazon EBS Encryption] in the Amazon Elastic
	// Compute Cloud User Guide.
	//
	// [Amazon EBS Encryption]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html
	Encrypted *bool

	// An identifier for the symmetric KMS key to use when creating the encrypted
	// snapshot. This parameter is only required if you want to use a non-default KMS
	// key; if this parameter is not specified, the default KMS key for EBS is used. If
	// a KmsKeyId is specified, the Encrypted flag must also be set.
	//
	// The KMS key identifier may be provided in any of the following formats:
	//
	//   - Key ID
	//
	//   - Key alias
	//
	//   - ARN using key ID. The ID ARN contains the arn:aws:kms namespace, followed by
	//   the Region of the key, the Amazon Web Services account ID of the key owner, the
	//   key namespace, and then the key ID. For example,
	//   arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.
	//
	//   - ARN using key alias. The alias ARN contains the arn:aws:kms namespace,
	//   followed by the Region of the key, the Amazon Web Services account ID of the key
	//   owner, the alias namespace, and then the key alias. For example,
	//   arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// Amazon Web Services parses KmsKeyId asynchronously, meaning that the action you
	// call may appear to complete even though you provided an invalid identifier. This
	// action will eventually report failure.
	//
	// The specified KMS key must exist in the Region that the snapshot is being
	// copied to.
	//
	// Amazon EBS does not support asymmetric KMS keys.
	KmsKeyId *string

	// The name of the role to use when not using the default role, 'vmimport'.
	RoleName *string

	// The tags to apply to the import snapshot task during creation.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type ImportSnapshotOutput struct {

	// A description of the import snapshot task.
	Description *string

	// The ID of the import snapshot task.
	ImportTaskId *string

	// Information about the import snapshot task.
	SnapshotTaskDetail *types.SnapshotTaskDetail

	// Any tags assigned to the import snapshot task.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpImportSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpImportSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ImportSnapshot"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ImportSnapshot",
	}
}
