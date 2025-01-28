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

// DependencyMapControllerServer is a fake server for instances of the armmigrationdiscovery.DependencyMapControllerClient type.
type DependencyMapControllerServer struct {
	// BeginClientGroupMembers is the fake for method DependencyMapControllerClient.BeginClientGroupMembers
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginClientGroupMembers func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrationdiscovery.DependencyMapServiceMapextensionsClientGroupMembersRequest, options *armmigrationdiscovery.DependencyMapControllerClientBeginClientGroupMembersOptions) (resp azfake.PollerResponder[armmigrationdiscovery.DependencyMapControllerClientClientGroupMembersResponse], errResp azfake.ErrorResponder)

	// BeginExportDependencies is the fake for method DependencyMapControllerClient.BeginExportDependencies
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExportDependencies func(ctx context.Context, resourceGroupName string, siteName string, requestBody armmigrationdiscovery.DependencyMapServiceMapextensionsExportDependenciesRequest, options *armmigrationdiscovery.DependencyMapControllerClientBeginExportDependenciesOptions) (resp azfake.PollerResponder[armmigrationdiscovery.DependencyMapControllerClientExportDependenciesResponse], errResp azfake.ErrorResponder)

	// BeginGenerateCoarseMap is the fake for method DependencyMapControllerClient.BeginGenerateCoarseMap
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGenerateCoarseMap func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrationdiscovery.DependencyMapServiceMapextensionsScopeMapRequest, options *armmigrationdiscovery.DependencyMapControllerClientBeginGenerateCoarseMapOptions) (resp azfake.PollerResponder[armmigrationdiscovery.DependencyMapControllerClientGenerateCoarseMapResponse], errResp azfake.ErrorResponder)

	// BeginGenerateDetailedMap is the fake for method DependencyMapControllerClient.BeginGenerateDetailedMap
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGenerateDetailedMap func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrationdiscovery.DependencyMapServiceMapextensionsSingleMachineDetailedMapRequest, options *armmigrationdiscovery.DependencyMapControllerClientBeginGenerateDetailedMapOptions) (resp azfake.PollerResponder[armmigrationdiscovery.DependencyMapControllerClientGenerateDetailedMapResponse], errResp azfake.ErrorResponder)

	// BeginServerGroupMembers is the fake for method DependencyMapControllerClient.BeginServerGroupMembers
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginServerGroupMembers func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrationdiscovery.DependencyMapServiceMapextensionsServerGroupMembersRequest, options *armmigrationdiscovery.DependencyMapControllerClientBeginServerGroupMembersOptions) (resp azfake.PollerResponder[armmigrationdiscovery.DependencyMapControllerClientServerGroupMembersResponse], errResp azfake.ErrorResponder)
}

// NewDependencyMapControllerServerTransport creates a new instance of DependencyMapControllerServerTransport with the provided implementation.
// The returned DependencyMapControllerServerTransport instance is connected to an instance of armmigrationdiscovery.DependencyMapControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDependencyMapControllerServerTransport(srv *DependencyMapControllerServer) *DependencyMapControllerServerTransport {
	return &DependencyMapControllerServerTransport{
		srv:                      srv,
		beginClientGroupMembers:  newTracker[azfake.PollerResponder[armmigrationdiscovery.DependencyMapControllerClientClientGroupMembersResponse]](),
		beginExportDependencies:  newTracker[azfake.PollerResponder[armmigrationdiscovery.DependencyMapControllerClientExportDependenciesResponse]](),
		beginGenerateCoarseMap:   newTracker[azfake.PollerResponder[armmigrationdiscovery.DependencyMapControllerClientGenerateCoarseMapResponse]](),
		beginGenerateDetailedMap: newTracker[azfake.PollerResponder[armmigrationdiscovery.DependencyMapControllerClientGenerateDetailedMapResponse]](),
		beginServerGroupMembers:  newTracker[azfake.PollerResponder[armmigrationdiscovery.DependencyMapControllerClientServerGroupMembersResponse]](),
	}
}

// DependencyMapControllerServerTransport connects instances of armmigrationdiscovery.DependencyMapControllerClient to instances of DependencyMapControllerServer.
// Don't use this type directly, use NewDependencyMapControllerServerTransport instead.
type DependencyMapControllerServerTransport struct {
	srv                      *DependencyMapControllerServer
	beginClientGroupMembers  *tracker[azfake.PollerResponder[armmigrationdiscovery.DependencyMapControllerClientClientGroupMembersResponse]]
	beginExportDependencies  *tracker[azfake.PollerResponder[armmigrationdiscovery.DependencyMapControllerClientExportDependenciesResponse]]
	beginGenerateCoarseMap   *tracker[azfake.PollerResponder[armmigrationdiscovery.DependencyMapControllerClientGenerateCoarseMapResponse]]
	beginGenerateDetailedMap *tracker[azfake.PollerResponder[armmigrationdiscovery.DependencyMapControllerClientGenerateDetailedMapResponse]]
	beginServerGroupMembers  *tracker[azfake.PollerResponder[armmigrationdiscovery.DependencyMapControllerClientServerGroupMembersResponse]]
}

// Do implements the policy.Transporter interface for DependencyMapControllerServerTransport.
func (d *DependencyMapControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DependencyMapControllerClient.BeginClientGroupMembers":
		resp, err = d.dispatchBeginClientGroupMembers(req)
	case "DependencyMapControllerClient.BeginExportDependencies":
		resp, err = d.dispatchBeginExportDependencies(req)
	case "DependencyMapControllerClient.BeginGenerateCoarseMap":
		resp, err = d.dispatchBeginGenerateCoarseMap(req)
	case "DependencyMapControllerClient.BeginGenerateDetailedMap":
		resp, err = d.dispatchBeginGenerateDetailedMap(req)
	case "DependencyMapControllerClient.BeginServerGroupMembers":
		resp, err = d.dispatchBeginServerGroupMembers(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DependencyMapControllerServerTransport) dispatchBeginClientGroupMembers(req *http.Request) (*http.Response, error) {
	if d.srv.BeginClientGroupMembers == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginClientGroupMembers not implemented")}
	}
	beginClientGroupMembers := d.beginClientGroupMembers.get(req)
	if beginClientGroupMembers == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clientGroupMembers`
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
		respr, errRespr := d.srv.BeginClientGroupMembers(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginClientGroupMembers = &respr
		d.beginClientGroupMembers.add(req, beginClientGroupMembers)
	}

	resp, err := server.PollerResponderNext(beginClientGroupMembers, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginClientGroupMembers.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginClientGroupMembers) {
		d.beginClientGroupMembers.remove(req)
	}

	return resp, nil
}

func (d *DependencyMapControllerServerTransport) dispatchBeginExportDependencies(req *http.Request) (*http.Response, error) {
	if d.srv.BeginExportDependencies == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExportDependencies not implemented")}
	}
	beginExportDependencies := d.beginExportDependencies.get(req)
	if beginExportDependencies == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportDependencies`
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
		respr, errRespr := d.srv.BeginExportDependencies(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginExportDependencies = &respr
		d.beginExportDependencies.add(req, beginExportDependencies)
	}

	resp, err := server.PollerResponderNext(beginExportDependencies, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginExportDependencies.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExportDependencies) {
		d.beginExportDependencies.remove(req)
	}

	return resp, nil
}

func (d *DependencyMapControllerServerTransport) dispatchBeginGenerateCoarseMap(req *http.Request) (*http.Response, error) {
	if d.srv.BeginGenerateCoarseMap == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGenerateCoarseMap not implemented")}
	}
	beginGenerateCoarseMap := d.beginGenerateCoarseMap.get(req)
	if beginGenerateCoarseMap == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/generateCoarseMap`
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
		respr, errRespr := d.srv.BeginGenerateCoarseMap(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGenerateCoarseMap = &respr
		d.beginGenerateCoarseMap.add(req, beginGenerateCoarseMap)
	}

	resp, err := server.PollerResponderNext(beginGenerateCoarseMap, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginGenerateCoarseMap.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGenerateCoarseMap) {
		d.beginGenerateCoarseMap.remove(req)
	}

	return resp, nil
}

func (d *DependencyMapControllerServerTransport) dispatchBeginGenerateDetailedMap(req *http.Request) (*http.Response, error) {
	if d.srv.BeginGenerateDetailedMap == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGenerateDetailedMap not implemented")}
	}
	beginGenerateDetailedMap := d.beginGenerateDetailedMap.get(req)
	if beginGenerateDetailedMap == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/generateDetailedMap`
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
		respr, errRespr := d.srv.BeginGenerateDetailedMap(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGenerateDetailedMap = &respr
		d.beginGenerateDetailedMap.add(req, beginGenerateDetailedMap)
	}

	resp, err := server.PollerResponderNext(beginGenerateDetailedMap, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginGenerateDetailedMap.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGenerateDetailedMap) {
		d.beginGenerateDetailedMap.remove(req)
	}

	return resp, nil
}

func (d *DependencyMapControllerServerTransport) dispatchBeginServerGroupMembers(req *http.Request) (*http.Response, error) {
	if d.srv.BeginServerGroupMembers == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginServerGroupMembers not implemented")}
	}
	beginServerGroupMembers := d.beginServerGroupMembers.get(req)
	if beginServerGroupMembers == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serverGroupMembers`
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
		respr, errRespr := d.srv.BeginServerGroupMembers(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginServerGroupMembers = &respr
		d.beginServerGroupMembers.add(req, beginServerGroupMembers)
	}

	resp, err := server.PollerResponderNext(beginServerGroupMembers, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginServerGroupMembers.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginServerGroupMembers) {
		d.beginServerGroupMembers.remove(req)
	}

	return resp, nil
}
