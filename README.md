[![Build Status](https://travis-ci.com/pulumi/pulumi-github.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master)](https://travis-ci.com/pulumi/pulumi-github)

# GitHub provider

The GitHub resource provider for Pulumi lets you use GitHub resources in your infrastructure programs. 
To use this package, please [install the Pulumi CLI first](https://pulumi.io/reference/cli/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/github

or `yarn`:

    $ yarn add @pulumi/github

### Python

To use from Python, install using `pip`:

    $ pip install pulumi-github

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-github/sdk/v5

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Github

## Configuration

The following configuration points are available:

- `github:token` - (Optional) This is the GitHub personal access token. It can also be sourced from the `GITHUB_TOKEN`
  environment variable. If anonymous is false, token is required.
- `github:baseUrl` - (Optional) This is the target GitHub base API endpoint. Providing a value is a requirement when
  working with GitHub Enterprise. It is optional to provide this value and it can also be sourced from the `GITHUB_BASE_URL`
  environment variable. The value must end with a slash, and generally includes the API version, for instance
  `https://github.someorg.example/api/v3/`.
- `github:owner` - (Optional) This is the target GitHub organization or individual user account to manage. For example, 
  `torvalds` and `github` are valid owners. It is optional to provide this value and it can also be sourced from the 
  `GITHUB_OWNER` environment variable. When not provided and a token is available, the individual user account owning 
  the token will be used. When not provided and no token is available, the provider may not function correctly.
- `github:organization` - (Deprecated) This behaves the same as owner, which should be used instead. This value can also 
  be sourced from the `GITHUB_ORGANIZATION` environment variable.

## Reference

For further information, please visit [the GitHub provider docs](https://www.pulumi.com/docs/intro/cloud-providers/github)
or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/github).
