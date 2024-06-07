// Copyright (C) 2019-2022 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package ledger

import (
	"context"
	"fmt"
	"testing"

	"github.com/algorand/go-deadlock"
	"github.com/stretchr/testify/require"

	"github.com/Orca18/go-novarand/agreement"
	"github.com/Orca18/go-novarand/config"
	"github.com/Orca18/go-novarand/crypto"
	"github.com/Orca18/go-novarand/data/basics"
	"github.com/Orca18/go-novarand/data/bookkeeping"
	"github.com/Orca18/go-novarand/data/transactions"
	ledgertesting "github.com/Orca18/go-novarand/ledger/testing"
	"github.com/Orca18/go-novarand/logging"
	"github.com/Orca18/go-novarand/protocol"
	"github.com/Orca18/go-novarand/util/execpool"
)

func BenchmarkManyAccounts(b *testing.B) {
	deadlock.Opts.Disable = true

	b.StopTimer()

	genesisInitState, addrs, _ := ledgertesting.Genesis(1)
	addr := addrs[0]

	dbName := fmt.Sprintf("%s.%d", b.Name(), crypto.RandUint64())
	const inMem = true
	cfg := config.GetDefaultLocal()
	cfg.Archival = true
	l, err := OpenLedger(logging.Base(), dbName, inMem, genesisInitState, cfg)
	require.NoError(b, err)
	defer l.Close()

	blk := genesisInitState.Block
	for i := 0; i < b.N; i++ {
		blk = bookkeeping.MakeBlock(blk.BlockHeader)

		proto, ok := config.Consensus[blk.CurrentProtocol]
		require.True(b, ok)

		var txbytes int
		for {
			var st transactions.SignedTxn
			crypto.RandBytes(st.Sig[:])
			st.Txn.Type = protocol.PaymentTx
			st.Txn.Sender = addr
			st.Txn.Fee = basics.MicroNovas{Raw: 1}
			st.Txn.Amount = basics.MicroNovas{Raw: 1}
			crypto.RandBytes(st.Txn.Receiver[:])

			txib, err := blk.EncodeSignedTxn(st, transactions.ApplyData{})
			require.NoError(b, err)

			txlen := len(protocol.Encode(&txib))
			if txbytes+txlen > proto.MaxTxnBytesPerBlock {
				break
			}

			txbytes += txlen
			blk.Payset = append(blk.Payset, txib)
		}

		var c agreement.Certificate
		b.StartTimer()
		err := l.AddBlock(blk, c)
		b.StopTimer()
		require.NoError(b, err)
	}
}

func BenchmarkValidate(b *testing.B) {
	b.StopTimer()

	genesisInitState, addrs, keys := ledgertesting.Genesis(10000)

	backlogPool := execpool.MakeBacklog(nil, 0, execpool.LowPriority, nil)
	defer backlogPool.Shutdown()

	dbName := fmt.Sprintf("%s.%d", b.Name(), crypto.RandUint64())
	const inMem = true
	cfg := config.GetDefaultLocal()
	cfg.Archival = true
	l, err := OpenLedger(logging.Base(), dbName, inMem, genesisInitState, cfg)
	require.NoError(b, err)
	defer l.Close()

	blk := genesisInitState.Block
	for i := 0; i < b.N; i++ {
		newblk := bookkeeping.MakeBlock(blk.BlockHeader)

		proto, ok := config.Consensus[newblk.CurrentProtocol]
		require.True(b, ok)

		var txbytes int
		for i := 0; i < 10000; i++ {
			t := transactions.Transaction{
				Type: protocol.PaymentTx,
				Header: transactions.Header{
					Sender:     addrs[i],
					Fee:        basics.MicroNovas{Raw: 1},
					FirstValid: newblk.Round(),
					LastValid:  newblk.Round(),
				},
				PaymentTxnFields: transactions.PaymentTxnFields{
					Amount: basics.MicroNovas{Raw: 1},
				},
			}
			crypto.RandBytes(t.Receiver[:])
			st := t.Sign(keys[i])

			txib, err := newblk.EncodeSignedTxn(st, transactions.ApplyData{})
			require.NoError(b, err)

			txlen := len(protocol.Encode(&txib))
			if txbytes+txlen > proto.MaxTxnBytesPerBlock {
				break
			}

			txbytes += txlen
			newblk.Payset = append(newblk.Payset, txib)
		}

		newblk.BlockHeader.TxnCommitments, err = newblk.PaysetCommit()
		require.NoError(b, err)

		b.StartTimer()
		_, err = l.Validate(context.Background(), newblk, backlogPool)
		b.StopTimer()
		require.NoError(b, err)
	}
}
