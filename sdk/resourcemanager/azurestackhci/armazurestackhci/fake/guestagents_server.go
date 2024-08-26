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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
	"net/http"
	"net/url"
	"regexp"
)

// GuestAgentsServer is a fake server for instances of the armazurestackhci.GuestAgentsClient type.
type GuestAgentsServer struct {
	// NewListPager is the fake for method GuestAgentsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceURI string, options *armazurestackhci.GuestAgentsClientListOptions) (resp azfake.PagerResponder[armazurestackhci.GuestAgentsClientListResponse])
}

// NewGuestAgentsServerTransport creates a new instance of GuestAgentsServerTransport with the provided implementation.
// The returned GuestAgentsServerTransport instance is connected to an instance of armazurestackhci.GuestAgentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGuestAgentsServerTransport(srv *GuestAgentsServer) *GuestAgentsServerTransport {
	return &GuestAgentsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armazurestackhci.GuestAgentsClientListResponse]](),
	}
}

// GuestAgentsServerTransport connects instances of armazurestackhci.GuestAgentsClient to instances of GuestAgentsServer.
// Don't use this type directly, use NewGuestAgentsServerTransport instead.
type GuestAgentsServerTransport struct {
	srv          *GuestAgentsServer
	newListPager *tracker[azfake.PagerResponder[armazurestackhci.GuestAgentsClientListResponse]]
}

// Do implements the policy.Transporter interface for GuestAgentsServerTransport.
func (g *GuestAgentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GuestAgentsClient.NewListPager":
		resp, err = g.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GuestAgentsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := g.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHCI/virtualMachineInstances/default/guestAgents`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		resp := g.srv.NewListPager(resourceURIParam, nil)
		newListPager = &resp
		g.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armazurestackhci.GuestAgentsClientListResponse, createLink func() string) {
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
