package main

import (
	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		repo, err := github.NewRepository(ctx, "demo-repo", &github.RepositoryArgs{
			Description: pulumi.String("Generated from automated test"),
			Visibility:  pulumi.String("private"),
		})
		if err != nil {
			return err
		}

		ctx.Export("repo", repo.Name)
		return nil
	})
}
