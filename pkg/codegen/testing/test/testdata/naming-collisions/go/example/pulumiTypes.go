// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Object struct {
	Bar *string `pulumi:"bar"`
}

// ObjectInput is an input type that accepts ObjectArgs and ObjectOutput values.
// You can construct a concrete instance of `ObjectInput` via:
//
//          ObjectArgs{...}
type ObjectInput interface {
	pulumi.Input

	ToObjectOutput() ObjectOutput
	ToObjectOutputWithContext(context.Context) ObjectOutput
}

type ObjectArgs struct {
	Bar pulumi.StringPtrInput `pulumi:"bar"`
}

func (ObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Object)(nil)).Elem()
}

func (i ObjectArgs) ToObjectOutput() ObjectOutput {
	return i.ToObjectOutputWithContext(context.Background())
}

func (i ObjectArgs) ToObjectOutputWithContext(ctx context.Context) ObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectOutput)
}

type ObjectOutput struct{ *pulumi.OutputState }

func (ObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Object)(nil)).Elem()
}

func (o ObjectOutput) ToObjectOutput() ObjectOutput {
	return o
}

func (o ObjectOutput) ToObjectOutputWithContext(ctx context.Context) ObjectOutput {
	return o
}

func (o ObjectOutput) Bar() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Object) *string { return v.Bar }).(pulumi.StringPtrOutput)
}

type ObjectInputType struct {
	Bar *string `pulumi:"bar"`
}

// ObjectInputTypeInput is an input type that accepts ObjectInputTypeArgs and ObjectInputTypeOutput values.
// You can construct a concrete instance of `ObjectInputTypeInput` via:
//
//          ObjectInputTypeArgs{...}
type ObjectInputTypeInput interface {
	pulumi.Input

	ToObjectInputTypeOutput() ObjectInputTypeOutput
	ToObjectInputTypeOutputWithContext(context.Context) ObjectInputTypeOutput
}

type ObjectInputTypeArgs struct {
	Bar pulumi.StringPtrInput `pulumi:"bar"`
}

func (ObjectInputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectInputType)(nil)).Elem()
}

func (i ObjectInputTypeArgs) ToObjectInputTypeOutput() ObjectInputTypeOutput {
	return i.ToObjectInputTypeOutputWithContext(context.Background())
}

func (i ObjectInputTypeArgs) ToObjectInputTypeOutputWithContext(ctx context.Context) ObjectInputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectInputTypeOutput)
}

type ObjectInputTypeOutput struct{ *pulumi.OutputState }

func (ObjectInputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectInputType)(nil)).Elem()
}

func (o ObjectInputTypeOutput) ToObjectInputTypeOutput() ObjectInputTypeOutput {
	return o
}

func (o ObjectInputTypeOutput) ToObjectInputTypeOutputWithContext(ctx context.Context) ObjectInputTypeOutput {
	return o
}

func (o ObjectInputTypeOutput) Bar() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectInputType) *string { return v.Bar }).(pulumi.StringPtrOutput)
}

type ResourceType struct {
	Name *string `pulumi:"name"`
}

// ResourceTypeInput is an input type that accepts ResourceTypeArgs and ResourceTypeOutput values.
// You can construct a concrete instance of `ResourceTypeInput` via:
//
//          ResourceTypeArgs{...}
type ResourceTypeInput interface {
	pulumi.Input

	ToResourceTypeOutput() ResourceTypeOutput
	ToResourceTypeOutputWithContext(context.Context) ResourceTypeOutput
}

type ResourceTypeArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ResourceTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceType)(nil)).Elem()
}

func (i ResourceTypeArgs) ToResourceTypeOutput() ResourceTypeOutput {
	return i.ToResourceTypeOutputWithContext(context.Background())
}

func (i ResourceTypeArgs) ToResourceTypeOutputWithContext(ctx context.Context) ResourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceTypeOutput)
}

type ResourceTypeOutput struct{ *pulumi.OutputState }

func (ResourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceType)(nil)).Elem()
}

func (o ResourceTypeOutput) ToResourceTypeOutput() ResourceTypeOutput {
	return o
}

func (o ResourceTypeOutput) ToResourceTypeOutputWithContext(ctx context.Context) ResourceTypeOutput {
	return o
}

func (o ResourceTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectInput)(nil)).Elem(), ObjectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectInputTypeInput)(nil)).Elem(), ObjectInputTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceTypeInput)(nil)).Elem(), ResourceTypeArgs{})
	pulumi.RegisterOutputType(ObjectOutput{})
	pulumi.RegisterOutputType(ObjectInputTypeOutput{})
	pulumi.RegisterOutputType(ResourceTypeOutput{})
}
