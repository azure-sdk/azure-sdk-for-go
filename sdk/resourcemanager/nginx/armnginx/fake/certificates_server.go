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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v3"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// CertificatesServer is a fake server for instances of the armnginx.CertificatesClient type.
type CertificatesServer struct {
	// BeginCreateOrUpdate is the fake for method CertificatesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, deploymentName string, certificateName string, options *armnginx.CertificatesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnginx.CertificatesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method CertificatesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, deploymentName string, certificateName string, options *armnginx.CertificatesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnginx.CertificatesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CertificatesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, deploymentName string, certificateName string, options *armnginx.CertificatesClientGetOptions) (resp azfake.Responder[armnginx.CertificatesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method CertificatesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, deploymentName string, options *armnginx.CertificatesClientListOptions) (resp azfake.PagerResponder[armnginx.CertificatesClientListResponse])
}

// NewCertificatesServerTransport creates a new instance of CertificatesServerTransport with the provided implementation.
// The returned CertificatesServerTransport instance is connected to an instance of armnginx.CertificatesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCertificatesServerTransport(srv *CertificatesServer) *CertificatesServerTransport {
	return &CertificatesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnginx.CertificatesClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnginx.CertificatesClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnginx.CertificatesClientListResponse]](),
	}
}

// CertificatesServerTransport connects instances of armnginx.CertificatesClient to instances of CertificatesServer.
// Don't use this type directly, use NewCertificatesServerTransport instead.
type CertificatesServerTransport struct {
	srv                 *CertificatesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnginx.CertificatesClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnginx.CertificatesClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnginx.CertificatesClientListResponse]]
}

// Do implements the policy.Transporter interface for CertificatesServerTransport.
func (c *CertificatesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *CertificatesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if certificatesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = certificatesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "CertificatesClient.BeginCreateOrUpdate":
				res.resp, res.err = c.dispatchBeginCreateOrUpdate(req)
			case "CertificatesClient.BeginDelete":
				res.resp, res.err = c.dispatchBeginDelete(req)
			case "CertificatesClient.Get":
				res.resp, res.err = c.dispatchGet(req)
			case "CertificatesClient.NewListPager":
				res.resp, res.err = c.dispatchNewListPager(req)
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

func (c *CertificatesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Nginx\.NginxPlus/nginxDeployments/(?P<deploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/certificates/(?P<certificateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnginx.Certificate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		deploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentName")])
		if err != nil {
			return nil, err
		}
		certificateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("certificateName")])
		if err != nil {
			return nil, err
		}
		var options *armnginx.CertificatesClientBeginCreateOrUpdateOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armnginx.CertificatesClientBeginCreateOrUpdateOptions{
				Body: &body,
			}
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, deploymentNameParam, certificateNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		c.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		c.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (c *CertificatesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Nginx\.NginxPlus/nginxDeployments/(?P<deploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/certificates/(?P<certificateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		deploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentName")])
		if err != nil {
			return nil, err
		}
		certificateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("certificateName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, deploymentNameParam, certificateNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *CertificatesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Nginx\.NginxPlus/nginxDeployments/(?P<deploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/certificates/(?P<certificateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	deploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentName")])
	if err != nil {
		return nil, err
	}
	certificateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("certificateName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, deploymentNameParam, certificateNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Certificate, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CertificatesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Nginx\.NginxPlus/nginxDeployments/(?P<deploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/certificates`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		deploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListPager(resourceGroupNameParam, deploymentNameParam, nil)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnginx.CertificatesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		c.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to CertificatesServerTransport
var certificatesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
