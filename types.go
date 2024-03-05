/*******************************************************************************
*   (c) 2018 - 2022 ZondaX AG
*
*  Licensed under the Apache License, Version 2.0 (the "License");
*  you may not use this file except in compliance with the License.
*  You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*  distributed under the License is distributed on an "AS IS" BASIS,
*  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*  See the License for the specific language governing permissions and
*  limitations under the License.
********************************************************************************/

package ledger_filecoin_go

import (
	// "fmt"
	"github.com/zondax/ledger-go"
)

const (
	CLA                 = 0x06

	INSGetVersion       = 0
	INSGetAddrSECP256K1 = 1
	INSSignSECP256K1    = 2
	INSSignDataCap      = 5
	INSSignClientDeal   = 6
	INSSignRawBytes     = 7
)

const (
	PayloadChunkInit = 0
	PayloadChunkAdd  = 1
	PayloadChunkLast = 2
)

const HardenCount int = 2

// LedgerFilecoin represents a connection to the Ledger app
type LedgerFilecoin struct {
	api     ledger_go.LedgerDevice
	version VersionInfo
}

type SignatureAnswer struct {
	r            []byte
	s            []byte
	v            uint8
	derSignature []byte
}

