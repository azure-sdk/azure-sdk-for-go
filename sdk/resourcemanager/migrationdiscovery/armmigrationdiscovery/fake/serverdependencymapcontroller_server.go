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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationdiscovery/armmigrationdiscovery"
	"net/http"
	"net/url"
	"regexp"
)

// ServerDependencyMapControllerServer is a fake server for instances of the armmigrationdiscovery.ServerDependencyMapControllerClient type.
type ServerDependencyMapControllerServer struct {
	// BeginClientGroupMembers is the fake for method ServerDependencyMapControllerClient.BeginClientGroupMembers
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginClientGroupMembers func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrationdiscovery.DependencyMapServiceMapextensionsClientGroupMembersRequest, options *armmigrationdiscovery.ServerDependencyMapControllerClientBeginClientGroupMembersOptions) (resp azfake.PollerResponder[armmigrationdiscovery.ServerDependencyMapControllerClientClientGroupMembersResponse], errResp azfake.ErrorResponder)

	// BeginExportDependencies is the fake for method ServerDependencyMapControllerClient.BeginExportDependencies
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExportDependencies func(ctx context.Context, resourceGroupName string, siteName string, requestBody armmigrationdiscovery.DependencyMapServiceMapextensionsExportDependenciesRequest, options *armmigrationdiscovery.ServerDependencyMapControllerClientBeginExportDependenciesOptions) (resp azfake.PollerResponder[armmigrationdiscovery.ServerDependencyMapControllerClientExportDependenciesResponse], errResp azfake.ErrorResponder)

	// BeginGenerateCoarseMap is the fake for method ServerDependencyMapControllerClient.BeginGenerateCoarseMap
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGenerateCoarseMap func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrationdiscovery.DependencyMapServiceMapextensionsScopeMapRequest, options *armmigrationdiscovery.ServerDependencyMapControllerClientBeginGenerateCoarseMapOptions) (resp azfake.PollerResponder[armmigrationdiscovery.ServerDependencyMapControllerClientGenerateCoarseMapResponse], errResp azfake.ErrorResponder)

	// BeginGenerateDetailedMap is the fake for method ServerDependencyMapControllerClient.BeginGenerateDetailedMap
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGenerateDetailedMap func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrationdiscovery.DependencyMapServiceMapextensionsSingleMachineDetailedMapRequest, options *armmigrationdiscovery.ServerDependencyMapControllerClientBeginGenerateDetailedMapOptions) (resp azfake.PollerResponder[armmigrationdiscovery.ServerDependencyMapControllerClientGenerateDetailedMapResponse], errResp azfake.ErrorResponder)

	// BeginServerGroupMembers is the fake for method ServerDependencyMapControllerClient.BeginServerGroupMembers
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginServerGroupMembers func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrationdiscovery.DependencyMapServiceMapextensionsServerGroupMembersRequest, options *armmigrationdiscovery.ServerDependencyMapControllerClientBeginServerGroupMembersOptions) (resp azfake.PollerResponder[armmigrationdiscovery.ServerDependencyMapControllerClientServerGroupMembersResponse], errResp azfake.ErrorResponder)
}

// NewServerDependencyMapControllerServerTransport creates a new instance of ServerDependencyMapControllerServerTransport with the provided implementation.
// The returned ServerDependencyMapControllerServerTransport instance is connected to an instance of armmigrationdiscovery.ServerDependencyMapControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerDependencyMapControllerServerTransport(srv *ServerDependencyMapControllerServer) *ServerDependencyMapControllerServerTransport {
	return &ServerDependencyMapControllerServerTransport{
		srv:                      srv,
		beginClientGroupMembers:  newTracker[azfake.PollerResponder[armmigrationdiscovery.ServerDependencyMapControllerClientClientGroupMembersResponse]](),
		beginExportDependencies:  newTracker[azfake.PollerResponder[armmigrationdiscovery.ServerDependencyMapControllerClientExportDependenciesResponse]](),
		beginGenerateCoarseMap:   newTracker[azfake.PollerResponder[armmigrationdiscovery.ServerDependencyMapControllerClientGenerateCoarseMapResponse]](),
		beginGenerateDetailedMap: newTracker[azfake.PollerResponder[armmigrationdiscovery.ServerDependencyMapControllerClientGenerateDetailedMapResponse]](),
		beginServerGroupMembers:  newTracker[azfake.PollerResponder[armmigrationdiscovery.ServerDependencyMapControllerClientServerGroupMembersResponse]](),
	}
}

// ServerDependencyMapControllerServerTransport connects instances of armmigrationdiscovery.ServerDependencyMapControllerClient to instances of ServerDependencyMapControllerServer.
// Don't use this type directly, use NewServerDependencyMapControllerServerTransport instead.
type ServerDependencyMapControllerServerTransport struct {
	srv                      *ServerDependencyMapControllerServer
	beginClientGroupMembers  *tracker[azfake.PollerResponder[armmigrationdiscovery.ServerDependencyMapControllerClientClientGroupMembersResponse]]
	beginExportDependencies  *tracker[azfake.PollerResponder[armmigrationdiscovery.ServerDependencyMapControllerClientExportDependenciesResponse]]
	beginGenerateCoarseMap   *tracker[azfake.PollerResponder[armmigrationdiscovery.ServerDependencyMapControllerClientGenerateCoarseMapResponse]]
	beginGenerateDetailedMap *tracker[azfake.PollerResponder[armmigrationdiscovery.ServerDependencyMapControllerClientGenerateDetailedMapResponse]]
	beginServerGroupMembers  *tracker[azfake.PollerResponder[armmigrationdiscovery.ServerDependencyMapControllerClientServerGroupMembersResponse]]
}

// Do implements the policy.Transporter interface for ServerDependencyMapControllerServerTransport.
func (s *ServerDependencyMapControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ServerDependencyMapControllerClient.BeginClientGroupMembers":
		resp, err = s.dispatchBeginClientGroupMembers(req)
	case "ServerDependencyMapControllerClient.BeginExportDependencies":
		resp, err = s.dispatchBeginExportDependencies(req)
	case "ServerDependencyMapControllerClient.BeginGenerateCoarseMap":
		resp, err = s.dispatchBeginGenerateCoarseMap(req)
	case "ServerDependencyMapControllerClient.BeginGenerateDetailedMap":
		resp, err = s.dispatchBeginGenerateDetailedMap(req)
	case "ServerDependencyMapControllerClient.BeginServerGroupMembers":
		resp, err = s.dispatchBeginServerGroupMembers(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServerDependencyMapControllerServerTransport) dispatchBeginClientGroupMembers(req *http.Request) (*http.Response, error) {
	if s.srv.BeginClientGroupMembers == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginClientGroupMembers not implemented")}
	}
	beginClientGroupMembers := s.beginClientGroupMembers.get(req)
	if beginClientGroupMembers == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clientGroupMembers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrationdiscovery.DependencyMapServiceMapextensionsClientGroupMembersRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginClientGroupMembers(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginClientGroupMembers = &respr
		s.beginClientGroupMembers.add(req, beginClientGroupMembers)
	}

	resp, err := server.PollerResponderNext(beginClientGroupMembers, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginClientGroupMembers.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginClientGroupMembers) {
		s.beginClientGroupMembers.remove(req)
	}

	return resp, nil
}

func (s *ServerDependencyMapControllerServerTransport) dispatchBeginExportDependencies(req *http.Request) (*http.Response, error) {
	if s.srv.BeginExportDependencies == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExportDependencies not implemented")}
	}
	beginExportDependencies := s.beginExportDependencies.get(req)
	if beginExportDependencies == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportDependencies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrationdiscovery.DependencyMapServiceMapextensionsExportDependenciesRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginExportDependencies(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginExportDependencies = &respr
		s.beginExportDependencies.add(req, beginExportDependencies)
	}

	resp, err := server.PollerResponderNext(beginExportDependencies, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginExportDependencies.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExportDependencies) {
		s.beginExportDependencies.remove(req)
	}

	return resp, nil
}

func (s *ServerDependencyMapControllerServerTransport) dispatchBeginGenerateCoarseMap(req *http.Request) (*http.Response, error) {
	if s.srv.BeginGenerateCoarseMap == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGenerateCoarseMap not implemented")}
	}
	beginGenerateCoarseMap := s.beginGenerateCoarseMap.get(req)
	if beginGenerateCoarseMap == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/generateCoarseMap`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrationdiscovery.DependencyMapServiceMapextensionsScopeMapRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginGenerateCoarseMap(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGenerateCoarseMap = &respr
		s.beginGenerateCoarseMap.add(req, beginGenerateCoarseMap)
	}

	resp, err := server.PollerResponderNext(beginGenerateCoarseMap, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginGenerateCoarseMap.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGenerateCoarseMap) {
		s.beginGenerateCoarseMap.remove(req)
	}

	return resp, nil
}

func (s *ServerDependencyMapControllerServerTransport) dispatchBeginGenerateDetailedMap(req *http.Request) (*http.Response, error) {
	if s.srv.BeginGenerateDetailedMap == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGenerateDetailedMap not implemented")}
	}
	beginGenerateDetailedMap := s.beginGenerateDetailedMap.get(req)
	if beginGenerateDetailedMap == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/generateDetailedMap`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrationdiscovery.DependencyMapServiceMapextensionsSingleMachineDetailedMapRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginGenerateDetailedMap(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGenerateDetailedMap = &respr
		s.beginGenerateDetailedMap.add(req, beginGenerateDetailedMap)
	}

	resp, err := server.PollerResponderNext(beginGenerateDetailedMap, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginGenerateDetailedMap.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGenerateDetailedMap) {
		s.beginGenerateDetailedMap.remove(req)
	}

	return resp, nil
}

func (s *ServerDependencyMapControllerServerTransport) dispatchBeginServerGroupMembers(req *http.Request) (*http.Response, error) {
	if s.srv.BeginServerGroupMembers == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginServerGroupMembers not implemented")}
	}
	beginServerGroupMembers := s.beginServerGroupMembers.get(req)
	if beginServerGroupMembers == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serverGroupMembers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrationdiscovery.DependencyMapServiceMapextensionsServerGroupMembersRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginServerGroupMembers(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginServerGroupMembers = &respr
		s.beginServerGroupMembers.add(req, beginServerGroupMembers)
	}

	resp, err := server.PollerResponderNext(beginServerGroupMembers, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginServerGroupMembers.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginServerGroupMembers) {
		s.beginServerGroupMembers.remove(req)
	}

	return resp, nil
}
