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

package filter

import (
	"context"
	"encoding/json"

	"github.com/erda-project/erda/apistructs"
	protocol "github.com/erda-project/erda/modules/openapi/component-protocol"
)

func (i *ComponentFilter) unmarshal(c *apistructs.Component) error {
	stateValue, err := json.Marshal(c.State)
	if err != nil {
		return err
	}
	var state State
	err = json.Unmarshal(stateValue, &state)
	if err != nil {
		return err
	}
	i.State = state
	return nil
}

func (i *ComponentFilter) marshal(c *apistructs.Component) error {
	stateValue, err := json.Marshal(i.State)
	if err != nil {
		return err
	}
	var state map[string]interface{}
	err = json.Unmarshal(stateValue, &state)
	if err != nil {
		return err
	}
	c.State = state //why not work
	c.Props = i.Props
	c.Operations = i.Operations
	return nil
}

func (i *ComponentFilter) Render(ctx context.Context, c *apistructs.Component, _ apistructs.ComponentProtocolScenario, event apistructs.ComponentEvent, gs *apistructs.GlobalStateData) (err error) {
	if event.Operation != apistructs.InitializeOperation && event.Operation != apistructs.RenderingOperation {
		err = i.unmarshal(c)
		if err != nil {
			return err
		}
	}

	defer func() {
		fail := i.marshal(c)
		if err == nil && fail != nil {
			err = fail
		}
	}()

	i.initProperty()
	if event.Operation == apistructs.FilterOrgsOperationKey {
		if v, ok := i.State.Values["title"]; ok {
			i.State.SearchEntry = v
			i.State.SearchRefresh = true
		}
	}
	return nil
}

func (i *ComponentFilter) initProperty() {
	i.Props = Props{Delay: 1000}
	i.Operations = map[string]interface{}{
		"filter": Operation{
			Key:    "filter",
			Reload: true,
		},
	}
	i.State.Conditions = []Condition{
		{
			Key:         "title",
			Label:       "标题",
			EmptyText:   "全部",
			Fixed:       true,
			ShowIndex:   2,
			Placeholder: "搜索",
			Type:        "input",
		},
	}
}

func RenderCreator() protocol.CompRender {
	return &ComponentFilter{}
}
