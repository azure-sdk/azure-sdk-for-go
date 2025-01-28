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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
	"net/http"
	"net/url"
	"regexp"
)

// HypervClusterControllerServer is a fake server for instances of the armmigrate.HypervClusterControllerClient type.
type HypervClusterControllerServer struct {
	// BeginCreateCluster is the fake for method HypervClusterControllerClient.BeginCreateCluster
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateCluster func(ctx context.Context, resourceGroupName string, siteName string, clusterName string, body armmigrate.HypervCluster, options *armmigrate.HypervClusterControllerClientBeginCreateClusterOptions) (resp azfake.PollerResponder[armmigrate.HypervClusterControllerClientCreateClusterResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method HypervClusterControllerClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, siteName string, clusterName string, options *armmigrate.HypervClusterControllerClientDeleteOptions) (resp azfake.Responder[armmigrate.HypervClusterControllerClientDeleteResponse], errResp azfake.ErrorResponder)

	// GetCluster is the fake for method HypervClusterControllerClient.GetCluster
	// HTTP status codes to indicate success: http.StatusOK
	GetCluster func(ctx context.Context, resourceGroupName string, siteName string, clusterName string, options *armmigrate.HypervClusterControllerClientGetClusterOptions) (resp azfake.Responder[armmigrate.HypervClusterControllerClientGetClusterResponse], errResp azfake.ErrorResponder)

	// NewListByHypervSitePager is the fake for method HypervClusterControllerClient.NewListByHypervSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHypervSitePager func(resourceGroupName string, siteName string, options *armmigrate.HypervClusterControllerClientListByHypervSiteOptions) (resp azfake.PagerResponder[armmigrate.HypervClusterControllerClientListByHypervSiteResponse])
}

// NewHypervClusterControllerServerTransport creates a new instance of HypervClusterControllerServerTransport with the provided implementation.
// The returned HypervClusterControllerServerTransport instance is connected to an instance of armmigrate.HypervClusterControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHypervClusterControllerServerTransport(srv *HypervClusterControllerServer) *HypervClusterControllerServerTransport {
	return &HypervClusterControllerServerTransport{
		srv:                      srv,
		beginCreateCluster:       newTracker[azfake.PollerResponder[armmigrate.HypervClusterControllerClientCreateClusterResponse]](),
		newListByHypervSitePager: newTracker[azfake.PagerResponder[armmigrate.HypervClusterControllerClientListByHypervSiteResponse]](),
	}
}

// HypervClusterControllerServerTransport connects instances of armmigrate.HypervClusterControllerClient to instances of HypervClusterControllerServer.
// Don't use this type directly, use NewHypervClusterControllerServerTransport instead.
type HypervClusterControllerServerTransport struct {
	srv                      *HypervClusterControllerServer
	beginCreateCluster       *tracker[azfake.PollerResponder[armmigrate.HypervClusterControllerClientCreateClusterResponse]]
	newListByHypervSitePager *tracker[azfake.PagerResponder[armmigrate.HypervClusterControllerClientListByHypervSiteResponse]]
}

// Do implements the policy.Transporter interface for HypervClusterControllerServerTransport.
func (h *HypervClusterControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HypervClusterControllerClient.BeginCreateCluster":
		resp, err = h.dispatchBeginCreateCluster(req)
	case "HypervClusterControllerClient.Delete":
		resp, err = h.dispatchDelete(req)
	case "HypervClusterControllerClient.GetCluster":
		resp, err = h.dispatchGetCluster(req)
	case "HypervClusterControllerClient.NewListByHypervSitePager":
		resp, err = h.dispatchNewListByHypervSitePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HypervClusterControllerServerTransport) dispatchBeginCreateCluster(req *http.Request) (*http.Response, error) {
	if h.srv.BeginCreateCluster == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateCluster not implemented")}
	}
	beginCreateCluster := h.beginCreateCluster.get(req)
	if beginCreateCluster == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrate.HypervCluster](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := h.srv.BeginCreateCluster(req.Context(), resourceGroupNameParam, siteNameParam, clusterNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateCluster = &respr
		h.beginCreateCluster.add(req, beginCreateCluster)
	}

	resp, err := server.PollerResponderNext(beginCreateCluster, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		h.beginCreateCluster.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateCluster) {
		h.beginCreateCluster.remove(req)
	}

	return resp, nil
}

func (h *HypervClusterControllerServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if h.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Delete(req.Context(), resourceGroupNameParam, siteNameParam, clusterNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HypervClusterControllerServerTransport) dispatchGetCluster(req *http.Request) (*http.Response, error) {
	if h.srv.GetCluster == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetCluster not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.GetCluster(req.Context(), resourceGroupNameParam, siteNameParam, clusterNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HypervCluster, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HypervClusterControllerServerTransport) dispatchNewListByHypervSitePager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListByHypervSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHypervSitePager not implemented")}
	}
	newListByHypervSitePager := h.newListByHypervSitePager.get(req)
	if newListByHypervSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		var options *armmigrate.HypervClusterControllerClientListByHypervSiteOptions
		if filterParam != nil {
			options = &armmigrate.HypervClusterControllerClientListByHypervSiteOptions{
				Filter: filterParam,
			}
		}
		resp := h.srv.NewListByHypervSitePager(resourceGroupNameParam, siteNameParam, options)
		newListByHypervSitePager = &resp
		h.newListByHypervSitePager.add(req, newListByHypervSitePager)
		server.PagerResponderInjectNextLinks(newListByHypervSitePager, req, func(page *armmigrate.HypervClusterControllerClientListByHypervSiteResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByHypervSitePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		h.newListByHypervSitePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByHypervSitePager) {
		h.newListByHypervSitePager.remove(req)
	}
	return resp, nil
}
