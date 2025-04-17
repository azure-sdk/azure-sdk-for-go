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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
	"net/http"
	"net/url"
	"regexp"
)

// VHDsServer is a fake server for instances of the armtestbase.VHDsClient type.
type VHDsServer struct {
	// Delete is the fake for method VHDsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, testBaseAccountName string, vhdName string, options *armtestbase.VHDsClientDeleteOptions) (resp azfake.Responder[armtestbase.VHDsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VHDsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, testBaseAccountName string, vhdName string, options *armtestbase.VHDsClientGetOptions) (resp azfake.Responder[armtestbase.VHDsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByTestBaseAccountPager is the fake for method VHDsClient.NewListByTestBaseAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByTestBaseAccountPager func(resourceGroupName string, testBaseAccountName string, options *armtestbase.VHDsClientListByTestBaseAccountOptions) (resp azfake.PagerResponder[armtestbase.VHDsClientListByTestBaseAccountResponse])
}

// NewVHDsServerTransport creates a new instance of VHDsServerTransport with the provided implementation.
// The returned VHDsServerTransport instance is connected to an instance of armtestbase.VHDsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVHDsServerTransport(srv *VHDsServer) *VHDsServerTransport {
	return &VHDsServerTransport{
		srv:                           srv,
		newListByTestBaseAccountPager: newTracker[azfake.PagerResponder[armtestbase.VHDsClientListByTestBaseAccountResponse]](),
	}
}

// VHDsServerTransport connects instances of armtestbase.VHDsClient to instances of VHDsServer.
// Don't use this type directly, use NewVHDsServerTransport instead.
type VHDsServerTransport struct {
	srv                           *VHDsServer
	newListByTestBaseAccountPager *tracker[azfake.PagerResponder[armtestbase.VHDsClientListByTestBaseAccountResponse]]
}

// Do implements the policy.Transporter interface for VHDsServerTransport.
func (v *VHDsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *VHDsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if vhDsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = vhDsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "VHDsClient.Delete":
				res.resp, res.err = v.dispatchDelete(req)
			case "VHDsClient.Get":
				res.resp, res.err = v.dispatchGet(req)
			case "VHDsClient.NewListByTestBaseAccountPager":
				res.resp, res.err = v.dispatchNewListByTestBaseAccountPager(req)
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

func (v *VHDsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if v.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.TestBase/testBaseAccounts/(?P<testBaseAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vhds/(?P<vhdName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	testBaseAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("testBaseAccountName")])
	if err != nil {
		return nil, err
	}
	vhdNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vhdName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Delete(req.Context(), resourceGroupNameParam, testBaseAccountNameParam, vhdNameParam, nil)
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

func (v *VHDsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.TestBase/testBaseAccounts/(?P<testBaseAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vhds/(?P<vhdName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	testBaseAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("testBaseAccountName")])
	if err != nil {
		return nil, err
	}
	vhdNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vhdName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, testBaseAccountNameParam, vhdNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VHDResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VHDsServerTransport) dispatchNewListByTestBaseAccountPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByTestBaseAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByTestBaseAccountPager not implemented")}
	}
	newListByTestBaseAccountPager := v.newListByTestBaseAccountPager.get(req)
	if newListByTestBaseAccountPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.TestBase/testBaseAccounts/(?P<testBaseAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vhds`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		testBaseAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("testBaseAccountName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListByTestBaseAccountPager(resourceGroupNameParam, testBaseAccountNameParam, nil)
		newListByTestBaseAccountPager = &resp
		v.newListByTestBaseAccountPager.add(req, newListByTestBaseAccountPager)
		server.PagerResponderInjectNextLinks(newListByTestBaseAccountPager, req, func(page *armtestbase.VHDsClientListByTestBaseAccountResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByTestBaseAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListByTestBaseAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByTestBaseAccountPager) {
		v.newListByTestBaseAccountPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to VHDsServerTransport
var vhDsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
