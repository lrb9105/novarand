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

package catchup

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/Orca18/go-novarand/config"
	"github.com/Orca18/go-novarand/crypto"
	"github.com/Orca18/go-novarand/data"
	"github.com/Orca18/go-novarand/data/account"
	"github.com/Orca18/go-novarand/data/basics"
	"github.com/Orca18/go-novarand/data/bookkeeping"
	"github.com/Orca18/go-novarand/data/datatest"
	"github.com/Orca18/go-novarand/logging"
	"github.com/Orca18/go-novarand/protocol"
	"github.com/Orca18/go-novarand/rpcs"
	"github.com/Orca18/go-novarand/util/db"
)

func BenchmarkServiceFetchBlocks(b *testing.B) {
	b.StopTimer()
	// Make Ledger
	remote, local, release, genesisBalances := benchenv(b, 100, 500)
	defer release()

	require.NotNil(b, remote)
	require.NotNil(b, local)

	// Create a network and block service
	net := &httpTestPeerSource{}
	ls := rpcs.MakeBlockService(logging.TestingLog(b), config.GetDefaultLocal(), remote, net, "test genesisID")
	nodeA := basicRPCNode{}
	nodeA.RegisterHTTPHandler(rpcs.BlockServiceBlockPath, ls)
	nodeA.start()
	defer nodeA.stop()
	rootURL := nodeA.rootURL()
	net.addPeer(rootURL)

	cfg := config.GetDefaultLocal()
	cfg.Archival = true

	for i := 0; i < b.N; i++ {
		inMem := true
		local, err := data.LoadLedger(logging.TestingLog(b), b.Name()+"empty"+strconv.Itoa(i), inMem, protocol.ConsensusCurrentVersion, genesisBalances, "", crypto.Digest{}, nil, nil, cfg)
		require.NoError(b, err)

		// Make Service
		syncer := MakeService(logging.TestingLog(b), defaultConfig, net, local, new(mockedAuthenticator), nil, nil)
		b.StartTimer()
		syncer.Start()
		for w := 0; w < 1000; w++ {
			if remote.LastRound() == local.LastRound() {
				break
			}
			time.Sleep(10 * time.Millisecond)
		}
		b.StopTimer()
		syncer.Stop()
		require.Equal(b, remote.LastRound(), local.LastRound())
		local.Close()
	}
}

// one service
func benchenv(t testing.TB, numAccounts, numBlocks int) (ledger, emptyLedger *data.Ledger, release func(), genesisBalances bookkeeping.GenesisBalances) {
	P := numAccounts                                  // n accounts
	maxMoneyAtStart := uint64(10 * defaultRewardUnit) // max money start
	minMoneyAtStart := uint64(defaultRewardUnit)      // min money start

	accesssors := make([]db.Accessor, 0)
	release = func() {
		ledger.Close()
		emptyLedger.Close()
		for _, acc := range accesssors {
			acc.Close()
		}
	}
	// generate accounts
	genesis := make(map[basics.Address]basics.AccountData)
	gen := rand.New(rand.NewSource(2))
	parts := make([]account.Participation, P)
	for i := 0; i < P; i++ {
		access, err := db.MakeAccessor(t.Name()+"_root_benchenv"+strconv.Itoa(i), false, true)
		if err != nil {
			panic(err)
		}
		accesssors = append(accesssors, access)
		root, err := account.GenerateRoot(access)
		if err != nil {
			panic(err)
		}

		access, err = db.MakeAccessor(t.Name()+"_part_benchenv"+strconv.Itoa(i), false, true)
		if err != nil {
			panic(err)
		}
		accesssors = append(accesssors, access)
		part, err := account.FillDBWithParticipationKeys(access, root.Address(), 0, basics.Round(numBlocks),
			config.Consensus[protocol.ConsensusCurrentVersion].DefaultKeyDilution)
		if err != nil {
			panic(err)
		}

		startamt := basics.AccountData{
			Status:      basics.Online,
			MicroNovas:  basics.MicroNovas{Raw: uint64(minMoneyAtStart + (gen.Uint64() % (maxMoneyAtStart - minMoneyAtStart)))},
			SelectionID: part.VRFSecrets().PK,
			VoteID:      part.VotingSecrets().OneTimeSignatureVerifier,
		}
		short := root.Address()

		parts[i] = part.Participation
		genesis[short] = startamt
		part.Close()
	}

	genesis[basics.Address(sinkAddr)] = basics.AccountData{
		Status:     basics.NotParticipating,
		MicroNovas: basics.MicroNovas{Raw: uint64(1e3 * minMoneyAtStart)},
	}
	genesis[basics.Address(poolAddr)] = basics.AccountData{
		Status:     basics.NotParticipating,
		MicroNovas: basics.MicroNovas{Raw: uint64(1e3 * minMoneyAtStart)},
	}

	var err error
	genesisBalances = bookkeeping.MakeGenesisBalances(genesis, sinkAddr, poolAddr)
	const inMem = true
	cfg := config.GetDefaultLocal()
	cfg.Archival = true
	emptyLedger, err = data.LoadLedger(logging.TestingLog(t), t.Name()+"empty", inMem, protocol.ConsensusCurrentVersion, genesisBalances, "", crypto.Digest{}, nil, nil, cfg)
	require.NoError(t, err)

	ledger, err = datatest.FabricateLedger(logging.TestingLog(t), t.Name(), parts, genesisBalances, emptyLedger.LastRound()+basics.Round(numBlocks))
	require.NoError(t, err)
	require.Equal(t, ledger.LastRound(), emptyLedger.LastRound()+basics.Round(numBlocks))
	return ledger, emptyLedger, release, genesisBalances
}
