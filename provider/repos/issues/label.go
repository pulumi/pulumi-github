// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package issues

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/pulumi/lumi/pkg/resource"
	"github.com/pulumi/lumi/pkg/resource/provider"
	"github.com/pulumi/lumi/pkg/util/contract"
	"github.com/pulumi/lumi/sdk/go/pkg/lumirpc"
	"golang.org/x/net/context"

	"github.com/pulumi/lumi-github/provider/ghctx"
	"github.com/pulumi/lumi-github/rpc/repos/issues"
)

const LabelToken = issues.LabelToken

// NewLabelProvider creates a provider that handles GitHub label operations.
func NewLabelProvider(host *provider.HostClient) lumirpc.ResourceProviderServer {
	return issues.NewLabelProvider(&lblProvider{host: host})
}

type lblProvider struct {
	host *provider.HostClient
}

func (p *lblProvider) baseURL(repo string) (string, error) {
	base, err := defaultBaseURL(p.host, repo)
	if err != nil {
		return "", err
	}
	return base + "/labels", nil
}

func newLabelID(repo string, name string) resource.ID {
	return resource.ID(repo + ":" + name)
}

func fromLabelID(id resource.ID) (string, string) {
	s := string(id)
	ix := strings.Index(s, ":")
	contract.Assert(ix != -1)
	return s[:ix], s[ix+1:]
}

// Name creates a unique name for this resource.
func (p *lblProvider) Name(ctx context.Context, obj *issues.Label) (string, error) {
	name := obj.Name
	if obj.Repo != nil {
		name = *obj.Repo + ":" + name
	}
	return name, nil
}

// Check validates that the given property bag is valid for a resource of the given type.
func (p *lblProvider) Check(ctx context.Context, obj *issues.Label, property string) error {
	if property == "" {
		return checkCommon(p.host, obj.Repo)
	}
	return nil
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.  (The input ID
// must be blank.)  If this call fails, the resource must not have been created (i.e., it is "transacational").
func (p *lblProvider) Create(ctx context.Context, obj *issues.Label) (resource.ID, error) {
	// Manufacture and POST a payload to the right endpoint.
	body, err := json.Marshal(struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}{
		Name:  obj.Name,
		Color: obj.Color,
	})
	if err != nil {
		return "", err
	}
	repo, err := defaultRepo(p.host, obj.Repo)
	if err != nil {
		return "", err
	}
	url, err := p.baseURL(repo)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	} else if !ghctx.HTTPSuccess(resp.StatusCode) {
		respbody, err := ioutil.ReadAll(resp.Body)
		contract.IgnoreError(err)
		return "", errors.Errorf(
			"GitHub POST did not reply with the expected 201 Created; got %v: %v", resp.Status, string(respbody))
	}
	return newLabelID(repo, obj.Name), err
}

// Get reads the instance state identified by ID, returning a populated resource object, or an error if not found.
func (p *lblProvider) Get(ctx context.Context, id resource.ID) (*issues.Label, error) {
	// Simply perform a GET to fetch the information.
	repo, name := fromLabelID(id)
	url, err := p.baseURL(repo)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url + "/" + name)
	if err != nil {
		return nil, err
	} else if !ghctx.HTTPSuccess(resp.StatusCode) {
		respbody, err := ioutil.ReadAll(resp.Body)
		contract.IgnoreError(err)
		return nil, errors.Errorf(
			"GitHub GET did not reply with the expected 200 OK; got %v: %v", resp.Status, string(respbody))
	}

	// Parse the reply and pluck out the color property.
	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var reply struct {
		Color string `json:"color"`
	}
	if err = json.Unmarshal(respbody, &reply); err != nil {
		return nil, err
	}
	return &issues.Label{
		Name:  name,
		Color: reply.Color,
		Repo:  &repo,
	}, nil
}

// InspectChange checks what impacts a hypothetical update will have on the resource's properties.
func (p *lblProvider) InspectChange(ctx context.Context, id resource.ID,
	old *issues.Label, new *issues.Label, diff *resource.ObjectDiff) ([]string, error) {
	return nil, nil
}

// Update updates an existing resource with new values.  Only those values in the provided property bag are updated
// to new values.  The resource ID is returned and may be different if the resource had to be recreated.
func (p *lblProvider) Update(ctx context.Context, id resource.ID,
	old *issues.Label, new *issues.Label, diff *resource.ObjectDiff) error {
	// Just do a PATCH with the full object state to update the target.
	repo, name := fromLabelID(id)
	body, err := json.Marshal(struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}{
		Name:  name,
		Color: new.Color,
	})
	if err != nil {
		return err
	}
	url, err := p.baseURL(repo)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPatch, url+"/"+name, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	} else if !ghctx.HTTPSuccess(resp.StatusCode) {
		respbody, err := ioutil.ReadAll(resp.Body)
		contract.IgnoreError(err)
		return errors.Errorf(
			"GitHub PATCH did not reply with the expected 200 OK; got %v: %v", resp.Status, string(respbody))
	}
	return nil
}

// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed to still exist.
func (p *lblProvider) Delete(ctx context.Context, id resource.ID, obj issues.Label) error {
	// DELETE the resource using its ID (name).
	repo, name := fromLabelID(id)
	url, err := p.baseURL(repo)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodDelete, url+"/"+name, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	} else if !ghctx.HTTPSuccess(resp.StatusCode) {
		respbody, err := ioutil.ReadAll(resp.Body)
		contract.IgnoreError(err)
		return errors.Errorf(
			"GitHub DELETE did not reply with the expected 204 No Content; got %v: %v", resp.Status, string(respbody))
	}
	return nil
}
