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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v4"
	"net/http"
	"net/url"
	"regexp"
)

// ExtendedUeInformationServer is a fake server for instances of the armmobilenetwork.ExtendedUeInformationClient type.
type ExtendedUeInformationServer struct {
	// Get is the fake for method ExtendedUeInformationClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, ueID string, options *armmobilenetwork.ExtendedUeInformationClientGetOptions) (resp azfake.Responder[armmobilenetwork.ExtendedUeInformationClientGetResponse], errResp azfake.ErrorResponder)
}

// NewExtendedUeInformationServerTransport creates a new instance of ExtendedUeInformationServerTransport with the provided implementation.
// The returned ExtendedUeInformationServerTransport instance is connected to an instance of armmobilenetwork.ExtendedUeInformationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExtendedUeInformationServerTransport(srv *ExtendedUeInformationServer) *ExtendedUeInformationServerTransport {
	return &ExtendedUeInformationServerTransport{srv: srv}
}

// ExtendedUeInformationServerTransport connects instances of armmobilenetwork.ExtendedUeInformationClient to instances of ExtendedUeInformationServer.
// Don't use this type directly, use NewExtendedUeInformationServerTransport instead.
type ExtendedUeInformationServerTransport struct {
	srv *ExtendedUeInformationServer
}

// Do implements the policy.Transporter interface for ExtendedUeInformationServerTransport.
func (e *ExtendedUeInformationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *ExtendedUeInformationServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if extendedUeInformationServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = extendedUeInformationServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ExtendedUeInformationClient.Get":
				res.resp, res.err = e.dispatchGet(req)
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

func (e *ExtendedUeInformationServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MobileNetwork/packetCoreControlPlanes/(?P<packetCoreControlPlaneName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ues/(?P<ueId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extendedInformation/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	packetCoreControlPlaneNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("packetCoreControlPlaneName")])
	if err != nil {
		return nil, err
	}
	ueIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("ueId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, packetCoreControlPlaneNameParam, ueIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExtendedUeInfo, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ExtendedUeInformationServerTransport
var extendedUeInformationServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
