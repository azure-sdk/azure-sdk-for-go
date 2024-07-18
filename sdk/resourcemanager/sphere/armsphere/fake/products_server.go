// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sphere/armsphere"
	"net/http"
	"net/url"
	"regexp"
)

// ProductsServer is a fake server for instances of the armsphere.ProductsClient type.
type ProductsServer struct {
	// CountDevices is the fake for method ProductsClient.CountDevices
	// HTTP status codes to indicate success: http.StatusOK
	CountDevices func(ctx context.Context, resourceGroupName string, catalogName string, productName string, options *armsphere.ProductsClientCountDevicesOptions) (resp azfake.Responder[armsphere.ProductsClientCountDevicesResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method ProductsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, catalogName string, productName string, resource armsphere.Product, options *armsphere.ProductsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsphere.ProductsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ProductsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, catalogName string, productName string, options *armsphere.ProductsClientBeginDeleteOptions) (resp azfake.PollerResponder[armsphere.ProductsClientDeleteResponse], errResp azfake.ErrorResponder)

	// NewGenerateDefaultDeviceGroupsPager is the fake for method ProductsClient.NewGenerateDefaultDeviceGroupsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGenerateDefaultDeviceGroupsPager func(resourceGroupName string, catalogName string, productName string, options *armsphere.ProductsClientGenerateDefaultDeviceGroupsOptions) (resp azfake.PagerResponder[armsphere.ProductsClientGenerateDefaultDeviceGroupsResponse])

	// Get is the fake for method ProductsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, catalogName string, productName string, options *armsphere.ProductsClientGetOptions) (resp azfake.Responder[armsphere.ProductsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByCatalogPager is the fake for method ProductsClient.NewListByCatalogPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByCatalogPager func(resourceGroupName string, catalogName string, options *armsphere.ProductsClientListByCatalogOptions) (resp azfake.PagerResponder[armsphere.ProductsClientListByCatalogResponse])

	// BeginUpdate is the fake for method ProductsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, catalogName string, productName string, properties armsphere.ProductUpdate, options *armsphere.ProductsClientBeginUpdateOptions) (resp azfake.PollerResponder[armsphere.ProductsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewProductsServerTransport creates a new instance of ProductsServerTransport with the provided implementation.
// The returned ProductsServerTransport instance is connected to an instance of armsphere.ProductsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProductsServerTransport(srv *ProductsServer) *ProductsServerTransport {
	return &ProductsServerTransport{
		srv:                                 srv,
		beginCreateOrUpdate:                 newTracker[azfake.PollerResponder[armsphere.ProductsClientCreateOrUpdateResponse]](),
		beginDelete:                         newTracker[azfake.PollerResponder[armsphere.ProductsClientDeleteResponse]](),
		newGenerateDefaultDeviceGroupsPager: newTracker[azfake.PagerResponder[armsphere.ProductsClientGenerateDefaultDeviceGroupsResponse]](),
		newListByCatalogPager:               newTracker[azfake.PagerResponder[armsphere.ProductsClientListByCatalogResponse]](),
		beginUpdate:                         newTracker[azfake.PollerResponder[armsphere.ProductsClientUpdateResponse]](),
	}
}

// ProductsServerTransport connects instances of armsphere.ProductsClient to instances of ProductsServer.
// Don't use this type directly, use NewProductsServerTransport instead.
type ProductsServerTransport struct {
	srv                                 *ProductsServer
	beginCreateOrUpdate                 *tracker[azfake.PollerResponder[armsphere.ProductsClientCreateOrUpdateResponse]]
	beginDelete                         *tracker[azfake.PollerResponder[armsphere.ProductsClientDeleteResponse]]
	newGenerateDefaultDeviceGroupsPager *tracker[azfake.PagerResponder[armsphere.ProductsClientGenerateDefaultDeviceGroupsResponse]]
	newListByCatalogPager               *tracker[azfake.PagerResponder[armsphere.ProductsClientListByCatalogResponse]]
	beginUpdate                         *tracker[azfake.PollerResponder[armsphere.ProductsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ProductsServerTransport.
func (p *ProductsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *ProductsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ProductsClient.CountDevices":
		resp, err = p.dispatchCountDevices(req)
	case "ProductsClient.BeginCreateOrUpdate":
		resp, err = p.dispatchBeginCreateOrUpdate(req)
	case "ProductsClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "ProductsClient.NewGenerateDefaultDeviceGroupsPager":
		resp, err = p.dispatchNewGenerateDefaultDeviceGroupsPager(req)
	case "ProductsClient.Get":
		resp, err = p.dispatchGet(req)
	case "ProductsClient.NewListByCatalogPager":
		resp, err = p.dispatchNewListByCatalogPager(req)
	case "ProductsClient.BeginUpdate":
		resp, err = p.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (p *ProductsServerTransport) dispatchCountDevices(req *http.Request) (*http.Response, error) {
	if p.srv.CountDevices == nil {
		return nil, &nonRetriableError{errors.New("fake for method CountDevices not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/countDevices`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
	if err != nil {
		return nil, err
	}
	productNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("productName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.CountDevices(req.Context(), resourceGroupNameParam, catalogNameParam, productNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CountDevicesResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProductsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsphere.Product](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		productNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("productName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, catalogNameParam, productNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		p.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		p.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		p.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (p *ProductsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		productNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("productName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, catalogNameParam, productNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		p.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *ProductsServerTransport) dispatchNewGenerateDefaultDeviceGroupsPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGenerateDefaultDeviceGroupsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGenerateDefaultDeviceGroupsPager not implemented")}
	}
	newGenerateDefaultDeviceGroupsPager := p.newGenerateDefaultDeviceGroupsPager.get(req)
	if newGenerateDefaultDeviceGroupsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/generateDefaultDeviceGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		productNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("productName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewGenerateDefaultDeviceGroupsPager(resourceGroupNameParam, catalogNameParam, productNameParam, nil)
		newGenerateDefaultDeviceGroupsPager = &resp
		p.newGenerateDefaultDeviceGroupsPager.add(req, newGenerateDefaultDeviceGroupsPager)
		server.PagerResponderInjectNextLinks(newGenerateDefaultDeviceGroupsPager, req, func(page *armsphere.ProductsClientGenerateDefaultDeviceGroupsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newGenerateDefaultDeviceGroupsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newGenerateDefaultDeviceGroupsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGenerateDefaultDeviceGroupsPager) {
		p.newGenerateDefaultDeviceGroupsPager.remove(req)
	}
	return resp, nil
}

func (p *ProductsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
	if err != nil {
		return nil, err
	}
	productNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("productName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, catalogNameParam, productNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Product, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProductsServerTransport) dispatchNewListByCatalogPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByCatalogPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByCatalogPager not implemented")}
	}
	newListByCatalogPager := p.newListByCatalogPager.get(req)
	if newListByCatalogPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByCatalogPager(resourceGroupNameParam, catalogNameParam, nil)
		newListByCatalogPager = &resp
		p.newListByCatalogPager.add(req, newListByCatalogPager)
		server.PagerResponderInjectNextLinks(newListByCatalogPager, req, func(page *armsphere.ProductsClientListByCatalogResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByCatalogPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByCatalogPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByCatalogPager) {
		p.newListByCatalogPager.remove(req)
	}
	return resp, nil
}

func (p *ProductsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := p.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsphere.ProductUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		productNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("productName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginUpdate(req.Context(), resourceGroupNameParam, catalogNameParam, productNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		p.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		p.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		p.beginUpdate.remove(req)
	}

	return resp, nil
}
