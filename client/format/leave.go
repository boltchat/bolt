// bolt.chat
// Copyright (C) 2021  Kees van Voorthuizen
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package format

import (
	"encoding/json"
	"fmt"

	"github.com/bolt-chat/protocol/events"
	"github.com/fatih/color"
	"github.com/gdamore/tcell/v2"
)

func FormatLeave(e *events.BaseEvent) string {
	leaveEvt := &events.JoinEvent{}
	json.Unmarshal(*e.Raw, leaveEvt)

	return color.HiMagentaString(
		fmt.Sprintf("%s %s left the room.", string(tcell.RuneDiamond), leaveEvt.User.Nickname),
	)
}
