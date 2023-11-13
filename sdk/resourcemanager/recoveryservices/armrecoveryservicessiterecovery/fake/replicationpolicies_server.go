//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ReplicationPoliciesServer is a fake server for instances of the armrecoveryservicessiterecovery.ReplicationPoliciesClient type.
type ReplicationPoliciesServer struct {
	// BeginCreate is the fake for method ReplicationPoliciesClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceName string, resourceGroupName string, policyName string, input armrecoveryservicessiterecovery.CreatePolicyInput, options *armrecoveryservicessiterecovery.ReplicationPoliciesClientBeginCreateOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationPoliciesClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ReplicationPoliciesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceName string, resourceGroupName string, policyName string, options *armrecoveryservicessiterecovery.ReplicationPoliciesClientBeginDeleteOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationPoliciesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ReplicationPoliciesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceName string, resourceGroupName string, policyName string, options *armrecoveryservicessiterecovery.ReplicationPoliciesClientGetOptions) (resp azfake.Responder[armrecoveryservicessiterecovery.ReplicationPoliciesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ReplicationPoliciesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceName string, resourceGroupName string, options *armrecoveryservicessiterecovery.ReplicationPoliciesClientListOptions) (resp azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationPoliciesClientListResponse])

	// BeginUpdate is the fake for method ReplicationPoliciesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceName string, resourceGroupName string, policyName string, input armrecoveryservicessiterecovery.UpdatePolicyInput, options *armrecoveryservicessiterecovery.ReplicationPoliciesClientBeginUpdateOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationPoliciesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewReplicationPoliciesServerTransport creates a new instance of ReplicationPoliciesServerTransport with the provided implementation.
// The returned ReplicationPoliciesServerTransport instance is connected to an instance of armrecoveryservicessiterecovery.ReplicationPoliciesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReplicationPoliciesServerTransport(srv *ReplicationPoliciesServer) *ReplicationPoliciesServerTransport {
	return &ReplicationPoliciesServerTransport{
		srv:          srv,
		beginCreate:  newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationPoliciesClientCreateResponse]](),
		beginDelete:  newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationPoliciesClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationPoliciesClientListResponse]](),
		beginUpdate:  newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationPoliciesClientUpdateResponse]](),
	}
}

// ReplicationPoliciesServerTransport connects instances of armrecoveryservicessiterecovery.ReplicationPoliciesClient to instances of ReplicationPoliciesServer.
// Don't use this type directly, use NewReplicationPoliciesServerTransport instead.
type ReplicationPoliciesServerTransport struct {
	srv          *ReplicationPoliciesServer
	beginCreate  *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationPoliciesClientCreateResponse]]
	beginDelete  *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationPoliciesClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationPoliciesClientListResponse]]
	beginUpdate  *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationPoliciesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ReplicationPoliciesServerTransport.
func (r *ReplicationPoliciesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ReplicationPoliciesClient.BeginCreate":
		resp, err = r.dispatchBeginCreate(req)
	case "ReplicationPoliciesClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "ReplicationPoliciesClient.Get":
		resp, err = r.dispatchGet(req)
	case "ReplicationPoliciesClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	case "ReplicationPoliciesClient.BeginUpdate":
		resp, err = r.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ReplicationPoliciesServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := r.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicessiterecovery.CreatePolicyInput](req)
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		policyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreate(req.Context(), resourceNameParam, resourceGroupNameParam, policyNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		r.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		r.beginCreate.remove(req)
	}

	return resp, nil
}

func (r *ReplicationPoliciesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		policyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceNameParam, resourceGroupNameParam, policyNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		r.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		r.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		r.beginDelete.remove(req)
	}

	return resp, nil
}

func (r *ReplicationPoliciesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	policyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceNameParam, resourceGroupNameParam, policyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Policy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ReplicationPoliciesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListPager(resourceNameParam, resourceGroupNameParam, nil)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armrecoveryservicessiterecovery.ReplicationPoliciesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}

func (r *ReplicationPoliciesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := r.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicessiterecovery.UpdatePolicyInput](req)
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		policyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginUpdate(req.Context(), resourceNameParam, resourceGroupNameParam, policyNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		r.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		r.beginUpdate.remove(req)
	}

	return resp, nil
}
