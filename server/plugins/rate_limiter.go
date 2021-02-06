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

package plugins

import (
	"errors"
	"time"

	"github.com/boltchat/protocol/errs"
	"github.com/boltchat/protocol/events"
	"github.com/boltchat/server/pools"
)

type RateLimiterPlugin struct {
	Amount int
	Time   time.Duration
}

func (RateLimiterPlugin) GetInfo() *PluginInfo {
	return &PluginInfo{
		Id: "rate-limiter",
	}
}

func (p RateLimiterPlugin) OnMessage(msg *events.MessageData, c *pools.Connection) error {
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

func (p RateLimiterPlugin) OnIdentify(data *events.JoinData, c *pools.Connection) error {
	return nil
}
