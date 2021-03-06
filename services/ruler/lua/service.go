// Copyright © 2020 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lua

import (
	"github.com/wealdtech/walletd/core"
	"github.com/wealdtech/walletd/services/locker"
	"github.com/wealdtech/walletd/services/storage"
)

// Service is the ruler service.
type Service struct {
	locker *locker.Service
	store  storage.Service
	rules  []*core.Rule
}

// New creates a new ruler service
func New(locker *locker.Service, store storage.Service, rules []*core.Rule) (*Service, error) {
	return &Service{
		locker: locker,
		store:  store,
		rules:  rules,
	}, nil
}
