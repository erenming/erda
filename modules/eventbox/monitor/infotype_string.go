// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Code generated by "stringer -type=InfoType"; DO NOT EDIT.

package monitor

import "strconv"

const _InfoType_name = "EtcdInputEtcdInputDropHTTPInputDINGDINGOutputDINGDINGWorkNoticeOutputMYSQLOutputHTTPOutputLastType"

var _InfoType_index = [...]uint8{0, 9, 22, 31, 45, 69, 80, 90, 98}

func (i InfoType) String() string {
	if i < 0 || i >= InfoType(len(_InfoType_index)-1) {
		return "InfoType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _InfoType_name[_InfoType_index[i]:_InfoType_index[i+1]]
}
