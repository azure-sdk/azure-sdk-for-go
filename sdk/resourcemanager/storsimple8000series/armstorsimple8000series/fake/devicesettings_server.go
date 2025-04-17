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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
	"net/http"
	"net/url"
	"regexp"
)

// DeviceSettingsServer is a fake server for instances of the armstorsimple8000series.DeviceSettingsClient type.
type DeviceSettingsServer struct {
	// BeginCreateOrUpdateAlertSettings is the fake for method DeviceSettingsClient.BeginCreateOrUpdateAlertSettings
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdateAlertSettings func(ctx context.Context, deviceName string, resourceGroupName string, managerName string, parameters armstorsimple8000series.AlertSettings, options *armstorsimple8000series.DeviceSettingsClientBeginCreateOrUpdateAlertSettingsOptions) (resp azfake.PollerResponder[armstorsimple8000series.DeviceSettingsClientCreateOrUpdateAlertSettingsResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdateTimeSettings is the fake for method DeviceSettingsClient.BeginCreateOrUpdateTimeSettings
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdateTimeSettings func(ctx context.Context, deviceName string, resourceGroupName string, managerName string, parameters armstorsimple8000series.TimeSettings, options *armstorsimple8000series.DeviceSettingsClientBeginCreateOrUpdateTimeSettingsOptions) (resp azfake.PollerResponder[armstorsimple8000series.DeviceSettingsClientCreateOrUpdateTimeSettingsResponse], errResp azfake.ErrorResponder)

	// GetAlertSettings is the fake for method DeviceSettingsClient.GetAlertSettings
	// HTTP status codes to indicate success: http.StatusOK
	GetAlertSettings func(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *armstorsimple8000series.DeviceSettingsClientGetAlertSettingsOptions) (resp azfake.Responder[armstorsimple8000series.DeviceSettingsClientGetAlertSettingsResponse], errResp azfake.ErrorResponder)

	// GetNetworkSettings is the fake for method DeviceSettingsClient.GetNetworkSettings
	// HTTP status codes to indicate success: http.StatusOK
	GetNetworkSettings func(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *armstorsimple8000series.DeviceSettingsClientGetNetworkSettingsOptions) (resp azfake.Responder[armstorsimple8000series.DeviceSettingsClientGetNetworkSettingsResponse], errResp azfake.ErrorResponder)

	// GetSecuritySettings is the fake for method DeviceSettingsClient.GetSecuritySettings
	// HTTP status codes to indicate success: http.StatusOK
	GetSecuritySettings func(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *armstorsimple8000series.DeviceSettingsClientGetSecuritySettingsOptions) (resp azfake.Responder[armstorsimple8000series.DeviceSettingsClientGetSecuritySettingsResponse], errResp azfake.ErrorResponder)

	// GetTimeSettings is the fake for method DeviceSettingsClient.GetTimeSettings
	// HTTP status codes to indicate success: http.StatusOK
	GetTimeSettings func(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *armstorsimple8000series.DeviceSettingsClientGetTimeSettingsOptions) (resp azfake.Responder[armstorsimple8000series.DeviceSettingsClientGetTimeSettingsResponse], errResp azfake.ErrorResponder)

	// BeginSyncRemotemanagementCertificate is the fake for method DeviceSettingsClient.BeginSyncRemotemanagementCertificate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginSyncRemotemanagementCertificate func(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *armstorsimple8000series.DeviceSettingsClientBeginSyncRemotemanagementCertificateOptions) (resp azfake.PollerResponder[armstorsimple8000series.DeviceSettingsClientSyncRemotemanagementCertificateResponse], errResp azfake.ErrorResponder)

	// BeginUpdateNetworkSettings is the fake for method DeviceSettingsClient.BeginUpdateNetworkSettings
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateNetworkSettings func(ctx context.Context, deviceName string, resourceGroupName string, managerName string, parameters armstorsimple8000series.NetworkSettingsPatch, options *armstorsimple8000series.DeviceSettingsClientBeginUpdateNetworkSettingsOptions) (resp azfake.PollerResponder[armstorsimple8000series.DeviceSettingsClientUpdateNetworkSettingsResponse], errResp azfake.ErrorResponder)

	// BeginUpdateSecuritySettings is the fake for method DeviceSettingsClient.BeginUpdateSecuritySettings
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateSecuritySettings func(ctx context.Context, deviceName string, resourceGroupName string, managerName string, parameters armstorsimple8000series.SecuritySettingsPatch, options *armstorsimple8000series.DeviceSettingsClientBeginUpdateSecuritySettingsOptions) (resp azfake.PollerResponder[armstorsimple8000series.DeviceSettingsClientUpdateSecuritySettingsResponse], errResp azfake.ErrorResponder)
}

// NewDeviceSettingsServerTransport creates a new instance of DeviceSettingsServerTransport with the provided implementation.
// The returned DeviceSettingsServerTransport instance is connected to an instance of armstorsimple8000series.DeviceSettingsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDeviceSettingsServerTransport(srv *DeviceSettingsServer) *DeviceSettingsServerTransport {
	return &DeviceSettingsServerTransport{
		srv:                                  srv,
		beginCreateOrUpdateAlertSettings:     newTracker[azfake.PollerResponder[armstorsimple8000series.DeviceSettingsClientCreateOrUpdateAlertSettingsResponse]](),
		beginCreateOrUpdateTimeSettings:      newTracker[azfake.PollerResponder[armstorsimple8000series.DeviceSettingsClientCreateOrUpdateTimeSettingsResponse]](),
		beginSyncRemotemanagementCertificate: newTracker[azfake.PollerResponder[armstorsimple8000series.DeviceSettingsClientSyncRemotemanagementCertificateResponse]](),
		beginUpdateNetworkSettings:           newTracker[azfake.PollerResponder[armstorsimple8000series.DeviceSettingsClientUpdateNetworkSettingsResponse]](),
		beginUpdateSecuritySettings:          newTracker[azfake.PollerResponder[armstorsimple8000series.DeviceSettingsClientUpdateSecuritySettingsResponse]](),
	}
}

// DeviceSettingsServerTransport connects instances of armstorsimple8000series.DeviceSettingsClient to instances of DeviceSettingsServer.
// Don't use this type directly, use NewDeviceSettingsServerTransport instead.
type DeviceSettingsServerTransport struct {
	srv                                  *DeviceSettingsServer
	beginCreateOrUpdateAlertSettings     *tracker[azfake.PollerResponder[armstorsimple8000series.DeviceSettingsClientCreateOrUpdateAlertSettingsResponse]]
	beginCreateOrUpdateTimeSettings      *tracker[azfake.PollerResponder[armstorsimple8000series.DeviceSettingsClientCreateOrUpdateTimeSettingsResponse]]
	beginSyncRemotemanagementCertificate *tracker[azfake.PollerResponder[armstorsimple8000series.DeviceSettingsClientSyncRemotemanagementCertificateResponse]]
	beginUpdateNetworkSettings           *tracker[azfake.PollerResponder[armstorsimple8000series.DeviceSettingsClientUpdateNetworkSettingsResponse]]
	beginUpdateSecuritySettings          *tracker[azfake.PollerResponder[armstorsimple8000series.DeviceSettingsClientUpdateSecuritySettingsResponse]]
}

// Do implements the policy.Transporter interface for DeviceSettingsServerTransport.
func (d *DeviceSettingsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DeviceSettingsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if deviceSettingsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = deviceSettingsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DeviceSettingsClient.BeginCreateOrUpdateAlertSettings":
				res.resp, res.err = d.dispatchBeginCreateOrUpdateAlertSettings(req)
			case "DeviceSettingsClient.BeginCreateOrUpdateTimeSettings":
				res.resp, res.err = d.dispatchBeginCreateOrUpdateTimeSettings(req)
			case "DeviceSettingsClient.GetAlertSettings":
				res.resp, res.err = d.dispatchGetAlertSettings(req)
			case "DeviceSettingsClient.GetNetworkSettings":
				res.resp, res.err = d.dispatchGetNetworkSettings(req)
			case "DeviceSettingsClient.GetSecuritySettings":
				res.resp, res.err = d.dispatchGetSecuritySettings(req)
			case "DeviceSettingsClient.GetTimeSettings":
				res.resp, res.err = d.dispatchGetTimeSettings(req)
			case "DeviceSettingsClient.BeginSyncRemotemanagementCertificate":
				res.resp, res.err = d.dispatchBeginSyncRemotemanagementCertificate(req)
			case "DeviceSettingsClient.BeginUpdateNetworkSettings":
				res.resp, res.err = d.dispatchBeginUpdateNetworkSettings(req)
			case "DeviceSettingsClient.BeginUpdateSecuritySettings":
				res.resp, res.err = d.dispatchBeginUpdateSecuritySettings(req)
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

func (d *DeviceSettingsServerTransport) dispatchBeginCreateOrUpdateAlertSettings(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdateAlertSettings == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdateAlertSettings not implemented")}
	}
	beginCreateOrUpdateAlertSettings := d.beginCreateOrUpdateAlertSettings.get(req)
	if beginCreateOrUpdateAlertSettings == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/alertSettings/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstorsimple8000series.AlertSettings](req)
		if err != nil {
			return nil, err
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdateAlertSettings(req.Context(), deviceNameParam, resourceGroupNameParam, managerNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdateAlertSettings = &respr
		d.beginCreateOrUpdateAlertSettings.add(req, beginCreateOrUpdateAlertSettings)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdateAlertSettings, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginCreateOrUpdateAlertSettings.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdateAlertSettings) {
		d.beginCreateOrUpdateAlertSettings.remove(req)
	}

	return resp, nil
}

func (d *DeviceSettingsServerTransport) dispatchBeginCreateOrUpdateTimeSettings(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdateTimeSettings == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdateTimeSettings not implemented")}
	}
	beginCreateOrUpdateTimeSettings := d.beginCreateOrUpdateTimeSettings.get(req)
	if beginCreateOrUpdateTimeSettings == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/timeSettings/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstorsimple8000series.TimeSettings](req)
		if err != nil {
			return nil, err
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdateTimeSettings(req.Context(), deviceNameParam, resourceGroupNameParam, managerNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdateTimeSettings = &respr
		d.beginCreateOrUpdateTimeSettings.add(req, beginCreateOrUpdateTimeSettings)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdateTimeSettings, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginCreateOrUpdateTimeSettings.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdateTimeSettings) {
		d.beginCreateOrUpdateTimeSettings.remove(req)
	}

	return resp, nil
}

func (d *DeviceSettingsServerTransport) dispatchGetAlertSettings(req *http.Request) (*http.Response, error) {
	if d.srv.GetAlertSettings == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAlertSettings not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/alertSettings/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.GetAlertSettings(req.Context(), deviceNameParam, resourceGroupNameParam, managerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AlertSettings, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeviceSettingsServerTransport) dispatchGetNetworkSettings(req *http.Request) (*http.Response, error) {
	if d.srv.GetNetworkSettings == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNetworkSettings not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSettings/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.GetNetworkSettings(req.Context(), deviceNameParam, resourceGroupNameParam, managerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkSettings, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeviceSettingsServerTransport) dispatchGetSecuritySettings(req *http.Request) (*http.Response, error) {
	if d.srv.GetSecuritySettings == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSecuritySettings not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securitySettings/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.GetSecuritySettings(req.Context(), deviceNameParam, resourceGroupNameParam, managerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SecuritySettings, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeviceSettingsServerTransport) dispatchGetTimeSettings(req *http.Request) (*http.Response, error) {
	if d.srv.GetTimeSettings == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetTimeSettings not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/timeSettings/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.GetTimeSettings(req.Context(), deviceNameParam, resourceGroupNameParam, managerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TimeSettings, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeviceSettingsServerTransport) dispatchBeginSyncRemotemanagementCertificate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginSyncRemotemanagementCertificate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginSyncRemotemanagementCertificate not implemented")}
	}
	beginSyncRemotemanagementCertificate := d.beginSyncRemotemanagementCertificate.get(req)
	if beginSyncRemotemanagementCertificate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securitySettings/default/syncRemoteManagementCertificate`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginSyncRemotemanagementCertificate(req.Context(), deviceNameParam, resourceGroupNameParam, managerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginSyncRemotemanagementCertificate = &respr
		d.beginSyncRemotemanagementCertificate.add(req, beginSyncRemotemanagementCertificate)
	}

	resp, err := server.PollerResponderNext(beginSyncRemotemanagementCertificate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		d.beginSyncRemotemanagementCertificate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginSyncRemotemanagementCertificate) {
		d.beginSyncRemotemanagementCertificate.remove(req)
	}

	return resp, nil
}

func (d *DeviceSettingsServerTransport) dispatchBeginUpdateNetworkSettings(req *http.Request) (*http.Response, error) {
	if d.srv.BeginUpdateNetworkSettings == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateNetworkSettings not implemented")}
	}
	beginUpdateNetworkSettings := d.beginUpdateNetworkSettings.get(req)
	if beginUpdateNetworkSettings == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSettings/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstorsimple8000series.NetworkSettingsPatch](req)
		if err != nil {
			return nil, err
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginUpdateNetworkSettings(req.Context(), deviceNameParam, resourceGroupNameParam, managerNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateNetworkSettings = &respr
		d.beginUpdateNetworkSettings.add(req, beginUpdateNetworkSettings)
	}

	resp, err := server.PollerResponderNext(beginUpdateNetworkSettings, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginUpdateNetworkSettings.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateNetworkSettings) {
		d.beginUpdateNetworkSettings.remove(req)
	}

	return resp, nil
}

func (d *DeviceSettingsServerTransport) dispatchBeginUpdateSecuritySettings(req *http.Request) (*http.Response, error) {
	if d.srv.BeginUpdateSecuritySettings == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateSecuritySettings not implemented")}
	}
	beginUpdateSecuritySettings := d.beginUpdateSecuritySettings.get(req)
	if beginUpdateSecuritySettings == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securitySettings/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstorsimple8000series.SecuritySettingsPatch](req)
		if err != nil {
			return nil, err
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginUpdateSecuritySettings(req.Context(), deviceNameParam, resourceGroupNameParam, managerNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateSecuritySettings = &respr
		d.beginUpdateSecuritySettings.add(req, beginUpdateSecuritySettings)
	}

	resp, err := server.PollerResponderNext(beginUpdateSecuritySettings, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginUpdateSecuritySettings.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateSecuritySettings) {
		d.beginUpdateSecuritySettings.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to DeviceSettingsServerTransport
var deviceSettingsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
