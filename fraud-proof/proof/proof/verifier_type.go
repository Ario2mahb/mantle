// Copyright 2022, Specular contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proof

type VerifierType int

const (
	VerifierTypeStackOp         VerifierType = 0
	VerifierTypeEnvironmentalOp VerifierType = 1
	VerifierTypeMemoryOp        VerifierType = 2
	VerifierTypeStorageOp       VerifierType = 3
	VerifierTypeCallOp          VerifierType = 4
	VerifierTypeInvalidOp       VerifierType = 5
	VerifierTypeInterTx         VerifierType = 6
	VerifierTypeBlockInit       VerifierType = 7
)
