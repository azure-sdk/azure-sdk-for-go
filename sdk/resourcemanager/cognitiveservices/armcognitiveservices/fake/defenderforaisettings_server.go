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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"net/http"
	"net/url"
	"regexp"
)

// DefenderForAISettingsServer is a fake server for instances of the armcognitiveservices.DefenderForAISettingsClient type.
type DefenderForAISettingsServer struct {
	// CreateOrUpdate is the fake for method DefenderForAISettingsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, accountName string, defenderForAISettingName string, defenderForAISettings armcognitiveservices.DefenderForAISetting, options *armcognitiveservices.DefenderForAISettingsClientCreateOrUpdateOptions) (resp azfake.Responder[armcognitiveservices.DefenderForAISettingsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DefenderForAISettingsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, defenderForAISettingName string, options *armcognitiveservices.DefenderForAISettingsClientGetOptions) (resp azfake.Responder[armcognitiveservices.DefenderForAISettingsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DefenderForAISettingsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, accountName string, options *armcognitiveservices.DefenderForAISettingsClientListOptions) (resp azfake.PagerResponder[armcognitiveservices.DefenderForAISettingsClientListResponse])

	// Update is the fake for method DefenderForAISettingsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, accountName string, defenderForAISettingName string, defenderForAISettings armcognitiveservices.DefenderForAISetting, options *armcognitiveservices.DefenderForAISettingsClientUpdateOptions) (resp azfake.Responder[armcognitiveservices.DefenderForAISettingsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewDefenderForAISettingsServerTransport creates a new instance of DefenderForAISettingsServerTransport with the provided implementation.
// The returned DefenderForAISettingsServerTransport instance is connected to an instance of armcognitiveservices.DefenderForAISettingsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDefenderForAISettingsServerTransport(srv *DefenderForAISettingsServer) *DefenderForAISettingsServerTransport {
	return &DefenderForAISettingsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armcognitiveservices.DefenderForAISettingsClientListResponse]](),
	}
}

// DefenderForAISettingsServerTransport connects instances of armcognitiveservices.DefenderForAISettingsClient to instances of DefenderForAISettingsServer.
// Don't use this type directly, use NewDefenderForAISettingsServerTransport instead.
type DefenderForAISettingsServerTransport struct {
	srv          *DefenderForAISettingsServer
	newListPager *tracker[azfake.PagerResponder[armcognitiveservices.DefenderForAISettingsClientListResponse]]
}

// Do implements the policy.Transporter interface for DefenderForAISettingsServerTransport.
func (d *DefenderForAISettingsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DefenderForAISettingsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if defenderForAiSettingsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = defenderForAiSettingsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DefenderForAISettingsClient.CreateOrUpdate":
				res.resp, res.err = d.dispatchCreateOrUpdate(req)
			case "DefenderForAISettingsClient.Get":
				res.resp, res.err = d.dispatchGet(req)
			case "DefenderForAISettingsClient.NewListPager":
				res.resp, res.err = d.dispatchNewListPager(req)
			case "DefenderForAISettingsClient.Update":
				res.resp, res.err = d.dispatchUpdate(req)
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

func (d *DefenderForAISettingsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CognitiveServices/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/defenderForAISettings/(?P<defenderForAISettingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcognitiveservices.DefenderForAISetting](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	defenderForAISettingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("defenderForAISettingName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, accountNameParam, defenderForAISettingNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DefenderForAISetting, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DefenderForAISettingsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CognitiveServices/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/defenderForAISettings/(?P<defenderForAISettingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	defenderForAISettingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("defenderForAISettingName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, defenderForAISettingNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DefenderForAISetting, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DefenderForAISettingsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CognitiveServices/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/defenderForAISettings`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListPager(resourceGroupNameParam, accountNameParam, nil)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcognitiveservices.DefenderForAISettingsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		d.newListPager.remove(req)
	}
	return resp, nil
}

func (d *DefenderForAISettingsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CognitiveServices/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/defenderForAISettings/(?P<defenderForAISettingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcognitiveservices.DefenderForAISetting](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	defenderForAISettingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("defenderForAISettingName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Update(req.Context(), resourceGroupNameParam, accountNameParam, defenderForAISettingNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DefenderForAISetting, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to DefenderForAISettingsServerTransport
var defenderForAiSettingsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
