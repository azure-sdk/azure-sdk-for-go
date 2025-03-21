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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
	"net/http"
	"net/url"
	"regexp"
)

// TieringCostOperationStatusServer is a fake server for instances of the armrecoveryservicesbackup.TieringCostOperationStatusClient type.
type TieringCostOperationStatusServer struct {
	// Get is the fake for method TieringCostOperationStatusClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vaultName string, operationID string, options *armrecoveryservicesbackup.TieringCostOperationStatusClientGetOptions) (resp azfake.Responder[armrecoveryservicesbackup.TieringCostOperationStatusClientGetResponse], errResp azfake.ErrorResponder)
}

// NewTieringCostOperationStatusServerTransport creates a new instance of TieringCostOperationStatusServerTransport with the provided implementation.
// The returned TieringCostOperationStatusServerTransport instance is connected to an instance of armrecoveryservicesbackup.TieringCostOperationStatusClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTieringCostOperationStatusServerTransport(srv *TieringCostOperationStatusServer) *TieringCostOperationStatusServerTransport {
	return &TieringCostOperationStatusServerTransport{srv: srv}
}

// TieringCostOperationStatusServerTransport connects instances of armrecoveryservicesbackup.TieringCostOperationStatusClient to instances of TieringCostOperationStatusServer.
// Don't use this type directly, use NewTieringCostOperationStatusServerTransport instead.
type TieringCostOperationStatusServerTransport struct {
	srv *TieringCostOperationStatusServer
}

// Do implements the policy.Transporter interface for TieringCostOperationStatusServerTransport.
func (t *TieringCostOperationStatusServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return t.dispatchToMethodFake(req, method)
}

func (t *TieringCostOperationStatusServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if tieringCostOperationStatusServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = tieringCostOperationStatusServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "TieringCostOperationStatusClient.Get":
				res.resp, res.err = t.dispatchGet(req)
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

func (t *TieringCostOperationStatusServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupTieringCost/default/operationsStatus/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), resourceGroupNameParam, vaultNameParam, operationIDParam, nil)
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

// set this to conditionally intercept incoming requests to TieringCostOperationStatusServerTransport
var tieringCostOperationStatusServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
