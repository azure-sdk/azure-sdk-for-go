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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
	"net/http"
	"net/url"
	"regexp"
)

// FetchSecondaryRecoveryPointsServer is a fake server for instances of the armdataprotection.FetchSecondaryRecoveryPointsClient type.
type FetchSecondaryRecoveryPointsServer struct {
	// NewListPager is the fake for method FetchSecondaryRecoveryPointsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, location string, parameters armdataprotection.FetchSecondaryRPsRequestParameters, options *armdataprotection.FetchSecondaryRecoveryPointsClientListOptions) (resp azfake.PagerResponder[armdataprotection.FetchSecondaryRecoveryPointsClientListResponse])
}

// NewFetchSecondaryRecoveryPointsServerTransport creates a new instance of FetchSecondaryRecoveryPointsServerTransport with the provided implementation.
// The returned FetchSecondaryRecoveryPointsServerTransport instance is connected to an instance of armdataprotection.FetchSecondaryRecoveryPointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFetchSecondaryRecoveryPointsServerTransport(srv *FetchSecondaryRecoveryPointsServer) *FetchSecondaryRecoveryPointsServerTransport {
	return &FetchSecondaryRecoveryPointsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armdataprotection.FetchSecondaryRecoveryPointsClientListResponse]](),
	}
}

// FetchSecondaryRecoveryPointsServerTransport connects instances of armdataprotection.FetchSecondaryRecoveryPointsClient to instances of FetchSecondaryRecoveryPointsServer.
// Don't use this type directly, use NewFetchSecondaryRecoveryPointsServerTransport instead.
type FetchSecondaryRecoveryPointsServerTransport struct {
	srv          *FetchSecondaryRecoveryPointsServer
	newListPager *tracker[azfake.PagerResponder[armdataprotection.FetchSecondaryRecoveryPointsClientListResponse]]
}

// Do implements the policy.Transporter interface for FetchSecondaryRecoveryPointsServerTransport.
func (f *FetchSecondaryRecoveryPointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return f.dispatchToMethodFake(req, method)
}

func (f *FetchSecondaryRecoveryPointsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if fetchSecondaryRecoveryPointsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = fetchSecondaryRecoveryPointsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "FetchSecondaryRecoveryPointsClient.NewListPager":
				res.resp, res.err = f.dispatchNewListPager(req)
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

func (f *FetchSecondaryRecoveryPointsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := f.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataProtection/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fetchSecondaryRecoveryPoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[armdataprotection.FetchSecondaryRPsRequestParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
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
		var options *armdataprotection.FetchSecondaryRecoveryPointsClientListOptions
		if filterParam != nil || skipTokenParam != nil {
			options = &armdataprotection.FetchSecondaryRecoveryPointsClientListOptions{
				Filter:    filterParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := f.srv.NewListPager(resourceGroupNameParam, locationParam, body, options)
		newListPager = &resp
		f.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdataprotection.FetchSecondaryRecoveryPointsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		f.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to FetchSecondaryRecoveryPointsServerTransport
var fetchSecondaryRecoveryPointsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
