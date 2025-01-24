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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
	"net/http"
	"net/url"
	"regexp"
)

// HypervDependencyMapControllerServer is a fake server for instances of the armmigrate.HypervDependencyMapControllerClient type.
type HypervDependencyMapControllerServer struct {
	// BeginClientGroupMembers is the fake for method HypervDependencyMapControllerClient.BeginClientGroupMembers
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginClientGroupMembers func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrate.DependencyMapServiceMapextensionsClientGroupMembersRequest, options *armmigrate.HypervDependencyMapControllerClientBeginClientGroupMembersOptions) (resp azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientClientGroupMembersResponse], errResp azfake.ErrorResponder)

	// BeginExportDependencies is the fake for method HypervDependencyMapControllerClient.BeginExportDependencies
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExportDependencies func(ctx context.Context, resourceGroupName string, siteName string, requestBody armmigrate.DependencyMapServiceMapextensionsExportDependenciesRequest, options *armmigrate.HypervDependencyMapControllerClientBeginExportDependenciesOptions) (resp azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientExportDependenciesResponse], errResp azfake.ErrorResponder)

	// BeginGenerateCoarseMap is the fake for method HypervDependencyMapControllerClient.BeginGenerateCoarseMap
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGenerateCoarseMap func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrate.DependencyMapServiceMapextensionsScopeMapRequest, options *armmigrate.HypervDependencyMapControllerClientBeginGenerateCoarseMapOptions) (resp azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientGenerateCoarseMapResponse], errResp azfake.ErrorResponder)

	// BeginGenerateDetailedMap is the fake for method HypervDependencyMapControllerClient.BeginGenerateDetailedMap
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGenerateDetailedMap func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrate.DependencyMapServiceMapextensionsSingleMachineDetailedMapRequest, options *armmigrate.HypervDependencyMapControllerClientBeginGenerateDetailedMapOptions) (resp azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientGenerateDetailedMapResponse], errResp azfake.ErrorResponder)

	// BeginServerGroupMembers is the fake for method HypervDependencyMapControllerClient.BeginServerGroupMembers
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginServerGroupMembers func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrate.DependencyMapServiceMapextensionsServerGroupMembersRequest, options *armmigrate.HypervDependencyMapControllerClientBeginServerGroupMembersOptions) (resp azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientServerGroupMembersResponse], errResp azfake.ErrorResponder)

	// BeginUpdateDependencyMapStatus is the fake for method HypervDependencyMapControllerClient.BeginUpdateDependencyMapStatus
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateDependencyMapStatus func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrate.UpdateMachineDepMapStatus, options *armmigrate.HypervDependencyMapControllerClientBeginUpdateDependencyMapStatusOptions) (resp azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientUpdateDependencyMapStatusResponse], errResp azfake.ErrorResponder)
}

// NewHypervDependencyMapControllerServerTransport creates a new instance of HypervDependencyMapControllerServerTransport with the provided implementation.
// The returned HypervDependencyMapControllerServerTransport instance is connected to an instance of armmigrate.HypervDependencyMapControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHypervDependencyMapControllerServerTransport(srv *HypervDependencyMapControllerServer) *HypervDependencyMapControllerServerTransport {
	return &HypervDependencyMapControllerServerTransport{
		srv:                            srv,
		beginClientGroupMembers:        newTracker[azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientClientGroupMembersResponse]](),
		beginExportDependencies:        newTracker[azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientExportDependenciesResponse]](),
		beginGenerateCoarseMap:         newTracker[azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientGenerateCoarseMapResponse]](),
		beginGenerateDetailedMap:       newTracker[azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientGenerateDetailedMapResponse]](),
		beginServerGroupMembers:        newTracker[azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientServerGroupMembersResponse]](),
		beginUpdateDependencyMapStatus: newTracker[azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientUpdateDependencyMapStatusResponse]](),
	}
}

// HypervDependencyMapControllerServerTransport connects instances of armmigrate.HypervDependencyMapControllerClient to instances of HypervDependencyMapControllerServer.
// Don't use this type directly, use NewHypervDependencyMapControllerServerTransport instead.
type HypervDependencyMapControllerServerTransport struct {
	srv                            *HypervDependencyMapControllerServer
	beginClientGroupMembers        *tracker[azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientClientGroupMembersResponse]]
	beginExportDependencies        *tracker[azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientExportDependenciesResponse]]
	beginGenerateCoarseMap         *tracker[azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientGenerateCoarseMapResponse]]
	beginGenerateDetailedMap       *tracker[azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientGenerateDetailedMapResponse]]
	beginServerGroupMembers        *tracker[azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientServerGroupMembersResponse]]
	beginUpdateDependencyMapStatus *tracker[azfake.PollerResponder[armmigrate.HypervDependencyMapControllerClientUpdateDependencyMapStatusResponse]]
}

// Do implements the policy.Transporter interface for HypervDependencyMapControllerServerTransport.
func (h *HypervDependencyMapControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HypervDependencyMapControllerClient.BeginClientGroupMembers":
		resp, err = h.dispatchBeginClientGroupMembers(req)
	case "HypervDependencyMapControllerClient.BeginExportDependencies":
		resp, err = h.dispatchBeginExportDependencies(req)
	case "HypervDependencyMapControllerClient.BeginGenerateCoarseMap":
		resp, err = h.dispatchBeginGenerateCoarseMap(req)
	case "HypervDependencyMapControllerClient.BeginGenerateDetailedMap":
		resp, err = h.dispatchBeginGenerateDetailedMap(req)
	case "HypervDependencyMapControllerClient.BeginServerGroupMembers":
		resp, err = h.dispatchBeginServerGroupMembers(req)
	case "HypervDependencyMapControllerClient.BeginUpdateDependencyMapStatus":
		resp, err = h.dispatchBeginUpdateDependencyMapStatus(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HypervDependencyMapControllerServerTransport) dispatchBeginClientGroupMembers(req *http.Request) (*http.Response, error) {
	if h.srv.BeginClientGroupMembers == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginClientGroupMembers not implemented")}
	}
	beginClientGroupMembers := h.beginClientGroupMembers.get(req)
	if beginClientGroupMembers == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clientGroupMembers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrate.DependencyMapServiceMapextensionsClientGroupMembersRequest](req)
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
		respr, errRespr := h.srv.BeginClientGroupMembers(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginClientGroupMembers = &respr
		h.beginClientGroupMembers.add(req, beginClientGroupMembers)
	}

	resp, err := server.PollerResponderNext(beginClientGroupMembers, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		h.beginClientGroupMembers.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginClientGroupMembers) {
		h.beginClientGroupMembers.remove(req)
	}

	return resp, nil
}

func (h *HypervDependencyMapControllerServerTransport) dispatchBeginExportDependencies(req *http.Request) (*http.Response, error) {
	if h.srv.BeginExportDependencies == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExportDependencies not implemented")}
	}
	beginExportDependencies := h.beginExportDependencies.get(req)
	if beginExportDependencies == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportDependencies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrate.DependencyMapServiceMapextensionsExportDependenciesRequest](req)
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
		respr, errRespr := h.srv.BeginExportDependencies(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginExportDependencies = &respr
		h.beginExportDependencies.add(req, beginExportDependencies)
	}

	resp, err := server.PollerResponderNext(beginExportDependencies, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		h.beginExportDependencies.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExportDependencies) {
		h.beginExportDependencies.remove(req)
	}

	return resp, nil
}

func (h *HypervDependencyMapControllerServerTransport) dispatchBeginGenerateCoarseMap(req *http.Request) (*http.Response, error) {
	if h.srv.BeginGenerateCoarseMap == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGenerateCoarseMap not implemented")}
	}
	beginGenerateCoarseMap := h.beginGenerateCoarseMap.get(req)
	if beginGenerateCoarseMap == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/generateCoarseMap`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrate.DependencyMapServiceMapextensionsScopeMapRequest](req)
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
		respr, errRespr := h.srv.BeginGenerateCoarseMap(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGenerateCoarseMap = &respr
		h.beginGenerateCoarseMap.add(req, beginGenerateCoarseMap)
	}

	resp, err := server.PollerResponderNext(beginGenerateCoarseMap, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		h.beginGenerateCoarseMap.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGenerateCoarseMap) {
		h.beginGenerateCoarseMap.remove(req)
	}

	return resp, nil
}

func (h *HypervDependencyMapControllerServerTransport) dispatchBeginGenerateDetailedMap(req *http.Request) (*http.Response, error) {
	if h.srv.BeginGenerateDetailedMap == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGenerateDetailedMap not implemented")}
	}
	beginGenerateDetailedMap := h.beginGenerateDetailedMap.get(req)
	if beginGenerateDetailedMap == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/generateDetailedMap`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrate.DependencyMapServiceMapextensionsSingleMachineDetailedMapRequest](req)
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
		respr, errRespr := h.srv.BeginGenerateDetailedMap(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGenerateDetailedMap = &respr
		h.beginGenerateDetailedMap.add(req, beginGenerateDetailedMap)
	}

	resp, err := server.PollerResponderNext(beginGenerateDetailedMap, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		h.beginGenerateDetailedMap.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGenerateDetailedMap) {
		h.beginGenerateDetailedMap.remove(req)
	}

	return resp, nil
}

func (h *HypervDependencyMapControllerServerTransport) dispatchBeginServerGroupMembers(req *http.Request) (*http.Response, error) {
	if h.srv.BeginServerGroupMembers == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginServerGroupMembers not implemented")}
	}
	beginServerGroupMembers := h.beginServerGroupMembers.get(req)
	if beginServerGroupMembers == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serverGroupMembers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrate.DependencyMapServiceMapextensionsServerGroupMembersRequest](req)
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
		respr, errRespr := h.srv.BeginServerGroupMembers(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginServerGroupMembers = &respr
		h.beginServerGroupMembers.add(req, beginServerGroupMembers)
	}

	resp, err := server.PollerResponderNext(beginServerGroupMembers, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		h.beginServerGroupMembers.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginServerGroupMembers) {
		h.beginServerGroupMembers.remove(req)
	}

	return resp, nil
}

func (h *HypervDependencyMapControllerServerTransport) dispatchBeginUpdateDependencyMapStatus(req *http.Request) (*http.Response, error) {
	if h.srv.BeginUpdateDependencyMapStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateDependencyMapStatus not implemented")}
	}
	beginUpdateDependencyMapStatus := h.beginUpdateDependencyMapStatus.get(req)
	if beginUpdateDependencyMapStatus == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateDependencyMapStatus`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrate.UpdateMachineDepMapStatus](req)
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
		respr, errRespr := h.srv.BeginUpdateDependencyMapStatus(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateDependencyMapStatus = &respr
		h.beginUpdateDependencyMapStatus.add(req, beginUpdateDependencyMapStatus)
	}

	resp, err := server.PollerResponderNext(beginUpdateDependencyMapStatus, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		h.beginUpdateDependencyMapStatus.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateDependencyMapStatus) {
		h.beginUpdateDependencyMapStatus.remove(req)
	}

	return resp, nil
}
