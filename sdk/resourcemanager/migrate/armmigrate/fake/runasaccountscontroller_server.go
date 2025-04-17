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

// RunAsAccountsControllerServer is a fake server for instances of the armmigrate.RunAsAccountsControllerClient type.
type RunAsAccountsControllerServer struct {
	// Get is the fake for method RunAsAccountsControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, accountName string, options *armmigrate.RunAsAccountsControllerClientGetOptions) (resp azfake.Responder[armmigrate.RunAsAccountsControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByVmwareSitePager is the fake for method RunAsAccountsControllerClient.NewListByVmwareSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByVmwareSitePager func(resourceGroupName string, siteName string, options *armmigrate.RunAsAccountsControllerClientListByVmwareSiteOptions) (resp azfake.PagerResponder[armmigrate.RunAsAccountsControllerClientListByVmwareSiteResponse])
}

// NewRunAsAccountsControllerServerTransport creates a new instance of RunAsAccountsControllerServerTransport with the provided implementation.
// The returned RunAsAccountsControllerServerTransport instance is connected to an instance of armmigrate.RunAsAccountsControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRunAsAccountsControllerServerTransport(srv *RunAsAccountsControllerServer) *RunAsAccountsControllerServerTransport {
	return &RunAsAccountsControllerServerTransport{
		srv:                      srv,
		newListByVmwareSitePager: newTracker[azfake.PagerResponder[armmigrate.RunAsAccountsControllerClientListByVmwareSiteResponse]](),
	}
}

// RunAsAccountsControllerServerTransport connects instances of armmigrate.RunAsAccountsControllerClient to instances of RunAsAccountsControllerServer.
// Don't use this type directly, use NewRunAsAccountsControllerServerTransport instead.
type RunAsAccountsControllerServerTransport struct {
	srv                      *RunAsAccountsControllerServer
	newListByVmwareSitePager *tracker[azfake.PagerResponder[armmigrate.RunAsAccountsControllerClientListByVmwareSiteResponse]]
}

// Do implements the policy.Transporter interface for RunAsAccountsControllerServerTransport.
func (r *RunAsAccountsControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *RunAsAccountsControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if runAsAccountsControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = runAsAccountsControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "RunAsAccountsControllerClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "RunAsAccountsControllerClient.NewListByVmwareSitePager":
				res.resp, res.err = r.dispatchNewListByVmwareSitePager(req)
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

func (r *RunAsAccountsControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runAsAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, accountNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VmwareRunAsAccountResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RunAsAccountsControllerServerTransport) dispatchNewListByVmwareSitePager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByVmwareSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByVmwareSitePager not implemented")}
	}
	newListByVmwareSitePager := r.newListByVmwareSitePager.get(req)
	if newListByVmwareSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runAsAccounts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
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
		resp := r.srv.NewListByVmwareSitePager(resourceGroupNameParam, siteNameParam, nil)
		newListByVmwareSitePager = &resp
		r.newListByVmwareSitePager.add(req, newListByVmwareSitePager)
		server.PagerResponderInjectNextLinks(newListByVmwareSitePager, req, func(page *armmigrate.RunAsAccountsControllerClientListByVmwareSiteResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByVmwareSitePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByVmwareSitePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByVmwareSitePager) {
		r.newListByVmwareSitePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to RunAsAccountsControllerServerTransport
var runAsAccountsControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
