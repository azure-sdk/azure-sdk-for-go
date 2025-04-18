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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasefleetmanager/armdatabasefleetmanager"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// FleetspacesServer is a fake server for instances of the armdatabasefleetmanager.FleetspacesClient type.
type FleetspacesServer struct {
	// BeginCreateOrUpdate is the fake for method FleetspacesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, fleetName string, fleetspaceName string, resource armdatabasefleetmanager.Fleetspace, options *armdatabasefleetmanager.FleetspacesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdatabasefleetmanager.FleetspacesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method FleetspacesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, fleetName string, fleetspaceName string, options *armdatabasefleetmanager.FleetspacesClientBeginDeleteOptions) (resp azfake.PollerResponder[armdatabasefleetmanager.FleetspacesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FleetspacesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, fleetName string, fleetspaceName string, options *armdatabasefleetmanager.FleetspacesClientGetOptions) (resp azfake.Responder[armdatabasefleetmanager.FleetspacesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByFleetPager is the fake for method FleetspacesClient.NewListByFleetPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByFleetPager func(resourceGroupName string, fleetName string, options *armdatabasefleetmanager.FleetspacesClientListByFleetOptions) (resp azfake.PagerResponder[armdatabasefleetmanager.FleetspacesClientListByFleetResponse])

	// BeginRegisterServer is the fake for method FleetspacesClient.BeginRegisterServer
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginRegisterServer func(ctx context.Context, resourceGroupName string, fleetName string, fleetspaceName string, body armdatabasefleetmanager.RegisterServerProperties, options *armdatabasefleetmanager.FleetspacesClientBeginRegisterServerOptions) (resp azfake.PollerResponder[armdatabasefleetmanager.FleetspacesClientRegisterServerResponse], errResp azfake.ErrorResponder)

	// BeginUnregister is the fake for method FleetspacesClient.BeginUnregister
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginUnregister func(ctx context.Context, resourceGroupName string, fleetName string, fleetspaceName string, options *armdatabasefleetmanager.FleetspacesClientBeginUnregisterOptions) (resp azfake.PollerResponder[armdatabasefleetmanager.FleetspacesClientUnregisterResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method FleetspacesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, fleetName string, fleetspaceName string, properties armdatabasefleetmanager.Fleetspace, options *armdatabasefleetmanager.FleetspacesClientBeginUpdateOptions) (resp azfake.PollerResponder[armdatabasefleetmanager.FleetspacesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewFleetspacesServerTransport creates a new instance of FleetspacesServerTransport with the provided implementation.
// The returned FleetspacesServerTransport instance is connected to an instance of armdatabasefleetmanager.FleetspacesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFleetspacesServerTransport(srv *FleetspacesServer) *FleetspacesServerTransport {
	return &FleetspacesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armdatabasefleetmanager.FleetspacesClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armdatabasefleetmanager.FleetspacesClientDeleteResponse]](),
		newListByFleetPager: newTracker[azfake.PagerResponder[armdatabasefleetmanager.FleetspacesClientListByFleetResponse]](),
		beginRegisterServer: newTracker[azfake.PollerResponder[armdatabasefleetmanager.FleetspacesClientRegisterServerResponse]](),
		beginUnregister:     newTracker[azfake.PollerResponder[armdatabasefleetmanager.FleetspacesClientUnregisterResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armdatabasefleetmanager.FleetspacesClientUpdateResponse]](),
	}
}

// FleetspacesServerTransport connects instances of armdatabasefleetmanager.FleetspacesClient to instances of FleetspacesServer.
// Don't use this type directly, use NewFleetspacesServerTransport instead.
type FleetspacesServerTransport struct {
	srv                 *FleetspacesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armdatabasefleetmanager.FleetspacesClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armdatabasefleetmanager.FleetspacesClientDeleteResponse]]
	newListByFleetPager *tracker[azfake.PagerResponder[armdatabasefleetmanager.FleetspacesClientListByFleetResponse]]
	beginRegisterServer *tracker[azfake.PollerResponder[armdatabasefleetmanager.FleetspacesClientRegisterServerResponse]]
	beginUnregister     *tracker[azfake.PollerResponder[armdatabasefleetmanager.FleetspacesClientUnregisterResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armdatabasefleetmanager.FleetspacesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for FleetspacesServerTransport.
func (f *FleetspacesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return f.dispatchToMethodFake(req, method)
}

func (f *FleetspacesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if fleetspacesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = fleetspacesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "FleetspacesClient.BeginCreateOrUpdate":
				res.resp, res.err = f.dispatchBeginCreateOrUpdate(req)
			case "FleetspacesClient.BeginDelete":
				res.resp, res.err = f.dispatchBeginDelete(req)
			case "FleetspacesClient.Get":
				res.resp, res.err = f.dispatchGet(req)
			case "FleetspacesClient.NewListByFleetPager":
				res.resp, res.err = f.dispatchNewListByFleetPager(req)
			case "FleetspacesClient.BeginRegisterServer":
				res.resp, res.err = f.dispatchBeginRegisterServer(req)
			case "FleetspacesClient.BeginUnregister":
				res.resp, res.err = f.dispatchBeginUnregister(req)
			case "FleetspacesClient.BeginUpdate":
				res.resp, res.err = f.dispatchBeginUpdate(req)
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

func (f *FleetspacesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := f.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseFleetManager/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fleetspaces/(?P<fleetspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatabasefleetmanager.Fleetspace](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		fleetspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetspaceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, fleetNameParam, fleetspaceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		f.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		f.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		f.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (f *FleetspacesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if f.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := f.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseFleetManager/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fleetspaces/(?P<fleetspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		fleetspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetspaceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginDelete(req.Context(), resourceGroupNameParam, fleetNameParam, fleetspaceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		f.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		f.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		f.beginDelete.remove(req)
	}

	return resp, nil
}

func (f *FleetspacesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseFleetManager/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fleetspaces/(?P<fleetspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	fleetspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetspaceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), resourceGroupNameParam, fleetNameParam, fleetspaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Fleetspace, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FleetspacesServerTransport) dispatchNewListByFleetPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListByFleetPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByFleetPager not implemented")}
	}
	newListByFleetPager := f.newListByFleetPager.get(req)
	if newListByFleetPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseFleetManager/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fleetspaces`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		skiptokenUnescaped, err := url.QueryUnescape(qp.Get("$skiptoken"))
		if err != nil {
			return nil, err
		}
		skiptokenParam := getOptional(skiptokenUnescaped)
		var options *armdatabasefleetmanager.FleetspacesClientListByFleetOptions
		if skipParam != nil || topParam != nil || skiptokenParam != nil {
			options = &armdatabasefleetmanager.FleetspacesClientListByFleetOptions{
				Skip:      skipParam,
				Top:       topParam,
				Skiptoken: skiptokenParam,
			}
		}
		resp := f.srv.NewListByFleetPager(resourceGroupNameParam, fleetNameParam, options)
		newListByFleetPager = &resp
		f.newListByFleetPager.add(req, newListByFleetPager)
		server.PagerResponderInjectNextLinks(newListByFleetPager, req, func(page *armdatabasefleetmanager.FleetspacesClientListByFleetResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByFleetPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListByFleetPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByFleetPager) {
		f.newListByFleetPager.remove(req)
	}
	return resp, nil
}

func (f *FleetspacesServerTransport) dispatchBeginRegisterServer(req *http.Request) (*http.Response, error) {
	if f.srv.BeginRegisterServer == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRegisterServer not implemented")}
	}
	beginRegisterServer := f.beginRegisterServer.get(req)
	if beginRegisterServer == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseFleetManager/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fleetspaces/(?P<fleetspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/registerServer`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatabasefleetmanager.RegisterServerProperties](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		fleetspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetspaceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginRegisterServer(req.Context(), resourceGroupNameParam, fleetNameParam, fleetspaceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRegisterServer = &respr
		f.beginRegisterServer.add(req, beginRegisterServer)
	}

	resp, err := server.PollerResponderNext(beginRegisterServer, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		f.beginRegisterServer.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRegisterServer) {
		f.beginRegisterServer.remove(req)
	}

	return resp, nil
}

func (f *FleetspacesServerTransport) dispatchBeginUnregister(req *http.Request) (*http.Response, error) {
	if f.srv.BeginUnregister == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUnregister not implemented")}
	}
	beginUnregister := f.beginUnregister.get(req)
	if beginUnregister == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseFleetManager/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fleetspaces/(?P<fleetspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/unregister`
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
		fleetspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetspaceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginUnregister(req.Context(), resourceGroupNameParam, fleetNameParam, fleetspaceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUnregister = &respr
		f.beginUnregister.add(req, beginUnregister)
	}

	resp, err := server.PollerResponderNext(beginUnregister, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		f.beginUnregister.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUnregister) {
		f.beginUnregister.remove(req)
	}

	return resp, nil
}

func (f *FleetspacesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := f.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseFleetManager/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fleetspaces/(?P<fleetspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatabasefleetmanager.Fleetspace](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		fleetspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetspaceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginUpdate(req.Context(), resourceGroupNameParam, fleetNameParam, fleetspaceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		f.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		f.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		f.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to FleetspacesServerTransport
var fleetspacesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
