// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides details about an EC2 Local Gateway Virtual Interface. More information can be found in the [Outposts User Guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#routing).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_ := "TODO: For expression"
//			return nil
//		})
//	}
//
// ```
func GetLocalGatewayVirtualInterface(ctx *pulumi.Context, args *GetLocalGatewayVirtualInterfaceArgs, opts ...pulumi.InvokeOption) (*GetLocalGatewayVirtualInterfaceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLocalGatewayVirtualInterfaceResult
	err := ctx.Invoke("aws:ec2/getLocalGatewayVirtualInterface:getLocalGatewayVirtualInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalGatewayVirtualInterface.
type GetLocalGatewayVirtualInterfaceArgs struct {
	// One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLocalGatewayVirtualInterfaces.html) for supported filters. Detailed below.
	Filters []GetLocalGatewayVirtualInterfaceFilter `pulumi:"filters"`
	// Identifier of EC2 Local Gateway Virtual Interface.
	Id *string `pulumi:"id"`
	// Key-value map of resource tags, each pair of which must exactly match a pair on the desired local gateway route table.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getLocalGatewayVirtualInterface.
type GetLocalGatewayVirtualInterfaceResult struct {
	Filters []GetLocalGatewayVirtualInterfaceFilter `pulumi:"filters"`
	Id      string                                  `pulumi:"id"`
	// Local address.
	LocalAddress string `pulumi:"localAddress"`
	// Border Gateway Protocol (BGP) Autonomous System Number (ASN) of the EC2 Local Gateway.
	LocalBgpAsn int `pulumi:"localBgpAsn"`
	// Identifier of the EC2 Local Gateway.
	LocalGatewayId                  string   `pulumi:"localGatewayId"`
	LocalGatewayVirtualInterfaceIds []string `pulumi:"localGatewayVirtualInterfaceIds"`
	// Peer address.
	PeerAddress string `pulumi:"peerAddress"`
	// Border Gateway Protocol (BGP) Autonomous System Number (ASN) of the peer.
	PeerBgpAsn int               `pulumi:"peerBgpAsn"`
	Tags       map[string]string `pulumi:"tags"`
	// Virtual Local Area Network.
	Vlan int `pulumi:"vlan"`
}

func GetLocalGatewayVirtualInterfaceOutput(ctx *pulumi.Context, args GetLocalGatewayVirtualInterfaceOutputArgs, opts ...pulumi.InvokeOption) GetLocalGatewayVirtualInterfaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLocalGatewayVirtualInterfaceResult, error) {
			args := v.(GetLocalGatewayVirtualInterfaceArgs)
			r, err := GetLocalGatewayVirtualInterface(ctx, &args, opts...)
			var s GetLocalGatewayVirtualInterfaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLocalGatewayVirtualInterfaceResultOutput)
}

// A collection of arguments for invoking getLocalGatewayVirtualInterface.
type GetLocalGatewayVirtualInterfaceOutputArgs struct {
	// One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLocalGatewayVirtualInterfaces.html) for supported filters. Detailed below.
	Filters GetLocalGatewayVirtualInterfaceFilterArrayInput `pulumi:"filters"`
	// Identifier of EC2 Local Gateway Virtual Interface.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Key-value map of resource tags, each pair of which must exactly match a pair on the desired local gateway route table.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetLocalGatewayVirtualInterfaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalGatewayVirtualInterfaceArgs)(nil)).Elem()
}

// A collection of values returned by getLocalGatewayVirtualInterface.
type GetLocalGatewayVirtualInterfaceResultOutput struct{ *pulumi.OutputState }

func (GetLocalGatewayVirtualInterfaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalGatewayVirtualInterfaceResult)(nil)).Elem()
}

func (o GetLocalGatewayVirtualInterfaceResultOutput) ToGetLocalGatewayVirtualInterfaceResultOutput() GetLocalGatewayVirtualInterfaceResultOutput {
	return o
}

func (o GetLocalGatewayVirtualInterfaceResultOutput) ToGetLocalGatewayVirtualInterfaceResultOutputWithContext(ctx context.Context) GetLocalGatewayVirtualInterfaceResultOutput {
	return o
}

func (o GetLocalGatewayVirtualInterfaceResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetLocalGatewayVirtualInterfaceResult] {
	return pulumix.Output[GetLocalGatewayVirtualInterfaceResult]{
		OutputState: o.OutputState,
	}
}

func (o GetLocalGatewayVirtualInterfaceResultOutput) Filters() GetLocalGatewayVirtualInterfaceFilterArrayOutput {
	return o.ApplyT(func(v GetLocalGatewayVirtualInterfaceResult) []GetLocalGatewayVirtualInterfaceFilter {
		return v.Filters
	}).(GetLocalGatewayVirtualInterfaceFilterArrayOutput)
}

func (o GetLocalGatewayVirtualInterfaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalGatewayVirtualInterfaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Local address.
func (o GetLocalGatewayVirtualInterfaceResultOutput) LocalAddress() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalGatewayVirtualInterfaceResult) string { return v.LocalAddress }).(pulumi.StringOutput)
}

// Border Gateway Protocol (BGP) Autonomous System Number (ASN) of the EC2 Local Gateway.
func (o GetLocalGatewayVirtualInterfaceResultOutput) LocalBgpAsn() pulumi.IntOutput {
	return o.ApplyT(func(v GetLocalGatewayVirtualInterfaceResult) int { return v.LocalBgpAsn }).(pulumi.IntOutput)
}

// Identifier of the EC2 Local Gateway.
func (o GetLocalGatewayVirtualInterfaceResultOutput) LocalGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalGatewayVirtualInterfaceResult) string { return v.LocalGatewayId }).(pulumi.StringOutput)
}

func (o GetLocalGatewayVirtualInterfaceResultOutput) LocalGatewayVirtualInterfaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalGatewayVirtualInterfaceResult) []string { return v.LocalGatewayVirtualInterfaceIds }).(pulumi.StringArrayOutput)
}

// Peer address.
func (o GetLocalGatewayVirtualInterfaceResultOutput) PeerAddress() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalGatewayVirtualInterfaceResult) string { return v.PeerAddress }).(pulumi.StringOutput)
}

// Border Gateway Protocol (BGP) Autonomous System Number (ASN) of the peer.
func (o GetLocalGatewayVirtualInterfaceResultOutput) PeerBgpAsn() pulumi.IntOutput {
	return o.ApplyT(func(v GetLocalGatewayVirtualInterfaceResult) int { return v.PeerBgpAsn }).(pulumi.IntOutput)
}

func (o GetLocalGatewayVirtualInterfaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetLocalGatewayVirtualInterfaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Virtual Local Area Network.
func (o GetLocalGatewayVirtualInterfaceResultOutput) Vlan() pulumi.IntOutput {
	return o.ApplyT(func(v GetLocalGatewayVirtualInterfaceResult) int { return v.Vlan }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLocalGatewayVirtualInterfaceResultOutput{})
}
