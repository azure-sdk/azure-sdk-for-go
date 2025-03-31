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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
	"net/http"
	"net/url"
	"regexp"
)

// InboundSecurityRuleServer is a fake server for instances of the armnetwork.InboundSecurityRuleClient type.
type InboundSecurityRuleServer struct {
	// BeginCreateOrUpdate is the fake for method InboundSecurityRuleClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, ruleCollectionName string, parameters armnetwork.InboundSecurityRule, options *armnetwork.InboundSecurityRuleClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.InboundSecurityRuleClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method InboundSecurityRuleClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, ruleCollectionName string, options *armnetwork.InboundSecurityRuleClientGetOptions) (resp azfake.Responder[armnetwork.InboundSecurityRuleClientGetResponse], errResp azfake.ErrorResponder)
}

// NewInboundSecurityRuleServerTransport creates a new instance of InboundSecurityRuleServerTransport with the provided implementation.
// The returned InboundSecurityRuleServerTransport instance is connected to an instance of armnetwork.InboundSecurityRuleClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewInboundSecurityRuleServerTransport(srv *InboundSecurityRuleServer) *InboundSecurityRuleServerTransport {
	return &InboundSecurityRuleServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.InboundSecurityRuleClientCreateOrUpdateResponse]](),
	}
}

// InboundSecurityRuleServerTransport connects instances of armnetwork.InboundSecurityRuleClient to instances of InboundSecurityRuleServer.
// Don't use this type directly, use NewInboundSecurityRuleServerTransport instead.
type InboundSecurityRuleServerTransport struct {
	srv                 *InboundSecurityRuleServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.InboundSecurityRuleClientCreateOrUpdateResponse]]
}

// Do implements the policy.Transporter interface for InboundSecurityRuleServerTransport.
func (i *InboundSecurityRuleServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return i.dispatchToMethodFake(req, method)
}

func (i *InboundSecurityRuleServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if inboundSecurityRuleServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = inboundSecurityRuleServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "InboundSecurityRuleClient.BeginCreateOrUpdate":
				res.resp, res.err = i.dispatchBeginCreateOrUpdate(req)
			case "InboundSecurityRuleClient.Get":
				res.resp, res.err = i.dispatchGet(req)
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

func (i *InboundSecurityRuleServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := i.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inboundSecurityRules/(?P<ruleCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.InboundSecurityRule](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
		if err != nil {
			return nil, err
		}
		ruleCollectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleCollectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, networkVirtualApplianceNameParam, ruleCollectionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		i.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		i.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		i.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (i *InboundSecurityRuleServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inboundSecurityRules/(?P<ruleCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
	if err != nil {
		return nil, err
	}
	ruleCollectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleCollectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, networkVirtualApplianceNameParam, ruleCollectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InboundSecurityRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to InboundSecurityRuleServerTransport
var inboundSecurityRuleServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
