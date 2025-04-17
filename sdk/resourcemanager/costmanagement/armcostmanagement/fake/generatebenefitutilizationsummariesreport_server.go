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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
	"net/http"
	"net/url"
	"regexp"
)

// GenerateBenefitUtilizationSummariesReportServer is a fake server for instances of the armcostmanagement.GenerateBenefitUtilizationSummariesReportClient type.
type GenerateBenefitUtilizationSummariesReportServer struct {
	// BeginGenerateByBillingAccount is the fake for method GenerateBenefitUtilizationSummariesReportClient.BeginGenerateByBillingAccount
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGenerateByBillingAccount func(ctx context.Context, billingAccountID string, benefitUtilizationSummariesRequest armcostmanagement.BenefitUtilizationSummariesRequest, options *armcostmanagement.GenerateBenefitUtilizationSummariesReportClientBeginGenerateByBillingAccountOptions) (resp azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateByBillingAccountResponse], errResp azfake.ErrorResponder)

	// BeginGenerateByBillingProfile is the fake for method GenerateBenefitUtilizationSummariesReportClient.BeginGenerateByBillingProfile
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGenerateByBillingProfile func(ctx context.Context, billingAccountID string, billingProfileID string, benefitUtilizationSummariesRequest armcostmanagement.BenefitUtilizationSummariesRequest, options *armcostmanagement.GenerateBenefitUtilizationSummariesReportClientBeginGenerateByBillingProfileOptions) (resp azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateByBillingProfileResponse], errResp azfake.ErrorResponder)

	// BeginGenerateByReservationID is the fake for method GenerateBenefitUtilizationSummariesReportClient.BeginGenerateByReservationID
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGenerateByReservationID func(ctx context.Context, reservationOrderID string, reservationID string, benefitUtilizationSummariesRequest armcostmanagement.BenefitUtilizationSummariesRequest, options *armcostmanagement.GenerateBenefitUtilizationSummariesReportClientBeginGenerateByReservationIDOptions) (resp azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateByReservationIDResponse], errResp azfake.ErrorResponder)

	// BeginGenerateByReservationOrderID is the fake for method GenerateBenefitUtilizationSummariesReportClient.BeginGenerateByReservationOrderID
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGenerateByReservationOrderID func(ctx context.Context, reservationOrderID string, benefitUtilizationSummariesRequest armcostmanagement.BenefitUtilizationSummariesRequest, options *armcostmanagement.GenerateBenefitUtilizationSummariesReportClientBeginGenerateByReservationOrderIDOptions) (resp azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateByReservationOrderIDResponse], errResp azfake.ErrorResponder)

	// BeginGenerateBySavingsPlanID is the fake for method GenerateBenefitUtilizationSummariesReportClient.BeginGenerateBySavingsPlanID
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGenerateBySavingsPlanID func(ctx context.Context, savingsPlanOrderID string, savingsPlanID string, benefitUtilizationSummariesRequest armcostmanagement.BenefitUtilizationSummariesRequest, options *armcostmanagement.GenerateBenefitUtilizationSummariesReportClientBeginGenerateBySavingsPlanIDOptions) (resp azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateBySavingsPlanIDResponse], errResp azfake.ErrorResponder)

	// BeginGenerateBySavingsPlanOrderID is the fake for method GenerateBenefitUtilizationSummariesReportClient.BeginGenerateBySavingsPlanOrderID
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGenerateBySavingsPlanOrderID func(ctx context.Context, savingsPlanOrderID string, benefitUtilizationSummariesRequest armcostmanagement.BenefitUtilizationSummariesRequest, options *armcostmanagement.GenerateBenefitUtilizationSummariesReportClientBeginGenerateBySavingsPlanOrderIDOptions) (resp azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateBySavingsPlanOrderIDResponse], errResp azfake.ErrorResponder)
}

// NewGenerateBenefitUtilizationSummariesReportServerTransport creates a new instance of GenerateBenefitUtilizationSummariesReportServerTransport with the provided implementation.
// The returned GenerateBenefitUtilizationSummariesReportServerTransport instance is connected to an instance of armcostmanagement.GenerateBenefitUtilizationSummariesReportClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGenerateBenefitUtilizationSummariesReportServerTransport(srv *GenerateBenefitUtilizationSummariesReportServer) *GenerateBenefitUtilizationSummariesReportServerTransport {
	return &GenerateBenefitUtilizationSummariesReportServerTransport{
		srv:                               srv,
		beginGenerateByBillingAccount:     newTracker[azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateByBillingAccountResponse]](),
		beginGenerateByBillingProfile:     newTracker[azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateByBillingProfileResponse]](),
		beginGenerateByReservationID:      newTracker[azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateByReservationIDResponse]](),
		beginGenerateByReservationOrderID: newTracker[azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateByReservationOrderIDResponse]](),
		beginGenerateBySavingsPlanID:      newTracker[azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateBySavingsPlanIDResponse]](),
		beginGenerateBySavingsPlanOrderID: newTracker[azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateBySavingsPlanOrderIDResponse]](),
	}
}

// GenerateBenefitUtilizationSummariesReportServerTransport connects instances of armcostmanagement.GenerateBenefitUtilizationSummariesReportClient to instances of GenerateBenefitUtilizationSummariesReportServer.
// Don't use this type directly, use NewGenerateBenefitUtilizationSummariesReportServerTransport instead.
type GenerateBenefitUtilizationSummariesReportServerTransport struct {
	srv                               *GenerateBenefitUtilizationSummariesReportServer
	beginGenerateByBillingAccount     *tracker[azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateByBillingAccountResponse]]
	beginGenerateByBillingProfile     *tracker[azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateByBillingProfileResponse]]
	beginGenerateByReservationID      *tracker[azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateByReservationIDResponse]]
	beginGenerateByReservationOrderID *tracker[azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateByReservationOrderIDResponse]]
	beginGenerateBySavingsPlanID      *tracker[azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateBySavingsPlanIDResponse]]
	beginGenerateBySavingsPlanOrderID *tracker[azfake.PollerResponder[armcostmanagement.GenerateBenefitUtilizationSummariesReportClientGenerateBySavingsPlanOrderIDResponse]]
}

// Do implements the policy.Transporter interface for GenerateBenefitUtilizationSummariesReportServerTransport.
func (g *GenerateBenefitUtilizationSummariesReportServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return g.dispatchToMethodFake(req, method)
}

func (g *GenerateBenefitUtilizationSummariesReportServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if generateBenefitUtilizationSummariesReportServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = generateBenefitUtilizationSummariesReportServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "GenerateBenefitUtilizationSummariesReportClient.BeginGenerateByBillingAccount":
				res.resp, res.err = g.dispatchBeginGenerateByBillingAccount(req)
			case "GenerateBenefitUtilizationSummariesReportClient.BeginGenerateByBillingProfile":
				res.resp, res.err = g.dispatchBeginGenerateByBillingProfile(req)
			case "GenerateBenefitUtilizationSummariesReportClient.BeginGenerateByReservationID":
				res.resp, res.err = g.dispatchBeginGenerateByReservationID(req)
			case "GenerateBenefitUtilizationSummariesReportClient.BeginGenerateByReservationOrderID":
				res.resp, res.err = g.dispatchBeginGenerateByReservationOrderID(req)
			case "GenerateBenefitUtilizationSummariesReportClient.BeginGenerateBySavingsPlanID":
				res.resp, res.err = g.dispatchBeginGenerateBySavingsPlanID(req)
			case "GenerateBenefitUtilizationSummariesReportClient.BeginGenerateBySavingsPlanOrderID":
				res.resp, res.err = g.dispatchBeginGenerateBySavingsPlanOrderID(req)
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

func (g *GenerateBenefitUtilizationSummariesReportServerTransport) dispatchBeginGenerateByBillingAccount(req *http.Request) (*http.Response, error) {
	if g.srv.BeginGenerateByBillingAccount == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGenerateByBillingAccount not implemented")}
	}
	beginGenerateByBillingAccount := g.beginGenerateByBillingAccount.get(req)
	if beginGenerateByBillingAccount == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/generateBenefitUtilizationSummariesReport`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcostmanagement.BenefitUtilizationSummariesRequest](req)
		if err != nil {
			return nil, err
		}
		billingAccountIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginGenerateByBillingAccount(req.Context(), billingAccountIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGenerateByBillingAccount = &respr
		g.beginGenerateByBillingAccount.add(req, beginGenerateByBillingAccount)
	}

	resp, err := server.PollerResponderNext(beginGenerateByBillingAccount, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		g.beginGenerateByBillingAccount.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGenerateByBillingAccount) {
		g.beginGenerateByBillingAccount.remove(req)
	}

	return resp, nil
}

func (g *GenerateBenefitUtilizationSummariesReportServerTransport) dispatchBeginGenerateByBillingProfile(req *http.Request) (*http.Response, error) {
	if g.srv.BeginGenerateByBillingProfile == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGenerateByBillingProfile not implemented")}
	}
	beginGenerateByBillingProfile := g.beginGenerateByBillingProfile.get(req)
	if beginGenerateByBillingProfile == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/generateBenefitUtilizationSummariesReport`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcostmanagement.BenefitUtilizationSummariesRequest](req)
		if err != nil {
			return nil, err
		}
		billingAccountIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountId")])
		if err != nil {
			return nil, err
		}
		billingProfileIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginGenerateByBillingProfile(req.Context(), billingAccountIDParam, billingProfileIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGenerateByBillingProfile = &respr
		g.beginGenerateByBillingProfile.add(req, beginGenerateByBillingProfile)
	}

	resp, err := server.PollerResponderNext(beginGenerateByBillingProfile, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		g.beginGenerateByBillingProfile.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGenerateByBillingProfile) {
		g.beginGenerateByBillingProfile.remove(req)
	}

	return resp, nil
}

func (g *GenerateBenefitUtilizationSummariesReportServerTransport) dispatchBeginGenerateByReservationID(req *http.Request) (*http.Response, error) {
	if g.srv.BeginGenerateByReservationID == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGenerateByReservationID not implemented")}
	}
	beginGenerateByReservationID := g.beginGenerateByReservationID.get(req)
	if beginGenerateByReservationID == nil {
		const regexStr = `/providers/Microsoft\.Capacity/reservationorders/(?P<reservationOrderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reservations/(?P<reservationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/generateBenefitUtilizationSummariesReport`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcostmanagement.BenefitUtilizationSummariesRequest](req)
		if err != nil {
			return nil, err
		}
		reservationOrderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reservationOrderId")])
		if err != nil {
			return nil, err
		}
		reservationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reservationId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginGenerateByReservationID(req.Context(), reservationOrderIDParam, reservationIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGenerateByReservationID = &respr
		g.beginGenerateByReservationID.add(req, beginGenerateByReservationID)
	}

	resp, err := server.PollerResponderNext(beginGenerateByReservationID, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		g.beginGenerateByReservationID.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGenerateByReservationID) {
		g.beginGenerateByReservationID.remove(req)
	}

	return resp, nil
}

func (g *GenerateBenefitUtilizationSummariesReportServerTransport) dispatchBeginGenerateByReservationOrderID(req *http.Request) (*http.Response, error) {
	if g.srv.BeginGenerateByReservationOrderID == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGenerateByReservationOrderID not implemented")}
	}
	beginGenerateByReservationOrderID := g.beginGenerateByReservationOrderID.get(req)
	if beginGenerateByReservationOrderID == nil {
		const regexStr = `/providers/Microsoft\.Capacity/reservationorders/(?P<reservationOrderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/generateBenefitUtilizationSummariesReport`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcostmanagement.BenefitUtilizationSummariesRequest](req)
		if err != nil {
			return nil, err
		}
		reservationOrderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reservationOrderId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginGenerateByReservationOrderID(req.Context(), reservationOrderIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGenerateByReservationOrderID = &respr
		g.beginGenerateByReservationOrderID.add(req, beginGenerateByReservationOrderID)
	}

	resp, err := server.PollerResponderNext(beginGenerateByReservationOrderID, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		g.beginGenerateByReservationOrderID.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGenerateByReservationOrderID) {
		g.beginGenerateByReservationOrderID.remove(req)
	}

	return resp, nil
}

func (g *GenerateBenefitUtilizationSummariesReportServerTransport) dispatchBeginGenerateBySavingsPlanID(req *http.Request) (*http.Response, error) {
	if g.srv.BeginGenerateBySavingsPlanID == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGenerateBySavingsPlanID not implemented")}
	}
	beginGenerateBySavingsPlanID := g.beginGenerateBySavingsPlanID.get(req)
	if beginGenerateBySavingsPlanID == nil {
		const regexStr = `/providers/Microsoft\.BillingBenefits/savingsPlanOrders/(?P<savingsPlanOrderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/savingsPlans/(?P<savingsPlanId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/generateBenefitUtilizationSummariesReport`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcostmanagement.BenefitUtilizationSummariesRequest](req)
		if err != nil {
			return nil, err
		}
		savingsPlanOrderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("savingsPlanOrderId")])
		if err != nil {
			return nil, err
		}
		savingsPlanIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("savingsPlanId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginGenerateBySavingsPlanID(req.Context(), savingsPlanOrderIDParam, savingsPlanIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGenerateBySavingsPlanID = &respr
		g.beginGenerateBySavingsPlanID.add(req, beginGenerateBySavingsPlanID)
	}

	resp, err := server.PollerResponderNext(beginGenerateBySavingsPlanID, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		g.beginGenerateBySavingsPlanID.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGenerateBySavingsPlanID) {
		g.beginGenerateBySavingsPlanID.remove(req)
	}

	return resp, nil
}

func (g *GenerateBenefitUtilizationSummariesReportServerTransport) dispatchBeginGenerateBySavingsPlanOrderID(req *http.Request) (*http.Response, error) {
	if g.srv.BeginGenerateBySavingsPlanOrderID == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGenerateBySavingsPlanOrderID not implemented")}
	}
	beginGenerateBySavingsPlanOrderID := g.beginGenerateBySavingsPlanOrderID.get(req)
	if beginGenerateBySavingsPlanOrderID == nil {
		const regexStr = `/providers/Microsoft\.BillingBenefits/savingsPlanOrders/(?P<savingsPlanOrderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/generateBenefitUtilizationSummariesReport`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcostmanagement.BenefitUtilizationSummariesRequest](req)
		if err != nil {
			return nil, err
		}
		savingsPlanOrderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("savingsPlanOrderId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginGenerateBySavingsPlanOrderID(req.Context(), savingsPlanOrderIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGenerateBySavingsPlanOrderID = &respr
		g.beginGenerateBySavingsPlanOrderID.add(req, beginGenerateBySavingsPlanOrderID)
	}

	resp, err := server.PollerResponderNext(beginGenerateBySavingsPlanOrderID, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		g.beginGenerateBySavingsPlanOrderID.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGenerateBySavingsPlanOrderID) {
		g.beginGenerateBySavingsPlanOrderID.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to GenerateBenefitUtilizationSummariesReportServerTransport
var generateBenefitUtilizationSummariesReportServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
