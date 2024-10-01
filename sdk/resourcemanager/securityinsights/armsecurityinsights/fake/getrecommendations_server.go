//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
	"net/http"
	"net/url"
	"regexp"
)

// GetRecommendationsServer is a fake server for instances of the armsecurityinsights.GetRecommendationsClient type.
type GetRecommendationsServer struct {
	// NewListPager is the fake for method GetRecommendationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, options *armsecurityinsights.GetRecommendationsClientListOptions) (resp azfake.PagerResponder[armsecurityinsights.GetRecommendationsClientListResponse])
}

// NewGetRecommendationsServerTransport creates a new instance of GetRecommendationsServerTransport with the provided implementation.
// The returned GetRecommendationsServerTransport instance is connected to an instance of armsecurityinsights.GetRecommendationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGetRecommendationsServerTransport(srv *GetRecommendationsServer) *GetRecommendationsServerTransport {
	return &GetRecommendationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsecurityinsights.GetRecommendationsClientListResponse]](),
	}
}

// GetRecommendationsServerTransport connects instances of armsecurityinsights.GetRecommendationsClient to instances of GetRecommendationsServer.
// Don't use this type directly, use NewGetRecommendationsServerTransport instead.
type GetRecommendationsServerTransport struct {
	srv          *GetRecommendationsServer
	newListPager *tracker[azfake.PagerResponder[armsecurityinsights.GetRecommendationsClientListResponse]]
}

// Do implements the policy.Transporter interface for GetRecommendationsServerTransport.
func (g *GetRecommendationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GetRecommendationsClient.NewListPager":
		resp, err = g.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GetRecommendationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := g.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/recommendations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		resp := g.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, nil)
		newListPager = &resp
		g.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsecurityinsights.GetRecommendationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		g.newListPager.remove(req)
	}
	return resp, nil
}
