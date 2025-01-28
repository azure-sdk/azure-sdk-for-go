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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationmodernization/armmigrationmodernization"
	"net/http"
	"net/url"
	"regexp"
)

// DeployedResourceServer is a fake server for instances of the armmigrationmodernization.DeployedResourceClient type.
type DeployedResourceServer struct {
	// Get is the fake for method DeployedResourceClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, deployedResourceName string, options *armmigrationmodernization.DeployedResourceClientGetOptions) (resp azfake.Responder[armmigrationmodernization.DeployedResourceClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DeployedResourceClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(subscriptionID string, resourceGroupName string, modernizeProjectName string, options *armmigrationmodernization.DeployedResourceClientListOptions) (resp azfake.PagerResponder[armmigrationmodernization.DeployedResourceClientListResponse])
}

// NewDeployedResourceServerTransport creates a new instance of DeployedResourceServerTransport with the provided implementation.
// The returned DeployedResourceServerTransport instance is connected to an instance of armmigrationmodernization.DeployedResourceClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDeployedResourceServerTransport(srv *DeployedResourceServer) *DeployedResourceServerTransport {
	return &DeployedResourceServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armmigrationmodernization.DeployedResourceClientListResponse]](),
	}
}

// DeployedResourceServerTransport connects instances of armmigrationmodernization.DeployedResourceClient to instances of DeployedResourceServer.
// Don't use this type directly, use NewDeployedResourceServerTransport instead.
type DeployedResourceServerTransport struct {
	srv          *DeployedResourceServer
	newListPager *tracker[azfake.PagerResponder[armmigrationmodernization.DeployedResourceClientListResponse]]
}

// Do implements the policy.Transporter interface for DeployedResourceServerTransport.
func (d *DeployedResourceServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DeployedResourceClient.Get":
		resp, err = d.dispatchGet(req)
	case "DeployedResourceClient.NewListPager":
		resp, err = d.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DeployedResourceServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/modernizeProjects/(?P<modernizeProjectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deployedResources/(?P<deployedResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	modernizeProjectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("modernizeProjectName")])
	if err != nil {
		return nil, err
	}
	deployedResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deployedResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), subscriptionIDParam, resourceGroupNameParam, modernizeProjectNameParam, deployedResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeployedResourceModel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeployedResourceServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/modernizeProjects/(?P<modernizeProjectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deployedResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		modernizeProjectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("modernizeProjectName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListPager(subscriptionIDParam, resourceGroupNameParam, modernizeProjectNameParam, nil)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmigrationmodernization.DeployedResourceClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		d.newListPager.remove(req)
	}
	return resp, nil
}
