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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
	"net/http"
	"net/url"
	"regexp"
)

// GenerateCostDetailsReportServer is a fake server for instances of the armcostmanagement.GenerateCostDetailsReportClient type.
type GenerateCostDetailsReportServer struct {
	// BeginCreateOperation is the fake for method GenerateCostDetailsReportClient.BeginCreateOperation
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginCreateOperation func(ctx context.Context, scope string, parameters armcostmanagement.GenerateCostDetailsReportRequestDefinition, options *armcostmanagement.GenerateCostDetailsReportClientBeginCreateOperationOptions) (resp azfake.PollerResponder[armcostmanagement.GenerateCostDetailsReportClientCreateOperationResponse], errResp azfake.ErrorResponder)

	// BeginGetOperationResults is the fake for method GenerateCostDetailsReportClient.BeginGetOperationResults
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGetOperationResults func(ctx context.Context, scope string, operationID string, options *armcostmanagement.GenerateCostDetailsReportClientBeginGetOperationResultsOptions) (resp azfake.PollerResponder[armcostmanagement.GenerateCostDetailsReportClientGetOperationResultsResponse], errResp azfake.ErrorResponder)
}

// NewGenerateCostDetailsReportServerTransport creates a new instance of GenerateCostDetailsReportServerTransport with the provided implementation.
// The returned GenerateCostDetailsReportServerTransport instance is connected to an instance of armcostmanagement.GenerateCostDetailsReportClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGenerateCostDetailsReportServerTransport(srv *GenerateCostDetailsReportServer) *GenerateCostDetailsReportServerTransport {
	return &GenerateCostDetailsReportServerTransport{
		srv:                      srv,
		beginCreateOperation:     newTracker[azfake.PollerResponder[armcostmanagement.GenerateCostDetailsReportClientCreateOperationResponse]](),
		beginGetOperationResults: newTracker[azfake.PollerResponder[armcostmanagement.GenerateCostDetailsReportClientGetOperationResultsResponse]](),
	}
}

// GenerateCostDetailsReportServerTransport connects instances of armcostmanagement.GenerateCostDetailsReportClient to instances of GenerateCostDetailsReportServer.
// Don't use this type directly, use NewGenerateCostDetailsReportServerTransport instead.
type GenerateCostDetailsReportServerTransport struct {
	srv                      *GenerateCostDetailsReportServer
	beginCreateOperation     *tracker[azfake.PollerResponder[armcostmanagement.GenerateCostDetailsReportClientCreateOperationResponse]]
	beginGetOperationResults *tracker[azfake.PollerResponder[armcostmanagement.GenerateCostDetailsReportClientGetOperationResultsResponse]]
}

// Do implements the policy.Transporter interface for GenerateCostDetailsReportServerTransport.
func (g *GenerateCostDetailsReportServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return g.dispatchToMethodFake(req, method)
}

func (g *GenerateCostDetailsReportServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if generateCostDetailsReportServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = generateCostDetailsReportServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "GenerateCostDetailsReportClient.BeginCreateOperation":
				res.resp, res.err = g.dispatchBeginCreateOperation(req)
			case "GenerateCostDetailsReportClient.BeginGetOperationResults":
				res.resp, res.err = g.dispatchBeginGetOperationResults(req)
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

func (g *GenerateCostDetailsReportServerTransport) dispatchBeginCreateOperation(req *http.Request) (*http.Response, error) {
	if g.srv.BeginCreateOperation == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOperation not implemented")}
	}
	beginCreateOperation := g.beginCreateOperation.get(req)
	if beginCreateOperation == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/generateCostDetailsReport`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcostmanagement.GenerateCostDetailsReportRequestDefinition](req)
		if err != nil {
			return nil, err
		}
		scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginCreateOperation(req.Context(), scopeParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOperation = &respr
		g.beginCreateOperation.add(req, beginCreateOperation)
	}

	resp, err := server.PollerResponderNext(beginCreateOperation, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		g.beginCreateOperation.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOperation) {
		g.beginCreateOperation.remove(req)
	}

	return resp, nil
}

func (g *GenerateCostDetailsReportServerTransport) dispatchBeginGetOperationResults(req *http.Request) (*http.Response, error) {
	if g.srv.BeginGetOperationResults == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGetOperationResults not implemented")}
	}
	beginGetOperationResults := g.beginGetOperationResults.get(req)
	if beginGetOperationResults == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/costDetailsOperationResults/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
		if err != nil {
			return nil, err
		}
		operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginGetOperationResults(req.Context(), scopeParam, operationIDParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGetOperationResults = &respr
		g.beginGetOperationResults.add(req, beginGetOperationResults)
	}

	resp, err := server.PollerResponderNext(beginGetOperationResults, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		g.beginGetOperationResults.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGetOperationResults) {
		g.beginGetOperationResults.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to GenerateCostDetailsReportServerTransport
var generateCostDetailsReportServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
