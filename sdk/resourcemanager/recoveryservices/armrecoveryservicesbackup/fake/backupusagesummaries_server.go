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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
	"net/http"
	"net/url"
	"regexp"
)

// BackupUsageSummariesServer is a fake server for instances of the armrecoveryservicesbackup.BackupUsageSummariesClient type.
type BackupUsageSummariesServer struct {
	// NewListPager is the fake for method BackupUsageSummariesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(vaultName string, resourceGroupName string, options *armrecoveryservicesbackup.BackupUsageSummariesClientListOptions) (resp azfake.PagerResponder[armrecoveryservicesbackup.BackupUsageSummariesClientListResponse])
}

// NewBackupUsageSummariesServerTransport creates a new instance of BackupUsageSummariesServerTransport with the provided implementation.
// The returned BackupUsageSummariesServerTransport instance is connected to an instance of armrecoveryservicesbackup.BackupUsageSummariesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBackupUsageSummariesServerTransport(srv *BackupUsageSummariesServer) *BackupUsageSummariesServerTransport {
	return &BackupUsageSummariesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armrecoveryservicesbackup.BackupUsageSummariesClientListResponse]](),
	}
}

// BackupUsageSummariesServerTransport connects instances of armrecoveryservicesbackup.BackupUsageSummariesClient to instances of BackupUsageSummariesServer.
// Don't use this type directly, use NewBackupUsageSummariesServerTransport instead.
type BackupUsageSummariesServerTransport struct {
	srv          *BackupUsageSummariesServer
	newListPager *tracker[azfake.PagerResponder[armrecoveryservicesbackup.BackupUsageSummariesClientListResponse]]
}

// Do implements the policy.Transporter interface for BackupUsageSummariesServerTransport.
func (b *BackupUsageSummariesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BackupUsageSummariesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if backupUsageSummariesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = backupUsageSummariesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "BackupUsageSummariesClient.NewListPager":
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

func (b *BackupUsageSummariesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/Subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupUsageSummaries`
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
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armrecoveryservicesbackup.BackupUsageSummariesClientListOptions
		if filterParam != nil || skipTokenParam != nil {
			options = &armrecoveryservicesbackup.BackupUsageSummariesClientListOptions{
				Filter:    filterParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := b.srv.NewListPager(vaultNameParam, resourceGroupNameParam, options)
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

// set this to conditionally intercept incoming requests to BackupUsageSummariesServerTransport
var backupUsageSummariesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
