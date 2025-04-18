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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
	"net/http"
	"net/url"
	"regexp"
)

// OperationStatusesServer is a fake server for instances of the armmediaservices.OperationStatusesClient type.
type OperationStatusesServer struct {
	// Get is the fake for method OperationStatusesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, locationName string, operationID string, options *armmediaservices.OperationStatusesClientGetOptions) (resp azfake.Responder[armmediaservices.OperationStatusesClientGetResponse], errResp azfake.ErrorResponder)
}

// NewOperationStatusesServerTransport creates a new instance of OperationStatusesServerTransport with the provided implementation.
// The returned OperationStatusesServerTransport instance is connected to an instance of armmediaservices.OperationStatusesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOperationStatusesServerTransport(srv *OperationStatusesServer) *OperationStatusesServerTransport {
	return &OperationStatusesServerTransport{srv: srv}
}

// OperationStatusesServerTransport connects instances of armmediaservices.OperationStatusesClient to instances of OperationStatusesServer.
// Don't use this type directly, use NewOperationStatusesServerTransport instead.
type OperationStatusesServerTransport struct {
	srv *OperationStatusesServer
}

// Do implements the policy.Transporter interface for OperationStatusesServerTransport.
func (o *OperationStatusesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return o.dispatchToMethodFake(req, method)
}

func (o *OperationStatusesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if operationStatusesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = operationStatusesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "OperationStatusesClient.Get":
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

func (o *OperationStatusesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Media/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/mediaServicesOperationStatuses/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Get(req.Context(), locationNameParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).MediaServiceOperationStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to OperationStatusesServerTransport
var operationStatusesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
