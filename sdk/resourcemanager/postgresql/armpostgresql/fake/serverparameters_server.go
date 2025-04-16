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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
	"net/http"
	"net/url"
	"regexp"
)

// ServerParametersServer is a fake server for instances of the armpostgresql.ServerParametersClient type.
type ServerParametersServer struct {
	// BeginListUpdateConfigurations is the fake for method ServerParametersClient.BeginListUpdateConfigurations
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginListUpdateConfigurations func(ctx context.Context, resourceGroupName string, serverName string, value armpostgresql.ConfigurationListResult, options *armpostgresql.ServerParametersClientBeginListUpdateConfigurationsOptions) (resp azfake.PollerResponder[armpostgresql.ServerParametersClientListUpdateConfigurationsResponse], errResp azfake.ErrorResponder)
}

// NewServerParametersServerTransport creates a new instance of ServerParametersServerTransport with the provided implementation.
// The returned ServerParametersServerTransport instance is connected to an instance of armpostgresql.ServerParametersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerParametersServerTransport(srv *ServerParametersServer) *ServerParametersServerTransport {
	return &ServerParametersServerTransport{
		srv:                           srv,
		beginListUpdateConfigurations: newTracker[azfake.PollerResponder[armpostgresql.ServerParametersClientListUpdateConfigurationsResponse]](),
	}
}

// ServerParametersServerTransport connects instances of armpostgresql.ServerParametersClient to instances of ServerParametersServer.
// Don't use this type directly, use NewServerParametersServerTransport instead.
type ServerParametersServerTransport struct {
	srv                           *ServerParametersServer
	beginListUpdateConfigurations *tracker[azfake.PollerResponder[armpostgresql.ServerParametersClientListUpdateConfigurationsResponse]]
}

// Do implements the policy.Transporter interface for ServerParametersServerTransport.
func (s *ServerParametersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServerParametersServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if serverParametersServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = serverParametersServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ServerParametersClient.BeginListUpdateConfigurations":
				res.resp, res.err = s.dispatchBeginListUpdateConfigurations(req)
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

func (s *ServerParametersServerTransport) dispatchBeginListUpdateConfigurations(req *http.Request) (*http.Response, error) {
	if s.srv.BeginListUpdateConfigurations == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginListUpdateConfigurations not implemented")}
	}
	beginListUpdateConfigurations := s.beginListUpdateConfigurations.get(req)
	if beginListUpdateConfigurations == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armpostgresql.ConfigurationListResult](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginListUpdateConfigurations(req.Context(), resourceGroupNameParam, serverNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginListUpdateConfigurations = &respr
		s.beginListUpdateConfigurations.add(req, beginListUpdateConfigurations)
	}

	resp, err := server.PollerResponderNext(beginListUpdateConfigurations, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginListUpdateConfigurations.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginListUpdateConfigurations) {
		s.beginListUpdateConfigurations.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ServerParametersServerTransport
var serverParametersServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
