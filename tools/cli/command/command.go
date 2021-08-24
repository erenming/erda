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

package command

type Command struct {
	// default: rootcmd
	ParentName string
	Name       string
	ShortHelp  string
	LongHelp   string
	Example    string
	// dont show up in list of cmds
	Hidden         bool
	DontHideCursor bool
	/* args:
	        []Arg {
		       IPArg{},
		       StringArg{},
		       BoolArg{},
		}
	*/
	Args []Arg
	/* flags:
	        []Flag{
	                StringFlag{"H", "host", "1.2.3.4", "doc"},
	                BoolFlag{"A", "another", true, "doc"},
	                IntFlag{"O", "ohyoyo", 1, "doc"},
		}
	*/
	Flags []Flag
	/* actually type:
		func(ctx Context, arg1 IPArg, arg2 string, arg3 bool, host string, anotherone bool, ohyoyo int) error

	`host`, `anotherone`, `ohyoyo` is generated by `Flags`
	*/
	Run interface{}
}
