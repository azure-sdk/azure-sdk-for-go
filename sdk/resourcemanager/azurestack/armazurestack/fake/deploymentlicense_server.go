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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestack/armazurestack"
	"net/http"
	"regexp"
)

// DeploymentLicenseServer is a fake server for instances of the armazurestack.DeploymentLicenseClient type.
type DeploymentLicenseServer struct {
	// Create is the fake for method DeploymentLicenseClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, deploymentLicenseRequest armazurestack.DeploymentLicenseRequest, options *armazurestack.DeploymentLicenseClientCreateOptions) (resp azfake.Responder[armazurestack.DeploymentLicenseClientCreateResponse], errResp azfake.ErrorResponder)
}

// NewDeploymentLicenseServerTransport creates a new instance of DeploymentLicenseServerTransport with the provided implementation.
// The returned DeploymentLicenseServerTransport instance is connected to an instance of armazurestack.DeploymentLicenseClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDeploymentLicenseServerTransport(srv *DeploymentLicenseServer) *DeploymentLicenseServerTransport {
	return &DeploymentLicenseServerTransport{srv: srv}
}

// DeploymentLicenseServerTransport connects instances of armazurestack.DeploymentLicenseClient to instances of DeploymentLicenseServer.
// Don't use this type directly, use NewDeploymentLicenseServerTransport instead.
type DeploymentLicenseServerTransport struct {
	srv *DeploymentLicenseServer
}

// Do implements the policy.Transporter interface for DeploymentLicenseServerTransport.
func (d *DeploymentLicenseServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DeploymentLicenseServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if deploymentLicenseServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = deploymentLicenseServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DeploymentLicenseClient.Create":
				res.resp, res.err = d.dispatchCreate(req)
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

func (d *DeploymentLicenseServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if d.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStack/generateDeploymentLicense`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armazurestack.DeploymentLicenseRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Create(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeploymentLicenseResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to DeploymentLicenseServerTransport
var deploymentLicenseServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
