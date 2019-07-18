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

package kube

import (
	"github.com/PureStorage-OpenConnect/pure1-unplugged/pkg/puctl/exechelper"
)

// RunHelm runs helm and injects in our kubeconfig info
func RunHelm(exec exechelper.Executor, timeout int, helmArgs ...string) (string, error) {
	params := exechelper.ExecParams{
		CmdName: "helm",
		CmdArgs: append(helmArgs, "--kubeconfig="+kubeconf),
		Timeout: timeout,
	}
	result := exec.RunCommand(params)
	combinedOutput := result.OutBuf.String() + " " + result.ErrBuf.String()
	return combinedOutput, result.Error
}