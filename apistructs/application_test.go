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

package apistructs

import "testing"

func TestCheckAppMode(t *testing.T) {
	type args struct {
		mode string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "otherName",
			args: args{
				mode: "_&((()))",
			},
			wantErr: true,
		},
		{
			name: string(ApplicationModeService),
			args: args{
				mode: string(ApplicationModeService),
			},
			wantErr: false,
		},
		{
			name: string(ApplicationModeLibrary),
			args: args{
				mode: string(ApplicationModeLibrary),
			},
			wantErr: false,
		},
		{
			name: string(ApplicationModeMobile),
			args: args{
				mode: string(ApplicationModeMobile),
			},
			wantErr: false,
		},
		{
			name: string(ApplicationModeProjectService),
			args: args{
				mode: string(ApplicationModeProjectService),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ApplicationMode(tt.args.mode).CheckAppMode(); (err != nil) != tt.wantErr {
				t.Errorf("CheckAppMode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
