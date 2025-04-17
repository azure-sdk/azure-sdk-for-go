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

// ServerFactory is a fake server for instances of the armcdn.ClientFactory type.
type ServerFactory struct {
	// AFDCustomDomainsServer contains the fakes for client AFDCustomDomainsClient
	AFDCustomDomainsServer AFDCustomDomainsServer

	// AFDEndpointsServer contains the fakes for client AFDEndpointsClient
	AFDEndpointsServer AFDEndpointsServer

	// AFDOriginGroupsServer contains the fakes for client AFDOriginGroupsClient
	AFDOriginGroupsServer AFDOriginGroupsServer

	// AFDOriginsServer contains the fakes for client AFDOriginsClient
	AFDOriginsServer AFDOriginsServer

	// AFDProfilesServer contains the fakes for client AFDProfilesClient
	AFDProfilesServer AFDProfilesServer

	// CustomDomainsServer contains the fakes for client CustomDomainsClient
	CustomDomainsServer CustomDomainsServer

	// EdgeNodesServer contains the fakes for client EdgeNodesClient
	EdgeNodesServer EdgeNodesServer

	// EndpointsServer contains the fakes for client EndpointsClient
	EndpointsServer EndpointsServer

	// LogAnalyticsServer contains the fakes for client LogAnalyticsClient
	LogAnalyticsServer LogAnalyticsServer

	// ManagedRuleSetsServer contains the fakes for client ManagedRuleSetsClient
	ManagedRuleSetsServer ManagedRuleSetsServer

	// ManagementServer contains the fakes for client ManagementClient
	ManagementServer ManagementServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// OriginGroupsServer contains the fakes for client OriginGroupsClient
	OriginGroupsServer OriginGroupsServer

	// OriginsServer contains the fakes for client OriginsClient
	OriginsServer OriginsServer

	// PoliciesServer contains the fakes for client PoliciesClient
	PoliciesServer PoliciesServer

	// ProfilesServer contains the fakes for client ProfilesClient
	ProfilesServer ProfilesServer

	// ResourceUsageServer contains the fakes for client ResourceUsageClient
	ResourceUsageServer ResourceUsageServer

	// RoutesServer contains the fakes for client RoutesClient
	RoutesServer RoutesServer

	// RuleSetsServer contains the fakes for client RuleSetsClient
	RuleSetsServer RuleSetsServer

	// RulesServer contains the fakes for client RulesClient
	RulesServer RulesServer

	// SecretsServer contains the fakes for client SecretsClient
	SecretsServer SecretsServer

	// SecurityPoliciesServer contains the fakes for client SecurityPoliciesClient
	SecurityPoliciesServer SecurityPoliciesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armcdn.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armcdn.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                      *ServerFactory
	trMu                     sync.Mutex
	trAFDCustomDomainsServer *AFDCustomDomainsServerTransport
	trAFDEndpointsServer     *AFDEndpointsServerTransport
	trAFDOriginGroupsServer  *AFDOriginGroupsServerTransport
	trAFDOriginsServer       *AFDOriginsServerTransport
	trAFDProfilesServer      *AFDProfilesServerTransport
	trCustomDomainsServer    *CustomDomainsServerTransport
	trEdgeNodesServer        *EdgeNodesServerTransport
	trEndpointsServer        *EndpointsServerTransport
	trLogAnalyticsServer     *LogAnalyticsServerTransport
	trManagedRuleSetsServer  *ManagedRuleSetsServerTransport
	trManagementServer       *ManagementServerTransport
	trOperationsServer       *OperationsServerTransport
	trOriginGroupsServer     *OriginGroupsServerTransport
	trOriginsServer          *OriginsServerTransport
	trPoliciesServer         *PoliciesServerTransport
	trProfilesServer         *ProfilesServerTransport
	trResourceUsageServer    *ResourceUsageServerTransport
	trRoutesServer           *RoutesServerTransport
	trRuleSetsServer         *RuleSetsServerTransport
	trRulesServer            *RulesServerTransport
	trSecretsServer          *SecretsServerTransport
	trSecurityPoliciesServer *SecurityPoliciesServerTransport
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
	case "AFDCustomDomainsClient":
		initServer(s, &s.trAFDCustomDomainsServer, func() *AFDCustomDomainsServerTransport {
			return NewAFDCustomDomainsServerTransport(&s.srv.AFDCustomDomainsServer)
		})
		resp, err = s.trAFDCustomDomainsServer.Do(req)
	case "AFDEndpointsClient":
		initServer(s, &s.trAFDEndpointsServer, func() *AFDEndpointsServerTransport { return NewAFDEndpointsServerTransport(&s.srv.AFDEndpointsServer) })
		resp, err = s.trAFDEndpointsServer.Do(req)
	case "AFDOriginGroupsClient":
		initServer(s, &s.trAFDOriginGroupsServer, func() *AFDOriginGroupsServerTransport {
			return NewAFDOriginGroupsServerTransport(&s.srv.AFDOriginGroupsServer)
		})
		resp, err = s.trAFDOriginGroupsServer.Do(req)
	case "AFDOriginsClient":
		initServer(s, &s.trAFDOriginsServer, func() *AFDOriginsServerTransport { return NewAFDOriginsServerTransport(&s.srv.AFDOriginsServer) })
		resp, err = s.trAFDOriginsServer.Do(req)
	case "AFDProfilesClient":
		initServer(s, &s.trAFDProfilesServer, func() *AFDProfilesServerTransport { return NewAFDProfilesServerTransport(&s.srv.AFDProfilesServer) })
		resp, err = s.trAFDProfilesServer.Do(req)
	case "CustomDomainsClient":
		initServer(s, &s.trCustomDomainsServer, func() *CustomDomainsServerTransport {
			return NewCustomDomainsServerTransport(&s.srv.CustomDomainsServer)
		})
		resp, err = s.trCustomDomainsServer.Do(req)
	case "EdgeNodesClient":
		initServer(s, &s.trEdgeNodesServer, func() *EdgeNodesServerTransport { return NewEdgeNodesServerTransport(&s.srv.EdgeNodesServer) })
		resp, err = s.trEdgeNodesServer.Do(req)
	case "EndpointsClient":
		initServer(s, &s.trEndpointsServer, func() *EndpointsServerTransport { return NewEndpointsServerTransport(&s.srv.EndpointsServer) })
		resp, err = s.trEndpointsServer.Do(req)
	case "LogAnalyticsClient":
		initServer(s, &s.trLogAnalyticsServer, func() *LogAnalyticsServerTransport { return NewLogAnalyticsServerTransport(&s.srv.LogAnalyticsServer) })
		resp, err = s.trLogAnalyticsServer.Do(req)
	case "ManagedRuleSetsClient":
		initServer(s, &s.trManagedRuleSetsServer, func() *ManagedRuleSetsServerTransport {
			return NewManagedRuleSetsServerTransport(&s.srv.ManagedRuleSetsServer)
		})
		resp, err = s.trManagedRuleSetsServer.Do(req)
	case "ManagementClient":
		initServer(s, &s.trManagementServer, func() *ManagementServerTransport { return NewManagementServerTransport(&s.srv.ManagementServer) })
		resp, err = s.trManagementServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "OriginGroupsClient":
		initServer(s, &s.trOriginGroupsServer, func() *OriginGroupsServerTransport { return NewOriginGroupsServerTransport(&s.srv.OriginGroupsServer) })
		resp, err = s.trOriginGroupsServer.Do(req)
	case "OriginsClient":
		initServer(s, &s.trOriginsServer, func() *OriginsServerTransport { return NewOriginsServerTransport(&s.srv.OriginsServer) })
		resp, err = s.trOriginsServer.Do(req)
	case "PoliciesClient":
		initServer(s, &s.trPoliciesServer, func() *PoliciesServerTransport { return NewPoliciesServerTransport(&s.srv.PoliciesServer) })
		resp, err = s.trPoliciesServer.Do(req)
	case "ProfilesClient":
		initServer(s, &s.trProfilesServer, func() *ProfilesServerTransport { return NewProfilesServerTransport(&s.srv.ProfilesServer) })
		resp, err = s.trProfilesServer.Do(req)
	case "ResourceUsageClient":
		initServer(s, &s.trResourceUsageServer, func() *ResourceUsageServerTransport {
			return NewResourceUsageServerTransport(&s.srv.ResourceUsageServer)
		})
		resp, err = s.trResourceUsageServer.Do(req)
	case "RoutesClient":
		initServer(s, &s.trRoutesServer, func() *RoutesServerTransport { return NewRoutesServerTransport(&s.srv.RoutesServer) })
		resp, err = s.trRoutesServer.Do(req)
	case "RuleSetsClient":
		initServer(s, &s.trRuleSetsServer, func() *RuleSetsServerTransport { return NewRuleSetsServerTransport(&s.srv.RuleSetsServer) })
		resp, err = s.trRuleSetsServer.Do(req)
	case "RulesClient":
		initServer(s, &s.trRulesServer, func() *RulesServerTransport { return NewRulesServerTransport(&s.srv.RulesServer) })
		resp, err = s.trRulesServer.Do(req)
	case "SecretsClient":
		initServer(s, &s.trSecretsServer, func() *SecretsServerTransport { return NewSecretsServerTransport(&s.srv.SecretsServer) })
		resp, err = s.trSecretsServer.Do(req)
	case "SecurityPoliciesClient":
		initServer(s, &s.trSecurityPoliciesServer, func() *SecurityPoliciesServerTransport {
			return NewSecurityPoliciesServerTransport(&s.srv.SecurityPoliciesServer)
		})
		resp, err = s.trSecurityPoliciesServer.Do(req)
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
