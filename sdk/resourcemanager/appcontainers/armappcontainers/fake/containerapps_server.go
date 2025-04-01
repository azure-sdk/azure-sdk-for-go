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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ContainerAppsServer is a fake server for instances of the armappcontainers.ContainerAppsClient type.
type ContainerAppsServer struct {
	// BeginCreateOrUpdate is the fake for method ContainerAppsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, containerAppName string, containerAppEnvelope armappcontainers.ContainerApp, options *armappcontainers.ContainerAppsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armappcontainers.ContainerAppsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ContainerAppsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, containerAppName string, options *armappcontainers.ContainerAppsClientBeginDeleteOptions) (resp azfake.PollerResponder[armappcontainers.ContainerAppsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ContainerAppsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, containerAppName string, options *armappcontainers.ContainerAppsClientGetOptions) (resp azfake.Responder[armappcontainers.ContainerAppsClientGetResponse], errResp azfake.ErrorResponder)

	// GetAuthToken is the fake for method ContainerAppsClient.GetAuthToken
	// HTTP status codes to indicate success: http.StatusOK
	GetAuthToken func(ctx context.Context, resourceGroupName string, containerAppName string, options *armappcontainers.ContainerAppsClientGetAuthTokenOptions) (resp azfake.Responder[armappcontainers.ContainerAppsClientGetAuthTokenResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method ContainerAppsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armappcontainers.ContainerAppsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armappcontainers.ContainerAppsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method ContainerAppsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armappcontainers.ContainerAppsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armappcontainers.ContainerAppsClientListBySubscriptionResponse])

	// ListCustomHostNameAnalysis is the fake for method ContainerAppsClient.ListCustomHostNameAnalysis
	// HTTP status codes to indicate success: http.StatusOK
	ListCustomHostNameAnalysis func(ctx context.Context, resourceGroupName string, containerAppName string, options *armappcontainers.ContainerAppsClientListCustomHostNameAnalysisOptions) (resp azfake.Responder[armappcontainers.ContainerAppsClientListCustomHostNameAnalysisResponse], errResp azfake.ErrorResponder)

	// ListSecrets is the fake for method ContainerAppsClient.ListSecrets
	// HTTP status codes to indicate success: http.StatusOK
	ListSecrets func(ctx context.Context, resourceGroupName string, containerAppName string, options *armappcontainers.ContainerAppsClientListSecretsOptions) (resp azfake.Responder[armappcontainers.ContainerAppsClientListSecretsResponse], errResp azfake.ErrorResponder)

	// BeginStart is the fake for method ContainerAppsClient.BeginStart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStart func(ctx context.Context, resourceGroupName string, containerAppName string, options *armappcontainers.ContainerAppsClientBeginStartOptions) (resp azfake.PollerResponder[armappcontainers.ContainerAppsClientStartResponse], errResp azfake.ErrorResponder)

	// BeginStop is the fake for method ContainerAppsClient.BeginStop
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStop func(ctx context.Context, resourceGroupName string, containerAppName string, options *armappcontainers.ContainerAppsClientBeginStopOptions) (resp azfake.PollerResponder[armappcontainers.ContainerAppsClientStopResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method ContainerAppsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, containerAppName string, containerAppEnvelope armappcontainers.ContainerApp, options *armappcontainers.ContainerAppsClientBeginUpdateOptions) (resp azfake.PollerResponder[armappcontainers.ContainerAppsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewContainerAppsServerTransport creates a new instance of ContainerAppsServerTransport with the provided implementation.
// The returned ContainerAppsServerTransport instance is connected to an instance of armappcontainers.ContainerAppsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewContainerAppsServerTransport(srv *ContainerAppsServer) *ContainerAppsServerTransport {
	return &ContainerAppsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armappcontainers.ContainerAppsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armappcontainers.ContainerAppsClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armappcontainers.ContainerAppsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armappcontainers.ContainerAppsClientListBySubscriptionResponse]](),
		beginStart:                  newTracker[azfake.PollerResponder[armappcontainers.ContainerAppsClientStartResponse]](),
		beginStop:                   newTracker[azfake.PollerResponder[armappcontainers.ContainerAppsClientStopResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armappcontainers.ContainerAppsClientUpdateResponse]](),
	}
}

// ContainerAppsServerTransport connects instances of armappcontainers.ContainerAppsClient to instances of ContainerAppsServer.
// Don't use this type directly, use NewContainerAppsServerTransport instead.
type ContainerAppsServerTransport struct {
	srv                         *ContainerAppsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armappcontainers.ContainerAppsClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armappcontainers.ContainerAppsClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armappcontainers.ContainerAppsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armappcontainers.ContainerAppsClientListBySubscriptionResponse]]
	beginStart                  *tracker[azfake.PollerResponder[armappcontainers.ContainerAppsClientStartResponse]]
	beginStop                   *tracker[azfake.PollerResponder[armappcontainers.ContainerAppsClientStopResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armappcontainers.ContainerAppsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ContainerAppsServerTransport.
func (c *ContainerAppsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ContainerAppsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if containerAppsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = containerAppsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ContainerAppsClient.BeginCreateOrUpdate":
				res.resp, res.err = c.dispatchBeginCreateOrUpdate(req)
			case "ContainerAppsClient.BeginDelete":
				res.resp, res.err = c.dispatchBeginDelete(req)
			case "ContainerAppsClient.Get":
				res.resp, res.err = c.dispatchGet(req)
			case "ContainerAppsClient.GetAuthToken":
				res.resp, res.err = c.dispatchGetAuthToken(req)
			case "ContainerAppsClient.NewListByResourceGroupPager":
				res.resp, res.err = c.dispatchNewListByResourceGroupPager(req)
			case "ContainerAppsClient.NewListBySubscriptionPager":
				res.resp, res.err = c.dispatchNewListBySubscriptionPager(req)
			case "ContainerAppsClient.ListCustomHostNameAnalysis":
				res.resp, res.err = c.dispatchListCustomHostNameAnalysis(req)
			case "ContainerAppsClient.ListSecrets":
				res.resp, res.err = c.dispatchListSecrets(req)
			case "ContainerAppsClient.BeginStart":
				res.resp, res.err = c.dispatchBeginStart(req)
			case "ContainerAppsClient.BeginStop":
				res.resp, res.err = c.dispatchBeginStop(req)
			case "ContainerAppsClient.BeginUpdate":
				res.resp, res.err = c.dispatchBeginUpdate(req)
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

func (c *ContainerAppsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappcontainers.ContainerApp](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, containerAppNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		c.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		c.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (c *ContainerAppsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, containerAppNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *ContainerAppsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, containerAppNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ContainerApp, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerAppsServerTransport) dispatchGetAuthToken(req *http.Request) (*http.Response, error) {
	if c.srv.GetAuthToken == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAuthToken not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getAuthtoken`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetAuthToken(req.Context(), resourceGroupNameParam, containerAppNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ContainerAppAuthToken, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerAppsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := c.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		c.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armappcontainers.ContainerAppsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		c.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (c *ContainerAppsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := c.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		c.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armappcontainers.ContainerAppsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		c.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (c *ContainerAppsServerTransport) dispatchListCustomHostNameAnalysis(req *http.Request) (*http.Response, error) {
	if c.srv.ListCustomHostNameAnalysis == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListCustomHostNameAnalysis not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listCustomHostNameAnalysis`
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
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	customHostnameUnescaped, err := url.QueryUnescape(qp.Get("customHostname"))
	if err != nil {
		return nil, err
	}
	customHostnameParam := getOptional(customHostnameUnescaped)
	var options *armappcontainers.ContainerAppsClientListCustomHostNameAnalysisOptions
	if customHostnameParam != nil {
		options = &armappcontainers.ContainerAppsClientListCustomHostNameAnalysisOptions{
			CustomHostname: customHostnameParam,
		}
	}
	respr, errRespr := c.srv.ListCustomHostNameAnalysis(req.Context(), resourceGroupNameParam, containerAppNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CustomHostnameAnalysisResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerAppsServerTransport) dispatchListSecrets(req *http.Request) (*http.Response, error) {
	if c.srv.ListSecrets == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListSecrets not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listSecrets`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.ListSecrets(req.Context(), resourceGroupNameParam, containerAppNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SecretsCollection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerAppsServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if c.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStart not implemented")}
	}
	beginStart := c.beginStart.get(req)
	if beginStart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/start`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginStart(req.Context(), resourceGroupNameParam, containerAppNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStart = &respr
		c.beginStart.add(req, beginStart)
	}

	resp, err := server.PollerResponderNext(beginStart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginStart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStart) {
		c.beginStart.remove(req)
	}

	return resp, nil
}

func (c *ContainerAppsServerTransport) dispatchBeginStop(req *http.Request) (*http.Response, error) {
	if c.srv.BeginStop == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStop not implemented")}
	}
	beginStop := c.beginStop.get(req)
	if beginStop == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stop`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginStop(req.Context(), resourceGroupNameParam, containerAppNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStop = &respr
		c.beginStop.add(req, beginStop)
	}

	resp, err := server.PollerResponderNext(beginStop, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginStop.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStop) {
		c.beginStop.remove(req)
	}

	return resp, nil
}

func (c *ContainerAppsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := c.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappcontainers.ContainerApp](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginUpdate(req.Context(), resourceGroupNameParam, containerAppNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		c.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		c.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ContainerAppsServerTransport
var containerAppsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
