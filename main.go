package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pyka-project/pulumi-go-pkg/pkg/spacelift/stack"
	"github.com/spacelift-io/pulumi-spacelift/sdk/v2/go/spacelift"
)

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {
		stack, err := stack.Create(ctx, &spacelift.StackArgs{
			Name:                pulumi.String("spacelift-management"),
			Repository:          pulumi.String("https://github.com/pyka-project/spacelift-management"),
			Branch:              pulumi.String("main"),
			WorkerPoolId:        pulumi.String("default"),
			Administrative:      pulumi.Bool(true),
			RunnerImage:         pulumi.String("public.ecr.aws/spacelift/runner-pulumi-golang:latest"),
			Autodeploy:          pulumi.Bool(true),
			Autoretry:           pulumi.Bool(true),
			EnableLocalPreview:  pulumi.Bool(true),
			ProtectFromDeletion: pulumi.Bool(true),
		})

		if err != nil {
			return err
		}
		ctx.Export("stackName", stack.Name)
		ctx.Export("stackRepository", stack.Repository)
		return nil
	})
}
