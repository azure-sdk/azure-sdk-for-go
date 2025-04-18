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

// ServerFactory is a fake server for instances of the armhybridnetwork.ClientFactory type.
type ServerFactory struct {
	// ArtifactManifestsServer contains the fakes for client ArtifactManifestsClient
	ArtifactManifestsServer ArtifactManifestsServer

	// ArtifactStoresServer contains the fakes for client ArtifactStoresClient
	ArtifactStoresServer ArtifactStoresServer

	// ComponentsServer contains the fakes for client ComponentsClient
	ComponentsServer ComponentsServer

	// ConfigurationGroupSchemasServer contains the fakes for client ConfigurationGroupSchemasClient
	ConfigurationGroupSchemasServer ConfigurationGroupSchemasServer

	// ConfigurationGroupValuesServer contains the fakes for client ConfigurationGroupValuesClient
	ConfigurationGroupValuesServer ConfigurationGroupValuesServer

	// NetworkFunctionDefinitionGroupsServer contains the fakes for client NetworkFunctionDefinitionGroupsClient
	NetworkFunctionDefinitionGroupsServer NetworkFunctionDefinitionGroupsServer

	// NetworkFunctionDefinitionVersionsServer contains the fakes for client NetworkFunctionDefinitionVersionsClient
	NetworkFunctionDefinitionVersionsServer NetworkFunctionDefinitionVersionsServer

	// NetworkFunctionsServer contains the fakes for client NetworkFunctionsClient
	NetworkFunctionsServer NetworkFunctionsServer

	// NetworkServiceDesignGroupsServer contains the fakes for client NetworkServiceDesignGroupsClient
	NetworkServiceDesignGroupsServer NetworkServiceDesignGroupsServer

	// NetworkServiceDesignVersionsServer contains the fakes for client NetworkServiceDesignVersionsClient
	NetworkServiceDesignVersionsServer NetworkServiceDesignVersionsServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// ProxyArtifactServer contains the fakes for client ProxyArtifactClient
	ProxyArtifactServer ProxyArtifactServer

	// PublishersServer contains the fakes for client PublishersClient
	PublishersServer PublishersServer

	// SiteNetworkServicesServer contains the fakes for client SiteNetworkServicesClient
	SiteNetworkServicesServer SiteNetworkServicesServer

	// SitesServer contains the fakes for client SitesClient
	SitesServer SitesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armhybridnetwork.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armhybridnetwork.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                       *ServerFactory
	trMu                                      sync.Mutex
	trArtifactManifestsServer                 *ArtifactManifestsServerTransport
	trArtifactStoresServer                    *ArtifactStoresServerTransport
	trComponentsServer                        *ComponentsServerTransport
	trConfigurationGroupSchemasServer         *ConfigurationGroupSchemasServerTransport
	trConfigurationGroupValuesServer          *ConfigurationGroupValuesServerTransport
	trNetworkFunctionDefinitionGroupsServer   *NetworkFunctionDefinitionGroupsServerTransport
	trNetworkFunctionDefinitionVersionsServer *NetworkFunctionDefinitionVersionsServerTransport
	trNetworkFunctionsServer                  *NetworkFunctionsServerTransport
	trNetworkServiceDesignGroupsServer        *NetworkServiceDesignGroupsServerTransport
	trNetworkServiceDesignVersionsServer      *NetworkServiceDesignVersionsServerTransport
	trOperationsServer                        *OperationsServerTransport
	trProxyArtifactServer                     *ProxyArtifactServerTransport
	trPublishersServer                        *PublishersServerTransport
	trSiteNetworkServicesServer               *SiteNetworkServicesServerTransport
	trSitesServer                             *SitesServerTransport
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
	case "ArtifactManifestsClient":
		initServer(s, &s.trArtifactManifestsServer, func() *ArtifactManifestsServerTransport {
			return NewArtifactManifestsServerTransport(&s.srv.ArtifactManifestsServer)
		})
		resp, err = s.trArtifactManifestsServer.Do(req)
	case "ArtifactStoresClient":
		initServer(s, &s.trArtifactStoresServer, func() *ArtifactStoresServerTransport {
			return NewArtifactStoresServerTransport(&s.srv.ArtifactStoresServer)
		})
		resp, err = s.trArtifactStoresServer.Do(req)
	case "ComponentsClient":
		initServer(s, &s.trComponentsServer, func() *ComponentsServerTransport { return NewComponentsServerTransport(&s.srv.ComponentsServer) })
		resp, err = s.trComponentsServer.Do(req)
	case "ConfigurationGroupSchemasClient":
		initServer(s, &s.trConfigurationGroupSchemasServer, func() *ConfigurationGroupSchemasServerTransport {
			return NewConfigurationGroupSchemasServerTransport(&s.srv.ConfigurationGroupSchemasServer)
		})
		resp, err = s.trConfigurationGroupSchemasServer.Do(req)
	case "ConfigurationGroupValuesClient":
		initServer(s, &s.trConfigurationGroupValuesServer, func() *ConfigurationGroupValuesServerTransport {
			return NewConfigurationGroupValuesServerTransport(&s.srv.ConfigurationGroupValuesServer)
		})
		resp, err = s.trConfigurationGroupValuesServer.Do(req)
	case "NetworkFunctionDefinitionGroupsClient":
		initServer(s, &s.trNetworkFunctionDefinitionGroupsServer, func() *NetworkFunctionDefinitionGroupsServerTransport {
			return NewNetworkFunctionDefinitionGroupsServerTransport(&s.srv.NetworkFunctionDefinitionGroupsServer)
		})
		resp, err = s.trNetworkFunctionDefinitionGroupsServer.Do(req)
	case "NetworkFunctionDefinitionVersionsClient":
		initServer(s, &s.trNetworkFunctionDefinitionVersionsServer, func() *NetworkFunctionDefinitionVersionsServerTransport {
			return NewNetworkFunctionDefinitionVersionsServerTransport(&s.srv.NetworkFunctionDefinitionVersionsServer)
		})
		resp, err = s.trNetworkFunctionDefinitionVersionsServer.Do(req)
	case "NetworkFunctionsClient":
		initServer(s, &s.trNetworkFunctionsServer, func() *NetworkFunctionsServerTransport {
			return NewNetworkFunctionsServerTransport(&s.srv.NetworkFunctionsServer)
		})
		resp, err = s.trNetworkFunctionsServer.Do(req)
	case "NetworkServiceDesignGroupsClient":
		initServer(s, &s.trNetworkServiceDesignGroupsServer, func() *NetworkServiceDesignGroupsServerTransport {
			return NewNetworkServiceDesignGroupsServerTransport(&s.srv.NetworkServiceDesignGroupsServer)
		})
		resp, err = s.trNetworkServiceDesignGroupsServer.Do(req)
	case "NetworkServiceDesignVersionsClient":
		initServer(s, &s.trNetworkServiceDesignVersionsServer, func() *NetworkServiceDesignVersionsServerTransport {
			return NewNetworkServiceDesignVersionsServerTransport(&s.srv.NetworkServiceDesignVersionsServer)
		})
		resp, err = s.trNetworkServiceDesignVersionsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "ProxyArtifactClient":
		initServer(s, &s.trProxyArtifactServer, func() *ProxyArtifactServerTransport {
			return NewProxyArtifactServerTransport(&s.srv.ProxyArtifactServer)
		})
		resp, err = s.trProxyArtifactServer.Do(req)
	case "PublishersClient":
		initServer(s, &s.trPublishersServer, func() *PublishersServerTransport { return NewPublishersServerTransport(&s.srv.PublishersServer) })
		resp, err = s.trPublishersServer.Do(req)
	case "SiteNetworkServicesClient":
		initServer(s, &s.trSiteNetworkServicesServer, func() *SiteNetworkServicesServerTransport {
			return NewSiteNetworkServicesServerTransport(&s.srv.SiteNetworkServicesServer)
		})
		resp, err = s.trSiteNetworkServicesServer.Do(req)
	case "SitesClient":
		initServer(s, &s.trSitesServer, func() *SitesServerTransport { return NewSitesServerTransport(&s.srv.SitesServer) })
		resp, err = s.trSitesServer.Do(req)
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
