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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
	"net/http"
	"regexp"
)

// TokensServer is a fake server for instances of the armpolicy.TokensClient type.
type TokensServer struct {
	// BeginAcquire is the fake for method TokensClient.BeginAcquire
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginAcquire func(ctx context.Context, parameters armpolicy.TokenRequest, options *armpolicy.TokensClientBeginAcquireOptions) (resp azfake.PollerResponder[armpolicy.TokensClientAcquireResponse], errResp azfake.ErrorResponder)
}

// NewTokensServerTransport creates a new instance of TokensServerTransport with the provided implementation.
// The returned TokensServerTransport instance is connected to an instance of armpolicy.TokensClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTokensServerTransport(srv *TokensServer) *TokensServerTransport {
	return &TokensServerTransport{
		srv:          srv,
		beginAcquire: newTracker[azfake.PollerResponder[armpolicy.TokensClientAcquireResponse]](),
	}
}

// TokensServerTransport connects instances of armpolicy.TokensClient to instances of TokensServer.
// Don't use this type directly, use NewTokensServerTransport instead.
type TokensServerTransport struct {
	srv          *TokensServer
	beginAcquire *tracker[azfake.PollerResponder[armpolicy.TokensClientAcquireResponse]]
}

// Do implements the policy.Transporter interface for TokensServerTransport.
func (t *TokensServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return t.dispatchToMethodFake(req, method)
}

func (t *TokensServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if tokensServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = tokensServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "TokensClient.BeginAcquire":
				res.resp, res.err = t.dispatchBeginAcquire(req)
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

func (t *TokensServerTransport) dispatchBeginAcquire(req *http.Request) (*http.Response, error) {
	if t.srv.BeginAcquire == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginAcquire not implemented")}
	}
	beginAcquire := t.beginAcquire.get(req)
	if beginAcquire == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/acquirePolicyToken`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armpolicy.TokenRequest](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginAcquire(req.Context(), body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginAcquire = &respr
		t.beginAcquire.add(req, beginAcquire)
	}

	resp, err := server.PollerResponderNext(beginAcquire, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		t.beginAcquire.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginAcquire) {
		t.beginAcquire.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to TokensServerTransport
var tokensServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
