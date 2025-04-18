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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"net/http"
	"net/url"
	"regexp"
)

// SQLPoolGeoBackupPoliciesServer is a fake server for instances of the armsynapse.SQLPoolGeoBackupPoliciesClient type.
type SQLPoolGeoBackupPoliciesServer struct {
	// CreateOrUpdate is the fake for method SQLPoolGeoBackupPoliciesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, geoBackupPolicyName armsynapse.GeoBackupPolicyName, parameters armsynapse.GeoBackupPolicy, options *armsynapse.SQLPoolGeoBackupPoliciesClientCreateOrUpdateOptions) (resp azfake.Responder[armsynapse.SQLPoolGeoBackupPoliciesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SQLPoolGeoBackupPoliciesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, geoBackupPolicyName armsynapse.GeoBackupPolicyName, options *armsynapse.SQLPoolGeoBackupPoliciesClientGetOptions) (resp azfake.Responder[armsynapse.SQLPoolGeoBackupPoliciesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SQLPoolGeoBackupPoliciesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, sqlPoolName string, options *armsynapse.SQLPoolGeoBackupPoliciesClientListOptions) (resp azfake.PagerResponder[armsynapse.SQLPoolGeoBackupPoliciesClientListResponse])
}

// NewSQLPoolGeoBackupPoliciesServerTransport creates a new instance of SQLPoolGeoBackupPoliciesServerTransport with the provided implementation.
// The returned SQLPoolGeoBackupPoliciesServerTransport instance is connected to an instance of armsynapse.SQLPoolGeoBackupPoliciesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSQLPoolGeoBackupPoliciesServerTransport(srv *SQLPoolGeoBackupPoliciesServer) *SQLPoolGeoBackupPoliciesServerTransport {
	return &SQLPoolGeoBackupPoliciesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsynapse.SQLPoolGeoBackupPoliciesClientListResponse]](),
	}
}

// SQLPoolGeoBackupPoliciesServerTransport connects instances of armsynapse.SQLPoolGeoBackupPoliciesClient to instances of SQLPoolGeoBackupPoliciesServer.
// Don't use this type directly, use NewSQLPoolGeoBackupPoliciesServerTransport instead.
type SQLPoolGeoBackupPoliciesServerTransport struct {
	srv          *SQLPoolGeoBackupPoliciesServer
	newListPager *tracker[azfake.PagerResponder[armsynapse.SQLPoolGeoBackupPoliciesClientListResponse]]
}

// Do implements the policy.Transporter interface for SQLPoolGeoBackupPoliciesServerTransport.
func (s *SQLPoolGeoBackupPoliciesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SQLPoolGeoBackupPoliciesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if sqlPoolGeoBackupPoliciesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = sqlPoolGeoBackupPoliciesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SQLPoolGeoBackupPoliciesClient.CreateOrUpdate":
				res.resp, res.err = s.dispatchCreateOrUpdate(req)
			case "SQLPoolGeoBackupPoliciesClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "SQLPoolGeoBackupPoliciesClient.NewListPager":
				res.resp, res.err = s.dispatchNewListPager(req)
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

func (s *SQLPoolGeoBackupPoliciesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlPools/(?P<sqlPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/geoBackupPolicies/(?P<geoBackupPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsynapse.GeoBackupPolicy](req)
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
	sqlPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlPoolName")])
	if err != nil {
		return nil, err
	}
	geoBackupPolicyNameParam, err := parseWithCast(matches[regex.SubexpIndex("geoBackupPolicyName")], func(v string) (armsynapse.GeoBackupPolicyName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsynapse.GeoBackupPolicyName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, sqlPoolNameParam, geoBackupPolicyNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GeoBackupPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SQLPoolGeoBackupPoliciesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlPools/(?P<sqlPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/geoBackupPolicies/(?P<geoBackupPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	sqlPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlPoolName")])
	if err != nil {
		return nil, err
	}
	geoBackupPolicyNameParam, err := parseWithCast(matches[regex.SubexpIndex("geoBackupPolicyName")], func(v string) (armsynapse.GeoBackupPolicyName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsynapse.GeoBackupPolicyName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, sqlPoolNameParam, geoBackupPolicyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GeoBackupPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SQLPoolGeoBackupPoliciesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlPools/(?P<sqlPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/geoBackupPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
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
		sqlPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlPoolName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, sqlPoolNameParam, nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to SQLPoolGeoBackupPoliciesServerTransport
var sqlPoolGeoBackupPoliciesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
