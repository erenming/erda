// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd_test

import (
	"testing"

	"github.com/erda-project/erda/tools/cli/cmd"
	"github.com/erda-project/erda/tools/cli/command"
)

func TestHepaHostFromOpenapi(t *testing.T) {
	var ctx = command.Context{
		CurrentHost: "https://openapi.erda.cloud",
	}
	var hepa = "https://hepa.erda.cloud"
	if err := cmd.HepaHostFromOpenapi(&ctx, ""); err != nil {
		t.Fatal(err)
	}
	if ctx.CurrentHost != hepa {
		t.Fatal("ctx.CurrentHost != hepa")
	}

	ctx.CurrentHost = "https://openapi.erda.cloud"
	if err := cmd.HepaHostFromOpenapi(&ctx, hepa); err != nil {
		t.Fatal(err)
	}
	if ctx.CurrentHost != hepa {
		t.Fatal("ctx.CurrentHost != hepa")
	}
}
