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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
	"net/http"
	"net/url"
	"regexp"
)

// UpdatesServer is a fake server for instances of the armmaintenance.UpdatesClient type.
type UpdatesServer struct {
	// NewListPager is the fake for method UpdatesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, providerName string, resourceType string, resourceName string, options *armmaintenance.UpdatesClientListOptions) (resp azfake.PagerResponder[armmaintenance.UpdatesClientListResponse])

	// NewListParentPager is the fake for method UpdatesClient.NewListParentPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListParentPager func(resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, options *armmaintenance.UpdatesClientListParentOptions) (resp azfake.PagerResponder[armmaintenance.UpdatesClientListParentResponse])
}

// NewUpdatesServerTransport creates a new instance of UpdatesServerTransport with the provided implementation.
// The returned UpdatesServerTransport instance is connected to an instance of armmaintenance.UpdatesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewUpdatesServerTransport(srv *UpdatesServer) *UpdatesServerTransport {
	return &UpdatesServerTransport{
		srv:                srv,
		newListPager:       newTracker[azfake.PagerResponder[armmaintenance.UpdatesClientListResponse]](),
		newListParentPager: newTracker[azfake.PagerResponder[armmaintenance.UpdatesClientListParentResponse]](),
	}
}

// UpdatesServerTransport connects instances of armmaintenance.UpdatesClient to instances of UpdatesServer.
// Don't use this type directly, use NewUpdatesServerTransport instead.
type UpdatesServerTransport struct {
	srv                *UpdatesServer
	newListPager       *tracker[azfake.PagerResponder[armmaintenance.UpdatesClientListResponse]]
	newListParentPager *tracker[azfake.PagerResponder[armmaintenance.UpdatesClientListParentResponse]]
}

// Do implements the policy.Transporter interface for UpdatesServerTransport.
func (u *UpdatesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return u.dispatchToMethodFake(req, method)
}

func (u *UpdatesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if updatesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = updatesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "UpdatesClient.NewListPager":
				res.resp, res.err = u.dispatchNewListPager(req)
			case "UpdatesClient.NewListParentPager":
				res.resp, res.err = u.dispatchNewListParentPager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (u *UpdatesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if u.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := u.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<providerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Maintenance/updates`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		providerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerName")])
		if err != nil {
			return nil, err
		}
		resourceTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceType")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resp := u.srv.NewListPager(resourceGroupNameParam, providerNameParam, resourceTypeParam, resourceNameParam, nil)
		newListPager = &resp
		u.newListPager.add(req, newListPager)
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		u.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		u.newListPager.remove(req)
	}
	return resp, nil
}

func (u *UpdatesServerTransport) dispatchNewListParentPager(req *http.Request) (*http.Response, error) {
	if u.srv.NewListParentPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListParentPager not implemented")}
	}
	newListParentPager := u.newListParentPager.get(req)
	if newListParentPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<providerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceParentType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceParentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Maintenance/updates`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 7 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		providerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerName")])
		if err != nil {
			return nil, err
		}
		resourceParentTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceParentType")])
		if err != nil {
			return nil, err
		}
		resourceParentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceParentName")])
		if err != nil {
			return nil, err
		}
		resourceTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceType")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resp := u.srv.NewListParentPager(resourceGroupNameParam, providerNameParam, resourceParentTypeParam, resourceParentNameParam, resourceTypeParam, resourceNameParam, nil)
		newListParentPager = &resp
		u.newListParentPager.add(req, newListParentPager)
	}
	resp, err := server.PagerResponderNext(newListParentPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		u.newListParentPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListParentPager) {
		u.newListParentPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to UpdatesServerTransport
var updatesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
