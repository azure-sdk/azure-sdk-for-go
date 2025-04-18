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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
	"net/http"
	"net/url"
	"regexp"
)

// CustomizedAcceleratorsServer is a fake server for instances of the armappplatform.CustomizedAcceleratorsClient type.
type CustomizedAcceleratorsServer struct {
	// BeginCreateOrUpdate is the fake for method CustomizedAcceleratorsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, customizedAcceleratorResource armappplatform.CustomizedAcceleratorResource, options *armappplatform.CustomizedAcceleratorsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armappplatform.CustomizedAcceleratorsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method CustomizedAcceleratorsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, options *armappplatform.CustomizedAcceleratorsClientBeginDeleteOptions) (resp azfake.PollerResponder[armappplatform.CustomizedAcceleratorsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CustomizedAcceleratorsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, options *armappplatform.CustomizedAcceleratorsClientGetOptions) (resp azfake.Responder[armappplatform.CustomizedAcceleratorsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method CustomizedAcceleratorsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, serviceName string, applicationAcceleratorName string, options *armappplatform.CustomizedAcceleratorsClientListOptions) (resp azfake.PagerResponder[armappplatform.CustomizedAcceleratorsClientListResponse])

	// BeginValidate is the fake for method CustomizedAcceleratorsClient.BeginValidate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginValidate func(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, properties armappplatform.CustomizedAcceleratorProperties, options *armappplatform.CustomizedAcceleratorsClientBeginValidateOptions) (resp azfake.PollerResponder[armappplatform.CustomizedAcceleratorsClientValidateResponse], errResp azfake.ErrorResponder)
}

// NewCustomizedAcceleratorsServerTransport creates a new instance of CustomizedAcceleratorsServerTransport with the provided implementation.
// The returned CustomizedAcceleratorsServerTransport instance is connected to an instance of armappplatform.CustomizedAcceleratorsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCustomizedAcceleratorsServerTransport(srv *CustomizedAcceleratorsServer) *CustomizedAcceleratorsServerTransport {
	return &CustomizedAcceleratorsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armappplatform.CustomizedAcceleratorsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armappplatform.CustomizedAcceleratorsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armappplatform.CustomizedAcceleratorsClientListResponse]](),
		beginValidate:       newTracker[azfake.PollerResponder[armappplatform.CustomizedAcceleratorsClientValidateResponse]](),
	}
}

// CustomizedAcceleratorsServerTransport connects instances of armappplatform.CustomizedAcceleratorsClient to instances of CustomizedAcceleratorsServer.
// Don't use this type directly, use NewCustomizedAcceleratorsServerTransport instead.
type CustomizedAcceleratorsServerTransport struct {
	srv                 *CustomizedAcceleratorsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armappplatform.CustomizedAcceleratorsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armappplatform.CustomizedAcceleratorsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armappplatform.CustomizedAcceleratorsClientListResponse]]
	beginValidate       *tracker[azfake.PollerResponder[armappplatform.CustomizedAcceleratorsClientValidateResponse]]
}

// Do implements the policy.Transporter interface for CustomizedAcceleratorsServerTransport.
func (c *CustomizedAcceleratorsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *CustomizedAcceleratorsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if customizedAcceleratorsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = customizedAcceleratorsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "CustomizedAcceleratorsClient.BeginCreateOrUpdate":
				res.resp, res.err = c.dispatchBeginCreateOrUpdate(req)
			case "CustomizedAcceleratorsClient.BeginDelete":
				res.resp, res.err = c.dispatchBeginDelete(req)
			case "CustomizedAcceleratorsClient.Get":
				res.resp, res.err = c.dispatchGet(req)
			case "CustomizedAcceleratorsClient.NewListPager":
				res.resp, res.err = c.dispatchNewListPager(req)
			case "CustomizedAcceleratorsClient.BeginValidate":
				res.resp, res.err = c.dispatchBeginValidate(req)
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

func (c *CustomizedAcceleratorsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationAccelerators/(?P<applicationAcceleratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/customizedAccelerators/(?P<customizedAcceleratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappplatform.CustomizedAcceleratorResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		applicationAcceleratorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationAcceleratorName")])
		if err != nil {
			return nil, err
		}
		customizedAcceleratorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("customizedAcceleratorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, applicationAcceleratorNameParam, customizedAcceleratorNameParam, body, nil)
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

func (c *CustomizedAcceleratorsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationAccelerators/(?P<applicationAcceleratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/customizedAccelerators/(?P<customizedAcceleratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		applicationAcceleratorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationAcceleratorName")])
		if err != nil {
			return nil, err
		}
		customizedAcceleratorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("customizedAcceleratorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, serviceNameParam, applicationAcceleratorNameParam, customizedAcceleratorNameParam, nil)
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

func (c *CustomizedAcceleratorsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationAccelerators/(?P<applicationAcceleratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/customizedAccelerators/(?P<customizedAcceleratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	applicationAcceleratorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationAcceleratorName")])
	if err != nil {
		return nil, err
	}
	customizedAcceleratorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("customizedAcceleratorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, applicationAcceleratorNameParam, customizedAcceleratorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CustomizedAcceleratorResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CustomizedAcceleratorsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationAccelerators/(?P<applicationAcceleratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/customizedAccelerators`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		applicationAcceleratorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationAcceleratorName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListPager(resourceGroupNameParam, serviceNameParam, applicationAcceleratorNameParam, nil)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappplatform.CustomizedAcceleratorsClientListResponse, createLink func() string) {
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

func (c *CustomizedAcceleratorsServerTransport) dispatchBeginValidate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginValidate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginValidate not implemented")}
	}
	beginValidate := c.beginValidate.get(req)
	if beginValidate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationAccelerators/(?P<applicationAcceleratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/customizedAccelerators/(?P<customizedAcceleratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/validate`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappplatform.CustomizedAcceleratorProperties](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		applicationAcceleratorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationAcceleratorName")])
		if err != nil {
			return nil, err
		}
		customizedAcceleratorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("customizedAcceleratorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginValidate(req.Context(), resourceGroupNameParam, serviceNameParam, applicationAcceleratorNameParam, customizedAcceleratorNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginValidate = &respr
		c.beginValidate.add(req, beginValidate)
	}

	resp, err := server.PollerResponderNext(beginValidate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginValidate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginValidate) {
		c.beginValidate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to CustomizedAcceleratorsServerTransport
var customizedAcceleratorsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
