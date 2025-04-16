// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
	"net/http"
	"net/url"
	"regexp"
)

// BackupPoliciesServer is a fake server for instances of the armrecoveryservicesbackup.BackupPoliciesClient type.
type BackupPoliciesServer struct {
	// NewListPager is the fake for method BackupPoliciesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(vaultName string, resourceGroupName string, options *armrecoveryservicesbackup.BackupPoliciesClientListOptions) (resp azfake.PagerResponder[armrecoveryservicesbackup.BackupPoliciesClientListResponse])
}

// NewBackupPoliciesServerTransport creates a new instance of BackupPoliciesServerTransport with the provided implementation.
// The returned BackupPoliciesServerTransport instance is connected to an instance of armrecoveryservicesbackup.BackupPoliciesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBackupPoliciesServerTransport(srv *BackupPoliciesServer) *BackupPoliciesServerTransport {
	return &BackupPoliciesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armrecoveryservicesbackup.BackupPoliciesClientListResponse]](),
	}
}

// BackupPoliciesServerTransport connects instances of armrecoveryservicesbackup.BackupPoliciesClient to instances of BackupPoliciesServer.
// Don't use this type directly, use NewBackupPoliciesServerTransport instead.
type BackupPoliciesServerTransport struct {
	srv          *BackupPoliciesServer
	newListPager *tracker[azfake.PagerResponder[armrecoveryservicesbackup.BackupPoliciesClientListResponse]]
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
			case "BackupPoliciesClient.NewListPager":
				res.resp, res.err = b.dispatchNewListPager(req)
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

func (b *BackupPoliciesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armrecoveryservicesbackup.BackupPoliciesClientListOptions
		if filterParam != nil {
			options = &armrecoveryservicesbackup.BackupPoliciesClientListOptions{
				Filter: filterParam,
			}
		}
		resp := b.srv.NewListPager(vaultNameParam, resourceGroupNameParam, options)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armrecoveryservicesbackup.BackupPoliciesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
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

// set this to conditionally intercept incoming requests to BackupPoliciesServerTransport
var backupPoliciesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
