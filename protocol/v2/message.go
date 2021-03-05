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

package protocol

/*
Message represents a message that is
either transmitted or stored locally.
*/
type Message struct {
	Content     string `msgpack:"body"`
	Signature   string `msgpack:"sig,omitempty"`
	Fingerprint string `msgpack:"fprint,omitempty"`
	User        *User  `msgpack:"user"`
}
