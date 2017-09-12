// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package ghctx

// BaseURL returns the base URL to use for all GitHub API calls.
func BaseURL() (string, error) {
	user, err := User()
	if err != nil {
		return "", err
	}
	token, err := Token()
	if err != nil {
		return "", err
	}
	return "https://" + user + ":" + token + "@api.github.com", nil
}
