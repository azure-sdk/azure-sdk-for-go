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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
	"net/http"
	"net/url"
	"regexp"
)

// ConnectionRaiPolicyServer is a fake server for instances of the armmachinelearning.ConnectionRaiPolicyClient type.
type ConnectionRaiPolicyServer struct {
	// BeginCreate is the fake for method ConnectionRaiPolicyClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, raiPolicyName string, body armmachinelearning.RaiPolicyPropertiesBasicResource, options *armmachinelearning.ConnectionRaiPolicyClientBeginCreateOptions) (resp azfake.PollerResponder[armmachinelearning.ConnectionRaiPolicyClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ConnectionRaiPolicyClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, raiPolicyName string, options *armmachinelearning.ConnectionRaiPolicyClientBeginDeleteOptions) (resp azfake.PollerResponder[armmachinelearning.ConnectionRaiPolicyClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ConnectionRaiPolicyClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, raiPolicyName string, options *armmachinelearning.ConnectionRaiPolicyClientGetOptions) (resp azfake.Responder[armmachinelearning.ConnectionRaiPolicyClientGetResponse], errResp azfake.ErrorResponder)
}

// NewConnectionRaiPolicyServerTransport creates a new instance of ConnectionRaiPolicyServerTransport with the provided implementation.
// The returned ConnectionRaiPolicyServerTransport instance is connected to an instance of armmachinelearning.ConnectionRaiPolicyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConnectionRaiPolicyServerTransport(srv *ConnectionRaiPolicyServer) *ConnectionRaiPolicyServerTransport {
	return &ConnectionRaiPolicyServerTransport{
		srv:         srv,
		beginCreate: newTracker[azfake.PollerResponder[armmachinelearning.ConnectionRaiPolicyClientCreateResponse]](),
		beginDelete: newTracker[azfake.PollerResponder[armmachinelearning.ConnectionRaiPolicyClientDeleteResponse]](),
	}
}

// ConnectionRaiPolicyServerTransport connects instances of armmachinelearning.ConnectionRaiPolicyClient to instances of ConnectionRaiPolicyServer.
// Don't use this type directly, use NewConnectionRaiPolicyServerTransport instead.
type ConnectionRaiPolicyServerTransport struct {
	srv         *ConnectionRaiPolicyServer
	beginCreate *tracker[azfake.PollerResponder[armmachinelearning.ConnectionRaiPolicyClientCreateResponse]]
	beginDelete *tracker[azfake.PollerResponder[armmachinelearning.ConnectionRaiPolicyClientDeleteResponse]]
}

// Do implements the policy.Transporter interface for ConnectionRaiPolicyServerTransport.
func (c *ConnectionRaiPolicyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ConnectionRaiPolicyServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if connectionRaiPolicyServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = connectionRaiPolicyServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ConnectionRaiPolicyClient.BeginCreate":
				res.resp, res.err = c.dispatchBeginCreate(req)
			case "ConnectionRaiPolicyClient.BeginDelete":
				res.resp, res.err = c.dispatchBeginDelete(req)
			case "ConnectionRaiPolicyClient.Get":
				res.resp, res.err = c.dispatchGet(req)
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

func (c *ConnectionRaiPolicyServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := c.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/raiPolicies/(?P<raiPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[armmachinelearning.RaiPolicyPropertiesBasicResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		raiPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("raiPolicyName")])
		if err != nil {
			return nil, err
		}
		proxyAPIVersionUnescaped, err := url.QueryUnescape(qp.Get("proxy-api-version"))
		if err != nil {
			return nil, err
		}
		proxyAPIVersionParam := getOptional(proxyAPIVersionUnescaped)
		var options *armmachinelearning.ConnectionRaiPolicyClientBeginCreateOptions
		if proxyAPIVersionParam != nil {
			options = &armmachinelearning.ConnectionRaiPolicyClientBeginCreateOptions{
				ProxyAPIVersion: proxyAPIVersionParam,
			}
		}
		respr, errRespr := c.srv.BeginCreate(req.Context(), resourceGroupNameParam, workspaceNameParam, connectionNameParam, raiPolicyNameParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		c.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		c.beginCreate.remove(req)
	}

	return resp, nil
}

func (c *ConnectionRaiPolicyServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/raiPolicies/(?P<raiPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		raiPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("raiPolicyName")])
		if err != nil {
			return nil, err
		}
		proxyAPIVersionUnescaped, err := url.QueryUnescape(qp.Get("proxy-api-version"))
		if err != nil {
			return nil, err
		}
		proxyAPIVersionParam := getOptional(proxyAPIVersionUnescaped)
		var options *armmachinelearning.ConnectionRaiPolicyClientBeginDeleteOptions
		if proxyAPIVersionParam != nil {
			options = &armmachinelearning.ConnectionRaiPolicyClientBeginDeleteOptions{
				ProxyAPIVersion: proxyAPIVersionParam,
			}
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, workspaceNameParam, connectionNameParam, raiPolicyNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *ConnectionRaiPolicyServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/raiPolicies/(?P<raiPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
	if err != nil {
		return nil, err
	}
	raiPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("raiPolicyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, connectionNameParam, raiPolicyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RaiPolicyPropertiesBasicResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ConnectionRaiPolicyServerTransport
var connectionRaiPolicyServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
