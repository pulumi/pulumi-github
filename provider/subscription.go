// Licensed to Pulumi Corporation ("Pulumi") under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// Pulumi licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/pulumi/lumi/pkg/resource"
	"github.com/pulumi/lumi/sdk/go/pkg/lumirpc"
	"golang.org/x/net/context"

	"github.com/pulumi/github-lumipack/rpc/webhooks"
)

const SubscriptionToken = webhooks.SubscriptionToken

// NewSubscriptionProvider creates a provider that handles EC2 security group operations.
func NewSubscriptionProvider() lumirpc.ResourceProviderServer {
	return webhooks.NewSubscriptionProvider(&subProvider{})
}

type subProvider struct {
}

func (p *subProvider) baseWebHooksURL() string {
	// TODO: fetch these from real configuration.
	owner := "joeduffy"
	repo := "pulumi/lumi"
	return "https://" + owner + "@api.github.com/repos/" + repo + "/hooks"
}

// Check validates that the given property bag is valid for a resource of the given type.
func (p *subProvider) Check(ctx context.Context, obj *webhooks.Subscription, property string) error {
	return nil
}

type subBody struct {
	Name   string    `json:"name"`
	Active *bool     `json:"active,omitempty"`
	Events *[]string `json:"events,omitempty"`
	Config subConfig `json:"config"`
}

type subConfig struct {
	URL         string  `json:"url"`
	ContentType *string `json:"content_type,omitempty"`
	Secret      *string `json:"secret,omitempty"`
	InsecureSSL *string `json:"insecure_ssl,omitempty"`
}

// isSuccessHTTPCode returns true if the status represents success.
func isSuccessHTTPCode(status int) bool {
	return status >= 200 && status < 300
}

// subToSubBody translates a subscription from RPC into a JSON serializable structure.
func subToSubBody(obj *webhooks.Subscription) subBody {
	var contentType *string
	if obj.Config.ContentType != nil {
		ct := string(*obj.Config.ContentType)
		contentType = &ct
	}
	var insecureSSL *string
	if obj.Config.InsecureSSL != nil {
		var ins string
		if *obj.Config.InsecureSSL {
			ins = "1"
		} else {
			ins = "0"
		}
		insecureSSL = &ins
	}
	return subBody{
		Name: obj.Service,
		Config: subConfig{
			URL:         obj.Config.URL,
			ContentType: contentType,
			Secret:      obj.Config.Secret,
			InsecureSSL: insecureSSL,
		},
		Events: obj.Events,
		Active: obj.Active,
	}
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.  (The input ID
// must be blank.)  If this call fails, the resource must not have been created (i.e., it is "transacational").
func (p *subProvider) Create(ctx context.Context, obj *webhooks.Subscription) (resource.ID, error) {
	// Manufacture and POST a payload to the right endpoint.
	sub := subToSubBody(obj)
	body, err := json.Marshal(sub)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(p.baseWebHooksURL(), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	} else if !isSuccessHTTPCode(resp.StatusCode) {
		return "", errors.Errorf("GitHub POST did not reply with the expected 201 Created; got %v", resp.Status)
	}

	// Parse the reply and pluck out the webhook ID.
	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	reply := struct {
		ID float64 `json:"id"`
	}{}
	if err = json.Unmarshal(respbody, &reply); err != nil {
		return "", err
	}
	return resource.ID(strconv.FormatInt(int64(reply.ID), 10)), err
}

// Get reads the instance state identified by ID, returning a populated resource object, or an error if not found.
func (p *subProvider) Get(ctx context.Context, id resource.ID) (*webhooks.Subscription, error) {
	// Simply perform a GET to fetch the information.
	resp, err := http.Get(p.baseWebHooksURL() + "/" + string(id))
	if err != nil {
		return nil, err
	} else if !isSuccessHTTPCode(resp.StatusCode) {
		return nil, errors.Errorf("GitHub GET did not reply with the expected 200 OK; got %v", resp.Status)
	}

	// Parse the reply and pluck out the webhook ID.
	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var reply subBody
	if err = json.Unmarshal(respbody, &reply); err != nil {
		return nil, err
	}
	var contentType *webhooks.ContentType
	if reply.Config.ContentType != nil {
		ct := webhooks.ContentType(*reply.Config.ContentType)
		contentType = &ct
	}
	return &webhooks.Subscription{
		Service: reply.Name,
		Events:  reply.Events,
		Active:  reply.Active,
		Config: webhooks.Config{
			URL:         reply.Config.URL,
			ContentType: contentType,
		},
	}, nil
}

// InspectChange checks what impacts a hypothetical update will have on the resource's properties.
func (p *subProvider) InspectChange(ctx context.Context, id resource.ID,
	old *webhooks.Subscription, new *webhooks.Subscription, diff *resource.ObjectDiff) ([]string, error) {
	return nil, nil
}

// Update updates an existing resource with new values.  Only those values in the provided property bag are updated
// to new values.  The resource ID is returned and may be different if the resource had to be recreated.
func (p *subProvider) Update(ctx context.Context, id resource.ID,
	old *webhooks.Subscription, new *webhooks.Subscription, diff *resource.ObjectDiff) error {
	// Just do a PATCH with the full object state to update the target.
	sub := subToSubBody(new)
	body, err := json.Marshal(sub)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPatch, p.baseWebHooksURL()+"/"+string(id), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	} else if !isSuccessHTTPCode(resp.StatusCode) {
		return errors.Errorf("GitHub PATCH did not reply with the expected 200 OK; got %v", resp.Status)
	}
	return nil
}

// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed to still exist.
func (p *subProvider) Delete(ctx context.Context, id resource.ID) error {
	// DELETE the resource using its ID.
	resp, err := http.Get(p.baseWebHooksURL() + "/" + string(id))
	if err != nil {
		return err
	} else if !isSuccessHTTPCode(resp.StatusCode) {
		return errors.Errorf("GitHub DELETE did not reply with the expected 204 No Content; got %v", resp.Status)
	}
	return nil
}
