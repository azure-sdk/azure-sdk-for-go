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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
	"net/http"
	"net/url"
	"regexp"
)

// ServiceReplicaServer is a fake server for instances of the armservicefabricmesh.ServiceReplicaClient type.
type ServiceReplicaServer struct {
	// Get is the fake for method ServiceReplicaClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string, replicaName string, options *armservicefabricmesh.ServiceReplicaClientGetOptions) (resp azfake.Responder[armservicefabricmesh.ServiceReplicaClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ServiceReplicaClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, applicationResourceName string, serviceResourceName string, options *armservicefabricmesh.ServiceReplicaClientListOptions) (resp azfake.PagerResponder[armservicefabricmesh.ServiceReplicaClientListResponse])
}

// NewServiceReplicaServerTransport creates a new instance of ServiceReplicaServerTransport with the provided implementation.
// The returned ServiceReplicaServerTransport instance is connected to an instance of armservicefabricmesh.ServiceReplicaClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServiceReplicaServerTransport(srv *ServiceReplicaServer) *ServiceReplicaServerTransport {
	return &ServiceReplicaServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armservicefabricmesh.ServiceReplicaClientListResponse]](),
	}
}

// ServiceReplicaServerTransport connects instances of armservicefabricmesh.ServiceReplicaClient to instances of ServiceReplicaServer.
// Don't use this type directly, use NewServiceReplicaServerTransport instead.
type ServiceReplicaServerTransport struct {
	srv          *ServiceReplicaServer
	newListPager *tracker[azfake.PagerResponder[armservicefabricmesh.ServiceReplicaClientListResponse]]
}

// Do implements the policy.Transporter interface for ServiceReplicaServerTransport.
func (s *ServiceReplicaServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServiceReplicaServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if serviceReplicaServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = serviceReplicaServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ServiceReplicaClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "ServiceReplicaClient.NewListPager":
				res.resp, res.err = s.dispatchNewListPager(req)
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

func (s *ServiceReplicaServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabricMesh/applications/(?P<applicationResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/services/(?P<serviceResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicas/(?P<replicaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	applicationResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationResourceName")])
	if err != nil {
		return nil, err
	}
	serviceResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceResourceName")])
	if err != nil {
		return nil, err
	}
	replicaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("replicaName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, applicationResourceNameParam, serviceResourceNameParam, replicaNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServiceReplicaDescription, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServiceReplicaServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabricMesh/applications/(?P<applicationResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/services/(?P<serviceResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicas`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		applicationResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationResourceName")])
		if err != nil {
			return nil, err
		}
		serviceResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceResourceName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListPager(resourceGroupNameParam, applicationResourceNameParam, serviceResourceNameParam, nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armservicefabricmesh.ServiceReplicaClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ServiceReplicaServerTransport
var serviceReplicaServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
