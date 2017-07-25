// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package issues

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
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

const MilestoneToken = issues.MilestoneToken

// NewMilestoneProvider creates a provider that handles GitHub label operations.
func NewMilestoneProvider(host *provider.HostClient) lumirpc.ResourceProviderServer {
	return issues.NewMilestoneProvider(&msProvider{host: host})
}

type msProvider struct {
	host *provider.HostClient
}

func (p *msProvider) baseURL(repo string) (string, error) {
	base, err := defaultBaseURL(p.host, repo)
	if err != nil {
		return "", err
	}
	return base + "/milestones", nil
}

func newMilestoneID(repo string, id int) resource.ID {
	return resource.ID(repo + ":" + strconv.Itoa(id))
}

func fromMilestoneID(id resource.ID) (string, int) {
	s := string(id)
	ix := strings.Index(s, ":")
	contract.Assert(ix != -1)
	n, err := strconv.Atoi(s[ix+1:])
	contract.Assert(err == nil)
	return s[:ix], n
}

type milestone struct {
	Title       string  `json:"title"`
	DueOn       string  `json:"due_on"`
	Description *string `json:"description,omitempty"`
	State       *string `json:"state,omitempty"`
}

func stateToString(s *issues.MilestoneState) *string {
	if s == nil {
		return nil
	}
	st := string(*s)
	return &st
}

func stringToState(s *string) *issues.MilestoneState {
	if s == nil {
		return nil
	}
	st := issues.MilestoneState(*s)
	return &st
}

// Name creates a unique name for this resource.
func (p *msProvider) Name(ctx context.Context, obj *issues.Milestone) (string, error) {
	name := obj.Title
	if obj.Repo != nil {
		name = *obj.Repo + ":" + name
	}
	return name, nil
}

// Check validates that the given property bag is valid for a resource of the given type.
func (p *msProvider) Check(ctx context.Context, obj *issues.Milestone, property string) error {
	if property == "" {
		return checkCommon(p.host, obj.Repo)
	}
	return nil
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.  (The input ID
// must be blank.)  If this call fails, the resource must not have been created (i.e., it is "transacational").
func (p *msProvider) Create(ctx context.Context, obj *issues.Milestone) (resource.ID, error) {
	// Manufacture and POST a payload to the right endpoint.
	body, err := json.Marshal(milestone{
		Title:       obj.Title,
		DueOn:       obj.DueOn,
		Description: obj.Description,
		State:       stateToString(obj.State),
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

	// Parse the reply and pluck out the milestone number.
	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var reply struct {
		Number int `json:"number"`
	}
	if err = json.Unmarshal(respbody, &reply); err != nil {
		return "", err
	}
	return newMilestoneID(repo, reply.Number), err
}

// Get reads the instance state identified by ID, returning a populated resource object, or an error if not found.
func (p *msProvider) Get(ctx context.Context, id resource.ID) (*issues.Milestone, error) {
	// Simply perform a GET to fetch the information.
	repo, number := fromMilestoneID(id)
	url, err := p.baseURL(repo)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url + "/" + strconv.Itoa(number))
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
	var ms milestone
	if err = json.Unmarshal(respbody, &ms); err != nil {
		return nil, err
	}
	return &issues.Milestone{
		Title:       ms.Title,
		DueOn:       ms.DueOn,
		Description: ms.Description,
		State:       stringToState(ms.State),
		Repo:        &repo,
	}, nil
}

// InspectChange checks what impacts a hypothetical update will have on the resource's properties.
func (p *msProvider) InspectChange(ctx context.Context, id resource.ID,
	old *issues.Milestone, new *issues.Milestone, diff *resource.ObjectDiff) ([]string, error) {
	return nil, nil
}

// Update updates an existing resource with new values.  Only those values in the provided property bag are updated
// to new values.  The resource ID is returned and may be different if the resource had to be recreated.
func (p *msProvider) Update(ctx context.Context, id resource.ID,
	old *issues.Milestone, new *issues.Milestone, diff *resource.ObjectDiff) error {
	// Just do a PATCH with the full object state to update the target.
	repo, number := fromMilestoneID(id)
	body, err := json.Marshal(milestone{
		Title:       new.Title,
		DueOn:       new.DueOn,
		Description: new.Description,
		State:       stateToString(new.State),
	})
	if err != nil {
		return err
	}
	url, err := p.baseURL(repo)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPatch, url+"/"+strconv.Itoa(number), bytes.NewBuffer(body))
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
func (p *msProvider) Delete(ctx context.Context, id resource.ID, obj issues.Milestone) error {
	// DELETE the resource using its ID (number).
	repo, number := fromMilestoneID(id)
	url, err := p.baseURL(repo)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodDelete, url+"/"+strconv.Itoa(number), nil)
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
