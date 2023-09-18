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

// Provides information for multiple EC2 Local Gateway Route Tables, such as their identifiers.
//
// ## Example Usage
//
// The following shows outputting all Local Gateway Route Table Ids.
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
//			fooLocalGatewayRouteTables, err := ec2.GetLocalGatewayRouteTables(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("foo", fooLocalGatewayRouteTables.Ids)
//			return nil
//		})
//	}
//
// ```
func GetLocalGatewayRouteTables(ctx *pulumi.Context, args *GetLocalGatewayRouteTablesArgs, opts ...pulumi.InvokeOption) (*GetLocalGatewayRouteTablesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLocalGatewayRouteTablesResult
	err := ctx.Invoke("aws:ec2/getLocalGatewayRouteTables:getLocalGatewayRouteTables", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalGatewayRouteTables.
type GetLocalGatewayRouteTablesArgs struct {
	// Custom filter block as described below.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Filters []GetLocalGatewayRouteTablesFilter `pulumi:"filters"`
	// Mapping of tags, each pair of which must exactly match
	// a pair on the desired local gateway route table.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getLocalGatewayRouteTables.
type GetLocalGatewayRouteTablesResult struct {
	Filters []GetLocalGatewayRouteTablesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Set of Local Gateway Route Table identifiers
	Ids  []string          `pulumi:"ids"`
	Tags map[string]string `pulumi:"tags"`
}

func GetLocalGatewayRouteTablesOutput(ctx *pulumi.Context, args GetLocalGatewayRouteTablesOutputArgs, opts ...pulumi.InvokeOption) GetLocalGatewayRouteTablesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLocalGatewayRouteTablesResult, error) {
			args := v.(GetLocalGatewayRouteTablesArgs)
			r, err := GetLocalGatewayRouteTables(ctx, &args, opts...)
			var s GetLocalGatewayRouteTablesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLocalGatewayRouteTablesResultOutput)
}

// A collection of arguments for invoking getLocalGatewayRouteTables.
type GetLocalGatewayRouteTablesOutputArgs struct {
	// Custom filter block as described below.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Filters GetLocalGatewayRouteTablesFilterArrayInput `pulumi:"filters"`
	// Mapping of tags, each pair of which must exactly match
	// a pair on the desired local gateway route table.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetLocalGatewayRouteTablesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalGatewayRouteTablesArgs)(nil)).Elem()
}

// A collection of values returned by getLocalGatewayRouteTables.
type GetLocalGatewayRouteTablesResultOutput struct{ *pulumi.OutputState }

func (GetLocalGatewayRouteTablesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalGatewayRouteTablesResult)(nil)).Elem()
}

func (o GetLocalGatewayRouteTablesResultOutput) ToGetLocalGatewayRouteTablesResultOutput() GetLocalGatewayRouteTablesResultOutput {
	return o
}

func (o GetLocalGatewayRouteTablesResultOutput) ToGetLocalGatewayRouteTablesResultOutputWithContext(ctx context.Context) GetLocalGatewayRouteTablesResultOutput {
	return o
}

func (o GetLocalGatewayRouteTablesResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetLocalGatewayRouteTablesResult] {
	return pulumix.Output[GetLocalGatewayRouteTablesResult]{
		OutputState: o.OutputState,
	}
}

func (o GetLocalGatewayRouteTablesResultOutput) Filters() GetLocalGatewayRouteTablesFilterArrayOutput {
	return o.ApplyT(func(v GetLocalGatewayRouteTablesResult) []GetLocalGatewayRouteTablesFilter { return v.Filters }).(GetLocalGatewayRouteTablesFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLocalGatewayRouteTablesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalGatewayRouteTablesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Set of Local Gateway Route Table identifiers
func (o GetLocalGatewayRouteTablesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalGatewayRouteTablesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetLocalGatewayRouteTablesResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetLocalGatewayRouteTablesResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLocalGatewayRouteTablesResultOutput{})
}
