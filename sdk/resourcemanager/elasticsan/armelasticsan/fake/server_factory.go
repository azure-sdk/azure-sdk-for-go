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

// ServerFactory is a fake server for instances of the armelasticsan.ClientFactory type.
type ServerFactory struct {
	// ElasticSansServer contains the fakes for client ElasticSansClient
	ElasticSansServer ElasticSansServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// PrivateEndpointConnectionsServer contains the fakes for client PrivateEndpointConnectionsClient
	PrivateEndpointConnectionsServer PrivateEndpointConnectionsServer

	// PrivateLinkResourcesServer contains the fakes for client PrivateLinkResourcesClient
	PrivateLinkResourcesServer PrivateLinkResourcesServer

	// SKUsServer contains the fakes for client SKUsClient
	SKUsServer SKUsServer

	// VolumeGroupsServer contains the fakes for client VolumeGroupsClient
	VolumeGroupsServer VolumeGroupsServer

	// VolumeSnapshotsServer contains the fakes for client VolumeSnapshotsClient
	VolumeSnapshotsServer VolumeSnapshotsServer

	// VolumesServer contains the fakes for client VolumesClient
	VolumesServer VolumesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armelasticsan.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armelasticsan.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                *ServerFactory
	trMu                               sync.Mutex
	trElasticSansServer                *ElasticSansServerTransport
	trOperationsServer                 *OperationsServerTransport
	trPrivateEndpointConnectionsServer *PrivateEndpointConnectionsServerTransport
	trPrivateLinkResourcesServer       *PrivateLinkResourcesServerTransport
	trSKUsServer                       *SKUsServerTransport
	trVolumeGroupsServer               *VolumeGroupsServerTransport
	trVolumeSnapshotsServer            *VolumeSnapshotsServerTransport
	trVolumesServer                    *VolumesServerTransport
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
	case "ElasticSansClient":
		initServer(s, &s.trElasticSansServer, func() *ElasticSansServerTransport { return NewElasticSansServerTransport(&s.srv.ElasticSansServer) })
		resp, err = s.trElasticSansServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PrivateEndpointConnectionsClient":
		initServer(s, &s.trPrivateEndpointConnectionsServer, func() *PrivateEndpointConnectionsServerTransport {
			return NewPrivateEndpointConnectionsServerTransport(&s.srv.PrivateEndpointConnectionsServer)
		})
		resp, err = s.trPrivateEndpointConnectionsServer.Do(req)
	case "PrivateLinkResourcesClient":
		initServer(s, &s.trPrivateLinkResourcesServer, func() *PrivateLinkResourcesServerTransport {
			return NewPrivateLinkResourcesServerTransport(&s.srv.PrivateLinkResourcesServer)
		})
		resp, err = s.trPrivateLinkResourcesServer.Do(req)
	case "SKUsClient":
		initServer(s, &s.trSKUsServer, func() *SKUsServerTransport { return NewSKUsServerTransport(&s.srv.SKUsServer) })
		resp, err = s.trSKUsServer.Do(req)
	case "VolumeGroupsClient":
		initServer(s, &s.trVolumeGroupsServer, func() *VolumeGroupsServerTransport { return NewVolumeGroupsServerTransport(&s.srv.VolumeGroupsServer) })
		resp, err = s.trVolumeGroupsServer.Do(req)
	case "VolumeSnapshotsClient":
		initServer(s, &s.trVolumeSnapshotsServer, func() *VolumeSnapshotsServerTransport {
			return NewVolumeSnapshotsServerTransport(&s.srv.VolumeSnapshotsServer)
		})
		resp, err = s.trVolumeSnapshotsServer.Do(req)
	case "VolumesClient":
		initServer(s, &s.trVolumesServer, func() *VolumesServerTransport { return NewVolumesServerTransport(&s.srv.VolumesServer) })
		resp, err = s.trVolumesServer.Do(req)
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
