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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity/v2"
	"net/http"
	"regexp"
)

// GenerateAwsTemplateServer is a fake server for instances of the armhybridconnectivity.GenerateAwsTemplateClient type.
type GenerateAwsTemplateServer struct {
	// Post is the fake for method GenerateAwsTemplateClient.Post
	// HTTP status codes to indicate success: http.StatusOK
	Post func(ctx context.Context, generateAwsTemplateRequest armhybridconnectivity.GenerateAwsTemplateRequest, options *armhybridconnectivity.GenerateAwsTemplateClientPostOptions) (resp azfake.Responder[armhybridconnectivity.GenerateAwsTemplateClientPostResponse], errResp azfake.ErrorResponder)
}

// NewGenerateAwsTemplateServerTransport creates a new instance of GenerateAwsTemplateServerTransport with the provided implementation.
// The returned GenerateAwsTemplateServerTransport instance is connected to an instance of armhybridconnectivity.GenerateAwsTemplateClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGenerateAwsTemplateServerTransport(srv *GenerateAwsTemplateServer) *GenerateAwsTemplateServerTransport {
	return &GenerateAwsTemplateServerTransport{srv: srv}
}

// GenerateAwsTemplateServerTransport connects instances of armhybridconnectivity.GenerateAwsTemplateClient to instances of GenerateAwsTemplateServer.
// Don't use this type directly, use NewGenerateAwsTemplateServerTransport instead.
type GenerateAwsTemplateServerTransport struct {
	srv *GenerateAwsTemplateServer
}

// Do implements the policy.Transporter interface for GenerateAwsTemplateServerTransport.
func (g *GenerateAwsTemplateServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return g.dispatchToMethodFake(req, method)
}

func (g *GenerateAwsTemplateServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if generateAwsTemplateServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = generateAwsTemplateServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "GenerateAwsTemplateClient.Post":
				res.resp, res.err = g.dispatchPost(req)
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

func (g *GenerateAwsTemplateServerTransport) dispatchPost(req *http.Request) (*http.Response, error) {
	if g.srv.Post == nil {
		return nil, &nonRetriableError{errors.New("fake for method Post not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridConnectivity/generateAwsTemplate`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armhybridconnectivity.GenerateAwsTemplateRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Post(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PostResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to GenerateAwsTemplateServerTransport
var generateAwsTemplateServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}