// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package config

import (
	"encoding/json"
	"path/filepath"

	"github.com/Azure/azure-sdk-for-go/eng/tools/generator/cmd/v2/common"
	"github.com/Azure/azure-sdk-for-go/eng/tools/generator/typespec"
)

type TypeSpecReleaseRequests map[string][]Track2Request

func (c TypeSpecReleaseRequests) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func (c TypeSpecReleaseRequests) Add(readme string, info Track2Request) {
	if !c.Contains(readme) {
		c[readme] = make([]Track2Request, 0)
	}
	c[readme] = append(c[readme], info)
}

func (c TypeSpecReleaseRequests) Contains(readme string) bool {
	_, ok := c[readme]
	return ok
}

type TypeSpecPakcageInfo struct {
	common.PackageInfo
	TspConfigPath string
}

func GetTypeSpecProjectsFromConfig(config *Config, specRoot string) (tspProjects map[string][]TypeSpecPakcageInfo, errResult error) {
	tspProjects = make(map[string][]TypeSpecPakcageInfo)
	for tspConfigPath, typespecRequests := range config.TypeSpecRequests {
		for _, releaseRequestInfo := range typespecRequests {
			localTspConfigPath := filepath.Join(specRoot, tspConfigPath)
			tspConfig, err := typespec.ParseTypeSpecConfig(localTspConfigPath)
			if err != nil {
				return nil, err
			}
			module, err := tspConfig.GetModuleName()
			if err != nil {
				return nil, err
			}

			tspProjects[module[0]] = append(tspProjects[localTspConfigPath], TypeSpecPakcageInfo{
				PackageInfo: common.PackageInfo{
					Name:        module[1],
					RequestLink: releaseRequestInfo.RequestLink,
					ReleaseDate: releaseRequestInfo.TargetDate,
					Tag:         releaseRequestInfo.PackageFlag,
				},
				TspConfigPath: localTspConfigPath,
			})
		}
	}

	return
}
