// Copyright 2021 The boltchat Authors
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

package version

import (
	"fmt"
)

func FormatVersion(versions []*Version) string {
	str := "Copyright (c) 2021 The boltchat Authors\n"

	for i, v := range versions {
		str += fmt.Sprintf("%s version %s", v.Type, v.VersionString)

		if i < len(versions)-1 {
			str += "\n"
		}
	}

	return str
}
