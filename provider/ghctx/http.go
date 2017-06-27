// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package ghctx

// HTTPSuccess returns true if the status represents success (i.e., it is a 2xx code).
func HTTPSuccess(status int) bool {
	return status >= 200 && status < 300
}
