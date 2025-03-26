// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
	"net/http"
	"net/url"
	"regexp"
)

// OperationResultsServer is a fake server for instances of the armrecoveryservicesdatareplication.OperationResultsClient type.
type OperationResultsServer struct {
	// Get is the fake for method OperationResultsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, operationID string, options *armrecoveryservicesdatareplication.OperationResultsClientGetOptions) (resp azfake.Responder[armrecoveryservicesdatareplication.OperationResultsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewOperationResultsServerTransport creates a new instance of OperationResultsServerTransport with the provided implementation.
// The returned OperationResultsServerTransport instance is connected to an instance of armrecoveryservicesdatareplication.OperationResultsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOperationResultsServerTransport(srv *OperationResultsServer) *OperationResultsServerTransport {
	return &OperationResultsServerTransport{srv: srv}
}

// OperationResultsServerTransport connects instances of armrecoveryservicesdatareplication.OperationResultsClient to instances of OperationResultsServer.
// Don't use this type directly, use NewOperationResultsServerTransport instead.
type OperationResultsServerTransport struct {
	srv *OperationResultsServer
}

// Do implements the policy.Transporter interface for OperationResultsServerTransport.
func (o *OperationResultsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return o.dispatchToMethodFake(req, method)
}

func (o *OperationResultsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if operationResultsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = operationResultsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "OperationResultsClient.Get":
				res.resp, res.err = o.dispatchGet(req)
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

func (o *OperationResultsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/operationResults/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/{operationId}`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Get(req.Context(), resourceGroupNameParam, operationIDParam, nil)
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

// set this to conditionally intercept incoming requests to OperationResultsServerTransport
var operationResultsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
