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

// BackupsUnderVolumeServer is a fake server for instances of the armnetapp.BackupsUnderVolumeClient type.
type BackupsUnderVolumeServer struct {
	// BeginMigrateBackups is the fake for method BackupsUnderVolumeClient.BeginMigrateBackups
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginMigrateBackups func(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, body armnetapp.BackupsMigrationRequest, options *armnetapp.BackupsUnderVolumeClientBeginMigrateBackupsOptions) (resp azfake.PollerResponder[armnetapp.BackupsUnderVolumeClientMigrateBackupsResponse], errResp azfake.ErrorResponder)
}

// NewBackupsUnderVolumeServerTransport creates a new instance of BackupsUnderVolumeServerTransport with the provided implementation.
// The returned BackupsUnderVolumeServerTransport instance is connected to an instance of armnetapp.BackupsUnderVolumeClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBackupsUnderVolumeServerTransport(srv *BackupsUnderVolumeServer) *BackupsUnderVolumeServerTransport {
	return &BackupsUnderVolumeServerTransport{
		srv:                 srv,
		beginMigrateBackups: newTracker[azfake.PollerResponder[armnetapp.BackupsUnderVolumeClientMigrateBackupsResponse]](),
	}
}

// BackupsUnderVolumeServerTransport connects instances of armnetapp.BackupsUnderVolumeClient to instances of BackupsUnderVolumeServer.
// Don't use this type directly, use NewBackupsUnderVolumeServerTransport instead.
type BackupsUnderVolumeServerTransport struct {
	srv                 *BackupsUnderVolumeServer
	beginMigrateBackups *tracker[azfake.PollerResponder[armnetapp.BackupsUnderVolumeClientMigrateBackupsResponse]]
}

// Do implements the policy.Transporter interface for BackupsUnderVolumeServerTransport.
func (b *BackupsUnderVolumeServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BackupsUnderVolumeServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if backupsUnderVolumeServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = backupsUnderVolumeServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "BackupsUnderVolumeClient.BeginMigrateBackups":
				res.resp, res.err = b.dispatchBeginMigrateBackups(req)
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

func (b *BackupsUnderVolumeServerTransport) dispatchBeginMigrateBackups(req *http.Request) (*http.Response, error) {
	if b.srv.BeginMigrateBackups == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginMigrateBackups not implemented")}
	}
	beginMigrateBackups := b.beginMigrateBackups.get(req)
	if beginMigrateBackups == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/capacityPools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumes/(?P<volumeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/migrateBackups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetapp.BackupsMigrationRequest](req)
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
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		volumeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginMigrateBackups(req.Context(), resourceGroupNameParam, accountNameParam, poolNameParam, volumeNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginMigrateBackups = &respr
		b.beginMigrateBackups.add(req, beginMigrateBackups)
	}

	resp, err := server.PollerResponderNext(beginMigrateBackups, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		b.beginMigrateBackups.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginMigrateBackups) {
		b.beginMigrateBackups.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to BackupsUnderVolumeServerTransport
var backupsUnderVolumeServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
