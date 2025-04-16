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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ExtensionsServer is a fake server for instances of the armazurestackhci.ExtensionsClient type.
type ExtensionsServer struct {
	// BeginCreate is the fake for method ExtensionsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extension armazurestackhci.Extension, options *armazurestackhci.ExtensionsClientBeginCreateOptions) (resp azfake.PollerResponder[armazurestackhci.ExtensionsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ExtensionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, options *armazurestackhci.ExtensionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armazurestackhci.ExtensionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ExtensionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, options *armazurestackhci.ExtensionsClientGetOptions) (resp azfake.Responder[armazurestackhci.ExtensionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByArcSettingPager is the fake for method ExtensionsClient.NewListByArcSettingPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByArcSettingPager func(resourceGroupName string, clusterName string, arcSettingName string, options *armazurestackhci.ExtensionsClientListByArcSettingOptions) (resp azfake.PagerResponder[armazurestackhci.ExtensionsClientListByArcSettingResponse])

	// BeginUpdate is the fake for method ExtensionsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extension armazurestackhci.ExtensionPatch, options *armazurestackhci.ExtensionsClientBeginUpdateOptions) (resp azfake.PollerResponder[armazurestackhci.ExtensionsClientUpdateResponse], errResp azfake.ErrorResponder)

	// BeginUpgrade is the fake for method ExtensionsClient.BeginUpgrade
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginUpgrade func(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extensionUpgradeParameters armazurestackhci.ExtensionUpgradeParameters, options *armazurestackhci.ExtensionsClientBeginUpgradeOptions) (resp azfake.PollerResponder[armazurestackhci.ExtensionsClientUpgradeResponse], errResp azfake.ErrorResponder)
}

// NewExtensionsServerTransport creates a new instance of ExtensionsServerTransport with the provided implementation.
// The returned ExtensionsServerTransport instance is connected to an instance of armazurestackhci.ExtensionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExtensionsServerTransport(srv *ExtensionsServer) *ExtensionsServerTransport {
	return &ExtensionsServerTransport{
		srv:                      srv,
		beginCreate:              newTracker[azfake.PollerResponder[armazurestackhci.ExtensionsClientCreateResponse]](),
		beginDelete:              newTracker[azfake.PollerResponder[armazurestackhci.ExtensionsClientDeleteResponse]](),
		newListByArcSettingPager: newTracker[azfake.PagerResponder[armazurestackhci.ExtensionsClientListByArcSettingResponse]](),
		beginUpdate:              newTracker[azfake.PollerResponder[armazurestackhci.ExtensionsClientUpdateResponse]](),
		beginUpgrade:             newTracker[azfake.PollerResponder[armazurestackhci.ExtensionsClientUpgradeResponse]](),
	}
}

// ExtensionsServerTransport connects instances of armazurestackhci.ExtensionsClient to instances of ExtensionsServer.
// Don't use this type directly, use NewExtensionsServerTransport instead.
type ExtensionsServerTransport struct {
	srv                      *ExtensionsServer
	beginCreate              *tracker[azfake.PollerResponder[armazurestackhci.ExtensionsClientCreateResponse]]
	beginDelete              *tracker[azfake.PollerResponder[armazurestackhci.ExtensionsClientDeleteResponse]]
	newListByArcSettingPager *tracker[azfake.PagerResponder[armazurestackhci.ExtensionsClientListByArcSettingResponse]]
	beginUpdate              *tracker[azfake.PollerResponder[armazurestackhci.ExtensionsClientUpdateResponse]]
	beginUpgrade             *tracker[azfake.PollerResponder[armazurestackhci.ExtensionsClientUpgradeResponse]]
}

// Do implements the policy.Transporter interface for ExtensionsServerTransport.
func (e *ExtensionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *ExtensionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if extensionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = extensionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ExtensionsClient.BeginCreate":
				res.resp, res.err = e.dispatchBeginCreate(req)
			case "ExtensionsClient.BeginDelete":
				res.resp, res.err = e.dispatchBeginDelete(req)
			case "ExtensionsClient.Get":
				res.resp, res.err = e.dispatchGet(req)
			case "ExtensionsClient.NewListByArcSettingPager":
				res.resp, res.err = e.dispatchNewListByArcSettingPager(req)
			case "ExtensionsClient.BeginUpdate":
				res.resp, res.err = e.dispatchBeginUpdate(req)
			case "ExtensionsClient.BeginUpgrade":
				res.resp, res.err = e.dispatchBeginUpgrade(req)
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

func (e *ExtensionsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := e.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHCI/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/arcSettings/(?P<arcSettingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<extensionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armazurestackhci.Extension](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		arcSettingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("arcSettingName")])
		if err != nil {
			return nil, err
		}
		extensionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("extensionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreate(req.Context(), resourceGroupNameParam, clusterNameParam, arcSettingNameParam, extensionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		e.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		e.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		e.beginCreate.remove(req)
	}

	return resp, nil
}

func (e *ExtensionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if e.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := e.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHCI/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/arcSettings/(?P<arcSettingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<extensionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		arcSettingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("arcSettingName")])
		if err != nil {
			return nil, err
		}
		extensionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("extensionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginDelete(req.Context(), resourceGroupNameParam, clusterNameParam, arcSettingNameParam, extensionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		e.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		e.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		e.beginDelete.remove(req)
	}

	return resp, nil
}

func (e *ExtensionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHCI/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/arcSettings/(?P<arcSettingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<extensionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	arcSettingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("arcSettingName")])
	if err != nil {
		return nil, err
	}
	extensionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("extensionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, clusterNameParam, arcSettingNameParam, extensionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Extension, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExtensionsServerTransport) dispatchNewListByArcSettingPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByArcSettingPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByArcSettingPager not implemented")}
	}
	newListByArcSettingPager := e.newListByArcSettingPager.get(req)
	if newListByArcSettingPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHCI/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/arcSettings/(?P<arcSettingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		arcSettingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("arcSettingName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListByArcSettingPager(resourceGroupNameParam, clusterNameParam, arcSettingNameParam, nil)
		newListByArcSettingPager = &resp
		e.newListByArcSettingPager.add(req, newListByArcSettingPager)
		server.PagerResponderInjectNextLinks(newListByArcSettingPager, req, func(page *armazurestackhci.ExtensionsClientListByArcSettingResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByArcSettingPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByArcSettingPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByArcSettingPager) {
		e.newListByArcSettingPager.remove(req)
	}
	return resp, nil
}

func (e *ExtensionsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := e.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHCI/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/arcSettings/(?P<arcSettingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<extensionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armazurestackhci.ExtensionPatch](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		arcSettingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("arcSettingName")])
		if err != nil {
			return nil, err
		}
		extensionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("extensionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginUpdate(req.Context(), resourceGroupNameParam, clusterNameParam, arcSettingNameParam, extensionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		e.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		e.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		e.beginUpdate.remove(req)
	}

	return resp, nil
}

func (e *ExtensionsServerTransport) dispatchBeginUpgrade(req *http.Request) (*http.Response, error) {
	if e.srv.BeginUpgrade == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpgrade not implemented")}
	}
	beginUpgrade := e.beginUpgrade.get(req)
	if beginUpgrade == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHCI/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/arcSettings/(?P<arcSettingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<extensionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/upgrade`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armazurestackhci.ExtensionUpgradeParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		arcSettingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("arcSettingName")])
		if err != nil {
			return nil, err
		}
		extensionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("extensionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginUpgrade(req.Context(), resourceGroupNameParam, clusterNameParam, arcSettingNameParam, extensionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpgrade = &respr
		e.beginUpgrade.add(req, beginUpgrade)
	}

	resp, err := server.PollerResponderNext(beginUpgrade, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		e.beginUpgrade.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpgrade) {
		e.beginUpgrade.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ExtensionsServerTransport
var extensionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
