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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
	"net/http"
	"net/url"
	"regexp"
)

// BackupPoliciesServer is a fake server for instances of the armnetapp.BackupPoliciesClient type.
type BackupPoliciesServer struct {
	// BeginCreate is the fake for method BackupPoliciesClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, accountName string, backupPolicyName string, body armnetapp.BackupPolicy, options *armnetapp.BackupPoliciesClientBeginCreateOptions) (resp azfake.PollerResponder[armnetapp.BackupPoliciesClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method BackupPoliciesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, accountName string, backupPolicyName string, options *armnetapp.BackupPoliciesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetapp.BackupPoliciesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method BackupPoliciesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, backupPolicyName string, options *armnetapp.BackupPoliciesClientGetOptions) (resp azfake.Responder[armnetapp.BackupPoliciesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method BackupPoliciesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, accountName string, options *armnetapp.BackupPoliciesClientListOptions) (resp azfake.PagerResponder[armnetapp.BackupPoliciesClientListResponse])

	// BeginUpdate is the fake for method BackupPoliciesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, accountName string, backupPolicyName string, body armnetapp.BackupPolicyPatch, options *armnetapp.BackupPoliciesClientBeginUpdateOptions) (resp azfake.PollerResponder[armnetapp.BackupPoliciesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewBackupPoliciesServerTransport creates a new instance of BackupPoliciesServerTransport with the provided implementation.
// The returned BackupPoliciesServerTransport instance is connected to an instance of armnetapp.BackupPoliciesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBackupPoliciesServerTransport(srv *BackupPoliciesServer) *BackupPoliciesServerTransport {
	return &BackupPoliciesServerTransport{
		srv:          srv,
		beginCreate:  newTracker[azfake.PollerResponder[armnetapp.BackupPoliciesClientCreateResponse]](),
		beginDelete:  newTracker[azfake.PollerResponder[armnetapp.BackupPoliciesClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armnetapp.BackupPoliciesClientListResponse]](),
		beginUpdate:  newTracker[azfake.PollerResponder[armnetapp.BackupPoliciesClientUpdateResponse]](),
	}
}

// BackupPoliciesServerTransport connects instances of armnetapp.BackupPoliciesClient to instances of BackupPoliciesServer.
// Don't use this type directly, use NewBackupPoliciesServerTransport instead.
type BackupPoliciesServerTransport struct {
	srv          *BackupPoliciesServer
	beginCreate  *tracker[azfake.PollerResponder[armnetapp.BackupPoliciesClientCreateResponse]]
	beginDelete  *tracker[azfake.PollerResponder[armnetapp.BackupPoliciesClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armnetapp.BackupPoliciesClientListResponse]]
	beginUpdate  *tracker[azfake.PollerResponder[armnetapp.BackupPoliciesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for BackupPoliciesServerTransport.
func (b *BackupPoliciesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BackupPoliciesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if backupPoliciesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = backupPoliciesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "BackupPoliciesClient.BeginCreate":
				res.resp, res.err = b.dispatchBeginCreate(req)
			case "BackupPoliciesClient.BeginDelete":
				res.resp, res.err = b.dispatchBeginDelete(req)
			case "BackupPoliciesClient.Get":
				res.resp, res.err = b.dispatchGet(req)
			case "BackupPoliciesClient.NewListPager":
				res.resp, res.err = b.dispatchNewListPager(req)
			case "BackupPoliciesClient.BeginUpdate":
				res.resp, res.err = b.dispatchBeginUpdate(req)
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

func (b *BackupPoliciesServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if b.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := b.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies/(?P<backupPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetapp.BackupPolicy](req)
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
		backupPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupPolicyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginCreate(req.Context(), resourceGroupNameParam, accountNameParam, backupPolicyNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		b.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		b.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		b.beginCreate.remove(req)
	}

	return resp, nil
}

func (b *BackupPoliciesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if b.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := b.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies/(?P<backupPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		backupPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupPolicyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginDelete(req.Context(), resourceGroupNameParam, accountNameParam, backupPolicyNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		b.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		b.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		b.beginDelete.remove(req)
	}

	return resp, nil
}

func (b *BackupPoliciesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies/(?P<backupPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	backupPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupPolicyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, backupPolicyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BackupPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BackupPoliciesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies`
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
		resp := b.srv.NewListPager(resourceGroupNameParam, accountNameParam, nil)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		b.newListPager.remove(req)
	}
	return resp, nil
}

func (b *BackupPoliciesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := b.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies/(?P<backupPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetapp.BackupPolicyPatch](req)
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
		backupPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupPolicyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginUpdate(req.Context(), resourceGroupNameParam, accountNameParam, backupPolicyNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		b.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		b.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		b.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to BackupPoliciesServerTransport
var backupPoliciesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
