//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armquota.ClientFactory type.
type ServerFactory struct {
	Server                                              Server
	GroupQuotaEnforcementServer                         GroupQuotaEnforcementServer
	GroupQuotaLimitsServer                              GroupQuotaLimitsServer
	GroupQuotaLimitsRequestsServer                      GroupQuotaLimitsRequestsServer
	GroupQuotaSubscriptionQuotaAllocationServer         GroupQuotaSubscriptionQuotaAllocationServer
	GroupQuotaSubscriptionQuotaAllocationRequestsServer GroupQuotaSubscriptionQuotaAllocationRequestsServer
	GroupQuotaSubscriptionsServer                       GroupQuotaSubscriptionsServer
	GroupQuotasServer                                   GroupQuotasServer
	OperationServer                                     OperationServer
	RequestStatusServer                                 RequestStatusServer
	SubscriptionRequestsServer                          SubscriptionRequestsServer
	UsagesServer                                        UsagesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armquota.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armquota.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                                   *ServerFactory
	trMu                                                  sync.Mutex
	trServer                                              *ServerTransport
	trGroupQuotaEnforcementServer                         *GroupQuotaEnforcementServerTransport
	trGroupQuotaLimitsServer                              *GroupQuotaLimitsServerTransport
	trGroupQuotaLimitsRequestsServer                      *GroupQuotaLimitsRequestsServerTransport
	trGroupQuotaSubscriptionQuotaAllocationServer         *GroupQuotaSubscriptionQuotaAllocationServerTransport
	trGroupQuotaSubscriptionQuotaAllocationRequestsServer *GroupQuotaSubscriptionQuotaAllocationRequestsServerTransport
	trGroupQuotaSubscriptionsServer                       *GroupQuotaSubscriptionsServerTransport
	trGroupQuotasServer                                   *GroupQuotasServerTransport
	trOperationServer                                     *OperationServerTransport
	trRequestStatusServer                                 *RequestStatusServerTransport
	trSubscriptionRequestsServer                          *SubscriptionRequestsServerTransport
	trUsagesServer                                        *UsagesServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "Client":
		initServer(s, &s.trServer, func() *ServerTransport { return NewServerTransport(&s.srv.Server) })
		resp, err = s.trServer.Do(req)
	case "GroupQuotaEnforcementClient":
		initServer(s, &s.trGroupQuotaEnforcementServer, func() *GroupQuotaEnforcementServerTransport {
			return NewGroupQuotaEnforcementServerTransport(&s.srv.GroupQuotaEnforcementServer)
		})
		resp, err = s.trGroupQuotaEnforcementServer.Do(req)
	case "GroupQuotaLimitsClient":
		initServer(s, &s.trGroupQuotaLimitsServer, func() *GroupQuotaLimitsServerTransport {
			return NewGroupQuotaLimitsServerTransport(&s.srv.GroupQuotaLimitsServer)
		})
		resp, err = s.trGroupQuotaLimitsServer.Do(req)
	case "GroupQuotaLimitsRequestsClient":
		initServer(s, &s.trGroupQuotaLimitsRequestsServer, func() *GroupQuotaLimitsRequestsServerTransport {
			return NewGroupQuotaLimitsRequestsServerTransport(&s.srv.GroupQuotaLimitsRequestsServer)
		})
		resp, err = s.trGroupQuotaLimitsRequestsServer.Do(req)
	case "GroupQuotaSubscriptionQuotaAllocationClient":
		initServer(s, &s.trGroupQuotaSubscriptionQuotaAllocationServer, func() *GroupQuotaSubscriptionQuotaAllocationServerTransport {
			return NewGroupQuotaSubscriptionQuotaAllocationServerTransport(&s.srv.GroupQuotaSubscriptionQuotaAllocationServer)
		})
		resp, err = s.trGroupQuotaSubscriptionQuotaAllocationServer.Do(req)
	case "GroupQuotaSubscriptionQuotaAllocationRequestsClient":
		initServer(s, &s.trGroupQuotaSubscriptionQuotaAllocationRequestsServer, func() *GroupQuotaSubscriptionQuotaAllocationRequestsServerTransport {
			return NewGroupQuotaSubscriptionQuotaAllocationRequestsServerTransport(&s.srv.GroupQuotaSubscriptionQuotaAllocationRequestsServer)
		})
		resp, err = s.trGroupQuotaSubscriptionQuotaAllocationRequestsServer.Do(req)
	case "GroupQuotaSubscriptionsClient":
		initServer(s, &s.trGroupQuotaSubscriptionsServer, func() *GroupQuotaSubscriptionsServerTransport {
			return NewGroupQuotaSubscriptionsServerTransport(&s.srv.GroupQuotaSubscriptionsServer)
		})
		resp, err = s.trGroupQuotaSubscriptionsServer.Do(req)
	case "GroupQuotasClient":
		initServer(s, &s.trGroupQuotasServer, func() *GroupQuotasServerTransport { return NewGroupQuotasServerTransport(&s.srv.GroupQuotasServer) })
		resp, err = s.trGroupQuotasServer.Do(req)
	case "OperationClient":
		initServer(s, &s.trOperationServer, func() *OperationServerTransport { return NewOperationServerTransport(&s.srv.OperationServer) })
		resp, err = s.trOperationServer.Do(req)
	case "RequestStatusClient":
		initServer(s, &s.trRequestStatusServer, func() *RequestStatusServerTransport {
			return NewRequestStatusServerTransport(&s.srv.RequestStatusServer)
		})
		resp, err = s.trRequestStatusServer.Do(req)
	case "SubscriptionRequestsClient":
		initServer(s, &s.trSubscriptionRequestsServer, func() *SubscriptionRequestsServerTransport {
			return NewSubscriptionRequestsServerTransport(&s.srv.SubscriptionRequestsServer)
		})
		resp, err = s.trSubscriptionRequestsServer.Do(req)
	case "UsagesClient":
		initServer(s, &s.trUsagesServer, func() *UsagesServerTransport { return NewUsagesServerTransport(&s.srv.UsagesServer) })
		resp, err = s.trUsagesServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
