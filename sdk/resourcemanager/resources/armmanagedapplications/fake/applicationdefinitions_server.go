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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications"
	"net/http"
	"net/url"
	"regexp"
)

// ApplicationDefinitionsServer is a fake server for instances of the armmanagedapplications.ApplicationDefinitionsClient type.
type ApplicationDefinitionsServer struct {
	// BeginCreateOrUpdate is the fake for method ApplicationDefinitionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters armmanagedapplications.ApplicationDefinition, options *armmanagedapplications.ApplicationDefinitionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armmanagedapplications.ApplicationDefinitionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdateByID is the fake for method ApplicationDefinitionsClient.BeginCreateOrUpdateByID
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdateByID func(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters armmanagedapplications.ApplicationDefinition, options *armmanagedapplications.ApplicationDefinitionsClientBeginCreateOrUpdateByIDOptions) (resp azfake.PollerResponder[armmanagedapplications.ApplicationDefinitionsClientCreateOrUpdateByIDResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ApplicationDefinitionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *armmanagedapplications.ApplicationDefinitionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armmanagedapplications.ApplicationDefinitionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDeleteByID is the fake for method ApplicationDefinitionsClient.BeginDeleteByID
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDeleteByID func(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *armmanagedapplications.ApplicationDefinitionsClientBeginDeleteByIDOptions) (resp azfake.PollerResponder[armmanagedapplications.ApplicationDefinitionsClientDeleteByIDResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ApplicationDefinitionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotFound
	Get func(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *armmanagedapplications.ApplicationDefinitionsClientGetOptions) (resp azfake.Responder[armmanagedapplications.ApplicationDefinitionsClientGetResponse], errResp azfake.ErrorResponder)

	// GetByID is the fake for method ApplicationDefinitionsClient.GetByID
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotFound
	GetByID func(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *armmanagedapplications.ApplicationDefinitionsClientGetByIDOptions) (resp azfake.Responder[armmanagedapplications.ApplicationDefinitionsClientGetByIDResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method ApplicationDefinitionsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armmanagedapplications.ApplicationDefinitionsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armmanagedapplications.ApplicationDefinitionsClientListByResourceGroupResponse])
}

// NewApplicationDefinitionsServerTransport creates a new instance of ApplicationDefinitionsServerTransport with the provided implementation.
// The returned ApplicationDefinitionsServerTransport instance is connected to an instance of armmanagedapplications.ApplicationDefinitionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewApplicationDefinitionsServerTransport(srv *ApplicationDefinitionsServer) *ApplicationDefinitionsServerTransport {
	return &ApplicationDefinitionsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armmanagedapplications.ApplicationDefinitionsClientCreateOrUpdateResponse]](),
		beginCreateOrUpdateByID:     newTracker[azfake.PollerResponder[armmanagedapplications.ApplicationDefinitionsClientCreateOrUpdateByIDResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armmanagedapplications.ApplicationDefinitionsClientDeleteResponse]](),
		beginDeleteByID:             newTracker[azfake.PollerResponder[armmanagedapplications.ApplicationDefinitionsClientDeleteByIDResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armmanagedapplications.ApplicationDefinitionsClientListByResourceGroupResponse]](),
	}
}

// ApplicationDefinitionsServerTransport connects instances of armmanagedapplications.ApplicationDefinitionsClient to instances of ApplicationDefinitionsServer.
// Don't use this type directly, use NewApplicationDefinitionsServerTransport instead.
type ApplicationDefinitionsServerTransport struct {
	srv                         *ApplicationDefinitionsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armmanagedapplications.ApplicationDefinitionsClientCreateOrUpdateResponse]]
	beginCreateOrUpdateByID     *tracker[azfake.PollerResponder[armmanagedapplications.ApplicationDefinitionsClientCreateOrUpdateByIDResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armmanagedapplications.ApplicationDefinitionsClientDeleteResponse]]
	beginDeleteByID             *tracker[azfake.PollerResponder[armmanagedapplications.ApplicationDefinitionsClientDeleteByIDResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armmanagedapplications.ApplicationDefinitionsClientListByResourceGroupResponse]]
}

// Do implements the policy.Transporter interface for ApplicationDefinitionsServerTransport.
func (a *ApplicationDefinitionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *ApplicationDefinitionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if applicationDefinitionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = applicationDefinitionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ApplicationDefinitionsClient.BeginCreateOrUpdate":
				res.resp, res.err = a.dispatchBeginCreateOrUpdate(req)
			case "ApplicationDefinitionsClient.BeginCreateOrUpdateByID":
				res.resp, res.err = a.dispatchBeginCreateOrUpdateByID(req)
			case "ApplicationDefinitionsClient.BeginDelete":
				res.resp, res.err = a.dispatchBeginDelete(req)
			case "ApplicationDefinitionsClient.BeginDeleteByID":
				res.resp, res.err = a.dispatchBeginDeleteByID(req)
			case "ApplicationDefinitionsClient.Get":
				res.resp, res.err = a.dispatchGet(req)
			case "ApplicationDefinitionsClient.GetByID":
				res.resp, res.err = a.dispatchGetByID(req)
			case "ApplicationDefinitionsClient.NewListByResourceGroupPager":
				res.resp, res.err = a.dispatchNewListByResourceGroupPager(req)
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

func (a *ApplicationDefinitionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := a.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Solutions/applicationDefinitions/(?P<applicationDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmanagedapplications.ApplicationDefinition](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		applicationDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationDefinitionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, applicationDefinitionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		a.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		a.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		a.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (a *ApplicationDefinitionsServerTransport) dispatchBeginCreateOrUpdateByID(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdateByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdateByID not implemented")}
	}
	beginCreateOrUpdateByID := a.beginCreateOrUpdateByID.get(req)
	if beginCreateOrUpdateByID == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Solutions/applicationDefinitions/(?P<applicationDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmanagedapplications.ApplicationDefinition](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		applicationDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationDefinitionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateOrUpdateByID(req.Context(), resourceGroupNameParam, applicationDefinitionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdateByID = &respr
		a.beginCreateOrUpdateByID.add(req, beginCreateOrUpdateByID)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdateByID, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		a.beginCreateOrUpdateByID.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdateByID) {
		a.beginCreateOrUpdateByID.remove(req)
	}

	return resp, nil
}

func (a *ApplicationDefinitionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Solutions/applicationDefinitions/(?P<applicationDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		applicationDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationDefinitionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, applicationDefinitionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		a.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *ApplicationDefinitionsServerTransport) dispatchBeginDeleteByID(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDeleteByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeleteByID not implemented")}
	}
	beginDeleteByID := a.beginDeleteByID.get(req)
	if beginDeleteByID == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Solutions/applicationDefinitions/(?P<applicationDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		applicationDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationDefinitionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDeleteByID(req.Context(), resourceGroupNameParam, applicationDefinitionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDeleteByID = &respr
		a.beginDeleteByID.add(req, beginDeleteByID)
	}

	resp, err := server.PollerResponderNext(beginDeleteByID, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDeleteByID.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDeleteByID) {
		a.beginDeleteByID.remove(req)
	}

	return resp, nil
}

func (a *ApplicationDefinitionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Solutions/applicationDefinitions/(?P<applicationDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	applicationDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, applicationDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotFound}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotFound", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationDefinitionsServerTransport) dispatchGetByID(req *http.Request) (*http.Response, error) {
	if a.srv.GetByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByID not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Solutions/applicationDefinitions/(?P<applicationDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	applicationDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.GetByID(req.Context(), resourceGroupNameParam, applicationDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotFound}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotFound", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationDefinitionsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := a.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Solutions/applicationDefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		a.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armmanagedapplications.ApplicationDefinitionsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		a.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ApplicationDefinitionsServerTransport
var applicationDefinitionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
