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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationdiscovery/armmigrationdiscovery"
	"net/http"
	"net/url"
	"regexp"
)

// ServerSoftwareInventoriesControllerServer is a fake server for instances of the armmigrationdiscovery.ServerSoftwareInventoriesControllerClient type.
type ServerSoftwareInventoriesControllerServer struct {
	// GetMachineSoftwareInventory is the fake for method ServerSoftwareInventoriesControllerClient.GetMachineSoftwareInventory
	// HTTP status codes to indicate success: http.StatusOK
	GetMachineSoftwareInventory func(ctx context.Context, resourceGroupName string, siteName string, machineName string, defaultParam armmigrationdiscovery.Default, options *armmigrationdiscovery.ServerSoftwareInventoriesControllerClientGetMachineSoftwareInventoryOptions) (resp azfake.Responder[armmigrationdiscovery.ServerSoftwareInventoriesControllerClientGetMachineSoftwareInventoryResponse], errResp azfake.ErrorResponder)

	// NewListByServerPager is the fake for method ServerSoftwareInventoriesControllerClient.NewListByServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServerPager func(resourceGroupName string, siteName string, machineName string, options *armmigrationdiscovery.ServerSoftwareInventoriesControllerClientListByServerOptions) (resp azfake.PagerResponder[armmigrationdiscovery.ServerSoftwareInventoriesControllerClientListByServerResponse])
}

// NewServerSoftwareInventoriesControllerServerTransport creates a new instance of ServerSoftwareInventoriesControllerServerTransport with the provided implementation.
// The returned ServerSoftwareInventoriesControllerServerTransport instance is connected to an instance of armmigrationdiscovery.ServerSoftwareInventoriesControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerSoftwareInventoriesControllerServerTransport(srv *ServerSoftwareInventoriesControllerServer) *ServerSoftwareInventoriesControllerServerTransport {
	return &ServerSoftwareInventoriesControllerServerTransport{
		srv:                  srv,
		newListByServerPager: newTracker[azfake.PagerResponder[armmigrationdiscovery.ServerSoftwareInventoriesControllerClientListByServerResponse]](),
	}
}

// ServerSoftwareInventoriesControllerServerTransport connects instances of armmigrationdiscovery.ServerSoftwareInventoriesControllerClient to instances of ServerSoftwareInventoriesControllerServer.
// Don't use this type directly, use NewServerSoftwareInventoriesControllerServerTransport instead.
type ServerSoftwareInventoriesControllerServerTransport struct {
	srv                  *ServerSoftwareInventoriesControllerServer
	newListByServerPager *tracker[azfake.PagerResponder[armmigrationdiscovery.ServerSoftwareInventoriesControllerClientListByServerResponse]]
}

// Do implements the policy.Transporter interface for ServerSoftwareInventoriesControllerServerTransport.
func (s *ServerSoftwareInventoriesControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServerSoftwareInventoriesControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if serverSoftwareInventoriesControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = serverSoftwareInventoriesControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ServerSoftwareInventoriesControllerClient.GetMachineSoftwareInventory":
				res.resp, res.err = s.dispatchGetMachineSoftwareInventory(req)
			case "ServerSoftwareInventoriesControllerClient.NewListByServerPager":
				res.resp, res.err = s.dispatchNewListByServerPager(req)
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

func (s *ServerSoftwareInventoriesControllerServerTransport) dispatchGetMachineSoftwareInventory(req *http.Request) (*http.Response, error) {
	if s.srv.GetMachineSoftwareInventory == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetMachineSoftwareInventory not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/softwareInventories/(?P<default>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
	if err != nil {
		return nil, err
	}
	defaultParamParam, err := parseWithCast(matches[regex.SubexpIndex("default")], func(v string) (armmigrationdiscovery.Default, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armmigrationdiscovery.Default(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetMachineSoftwareInventory(req.Context(), resourceGroupNameParam, siteNameParam, machineNameParam, defaultParamParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServerSoftwareInventory, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerSoftwareInventoriesControllerServerTransport) dispatchNewListByServerPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServerPager not implemented")}
	}
	newListByServerPager := s.newListByServerPager.get(req)
	if newListByServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/softwareinventories`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByServerPager(resourceGroupNameParam, siteNameParam, machineNameParam, nil)
		newListByServerPager = &resp
		s.newListByServerPager.add(req, newListByServerPager)
		server.PagerResponderInjectNextLinks(newListByServerPager, req, func(page *armmigrationdiscovery.ServerSoftwareInventoriesControllerClientListByServerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServerPager) {
		s.newListByServerPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ServerSoftwareInventoriesControllerServerTransport
var serverSoftwareInventoriesControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
