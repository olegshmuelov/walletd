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

package ruler

import (
	"context"

	"github.com/wealdtech/walletd/core"
)

var (
	// ActionSign is the action of signing data.
	ActionSign = "Sign"
	// ActionSignBeaconAttestation is the action of signing a beacon attestation.
	ActionSignBeaconAttestation = "Sign beacon attestation"
	// ActionSignBeaconProposal is the action of signing a beacon proposal.
	ActionSignBeaconProposal = "Sign beacon proposal"
	// ActionAccessAccount is the action of accessing an account.
	ActionAccessAccount = "Access account"
)

// Service provides an interface to check requests against a rules engine.
type Service interface {
	// RunRules runs a set of rules for the given information.
	RunRules(context.Context, string, string, string, []byte, interface{}) core.RulesResult
}
