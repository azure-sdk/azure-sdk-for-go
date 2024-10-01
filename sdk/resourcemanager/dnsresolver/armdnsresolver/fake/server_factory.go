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

// ServerFactory is a fake server for instances of the armdnsresolver.ClientFactory type.
type ServerFactory struct {
	DNSForwardingRulesetsServer     DNSForwardingRulesetsServer
	DNSResolversServer              DNSResolversServer
	DNSSecurityRulesServer          DNSSecurityRulesServer
	DomainListsServer               DomainListsServer
	ForwardingRulesServer           ForwardingRulesServer
	InboundEndpointsServer          InboundEndpointsServer
	OutboundEndpointsServer         OutboundEndpointsServer
	PoliciesServer                  PoliciesServer
	PolicyVirtualNetworkLinksServer PolicyVirtualNetworkLinksServer
	VirtualNetworkLinksServer       VirtualNetworkLinksServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armdnsresolver.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armdnsresolver.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                               *ServerFactory
	trMu                              sync.Mutex
	trDNSForwardingRulesetsServer     *DNSForwardingRulesetsServerTransport
	trDNSResolversServer              *DNSResolversServerTransport
	trDNSSecurityRulesServer          *DNSSecurityRulesServerTransport
	trDomainListsServer               *DomainListsServerTransport
	trForwardingRulesServer           *ForwardingRulesServerTransport
	trInboundEndpointsServer          *InboundEndpointsServerTransport
	trOutboundEndpointsServer         *OutboundEndpointsServerTransport
	trPoliciesServer                  *PoliciesServerTransport
	trPolicyVirtualNetworkLinksServer *PolicyVirtualNetworkLinksServerTransport
	trVirtualNetworkLinksServer       *VirtualNetworkLinksServerTransport
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
	case "DNSForwardingRulesetsClient":
		initServer(s, &s.trDNSForwardingRulesetsServer, func() *DNSForwardingRulesetsServerTransport {
			return NewDNSForwardingRulesetsServerTransport(&s.srv.DNSForwardingRulesetsServer)
		})
		resp, err = s.trDNSForwardingRulesetsServer.Do(req)
	case "DNSResolversClient":
		initServer(s, &s.trDNSResolversServer, func() *DNSResolversServerTransport { return NewDNSResolversServerTransport(&s.srv.DNSResolversServer) })
		resp, err = s.trDNSResolversServer.Do(req)
	case "DNSSecurityRulesClient":
		initServer(s, &s.trDNSSecurityRulesServer, func() *DNSSecurityRulesServerTransport {
			return NewDNSSecurityRulesServerTransport(&s.srv.DNSSecurityRulesServer)
		})
		resp, err = s.trDNSSecurityRulesServer.Do(req)
	case "DomainListsClient":
		initServer(s, &s.trDomainListsServer, func() *DomainListsServerTransport { return NewDomainListsServerTransport(&s.srv.DomainListsServer) })
		resp, err = s.trDomainListsServer.Do(req)
	case "ForwardingRulesClient":
		initServer(s, &s.trForwardingRulesServer, func() *ForwardingRulesServerTransport {
			return NewForwardingRulesServerTransport(&s.srv.ForwardingRulesServer)
		})
		resp, err = s.trForwardingRulesServer.Do(req)
	case "InboundEndpointsClient":
		initServer(s, &s.trInboundEndpointsServer, func() *InboundEndpointsServerTransport {
			return NewInboundEndpointsServerTransport(&s.srv.InboundEndpointsServer)
		})
		resp, err = s.trInboundEndpointsServer.Do(req)
	case "OutboundEndpointsClient":
		initServer(s, &s.trOutboundEndpointsServer, func() *OutboundEndpointsServerTransport {
			return NewOutboundEndpointsServerTransport(&s.srv.OutboundEndpointsServer)
		})
		resp, err = s.trOutboundEndpointsServer.Do(req)
	case "PoliciesClient":
		initServer(s, &s.trPoliciesServer, func() *PoliciesServerTransport { return NewPoliciesServerTransport(&s.srv.PoliciesServer) })
		resp, err = s.trPoliciesServer.Do(req)
	case "PolicyVirtualNetworkLinksClient":
		initServer(s, &s.trPolicyVirtualNetworkLinksServer, func() *PolicyVirtualNetworkLinksServerTransport {
			return NewPolicyVirtualNetworkLinksServerTransport(&s.srv.PolicyVirtualNetworkLinksServer)
		})
		resp, err = s.trPolicyVirtualNetworkLinksServer.Do(req)
	case "VirtualNetworkLinksClient":
		initServer(s, &s.trVirtualNetworkLinksServer, func() *VirtualNetworkLinksServerTransport {
			return NewVirtualNetworkLinksServerTransport(&s.srv.VirtualNetworkLinksServer)
		})
		resp, err = s.trVirtualNetworkLinksServer.Do(req)
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
