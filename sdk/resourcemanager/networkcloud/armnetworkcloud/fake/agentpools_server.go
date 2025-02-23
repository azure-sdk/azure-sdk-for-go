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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
	"net/http"
	"net/url"
	"regexp"
)

// AgentPoolsServer is a fake server for instances of the armnetworkcloud.AgentPoolsClient type.
type AgentPoolsServer struct {
	// BeginCreateOrUpdate is the fake for method AgentPoolsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, kubernetesClusterName string, agentPoolName string, agentPoolParameters armnetworkcloud.AgentPool, options *armnetworkcloud.AgentPoolsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetworkcloud.AgentPoolsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AgentPoolsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, kubernetesClusterName string, agentPoolName string, options *armnetworkcloud.AgentPoolsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetworkcloud.AgentPoolsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AgentPoolsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, kubernetesClusterName string, agentPoolName string, options *armnetworkcloud.AgentPoolsClientGetOptions) (resp azfake.Responder[armnetworkcloud.AgentPoolsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByKubernetesClusterPager is the fake for method AgentPoolsClient.NewListByKubernetesClusterPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByKubernetesClusterPager func(resourceGroupName string, kubernetesClusterName string, options *armnetworkcloud.AgentPoolsClientListByKubernetesClusterOptions) (resp azfake.PagerResponder[armnetworkcloud.AgentPoolsClientListByKubernetesClusterResponse])

	// BeginUpdate is the fake for method AgentPoolsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, kubernetesClusterName string, agentPoolName string, agentPoolUpdateParameters armnetworkcloud.AgentPoolPatchParameters, options *armnetworkcloud.AgentPoolsClientBeginUpdateOptions) (resp azfake.PollerResponder[armnetworkcloud.AgentPoolsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewAgentPoolsServerTransport creates a new instance of AgentPoolsServerTransport with the provided implementation.
// The returned AgentPoolsServerTransport instance is connected to an instance of armnetworkcloud.AgentPoolsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAgentPoolsServerTransport(srv *AgentPoolsServer) *AgentPoolsServerTransport {
	return &AgentPoolsServerTransport{
		srv:                             srv,
		beginCreateOrUpdate:             newTracker[azfake.PollerResponder[armnetworkcloud.AgentPoolsClientCreateOrUpdateResponse]](),
		beginDelete:                     newTracker[azfake.PollerResponder[armnetworkcloud.AgentPoolsClientDeleteResponse]](),
		newListByKubernetesClusterPager: newTracker[azfake.PagerResponder[armnetworkcloud.AgentPoolsClientListByKubernetesClusterResponse]](),
		beginUpdate:                     newTracker[azfake.PollerResponder[armnetworkcloud.AgentPoolsClientUpdateResponse]](),
	}
}

// AgentPoolsServerTransport connects instances of armnetworkcloud.AgentPoolsClient to instances of AgentPoolsServer.
// Don't use this type directly, use NewAgentPoolsServerTransport instead.
type AgentPoolsServerTransport struct {
	srv                             *AgentPoolsServer
	beginCreateOrUpdate             *tracker[azfake.PollerResponder[armnetworkcloud.AgentPoolsClientCreateOrUpdateResponse]]
	beginDelete                     *tracker[azfake.PollerResponder[armnetworkcloud.AgentPoolsClientDeleteResponse]]
	newListByKubernetesClusterPager *tracker[azfake.PagerResponder[armnetworkcloud.AgentPoolsClientListByKubernetesClusterResponse]]
	beginUpdate                     *tracker[azfake.PollerResponder[armnetworkcloud.AgentPoolsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for AgentPoolsServerTransport.
func (a *AgentPoolsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AgentPoolsClient.BeginCreateOrUpdate":
		resp, err = a.dispatchBeginCreateOrUpdate(req)
	case "AgentPoolsClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AgentPoolsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AgentPoolsClient.NewListByKubernetesClusterPager":
		resp, err = a.dispatchNewListByKubernetesClusterPager(req)
	case "AgentPoolsClient.BeginUpdate":
		resp, err = a.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AgentPoolsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := a.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/kubernetesClusters/(?P<kubernetesClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentPools/(?P<agentPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.AgentPool](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		kubernetesClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("kubernetesClusterName")])
		if err != nil {
			return nil, err
		}
		agentPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("agentPoolName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
		var options *armnetworkcloud.AgentPoolsClientBeginCreateOrUpdateOptions
		if ifMatchParam != nil || ifNoneMatchParam != nil {
			options = &armnetworkcloud.AgentPoolsClientBeginCreateOrUpdateOptions{
				IfMatch:     ifMatchParam,
				IfNoneMatch: ifNoneMatchParam,
			}
		}
		respr, errRespr := a.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, kubernetesClusterNameParam, agentPoolNameParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		a.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		a.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		a.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (a *AgentPoolsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/kubernetesClusters/(?P<kubernetesClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentPools/(?P<agentPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		kubernetesClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("kubernetesClusterName")])
		if err != nil {
			return nil, err
		}
		agentPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("agentPoolName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
		var options *armnetworkcloud.AgentPoolsClientBeginDeleteOptions
		if ifMatchParam != nil || ifNoneMatchParam != nil {
			options = &armnetworkcloud.AgentPoolsClientBeginDeleteOptions{
				IfMatch:     ifMatchParam,
				IfNoneMatch: ifNoneMatchParam,
			}
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, kubernetesClusterNameParam, agentPoolNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		a.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *AgentPoolsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/kubernetesClusters/(?P<kubernetesClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentPools/(?P<agentPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	kubernetesClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("kubernetesClusterName")])
	if err != nil {
		return nil, err
	}
	agentPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("agentPoolName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, kubernetesClusterNameParam, agentPoolNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AgentPool, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AgentPoolsServerTransport) dispatchNewListByKubernetesClusterPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByKubernetesClusterPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByKubernetesClusterPager not implemented")}
	}
	newListByKubernetesClusterPager := a.newListByKubernetesClusterPager.get(req)
	if newListByKubernetesClusterPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/kubernetesClusters/(?P<kubernetesClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentPools`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		kubernetesClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("kubernetesClusterName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByKubernetesClusterPager(resourceGroupNameParam, kubernetesClusterNameParam, nil)
		newListByKubernetesClusterPager = &resp
		a.newListByKubernetesClusterPager.add(req, newListByKubernetesClusterPager)
		server.PagerResponderInjectNextLinks(newListByKubernetesClusterPager, req, func(page *armnetworkcloud.AgentPoolsClientListByKubernetesClusterResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByKubernetesClusterPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByKubernetesClusterPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByKubernetesClusterPager) {
		a.newListByKubernetesClusterPager.remove(req)
	}
	return resp, nil
}

func (a *AgentPoolsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := a.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/kubernetesClusters/(?P<kubernetesClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentPools/(?P<agentPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.AgentPoolPatchParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		kubernetesClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("kubernetesClusterName")])
		if err != nil {
			return nil, err
		}
		agentPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("agentPoolName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
		var options *armnetworkcloud.AgentPoolsClientBeginUpdateOptions
		if ifMatchParam != nil || ifNoneMatchParam != nil {
			options = &armnetworkcloud.AgentPoolsClientBeginUpdateOptions{
				IfMatch:     ifMatchParam,
				IfNoneMatch: ifNoneMatchParam,
			}
		}
		respr, errRespr := a.srv.BeginUpdate(req.Context(), resourceGroupNameParam, kubernetesClusterNameParam, agentPoolNameParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		a.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		a.beginUpdate.remove(req)
	}

	return resp, nil
}
