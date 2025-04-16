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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
	"net/http"
	"net/url"
	"regexp"
)

// EndpointDeploymentServer is a fake server for instances of the armmachinelearning.EndpointDeploymentClient type.
type EndpointDeploymentServer struct {
	// BeginCreateOrUpdate is the fake for method EndpointDeploymentClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body armmachinelearning.EndpointDeploymentResourcePropertiesBasicResource, options *armmachinelearning.EndpointDeploymentClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armmachinelearning.EndpointDeploymentClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method EndpointDeploymentClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *armmachinelearning.EndpointDeploymentClientBeginDeleteOptions) (resp azfake.PollerResponder[armmachinelearning.EndpointDeploymentClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method EndpointDeploymentClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *armmachinelearning.EndpointDeploymentClientGetOptions) (resp azfake.Responder[armmachinelearning.EndpointDeploymentClientGetResponse], errResp azfake.ErrorResponder)

	// NewGetInWorkspacePager is the fake for method EndpointDeploymentClient.NewGetInWorkspacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetInWorkspacePager func(resourceGroupName string, workspaceName string, options *armmachinelearning.EndpointDeploymentClientGetInWorkspaceOptions) (resp azfake.PagerResponder[armmachinelearning.EndpointDeploymentClientGetInWorkspaceResponse])

	// NewListPager is the fake for method EndpointDeploymentClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, endpointName string, options *armmachinelearning.EndpointDeploymentClientListOptions) (resp azfake.PagerResponder[armmachinelearning.EndpointDeploymentClientListResponse])
}

// NewEndpointDeploymentServerTransport creates a new instance of EndpointDeploymentServerTransport with the provided implementation.
// The returned EndpointDeploymentServerTransport instance is connected to an instance of armmachinelearning.EndpointDeploymentClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEndpointDeploymentServerTransport(srv *EndpointDeploymentServer) *EndpointDeploymentServerTransport {
	return &EndpointDeploymentServerTransport{
		srv:                    srv,
		beginCreateOrUpdate:    newTracker[azfake.PollerResponder[armmachinelearning.EndpointDeploymentClientCreateOrUpdateResponse]](),
		beginDelete:            newTracker[azfake.PollerResponder[armmachinelearning.EndpointDeploymentClientDeleteResponse]](),
		newGetInWorkspacePager: newTracker[azfake.PagerResponder[armmachinelearning.EndpointDeploymentClientGetInWorkspaceResponse]](),
		newListPager:           newTracker[azfake.PagerResponder[armmachinelearning.EndpointDeploymentClientListResponse]](),
	}
}

// EndpointDeploymentServerTransport connects instances of armmachinelearning.EndpointDeploymentClient to instances of EndpointDeploymentServer.
// Don't use this type directly, use NewEndpointDeploymentServerTransport instead.
type EndpointDeploymentServerTransport struct {
	srv                    *EndpointDeploymentServer
	beginCreateOrUpdate    *tracker[azfake.PollerResponder[armmachinelearning.EndpointDeploymentClientCreateOrUpdateResponse]]
	beginDelete            *tracker[azfake.PollerResponder[armmachinelearning.EndpointDeploymentClientDeleteResponse]]
	newGetInWorkspacePager *tracker[azfake.PagerResponder[armmachinelearning.EndpointDeploymentClientGetInWorkspaceResponse]]
	newListPager           *tracker[azfake.PagerResponder[armmachinelearning.EndpointDeploymentClientListResponse]]
}

// Do implements the policy.Transporter interface for EndpointDeploymentServerTransport.
func (e *EndpointDeploymentServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *EndpointDeploymentServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if endpointDeploymentServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = endpointDeploymentServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "EndpointDeploymentClient.BeginCreateOrUpdate":
				res.resp, res.err = e.dispatchBeginCreateOrUpdate(req)
			case "EndpointDeploymentClient.BeginDelete":
				res.resp, res.err = e.dispatchBeginDelete(req)
			case "EndpointDeploymentClient.Get":
				res.resp, res.err = e.dispatchGet(req)
			case "EndpointDeploymentClient.NewGetInWorkspacePager":
				res.resp, res.err = e.dispatchNewGetInWorkspacePager(req)
			case "EndpointDeploymentClient.NewListPager":
				res.resp, res.err = e.dispatchNewListPager(req)
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

func (e *EndpointDeploymentServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := e.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deployments/(?P<deploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmachinelearning.EndpointDeploymentResourcePropertiesBasicResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		deploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, endpointNameParam, deploymentNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		e.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		e.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		e.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (e *EndpointDeploymentServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if e.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := e.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deployments/(?P<deploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		deploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentName")])
		if err != nil {
			return nil, err
		}
		proxyAPIVersionUnescaped, err := url.QueryUnescape(qp.Get("proxy-api-version"))
		if err != nil {
			return nil, err
		}
		proxyAPIVersionParam := getOptional(proxyAPIVersionUnescaped)
		var options *armmachinelearning.EndpointDeploymentClientBeginDeleteOptions
		if proxyAPIVersionParam != nil {
			options = &armmachinelearning.EndpointDeploymentClientBeginDeleteOptions{
				ProxyAPIVersion: proxyAPIVersionParam,
			}
		}
		respr, errRespr := e.srv.BeginDelete(req.Context(), resourceGroupNameParam, workspaceNameParam, endpointNameParam, deploymentNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		e.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		e.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		e.beginDelete.remove(req)
	}

	return resp, nil
}

func (e *EndpointDeploymentServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deployments/(?P<deploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	deploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, endpointNameParam, deploymentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EndpointDeploymentResourcePropertiesBasicResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EndpointDeploymentServerTransport) dispatchNewGetInWorkspacePager(req *http.Request) (*http.Response, error) {
	if e.srv.NewGetInWorkspacePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetInWorkspacePager not implemented")}
	}
	newGetInWorkspacePager := e.newGetInWorkspacePager.get(req)
	if newGetInWorkspacePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deployments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		endpointTypeUnescaped, err := url.QueryUnescape(qp.Get("endpointType"))
		if err != nil {
			return nil, err
		}
		endpointTypeParam := getOptional(armmachinelearning.EndpointType(endpointTypeUnescaped))
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam := getOptional(skipUnescaped)
		var options *armmachinelearning.EndpointDeploymentClientGetInWorkspaceOptions
		if endpointTypeParam != nil || skipParam != nil {
			options = &armmachinelearning.EndpointDeploymentClientGetInWorkspaceOptions{
				EndpointType: endpointTypeParam,
				Skip:         skipParam,
			}
		}
		resp := e.srv.NewGetInWorkspacePager(resourceGroupNameParam, workspaceNameParam, options)
		newGetInWorkspacePager = &resp
		e.newGetInWorkspacePager.add(req, newGetInWorkspacePager)
		server.PagerResponderInjectNextLinks(newGetInWorkspacePager, req, func(page *armmachinelearning.EndpointDeploymentClientGetInWorkspaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newGetInWorkspacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newGetInWorkspacePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGetInWorkspacePager) {
		e.newGetInWorkspacePager.remove(req)
	}
	return resp, nil
}

func (e *EndpointDeploymentServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := e.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deployments`
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
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, endpointNameParam, nil)
		newListPager = &resp
		e.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmachinelearning.EndpointDeploymentClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		e.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to EndpointDeploymentServerTransport
var endpointDeploymentServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
