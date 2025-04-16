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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
)

// RestorableDroppedManagedDatabasesServer is a fake server for instances of the armsql.RestorableDroppedManagedDatabasesClient type.
type RestorableDroppedManagedDatabasesServer struct {
	// Get is the fake for method RestorableDroppedManagedDatabasesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, options *armsql.RestorableDroppedManagedDatabasesClientGetOptions) (resp azfake.Responder[armsql.RestorableDroppedManagedDatabasesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByInstancePager is the fake for method RestorableDroppedManagedDatabasesClient.NewListByInstancePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByInstancePager func(resourceGroupName string, managedInstanceName string, options *armsql.RestorableDroppedManagedDatabasesClientListByInstanceOptions) (resp azfake.PagerResponder[armsql.RestorableDroppedManagedDatabasesClientListByInstanceResponse])
}

// NewRestorableDroppedManagedDatabasesServerTransport creates a new instance of RestorableDroppedManagedDatabasesServerTransport with the provided implementation.
// The returned RestorableDroppedManagedDatabasesServerTransport instance is connected to an instance of armsql.RestorableDroppedManagedDatabasesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRestorableDroppedManagedDatabasesServerTransport(srv *RestorableDroppedManagedDatabasesServer) *RestorableDroppedManagedDatabasesServerTransport {
	return &RestorableDroppedManagedDatabasesServerTransport{
		srv:                    srv,
		newListByInstancePager: newTracker[azfake.PagerResponder[armsql.RestorableDroppedManagedDatabasesClientListByInstanceResponse]](),
	}
}

// RestorableDroppedManagedDatabasesServerTransport connects instances of armsql.RestorableDroppedManagedDatabasesClient to instances of RestorableDroppedManagedDatabasesServer.
// Don't use this type directly, use NewRestorableDroppedManagedDatabasesServerTransport instead.
type RestorableDroppedManagedDatabasesServerTransport struct {
	srv                    *RestorableDroppedManagedDatabasesServer
	newListByInstancePager *tracker[azfake.PagerResponder[armsql.RestorableDroppedManagedDatabasesClientListByInstanceResponse]]
}

// Do implements the policy.Transporter interface for RestorableDroppedManagedDatabasesServerTransport.
func (r *RestorableDroppedManagedDatabasesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *RestorableDroppedManagedDatabasesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if restorableDroppedManagedDatabasesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = restorableDroppedManagedDatabasesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "RestorableDroppedManagedDatabasesClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "RestorableDroppedManagedDatabasesClient.NewListByInstancePager":
				res.resp, res.err = r.dispatchNewListByInstancePager(req)
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

func (r *RestorableDroppedManagedDatabasesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorableDroppedDatabases/(?P<restorableDroppedDatabaseId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
	if err != nil {
		return nil, err
	}
	restorableDroppedDatabaseIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("restorableDroppedDatabaseId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, managedInstanceNameParam, restorableDroppedDatabaseIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RestorableDroppedManagedDatabase, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RestorableDroppedManagedDatabasesServerTransport) dispatchNewListByInstancePager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByInstancePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByInstancePager not implemented")}
	}
	newListByInstancePager := r.newListByInstancePager.get(req)
	if newListByInstancePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorableDroppedDatabases`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByInstancePager(resourceGroupNameParam, managedInstanceNameParam, nil)
		newListByInstancePager = &resp
		r.newListByInstancePager.add(req, newListByInstancePager)
		server.PagerResponderInjectNextLinks(newListByInstancePager, req, func(page *armsql.RestorableDroppedManagedDatabasesClientListByInstanceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByInstancePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByInstancePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByInstancePager) {
		r.newListByInstancePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to RestorableDroppedManagedDatabasesServerTransport
var restorableDroppedManagedDatabasesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
