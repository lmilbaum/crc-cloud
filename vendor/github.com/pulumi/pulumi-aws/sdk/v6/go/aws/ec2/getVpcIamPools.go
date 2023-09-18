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

// `ec2.getVpcIpamPools` provides details about IPAM pools.
//
// This resource can prove useful when IPAM pools are created in another root
// module and you need the pool ids as input variables. For example, pools
// can be shared via RAM and used to create vpcs with CIDRs from that pool.
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
//			_, err := ec2.GetVpcIpamPools(ctx, &ec2.GetVpcIpamPoolsArgs{
//				Filters: []ec2.GetVpcIpamPoolsFilter{
//					{
//						Name: "description",
//						Values: []string{
//							"*test*",
//						},
//					},
//					{
//						Name: "address-family",
//						Values: []string{
//							"ipv4",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: aws.ec2/getvpciampools.getVpcIamPools has been deprecated in favor of aws.ec2/getvpcipampools.getVpcIpamPools
func GetVpcIamPools(ctx *pulumi.Context, args *GetVpcIamPoolsArgs, opts ...pulumi.InvokeOption) (*GetVpcIamPoolsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVpcIamPoolsResult
	err := ctx.Invoke("aws:ec2/getVpcIamPools:getVpcIamPools", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcIamPools.
type GetVpcIamPoolsArgs struct {
	// Custom filter block as described below.
	Filters []GetVpcIamPoolsFilter `pulumi:"filters"`
}

// A collection of values returned by getVpcIamPools.
type GetVpcIamPoolsResult struct {
	Filters []GetVpcIamPoolsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of IPAM pools and their attributes. See below for details
	IpamPools []GetVpcIamPoolsIpamPool `pulumi:"ipamPools"`
}

func GetVpcIamPoolsOutput(ctx *pulumi.Context, args GetVpcIamPoolsOutputArgs, opts ...pulumi.InvokeOption) GetVpcIamPoolsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVpcIamPoolsResult, error) {
			args := v.(GetVpcIamPoolsArgs)
			r, err := GetVpcIamPools(ctx, &args, opts...)
			var s GetVpcIamPoolsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVpcIamPoolsResultOutput)
}

// A collection of arguments for invoking getVpcIamPools.
type GetVpcIamPoolsOutputArgs struct {
	// Custom filter block as described below.
	Filters GetVpcIamPoolsFilterArrayInput `pulumi:"filters"`
}

func (GetVpcIamPoolsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcIamPoolsArgs)(nil)).Elem()
}

// A collection of values returned by getVpcIamPools.
type GetVpcIamPoolsResultOutput struct{ *pulumi.OutputState }

func (GetVpcIamPoolsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcIamPoolsResult)(nil)).Elem()
}

func (o GetVpcIamPoolsResultOutput) ToGetVpcIamPoolsResultOutput() GetVpcIamPoolsResultOutput {
	return o
}

func (o GetVpcIamPoolsResultOutput) ToGetVpcIamPoolsResultOutputWithContext(ctx context.Context) GetVpcIamPoolsResultOutput {
	return o
}

func (o GetVpcIamPoolsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetVpcIamPoolsResult] {
	return pulumix.Output[GetVpcIamPoolsResult]{
		OutputState: o.OutputState,
	}
}

func (o GetVpcIamPoolsResultOutput) Filters() GetVpcIamPoolsFilterArrayOutput {
	return o.ApplyT(func(v GetVpcIamPoolsResult) []GetVpcIamPoolsFilter { return v.Filters }).(GetVpcIamPoolsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVpcIamPoolsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcIamPoolsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of IPAM pools and their attributes. See below for details
func (o GetVpcIamPoolsResultOutput) IpamPools() GetVpcIamPoolsIpamPoolArrayOutput {
	return o.ApplyT(func(v GetVpcIamPoolsResult) []GetVpcIamPoolsIpamPool { return v.IpamPools }).(GetVpcIamPoolsIpamPoolArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpcIamPoolsResultOutput{})
}
