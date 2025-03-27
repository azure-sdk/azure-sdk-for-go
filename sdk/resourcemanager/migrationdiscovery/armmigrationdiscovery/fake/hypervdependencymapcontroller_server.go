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

// HypervDependencyMapControllerServer is a fake server for instances of the armmigrationdiscovery.HypervDependencyMapControllerClient type.
type HypervDependencyMapControllerServer struct {
	// BeginClientGroupMembers is the fake for method HypervDependencyMapControllerClient.BeginClientGroupMembers
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginClientGroupMembers func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrationdiscovery.DependencyMapServiceMapextensionsClientGroupMembersRequest, options *armmigrationdiscovery.HypervDependencyMapControllerClientBeginClientGroupMembersOptions) (resp azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientClientGroupMembersResponse], errResp azfake.ErrorResponder)

	// BeginExportDependencies is the fake for method HypervDependencyMapControllerClient.BeginExportDependencies
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExportDependencies func(ctx context.Context, resourceGroupName string, siteName string, requestBody armmigrationdiscovery.DependencyMapServiceMapextensionsExportDependenciesRequest, options *armmigrationdiscovery.HypervDependencyMapControllerClientBeginExportDependenciesOptions) (resp azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientExportDependenciesResponse], errResp azfake.ErrorResponder)

	// BeginGenerateCoarseMap is the fake for method HypervDependencyMapControllerClient.BeginGenerateCoarseMap
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGenerateCoarseMap func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrationdiscovery.DependencyMapServiceMapextensionsScopeMapRequest, options *armmigrationdiscovery.HypervDependencyMapControllerClientBeginGenerateCoarseMapOptions) (resp azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientGenerateCoarseMapResponse], errResp azfake.ErrorResponder)

	// BeginGenerateDetailedMap is the fake for method HypervDependencyMapControllerClient.BeginGenerateDetailedMap
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGenerateDetailedMap func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrationdiscovery.DependencyMapServiceMapextensionsSingleMachineDetailedMapRequest, options *armmigrationdiscovery.HypervDependencyMapControllerClientBeginGenerateDetailedMapOptions) (resp azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientGenerateDetailedMapResponse], errResp azfake.ErrorResponder)

	// BeginServerGroupMembers is the fake for method HypervDependencyMapControllerClient.BeginServerGroupMembers
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginServerGroupMembers func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrationdiscovery.DependencyMapServiceMapextensionsServerGroupMembersRequest, options *armmigrationdiscovery.HypervDependencyMapControllerClientBeginServerGroupMembersOptions) (resp azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientServerGroupMembersResponse], errResp azfake.ErrorResponder)

	// BeginUpdateDependencyMapStatus is the fake for method HypervDependencyMapControllerClient.BeginUpdateDependencyMapStatus
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateDependencyMapStatus func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrationdiscovery.UpdateMachineDepMapStatus, options *armmigrationdiscovery.HypervDependencyMapControllerClientBeginUpdateDependencyMapStatusOptions) (resp azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientUpdateDependencyMapStatusResponse], errResp azfake.ErrorResponder)
}

// NewHypervDependencyMapControllerServerTransport creates a new instance of HypervDependencyMapControllerServerTransport with the provided implementation.
// The returned HypervDependencyMapControllerServerTransport instance is connected to an instance of armmigrationdiscovery.HypervDependencyMapControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHypervDependencyMapControllerServerTransport(srv *HypervDependencyMapControllerServer) *HypervDependencyMapControllerServerTransport {
	return &HypervDependencyMapControllerServerTransport{
		srv:                            srv,
		beginClientGroupMembers:        newTracker[azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientClientGroupMembersResponse]](),
		beginExportDependencies:        newTracker[azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientExportDependenciesResponse]](),
		beginGenerateCoarseMap:         newTracker[azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientGenerateCoarseMapResponse]](),
		beginGenerateDetailedMap:       newTracker[azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientGenerateDetailedMapResponse]](),
		beginServerGroupMembers:        newTracker[azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientServerGroupMembersResponse]](),
		beginUpdateDependencyMapStatus: newTracker[azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientUpdateDependencyMapStatusResponse]](),
	}
}

// HypervDependencyMapControllerServerTransport connects instances of armmigrationdiscovery.HypervDependencyMapControllerClient to instances of HypervDependencyMapControllerServer.
// Don't use this type directly, use NewHypervDependencyMapControllerServerTransport instead.
type HypervDependencyMapControllerServerTransport struct {
	srv                            *HypervDependencyMapControllerServer
	beginClientGroupMembers        *tracker[azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientClientGroupMembersResponse]]
	beginExportDependencies        *tracker[azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientExportDependenciesResponse]]
	beginGenerateCoarseMap         *tracker[azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientGenerateCoarseMapResponse]]
	beginGenerateDetailedMap       *tracker[azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientGenerateDetailedMapResponse]]
	beginServerGroupMembers        *tracker[azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientServerGroupMembersResponse]]
	beginUpdateDependencyMapStatus *tracker[azfake.PollerResponder[armmigrationdiscovery.HypervDependencyMapControllerClientUpdateDependencyMapStatusResponse]]
}

// Do implements the policy.Transporter interface for HypervDependencyMapControllerServerTransport.
func (h *HypervDependencyMapControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return h.dispatchToMethodFake(req, method)
}

func (h *HypervDependencyMapControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if hypervDependencyMapControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = hypervDependencyMapControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "HypervDependencyMapControllerClient.BeginClientGroupMembers":
				res.resp, res.err = h.dispatchBeginClientGroupMembers(req)
			case "HypervDependencyMapControllerClient.BeginExportDependencies":
				res.resp, res.err = h.dispatchBeginExportDependencies(req)
			case "HypervDependencyMapControllerClient.BeginGenerateCoarseMap":
				res.resp, res.err = h.dispatchBeginGenerateCoarseMap(req)
			case "HypervDependencyMapControllerClient.BeginGenerateDetailedMap":
				res.resp, res.err = h.dispatchBeginGenerateDetailedMap(req)
			case "HypervDependencyMapControllerClient.BeginServerGroupMembers":
				res.resp, res.err = h.dispatchBeginServerGroupMembers(req)
			case "HypervDependencyMapControllerClient.BeginUpdateDependencyMapStatus":
				res.resp, res.err = h.dispatchBeginUpdateDependencyMapStatus(req)
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
		body, err := server.UnmarshalRequestAsJSON[armmigrationdiscovery.UpdateMachineDepMapStatus](req)
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

// set this to conditionally intercept incoming requests to HypervDependencyMapControllerServerTransport
var hypervDependencyMapControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
