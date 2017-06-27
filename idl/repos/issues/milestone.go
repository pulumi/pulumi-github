// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package issues

import (
	"github.com/pulumi/lumi/pkg/resource/idl"
)

// You can use milestones to track progress on groups of issues or pull requests in a repository.
type Milestone struct {
	idl.NamedResource
	// The milestone due date.  This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MMSSZ`.
	DueOn string `lumi:"dueOn"`
	// An optional description of the milestone.
	Description *string `lumi:"description,optional"`
	// The state of the milestone.  Either `open` or `closed`.  Default: `open`.
	State *MilestoneState `lumi:"state,optional"`
	// The repo in which this milestone exists; if empty, this is read from configuration.
	Repo *string `lumi:"repo,optional"`
}

type MilestoneState string

const (
	MilestoneOpen   MilestoneState = "open"
	MilestoneClosed MilestoneState = "closed"
)
