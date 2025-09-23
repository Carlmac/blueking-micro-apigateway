/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - 微网关(BlueKing - Micro APIGateway) available.
 * Copyright (C) 2025 Tencent. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 *     http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 * to the current version of the project delivered to anyone in the future.
 */

package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/TencentBlueKing/blueking-micro-apigateway/apiserver/pkg/constant"
)

func TestGetResourceSchema(t *testing.T) {
	tests := []struct {
		name       string
		version    constant.APISIXVersion
		shouldFail bool
	}{
		{
			name:       "APISIX 3.2 - Existing Resource",
			version:    constant.APISIXVersion32,
			shouldFail: false,
		},
		{
			name:       "APISIX 3.3 - Existing Resource",
			version:    constant.APISIXVersion33,
			shouldFail: false,
		},
		{
			name:       "APISIX 3.11 - Existing Resource",
			version:    constant.APISIXVersion311,
			shouldFail: false,
		},
		{
			name:       "APISIX 3.13 - Existing Resource",
			version:    constant.APISIXVersion313,
			shouldFail: false,
		},
		{
			name:       "Invalid Version",
			version:    "invalid_version",
			shouldFail: true,
		},
	}

	for _, tt := range tests {
		for _, resource := range constant.ResourceTypeList {
			t.Run(tt.name, func(t *testing.T) {
				result := GetResourceSchema(tt.version, resource.String())
				if tt.shouldFail {
					assert.Nil(t, result)
				} else {
					assert.NotNil(t, result)
				}
			})
		}
	}
}

func TestGetPluginSchema(t *testing.T) {
	tests := []struct {
		name       string
		pluginName string
		schemaType string
		shouldFail bool
	}{
		{
			name:       "Plugin Schema",
			pluginName: "key-auth",
			schemaType: "",
			shouldFail: false,
		},
		{
			name:       "Consumer Schema",
			pluginName: "basic-auth",
			schemaType: "consumer",
			shouldFail: false,
		},
		{
			name:       "Metadata Schema",
			pluginName: "authz-casbin",
			schemaType: "metadata",
			shouldFail: false,
		},
		{
			name:       "Stream Schema",
			pluginName: "mqtt-proxy",
			schemaType: "stream",
			shouldFail: false,
		},
		{
			name:       "Non-existing Plugin",
			pluginName: "non_existing",
			schemaType: "",
			shouldFail: true,
		},
	}

	// 查询所有版本的 schema
	for version := range schemaVersionMap {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := GetPluginSchema(version, tt.pluginName, tt.schemaType)
				if tt.shouldFail {
					assert.Nil(t, result)
				} else {
					assert.NotNil(t, result)
				}
			})
		}
	}
}
