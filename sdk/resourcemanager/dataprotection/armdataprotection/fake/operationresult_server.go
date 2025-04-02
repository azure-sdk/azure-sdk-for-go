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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// OperationResultServer is a fake server for instances of the armdataprotection.OperationResultClient type.
type OperationResultServer struct {
	// Get is the fake for method OperationResultClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	Get func(ctx context.Context, operationID string, location string, options *armdataprotection.OperationResultClientGetOptions) (resp azfake.Responder[armdataprotection.OperationResultClientGetResponse], errResp azfake.ErrorResponder)
}

// NewOperationResultServerTransport creates a new instance of OperationResultServerTransport with the provided implementation.
// The returned OperationResultServerTransport instance is connected to an instance of armdataprotection.OperationResultClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOperationResultServerTransport(srv *OperationResultServer) *OperationResultServerTransport {
	return &OperationResultServerTransport{srv: srv}
}

// OperationResultServerTransport connects instances of armdataprotection.OperationResultClient to instances of OperationResultServer.
// Don't use this type directly, use NewOperationResultServerTransport instead.
type OperationResultServerTransport struct {
	srv *OperationResultServer
}

// Do implements the policy.Transporter interface for OperationResultServerTransport.
func (o *OperationResultServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return o.dispatchToMethodFake(req, method)
}

func (o *OperationResultServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if operationResultServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = operationResultServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "OperationResultClient.Get":
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

func (o *OperationResultServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataProtection/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operationResults/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Get(req.Context(), operationIDParam, locationParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OperationJobExtendedInfo, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).AzureAsyncOperation; val != nil {
		resp.Header.Set("Azure-AsyncOperation", *val)
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", *val)
	}
	if val := server.GetResponse(respr).RetryAfter; val != nil {
		resp.Header.Set("Retry-After", strconv.FormatInt(int64(*val), 10))
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to OperationResultServerTransport
var operationResultServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
