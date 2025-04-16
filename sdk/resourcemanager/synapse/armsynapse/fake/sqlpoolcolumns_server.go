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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"net/http"
	"net/url"
	"regexp"
)

// SQLPoolColumnsServer is a fake server for instances of the armsynapse.SQLPoolColumnsClient type.
type SQLPoolColumnsServer struct {
	// Get is the fake for method SQLPoolColumnsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, schemaName string, tableName string, columnName string, options *armsynapse.SQLPoolColumnsClientGetOptions) (resp azfake.Responder[armsynapse.SQLPoolColumnsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewSQLPoolColumnsServerTransport creates a new instance of SQLPoolColumnsServerTransport with the provided implementation.
// The returned SQLPoolColumnsServerTransport instance is connected to an instance of armsynapse.SQLPoolColumnsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSQLPoolColumnsServerTransport(srv *SQLPoolColumnsServer) *SQLPoolColumnsServerTransport {
	return &SQLPoolColumnsServerTransport{srv: srv}
}

// SQLPoolColumnsServerTransport connects instances of armsynapse.SQLPoolColumnsClient to instances of SQLPoolColumnsServer.
// Don't use this type directly, use NewSQLPoolColumnsServerTransport instead.
type SQLPoolColumnsServerTransport struct {
	srv *SQLPoolColumnsServer
}

// Do implements the policy.Transporter interface for SQLPoolColumnsServerTransport.
func (s *SQLPoolColumnsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SQLPoolColumnsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if sqlPoolColumnsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = sqlPoolColumnsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SQLPoolColumnsClient.Get":
				res.resp, res.err = s.dispatchGet(req)
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

func (s *SQLPoolColumnsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlPools/(?P<sqlPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/columns/(?P<columnName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
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
	sqlPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlPoolName")])
	if err != nil {
		return nil, err
	}
	schemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaName")])
	if err != nil {
		return nil, err
	}
	tableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
	if err != nil {
		return nil, err
	}
	columnNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("columnName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, sqlPoolNameParam, schemaNameParam, tableNameParam, columnNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SQLPoolColumn, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to SQLPoolColumnsServerTransport
var sqlPoolColumnsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
