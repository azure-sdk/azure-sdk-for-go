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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
	"net/http"
	"net/url"
	"regexp"
)

// ServerSecurityAlertPoliciesServer is a fake server for instances of the armpostgresql.ServerSecurityAlertPoliciesClient type.
type ServerSecurityAlertPoliciesServer struct {
	// BeginCreateOrUpdate is the fake for method ServerSecurityAlertPoliciesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, serverName string, securityAlertPolicyName armpostgresql.SecurityAlertPolicyName, parameters armpostgresql.ServerSecurityAlertPolicy, options *armpostgresql.ServerSecurityAlertPoliciesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armpostgresql.ServerSecurityAlertPoliciesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ServerSecurityAlertPoliciesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, securityAlertPolicyName armpostgresql.SecurityAlertPolicyName, options *armpostgresql.ServerSecurityAlertPoliciesClientGetOptions) (resp azfake.Responder[armpostgresql.ServerSecurityAlertPoliciesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServerPager is the fake for method ServerSecurityAlertPoliciesClient.NewListByServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServerPager func(resourceGroupName string, serverName string, options *armpostgresql.ServerSecurityAlertPoliciesClientListByServerOptions) (resp azfake.PagerResponder[armpostgresql.ServerSecurityAlertPoliciesClientListByServerResponse])
}

// NewServerSecurityAlertPoliciesServerTransport creates a new instance of ServerSecurityAlertPoliciesServerTransport with the provided implementation.
// The returned ServerSecurityAlertPoliciesServerTransport instance is connected to an instance of armpostgresql.ServerSecurityAlertPoliciesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerSecurityAlertPoliciesServerTransport(srv *ServerSecurityAlertPoliciesServer) *ServerSecurityAlertPoliciesServerTransport {
	return &ServerSecurityAlertPoliciesServerTransport{
		srv:                  srv,
		beginCreateOrUpdate:  newTracker[azfake.PollerResponder[armpostgresql.ServerSecurityAlertPoliciesClientCreateOrUpdateResponse]](),
		newListByServerPager: newTracker[azfake.PagerResponder[armpostgresql.ServerSecurityAlertPoliciesClientListByServerResponse]](),
	}
}

// ServerSecurityAlertPoliciesServerTransport connects instances of armpostgresql.ServerSecurityAlertPoliciesClient to instances of ServerSecurityAlertPoliciesServer.
// Don't use this type directly, use NewServerSecurityAlertPoliciesServerTransport instead.
type ServerSecurityAlertPoliciesServerTransport struct {
	srv                  *ServerSecurityAlertPoliciesServer
	beginCreateOrUpdate  *tracker[azfake.PollerResponder[armpostgresql.ServerSecurityAlertPoliciesClientCreateOrUpdateResponse]]
	newListByServerPager *tracker[azfake.PagerResponder[armpostgresql.ServerSecurityAlertPoliciesClientListByServerResponse]]
}

// Do implements the policy.Transporter interface for ServerSecurityAlertPoliciesServerTransport.
func (s *ServerSecurityAlertPoliciesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServerSecurityAlertPoliciesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if serverSecurityAlertPoliciesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = serverSecurityAlertPoliciesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ServerSecurityAlertPoliciesClient.BeginCreateOrUpdate":
				res.resp, res.err = s.dispatchBeginCreateOrUpdate(req)
			case "ServerSecurityAlertPoliciesClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "ServerSecurityAlertPoliciesClient.NewListByServerPager":
				res.resp, res.err = s.dispatchNewListByServerPager(req)
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

func (s *ServerSecurityAlertPoliciesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := s.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityAlertPolicies/(?P<securityAlertPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armpostgresql.ServerSecurityAlertPolicy](req)
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
		securityAlertPolicyNameParam, err := parseWithCast(matches[regex.SubexpIndex("securityAlertPolicyName")], func(v string) (armpostgresql.SecurityAlertPolicyName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armpostgresql.SecurityAlertPolicyName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, serverNameParam, securityAlertPolicyNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		s.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		s.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (s *ServerSecurityAlertPoliciesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityAlertPolicies/(?P<securityAlertPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	securityAlertPolicyNameParam, err := parseWithCast(matches[regex.SubexpIndex("securityAlertPolicyName")], func(v string) (armpostgresql.SecurityAlertPolicyName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armpostgresql.SecurityAlertPolicyName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, securityAlertPolicyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServerSecurityAlertPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerSecurityAlertPoliciesServerTransport) dispatchNewListByServerPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServerPager not implemented")}
	}
	newListByServerPager := s.newListByServerPager.get(req)
	if newListByServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityAlertPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByServerPager(resourceGroupNameParam, serverNameParam, nil)
		newListByServerPager = &resp
		s.newListByServerPager.add(req, newListByServerPager)
		server.PagerResponderInjectNextLinks(newListByServerPager, req, func(page *armpostgresql.ServerSecurityAlertPoliciesClientListByServerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServerPager) {
		s.newListByServerPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ServerSecurityAlertPoliciesServerTransport
var serverSecurityAlertPoliciesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
