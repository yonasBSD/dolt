// Copyright 2022 Dolthub, Inc.
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

package docscmds

import (
	"context"
	"fmt"
	"io"

	"github.com/dolthub/go-mysql-server/sql"

	"github.com/dolthub/dolt/go/cmd/dolt/cli"
	"github.com/dolthub/dolt/go/cmd/dolt/commands"
	"github.com/dolthub/dolt/go/cmd/dolt/commands/engine"
	"github.com/dolthub/dolt/go/cmd/dolt/errhand"
	"github.com/dolthub/dolt/go/libraries/doltcore/doltdb"
	"github.com/dolthub/dolt/go/libraries/doltcore/env"
	"github.com/dolthub/dolt/go/libraries/utils/argparser"
)

var writeDocs = cli.CommandDocumentationContent{
	ShortDesc: "Reads Dolt docs to stdout",
	LongDesc:  "Reads Dolt docs to stdout",
	Synopsis: []string{
		"[{{.LessThan}}dolt doc{{.GreaterThan}}]",
	},
}

type WriteCmd struct{}

// Name implements cli.Command.
func (cmd WriteCmd) Name() string {
	return "write"
}

// Description implements cli.Command.
func (cmd WriteCmd) Description() string {
	return writeDocs.ShortDesc
}

// RequiresRepo implements cli.Command.
func (cmd WriteCmd) RequiresRepo() bool {
	return true
}

// Docs implements cli.Command.
func (cmd WriteCmd) Docs() *cli.CommandDocumentation {
	ap := cmd.ArgParser()
	return cli.NewCommandDocumentation(writeDocs, ap)
}

// ArgParser implements cli.Command.
func (cmd WriteCmd) ArgParser() *argparser.ArgParser {
	ap := argparser.NewArgParser()
	ap.ArgListHelp = append(ap.ArgListHelp, [2]string{"doc", "Dolt doc to be read."})
	return ap
}

// Exec implements cli.Command.
func (cmd WriteCmd) Exec(ctx context.Context, commandStr string, args []string, dEnv *env.DoltEnv) int {
	ap := cmd.ArgParser()
	help, usage := cli.HelpAndUsagePrinters(cli.CommandDocsForCommandString(commandStr, writeDocs, ap))
	apr := cli.ParseArgsOrDie(ap, args, help)

	if apr.NArg() != 1 {
		verr := errhand.BuildDError("dolt docs read takes exactly one argument").Build()
		return commands.HandleVErrAndExitCode(verr, usage)
	}

	var verr errhand.VerboseError
	if err := writeDoltDoc(ctx, dEnv, apr.Arg(0)); err != nil {
		verr = errhand.VerboseErrorFromError(err)
	}

	return commands.HandleVErrAndExitCode(verr, usage)
}

func writeDoltDoc(ctx context.Context, dEnv *env.DoltEnv, docName string) error {
	eng, err := engine.NewSqlEngineForEnv(ctx, dEnv)
	if err != nil {
		return err
	}

	doc, err := readDocFromTable(ctx, eng, docName)
	if err != nil {
		return err
	}

	cli.Print(doc)
	return nil
}

const (
	readDocTemplate = "SELECT " + doltdb.DocTextColumnName + " " +
		"FROM dolt_docs %s WHERE " + doltdb.DocPkColumnName + " = '%s'"
)

func readDocFromTable(ctx context.Context, eng *engine.SqlEngine, docName string) (string, error) {
	return readDocFromTableAsOf(ctx, eng, docName, "")
}

func readDocFromTableAsOf(ctx context.Context, eng *engine.SqlEngine, docName, asOf string) (doc string, err error) {
	var (
		sctx *sql.Context
		iter sql.RowIter
		row  sql.Row
	)

	if asOf != "" {
		asOf = fmt.Sprintf("AS OF '%s'", asOf)
	}
	query := fmt.Sprintf(readDocTemplate, asOf, docName)

	sctx, err = eng.NewContext(ctx)
	if err != nil {
		return "", err
	}
	sctx.Session.SetClient(sql.Client{User: "root", Address: "%", Capabilities: 0})

	_, iter, err = eng.Query(sctx, query)
	if err != nil {
		return "", err
	}
	defer func() {
		if cerr := iter.Close(sctx); err == nil {
			err = cerr
		}
	}()

	row, err = iter.Next(sctx)
	if err == io.EOF {
		// doc does not exist
		return "", nil
	}
	if err != nil {
		return "", err
	}

	doc = row[0].(string)

	_, eof := iter.Next(sctx)
	assertTrue(eof == io.EOF)
	return
}

func assertTrue(b bool) {
	if !b {
		panic("expected true")
	}
}
