//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
	"net/http"
	"net/url"
	"regexp"
)

// Server is a fake server for instances of the armsecurityinsights.Client type.
type Server struct {
	// ListGeodataByIP is the fake for method Client.ListGeodataByIP
	// HTTP status codes to indicate success: http.StatusOK
	ListGeodataByIP func(ctx context.Context, resourceGroupName string, workspaceName string, enrichmentType armsecurityinsights.EnrichmentType, ipAddressBody armsecurityinsights.EnrichmentIPAddressBody, options *armsecurityinsights.ClientListGeodataByIPOptions) (resp azfake.Responder[armsecurityinsights.ClientListGeodataByIPResponse], errResp azfake.ErrorResponder)

	// ListWhoisByDomain is the fake for method Client.ListWhoisByDomain
	// HTTP status codes to indicate success: http.StatusOK
	ListWhoisByDomain func(ctx context.Context, resourceGroupName string, workspaceName string, enrichmentType armsecurityinsights.EnrichmentType, domainBody armsecurityinsights.EnrichmentDomainBody, options *armsecurityinsights.ClientListWhoisByDomainOptions) (resp azfake.Responder[armsecurityinsights.ClientListWhoisByDomainResponse], errResp azfake.ErrorResponder)
}

// NewServerTransport creates a new instance of ServerTransport with the provided implementation.
// The returned ServerTransport instance is connected to an instance of armsecurityinsights.Client via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerTransport(srv *Server) *ServerTransport {
	return &ServerTransport{srv: srv}
}

// ServerTransport connects instances of armsecurityinsights.Client to instances of Server.
// Don't use this type directly, use NewServerTransport instead.
type ServerTransport struct {
	srv *Server
}

// Do implements the policy.Transporter interface for ServerTransport.
func (s *ServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "Client.ListGeodataByIP":
		resp, err = s.dispatchListGeodataByIP(req)
	case "Client.ListWhoisByDomain":
		resp, err = s.dispatchListWhoisByDomain(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServerTransport) dispatchListGeodataByIP(req *http.Request) (*http.Response, error) {
	if s.srv.ListGeodataByIP == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListGeodataByIP not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/enrichment/(?P<enrichmentType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listGeodataByIp`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsecurityinsights.EnrichmentIPAddressBody](req)
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
	enrichmentTypeParam, err := parseWithCast(matches[regex.SubexpIndex("enrichmentType")], func(v string) (armsecurityinsights.EnrichmentType, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsecurityinsights.EnrichmentType(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.ListGeodataByIP(req.Context(), resourceGroupNameParam, workspaceNameParam, enrichmentTypeParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EnrichmentIPGeodata, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchListWhoisByDomain(req *http.Request) (*http.Response, error) {
	if s.srv.ListWhoisByDomain == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListWhoisByDomain not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/enrichment/(?P<enrichmentType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listWhoisByDomain`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsecurityinsights.EnrichmentDomainBody](req)
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
	enrichmentTypeParam, err := parseWithCast(matches[regex.SubexpIndex("enrichmentType")], func(v string) (armsecurityinsights.EnrichmentType, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsecurityinsights.EnrichmentType(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.ListWhoisByDomain(req.Context(), resourceGroupNameParam, workspaceNameParam, enrichmentTypeParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EnrichmentDomainWhois, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
