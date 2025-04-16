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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// ServiceServer is a fake server for instances of the armapimanagement.ServiceClient type.
type ServiceServer struct {
	// BeginApplyNetworkConfigurationUpdates is the fake for method ServiceClient.BeginApplyNetworkConfigurationUpdates
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginApplyNetworkConfigurationUpdates func(ctx context.Context, resourceGroupName string, serviceName string, options *armapimanagement.ServiceClientBeginApplyNetworkConfigurationUpdatesOptions) (resp azfake.PollerResponder[armapimanagement.ServiceClientApplyNetworkConfigurationUpdatesResponse], errResp azfake.ErrorResponder)

	// BeginBackup is the fake for method ServiceClient.BeginBackup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginBackup func(ctx context.Context, resourceGroupName string, serviceName string, parameters armapimanagement.ServiceBackupRestoreParameters, options *armapimanagement.ServiceClientBeginBackupOptions) (resp azfake.PollerResponder[armapimanagement.ServiceClientBackupResponse], errResp azfake.ErrorResponder)

	// CheckNameAvailability is the fake for method ServiceClient.CheckNameAvailability
	// HTTP status codes to indicate success: http.StatusOK
	CheckNameAvailability func(ctx context.Context, parameters armapimanagement.ServiceCheckNameAvailabilityParameters, options *armapimanagement.ServiceClientCheckNameAvailabilityOptions) (resp azfake.Responder[armapimanagement.ServiceClientCheckNameAvailabilityResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method ServiceClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, parameters armapimanagement.ServiceResource, options *armapimanagement.ServiceClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armapimanagement.ServiceClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ServiceClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, serviceName string, options *armapimanagement.ServiceClientBeginDeleteOptions) (resp azfake.PollerResponder[armapimanagement.ServiceClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ServiceClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, options *armapimanagement.ServiceClientGetOptions) (resp azfake.Responder[armapimanagement.ServiceClientGetResponse], errResp azfake.ErrorResponder)

	// GetDomainOwnershipIdentifier is the fake for method ServiceClient.GetDomainOwnershipIdentifier
	// HTTP status codes to indicate success: http.StatusOK
	GetDomainOwnershipIdentifier func(ctx context.Context, options *armapimanagement.ServiceClientGetDomainOwnershipIdentifierOptions) (resp azfake.Responder[armapimanagement.ServiceClientGetDomainOwnershipIdentifierResponse], errResp azfake.ErrorResponder)

	// GetSsoToken is the fake for method ServiceClient.GetSsoToken
	// HTTP status codes to indicate success: http.StatusOK
	GetSsoToken func(ctx context.Context, resourceGroupName string, serviceName string, options *armapimanagement.ServiceClientGetSsoTokenOptions) (resp azfake.Responder[armapimanagement.ServiceClientGetSsoTokenResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ServiceClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armapimanagement.ServiceClientListOptions) (resp azfake.PagerResponder[armapimanagement.ServiceClientListResponse])

	// NewListByResourceGroupPager is the fake for method ServiceClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armapimanagement.ServiceClientListByResourceGroupOptions) (resp azfake.PagerResponder[armapimanagement.ServiceClientListByResourceGroupResponse])

	// BeginMigrateToStv2 is the fake for method ServiceClient.BeginMigrateToStv2
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginMigrateToStv2 func(ctx context.Context, resourceGroupName string, serviceName string, options *armapimanagement.ServiceClientBeginMigrateToStv2Options) (resp azfake.PollerResponder[armapimanagement.ServiceClientMigrateToStv2Response], errResp azfake.ErrorResponder)

	// BeginRestore is the fake for method ServiceClient.BeginRestore
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRestore func(ctx context.Context, resourceGroupName string, serviceName string, parameters armapimanagement.ServiceBackupRestoreParameters, options *armapimanagement.ServiceClientBeginRestoreOptions) (resp azfake.PollerResponder[armapimanagement.ServiceClientRestoreResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method ServiceClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, serviceName string, parameters armapimanagement.ServiceUpdateParameters, options *armapimanagement.ServiceClientBeginUpdateOptions) (resp azfake.PollerResponder[armapimanagement.ServiceClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewServiceServerTransport creates a new instance of ServiceServerTransport with the provided implementation.
// The returned ServiceServerTransport instance is connected to an instance of armapimanagement.ServiceClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServiceServerTransport(srv *ServiceServer) *ServiceServerTransport {
	return &ServiceServerTransport{
		srv:                                   srv,
		beginApplyNetworkConfigurationUpdates: newTracker[azfake.PollerResponder[armapimanagement.ServiceClientApplyNetworkConfigurationUpdatesResponse]](),
		beginBackup:                           newTracker[azfake.PollerResponder[armapimanagement.ServiceClientBackupResponse]](),
		beginCreateOrUpdate:                   newTracker[azfake.PollerResponder[armapimanagement.ServiceClientCreateOrUpdateResponse]](),
		beginDelete:                           newTracker[azfake.PollerResponder[armapimanagement.ServiceClientDeleteResponse]](),
		newListPager:                          newTracker[azfake.PagerResponder[armapimanagement.ServiceClientListResponse]](),
		newListByResourceGroupPager:           newTracker[azfake.PagerResponder[armapimanagement.ServiceClientListByResourceGroupResponse]](),
		beginMigrateToStv2:                    newTracker[azfake.PollerResponder[armapimanagement.ServiceClientMigrateToStv2Response]](),
		beginRestore:                          newTracker[azfake.PollerResponder[armapimanagement.ServiceClientRestoreResponse]](),
		beginUpdate:                           newTracker[azfake.PollerResponder[armapimanagement.ServiceClientUpdateResponse]](),
	}
}

// ServiceServerTransport connects instances of armapimanagement.ServiceClient to instances of ServiceServer.
// Don't use this type directly, use NewServiceServerTransport instead.
type ServiceServerTransport struct {
	srv                                   *ServiceServer
	beginApplyNetworkConfigurationUpdates *tracker[azfake.PollerResponder[armapimanagement.ServiceClientApplyNetworkConfigurationUpdatesResponse]]
	beginBackup                           *tracker[azfake.PollerResponder[armapimanagement.ServiceClientBackupResponse]]
	beginCreateOrUpdate                   *tracker[azfake.PollerResponder[armapimanagement.ServiceClientCreateOrUpdateResponse]]
	beginDelete                           *tracker[azfake.PollerResponder[armapimanagement.ServiceClientDeleteResponse]]
	newListPager                          *tracker[azfake.PagerResponder[armapimanagement.ServiceClientListResponse]]
	newListByResourceGroupPager           *tracker[azfake.PagerResponder[armapimanagement.ServiceClientListByResourceGroupResponse]]
	beginMigrateToStv2                    *tracker[azfake.PollerResponder[armapimanagement.ServiceClientMigrateToStv2Response]]
	beginRestore                          *tracker[azfake.PollerResponder[armapimanagement.ServiceClientRestoreResponse]]
	beginUpdate                           *tracker[azfake.PollerResponder[armapimanagement.ServiceClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ServiceServerTransport.
func (s *ServiceServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServiceServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if serviceServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = serviceServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ServiceClient.BeginApplyNetworkConfigurationUpdates":
				res.resp, res.err = s.dispatchBeginApplyNetworkConfigurationUpdates(req)
			case "ServiceClient.BeginBackup":
				res.resp, res.err = s.dispatchBeginBackup(req)
			case "ServiceClient.CheckNameAvailability":
				res.resp, res.err = s.dispatchCheckNameAvailability(req)
			case "ServiceClient.BeginCreateOrUpdate":
				res.resp, res.err = s.dispatchBeginCreateOrUpdate(req)
			case "ServiceClient.BeginDelete":
				res.resp, res.err = s.dispatchBeginDelete(req)
			case "ServiceClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "ServiceClient.GetDomainOwnershipIdentifier":
				res.resp, res.err = s.dispatchGetDomainOwnershipIdentifier(req)
			case "ServiceClient.GetSsoToken":
				res.resp, res.err = s.dispatchGetSsoToken(req)
			case "ServiceClient.NewListPager":
				res.resp, res.err = s.dispatchNewListPager(req)
			case "ServiceClient.NewListByResourceGroupPager":
				res.resp, res.err = s.dispatchNewListByResourceGroupPager(req)
			case "ServiceClient.BeginMigrateToStv2":
				res.resp, res.err = s.dispatchBeginMigrateToStv2(req)
			case "ServiceClient.BeginRestore":
				res.resp, res.err = s.dispatchBeginRestore(req)
			case "ServiceClient.BeginUpdate":
				res.resp, res.err = s.dispatchBeginUpdate(req)
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

func (s *ServiceServerTransport) dispatchBeginApplyNetworkConfigurationUpdates(req *http.Request) (*http.Response, error) {
	if s.srv.BeginApplyNetworkConfigurationUpdates == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginApplyNetworkConfigurationUpdates not implemented")}
	}
	beginApplyNetworkConfigurationUpdates := s.beginApplyNetworkConfigurationUpdates.get(req)
	if beginApplyNetworkConfigurationUpdates == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applynetworkconfigurationupdates`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armapimanagement.ServiceApplyNetworkConfigurationParameters](req)
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
		var options *armapimanagement.ServiceClientBeginApplyNetworkConfigurationUpdatesOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armapimanagement.ServiceClientBeginApplyNetworkConfigurationUpdatesOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := s.srv.BeginApplyNetworkConfigurationUpdates(req.Context(), resourceGroupNameParam, serviceNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginApplyNetworkConfigurationUpdates = &respr
		s.beginApplyNetworkConfigurationUpdates.add(req, beginApplyNetworkConfigurationUpdates)
	}

	resp, err := server.PollerResponderNext(beginApplyNetworkConfigurationUpdates, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginApplyNetworkConfigurationUpdates.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginApplyNetworkConfigurationUpdates) {
		s.beginApplyNetworkConfigurationUpdates.remove(req)
	}

	return resp, nil
}

func (s *ServiceServerTransport) dispatchBeginBackup(req *http.Request) (*http.Response, error) {
	if s.srv.BeginBackup == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginBackup not implemented")}
	}
	beginBackup := s.beginBackup.get(req)
	if beginBackup == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backup`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armapimanagement.ServiceBackupRestoreParameters](req)
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
		respr, errRespr := s.srv.BeginBackup(req.Context(), resourceGroupNameParam, serviceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginBackup = &respr
		s.beginBackup.add(req, beginBackup)
	}

	resp, err := server.PollerResponderNext(beginBackup, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginBackup.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginBackup) {
		s.beginBackup.remove(req)
	}

	return resp, nil
}

func (s *ServiceServerTransport) dispatchCheckNameAvailability(req *http.Request) (*http.Response, error) {
	if s.srv.CheckNameAvailability == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckNameAvailability not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/checkNameAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armapimanagement.ServiceCheckNameAvailabilityParameters](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CheckNameAvailability(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServiceNameAvailabilityResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServiceServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := s.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armapimanagement.ServiceResource](req)
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
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		s.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		s.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (s *ServiceServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
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
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, serviceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *ServiceServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
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
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServiceResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServiceServerTransport) dispatchGetDomainOwnershipIdentifier(req *http.Request) (*http.Response, error) {
	if s.srv.GetDomainOwnershipIdentifier == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDomainOwnershipIdentifier not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/getDomainOwnershipIdentifier`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := s.srv.GetDomainOwnershipIdentifier(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServiceGetDomainOwnershipIdentifierResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServiceServerTransport) dispatchGetSsoToken(req *http.Request) (*http.Response, error) {
	if s.srv.GetSsoToken == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSsoToken not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getssotoken`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
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
	respr, errRespr := s.srv.GetSsoToken(req.Context(), resourceGroupNameParam, serviceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServiceGetSsoTokenResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServiceServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := s.srv.NewListPager(nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armapimanagement.ServiceClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}

func (s *ServiceServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := s.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		s.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armapimanagement.ServiceClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		s.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (s *ServiceServerTransport) dispatchBeginMigrateToStv2(req *http.Request) (*http.Response, error) {
	if s.srv.BeginMigrateToStv2 == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginMigrateToStv2 not implemented")}
	}
	beginMigrateToStv2 := s.beginMigrateToStv2.get(req)
	if beginMigrateToStv2 == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/migrateToStv2`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armapimanagement.MigrateToStv2Contract](req)
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
		var options *armapimanagement.ServiceClientBeginMigrateToStv2Options
		if !reflect.ValueOf(body).IsZero() {
			options = &armapimanagement.ServiceClientBeginMigrateToStv2Options{
				Parameters: &body,
			}
		}
		respr, errRespr := s.srv.BeginMigrateToStv2(req.Context(), resourceGroupNameParam, serviceNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginMigrateToStv2 = &respr
		s.beginMigrateToStv2.add(req, beginMigrateToStv2)
	}

	resp, err := server.PollerResponderNext(beginMigrateToStv2, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginMigrateToStv2.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginMigrateToStv2) {
		s.beginMigrateToStv2.remove(req)
	}

	return resp, nil
}

func (s *ServiceServerTransport) dispatchBeginRestore(req *http.Request) (*http.Response, error) {
	if s.srv.BeginRestore == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRestore not implemented")}
	}
	beginRestore := s.beginRestore.get(req)
	if beginRestore == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restore`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armapimanagement.ServiceBackupRestoreParameters](req)
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
		respr, errRespr := s.srv.BeginRestore(req.Context(), resourceGroupNameParam, serviceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRestore = &respr
		s.beginRestore.add(req, beginRestore)
	}

	resp, err := server.PollerResponderNext(beginRestore, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginRestore.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRestore) {
		s.beginRestore.remove(req)
	}

	return resp, nil
}

func (s *ServiceServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := s.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armapimanagement.ServiceUpdateParameters](req)
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
		respr, errRespr := s.srv.BeginUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		s.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		s.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ServiceServerTransport
var serviceServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
