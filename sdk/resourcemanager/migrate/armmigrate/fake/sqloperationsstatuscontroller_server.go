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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
	"net/http"
	"net/url"
	"regexp"
)

// SQLOperationsStatusControllerServer is a fake server for instances of the armmigrate.SQLOperationsStatusControllerClient type.
type SQLOperationsStatusControllerServer struct {
	// GetSQLOperationStatus is the fake for method SQLOperationsStatusControllerClient.GetSQLOperationStatus
	// HTTP status codes to indicate success: http.StatusOK
	GetSQLOperationStatus func(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, operationStatusName string, options *armmigrate.SQLOperationsStatusControllerClientGetSQLOperationStatusOptions) (resp azfake.Responder[armmigrate.SQLOperationsStatusControllerClientGetSQLOperationStatusResponse], errResp azfake.ErrorResponder)
}

// NewSQLOperationsStatusControllerServerTransport creates a new instance of SQLOperationsStatusControllerServerTransport with the provided implementation.
// The returned SQLOperationsStatusControllerServerTransport instance is connected to an instance of armmigrate.SQLOperationsStatusControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSQLOperationsStatusControllerServerTransport(srv *SQLOperationsStatusControllerServer) *SQLOperationsStatusControllerServerTransport {
	return &SQLOperationsStatusControllerServerTransport{srv: srv}
}

// SQLOperationsStatusControllerServerTransport connects instances of armmigrate.SQLOperationsStatusControllerClient to instances of SQLOperationsStatusControllerServer.
// Don't use this type directly, use NewSQLOperationsStatusControllerServerTransport instead.
type SQLOperationsStatusControllerServerTransport struct {
	srv *SQLOperationsStatusControllerServer
}

// Do implements the policy.Transporter interface for SQLOperationsStatusControllerServerTransport.
func (s *SQLOperationsStatusControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SQLOperationsStatusControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if sqlOperationsStatusControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = sqlOperationsStatusControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SQLOperationsStatusControllerClient.GetSQLOperationStatus":
				res.resp, res.err = s.dispatchGetSQLOperationStatus(req)
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

func (s *SQLOperationsStatusControllerServerTransport) dispatchGetSQLOperationStatus(req *http.Request) (*http.Response, error) {
	if s.srv.GetSQLOperationStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSQLOperationStatus not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlSites/(?P<sqlSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operationsStatus/(?P<operationStatusName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	sqlSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlSiteName")])
	if err != nil {
		return nil, err
	}
	operationStatusNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationStatusName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetSQLOperationStatus(req.Context(), resourceGroupNameParam, siteNameParam, sqlSiteNameParam, operationStatusNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OperationStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to SQLOperationsStatusControllerServerTransport
var sqlOperationsStatusControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
