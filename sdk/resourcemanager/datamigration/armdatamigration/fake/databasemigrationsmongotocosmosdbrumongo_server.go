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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// DatabaseMigrationsMongoToCosmosDbRUMongoServer is a fake server for instances of the armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClient type.
type DatabaseMigrationsMongoToCosmosDbRUMongoServer struct {
	// BeginCreate is the fake for method DatabaseMigrationsMongoToCosmosDbRUMongoClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, targetResourceName string, migrationName string, parameters armdatamigration.DatabaseMigrationCosmosDbMongo, options *armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientBeginCreateOptions) (resp azfake.PollerResponder[armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method DatabaseMigrationsMongoToCosmosDbRUMongoClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, targetResourceName string, migrationName string, options *armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientBeginDeleteOptions) (resp azfake.PollerResponder[armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DatabaseMigrationsMongoToCosmosDbRUMongoClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, targetResourceName string, migrationName string, options *armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientGetOptions) (resp azfake.Responder[armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientGetResponse], errResp azfake.ErrorResponder)

	// NewGetForScopePager is the fake for method DatabaseMigrationsMongoToCosmosDbRUMongoClient.NewGetForScopePager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetForScopePager func(resourceGroupName string, targetResourceName string, options *armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientGetForScopeOptions) (resp azfake.PagerResponder[armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientGetForScopeResponse])
}

// NewDatabaseMigrationsMongoToCosmosDbRUMongoServerTransport creates a new instance of DatabaseMigrationsMongoToCosmosDbRUMongoServerTransport with the provided implementation.
// The returned DatabaseMigrationsMongoToCosmosDbRUMongoServerTransport instance is connected to an instance of armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDatabaseMigrationsMongoToCosmosDbRUMongoServerTransport(srv *DatabaseMigrationsMongoToCosmosDbRUMongoServer) *DatabaseMigrationsMongoToCosmosDbRUMongoServerTransport {
	return &DatabaseMigrationsMongoToCosmosDbRUMongoServerTransport{
		srv:                 srv,
		beginCreate:         newTracker[azfake.PollerResponder[armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientCreateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientDeleteResponse]](),
		newGetForScopePager: newTracker[azfake.PagerResponder[armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientGetForScopeResponse]](),
	}
}

// DatabaseMigrationsMongoToCosmosDbRUMongoServerTransport connects instances of armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClient to instances of DatabaseMigrationsMongoToCosmosDbRUMongoServer.
// Don't use this type directly, use NewDatabaseMigrationsMongoToCosmosDbRUMongoServerTransport instead.
type DatabaseMigrationsMongoToCosmosDbRUMongoServerTransport struct {
	srv                 *DatabaseMigrationsMongoToCosmosDbRUMongoServer
	beginCreate         *tracker[azfake.PollerResponder[armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientCreateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientDeleteResponse]]
	newGetForScopePager *tracker[azfake.PagerResponder[armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientGetForScopeResponse]]
}

// Do implements the policy.Transporter interface for DatabaseMigrationsMongoToCosmosDbRUMongoServerTransport.
func (d *DatabaseMigrationsMongoToCosmosDbRUMongoServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DatabaseMigrationsMongoToCosmosDbRUMongoServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if databaseMigrationsMongoToCosmosDbRuMongoServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = databaseMigrationsMongoToCosmosDbRuMongoServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DatabaseMigrationsMongoToCosmosDbRUMongoClient.BeginCreate":
				res.resp, res.err = d.dispatchBeginCreate(req)
			case "DatabaseMigrationsMongoToCosmosDbRUMongoClient.BeginDelete":
				res.resp, res.err = d.dispatchBeginDelete(req)
			case "DatabaseMigrationsMongoToCosmosDbRUMongoClient.Get":
				res.resp, res.err = d.dispatchGet(req)
			case "DatabaseMigrationsMongoToCosmosDbRUMongoClient.NewGetForScopePager":
				res.resp, res.err = d.dispatchNewGetForScopePager(req)
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

func (d *DatabaseMigrationsMongoToCosmosDbRUMongoServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := d.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<targetResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/databaseMigrations/(?P<migrationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatamigration.DatabaseMigrationCosmosDbMongo](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		targetResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetResourceName")])
		if err != nil {
			return nil, err
		}
		migrationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("migrationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreate(req.Context(), resourceGroupNameParam, targetResourceNameParam, migrationNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		d.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		d.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		d.beginCreate.remove(req)
	}

	return resp, nil
}

func (d *DatabaseMigrationsMongoToCosmosDbRUMongoServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := d.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<targetResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/databaseMigrations/(?P<migrationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		targetResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetResourceName")])
		if err != nil {
			return nil, err
		}
		migrationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("migrationName")])
		if err != nil {
			return nil, err
		}
		forceUnescaped, err := url.QueryUnescape(qp.Get("force"))
		if err != nil {
			return nil, err
		}
		forceParam, err := parseOptional(forceUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		var options *armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientBeginDeleteOptions
		if forceParam != nil {
			options = &armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientBeginDeleteOptions{
				Force: forceParam,
			}
		}
		respr, errRespr := d.srv.BeginDelete(req.Context(), resourceGroupNameParam, targetResourceNameParam, migrationNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		d.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		d.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		d.beginDelete.remove(req)
	}

	return resp, nil
}

func (d *DatabaseMigrationsMongoToCosmosDbRUMongoServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<targetResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/databaseMigrations/(?P<migrationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	targetResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetResourceName")])
	if err != nil {
		return nil, err
	}
	migrationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("migrationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, targetResourceNameParam, migrationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DatabaseMigrationCosmosDbMongo, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatabaseMigrationsMongoToCosmosDbRUMongoServerTransport) dispatchNewGetForScopePager(req *http.Request) (*http.Response, error) {
	if d.srv.NewGetForScopePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetForScopePager not implemented")}
	}
	newGetForScopePager := d.newGetForScopePager.get(req)
	if newGetForScopePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<targetResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/databaseMigrations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		targetResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetResourceName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewGetForScopePager(resourceGroupNameParam, targetResourceNameParam, nil)
		newGetForScopePager = &resp
		d.newGetForScopePager.add(req, newGetForScopePager)
		server.PagerResponderInjectNextLinks(newGetForScopePager, req, func(page *armdatamigration.DatabaseMigrationsMongoToCosmosDbRUMongoClientGetForScopeResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newGetForScopePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newGetForScopePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGetForScopePager) {
		d.newGetForScopePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to DatabaseMigrationsMongoToCosmosDbRUMongoServerTransport
var databaseMigrationsMongoToCosmosDbRuMongoServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
