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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
	"net/http"
)

// ServiceClassificationsNoSubscriptionServer is a fake server for instances of the armsupport.ServiceClassificationsNoSubscriptionClient type.
type ServiceClassificationsNoSubscriptionServer struct {
	// ClassifyServices is the fake for method ServiceClassificationsNoSubscriptionClient.ClassifyServices
	// HTTP status codes to indicate success: http.StatusOK
	ClassifyServices func(ctx context.Context, serviceClassificationRequest armsupport.ServiceClassificationRequest, options *armsupport.ServiceClassificationsNoSubscriptionClientClassifyServicesOptions) (resp azfake.Responder[armsupport.ServiceClassificationsNoSubscriptionClientClassifyServicesResponse], errResp azfake.ErrorResponder)
}

// NewServiceClassificationsNoSubscriptionServerTransport creates a new instance of ServiceClassificationsNoSubscriptionServerTransport with the provided implementation.
// The returned ServiceClassificationsNoSubscriptionServerTransport instance is connected to an instance of armsupport.ServiceClassificationsNoSubscriptionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServiceClassificationsNoSubscriptionServerTransport(srv *ServiceClassificationsNoSubscriptionServer) *ServiceClassificationsNoSubscriptionServerTransport {
	return &ServiceClassificationsNoSubscriptionServerTransport{srv: srv}
}

// ServiceClassificationsNoSubscriptionServerTransport connects instances of armsupport.ServiceClassificationsNoSubscriptionClient to instances of ServiceClassificationsNoSubscriptionServer.
// Don't use this type directly, use NewServiceClassificationsNoSubscriptionServerTransport instead.
type ServiceClassificationsNoSubscriptionServerTransport struct {
	srv *ServiceClassificationsNoSubscriptionServer
}

// Do implements the policy.Transporter interface for ServiceClassificationsNoSubscriptionServerTransport.
func (s *ServiceClassificationsNoSubscriptionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServiceClassificationsNoSubscriptionServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if serviceClassificationsNoSubscriptionServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = serviceClassificationsNoSubscriptionServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ServiceClassificationsNoSubscriptionClient.ClassifyServices":
				res.resp, res.err = s.dispatchClassifyServices(req)
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

func (s *ServiceClassificationsNoSubscriptionServerTransport) dispatchClassifyServices(req *http.Request) (*http.Response, error) {
	if s.srv.ClassifyServices == nil {
		return nil, &nonRetriableError{errors.New("fake for method ClassifyServices not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[armsupport.ServiceClassificationRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.ClassifyServices(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServiceClassificationOutput, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ServiceClassificationsNoSubscriptionServerTransport
var serviceClassificationsNoSubscriptionServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
