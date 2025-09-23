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

// Package version ...
package version

import (
	"fmt"
	"strings"

	"github.com/TencentBlueKing/blueking-micro-apigateway/apiserver/pkg/constant"
)

// ToXVersion 将完整版本号转换为 x 版本号，如 3.1 -> 3.1.x, 3.2.1 -> 3.2.x
func ToXVersion(version string) (constant.APISIXVersion, error) {
	parts := strings.Split(version, ".")
	if len(parts) < 2 || len(parts) > 3 {
		return "", fmt.Errorf("invalid version: %s", version)
	}
	// 如果版本号是两部分，则添加一个 x
	if len(parts) == 2 {
		parts = append(parts, "X")
	} else {
		// 如果版本号是三部分，则将第三部分替换为 x
		parts[2] = "X"
	}

	return constant.APISIXVersion(strings.Join(parts, ".")), nil
}
