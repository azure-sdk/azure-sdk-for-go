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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// SensitivityLabelsServer is a fake server for instances of the armsql.SensitivityLabelsClient type.
type SensitivityLabelsServer struct {
	// CreateOrUpdate is the fake for method SensitivityLabelsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, parameters armsql.SensitivityLabel, options *armsql.SensitivityLabelsClientCreateOrUpdateOptions) (resp azfake.Responder[armsql.SensitivityLabelsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method SensitivityLabelsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK
	Delete func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, options *armsql.SensitivityLabelsClientDeleteOptions) (resp azfake.Responder[armsql.SensitivityLabelsClientDeleteResponse], errResp azfake.ErrorResponder)

	// DisableRecommendation is the fake for method SensitivityLabelsClient.DisableRecommendation
	// HTTP status codes to indicate success: http.StatusOK
	DisableRecommendation func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, options *armsql.SensitivityLabelsClientDisableRecommendationOptions) (resp azfake.Responder[armsql.SensitivityLabelsClientDisableRecommendationResponse], errResp azfake.ErrorResponder)

	// EnableRecommendation is the fake for method SensitivityLabelsClient.EnableRecommendation
	// HTTP status codes to indicate success: http.StatusOK
	EnableRecommendation func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, options *armsql.SensitivityLabelsClientEnableRecommendationOptions) (resp azfake.Responder[armsql.SensitivityLabelsClientEnableRecommendationResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SensitivityLabelsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, sensitivityLabelSource armsql.SensitivityLabelSource, options *armsql.SensitivityLabelsClientGetOptions) (resp azfake.Responder[armsql.SensitivityLabelsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDatabasePager is the fake for method SensitivityLabelsClient.NewListByDatabasePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDatabasePager func(resourceGroupName string, serverName string, databaseName string, options *armsql.SensitivityLabelsClientListByDatabaseOptions) (resp azfake.PagerResponder[armsql.SensitivityLabelsClientListByDatabaseResponse])

	// NewListCurrentByDatabasePager is the fake for method SensitivityLabelsClient.NewListCurrentByDatabasePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListCurrentByDatabasePager func(resourceGroupName string, serverName string, databaseName string, options *armsql.SensitivityLabelsClientListCurrentByDatabaseOptions) (resp azfake.PagerResponder[armsql.SensitivityLabelsClientListCurrentByDatabaseResponse])

	// NewListRecommendedByDatabasePager is the fake for method SensitivityLabelsClient.NewListRecommendedByDatabasePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListRecommendedByDatabasePager func(resourceGroupName string, serverName string, databaseName string, options *armsql.SensitivityLabelsClientListRecommendedByDatabaseOptions) (resp azfake.PagerResponder[armsql.SensitivityLabelsClientListRecommendedByDatabaseResponse])

	// Update is the fake for method SensitivityLabelsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters armsql.SensitivityLabelUpdateList, options *armsql.SensitivityLabelsClientUpdateOptions) (resp azfake.Responder[armsql.SensitivityLabelsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewSensitivityLabelsServerTransport creates a new instance of SensitivityLabelsServerTransport with the provided implementation.
// The returned SensitivityLabelsServerTransport instance is connected to an instance of armsql.SensitivityLabelsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSensitivityLabelsServerTransport(srv *SensitivityLabelsServer) *SensitivityLabelsServerTransport {
	return &SensitivityLabelsServerTransport{
		srv:                               srv,
		newListByDatabasePager:            newTracker[azfake.PagerResponder[armsql.SensitivityLabelsClientListByDatabaseResponse]](),
		newListCurrentByDatabasePager:     newTracker[azfake.PagerResponder[armsql.SensitivityLabelsClientListCurrentByDatabaseResponse]](),
		newListRecommendedByDatabasePager: newTracker[azfake.PagerResponder[armsql.SensitivityLabelsClientListRecommendedByDatabaseResponse]](),
	}
}

// SensitivityLabelsServerTransport connects instances of armsql.SensitivityLabelsClient to instances of SensitivityLabelsServer.
// Don't use this type directly, use NewSensitivityLabelsServerTransport instead.
type SensitivityLabelsServerTransport struct {
	srv                               *SensitivityLabelsServer
	newListByDatabasePager            *tracker[azfake.PagerResponder[armsql.SensitivityLabelsClientListByDatabaseResponse]]
	newListCurrentByDatabasePager     *tracker[azfake.PagerResponder[armsql.SensitivityLabelsClientListCurrentByDatabaseResponse]]
	newListRecommendedByDatabasePager *tracker[azfake.PagerResponder[armsql.SensitivityLabelsClientListRecommendedByDatabaseResponse]]
}

// Do implements the policy.Transporter interface for SensitivityLabelsServerTransport.
func (s *SensitivityLabelsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SensitivityLabelsClient.CreateOrUpdate":
		resp, err = s.dispatchCreateOrUpdate(req)
	case "SensitivityLabelsClient.Delete":
		resp, err = s.dispatchDelete(req)
	case "SensitivityLabelsClient.DisableRecommendation":
		resp, err = s.dispatchDisableRecommendation(req)
	case "SensitivityLabelsClient.EnableRecommendation":
		resp, err = s.dispatchEnableRecommendation(req)
	case "SensitivityLabelsClient.Get":
		resp, err = s.dispatchGet(req)
	case "SensitivityLabelsClient.NewListByDatabasePager":
		resp, err = s.dispatchNewListByDatabasePager(req)
	case "SensitivityLabelsClient.NewListCurrentByDatabasePager":
		resp, err = s.dispatchNewListCurrentByDatabasePager(req)
	case "SensitivityLabelsClient.NewListRecommendedByDatabasePager":
		resp, err = s.dispatchNewListRecommendedByDatabasePager(req)
	case "SensitivityLabelsClient.Update":
		resp, err = s.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SensitivityLabelsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/columns/(?P<columnName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sensitivityLabels/(?P<sensitivityLabelSource>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsql.SensitivityLabel](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	schemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaName")])
	if err != nil {
		return nil, err
	}
	tableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
	if err != nil {
		return nil, err
	}
	columnNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("columnName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, schemaNameParam, tableNameParam, columnNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SensitivityLabel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SensitivityLabelsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/columns/(?P<columnName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sensitivityLabels/(?P<sensitivityLabelSource>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	schemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaName")])
	if err != nil {
		return nil, err
	}
	tableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
	if err != nil {
		return nil, err
	}
	columnNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("columnName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Delete(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, schemaNameParam, tableNameParam, columnNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SensitivityLabelsServerTransport) dispatchDisableRecommendation(req *http.Request) (*http.Response, error) {
	if s.srv.DisableRecommendation == nil {
		return nil, &nonRetriableError{errors.New("fake for method DisableRecommendation not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/columns/(?P<columnName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sensitivityLabels/(?P<sensitivityLabelSource>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disable`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	schemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaName")])
	if err != nil {
		return nil, err
	}
	tableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
	if err != nil {
		return nil, err
	}
	columnNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("columnName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.DisableRecommendation(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, schemaNameParam, tableNameParam, columnNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SensitivityLabelsServerTransport) dispatchEnableRecommendation(req *http.Request) (*http.Response, error) {
	if s.srv.EnableRecommendation == nil {
		return nil, &nonRetriableError{errors.New("fake for method EnableRecommendation not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/columns/(?P<columnName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sensitivityLabels/(?P<sensitivityLabelSource>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/enable`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	schemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaName")])
	if err != nil {
		return nil, err
	}
	tableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
	if err != nil {
		return nil, err
	}
	columnNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("columnName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.EnableRecommendation(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, schemaNameParam, tableNameParam, columnNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SensitivityLabelsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/columns/(?P<columnName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sensitivityLabels/(?P<sensitivityLabelSource>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 8 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	schemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaName")])
	if err != nil {
		return nil, err
	}
	tableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
	if err != nil {
		return nil, err
	}
	columnNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("columnName")])
	if err != nil {
		return nil, err
	}
	sensitivityLabelSourceParam, err := parseWithCast(matches[regex.SubexpIndex("sensitivityLabelSource")], func(v string) (armsql.SensitivityLabelSource, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsql.SensitivityLabelSource(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, schemaNameParam, tableNameParam, columnNameParam, sensitivityLabelSourceParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SensitivityLabel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SensitivityLabelsServerTransport) dispatchNewListByDatabasePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByDatabasePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDatabasePager not implemented")}
	}
	newListByDatabasePager := s.newListByDatabasePager.get(req)
	if newListByDatabasePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sensitivityLabels`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armsql.SensitivityLabelsClientListByDatabaseOptions
		if filterParam != nil {
			options = &armsql.SensitivityLabelsClientListByDatabaseOptions{
				Filter: filterParam,
			}
		}
		resp := s.srv.NewListByDatabasePager(resourceGroupNameParam, serverNameParam, databaseNameParam, options)
		newListByDatabasePager = &resp
		s.newListByDatabasePager.add(req, newListByDatabasePager)
		server.PagerResponderInjectNextLinks(newListByDatabasePager, req, func(page *armsql.SensitivityLabelsClientListByDatabaseResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByDatabasePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByDatabasePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDatabasePager) {
		s.newListByDatabasePager.remove(req)
	}
	return resp, nil
}

func (s *SensitivityLabelsServerTransport) dispatchNewListCurrentByDatabasePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListCurrentByDatabasePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListCurrentByDatabasePager not implemented")}
	}
	newListCurrentByDatabasePager := s.newListCurrentByDatabasePager.get(req)
	if newListCurrentByDatabasePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/currentSensitivityLabels`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		countUnescaped, err := url.QueryUnescape(qp.Get("$count"))
		if err != nil {
			return nil, err
		}
		countParam, err := parseOptional(countUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armsql.SensitivityLabelsClientListCurrentByDatabaseOptions
		if skipTokenParam != nil || countParam != nil || filterParam != nil {
			options = &armsql.SensitivityLabelsClientListCurrentByDatabaseOptions{
				SkipToken: skipTokenParam,
				Count:     countParam,
				Filter:    filterParam,
			}
		}
		resp := s.srv.NewListCurrentByDatabasePager(resourceGroupNameParam, serverNameParam, databaseNameParam, options)
		newListCurrentByDatabasePager = &resp
		s.newListCurrentByDatabasePager.add(req, newListCurrentByDatabasePager)
		server.PagerResponderInjectNextLinks(newListCurrentByDatabasePager, req, func(page *armsql.SensitivityLabelsClientListCurrentByDatabaseResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListCurrentByDatabasePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListCurrentByDatabasePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListCurrentByDatabasePager) {
		s.newListCurrentByDatabasePager.remove(req)
	}
	return resp, nil
}

func (s *SensitivityLabelsServerTransport) dispatchNewListRecommendedByDatabasePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListRecommendedByDatabasePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListRecommendedByDatabasePager not implemented")}
	}
	newListRecommendedByDatabasePager := s.newListRecommendedByDatabasePager.get(req)
	if newListRecommendedByDatabasePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/recommendedSensitivityLabels`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		includeDisabledRecommendationsUnescaped, err := url.QueryUnescape(qp.Get("includeDisabledRecommendations"))
		if err != nil {
			return nil, err
		}
		includeDisabledRecommendationsParam, err := parseOptional(includeDisabledRecommendationsUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armsql.SensitivityLabelsClientListRecommendedByDatabaseOptions
		if skipTokenParam != nil || includeDisabledRecommendationsParam != nil || filterParam != nil {
			options = &armsql.SensitivityLabelsClientListRecommendedByDatabaseOptions{
				SkipToken:                      skipTokenParam,
				IncludeDisabledRecommendations: includeDisabledRecommendationsParam,
				Filter:                         filterParam,
			}
		}
		resp := s.srv.NewListRecommendedByDatabasePager(resourceGroupNameParam, serverNameParam, databaseNameParam, options)
		newListRecommendedByDatabasePager = &resp
		s.newListRecommendedByDatabasePager.add(req, newListRecommendedByDatabasePager)
		server.PagerResponderInjectNextLinks(newListRecommendedByDatabasePager, req, func(page *armsql.SensitivityLabelsClientListRecommendedByDatabaseResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListRecommendedByDatabasePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListRecommendedByDatabasePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListRecommendedByDatabasePager) {
		s.newListRecommendedByDatabasePager.remove(req)
	}
	return resp, nil
}

func (s *SensitivityLabelsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/currentSensitivityLabels`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsql.SensitivityLabelUpdateList](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Update(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
