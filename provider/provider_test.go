package github_test

import (
	"context"
	"testing"

	"github.com/pulumi/providertest/replay"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	github "github.com/pulumi/pulumi-github/provider/v6"
	"github.com/pulumi/pulumi-github/provider/v6/pkg/version"
)

const repro586 = `[
{
        "method": "/pulumirpc.ResourceProvider/Diff",
        "request": {
            "id": "BPR_kwDOLcEl984C1wZz",
            "urn": "urn:pulumi:dev::ts::github:index/branchProtection:BranchProtection::debug-pulumi-github-protection",
            "olds": {
                "__meta": "{\"schema_version\":\"1\"}",
                "allowsDeletions": false,
                "allowsForcePushes": false,
                "blocksCreations": false,
                "enforceAdmins": true,
                "forcePushBypassers": [],
                "id": "BPR_kwDOLcEl984C1wZz",
                "lockBranch": false,
                "pattern": "main",
                "pushRestrictions": [
                    "/iwahbe"
                ],
                "repositoryId": "debug-pulumi-github",
                "requireConversationResolution": false,
                "requireSignedCommits": false,
                "requiredLinearHistory": false,
                "requiredPullRequestReviews": [
                    {
                        "dismissStaleReviews": true,
                        "dismissalRestrictions": [],
                        "pullRequestBypassers": [],
                        "requireCodeOwnerReviews": false,
                        "requireLastPushApproval": false,
                        "requiredApprovingReviewCount": 1,
                        "restrictDismissals": false
                    }
                ],
                "requiredStatusChecks": []
            },
            "news": {
                "__defaults": [
                    "allowsForcePushes",
                    "lockBranch",
                    "requireConversationResolution",
                    "requireSignedCommits",
                    "requiredLinearHistory"
                ],
                "allowsDeletions": false,
                "allowsForcePushes": false,
                "enforceAdmins": true,
                "lockBranch": false,
                "pattern": "main",
                "repositoryId": "debug-pulumi-github",
                "requireConversationResolution": false,
                "requireSignedCommits": false,
                "requiredLinearHistory": false,
                "requiredPullRequestReviews": [
                    {
                        "__defaults": [
                            "requireLastPushApproval"
                        ],
                        "dismissStaleReviews": true,
                        "requireLastPushApproval": false,
                        "requiredApprovingReviewCount": 1
                    }
                ],
                "restrictPushes": [
                    {
                        "__defaults": [
                            "blocksCreations"
                        ],
                        "blocksCreations": true,
                        "pushAllowances": [
                            "/iwahbe"
                        ]
                    }
                ]
            },
            "oldInputs": {
                "__defaults": [
                    "allowsForcePushes",
                    "blocksCreations",
                    "lockBranch",
                    "requireConversationResolution",
                    "requireSignedCommits",
                    "requiredLinearHistory"
                ],
                "allowsDeletions": false,
                "allowsForcePushes": false,
                "blocksCreations": false,
                "enforceAdmins": true,
                "lockBranch": false,
                "pattern": "main",
                "pushRestrictions": [
                    "/iwahbe"
                ],
                "repositoryId": "debug-pulumi-github",
                "requireConversationResolution": false,
                "requireSignedCommits": false,
                "requiredLinearHistory": false,
                "requiredPullRequestReviews": [
                    {
                        "__defaults": [
                            "requireLastPushApproval"
                        ],
                        "dismissStaleReviews": true,
                        "requireLastPushApproval": false,
                        "requiredApprovingReviewCount": 1
                    }
                ]
            }
        },
        "response": {
            "stables": [
                "repositoryId"
            ],
            "changes": "DIFF_SOME",
            "diffs": "*",
            "detailedDiff": {
                "restrictPushes": {
                }
            },
            "hasDetailedDiff": true
        },
        "metadata": {
            "kind": "resource",
            "mode": "client",
            "name": "github"
        }
    }
]
`

func TestRepros(t *testing.T) {
	t.Run("586", func(t *testing.T) {
		replay.ReplaySequence(t, server(t), repro586)
	})
}

func init() {
	version.Version = "6.0.0"
}

func server(*testing.T) pulumirpc.ResourceProviderServer {
	p := github.Provider()
	return tfbridge.NewProvider(context.Background(),
		nil, "github", version.Version, p.P, p, []byte("{}"))
}
