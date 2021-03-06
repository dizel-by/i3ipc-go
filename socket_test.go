// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package i3ipc

import (
	"testing"
)

func TestGetIPCSocket(t *testing.T) {
	startTestIPCSocket(testMessages[""])
	ipc, _ := GetIPCSocket()
	ipc.Close()
	if ipc.open {
		t.Error("IPC socket appears open after closing.")
	}
}

func TestRaw(t *testing.T) {
	ipc, _ := GetIPCSocket()

	go startTestIPCSocket(testMessages["workspaces"])

	_, err := ipc.Raw(I3GetWorkspaces, "")
	if err != nil {
		t.Errorf("Raw message sending failed: %v", err)
	}
}
