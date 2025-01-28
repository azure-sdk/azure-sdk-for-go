//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
	"net/http"
	"net/url"
	"regexp"
)

// ImportCollectorsOperationsServer is a fake server for instances of the armmigrationassessment.ImportCollectorsOperationsClient type.
type ImportCollectorsOperationsServer struct {
	// BeginCreate is the fake for method ImportCollectorsOperationsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, projectName string, importCollectorName string, resource armmigrationassessment.ImportCollector, options *armmigrationassessment.ImportCollectorsOperationsClientBeginCreateOptions) (resp azfake.PollerResponder[armmigrationassessment.ImportCollectorsOperationsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ImportCollectorsOperationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, projectName string, importCollectorName string, options *armmigrationassessment.ImportCollectorsOperationsClientDeleteOptions) (resp azfake.Responder[armmigrationassessment.ImportCollectorsOperationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ImportCollectorsOperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, projectName string, importCollectorName string, options *armmigrationassessment.ImportCollectorsOperationsClientGetOptions) (resp azfake.Responder[armmigrationassessment.ImportCollectorsOperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByAssessmentProjectPager is the fake for method ImportCollectorsOperationsClient.NewListByAssessmentProjectPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAssessmentProjectPager func(resourceGroupName string, projectName string, options *armmigrationassessment.ImportCollectorsOperationsClientListByAssessmentProjectOptions) (resp azfake.PagerResponder[armmigrationassessment.ImportCollectorsOperationsClientListByAssessmentProjectResponse])
}

// NewImportCollectorsOperationsServerTransport creates a new instance of ImportCollectorsOperationsServerTransport with the provided implementation.
// The returned ImportCollectorsOperationsServerTransport instance is connected to an instance of armmigrationassessment.ImportCollectorsOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewImportCollectorsOperationsServerTransport(srv *ImportCollectorsOperationsServer) *ImportCollectorsOperationsServerTransport {
	return &ImportCollectorsOperationsServerTransport{
		srv:                             srv,
		beginCreate:                     newTracker[azfake.PollerResponder[armmigrationassessment.ImportCollectorsOperationsClientCreateResponse]](),
		newListByAssessmentProjectPager: newTracker[azfake.PagerResponder[armmigrationassessment.ImportCollectorsOperationsClientListByAssessmentProjectResponse]](),
	}
}

// ImportCollectorsOperationsServerTransport connects instances of armmigrationassessment.ImportCollectorsOperationsClient to instances of ImportCollectorsOperationsServer.
// Don't use this type directly, use NewImportCollectorsOperationsServerTransport instead.
type ImportCollectorsOperationsServerTransport struct {
	srv                             *ImportCollectorsOperationsServer
	beginCreate                     *tracker[azfake.PollerResponder[armmigrationassessment.ImportCollectorsOperationsClientCreateResponse]]
	newListByAssessmentProjectPager *tracker[azfake.PagerResponder[armmigrationassessment.ImportCollectorsOperationsClientListByAssessmentProjectResponse]]
}

// Do implements the policy.Transporter interface for ImportCollectorsOperationsServerTransport.
func (i *ImportCollectorsOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ImportCollectorsOperationsClient.BeginCreate":
		resp, err = i.dispatchBeginCreate(req)
	case "ImportCollectorsOperationsClient.Delete":
		resp, err = i.dispatchDelete(req)
	case "ImportCollectorsOperationsClient.Get":
		resp, err = i.dispatchGet(req)
	case "ImportCollectorsOperationsClient.NewListByAssessmentProjectPager":
		resp, err = i.dispatchNewListByAssessmentProjectPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *ImportCollectorsOperationsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := i.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/importcollectors/(?P<importCollectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrationassessment.ImportCollector](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
		if err != nil {
			return nil, err
		}
		importCollectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("importCollectorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreate(req.Context(), resourceGroupNameParam, projectNameParam, importCollectorNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		i.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		i.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		i.beginCreate.remove(req)
	}

	return resp, nil
}

func (i *ImportCollectorsOperationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if i.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/importcollectors/(?P<importCollectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	importCollectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("importCollectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Delete(req.Context(), resourceGroupNameParam, projectNameParam, importCollectorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ImportCollectorsOperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/importcollectors/(?P<importCollectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	importCollectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("importCollectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, projectNameParam, importCollectorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImportCollector, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ImportCollectorsOperationsServerTransport) dispatchNewListByAssessmentProjectPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByAssessmentProjectPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAssessmentProjectPager not implemented")}
	}
	newListByAssessmentProjectPager := i.newListByAssessmentProjectPager.get(req)
	if newListByAssessmentProjectPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/importcollectors`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListByAssessmentProjectPager(resourceGroupNameParam, projectNameParam, nil)
		newListByAssessmentProjectPager = &resp
		i.newListByAssessmentProjectPager.add(req, newListByAssessmentProjectPager)
		server.PagerResponderInjectNextLinks(newListByAssessmentProjectPager, req, func(page *armmigrationassessment.ImportCollectorsOperationsClientListByAssessmentProjectResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByAssessmentProjectPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByAssessmentProjectPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAssessmentProjectPager) {
		i.newListByAssessmentProjectPager.remove(req)
	}
	return resp, nil
}
