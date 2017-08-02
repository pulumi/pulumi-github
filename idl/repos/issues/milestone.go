// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package issues

import (
	"github.com/pulumi/pulumi-fabric/pkg/resource/idl"
)

// You can use milestones to track progress on groups of issues or pull requests in a repository.
type Milestone struct {
	idl.Resource
	// The milestone's title.  This must be unique within a repo.
	Title string `lumi:"title,replaces"`
	// The milestone due date.  This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MMSSZ`.
	DueOn string `lumi:"dueOn"`
	// An optional description of the milestone.
	Description *string `lumi:"description,optional"`
	// The state of the milestone.  Either `open` or `closed`.  Default: `open`.
	State *MilestoneState `lumi:"state,optional"`
	// The repo in which this milestone exists; if empty, this is read from configuration.
	Repo *string `lumi:"repo,optional,replaces"`
}

type MilestoneState string

const (
	MilestoneOpen   MilestoneState = "open"
	MilestoneClosed MilestoneState = "closed"
)
