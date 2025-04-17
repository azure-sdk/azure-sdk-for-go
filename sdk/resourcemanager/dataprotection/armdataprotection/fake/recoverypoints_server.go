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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
	"net/http"
	"net/url"
	"regexp"
)

// RecoveryPointsServer is a fake server for instances of the armdataprotection.RecoveryPointsClient type.
type RecoveryPointsServer struct {
	// Get is the fake for method RecoveryPointsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vaultName string, backupInstanceName string, recoveryPointID string, options *armdataprotection.RecoveryPointsClientGetOptions) (resp azfake.Responder[armdataprotection.RecoveryPointsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RecoveryPointsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, vaultName string, backupInstanceName string, options *armdataprotection.RecoveryPointsClientListOptions) (resp azfake.PagerResponder[armdataprotection.RecoveryPointsClientListResponse])
}

// NewRecoveryPointsServerTransport creates a new instance of RecoveryPointsServerTransport with the provided implementation.
// The returned RecoveryPointsServerTransport instance is connected to an instance of armdataprotection.RecoveryPointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRecoveryPointsServerTransport(srv *RecoveryPointsServer) *RecoveryPointsServerTransport {
	return &RecoveryPointsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armdataprotection.RecoveryPointsClientListResponse]](),
	}
}

// RecoveryPointsServerTransport connects instances of armdataprotection.RecoveryPointsClient to instances of RecoveryPointsServer.
// Don't use this type directly, use NewRecoveryPointsServerTransport instead.
type RecoveryPointsServerTransport struct {
	srv          *RecoveryPointsServer
	newListPager *tracker[azfake.PagerResponder[armdataprotection.RecoveryPointsClientListResponse]]
}

// Do implements the policy.Transporter interface for RecoveryPointsServerTransport.
func (r *RecoveryPointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *RecoveryPointsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if recoveryPointsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = recoveryPointsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "RecoveryPointsClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "RecoveryPointsClient.NewListPager":
				res.resp, res.err = r.dispatchNewListPager(req)
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

func (r *RecoveryPointsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataProtection/backupVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupInstances/(?P<backupInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/recoveryPoints/(?P<recoveryPointId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	backupInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupInstanceName")])
	if err != nil {
		return nil, err
	}
	recoveryPointIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("recoveryPointId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, vaultNameParam, backupInstanceNameParam, recoveryPointIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureBackupRecoveryPointResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RecoveryPointsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataProtection/backupVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupInstances/(?P<backupInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/recoveryPoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		backupInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupInstanceName")])
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
		var options *armdataprotection.RecoveryPointsClientListOptions
		if filterParam != nil || skipTokenParam != nil {
			options = &armdataprotection.RecoveryPointsClientListOptions{
				Filter:    filterParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := r.srv.NewListPager(resourceGroupNameParam, vaultNameParam, backupInstanceNameParam, options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdataprotection.RecoveryPointsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to RecoveryPointsServerTransport
var recoveryPointsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
