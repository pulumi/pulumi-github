---
title: Github Provider
meta_desc: Provides an overview on how to configure the Pulumi Github provider.
layout: package
---
## Installation

The github provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/github`](https://www.npmjs.com/package/@pulumi/github)
* Python: [`pulumi-github`](https://pypi.org/project/pulumi-github/)
* Go: [`github.com/pulumi/pulumi-github/sdk/v6/go/github`](https://github.com/pulumi/pulumi-github)
* .NET: [`Pulumi.Github`](https://www.nuget.org/packages/Pulumi.Github)
* Java: [`com.pulumi/github`](https://central.sonatype.com/artifact/com.pulumi/github)
## Overview

The GitHub provider is used to interact with GitHub resources.

The provider allows you to manage your GitHub organization's members and teams easily.
It needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.
## Example Usage



{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as github from "@pulumi/github";

// Add a user to the organization
const membershipForUserX = new github.Membership("membership_for_user_x", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```
```python
import pulumi
import pulumi_github as github

# Add a user to the organization
membership_for_user_x = github.Membership("membership_for_user_x")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Github = Pulumi.Github;

return await Deployment.RunAsync(() =>
{
    // Add a user to the organization
    var membershipForUserX = new Github.Membership("membership_for_user_x");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```
```go
package main

import (
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Add a user to the organization
		_, err := github.NewMembership(ctx, "membership_for_user_x", nil)
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml

```
```yaml
resources:
  # Add a user to the organization
  membershipForUserX:
    type: github:Membership
    name: membership_for_user_x
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.github.Membership;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Add a user to the organization
        var membershipForUserX = new Membership("membershipForUserX");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}



{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as github from "@pulumi/github";

// Add a user to the organization
const membershipForUserX = new github.Membership("membership_for_user_x", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```
```python
import pulumi
import pulumi_github as github

# Add a user to the organization
membership_for_user_x = github.Membership("membership_for_user_x")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Github = Pulumi.Github;

return await Deployment.RunAsync(() =>
{
    // Add a user to the organization
    var membershipForUserX = new Github.Membership("membership_for_user_x");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```
```go
package main

import (
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Add a user to the organization
		_, err := github.NewMembership(ctx, "membership_for_user_x", nil)
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml

```
```yaml
resources:
  # Add a user to the organization
  membershipForUserX:
    type: github:Membership
    name: membership_for_user_x
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.github.Membership;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Add a user to the organization
        var membershipForUserX = new Membership("membershipForUserX");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication

The GitHub provider offers multiple ways to authenticate with GitHub API.
### GitHub CLI

The GitHub provider taps into [GitHub CLI](https://cli.github.com/) authentication, where it picks up the token issued by [`gh auth login`](https://cli.github.com/manual/gh_auth_login) command. It is possible to specify the path to the `gh` executable in the `GH_PATH` environment variable, which is useful for when the GitHub Pulumi provider can not properly determine its the path to GitHub CLI such as in the cygwin terminal.
### OAuth / Personal Access Token

To authenticate using OAuth tokens, ensure that the `token` argument or the `GITHUB_TOKEN` environment variable is set.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    github:token:
        value: 'TODO: var.token'

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    github:token:
        value: 'TODO: var.token'

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    github:token:
        value: 'TODO: var.token'

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    github:token:
        value: 'TODO: var.token'

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    github:token:
        value: 'TODO: var.token'

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    github:token:
        value: 'TODO: var.token'

```

{{% /choosable %}}
{{< /chooser >}}
### GitHub App Installation

To authenticate using a GitHub App installation, ensure that arguments in the `appAuth` block or the `GITHUB_APP_XXX` environment variables are set.
The `owner` parameter required in this situation. Leaving out will throw a `403 "Resource not accessible by integration"` error.

Some API operations may not be available when using a GitHub App installation configuration. For more information, refer to the list of [supported endpoints](https://docs.github.com/en/rest/overview/endpoints-available-for-github-apps).

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    github:owner:
        value: 'TODO: var.github_organization'

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    github:owner:
        value: 'TODO: var.github_organization'

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    github:owner:
        value: 'TODO: var.github_organization'

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    github:owner:
        value: 'TODO: var.github_organization'

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    github:owner:
        value: 'TODO: var.github_organization'

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    github:owner:
        value: 'TODO: var.github_organization'

```

{{% /choosable %}}
{{< /chooser >}}

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    github:owner:
        value: 'TODO: var.github_organization'

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    github:owner:
        value: 'TODO: var.github_organization'

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    github:owner:
        value: 'TODO: var.github_organization'

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    github:owner:
        value: 'TODO: var.github_organization'

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    github:owner:
        value: 'TODO: var.github_organization'

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    github:owner:
        value: 'TODO: var.github_organization'

```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported in the provider configuration:

* `token` - (Optional) A GitHub OAuth / Personal Access Token. When not provided or made available via the `GITHUB_TOKEN` environment variable, the provider can only access resources available anonymously.

* `baseUrl` - (Optional) This is the target GitHub base API endpoint. Providing a value is a requirement when working with GitHub Enterprise. It is optional to provide this value and it can also be sourced from the `GITHUB_BASE_URL` environment variable. The value must end with a slash, for example: `https://pulumitesting-ghe.westus.cloudapp.azure.com/`

* `owner` - (Optional) This is the target GitHub organization or individual user account to manage. For example, `torvalds` and `github` are valid owners. It is optional to provide this value and it can also be sourced from the `GITHUB_OWNER` environment variable. When not provided and a `token` is available, the individual user account owning the `token` will be used. When not provided and no `token` is available, the provider may not function correctly. It is required in case of GitHub App Installation.

* `organization` - (Deprecated) This behaves the same as `owner`, which should be used instead. This value can also be sourced from the `GITHUB_ORGANIZATION` environment variable.

* `appAuth` - (Optional) Configuration block to use GitHub App installation token. When not provided, the provider can only access resources available anonymously.
  * `id` - (Required) This is the ID of the GitHub App. It can sourced from the `GITHUB_APP_ID` environment variable.
  * `installationId` - (Required) This is the ID of the GitHub App installation. It can sourced from the `GITHUB_APP_INSTALLATION_ID` environment variable.
  * `pemFile` - (Required) This is the contents of the GitHub App private key PEM file. It can also be sourced from the `GITHUB_APP_PEM_FILE` environment variable and may use `\n` instead of actual new lines.

* `writeDelayMs` - (Optional) The number of milliseconds to sleep in between write operations in order to satisfy the GitHub API rate limits. Note that requests to the GraphQL API are implemented as `POST` requests under the hood, so this setting affects those calls as well. Defaults to 1000ms or 1 second if not provided.

* `retryDelayMs` - (Optional) Amount of time in milliseconds to sleep in between requests to GitHub API after an error response. Defaults to 1000ms or 1 second if not provided, the maxRetries must be set to greater than zero.

* `readDelayMs` - (Optional) The number of milliseconds to sleep in between non-write operations in order to satisfy the GitHub API rate limits. Defaults to 0ms.

* `retryableErrors` - (Optional) "Allow the provider to retry after receiving an error status code, the maxRetries should be set for this to work. Defaults to [500, 502, 503, 504]

* `maxRetries` - (Optional) Number of times to retry a request after receiving an error status code. Defaults to 3

Note: If you have a PEM file on disk, you can pass it in via `pemFile = file("path/to/file.pem")`.

For backwards compatibility, if more than one of `owner`, `organization`,
`GITHUB_OWNER` and `GITHUB_ORGANIZATION` are set, the first in this
list takes priority.

1. Setting `organization` in the GitHub provider configuration.
2. Setting the `GITHUB_ORGANIZATION` environment variable.
3. Setting the `GITHUB_OWNER` environment variable.
4. Setting `owner` in the GitHub provider configuration.

> It is a bug that `GITHUB_OWNER` takes precedence over `owner`, which may
be fixed in a future major release. For compatibility with future releases,
please set only one of `GITHUB_OWNER` and `owner`.