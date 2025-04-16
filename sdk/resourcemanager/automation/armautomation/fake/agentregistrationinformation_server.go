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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
	"net/http"
	"net/url"
	"regexp"
)

// AgentRegistrationInformationServer is a fake server for instances of the armautomation.AgentRegistrationInformationClient type.
type AgentRegistrationInformationServer struct {
	// Get is the fake for method AgentRegistrationInformationClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, automationAccountName string, options *armautomation.AgentRegistrationInformationClientGetOptions) (resp azfake.Responder[armautomation.AgentRegistrationInformationClientGetResponse], errResp azfake.ErrorResponder)

	// RegenerateKey is the fake for method AgentRegistrationInformationClient.RegenerateKey
	// HTTP status codes to indicate success: http.StatusOK
	RegenerateKey func(ctx context.Context, resourceGroupName string, automationAccountName string, parameters armautomation.AgentRegistrationRegenerateKeyParameter, options *armautomation.AgentRegistrationInformationClientRegenerateKeyOptions) (resp azfake.Responder[armautomation.AgentRegistrationInformationClientRegenerateKeyResponse], errResp azfake.ErrorResponder)
}

// NewAgentRegistrationInformationServerTransport creates a new instance of AgentRegistrationInformationServerTransport with the provided implementation.
// The returned AgentRegistrationInformationServerTransport instance is connected to an instance of armautomation.AgentRegistrationInformationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAgentRegistrationInformationServerTransport(srv *AgentRegistrationInformationServer) *AgentRegistrationInformationServerTransport {
	return &AgentRegistrationInformationServerTransport{srv: srv}
}

// AgentRegistrationInformationServerTransport connects instances of armautomation.AgentRegistrationInformationClient to instances of AgentRegistrationInformationServer.
// Don't use this type directly, use NewAgentRegistrationInformationServerTransport instead.
type AgentRegistrationInformationServerTransport struct {
	srv *AgentRegistrationInformationServer
}

// Do implements the policy.Transporter interface for AgentRegistrationInformationServerTransport.
func (a *AgentRegistrationInformationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *AgentRegistrationInformationServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if agentRegistrationInformationServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = agentRegistrationInformationServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "AgentRegistrationInformationClient.Get":
				res.resp, res.err = a.dispatchGet(req)
			case "AgentRegistrationInformationClient.RegenerateKey":
				res.resp, res.err = a.dispatchRegenerateKey(req)
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

func (a *AgentRegistrationInformationServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentRegistrationInformation`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, automationAccountNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AgentRegistration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AgentRegistrationInformationServerTransport) dispatchRegenerateKey(req *http.Request) (*http.Response, error) {
	if a.srv.RegenerateKey == nil {
		return nil, &nonRetriableError{errors.New("fake for method RegenerateKey not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentRegistrationInformation/regenerateKey`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armautomation.AgentRegistrationRegenerateKeyParameter](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.RegenerateKey(req.Context(), resourceGroupNameParam, automationAccountNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AgentRegistration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to AgentRegistrationInformationServerTransport
var agentRegistrationInformationServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
