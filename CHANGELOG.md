CHANGELOG
=========

## HEAD (Unreleased)
* Upgrade to v4.6.0 of the GitHub Terraform Provider

---

## 3.3.1 (2021-03-23)
* Upgrade to pulumi-terraform-bridge v2.22.1  
  **Please Note:** This includes a bug fix to the refresh operation regarding arrays

## 3.3.0 (2021-03-16)
* Upgrade to pulumi-terraform-bridge v2.21.0
* Release macOS arm64 binary

## 3.2.1 (2021-03-09)
* Upgrade to v4.5.1 of the GitHub Terraform Provider

## 3.2.0 (2021-02-19)
* Upgrade to v4.5.0 of the GitHub Terraform Provider

## 3.1.1 (2021-02-16)
* Upgrade to pulumi-terraform-bridge v2.19.0  
  **Please Note:** This includes a bug fix that stops mutating resources options in the nodejs provider

## 3.1.0 (2021-02-08)
* Upgrade to v4.4.0 of the GitHub Terraform Provider

## 3.0.2 (2021-02-03)
* Upgrade to v4.3.2 of the GitHub Terraform Provider

## 3.0.1 (2021-02-01)
* Avoid storing config from the environment into the state

## 3.0.0 (2021-01-29)
* Upgrading pulumi-terraform-bridge to v2.18.1  
  *PLEASE NOTE:* An earlier version of this package was tagged as a v3 go module incorrectly. There are no breaking  
  changes in this module, this major version release fixes that error

## 2.5.1 (2021-01-27)
* Upgrade to v4.3.1 of the GitHub Terraform Provider

## 2.5.0 (2021-01-19)
* Upgrade to v4.3.0 of the GitHub Terraform Provider

## 2.4.2 (2021-01-13)
* Upgrade to pulumi-terraform-bridge v2.17.0
* Upgrade to Pulumi v2.17.0

## 2.4.1 (2021-01-05)
* Upgrade to pulumi-terraform-bridge v2.13.2
  * This adds support for import specific examples in documentation

## 2.4.0 (2020-12-08)
* Upgrade to v4.1.0 of the GitHub Terraform Provider

## 2.3.0 (2020-11-26)
* Upgrade to v4.0.1 of the GitHub Terraform Provider
  **Please Note there is a breaking change to the provider:**
  * `github.branchProtection` `branch` has now been replaced by `pattern`

## 2.2.0 (2020-10-26)
* Upgrade to Pulumi v2.12.0 and pulumi-terraform-bridge v2.11.0
* Improving the accuracy of previews leading to a more accurate understanding of what will actually change rather than assuming all output properties will change.  
  ** PLEASE NOTE:**  
  This new preview functionality can be disabled by setting `PULUMI_DISABLE_PROVIDER_PREVIEW` to `1` or `false`.

## 2.1.0 (2020-10-12)
* Upgrade to v3.1.0 of the GitHub Terraform Provider  
* Upgrade to pulumi-terraform-bridge v2.8.0
* Upgrade to Pulumi v2.10.0

## 2.0.0 (2020-09-09)
* Upgrade to v3.0.0 of the GitHub Terraform Provider  
  **Please Note there are a number of changes to the provider configuration:**
  * Provider options `individual` and `anonymous` have been removed  
  * `token` is now optional and absence enables `anonymous` mode
  * Absense of `organization` now enables `individual` mode
  * `owner` is now inferred from `organization`

## 1.5.0 (2020-08-31)
* Upgrade to pulumi-terraform-bridge v2.7.3
* Upgrade to Pulumi v2.9.0, which adds type annotations and input/output classes to Python

## 1.4.2 (2020-07-15)
* Upgrade to v2.9.2 of the GitHub Terraform Provider

## 1.4.1 (2020-07-07)
* Upgrade to v2.9.1 of the GitHub Terraform Provider

## 1.4.0 (2020-06-30)
* Upgrade to v2.9.0 of the GitHub Terraform Provider

## 1.3.3 (2020-06-10)
* Switch to GitHub actions for build

## 1.3.2 (2020-06-10)
* Upgrade to v2.8.1 of the GitHub Terraform Provider

## 1.3.1 (2020-05-28)
* Upgrade to Pulumi v2.3.0
* Upgrade to pulumi-terraform-bridge v2.4.0

## 1.3.0 (2020-05-18)
* Upgrade to v2.8.0 of the GitHub Terraform Provider

## 1.2.1 (2020-05-12)
* Upgrade to pulumi-terraform-bridge v2.3.1

## 1.2.0 (2020-05-04)
* Upgrade to v2.7.0 of the GitHub Terraform Provider

## 1.1.0 (2020-04-28)
* Regenerate datasource examples to be async

## 1.0.0 (2020-04-27)
* Initial creation of the provider against v2.6.1 of the GitHub Terraform Provider
