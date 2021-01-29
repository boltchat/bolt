// boltchat
// Copyright (C) 2021  The boltchat Authors
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

package plugins

import (
	"errors"
	"time"

	"github.com/bolt-chat/protocol/errs"
	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/server/pools"
)

type RateLimiterPlugin struct {
	Amount int
	Time   time.Duration
}

func (p RateLimiterPlugin) OnMessage(msg *events.MessageEvent, c *pools.Connection) error {
	const amountKey string = "rate:a"
	const timeKey string = "rate:t"

	now := time.Now()

	if c.Data[amountKey] == nil {
		c.Data[amountKey] = 0
	}

	if c.Data[timeKey] == nil {
		c.Data[timeKey] = now
	}

	elapsed := now.Sub(c.Data[timeKey].(time.Time))

	if elapsed > p.Time {
		c.Data[timeKey] = now
		c.Data[amountKey] = 0
	} else if c.Data[amountKey].(int) >= p.Amount {
		return errors.New(errs.TooManyMessages)
	} else {
		c.Data[amountKey] = c.Data[amountKey].(int) + 1
	}

	return nil
}

func (RateLimiterPlugin) GetInfo() *PluginInfo {
	return &PluginInfo{
		Id: "rate-limiter",
	}
}
