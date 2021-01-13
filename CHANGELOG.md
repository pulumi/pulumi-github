CHANGELOG
=========

## HEAD (Unreleased)
_(none)_

---

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
