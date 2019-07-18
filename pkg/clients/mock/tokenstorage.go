// Copyright 2019, Pure Storage Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mock

import "github.com/PureStorage-OpenConnect/pure1-unplugged/pkg/common/resources"

// Type guard: ensure this implements the interface
var _ resources.APITokenStorage = (*APITokenStorageImpl)(nil)

// SaveToken is a mocked stub
func (a *APITokenStorageImpl) SaveToken(arrayID string, token string) error {
	args := a.Called(arrayID, token)
	return args.Error(0)
}

// HasToken is a mocked stub
func (a *APITokenStorageImpl) HasToken(arrayID string) (bool, error) {
	args := a.Called(arrayID)
	return args.Bool(0), args.Error(1)
}

// GetToken is a mocked stub
func (a *APITokenStorageImpl) GetToken(arrayID string) (string, error) {
	args := a.Called(arrayID)
	return args.String(0), args.Error(1)
}

// DeleteToken is a mocked stub
func (a *APITokenStorageImpl) DeleteToken(arrayID string) error {
	args := a.Called(arrayID)
	return args.Error(0)
}