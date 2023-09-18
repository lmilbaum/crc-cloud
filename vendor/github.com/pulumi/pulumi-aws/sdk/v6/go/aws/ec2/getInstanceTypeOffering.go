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

// Information about single EC2 Instance Type Offering.
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
//			_, err := ec2.GetInstanceTypeOffering(ctx, &ec2.GetInstanceTypeOfferingArgs{
//				Filters: []ec2.GetInstanceTypeOfferingFilter{
//					{
//						Name: "instance-type",
//						Values: []string{
//							"t2.micro",
//							"t3.micro",
//						},
//					},
//				},
//				PreferredInstanceTypes: []string{
//					"t3.micro",
//					"t2.micro",
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
func GetInstanceTypeOffering(ctx *pulumi.Context, args *GetInstanceTypeOfferingArgs, opts ...pulumi.InvokeOption) (*GetInstanceTypeOfferingResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInstanceTypeOfferingResult
	err := ctx.Invoke("aws:ec2/getInstanceTypeOffering:getInstanceTypeOffering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceTypeOffering.
type GetInstanceTypeOfferingArgs struct {
	// One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstanceTypeOfferings.html) for supported filters. Detailed below.
	Filters []GetInstanceTypeOfferingFilter `pulumi:"filters"`
	// Location type. Defaults to `region`. Valid values: `availability-zone`, `availability-zone-id`, and `region`.
	LocationType *string `pulumi:"locationType"`
	// Ordered list of preferred EC2 Instance Types. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned.
	PreferredInstanceTypes []string `pulumi:"preferredInstanceTypes"`
}

// A collection of values returned by getInstanceTypeOffering.
type GetInstanceTypeOfferingResult struct {
	Filters []GetInstanceTypeOfferingFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// EC2 Instance Type.
	InstanceType           string   `pulumi:"instanceType"`
	LocationType           *string  `pulumi:"locationType"`
	PreferredInstanceTypes []string `pulumi:"preferredInstanceTypes"`
}

func GetInstanceTypeOfferingOutput(ctx *pulumi.Context, args GetInstanceTypeOfferingOutputArgs, opts ...pulumi.InvokeOption) GetInstanceTypeOfferingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceTypeOfferingResult, error) {
			args := v.(GetInstanceTypeOfferingArgs)
			r, err := GetInstanceTypeOffering(ctx, &args, opts...)
			var s GetInstanceTypeOfferingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceTypeOfferingResultOutput)
}

// A collection of arguments for invoking getInstanceTypeOffering.
type GetInstanceTypeOfferingOutputArgs struct {
	// One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstanceTypeOfferings.html) for supported filters. Detailed below.
	Filters GetInstanceTypeOfferingFilterArrayInput `pulumi:"filters"`
	// Location type. Defaults to `region`. Valid values: `availability-zone`, `availability-zone-id`, and `region`.
	LocationType pulumi.StringPtrInput `pulumi:"locationType"`
	// Ordered list of preferred EC2 Instance Types. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned.
	PreferredInstanceTypes pulumi.StringArrayInput `pulumi:"preferredInstanceTypes"`
}

func (GetInstanceTypeOfferingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceTypeOfferingArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceTypeOffering.
type GetInstanceTypeOfferingResultOutput struct{ *pulumi.OutputState }

func (GetInstanceTypeOfferingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceTypeOfferingResult)(nil)).Elem()
}

func (o GetInstanceTypeOfferingResultOutput) ToGetInstanceTypeOfferingResultOutput() GetInstanceTypeOfferingResultOutput {
	return o
}

func (o GetInstanceTypeOfferingResultOutput) ToGetInstanceTypeOfferingResultOutputWithContext(ctx context.Context) GetInstanceTypeOfferingResultOutput {
	return o
}

func (o GetInstanceTypeOfferingResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetInstanceTypeOfferingResult] {
	return pulumix.Output[GetInstanceTypeOfferingResult]{
		OutputState: o.OutputState,
	}
}

func (o GetInstanceTypeOfferingResultOutput) Filters() GetInstanceTypeOfferingFilterArrayOutput {
	return o.ApplyT(func(v GetInstanceTypeOfferingResult) []GetInstanceTypeOfferingFilter { return v.Filters }).(GetInstanceTypeOfferingFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceTypeOfferingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceTypeOfferingResult) string { return v.Id }).(pulumi.StringOutput)
}

// EC2 Instance Type.
func (o GetInstanceTypeOfferingResultOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceTypeOfferingResult) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o GetInstanceTypeOfferingResultOutput) LocationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceTypeOfferingResult) *string { return v.LocationType }).(pulumi.StringPtrOutput)
}

func (o GetInstanceTypeOfferingResultOutput) PreferredInstanceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceTypeOfferingResult) []string { return v.PreferredInstanceTypes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceTypeOfferingResultOutput{})
}
