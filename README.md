[![Build Status](https://travis-ci.com/pulumi/pulumi-github.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master)](https://travis-ci.com/pulumi/pulumi-github)

# GitHub provider

The GitHub resource provider for Pulumi lets you use GitHub resources in your infrastructure 
programs. To use this package, please [install the Pulumi CLI first](https://pulumi.io/).

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

    $ go get github.com/pulumi/pulumi-github/sdk/v3/go/...

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Github

## Configuration

The following configuration points are available:

- `github:token` - (Optional) This is the GitHub personal access token. It can also be sourced from the `GITHUB_TOKEN`
environment variable. If anonymous is false, token is required.
- `github:organization` - (Optional) This is the target GitHub organization to manage. The account corresponding to the
token will need "owner" privileges for this organization. It can also be sourced from the `GITHUB_ORGANIZATION` environment 
variable. If individual is set to false, organization is required.
- `github:baseUrl` - (Optional) This is the target GitHub base API endpoint. Providing a value is a requirement when
working with GitHub Enterprise. It is optional to provide this value and it can also be sourced from the `GITHUB_BASE_URL`
environment variable. The value must end with a slash, and generally includes the API version, for instance 
`https://github.someorg.example/api/v3/`.
- `github:insecure` - (Optional) Whether server should be accessed without verifying the TLS certificate. As the name
suggests this is insecure and should not be used beyond experiments, accessing local (non-production) GHE instance etc.
There is a number of ways to obtain trusted certificate for free, e.g. from Let's Encrypt. Such trusted certificate does
not require this option to be enabled. Defaults to `false`.
- `github:individual` - (Optional) Run outside an organization. When individual is true, the provider will run outside
the scope of an organization. Defaults to `false`.
- `github:anonymous` - (Optional) Authenticate without a token. When anonymous is true, the provider will not be able to
access resources that require authentication. Setting to true will lead the GitHub provider to work in an anonymous mode
with the corresponding API rate limits. Defaults to `false`.

## Reference

For further information, please visit [the GitHub provider docs](https://www.pulumi.com/docs/intro/cloud-providers/github) or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/github).
