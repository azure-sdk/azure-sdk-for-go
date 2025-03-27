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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationdiscovery/armmigrationdiscovery"
	"net/http"
	"net/url"
	"regexp"
)

// VmwareOperationsStatusServer is a fake server for instances of the armmigrationdiscovery.VmwareOperationsStatusClient type.
type VmwareOperationsStatusServer struct {
	// GetVmwareOperationStatus is the fake for method VmwareOperationsStatusClient.GetVmwareOperationStatus
	// HTTP status codes to indicate success: http.StatusOK
	GetVmwareOperationStatus func(ctx context.Context, resourceGroupName string, siteName string, operationStatusName string, options *armmigrationdiscovery.VmwareOperationsStatusClientGetVmwareOperationStatusOptions) (resp azfake.Responder[armmigrationdiscovery.VmwareOperationsStatusClientGetVmwareOperationStatusResponse], errResp azfake.ErrorResponder)
}

// NewVmwareOperationsStatusServerTransport creates a new instance of VmwareOperationsStatusServerTransport with the provided implementation.
// The returned VmwareOperationsStatusServerTransport instance is connected to an instance of armmigrationdiscovery.VmwareOperationsStatusClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVmwareOperationsStatusServerTransport(srv *VmwareOperationsStatusServer) *VmwareOperationsStatusServerTransport {
	return &VmwareOperationsStatusServerTransport{srv: srv}
}

// VmwareOperationsStatusServerTransport connects instances of armmigrationdiscovery.VmwareOperationsStatusClient to instances of VmwareOperationsStatusServer.
// Don't use this type directly, use NewVmwareOperationsStatusServerTransport instead.
type VmwareOperationsStatusServerTransport struct {
	srv *VmwareOperationsStatusServer
}

// Do implements the policy.Transporter interface for VmwareOperationsStatusServerTransport.
func (v *VmwareOperationsStatusServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *VmwareOperationsStatusServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if vmwareOperationsStatusServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = vmwareOperationsStatusServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "VmwareOperationsStatusClient.GetVmwareOperationStatus":
				res.resp, res.err = v.dispatchGetVmwareOperationStatus(req)
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

func (v *VmwareOperationsStatusServerTransport) dispatchGetVmwareOperationStatus(req *http.Request) (*http.Response, error) {
	if v.srv.GetVmwareOperationStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetVmwareOperationStatus not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operationsStatus/(?P<operationStatusName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
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
	operationStatusNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationStatusName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.GetVmwareOperationStatus(req.Context(), resourceGroupNameParam, siteNameParam, operationStatusNameParam, nil)
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

// set this to conditionally intercept incoming requests to VmwareOperationsStatusServerTransport
var vmwareOperationsStatusServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
