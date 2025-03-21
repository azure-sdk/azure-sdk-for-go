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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
	"net/http"
	"net/url"
	"regexp"
)

// ProtectionPoliciesServer is a fake server for instances of the armrecoveryservicesbackup.ProtectionPoliciesClient type.
type ProtectionPoliciesServer struct {
	// CreateOrUpdate is the fake for method ProtectionPoliciesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	CreateOrUpdate func(ctx context.Context, vaultName string, resourceGroupName string, policyName string, parameters armrecoveryservicesbackup.ProtectionPolicyResource, options *armrecoveryservicesbackup.ProtectionPoliciesClientCreateOrUpdateOptions) (resp azfake.Responder[armrecoveryservicesbackup.ProtectionPoliciesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ProtectionPoliciesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	BeginDelete func(ctx context.Context, vaultName string, resourceGroupName string, policyName string, options *armrecoveryservicesbackup.ProtectionPoliciesClientBeginDeleteOptions) (resp azfake.PollerResponder[armrecoveryservicesbackup.ProtectionPoliciesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ProtectionPoliciesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, vaultName string, resourceGroupName string, policyName string, options *armrecoveryservicesbackup.ProtectionPoliciesClientGetOptions) (resp azfake.Responder[armrecoveryservicesbackup.ProtectionPoliciesClientGetResponse], errResp azfake.ErrorResponder)
}

// NewProtectionPoliciesServerTransport creates a new instance of ProtectionPoliciesServerTransport with the provided implementation.
// The returned ProtectionPoliciesServerTransport instance is connected to an instance of armrecoveryservicesbackup.ProtectionPoliciesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProtectionPoliciesServerTransport(srv *ProtectionPoliciesServer) *ProtectionPoliciesServerTransport {
	return &ProtectionPoliciesServerTransport{
		srv:         srv,
		beginDelete: newTracker[azfake.PollerResponder[armrecoveryservicesbackup.ProtectionPoliciesClientDeleteResponse]](),
	}
}

// ProtectionPoliciesServerTransport connects instances of armrecoveryservicesbackup.ProtectionPoliciesClient to instances of ProtectionPoliciesServer.
// Don't use this type directly, use NewProtectionPoliciesServerTransport instead.
type ProtectionPoliciesServerTransport struct {
	srv         *ProtectionPoliciesServer
	beginDelete *tracker[azfake.PollerResponder[armrecoveryservicesbackup.ProtectionPoliciesClientDeleteResponse]]
}

// Do implements the policy.Transporter interface for ProtectionPoliciesServerTransport.
func (p *ProtectionPoliciesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *ProtectionPoliciesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if protectionPoliciesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = protectionPoliciesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ProtectionPoliciesClient.CreateOrUpdate":
				res.resp, res.err = p.dispatchCreateOrUpdate(req)
			case "ProtectionPoliciesClient.BeginDelete":
				res.resp, res.err = p.dispatchBeginDelete(req)
			case "ProtectionPoliciesClient.Get":
				res.resp, res.err = p.dispatchGet(req)
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

func (p *ProtectionPoliciesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesbackup.ProtectionPolicyResource](req)
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	policyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyName")])
	if err != nil {
		return nil, err
	}
	xMSAuthorizationAuxiliaryParam := getOptional(getHeaderValue(req.Header, "x-ms-authorization-auxiliary"))
	var options *armrecoveryservicesbackup.ProtectionPoliciesClientCreateOrUpdateOptions
	if xMSAuthorizationAuxiliaryParam != nil {
		options = &armrecoveryservicesbackup.ProtectionPoliciesClientCreateOrUpdateOptions{
			XMSAuthorizationAuxiliary: xMSAuthorizationAuxiliaryParam,
		}
	}
	respr, errRespr := p.srv.CreateOrUpdate(req.Context(), vaultNameParam, resourceGroupNameParam, policyNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProtectionPolicyResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProtectionPoliciesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		policyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), vaultNameParam, resourceGroupNameParam, policyNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		p.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *ProtectionPoliciesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	policyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), vaultNameParam, resourceGroupNameParam, policyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProtectionPolicyResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ProtectionPoliciesServerTransport
var protectionPoliciesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
