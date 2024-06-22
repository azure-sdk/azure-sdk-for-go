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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v2"
	"net/http"
	"net/url"
	"regexp"
)

// UpdateRunsServer is a fake server for instances of the armcontainerservicefleet.UpdateRunsClient type.
type UpdateRunsServer struct {
	// BeginCreateOrUpdate is the fake for method UpdateRunsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, resource armcontainerservicefleet.UpdateRun, options *armcontainerservicefleet.UpdateRunsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcontainerservicefleet.UpdateRunsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method UpdateRunsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, options *armcontainerservicefleet.UpdateRunsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcontainerservicefleet.UpdateRunsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method UpdateRunsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, options *armcontainerservicefleet.UpdateRunsClientGetOptions) (resp azfake.Responder[armcontainerservicefleet.UpdateRunsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByFleetPager is the fake for method UpdateRunsClient.NewListByFleetPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByFleetPager func(resourceGroupName string, fleetName string, options *armcontainerservicefleet.UpdateRunsClientListByFleetOptions) (resp azfake.PagerResponder[armcontainerservicefleet.UpdateRunsClientListByFleetResponse])

	// BeginSkip is the fake for method UpdateRunsClient.BeginSkip
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginSkip func(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, body armcontainerservicefleet.SkipProperties, options *armcontainerservicefleet.UpdateRunsClientBeginSkipOptions) (resp azfake.PollerResponder[armcontainerservicefleet.UpdateRunsClientSkipResponse], errResp azfake.ErrorResponder)

	// BeginStart is the fake for method UpdateRunsClient.BeginStart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStart func(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, options *armcontainerservicefleet.UpdateRunsClientBeginStartOptions) (resp azfake.PollerResponder[armcontainerservicefleet.UpdateRunsClientStartResponse], errResp azfake.ErrorResponder)

	// BeginStop is the fake for method UpdateRunsClient.BeginStop
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStop func(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, options *armcontainerservicefleet.UpdateRunsClientBeginStopOptions) (resp azfake.PollerResponder[armcontainerservicefleet.UpdateRunsClientStopResponse], errResp azfake.ErrorResponder)
}

// NewUpdateRunsServerTransport creates a new instance of UpdateRunsServerTransport with the provided implementation.
// The returned UpdateRunsServerTransport instance is connected to an instance of armcontainerservicefleet.UpdateRunsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewUpdateRunsServerTransport(srv *UpdateRunsServer) *UpdateRunsServerTransport {
	return &UpdateRunsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armcontainerservicefleet.UpdateRunsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armcontainerservicefleet.UpdateRunsClientDeleteResponse]](),
		newListByFleetPager: newTracker[azfake.PagerResponder[armcontainerservicefleet.UpdateRunsClientListByFleetResponse]](),
		beginSkip:           newTracker[azfake.PollerResponder[armcontainerservicefleet.UpdateRunsClientSkipResponse]](),
		beginStart:          newTracker[azfake.PollerResponder[armcontainerservicefleet.UpdateRunsClientStartResponse]](),
		beginStop:           newTracker[azfake.PollerResponder[armcontainerservicefleet.UpdateRunsClientStopResponse]](),
	}
}

// UpdateRunsServerTransport connects instances of armcontainerservicefleet.UpdateRunsClient to instances of UpdateRunsServer.
// Don't use this type directly, use NewUpdateRunsServerTransport instead.
type UpdateRunsServerTransport struct {
	srv                 *UpdateRunsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armcontainerservicefleet.UpdateRunsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armcontainerservicefleet.UpdateRunsClientDeleteResponse]]
	newListByFleetPager *tracker[azfake.PagerResponder[armcontainerservicefleet.UpdateRunsClientListByFleetResponse]]
	beginSkip           *tracker[azfake.PollerResponder[armcontainerservicefleet.UpdateRunsClientSkipResponse]]
	beginStart          *tracker[azfake.PollerResponder[armcontainerservicefleet.UpdateRunsClientStartResponse]]
	beginStop           *tracker[azfake.PollerResponder[armcontainerservicefleet.UpdateRunsClientStopResponse]]
}

// Do implements the policy.Transporter interface for UpdateRunsServerTransport.
func (u *UpdateRunsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "UpdateRunsClient.BeginCreateOrUpdate":
		resp, err = u.dispatchBeginCreateOrUpdate(req)
	case "UpdateRunsClient.BeginDelete":
		resp, err = u.dispatchBeginDelete(req)
	case "UpdateRunsClient.Get":
		resp, err = u.dispatchGet(req)
	case "UpdateRunsClient.NewListByFleetPager":
		resp, err = u.dispatchNewListByFleetPager(req)
	case "UpdateRunsClient.BeginSkip":
		resp, err = u.dispatchBeginSkip(req)
	case "UpdateRunsClient.BeginStart":
		resp, err = u.dispatchBeginStart(req)
	case "UpdateRunsClient.BeginStop":
		resp, err = u.dispatchBeginStop(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *UpdateRunsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if u.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := u.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateRuns/(?P<updateRunName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcontainerservicefleet.UpdateRun](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		updateRunNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("updateRunName")])
		if err != nil {
			return nil, err
		}
		var options *armcontainerservicefleet.UpdateRunsClientBeginCreateOrUpdateOptions
		if ifMatchParam != nil || ifNoneMatchParam != nil {
			options = &armcontainerservicefleet.UpdateRunsClientBeginCreateOrUpdateOptions{
				IfMatch:     ifMatchParam,
				IfNoneMatch: ifNoneMatchParam,
			}
		}
		respr, errRespr := u.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, fleetNameParam, updateRunNameParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		u.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		u.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		u.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (u *UpdateRunsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if u.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := u.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateRuns/(?P<updateRunName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		updateRunNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("updateRunName")])
		if err != nil {
			return nil, err
		}
		var options *armcontainerservicefleet.UpdateRunsClientBeginDeleteOptions
		if ifMatchParam != nil {
			options = &armcontainerservicefleet.UpdateRunsClientBeginDeleteOptions{
				IfMatch: ifMatchParam,
			}
		}
		respr, errRespr := u.srv.BeginDelete(req.Context(), resourceGroupNameParam, fleetNameParam, updateRunNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		u.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		u.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		u.beginDelete.remove(req)
	}

	return resp, nil
}

func (u *UpdateRunsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if u.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateRuns/(?P<updateRunName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
	if err != nil {
		return nil, err
	}
	updateRunNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("updateRunName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := u.srv.Get(req.Context(), resourceGroupNameParam, fleetNameParam, updateRunNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UpdateRun, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *UpdateRunsServerTransport) dispatchNewListByFleetPager(req *http.Request) (*http.Response, error) {
	if u.srv.NewListByFleetPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByFleetPager not implemented")}
	}
	newListByFleetPager := u.newListByFleetPager.get(req)
	if newListByFleetPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateRuns`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		resp := u.srv.NewListByFleetPager(resourceGroupNameParam, fleetNameParam, nil)
		newListByFleetPager = &resp
		u.newListByFleetPager.add(req, newListByFleetPager)
		server.PagerResponderInjectNextLinks(newListByFleetPager, req, func(page *armcontainerservicefleet.UpdateRunsClientListByFleetResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByFleetPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		u.newListByFleetPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByFleetPager) {
		u.newListByFleetPager.remove(req)
	}
	return resp, nil
}

func (u *UpdateRunsServerTransport) dispatchBeginSkip(req *http.Request) (*http.Response, error) {
	if u.srv.BeginSkip == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginSkip not implemented")}
	}
	beginSkip := u.beginSkip.get(req)
	if beginSkip == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateRuns/(?P<updateRunName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/skip`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcontainerservicefleet.SkipProperties](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		updateRunNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("updateRunName")])
		if err != nil {
			return nil, err
		}
		var options *armcontainerservicefleet.UpdateRunsClientBeginSkipOptions
		if ifMatchParam != nil {
			options = &armcontainerservicefleet.UpdateRunsClientBeginSkipOptions{
				IfMatch: ifMatchParam,
			}
		}
		respr, errRespr := u.srv.BeginSkip(req.Context(), resourceGroupNameParam, fleetNameParam, updateRunNameParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginSkip = &respr
		u.beginSkip.add(req, beginSkip)
	}

	resp, err := server.PollerResponderNext(beginSkip, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		u.beginSkip.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginSkip) {
		u.beginSkip.remove(req)
	}

	return resp, nil
}

func (u *UpdateRunsServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if u.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStart not implemented")}
	}
	beginStart := u.beginStart.get(req)
	if beginStart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateRuns/(?P<updateRunName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/start`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		updateRunNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("updateRunName")])
		if err != nil {
			return nil, err
		}
		var options *armcontainerservicefleet.UpdateRunsClientBeginStartOptions
		if ifMatchParam != nil {
			options = &armcontainerservicefleet.UpdateRunsClientBeginStartOptions{
				IfMatch: ifMatchParam,
			}
		}
		respr, errRespr := u.srv.BeginStart(req.Context(), resourceGroupNameParam, fleetNameParam, updateRunNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStart = &respr
		u.beginStart.add(req, beginStart)
	}

	resp, err := server.PollerResponderNext(beginStart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		u.beginStart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStart) {
		u.beginStart.remove(req)
	}

	return resp, nil
}

func (u *UpdateRunsServerTransport) dispatchBeginStop(req *http.Request) (*http.Response, error) {
	if u.srv.BeginStop == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStop not implemented")}
	}
	beginStop := u.beginStop.get(req)
	if beginStop == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateRuns/(?P<updateRunName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stop`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		updateRunNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("updateRunName")])
		if err != nil {
			return nil, err
		}
		var options *armcontainerservicefleet.UpdateRunsClientBeginStopOptions
		if ifMatchParam != nil {
			options = &armcontainerservicefleet.UpdateRunsClientBeginStopOptions{
				IfMatch: ifMatchParam,
			}
		}
		respr, errRespr := u.srv.BeginStop(req.Context(), resourceGroupNameParam, fleetNameParam, updateRunNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStop = &respr
		u.beginStop.add(req, beginStop)
	}

	resp, err := server.PollerResponderNext(beginStop, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		u.beginStop.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStop) {
		u.beginStop.remove(req)
	}

	return resp, nil
}
