// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a resource to manage whether default EBS encryption is enabled for your AWS account in the current AWS region. To manage the default KMS key for the region, see the `ebs.DefaultKmsKey` resource.
//
// > **NOTE:** Removing this resource disables default EBS encryption.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ebs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ebs.NewEncryptionByDefault(ctx, "example", &ebs.EncryptionByDefaultArgs{
//				Enabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Using `pulumi import`, import the default EBS encryption state. For example:
//
// ```sh
//
//	$ pulumi import aws:ebs/encryptionByDefault:EncryptionByDefault example default
//
// ```
type EncryptionByDefault struct {
	pulumi.CustomResourceState

	// Whether or not default EBS encryption is enabled. Valid values are `true` or `false`. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
}

// NewEncryptionByDefault registers a new resource with the given unique name, arguments, and options.
func NewEncryptionByDefault(ctx *pulumi.Context,
	name string, args *EncryptionByDefaultArgs, opts ...pulumi.ResourceOption) (*EncryptionByDefault, error) {
	if args == nil {
		args = &EncryptionByDefaultArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EncryptionByDefault
	err := ctx.RegisterResource("aws:ebs/encryptionByDefault:EncryptionByDefault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEncryptionByDefault gets an existing EncryptionByDefault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEncryptionByDefault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EncryptionByDefaultState, opts ...pulumi.ResourceOption) (*EncryptionByDefault, error) {
	var resource EncryptionByDefault
	err := ctx.ReadResource("aws:ebs/encryptionByDefault:EncryptionByDefault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EncryptionByDefault resources.
type encryptionByDefaultState struct {
	// Whether or not default EBS encryption is enabled. Valid values are `true` or `false`. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
}

type EncryptionByDefaultState struct {
	// Whether or not default EBS encryption is enabled. Valid values are `true` or `false`. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
}

func (EncryptionByDefaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*encryptionByDefaultState)(nil)).Elem()
}

type encryptionByDefaultArgs struct {
	// Whether or not default EBS encryption is enabled. Valid values are `true` or `false`. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
}

// The set of arguments for constructing a EncryptionByDefault resource.
type EncryptionByDefaultArgs struct {
	// Whether or not default EBS encryption is enabled. Valid values are `true` or `false`. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
}

func (EncryptionByDefaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*encryptionByDefaultArgs)(nil)).Elem()
}

type EncryptionByDefaultInput interface {
	pulumi.Input

	ToEncryptionByDefaultOutput() EncryptionByDefaultOutput
	ToEncryptionByDefaultOutputWithContext(ctx context.Context) EncryptionByDefaultOutput
}

func (*EncryptionByDefault) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionByDefault)(nil)).Elem()
}

func (i *EncryptionByDefault) ToEncryptionByDefaultOutput() EncryptionByDefaultOutput {
	return i.ToEncryptionByDefaultOutputWithContext(context.Background())
}

func (i *EncryptionByDefault) ToEncryptionByDefaultOutputWithContext(ctx context.Context) EncryptionByDefaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionByDefaultOutput)
}

func (i *EncryptionByDefault) ToOutput(ctx context.Context) pulumix.Output[*EncryptionByDefault] {
	return pulumix.Output[*EncryptionByDefault]{
		OutputState: i.ToEncryptionByDefaultOutputWithContext(ctx).OutputState,
	}
}

// EncryptionByDefaultArrayInput is an input type that accepts EncryptionByDefaultArray and EncryptionByDefaultArrayOutput values.
// You can construct a concrete instance of `EncryptionByDefaultArrayInput` via:
//
//	EncryptionByDefaultArray{ EncryptionByDefaultArgs{...} }
type EncryptionByDefaultArrayInput interface {
	pulumi.Input

	ToEncryptionByDefaultArrayOutput() EncryptionByDefaultArrayOutput
	ToEncryptionByDefaultArrayOutputWithContext(context.Context) EncryptionByDefaultArrayOutput
}

type EncryptionByDefaultArray []EncryptionByDefaultInput

func (EncryptionByDefaultArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EncryptionByDefault)(nil)).Elem()
}

func (i EncryptionByDefaultArray) ToEncryptionByDefaultArrayOutput() EncryptionByDefaultArrayOutput {
	return i.ToEncryptionByDefaultArrayOutputWithContext(context.Background())
}

func (i EncryptionByDefaultArray) ToEncryptionByDefaultArrayOutputWithContext(ctx context.Context) EncryptionByDefaultArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionByDefaultArrayOutput)
}

func (i EncryptionByDefaultArray) ToOutput(ctx context.Context) pulumix.Output[[]*EncryptionByDefault] {
	return pulumix.Output[[]*EncryptionByDefault]{
		OutputState: i.ToEncryptionByDefaultArrayOutputWithContext(ctx).OutputState,
	}
}

// EncryptionByDefaultMapInput is an input type that accepts EncryptionByDefaultMap and EncryptionByDefaultMapOutput values.
// You can construct a concrete instance of `EncryptionByDefaultMapInput` via:
//
//	EncryptionByDefaultMap{ "key": EncryptionByDefaultArgs{...} }
type EncryptionByDefaultMapInput interface {
	pulumi.Input

	ToEncryptionByDefaultMapOutput() EncryptionByDefaultMapOutput
	ToEncryptionByDefaultMapOutputWithContext(context.Context) EncryptionByDefaultMapOutput
}

type EncryptionByDefaultMap map[string]EncryptionByDefaultInput

func (EncryptionByDefaultMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EncryptionByDefault)(nil)).Elem()
}

func (i EncryptionByDefaultMap) ToEncryptionByDefaultMapOutput() EncryptionByDefaultMapOutput {
	return i.ToEncryptionByDefaultMapOutputWithContext(context.Background())
}

func (i EncryptionByDefaultMap) ToEncryptionByDefaultMapOutputWithContext(ctx context.Context) EncryptionByDefaultMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionByDefaultMapOutput)
}

func (i EncryptionByDefaultMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*EncryptionByDefault] {
	return pulumix.Output[map[string]*EncryptionByDefault]{
		OutputState: i.ToEncryptionByDefaultMapOutputWithContext(ctx).OutputState,
	}
}

type EncryptionByDefaultOutput struct{ *pulumi.OutputState }

func (EncryptionByDefaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionByDefault)(nil)).Elem()
}

func (o EncryptionByDefaultOutput) ToEncryptionByDefaultOutput() EncryptionByDefaultOutput {
	return o
}

func (o EncryptionByDefaultOutput) ToEncryptionByDefaultOutputWithContext(ctx context.Context) EncryptionByDefaultOutput {
	return o
}

func (o EncryptionByDefaultOutput) ToOutput(ctx context.Context) pulumix.Output[*EncryptionByDefault] {
	return pulumix.Output[*EncryptionByDefault]{
		OutputState: o.OutputState,
	}
}

// Whether or not default EBS encryption is enabled. Valid values are `true` or `false`. Defaults to `true`.
func (o EncryptionByDefaultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionByDefault) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type EncryptionByDefaultArrayOutput struct{ *pulumi.OutputState }

func (EncryptionByDefaultArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EncryptionByDefault)(nil)).Elem()
}

func (o EncryptionByDefaultArrayOutput) ToEncryptionByDefaultArrayOutput() EncryptionByDefaultArrayOutput {
	return o
}

func (o EncryptionByDefaultArrayOutput) ToEncryptionByDefaultArrayOutputWithContext(ctx context.Context) EncryptionByDefaultArrayOutput {
	return o
}

func (o EncryptionByDefaultArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*EncryptionByDefault] {
	return pulumix.Output[[]*EncryptionByDefault]{
		OutputState: o.OutputState,
	}
}

func (o EncryptionByDefaultArrayOutput) Index(i pulumi.IntInput) EncryptionByDefaultOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EncryptionByDefault {
		return vs[0].([]*EncryptionByDefault)[vs[1].(int)]
	}).(EncryptionByDefaultOutput)
}

type EncryptionByDefaultMapOutput struct{ *pulumi.OutputState }

func (EncryptionByDefaultMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EncryptionByDefault)(nil)).Elem()
}

func (o EncryptionByDefaultMapOutput) ToEncryptionByDefaultMapOutput() EncryptionByDefaultMapOutput {
	return o
}

func (o EncryptionByDefaultMapOutput) ToEncryptionByDefaultMapOutputWithContext(ctx context.Context) EncryptionByDefaultMapOutput {
	return o
}

func (o EncryptionByDefaultMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*EncryptionByDefault] {
	return pulumix.Output[map[string]*EncryptionByDefault]{
		OutputState: o.OutputState,
	}
}

func (o EncryptionByDefaultMapOutput) MapIndex(k pulumi.StringInput) EncryptionByDefaultOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EncryptionByDefault {
		return vs[0].(map[string]*EncryptionByDefault)[vs[1].(string)]
	}).(EncryptionByDefaultOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionByDefaultInput)(nil)).Elem(), &EncryptionByDefault{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionByDefaultArrayInput)(nil)).Elem(), EncryptionByDefaultArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionByDefaultMapInput)(nil)).Elem(), EncryptionByDefaultMap{})
	pulumi.RegisterOutputType(EncryptionByDefaultOutput{})
	pulumi.RegisterOutputType(EncryptionByDefaultArrayOutput{})
	pulumi.RegisterOutputType(EncryptionByDefaultMapOutput{})
}
