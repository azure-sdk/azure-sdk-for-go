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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
	"net/http"
	"regexp"
)

// AssessmentTypesServer is a fake server for instances of the armadvisor.AssessmentTypesClient type.
type AssessmentTypesServer struct {
	// NewListPager is the fake for method AssessmentTypesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armadvisor.AssessmentTypesClientListOptions) (resp azfake.PagerResponder[armadvisor.AssessmentTypesClientListResponse])
}

// NewAssessmentTypesServerTransport creates a new instance of AssessmentTypesServerTransport with the provided implementation.
// The returned AssessmentTypesServerTransport instance is connected to an instance of armadvisor.AssessmentTypesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAssessmentTypesServerTransport(srv *AssessmentTypesServer) *AssessmentTypesServerTransport {
	return &AssessmentTypesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armadvisor.AssessmentTypesClientListResponse]](),
	}
}

// AssessmentTypesServerTransport connects instances of armadvisor.AssessmentTypesClient to instances of AssessmentTypesServer.
// Don't use this type directly, use NewAssessmentTypesServerTransport instead.
type AssessmentTypesServerTransport struct {
	srv          *AssessmentTypesServer
	newListPager *tracker[azfake.PagerResponder[armadvisor.AssessmentTypesClientListResponse]]
}

// Do implements the policy.Transporter interface for AssessmentTypesServerTransport.
func (a *AssessmentTypesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AssessmentTypesClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AssessmentTypesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/assessmentTypes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListPager(nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armadvisor.AssessmentTypesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}
