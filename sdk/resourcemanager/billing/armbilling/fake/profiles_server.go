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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ProfilesServer is a fake server for instances of the armbilling.ProfilesClient type.
type ProfilesServer struct {
	// BeginCreateOrUpdate is the fake for method ProfilesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, billingAccountName string, billingProfileName string, parameters armbilling.Profile, options *armbilling.ProfilesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armbilling.ProfilesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ProfilesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, billingAccountName string, billingProfileName string, options *armbilling.ProfilesClientBeginDeleteOptions) (resp azfake.PollerResponder[armbilling.ProfilesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ProfilesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, billingAccountName string, billingProfileName string, options *armbilling.ProfilesClientGetOptions) (resp azfake.Responder[armbilling.ProfilesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByBillingAccountPager is the fake for method ProfilesClient.NewListByBillingAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByBillingAccountPager func(billingAccountName string, options *armbilling.ProfilesClientListByBillingAccountOptions) (resp azfake.PagerResponder[armbilling.ProfilesClientListByBillingAccountResponse])

	// ValidateDeleteEligibility is the fake for method ProfilesClient.ValidateDeleteEligibility
	// HTTP status codes to indicate success: http.StatusOK
	ValidateDeleteEligibility func(ctx context.Context, billingAccountName string, billingProfileName string, options *armbilling.ProfilesClientValidateDeleteEligibilityOptions) (resp azfake.Responder[armbilling.ProfilesClientValidateDeleteEligibilityResponse], errResp azfake.ErrorResponder)
}

// NewProfilesServerTransport creates a new instance of ProfilesServerTransport with the provided implementation.
// The returned ProfilesServerTransport instance is connected to an instance of armbilling.ProfilesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProfilesServerTransport(srv *ProfilesServer) *ProfilesServerTransport {
	return &ProfilesServerTransport{
		srv:                          srv,
		beginCreateOrUpdate:          newTracker[azfake.PollerResponder[armbilling.ProfilesClientCreateOrUpdateResponse]](),
		beginDelete:                  newTracker[azfake.PollerResponder[armbilling.ProfilesClientDeleteResponse]](),
		newListByBillingAccountPager: newTracker[azfake.PagerResponder[armbilling.ProfilesClientListByBillingAccountResponse]](),
	}
}

// ProfilesServerTransport connects instances of armbilling.ProfilesClient to instances of ProfilesServer.
// Don't use this type directly, use NewProfilesServerTransport instead.
type ProfilesServerTransport struct {
	srv                          *ProfilesServer
	beginCreateOrUpdate          *tracker[azfake.PollerResponder[armbilling.ProfilesClientCreateOrUpdateResponse]]
	beginDelete                  *tracker[azfake.PollerResponder[armbilling.ProfilesClientDeleteResponse]]
	newListByBillingAccountPager *tracker[azfake.PagerResponder[armbilling.ProfilesClientListByBillingAccountResponse]]
}

// Do implements the policy.Transporter interface for ProfilesServerTransport.
func (p *ProfilesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *ProfilesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if profilesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = profilesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ProfilesClient.BeginCreateOrUpdate":
				res.resp, res.err = p.dispatchBeginCreateOrUpdate(req)
			case "ProfilesClient.BeginDelete":
				res.resp, res.err = p.dispatchBeginDelete(req)
			case "ProfilesClient.Get":
				res.resp, res.err = p.dispatchGet(req)
			case "ProfilesClient.NewListByBillingAccountPager":
				res.resp, res.err = p.dispatchNewListByBillingAccountPager(req)
			case "ProfilesClient.ValidateDeleteEligibility":
				res.resp, res.err = p.dispatchValidateDeleteEligibility(req)
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

func (p *ProfilesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armbilling.Profile](req)
		if err != nil {
			return nil, err
		}
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		billingProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), billingAccountNameParam, billingProfileNameParam, body, nil)
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

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		p.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		p.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (p *ProfilesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		billingProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), billingAccountNameParam, billingProfileNameParam, nil)
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

func (p *ProfilesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
	if err != nil {
		return nil, err
	}
	billingProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), billingAccountNameParam, billingProfileNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Profile, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProfilesServerTransport) dispatchNewListByBillingAccountPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByBillingAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByBillingAccountPager not implemented")}
	}
	newListByBillingAccountPager := p.newListByBillingAccountPager.get(req)
	if newListByBillingAccountPager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		includeDeletedUnescaped, err := url.QueryUnescape(qp.Get("includeDeleted"))
		if err != nil {
			return nil, err
		}
		includeDeletedParam, err := parseOptional(includeDeletedUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderByUnescaped, err := url.QueryUnescape(qp.Get("orderBy"))
		if err != nil {
			return nil, err
		}
		orderByParam := getOptional(orderByUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		countUnescaped, err := url.QueryUnescape(qp.Get("count"))
		if err != nil {
			return nil, err
		}
		countParam, err := parseOptional(countUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		searchUnescaped, err := url.QueryUnescape(qp.Get("search"))
		if err != nil {
			return nil, err
		}
		searchParam := getOptional(searchUnescaped)
		var options *armbilling.ProfilesClientListByBillingAccountOptions
		if includeDeletedParam != nil || filterParam != nil || orderByParam != nil || topParam != nil || skipParam != nil || countParam != nil || searchParam != nil {
			options = &armbilling.ProfilesClientListByBillingAccountOptions{
				IncludeDeleted: includeDeletedParam,
				Filter:         filterParam,
				OrderBy:        orderByParam,
				Top:            topParam,
				Skip:           skipParam,
				Count:          countParam,
				Search:         searchParam,
			}
		}
		resp := p.srv.NewListByBillingAccountPager(billingAccountNameParam, options)
		newListByBillingAccountPager = &resp
		p.newListByBillingAccountPager.add(req, newListByBillingAccountPager)
		server.PagerResponderInjectNextLinks(newListByBillingAccountPager, req, func(page *armbilling.ProfilesClientListByBillingAccountResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByBillingAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByBillingAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByBillingAccountPager) {
		p.newListByBillingAccountPager.remove(req)
	}
	return resp, nil
}

func (p *ProfilesServerTransport) dispatchValidateDeleteEligibility(req *http.Request) (*http.Response, error) {
	if p.srv.ValidateDeleteEligibility == nil {
		return nil, &nonRetriableError{errors.New("fake for method ValidateDeleteEligibility not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/validateDeleteEligibility`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
	if err != nil {
		return nil, err
	}
	billingProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.ValidateDeleteEligibility(req.Context(), billingAccountNameParam, billingProfileNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeleteBillingProfileEligibilityResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ProfilesServerTransport
var profilesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
