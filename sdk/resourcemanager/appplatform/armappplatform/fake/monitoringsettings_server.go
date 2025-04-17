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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
	"net/http"
	"net/url"
	"regexp"
)

// MonitoringSettingsServer is a fake server for instances of the armappplatform.MonitoringSettingsClient type.
type MonitoringSettingsServer struct {
	// Get is the fake for method MonitoringSettingsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, options *armappplatform.MonitoringSettingsClientGetOptions) (resp azfake.Responder[armappplatform.MonitoringSettingsClientGetResponse], errResp azfake.ErrorResponder)

	// BeginUpdatePatch is the fake for method MonitoringSettingsClient.BeginUpdatePatch
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdatePatch func(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource armappplatform.MonitoringSettingResource, options *armappplatform.MonitoringSettingsClientBeginUpdatePatchOptions) (resp azfake.PollerResponder[armappplatform.MonitoringSettingsClientUpdatePatchResponse], errResp azfake.ErrorResponder)

	// BeginUpdatePut is the fake for method MonitoringSettingsClient.BeginUpdatePut
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdatePut func(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource armappplatform.MonitoringSettingResource, options *armappplatform.MonitoringSettingsClientBeginUpdatePutOptions) (resp azfake.PollerResponder[armappplatform.MonitoringSettingsClientUpdatePutResponse], errResp azfake.ErrorResponder)
}

// NewMonitoringSettingsServerTransport creates a new instance of MonitoringSettingsServerTransport with the provided implementation.
// The returned MonitoringSettingsServerTransport instance is connected to an instance of armappplatform.MonitoringSettingsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMonitoringSettingsServerTransport(srv *MonitoringSettingsServer) *MonitoringSettingsServerTransport {
	return &MonitoringSettingsServerTransport{
		srv:              srv,
		beginUpdatePatch: newTracker[azfake.PollerResponder[armappplatform.MonitoringSettingsClientUpdatePatchResponse]](),
		beginUpdatePut:   newTracker[azfake.PollerResponder[armappplatform.MonitoringSettingsClientUpdatePutResponse]](),
	}
}

// MonitoringSettingsServerTransport connects instances of armappplatform.MonitoringSettingsClient to instances of MonitoringSettingsServer.
// Don't use this type directly, use NewMonitoringSettingsServerTransport instead.
type MonitoringSettingsServerTransport struct {
	srv              *MonitoringSettingsServer
	beginUpdatePatch *tracker[azfake.PollerResponder[armappplatform.MonitoringSettingsClientUpdatePatchResponse]]
	beginUpdatePut   *tracker[azfake.PollerResponder[armappplatform.MonitoringSettingsClientUpdatePutResponse]]
}

// Do implements the policy.Transporter interface for MonitoringSettingsServerTransport.
func (m *MonitoringSettingsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *MonitoringSettingsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if monitoringSettingsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = monitoringSettingsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "MonitoringSettingsClient.Get":
				res.resp, res.err = m.dispatchGet(req)
			case "MonitoringSettingsClient.BeginUpdatePatch":
				res.resp, res.err = m.dispatchBeginUpdatePatch(req)
			case "MonitoringSettingsClient.BeginUpdatePut":
				res.resp, res.err = m.dispatchBeginUpdatePut(req)
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

func (m *MonitoringSettingsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/monitoringSettings/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).MonitoringSettingResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MonitoringSettingsServerTransport) dispatchBeginUpdatePatch(req *http.Request) (*http.Response, error) {
	if m.srv.BeginUpdatePatch == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdatePatch not implemented")}
	}
	beginUpdatePatch := m.beginUpdatePatch.get(req)
	if beginUpdatePatch == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/monitoringSettings/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappplatform.MonitoringSettingResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginUpdatePatch(req.Context(), resourceGroupNameParam, serviceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdatePatch = &respr
		m.beginUpdatePatch.add(req, beginUpdatePatch)
	}

	resp, err := server.PollerResponderNext(beginUpdatePatch, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginUpdatePatch.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdatePatch) {
		m.beginUpdatePatch.remove(req)
	}

	return resp, nil
}

func (m *MonitoringSettingsServerTransport) dispatchBeginUpdatePut(req *http.Request) (*http.Response, error) {
	if m.srv.BeginUpdatePut == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdatePut not implemented")}
	}
	beginUpdatePut := m.beginUpdatePut.get(req)
	if beginUpdatePut == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/monitoringSettings/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappplatform.MonitoringSettingResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginUpdatePut(req.Context(), resourceGroupNameParam, serviceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdatePut = &respr
		m.beginUpdatePut.add(req, beginUpdatePut)
	}

	resp, err := server.PollerResponderNext(beginUpdatePut, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginUpdatePut.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdatePut) {
		m.beginUpdatePut.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to MonitoringSettingsServerTransport
var monitoringSettingsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
