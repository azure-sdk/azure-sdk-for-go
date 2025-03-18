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

// ManagedDatabaseTablesServer is a fake server for instances of the armsql.ManagedDatabaseTablesClient type.
type ManagedDatabaseTablesServer struct {
	// Get is the fake for method ManagedDatabaseTablesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, options *armsql.ManagedDatabaseTablesClientGetOptions) (resp azfake.Responder[armsql.ManagedDatabaseTablesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListBySchemaPager is the fake for method ManagedDatabaseTablesClient.NewListBySchemaPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySchemaPager func(resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, options *armsql.ManagedDatabaseTablesClientListBySchemaOptions) (resp azfake.PagerResponder[armsql.ManagedDatabaseTablesClientListBySchemaResponse])
}

// NewManagedDatabaseTablesServerTransport creates a new instance of ManagedDatabaseTablesServerTransport with the provided implementation.
// The returned ManagedDatabaseTablesServerTransport instance is connected to an instance of armsql.ManagedDatabaseTablesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedDatabaseTablesServerTransport(srv *ManagedDatabaseTablesServer) *ManagedDatabaseTablesServerTransport {
	return &ManagedDatabaseTablesServerTransport{
		srv:                  srv,
		newListBySchemaPager: newTracker[azfake.PagerResponder[armsql.ManagedDatabaseTablesClientListBySchemaResponse]](),
	}
}

// ManagedDatabaseTablesServerTransport connects instances of armsql.ManagedDatabaseTablesClient to instances of ManagedDatabaseTablesServer.
// Don't use this type directly, use NewManagedDatabaseTablesServerTransport instead.
type ManagedDatabaseTablesServerTransport struct {
	srv                  *ManagedDatabaseTablesServer
	newListBySchemaPager *tracker[azfake.PagerResponder[armsql.ManagedDatabaseTablesClientListBySchemaResponse]]
}

// Do implements the policy.Transporter interface for ManagedDatabaseTablesServerTransport.
func (m *ManagedDatabaseTablesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *ManagedDatabaseTablesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if managedDatabaseTablesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = managedDatabaseTablesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ManagedDatabaseTablesClient.Get":
				res.resp, res.err = m.dispatchGet(req)
			case "ManagedDatabaseTablesClient.NewListBySchemaPager":
				res.resp, res.err = m.dispatchNewListBySchemaPager(req)
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

func (m *ManagedDatabaseTablesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
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
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
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
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, schemaNameParam, tableNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DatabaseTable, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedDatabaseTablesServerTransport) dispatchNewListBySchemaPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListBySchemaPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySchemaPager not implemented")}
	}
	newListBySchemaPager := m.newListBySchemaPager.get(req)
	if newListBySchemaPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		schemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armsql.ManagedDatabaseTablesClientListBySchemaOptions
		if filterParam != nil {
			options = &armsql.ManagedDatabaseTablesClientListBySchemaOptions{
				Filter: filterParam,
			}
		}
		resp := m.srv.NewListBySchemaPager(resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, schemaNameParam, options)
		newListBySchemaPager = &resp
		m.newListBySchemaPager.add(req, newListBySchemaPager)
		server.PagerResponderInjectNextLinks(newListBySchemaPager, req, func(page *armsql.ManagedDatabaseTablesClientListBySchemaResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySchemaPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListBySchemaPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySchemaPager) {
		m.newListBySchemaPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ManagedDatabaseTablesServerTransport
var managedDatabaseTablesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
