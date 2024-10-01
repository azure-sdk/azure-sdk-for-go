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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
	"net/http"
	"net/url"
	"regexp"
)

// KubernetesClusterFeaturesServer is a fake server for instances of the armnetworkcloud.KubernetesClusterFeaturesClient type.
type KubernetesClusterFeaturesServer struct {
	// BeginCreateOrUpdate is the fake for method KubernetesClusterFeaturesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, kubernetesClusterName string, featureName string, kubernetesClusterFeatureParameters armnetworkcloud.KubernetesClusterFeature, options *armnetworkcloud.KubernetesClusterFeaturesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetworkcloud.KubernetesClusterFeaturesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method KubernetesClusterFeaturesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, kubernetesClusterName string, featureName string, options *armnetworkcloud.KubernetesClusterFeaturesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetworkcloud.KubernetesClusterFeaturesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method KubernetesClusterFeaturesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, kubernetesClusterName string, featureName string, options *armnetworkcloud.KubernetesClusterFeaturesClientGetOptions) (resp azfake.Responder[armnetworkcloud.KubernetesClusterFeaturesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByKubernetesClusterPager is the fake for method KubernetesClusterFeaturesClient.NewListByKubernetesClusterPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByKubernetesClusterPager func(resourceGroupName string, kubernetesClusterName string, options *armnetworkcloud.KubernetesClusterFeaturesClientListByKubernetesClusterOptions) (resp azfake.PagerResponder[armnetworkcloud.KubernetesClusterFeaturesClientListByKubernetesClusterResponse])

	// BeginUpdate is the fake for method KubernetesClusterFeaturesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, kubernetesClusterName string, featureName string, kubernetesClusterFeatureUpdateParameters armnetworkcloud.KubernetesClusterFeaturePatchParameters, options *armnetworkcloud.KubernetesClusterFeaturesClientBeginUpdateOptions) (resp azfake.PollerResponder[armnetworkcloud.KubernetesClusterFeaturesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewKubernetesClusterFeaturesServerTransport creates a new instance of KubernetesClusterFeaturesServerTransport with the provided implementation.
// The returned KubernetesClusterFeaturesServerTransport instance is connected to an instance of armnetworkcloud.KubernetesClusterFeaturesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewKubernetesClusterFeaturesServerTransport(srv *KubernetesClusterFeaturesServer) *KubernetesClusterFeaturesServerTransport {
	return &KubernetesClusterFeaturesServerTransport{
		srv:                             srv,
		beginCreateOrUpdate:             newTracker[azfake.PollerResponder[armnetworkcloud.KubernetesClusterFeaturesClientCreateOrUpdateResponse]](),
		beginDelete:                     newTracker[azfake.PollerResponder[armnetworkcloud.KubernetesClusterFeaturesClientDeleteResponse]](),
		newListByKubernetesClusterPager: newTracker[azfake.PagerResponder[armnetworkcloud.KubernetesClusterFeaturesClientListByKubernetesClusterResponse]](),
		beginUpdate:                     newTracker[azfake.PollerResponder[armnetworkcloud.KubernetesClusterFeaturesClientUpdateResponse]](),
	}
}

// KubernetesClusterFeaturesServerTransport connects instances of armnetworkcloud.KubernetesClusterFeaturesClient to instances of KubernetesClusterFeaturesServer.
// Don't use this type directly, use NewKubernetesClusterFeaturesServerTransport instead.
type KubernetesClusterFeaturesServerTransport struct {
	srv                             *KubernetesClusterFeaturesServer
	beginCreateOrUpdate             *tracker[azfake.PollerResponder[armnetworkcloud.KubernetesClusterFeaturesClientCreateOrUpdateResponse]]
	beginDelete                     *tracker[azfake.PollerResponder[armnetworkcloud.KubernetesClusterFeaturesClientDeleteResponse]]
	newListByKubernetesClusterPager *tracker[azfake.PagerResponder[armnetworkcloud.KubernetesClusterFeaturesClientListByKubernetesClusterResponse]]
	beginUpdate                     *tracker[azfake.PollerResponder[armnetworkcloud.KubernetesClusterFeaturesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for KubernetesClusterFeaturesServerTransport.
func (k *KubernetesClusterFeaturesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "KubernetesClusterFeaturesClient.BeginCreateOrUpdate":
		resp, err = k.dispatchBeginCreateOrUpdate(req)
	case "KubernetesClusterFeaturesClient.BeginDelete":
		resp, err = k.dispatchBeginDelete(req)
	case "KubernetesClusterFeaturesClient.Get":
		resp, err = k.dispatchGet(req)
	case "KubernetesClusterFeaturesClient.NewListByKubernetesClusterPager":
		resp, err = k.dispatchNewListByKubernetesClusterPager(req)
	case "KubernetesClusterFeaturesClient.BeginUpdate":
		resp, err = k.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (k *KubernetesClusterFeaturesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if k.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := k.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/kubernetesClusters/(?P<kubernetesClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/features/(?P<featureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.KubernetesClusterFeature](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		kubernetesClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("kubernetesClusterName")])
		if err != nil {
			return nil, err
		}
		featureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("featureName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := k.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, kubernetesClusterNameParam, featureNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		k.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		k.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		k.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (k *KubernetesClusterFeaturesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if k.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := k.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/kubernetesClusters/(?P<kubernetesClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/features/(?P<featureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		kubernetesClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("kubernetesClusterName")])
		if err != nil {
			return nil, err
		}
		featureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("featureName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := k.srv.BeginDelete(req.Context(), resourceGroupNameParam, kubernetesClusterNameParam, featureNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		k.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		k.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		k.beginDelete.remove(req)
	}

	return resp, nil
}

func (k *KubernetesClusterFeaturesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if k.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/kubernetesClusters/(?P<kubernetesClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/features/(?P<featureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	kubernetesClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("kubernetesClusterName")])
	if err != nil {
		return nil, err
	}
	featureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("featureName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := k.srv.Get(req.Context(), resourceGroupNameParam, kubernetesClusterNameParam, featureNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).KubernetesClusterFeature, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (k *KubernetesClusterFeaturesServerTransport) dispatchNewListByKubernetesClusterPager(req *http.Request) (*http.Response, error) {
	if k.srv.NewListByKubernetesClusterPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByKubernetesClusterPager not implemented")}
	}
	newListByKubernetesClusterPager := k.newListByKubernetesClusterPager.get(req)
	if newListByKubernetesClusterPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/kubernetesClusters/(?P<kubernetesClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/features`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		kubernetesClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("kubernetesClusterName")])
		if err != nil {
			return nil, err
		}
		resp := k.srv.NewListByKubernetesClusterPager(resourceGroupNameParam, kubernetesClusterNameParam, nil)
		newListByKubernetesClusterPager = &resp
		k.newListByKubernetesClusterPager.add(req, newListByKubernetesClusterPager)
		server.PagerResponderInjectNextLinks(newListByKubernetesClusterPager, req, func(page *armnetworkcloud.KubernetesClusterFeaturesClientListByKubernetesClusterResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByKubernetesClusterPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		k.newListByKubernetesClusterPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByKubernetesClusterPager) {
		k.newListByKubernetesClusterPager.remove(req)
	}
	return resp, nil
}

func (k *KubernetesClusterFeaturesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if k.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := k.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/kubernetesClusters/(?P<kubernetesClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/features/(?P<featureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.KubernetesClusterFeaturePatchParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		kubernetesClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("kubernetesClusterName")])
		if err != nil {
			return nil, err
		}
		featureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("featureName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := k.srv.BeginUpdate(req.Context(), resourceGroupNameParam, kubernetesClusterNameParam, featureNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		k.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		k.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		k.beginUpdate.remove(req)
	}

	return resp, nil
}
