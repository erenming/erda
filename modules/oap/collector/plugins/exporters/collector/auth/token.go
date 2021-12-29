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

package auth

import (
	"fmt"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

type TokenAuth struct {
	// token file or content
	Token string `mapstructure:"token_file"`
}

func (t *TokenAuth) Secure(req *http.Request) {
	// TODO
}

func NewTokenAuth(cfg map[string]interface{}) (*TokenAuth, error) {
	ta := TokenAuth{}
	err := mapstructure.Decode(cfg, &ta)
	if err != nil {
		return nil, fmt.Errorf("decode err: %w", err)
	}
	if ta.Token != "" {
		return nil, fmt.Errorf("empty token: %+v", ta)
	}
	return &ta, nil
}
