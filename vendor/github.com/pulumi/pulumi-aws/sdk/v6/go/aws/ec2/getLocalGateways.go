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

// Provides information for multiple EC2 Local Gateways, such as their identifiers.
//
// ## Example Usage
//
// The following example retrieves Local Gateways with a resource tag of `service` set to `production`.
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
//			fooLocalGateways, err := ec2.GetLocalGateways(ctx, &ec2.GetLocalGatewaysArgs{
//				Tags: map[string]interface{}{
//					"service": "production",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("foo", fooLocalGateways.Ids)
//			return nil
//		})
//	}
//
// ```
func GetLocalGateways(ctx *pulumi.Context, args *GetLocalGatewaysArgs, opts ...pulumi.InvokeOption) (*GetLocalGatewaysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLocalGatewaysResult
	err := ctx.Invoke("aws:ec2/getLocalGateways:getLocalGateways", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalGateways.
type GetLocalGatewaysArgs struct {
	// Custom filter block as described below.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Filters []GetLocalGatewaysFilter `pulumi:"filters"`
	// Mapping of tags, each pair of which must exactly match
	// a pair on the desired local_gateways.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getLocalGateways.
type GetLocalGatewaysResult struct {
	Filters []GetLocalGatewaysFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Set of all the Local Gateway identifiers
	Ids  []string          `pulumi:"ids"`
	Tags map[string]string `pulumi:"tags"`
}

func GetLocalGatewaysOutput(ctx *pulumi.Context, args GetLocalGatewaysOutputArgs, opts ...pulumi.InvokeOption) GetLocalGatewaysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLocalGatewaysResult, error) {
			args := v.(GetLocalGatewaysArgs)
			r, err := GetLocalGateways(ctx, &args, opts...)
			var s GetLocalGatewaysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLocalGatewaysResultOutput)
}

// A collection of arguments for invoking getLocalGateways.
type GetLocalGatewaysOutputArgs struct {
	// Custom filter block as described below.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Filters GetLocalGatewaysFilterArrayInput `pulumi:"filters"`
	// Mapping of tags, each pair of which must exactly match
	// a pair on the desired local_gateways.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetLocalGatewaysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalGatewaysArgs)(nil)).Elem()
}

// A collection of values returned by getLocalGateways.
type GetLocalGatewaysResultOutput struct{ *pulumi.OutputState }

func (GetLocalGatewaysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalGatewaysResult)(nil)).Elem()
}

func (o GetLocalGatewaysResultOutput) ToGetLocalGatewaysResultOutput() GetLocalGatewaysResultOutput {
	return o
}

func (o GetLocalGatewaysResultOutput) ToGetLocalGatewaysResultOutputWithContext(ctx context.Context) GetLocalGatewaysResultOutput {
	return o
}

func (o GetLocalGatewaysResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetLocalGatewaysResult] {
	return pulumix.Output[GetLocalGatewaysResult]{
		OutputState: o.OutputState,
	}
}

func (o GetLocalGatewaysResultOutput) Filters() GetLocalGatewaysFilterArrayOutput {
	return o.ApplyT(func(v GetLocalGatewaysResult) []GetLocalGatewaysFilter { return v.Filters }).(GetLocalGatewaysFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLocalGatewaysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalGatewaysResult) string { return v.Id }).(pulumi.StringOutput)
}

// Set of all the Local Gateway identifiers
func (o GetLocalGatewaysResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalGatewaysResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetLocalGatewaysResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetLocalGatewaysResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLocalGatewaysResultOutput{})
}
