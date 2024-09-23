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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
	"net/http"
	"net/url"
	"regexp"
)

// ManagedEnvironmentsServer is a fake server for instances of the armappcontainers.ManagedEnvironmentsClient type.
type ManagedEnvironmentsServer struct {
	// BeginCreateOrUpdate is the fake for method ManagedEnvironmentsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, environmentName string, environmentEnvelope armappcontainers.ManagedEnvironment, options *armappcontainers.ManagedEnvironmentsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armappcontainers.ManagedEnvironmentsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ManagedEnvironmentsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, environmentName string, options *armappcontainers.ManagedEnvironmentsClientBeginDeleteOptions) (resp azfake.PollerResponder[armappcontainers.ManagedEnvironmentsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagedEnvironmentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, environmentName string, options *armappcontainers.ManagedEnvironmentsClientGetOptions) (resp azfake.Responder[armappcontainers.ManagedEnvironmentsClientGetResponse], errResp azfake.ErrorResponder)

	// GetAuthToken is the fake for method ManagedEnvironmentsClient.GetAuthToken
	// HTTP status codes to indicate success: http.StatusOK
	GetAuthToken func(ctx context.Context, resourceGroupName string, environmentName string, options *armappcontainers.ManagedEnvironmentsClientGetAuthTokenOptions) (resp azfake.Responder[armappcontainers.ManagedEnvironmentsClientGetAuthTokenResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method ManagedEnvironmentsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armappcontainers.ManagedEnvironmentsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armappcontainers.ManagedEnvironmentsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method ManagedEnvironmentsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armappcontainers.ManagedEnvironmentsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armappcontainers.ManagedEnvironmentsClientListBySubscriptionResponse])

	// NewListWorkloadProfileStatesPager is the fake for method ManagedEnvironmentsClient.NewListWorkloadProfileStatesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListWorkloadProfileStatesPager func(resourceGroupName string, environmentName string, options *armappcontainers.ManagedEnvironmentsClientListWorkloadProfileStatesOptions) (resp azfake.PagerResponder[armappcontainers.ManagedEnvironmentsClientListWorkloadProfileStatesResponse])

	// BeginUpdate is the fake for method ManagedEnvironmentsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, environmentName string, environmentEnvelope armappcontainers.ManagedEnvironment, options *armappcontainers.ManagedEnvironmentsClientBeginUpdateOptions) (resp azfake.PollerResponder[armappcontainers.ManagedEnvironmentsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewManagedEnvironmentsServerTransport creates a new instance of ManagedEnvironmentsServerTransport with the provided implementation.
// The returned ManagedEnvironmentsServerTransport instance is connected to an instance of armappcontainers.ManagedEnvironmentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedEnvironmentsServerTransport(srv *ManagedEnvironmentsServer) *ManagedEnvironmentsServerTransport {
	return &ManagedEnvironmentsServerTransport{
		srv:                               srv,
		beginCreateOrUpdate:               newTracker[azfake.PollerResponder[armappcontainers.ManagedEnvironmentsClientCreateOrUpdateResponse]](),
		beginDelete:                       newTracker[azfake.PollerResponder[armappcontainers.ManagedEnvironmentsClientDeleteResponse]](),
		newListByResourceGroupPager:       newTracker[azfake.PagerResponder[armappcontainers.ManagedEnvironmentsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:        newTracker[azfake.PagerResponder[armappcontainers.ManagedEnvironmentsClientListBySubscriptionResponse]](),
		newListWorkloadProfileStatesPager: newTracker[azfake.PagerResponder[armappcontainers.ManagedEnvironmentsClientListWorkloadProfileStatesResponse]](),
		beginUpdate:                       newTracker[azfake.PollerResponder[armappcontainers.ManagedEnvironmentsClientUpdateResponse]](),
	}
}

// ManagedEnvironmentsServerTransport connects instances of armappcontainers.ManagedEnvironmentsClient to instances of ManagedEnvironmentsServer.
// Don't use this type directly, use NewManagedEnvironmentsServerTransport instead.
type ManagedEnvironmentsServerTransport struct {
	srv                               *ManagedEnvironmentsServer
	beginCreateOrUpdate               *tracker[azfake.PollerResponder[armappcontainers.ManagedEnvironmentsClientCreateOrUpdateResponse]]
	beginDelete                       *tracker[azfake.PollerResponder[armappcontainers.ManagedEnvironmentsClientDeleteResponse]]
	newListByResourceGroupPager       *tracker[azfake.PagerResponder[armappcontainers.ManagedEnvironmentsClientListByResourceGroupResponse]]
	newListBySubscriptionPager        *tracker[azfake.PagerResponder[armappcontainers.ManagedEnvironmentsClientListBySubscriptionResponse]]
	newListWorkloadProfileStatesPager *tracker[azfake.PagerResponder[armappcontainers.ManagedEnvironmentsClientListWorkloadProfileStatesResponse]]
	beginUpdate                       *tracker[azfake.PollerResponder[armappcontainers.ManagedEnvironmentsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ManagedEnvironmentsServerTransport.
func (m *ManagedEnvironmentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedEnvironmentsClient.BeginCreateOrUpdate":
		resp, err = m.dispatchBeginCreateOrUpdate(req)
	case "ManagedEnvironmentsClient.BeginDelete":
		resp, err = m.dispatchBeginDelete(req)
	case "ManagedEnvironmentsClient.Get":
		resp, err = m.dispatchGet(req)
	case "ManagedEnvironmentsClient.GetAuthToken":
		resp, err = m.dispatchGetAuthToken(req)
	case "ManagedEnvironmentsClient.NewListByResourceGroupPager":
		resp, err = m.dispatchNewListByResourceGroupPager(req)
	case "ManagedEnvironmentsClient.NewListBySubscriptionPager":
		resp, err = m.dispatchNewListBySubscriptionPager(req)
	case "ManagedEnvironmentsClient.NewListWorkloadProfileStatesPager":
		resp, err = m.dispatchNewListWorkloadProfileStatesPager(req)
	case "ManagedEnvironmentsClient.BeginUpdate":
		resp, err = m.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedEnvironmentsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := m.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappcontainers.ManagedEnvironment](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, environmentNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		m.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		m.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		m.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (m *ManagedEnvironmentsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if m.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := m.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginDelete(req.Context(), resourceGroupNameParam, environmentNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		m.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		m.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		m.beginDelete.remove(req)
	}

	return resp, nil
}

func (m *ManagedEnvironmentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, environmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedEnvironment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedEnvironmentsServerTransport) dispatchGetAuthToken(req *http.Request) (*http.Response, error) {
	if m.srv.GetAuthToken == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAuthToken not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getAuthtoken`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.GetAuthToken(req.Context(), resourceGroupNameParam, environmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EnvironmentAuthToken, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedEnvironmentsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := m.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		m.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armappcontainers.ManagedEnvironmentsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		m.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (m *ManagedEnvironmentsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := m.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := m.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		m.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armappcontainers.ManagedEnvironmentsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		m.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (m *ManagedEnvironmentsServerTransport) dispatchNewListWorkloadProfileStatesPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListWorkloadProfileStatesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListWorkloadProfileStatesPager not implemented")}
	}
	newListWorkloadProfileStatesPager := m.newListWorkloadProfileStatesPager.get(req)
	if newListWorkloadProfileStatesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workloadProfileStates`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListWorkloadProfileStatesPager(resourceGroupNameParam, environmentNameParam, nil)
		newListWorkloadProfileStatesPager = &resp
		m.newListWorkloadProfileStatesPager.add(req, newListWorkloadProfileStatesPager)
		server.PagerResponderInjectNextLinks(newListWorkloadProfileStatesPager, req, func(page *armappcontainers.ManagedEnvironmentsClientListWorkloadProfileStatesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListWorkloadProfileStatesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListWorkloadProfileStatesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListWorkloadProfileStatesPager) {
		m.newListWorkloadProfileStatesPager.remove(req)
	}
	return resp, nil
}

func (m *ManagedEnvironmentsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := m.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappcontainers.ManagedEnvironment](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginUpdate(req.Context(), resourceGroupNameParam, environmentNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		m.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		m.beginUpdate.remove(req)
	}

	return resp, nil
}
