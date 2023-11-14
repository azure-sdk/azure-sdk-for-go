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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ElasticPoolOperationsServer is a fake server for instances of the armsql.ElasticPoolOperationsClient type.
type ElasticPoolOperationsServer struct {
	// Cancel is the fake for method ElasticPoolOperationsClient.Cancel
	// HTTP status codes to indicate success: http.StatusOK
	Cancel func(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, operationID string, options *armsql.ElasticPoolOperationsClientCancelOptions) (resp azfake.Responder[armsql.ElasticPoolOperationsClientCancelResponse], errResp azfake.ErrorResponder)

	// NewListByElasticPoolPager is the fake for method ElasticPoolOperationsClient.NewListByElasticPoolPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByElasticPoolPager func(resourceGroupName string, serverName string, elasticPoolName string, options *armsql.ElasticPoolOperationsClientListByElasticPoolOptions) (resp azfake.PagerResponder[armsql.ElasticPoolOperationsClientListByElasticPoolResponse])
}

// NewElasticPoolOperationsServerTransport creates a new instance of ElasticPoolOperationsServerTransport with the provided implementation.
// The returned ElasticPoolOperationsServerTransport instance is connected to an instance of armsql.ElasticPoolOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewElasticPoolOperationsServerTransport(srv *ElasticPoolOperationsServer) *ElasticPoolOperationsServerTransport {
	return &ElasticPoolOperationsServerTransport{
		srv:                       srv,
		newListByElasticPoolPager: newTracker[azfake.PagerResponder[armsql.ElasticPoolOperationsClientListByElasticPoolResponse]](),
	}
}

// ElasticPoolOperationsServerTransport connects instances of armsql.ElasticPoolOperationsClient to instances of ElasticPoolOperationsServer.
// Don't use this type directly, use NewElasticPoolOperationsServerTransport instead.
type ElasticPoolOperationsServerTransport struct {
	srv                       *ElasticPoolOperationsServer
	newListByElasticPoolPager *tracker[azfake.PagerResponder[armsql.ElasticPoolOperationsClientListByElasticPoolResponse]]
}

// Do implements the policy.Transporter interface for ElasticPoolOperationsServerTransport.
func (e *ElasticPoolOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ElasticPoolOperationsClient.Cancel":
		resp, err = e.dispatchCancel(req)
	case "ElasticPoolOperationsClient.NewListByElasticPoolPager":
		resp, err = e.dispatchNewListByElasticPoolPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ElasticPoolOperationsServerTransport) dispatchCancel(req *http.Request) (*http.Response, error) {
	if e.srv.Cancel == nil {
		return nil, &nonRetriableError{errors.New("fake for method Cancel not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/elasticPools/(?P<elasticPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operations/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancel`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	elasticPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("elasticPoolName")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Cancel(req.Context(), resourceGroupNameParam, serverNameParam, elasticPoolNameParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ElasticPoolOperationsServerTransport) dispatchNewListByElasticPoolPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByElasticPoolPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByElasticPoolPager not implemented")}
	}
	newListByElasticPoolPager := e.newListByElasticPoolPager.get(req)
	if newListByElasticPoolPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/elasticPools/(?P<elasticPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		elasticPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("elasticPoolName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListByElasticPoolPager(resourceGroupNameParam, serverNameParam, elasticPoolNameParam, nil)
		newListByElasticPoolPager = &resp
		e.newListByElasticPoolPager.add(req, newListByElasticPoolPager)
		server.PagerResponderInjectNextLinks(newListByElasticPoolPager, req, func(page *armsql.ElasticPoolOperationsClientListByElasticPoolResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByElasticPoolPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByElasticPoolPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByElasticPoolPager) {
		e.newListByElasticPoolPager.remove(req)
	}
	return resp, nil
}
