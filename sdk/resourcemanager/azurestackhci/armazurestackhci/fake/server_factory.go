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

// ServerFactory is a fake server for instances of the armazurestackhci.ClientFactory type.
type ServerFactory struct {
	ArcSettingsServer        ArcSettingsServer
	ClustersServer           ClustersServer
	DeploymentSettingsServer DeploymentSettingsServer
	EdgeDeviceJobsServer     EdgeDeviceJobsServer
	EdgeDevicesServer        EdgeDevicesServer
	ExtensionsServer         ExtensionsServer
	OffersServer             OffersServer
	OperationsServer         OperationsServer
	PublishersServer         PublishersServer
	SKUsServer               SKUsServer
	SecuritySettingsServer   SecuritySettingsServer
	UpdateRunsServer         UpdateRunsServer
	UpdateSummariesServer    UpdateSummariesServer
	UpdatesServer            UpdatesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armazurestackhci.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armazurestackhci.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                        *ServerFactory
	trMu                       sync.Mutex
	trArcSettingsServer        *ArcSettingsServerTransport
	trClustersServer           *ClustersServerTransport
	trDeploymentSettingsServer *DeploymentSettingsServerTransport
	trEdgeDeviceJobsServer     *EdgeDeviceJobsServerTransport
	trEdgeDevicesServer        *EdgeDevicesServerTransport
	trExtensionsServer         *ExtensionsServerTransport
	trOffersServer             *OffersServerTransport
	trOperationsServer         *OperationsServerTransport
	trPublishersServer         *PublishersServerTransport
	trSKUsServer               *SKUsServerTransport
	trSecuritySettingsServer   *SecuritySettingsServerTransport
	trUpdateRunsServer         *UpdateRunsServerTransport
	trUpdateSummariesServer    *UpdateSummariesServerTransport
	trUpdatesServer            *UpdatesServerTransport
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
	case "ArcSettingsClient":
		initServer(s, &s.trArcSettingsServer, func() *ArcSettingsServerTransport { return NewArcSettingsServerTransport(&s.srv.ArcSettingsServer) })
		resp, err = s.trArcSettingsServer.Do(req)
	case "ClustersClient":
		initServer(s, &s.trClustersServer, func() *ClustersServerTransport { return NewClustersServerTransport(&s.srv.ClustersServer) })
		resp, err = s.trClustersServer.Do(req)
	case "DeploymentSettingsClient":
		initServer(s, &s.trDeploymentSettingsServer, func() *DeploymentSettingsServerTransport {
			return NewDeploymentSettingsServerTransport(&s.srv.DeploymentSettingsServer)
		})
		resp, err = s.trDeploymentSettingsServer.Do(req)
	case "EdgeDeviceJobsClient":
		initServer(s, &s.trEdgeDeviceJobsServer, func() *EdgeDeviceJobsServerTransport {
			return NewEdgeDeviceJobsServerTransport(&s.srv.EdgeDeviceJobsServer)
		})
		resp, err = s.trEdgeDeviceJobsServer.Do(req)
	case "EdgeDevicesClient":
		initServer(s, &s.trEdgeDevicesServer, func() *EdgeDevicesServerTransport { return NewEdgeDevicesServerTransport(&s.srv.EdgeDevicesServer) })
		resp, err = s.trEdgeDevicesServer.Do(req)
	case "ExtensionsClient":
		initServer(s, &s.trExtensionsServer, func() *ExtensionsServerTransport { return NewExtensionsServerTransport(&s.srv.ExtensionsServer) })
		resp, err = s.trExtensionsServer.Do(req)
	case "OffersClient":
		initServer(s, &s.trOffersServer, func() *OffersServerTransport { return NewOffersServerTransport(&s.srv.OffersServer) })
		resp, err = s.trOffersServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PublishersClient":
		initServer(s, &s.trPublishersServer, func() *PublishersServerTransport { return NewPublishersServerTransport(&s.srv.PublishersServer) })
		resp, err = s.trPublishersServer.Do(req)
	case "SKUsClient":
		initServer(s, &s.trSKUsServer, func() *SKUsServerTransport { return NewSKUsServerTransport(&s.srv.SKUsServer) })
		resp, err = s.trSKUsServer.Do(req)
	case "SecuritySettingsClient":
		initServer(s, &s.trSecuritySettingsServer, func() *SecuritySettingsServerTransport {
			return NewSecuritySettingsServerTransport(&s.srv.SecuritySettingsServer)
		})
		resp, err = s.trSecuritySettingsServer.Do(req)
	case "UpdateRunsClient":
		initServer(s, &s.trUpdateRunsServer, func() *UpdateRunsServerTransport { return NewUpdateRunsServerTransport(&s.srv.UpdateRunsServer) })
		resp, err = s.trUpdateRunsServer.Do(req)
	case "UpdateSummariesClient":
		initServer(s, &s.trUpdateSummariesServer, func() *UpdateSummariesServerTransport {
			return NewUpdateSummariesServerTransport(&s.srv.UpdateSummariesServer)
		})
		resp, err = s.trUpdateSummariesServer.Do(req)
	case "UpdatesClient":
		initServer(s, &s.trUpdatesServer, func() *UpdatesServerTransport { return NewUpdatesServerTransport(&s.srv.UpdatesServer) })
		resp, err = s.trUpdatesServer.Do(req)
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
