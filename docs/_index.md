---
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Github Provider
meta_desc: Provides an overview on how to configure the Pulumi Github provider.
layout: package
---

## Installation

The Github provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/github`](https://www.npmjs.com/package/@pulumi/github)
* Python: [`pulumi-github`](https://pypi.org/project/pulumi-github/)
* Go: [`github.com/pulumi/pulumi-github/sdk/v6/go/github`](https://github.com/pulumi/pulumi-github)
* .NET: [`Pulumi.Github`](https://www.nuget.org/packages/Pulumi.Github)
* Java: [`com.pulumi/github`](https://central.sonatype.com/artifact/com.pulumi/github)

## Overview

The GitHub Pulumi provider is used to interact with GitHub resources either as an authenticated client or anonymously.

> You **must** add a `requiredProviders` block to every module that will create resources with this provider. If you do not explicitly require `integrations/github` in a submodule, your Pulumi code run may break in hard-to-troubleshoot ways.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    github:owner:
        value: integrations

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as github from "@pulumi/github";

const example = github.getRepository({
    name: "pulumi-provider-github",
});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    github:owner:
        value: integrations

```

```python
import pulumi
import pulumi_github as github

example = github.get_repository(name="pulumi-provider-github")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    github:owner:
        value: integrations

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Github = Pulumi.Github;

return await Deployment.RunAsync(() =>
{
    var example = Github.GetRepository.Invoke(new()
    {
        Name = "pulumi-provider-github",
    });

});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    github:owner:
        value: integrations

```

```go
package main

import (
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := github.GetRepository(ctx, &github.LookupRepositoryArgs{
			Name: pulumi.StringRef("pulumi-provider-github"),
		}, nil)
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
config:
    github:owner:
        value: integrations

```

```yaml
variables:
  example:
    fn::invoke:
      function: github:getRepository
      arguments:
        name: pulumi-provider-github
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    github:owner:
        value: integrations

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.github.GithubFunctions;
import com.pulumi.github.inputs.GetRepositoryArgs;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        final var example = GithubFunctions.getRepository(GetRepositoryArgs.builder()
            .name("pulumi-provider-github")
            .build());

    }
}
```

{{% /choosable %}}
{{% choosable language hcl %}}
```hcl
pulumi {
  required_providers {
    github = {
      source = "pulumi/github"
    }
  }
}

data "github_getrepository" "example" {
  name = "pulumi-provider-github"
}

```

{{% /choosable %}}
{{< /chooser >}}
## Owner

For backwards compatibility; if more than one of `owner`, `organization`, `GITHUB_OWNER` and `GITHUB_ORGANIZATION` are set the first in this list takes priority.

1. Setting `organization` in the GitHub provider configuration.
2. Setting the `GITHUB_ORGANIZATION` environment variable.
3. Setting the `GITHUB_OWNER` environment variable.
4. Setting `owner` in the GitHub provider configuration.

> It is a bug that `GITHUB_OWNER` takes precedence over `owner`; this will be fixed in a future major release. For compatibility with future releases, please set only one of `GITHUB_OWNER` and `owner`.
## Authentication

The GitHub provider offers multiple ways to authenticate with GitHub API. It uses the following authentication fallback chain (first match wins):

1. **Explicit Token** — `token` argument or `GITHUB_TOKEN` environment variable
2. **GitHub App Installation** — `appAuth` block with `id`, `installationId`, and `pemFile`
3. **GitHub CLI** — Falls back to `gh auth token` if neither token nor appAuth is set
4. **Anonymous** — Read-only access when no credentials are available
### OAuth or Personal Access Token (PAT)

To authenticate using OAuth tokens, ensure that the `token` argument or the `GITHUB_TOKEN` environment variable is set.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    github:token:
        value: 'TODO: var.token'

```
### GitHub App Installation

To authenticate using a GitHub App installation, ensure that arguments in the `appAuth` block or the `GITHUB_APP_XXX` environment variables are set. The `owner` parameter required in this situation. Leaving out will throw a `403 "Resource not accessible by integration"` error.

Some API operations may not be available when using a GitHub App installation configuration. For more information, refer to the list of [supported endpoints](https://docs.github.com/en/rest/overview/endpoints-available-for-github-apps).

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    github:owner:
        value: 'TODO: var.github_organization'

```

> When using environment variables, an empty `appAuth` block is required to allow provider configurations from environment variables to be specified. See: <https://github.com/pulumi/pulumi-plugin-sdk/issues/142>
#### .env

```shell
export GITHUB_APP_ID="12332432" # Required: The GitHub App ID for authentication
export GITHUB_APP_INSTALLATION_ID="12435523" # Required: The GitHub App Installation ID for authentication
export GITHUB_APP_PEM_FILE="..." # Required: Contents of the PEM file for the GitHub App, not the path to the PEM file
```
#### main.tf

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    github:owner:
        value: 'TODO: var.github_organization'

```
### GitHub CLI Authentication

When using the GitHub CLI authentication fallback, you can optionally specify the path to the `gh` executable using the `GH_PATH` environment variable. This is useful when the provider cannot properly determine the path to GitHub CLI, such as in cygwin terminals. If not specified, the provider looks for `gh` in your system PATH.
#### .env

```shell
export GH_PATH="/path/to/gh" # Optional: Specify the path to the GitHub CLI executable if not in system PATH
```
#### main.tf

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```
## Configuration Reference

- `appAuth` (Block List, Max: 1) Authenticate using a GitHub App. (see below for nested schema)
- `baseUrl` (String) The base URL for the GitHub API; this defaults to the GitHub API URL. If you are using GitHub Enterprise Server (GHES) or GitHub Enterprise Cloud with Data Residency (GHEC-DR), this is required. This can also be set by the `GITHUB_BASE_URL` environment variable.
- `cachePath` (String) The path to the cache directory for persisting GitHub API requests between runs; if not set there will be no caching between runs. This can also be set by the `GITHUB_CACHE_PATH` environment variable.
- `insecure` (Boolean, Deprecated) Allow insecure server connections when using SSL.
- `legacyClient` (Boolean) Use the legacy GitHub client implementation; if set to `false`, the new client implementation is used. This can also be set by the `GITHUB_LEGACY_CLIENT` environment variable.
- `maxPerPage` (Number) The maximum number of results per page for paginated API requests; this defaults to `100`. This can also be set by the `GITHUB_MAX_PER_PAGE` environment variable.
- `maxRetries` (Number) The maximum number of retries for failed requests; this defaults to `3`.
- `organization` (String, Deprecated) GitHub organization to manage. This can also be set by the `GITHUB_ORGANIZATION` environment variable.
- `owner` (String) GitHub organization or user account to manage; this is required when authenticating using a GitHub App. If the owner is not provided and a token is provided, the provider will attempt to auto-detect the owner associated with the token. This can also be set by the `GITHUB_OWNER` environment variable.
- `parallelRequests` (Boolean) Allow the provider to make parallel API calls; this is experimental and may cause concurrency and rate limiting issues. This is ignored for the REST API when `legacyClient` is `false` since the new client implementation is designed to safely handle parallel requests.
- `readDelayMs` (Number) The delay in milliseconds between read operations; this defaults to `0`. This can be used to mitigate rate limiting issues when performing a large number of read operations. This is ignored for the REST API when `legacyClient` is `false` since the new client implementation is GitHub rate limit aware.
- `retryDelayMs` (Number) The delay in milliseconds between retry attempts; this defaults to `1000`. This setting only applies when `maxRetries` is greater than `0`.
- `retryableErrors` (List of Number) List of HTTP status codes that should be retried; if not set this uses the provider defaults. This setting only applies when `maxRetries` is greater than `0`. This is ignored for the REST API when `legacyClient` is `false` since the new client implementation handles the retry logic.
- `token` (String) GitHub OAuth or Personal Access Token (PAT) to use for authentication. This can also be set by the `GITHUB_TOKEN` environment variable.
- `writeDelayMs` (Number) The delay in milliseconds between write operations; this defaults to `1000`. This is used to mitigate the GitHub API's abuse rate limits when writing. Note that **ALL** requests to the GraphQL API are implemented as `POST` requests under the hood, so this setting affects those calls as well. This is ignored for the REST API when `legacyClient` is `false` since the new client implementation is GitHub rate limit aware.

<a id="nestedblock--app_auth"></a>
### Nested Schema for `appAuth`

Required:

- `id` (String) The GitHub App's identifier. This can also be set by the `GITHUB_APP_ID` environment variable.
- `installationId` (String) The GitHub App's installation identifier. This can also be set by the `GITHUB_APP_INSTALLATION_ID` environment variable.
- `pemFile` (String, Sensitive) The GitHub App's PEM file content; `\n` can be used for newlines. This can also be set by the `GITHUB_APP_PEM_FILE` environment variable.