// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DescribeDirectConnectGatewayAttachmentsRequest
type DescribeDirectConnectGatewayAttachmentsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string `locationName:"directConnectGatewayId" type:"string"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	//
	// If MaxResults is given a value larger than 100, only 100 results are returned.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The token provided in the previous call to retrieve the next page.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The ID of the virtual interface.
	VirtualInterfaceId *string `locationName:"virtualInterfaceId" type:"string"`
}

// String returns the string representation
func (s DescribeDirectConnectGatewayAttachmentsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DescribeDirectConnectGatewayAttachmentsResult
type DescribeDirectConnectGatewayAttachmentsOutput struct {
	_ struct{} `type:"structure"`

	// The attachments.
	DirectConnectGatewayAttachments []DirectConnectGatewayAttachment `locationName:"directConnectGatewayAttachments" type:"list"`

	// The token to retrieve the next page.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeDirectConnectGatewayAttachmentsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDirectConnectGatewayAttachments = "DescribeDirectConnectGatewayAttachments"

// DescribeDirectConnectGatewayAttachmentsRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Lists the attachments between your Direct Connect gateways and virtual interfaces.
// You must specify a Direct Connect gateway, a virtual interface, or both.
// If you specify a Direct Connect gateway, the response contains all virtual
// interfaces attached to the Direct Connect gateway. If you specify a virtual
// interface, the response contains all Direct Connect gateways attached to
// the virtual interface. If you specify both, the response contains the attachment
// between the Direct Connect gateway and the virtual interface.
//
//    // Example sending a request using DescribeDirectConnectGatewayAttachmentsRequest.
//    req := client.DescribeDirectConnectGatewayAttachmentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DescribeDirectConnectGatewayAttachments
func (c *Client) DescribeDirectConnectGatewayAttachmentsRequest(input *DescribeDirectConnectGatewayAttachmentsInput) DescribeDirectConnectGatewayAttachmentsRequest {
	op := &aws.Operation{
		Name:       opDescribeDirectConnectGatewayAttachments,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDirectConnectGatewayAttachmentsInput{}
	}

	req := c.newRequest(op, input, &DescribeDirectConnectGatewayAttachmentsOutput{})
	return DescribeDirectConnectGatewayAttachmentsRequest{Request: req, Input: input, Copy: c.DescribeDirectConnectGatewayAttachmentsRequest}
}

// DescribeDirectConnectGatewayAttachmentsRequest is the request type for the
// DescribeDirectConnectGatewayAttachments API operation.
type DescribeDirectConnectGatewayAttachmentsRequest struct {
	*aws.Request
	Input *DescribeDirectConnectGatewayAttachmentsInput
	Copy  func(*DescribeDirectConnectGatewayAttachmentsInput) DescribeDirectConnectGatewayAttachmentsRequest
}

// Send marshals and sends the DescribeDirectConnectGatewayAttachments API request.
func (r DescribeDirectConnectGatewayAttachmentsRequest) Send(ctx context.Context) (*DescribeDirectConnectGatewayAttachmentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDirectConnectGatewayAttachmentsResponse{
		DescribeDirectConnectGatewayAttachmentsOutput: r.Request.Data.(*DescribeDirectConnectGatewayAttachmentsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDirectConnectGatewayAttachmentsResponse is the response type for the
// DescribeDirectConnectGatewayAttachments API operation.
type DescribeDirectConnectGatewayAttachmentsResponse struct {
	*DescribeDirectConnectGatewayAttachmentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDirectConnectGatewayAttachments request.
func (r *DescribeDirectConnectGatewayAttachmentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}