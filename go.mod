module github.com/pulumi/pulumi-github

go 1.12

replace (
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
)

require (
	github.com/hashicorp/terraform v0.12.0-rc1.0.20190509225429-28b2383eacae
	github.com/pkg/errors v0.8.0
	github.com/pulumi/pulumi v0.17.6-0.20190410045519-ef5e148a73c0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190622203420-512f93e9d8c3
	github.com/stretchr/testify v1.3.1-0.20190311161405-34c6fa2dc709
	github.com/terraform-providers/terraform-provider-github v0.0.0-20190515134151-0542d36ee64d
)
