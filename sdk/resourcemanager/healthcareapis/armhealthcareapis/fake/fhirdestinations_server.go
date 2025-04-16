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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v2"
	"net/http"
	"net/url"
	"regexp"
)

// FhirDestinationsServer is a fake server for instances of the armhealthcareapis.FhirDestinationsClient type.
type FhirDestinationsServer struct {
	// NewListByIotConnectorPager is the fake for method FhirDestinationsClient.NewListByIotConnectorPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByIotConnectorPager func(resourceGroupName string, workspaceName string, iotConnectorName string, options *armhealthcareapis.FhirDestinationsClientListByIotConnectorOptions) (resp azfake.PagerResponder[armhealthcareapis.FhirDestinationsClientListByIotConnectorResponse])
}

// NewFhirDestinationsServerTransport creates a new instance of FhirDestinationsServerTransport with the provided implementation.
// The returned FhirDestinationsServerTransport instance is connected to an instance of armhealthcareapis.FhirDestinationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFhirDestinationsServerTransport(srv *FhirDestinationsServer) *FhirDestinationsServerTransport {
	return &FhirDestinationsServerTransport{
		srv:                        srv,
		newListByIotConnectorPager: newTracker[azfake.PagerResponder[armhealthcareapis.FhirDestinationsClientListByIotConnectorResponse]](),
	}
}

// FhirDestinationsServerTransport connects instances of armhealthcareapis.FhirDestinationsClient to instances of FhirDestinationsServer.
// Don't use this type directly, use NewFhirDestinationsServerTransport instead.
type FhirDestinationsServerTransport struct {
	srv                        *FhirDestinationsServer
	newListByIotConnectorPager *tracker[azfake.PagerResponder[armhealthcareapis.FhirDestinationsClientListByIotConnectorResponse]]
}

// Do implements the policy.Transporter interface for FhirDestinationsServerTransport.
func (f *FhirDestinationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return f.dispatchToMethodFake(req, method)
}

func (f *FhirDestinationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if fhirDestinationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = fhirDestinationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "FhirDestinationsClient.NewListByIotConnectorPager":
				res.resp, res.err = f.dispatchNewListByIotConnectorPager(req)
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

func (f *FhirDestinationsServerTransport) dispatchNewListByIotConnectorPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListByIotConnectorPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByIotConnectorPager not implemented")}
	}
	newListByIotConnectorPager := f.newListByIotConnectorPager.get(req)
	if newListByIotConnectorPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthcareApis/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iotconnectors/(?P<iotConnectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fhirdestinations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		iotConnectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iotConnectorName")])
		if err != nil {
			return nil, err
		}
		resp := f.srv.NewListByIotConnectorPager(resourceGroupNameParam, workspaceNameParam, iotConnectorNameParam, nil)
		newListByIotConnectorPager = &resp
		f.newListByIotConnectorPager.add(req, newListByIotConnectorPager)
		server.PagerResponderInjectNextLinks(newListByIotConnectorPager, req, func(page *armhealthcareapis.FhirDestinationsClientListByIotConnectorResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByIotConnectorPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListByIotConnectorPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByIotConnectorPager) {
		f.newListByIotConnectorPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to FhirDestinationsServerTransport
var fhirDestinationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
