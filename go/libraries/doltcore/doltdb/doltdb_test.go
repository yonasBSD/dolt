// Copyright 2019 Dolthub, Inc.
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

package doltdb

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dolthub/dolt/go/libraries/doltcore/dbfactory"
	"github.com/dolthub/dolt/go/libraries/doltcore/doltdb/durable"
	"github.com/dolthub/dolt/go/libraries/doltcore/ref"
	"github.com/dolthub/dolt/go/libraries/doltcore/row"
	"github.com/dolthub/dolt/go/libraries/doltcore/schema"
	"github.com/dolthub/dolt/go/libraries/utils/filesys"
	"github.com/dolthub/dolt/go/libraries/utils/test"
	"github.com/dolthub/dolt/go/store/datas"
	"github.com/dolthub/dolt/go/store/hash"
	"github.com/dolthub/dolt/go/store/prolly/tree"
	"github.com/dolthub/dolt/go/store/types"
)

const (
	idTag        = 0
	firstTag     = 1
	lastTag      = 2
	isMarriedTag = 3
	ageTag       = 4
	emptyTag     = 5
)
const testSchemaIndexName = "idx_name"
const testSchemaIndexAge = "idx_age"

var id0, _ = uuid.NewRandom()
var id1, _ = uuid.NewRandom()
var id2, _ = uuid.NewRandom()
var id3, _ = uuid.NewRandom()

func createTestSchema(t *testing.T) schema.Schema {
	colColl := schema.NewColCollection(
		schema.NewColumn("id", idTag, types.UUIDKind, true, schema.NotNullConstraint{}),
		schema.NewColumn("first", firstTag, types.StringKind, false, schema.NotNullConstraint{}),
		schema.NewColumn("last", lastTag, types.StringKind, false, schema.NotNullConstraint{}),
		schema.NewColumn("is_married", isMarriedTag, types.BoolKind, false),
		schema.NewColumn("age", ageTag, types.UintKind, false),
		schema.NewColumn("empty", emptyTag, types.IntKind, false),
	)
	sch, err := schema.SchemaFromCols(colColl)
	require.NoError(t, err)
	_, err = sch.Indexes().AddIndexByColTags(testSchemaIndexName, []uint64{firstTag, lastTag}, nil, schema.IndexProperties{IsUnique: false, Comment: ""})
	require.NoError(t, err)
	_, err = sch.Indexes().AddIndexByColTags(testSchemaIndexAge, []uint64{ageTag}, nil, schema.IndexProperties{IsUnique: false, Comment: ""})
	require.NoError(t, err)
	return sch
}

func CreateTestTable(vrw types.ValueReadWriter, ns tree.NodeStore, tSchema schema.Schema, rowData durable.Index) (*Table, error) {
	tbl, err := NewTable(context.Background(), vrw, ns, tSchema, rowData, nil, nil)

	if err != nil {
		return nil, err
	}

	return tbl, nil
}

func createTestRowData(t *testing.T, vrw types.ValueReadWriter, ns tree.NodeStore, sch schema.Schema) durable.Index {
	if types.Format_Default == types.Format_DOLT {
		idx, err := durable.NewEmptyPrimaryIndex(context.Background(), vrw, ns, sch)
		require.NoError(t, err)
		return idx
	}

	vals := []row.TaggedValues{
		{idTag: types.UUID(id0), firstTag: types.String("bill"), lastTag: types.String("billerson"), ageTag: types.Uint(53)},
		{idTag: types.UUID(id1), firstTag: types.String("eric"), lastTag: types.String("ericson"), isMarriedTag: types.Bool(true), ageTag: types.Uint(21)},
		{idTag: types.UUID(id2), firstTag: types.String("john"), lastTag: types.String("johnson"), isMarriedTag: types.Bool(false), ageTag: types.Uint(53)},
		{idTag: types.UUID(id3), firstTag: types.String("robert"), lastTag: types.String("robertson"), ageTag: types.Uint(36)},
	}

	var err error
	rows := make([]row.Row, len(vals))

	m, err := types.NewMap(context.Background(), vrw)
	assert.NoError(t, err)
	ed := m.Edit()

	for i, val := range vals {
		r, err := row.New(vrw.Format(), sch, val)
		require.NoError(t, err)
		rows[i] = r
		ed = ed.Set(r.NomsMapKey(sch), r.NomsMapValue(sch))
	}

	m, err = ed.Map(context.Background())
	assert.NoError(t, err)
	return durable.IndexFromNomsMap(m, vrw, ns)
}

func TestIsValidTableName(t *testing.T) {
	assert.True(t, IsValidTableName("a"))
	assert.True(t, IsValidTableName("a1"))
	assert.True(t, IsValidTableName("_a1"))
	assert.True(t, IsValidTableName("a1_b_c------1"))
	assert.True(t, IsValidTableName("Add-098234_lkjasdf0p98"))
	assert.True(t, IsValidTableName("1"))
	assert.True(t, IsValidTableName("-"))
	assert.True(t, IsValidTableName("-a"))
	assert.True(t, IsValidTableName("a1-"))
	assert.True(t, IsValidTableName("ab!!c"))
	assert.True(t, IsValidTableName("           space"))
	assert.True(t, IsValidTableName("this     is     ok"))
	assert.True(t, IsValidTableName(" ~!@#$%^&*()_+`-=[]{}|;':\",./<>?"))
	assert.True(t, IsValidTableName("あえいおう"))
	assert.True(t, IsValidTableName("might/be/problematic"))
	assert.False(t, IsValidTableName(""))
	assert.False(t, IsValidTableName(string(rune(0))))
	assert.False(t, IsValidTableName("𐌃𐌏𐌋𐌕"))
	assert.False(t, IsValidTableName("space            "))
}

// DO NOT CHANGE THIS TEST
// It is necessary to ensure consistent system table definitions
// for more info: https://github.com/dolthub/dolt/pull/663
func TestSystemTableTags(t *testing.T) {
	var sysTableMin uint64 = 1 << 51

	t.Run("asdf", func(t *testing.T) {
		assert.Equal(t, sysTableMin, schema.SystemTableReservedMin)
	})
	t.Run("dolt_doc tags", func(t *testing.T) {
		docTableMin := sysTableMin + uint64(5)
		assert.Equal(t, docTableMin+0, schema.DocNameTag)
		assert.Equal(t, docTableMin+1, schema.DocTextTag)
	})
	t.Run("dolt_history_ tags", func(t *testing.T) {
		doltHistoryMin := sysTableMin + uint64(1000)
		assert.Equal(t, doltHistoryMin+0, schema.HistoryCommitterTag)
		assert.Equal(t, doltHistoryMin+1, schema.HistoryCommitHashTag)
		assert.Equal(t, doltHistoryMin+2, schema.HistoryCommitDateTag)
	})
	t.Run("dolt_diff_ tags", func(t *testing.T) {
		diffTableMin := sysTableMin + uint64(2000)
		assert.Equal(t, diffTableMin+0, schema.DiffCommitTag)
	})
	t.Run("dolt_query_catalog tags", func(t *testing.T) {
		queryCatalogMin := sysTableMin + uint64(3005)
		assert.Equal(t, queryCatalogMin+0, schema.QueryCatalogIdTag)
		assert.Equal(t, queryCatalogMin+1, schema.QueryCatalogOrderTag)
		assert.Equal(t, queryCatalogMin+2, schema.QueryCatalogNameTag)
		assert.Equal(t, queryCatalogMin+3, schema.QueryCatalogQueryTag)
		assert.Equal(t, queryCatalogMin+4, schema.QueryCatalogDescriptionTag)
	})
	t.Run("dolt_schemas tags", func(t *testing.T) {
		doltSchemasMin := sysTableMin + uint64(4007)
		assert.Equal(t, doltSchemasMin+0, schema.DoltSchemasIdTag)
		assert.Equal(t, doltSchemasMin+1, schema.DoltSchemasTypeTag)
		assert.Equal(t, doltSchemasMin+2, schema.DoltSchemasNameTag)
		assert.Equal(t, doltSchemasMin+3, schema.DoltSchemasFragmentTag)
		assert.Equal(t, doltSchemasMin+4, schema.DoltSchemasExtraTag)
	})
	t.Run("dolt_conflicts_ tags", func(t *testing.T) {
		doltConflictsMin := sysTableMin + uint64(7000)
		assert.Equal(t, doltConflictsMin+0, schema.DoltConflictsOurDiffTypeTag)
		assert.Equal(t, doltConflictsMin+1, schema.DoltConflictsTheirDiffTypeTag)
		assert.Equal(t, doltConflictsMin+2, schema.DoltConflictsBaseCardinalityTag)
		assert.Equal(t, doltConflictsMin+3, schema.DoltConflictsOurCardinalityTag)
		assert.Equal(t, doltConflictsMin+4, schema.DoltConflictsTheirCardinalityTag)
	})
}

func TestEmptyInMemoryRepoCreation(t *testing.T) {
	ddb, err := LoadDoltDB(context.Background(), types.Format_Default, InMemDoltDB, filesys.LocalFS)

	if err != nil {
		t.Fatal("Failed to load db")
	}
	defer ddb.Close()

	err = ddb.WriteEmptyRepo(context.Background(), "master", "Bill Billerson", "bigbillieb@fake.horse")

	if err != nil {
		t.Fatal("Unexpected error creating empty repo", err)
	}

	cs, _ := NewCommitSpec("master")
	optCmt, err := ddb.Resolve(context.Background(), cs, nil)
	if err != nil {
		t.Fatal("Could not find commit")
	}
	commit, ok := optCmt.ToCommit()
	assert.True(t, ok)

	h, err := commit.HashOf()
	assert.NoError(t, err)
	cs2, _ := NewCommitSpec(h.String())
	_, err = ddb.Resolve(context.Background(), cs2, nil)

	if err != nil {
		t.Fatal("Failed to get commit by hash")
	}
}

func TestLoadNonExistentLocalFSRepo(t *testing.T) {
	_, err := test.ChangeToTestDir("TestLoadRepo")

	if err != nil {
		panic("Couldn't change the working directory to the test directory.")
	}

	ddb, err := LoadDoltDB(context.Background(), types.Format_Default, LocalDirDoltDB, filesys.LocalFS)
	assert.Nil(t, ddb, "Should return nil when loading a non-existent data dir")
	assert.Error(t, err, "Should see an error here")
}

func TestLoadBadLocalFSRepo(t *testing.T) {
	testDir, err := test.ChangeToTestDir("TestLoadRepo")

	if err != nil {
		panic("Couldn't change the working directory to the test directory.")
	}

	contents := []byte("not a directory")
	os.WriteFile(filepath.Join(testDir, dbfactory.DoltDataDir), contents, 0644)

	ddb, err := LoadDoltDB(context.Background(), types.Format_Default, LocalDirDoltDB, filesys.LocalFS)
	assert.Nil(t, ddb, "Should return nil when loading a non-directory data dir file")
	assert.Error(t, err, "Should see an error here")
}

func TestLDNoms(t *testing.T) {
	testDir, err := test.ChangeToTestDir("TestLoadRepo")

	if err != nil {
		panic("Couldn't change the working directory to the test directory.")
	}

	committerName := "Bill Billerson"
	committerEmail := "bigbillieb@fake.horse"

	// Create an empty repo in a temp dir on the filesys
	{
		err := filesys.LocalFS.MkDirs(filepath.Join(testDir, dbfactory.DoltDataDir))

		if err != nil {
			t.Fatal("Failed to create noms directory")
		}

		ddb, _ := LoadDoltDB(context.Background(), types.Format_Default, LocalDirDoltDB, filesys.LocalFS)
		err = ddb.WriteEmptyRepo(context.Background(), "master", committerName, committerEmail)

		if err != nil {
			t.Fatal("Unexpected error creating empty repo", err)
		}
	}

	//read the empty repo back and add a new table.  Write the value, but don't commit
	var valHash hash.Hash
	var tbl *Table
	{
		ddb, _ := LoadDoltDB(context.Background(), types.Format_Default, LocalDirDoltDB, filesys.LocalFS)
		cs, _ := NewCommitSpec("master")
		optCmt, err := ddb.Resolve(context.Background(), cs, nil)
		if err != nil {
			t.Fatal("Couldn't find commit")
		}
		commit, ok := optCmt.ToCommit()
		assert.True(t, ok)

		meta, err := commit.GetCommitMeta(context.Background())
		assert.NoError(t, err)

		if meta.Name != committerName || meta.Email != committerEmail {
			t.Error("Unexpected metadata")
		}

		root, err := commit.GetRootValue(context.Background())

		assert.NoError(t, err)

		names, err := root.GetTableNames(context.Background(), DefaultSchemaName)
		assert.NoError(t, err)
		if len(names) != 0 {
			t.Fatal("There should be no tables in empty db")
		}

		ctx := context.Background()
		tSchema := createTestSchema(t)
		rowData, err := durable.NewEmptyPrimaryIndex(ctx, ddb.vrw, ddb.ns, tSchema)
		if err != nil {
			t.Fatal("Failed to create new empty index")
		}

		tbl, err = CreateTestTable(ddb.vrw, ddb.ns, tSchema, rowData)
		if err != nil {
			t.Fatal("Failed to create test table with data")
		}

		root, err = root.PutTable(context.Background(), TableName{Name: "test"}, tbl)
		assert.NoError(t, err)

		root, valHash, err = ddb.WriteRootValue(context.Background(), root)
		assert.NoError(t, err)

		meta, err = datas.NewCommitMeta(committerName, committerEmail, "Sample data")
		if err != nil {
			t.Error("Failed to commit")
		}

		commit, err = ddb.Commit(context.Background(), valHash, ref.NewBranchRef("master"), meta)
		if err != nil {
			t.Error("Failed to commit")
		}

		rootHash, err := ddb.NomsRoot(context.Background())
		if err != nil {
			t.Error("Failed to get root hash")
		}
		branches, err := ddb.GetBranchesByRootHash(context.Background(), rootHash)
		if err != nil {
			t.Error("Failed to get branches by root hash")
		}
		assert.Equal(t, len(branches), 1)
		assert.Equal(t, branches[0].Ref.GetPath(), "master")

		numParents := commit.NumParents()
		assert.NoError(t, err)

		if numParents != 1 {
			t.Error("Unexpected ancestry")
		}

		root, err = commit.GetRootValue(context.Background())
		assert.NoError(t, err)

		readTable, ok, err := root.GetTable(context.Background(), TableName{Name: "test"})
		assert.NoError(t, err)

		if !ok {
			t.Error("Could not retrieve test table")
		}

		ts, err := tbl.GetSchema(context.Background())
		require.NoError(t, err)

		rs, err := readTable.GetSchema(context.Background())
		require.NoError(t, err)

		if !schema.SchemasAreEqual(ts, rs) {
			t.Error("Unexpected schema")
		}
	}
}
