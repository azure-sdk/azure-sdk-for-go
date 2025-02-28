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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
	"net/http"
	"net/url"
	"regexp"
)

// ConfigurationPolicyGroupsServer is a fake server for instances of the armnetwork.ConfigurationPolicyGroupsClient type.
type ConfigurationPolicyGroupsServer struct {
	// BeginCreateOrUpdate is the fake for method ConfigurationPolicyGroupsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string, vpnServerConfigurationPolicyGroupParameters armnetwork.VPNServerConfigurationPolicyGroup, options *armnetwork.ConfigurationPolicyGroupsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.ConfigurationPolicyGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ConfigurationPolicyGroupsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string, options *armnetwork.ConfigurationPolicyGroupsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.ConfigurationPolicyGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ConfigurationPolicyGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string, options *armnetwork.ConfigurationPolicyGroupsClientGetOptions) (resp azfake.Responder[armnetwork.ConfigurationPolicyGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByVPNServerConfigurationPager is the fake for method ConfigurationPolicyGroupsClient.NewListByVPNServerConfigurationPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByVPNServerConfigurationPager func(resourceGroupName string, vpnServerConfigurationName string, options *armnetwork.ConfigurationPolicyGroupsClientListByVPNServerConfigurationOptions) (resp azfake.PagerResponder[armnetwork.ConfigurationPolicyGroupsClientListByVPNServerConfigurationResponse])
}

// NewConfigurationPolicyGroupsServerTransport creates a new instance of ConfigurationPolicyGroupsServerTransport with the provided implementation.
// The returned ConfigurationPolicyGroupsServerTransport instance is connected to an instance of armnetwork.ConfigurationPolicyGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConfigurationPolicyGroupsServerTransport(srv *ConfigurationPolicyGroupsServer) *ConfigurationPolicyGroupsServerTransport {
	return &ConfigurationPolicyGroupsServerTransport{
		srv:                                  srv,
		beginCreateOrUpdate:                  newTracker[azfake.PollerResponder[armnetwork.ConfigurationPolicyGroupsClientCreateOrUpdateResponse]](),
		beginDelete:                          newTracker[azfake.PollerResponder[armnetwork.ConfigurationPolicyGroupsClientDeleteResponse]](),
		newListByVPNServerConfigurationPager: newTracker[azfake.PagerResponder[armnetwork.ConfigurationPolicyGroupsClientListByVPNServerConfigurationResponse]](),
	}
}

// ConfigurationPolicyGroupsServerTransport connects instances of armnetwork.ConfigurationPolicyGroupsClient to instances of ConfigurationPolicyGroupsServer.
// Don't use this type directly, use NewConfigurationPolicyGroupsServerTransport instead.
type ConfigurationPolicyGroupsServerTransport struct {
	srv                                  *ConfigurationPolicyGroupsServer
	beginCreateOrUpdate                  *tracker[azfake.PollerResponder[armnetwork.ConfigurationPolicyGroupsClientCreateOrUpdateResponse]]
	beginDelete                          *tracker[azfake.PollerResponder[armnetwork.ConfigurationPolicyGroupsClientDeleteResponse]]
	newListByVPNServerConfigurationPager *tracker[azfake.PagerResponder[armnetwork.ConfigurationPolicyGroupsClientListByVPNServerConfigurationResponse]]
}

// Do implements the policy.Transporter interface for ConfigurationPolicyGroupsServerTransport.
func (c *ConfigurationPolicyGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ConfigurationPolicyGroupsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if configurationPolicyGroupsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = configurationPolicyGroupsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ConfigurationPolicyGroupsClient.BeginCreateOrUpdate":
				res.resp, res.err = c.dispatchBeginCreateOrUpdate(req)
			case "ConfigurationPolicyGroupsClient.BeginDelete":
				res.resp, res.err = c.dispatchBeginDelete(req)
			case "ConfigurationPolicyGroupsClient.Get":
				res.resp, res.err = c.dispatchGet(req)
			case "ConfigurationPolicyGroupsClient.NewListByVPNServerConfigurationPager":
				res.resp, res.err = c.dispatchNewListByVPNServerConfigurationPager(req)
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

func (c *ConfigurationPolicyGroupsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnServerConfigurations/(?P<vpnServerConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/configurationPolicyGroups/(?P<configurationPolicyGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VPNServerConfigurationPolicyGroup](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vpnServerConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vpnServerConfigurationName")])
		if err != nil {
			return nil, err
		}
		configurationPolicyGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationPolicyGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, vpnServerConfigurationNameParam, configurationPolicyGroupNameParam, body, nil)
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

func (c *ConfigurationPolicyGroupsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnServerConfigurations/(?P<vpnServerConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/configurationPolicyGroups/(?P<configurationPolicyGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vpnServerConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vpnServerConfigurationName")])
		if err != nil {
			return nil, err
		}
		configurationPolicyGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationPolicyGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, vpnServerConfigurationNameParam, configurationPolicyGroupNameParam, nil)
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

func (c *ConfigurationPolicyGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnServerConfigurations/(?P<vpnServerConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/configurationPolicyGroups/(?P<configurationPolicyGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vpnServerConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vpnServerConfigurationName")])
	if err != nil {
		return nil, err
	}
	configurationPolicyGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationPolicyGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, vpnServerConfigurationNameParam, configurationPolicyGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VPNServerConfigurationPolicyGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConfigurationPolicyGroupsServerTransport) dispatchNewListByVPNServerConfigurationPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByVPNServerConfigurationPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByVPNServerConfigurationPager not implemented")}
	}
	newListByVPNServerConfigurationPager := c.newListByVPNServerConfigurationPager.get(req)
	if newListByVPNServerConfigurationPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnServerConfigurations/(?P<vpnServerConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/configurationPolicyGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vpnServerConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vpnServerConfigurationName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByVPNServerConfigurationPager(resourceGroupNameParam, vpnServerConfigurationNameParam, nil)
		newListByVPNServerConfigurationPager = &resp
		c.newListByVPNServerConfigurationPager.add(req, newListByVPNServerConfigurationPager)
		server.PagerResponderInjectNextLinks(newListByVPNServerConfigurationPager, req, func(page *armnetwork.ConfigurationPolicyGroupsClientListByVPNServerConfigurationResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByVPNServerConfigurationPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByVPNServerConfigurationPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByVPNServerConfigurationPager) {
		c.newListByVPNServerConfigurationPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ConfigurationPolicyGroupsServerTransport
var configurationPolicyGroupsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
