// Copyright 2021 Dolthub, Inc.
//
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

package globalstate

import (
	"context"
	"sync"

	"github.com/dolthub/dolt/go/libraries/doltcore/doltdb"
	"github.com/dolthub/dolt/go/libraries/doltcore/sqle/dsess"
	"github.com/dolthub/go-mysql-server/sql"

	"github.com/dolthub/dolt/go/libraries/doltcore/ref"
)

type StateProvider interface {
	GetGlobalState() GlobalState
}

func NewGlobalStateStore() GlobalState {
	return GlobalState{
		trackerMap: make(map[ref.WorkingSetRef]AutoIncrementTracker),
		mu:         &sync.Mutex{},
	}
}

func NewGlobalStateStoreForDb(ctx context.Context, db *doltdb.DoltDB) (GlobalState, error) {
	branches, err := db.GetBranches(ctx)
	if err != nil {
		return GlobalState{}, err
	}

	for _, b := range branches {
		ws, err := ref.WorkingSetRefForHead(b)
		if err != nil {
			return GlobalState{}, err
		}


	}

	return GlobalState{
		trackerMap: make(map[ref.WorkingSetRef]AutoIncrementTracker),
		mu:         &sync.Mutex{},
	}, nil
}

type GlobalState struct {
	trackerMap map[ref.WorkingSetRef]AutoIncrementTracker
	mu         *sync.Mutex
}

func (g GlobalState) GetAutoIncrementTracker(ctx *sql.Context, ws *doltdb.WorkingSet) (AutoIncrementTracker, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	perBranch, err := dsess.GetBooleanSystemVar(ctx, dsess.PerBranchAutoIncrement)
	if err != nil {
		return AutoIncrementTracker{}, err
	}

	if !perBranch {

	}

	ref := ws.Ref()
	ait, ok := g.trackerMap[ref]
	if ok {
		return ait, nil
	}

	ait, err = NewAutoIncrementTracker(ctx, ws)
	if err != nil {
		return AutoIncrementTracker{}, err
	}
	g.trackerMap[ref] = ait

	return ait, nil
}
