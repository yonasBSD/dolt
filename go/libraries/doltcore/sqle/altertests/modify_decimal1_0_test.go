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

package altertests

import (
	"testing"
)

func TestModifyTypeDecimal1(t *testing.T) {
	SkipByDefaultInCI(t)
	tests := []ModifyTypeTest{
		{
			"DECIMAL(1)",
			"TINYINT",
			`(0,"0")`,
			[]interface{}{int64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"SMALLINT",
			`(0,"0")`,
			[]interface{}{int64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"MEDIUMINT",
			`(0,"0")`,
			[]interface{}{int64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"INT",
			`(0,"0")`,
			[]interface{}{int64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIGINT",
			`(0,"0")`,
			[]interface{}{int64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"TINYINT UNSIGNED",
			`(0,"0")`,
			[]interface{}{uint64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"SMALLINT UNSIGNED",
			`(0,"0")`,
			[]interface{}{uint64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"MEDIUMINT UNSIGNED",
			`(0,"0")`,
			[]interface{}{uint64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"INT UNSIGNED",
			`(0,"0")`,
			[]interface{}{uint64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIGINT UNSIGNED",
			`(0,"0")`,
			[]interface{}{uint64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"FLOAT",
			`(0,"0")`,
			[]interface{}{float64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"DOUBLE",
			`(0,"0")`,
			[]interface{}{float64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(1,0)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(15,0)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(30,0)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(65,0)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(1,1)",
			`(0,"0")`,
			[]interface{}{"0.0"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(15,1)",
			`(0,"0")`,
			[]interface{}{"0.0"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(30,1)",
			`(0,"0")`,
			[]interface{}{"0.0"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(65,1)",
			`(0,"0")`,
			[]interface{}{"0.0"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(15,15)",
			`(0,"0")`,
			[]interface{}{"0.000000000000000"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(30,15)",
			`(0,"0")`,
			[]interface{}{"0.000000000000000"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(65,15)",
			`(0,"0")`,
			[]interface{}{"0.000000000000000"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(30,30)",
			`(0,"0")`,
			[]interface{}{"0.000000000000000000000000000000"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(65,30)",
			`(0,"0")`,
			[]interface{}{"0.000000000000000000000000000000"},
			false,
		},
		{
			"DECIMAL(1)",
			"BIT(1)",
			`(0,"0")`,
			[]interface{}{uint64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIT(8)",
			`(0,"0")`,
			[]interface{}{uint64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIT(16)",
			`(0,"0")`,
			[]interface{}{uint64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIT(24)",
			`(0,"0")`,
			[]interface{}{uint64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIT(32)",
			`(0,"0")`,
			[]interface{}{uint64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIT(48)",
			`(0,"0")`,
			[]interface{}{uint64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIT(64)",
			`(0,"0")`,
			[]interface{}{uint64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"TINYBLOB",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"BLOB",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"MEDIUMBLOB",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"LONGBLOB",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"TINYTEXT",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"TEXT",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"MEDIUMTEXT",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"LONGTEXT",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"CHAR(1)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"CHAR(10)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"CHAR(100)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"CHAR(255)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"BINARY(1)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"BINARY(10)",
			`(0,"0")`,
			[]interface{}{"0\x00\x00\x00\x00\x00\x00\x00\x00\x00"},
			false,
		},
		{
			"DECIMAL(1)",
			"BINARY(100)",
			`(0,"0")`,
			[]interface{}{"0\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"},
			false,
		},
		{
			"DECIMAL(1)",
			"BINARY(255)",
			`(0,"0")`,
			[]interface{}{"0\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(1)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(10)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(100)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(255)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(1023)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(4095)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(1)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(10)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(100)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(255)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(1023)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(4095)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(16383)",
			`(0,"0")`,
			[]interface{}{"0"},
			false,
		},
		{
			"DECIMAL(1)",
			"YEAR",
			`(0,"0")`,
			[]interface{}{int64(0)},
			false,
		},
		{
			"DECIMAL(1)",
			"DATE",
			`(0,"0")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"TIME",
			`(0,"0")`,
			[]interface{}{"00:00:00"},
			false,
		},
		{
			"DECIMAL(1)",
			"TIMESTAMP",
			`(0,"0")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"DATETIME",
			`(0,"0")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"ENUM('A')",
			`(0,"0")`,
			[]interface{}{""},
			false,
		},
		{
			"DECIMAL(1)",
			"ENUM('B')",
			`(0,"0")`,
			[]interface{}{""},
			false,
		},
		{
			"DECIMAL(1)",
			"ENUM('C')",
			`(0,"0")`,
			[]interface{}{""},
			false,
		},
		{
			"DECIMAL(1)",
			"ENUM('A','B')",
			`(0,"0")`,
			[]interface{}{""},
			false,
		},
		{
			"DECIMAL(1)",
			"ENUM('A','C')",
			`(0,"0")`,
			[]interface{}{""},
			false,
		},
		{
			"DECIMAL(1)",
			"ENUM('B','C')",
			`(0,"0")`,
			[]interface{}{""},
			false,
		},
		{
			"DECIMAL(1)",
			"ENUM('A','B','C')",
			`(0,"0")`,
			[]interface{}{""},
			false,
		},
		{
			"DECIMAL(1)",
			"ENUM('C','A','B')",
			`(0,"0")`,
			[]interface{}{""},
			false,
		},
		{
			"DECIMAL(1)",
			"SET('A')",
			`(0,"0")`,
			[]interface{}{""},
			false,
		},
		{
			"DECIMAL(1)",
			"SET('B')",
			`(0,"0")`,
			[]interface{}{""},
			false,
		},
		{
			"DECIMAL(1)",
			"SET('C')",
			`(0,"0")`,
			[]interface{}{""},
			false,
		},
		{
			"DECIMAL(1)",
			"SET('A','B')",
			`(0,"0")`,
			[]interface{}{""},
			false,
		},
		{
			"DECIMAL(1)",
			"SET('A','C')",
			`(0,"0")`,
			[]interface{}{""},
			false,
		},
		{
			"DECIMAL(1)",
			"SET('B','C')",
			`(0,"0")`,
			[]interface{}{""},
			false,
		},
		{
			"DECIMAL(1)",
			"SET('A','B','C')",
			`(0,"0")`,
			[]interface{}{""},
			false,
		},
		{
			"DECIMAL(1)",
			"SET('C','A','B')",
			`(0,"0")`,
			[]interface{}{""},
			false,
		},
		{
			"DECIMAL(1)",
			"TINYINT",
			`(0,"0"), (1,"1")`,
			[]interface{}{int64(0), int64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"SMALLINT",
			`(0,"0"), (1,"1")`,
			[]interface{}{int64(0), int64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"MEDIUMINT",
			`(0,"0"), (1,"1")`,
			[]interface{}{int64(0), int64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"INT",
			`(0,"0"), (1,"1")`,
			[]interface{}{int64(0), int64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIGINT",
			`(0,"0"), (1,"1")`,
			[]interface{}{int64(0), int64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"TINYINT UNSIGNED",
			`(0,"0"), (1,"1")`,
			[]interface{}{uint64(0), uint64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"SMALLINT UNSIGNED",
			`(0,"0"), (1,"1")`,
			[]interface{}{uint64(0), uint64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"MEDIUMINT UNSIGNED",
			`(0,"0"), (1,"1")`,
			[]interface{}{uint64(0), uint64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"INT UNSIGNED",
			`(0,"0"), (1,"1")`,
			[]interface{}{uint64(0), uint64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIGINT UNSIGNED",
			`(0,"0"), (1,"1")`,
			[]interface{}{uint64(0), uint64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"FLOAT",
			`(0,"0"), (1,"1")`,
			[]interface{}{float64(0), float64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"DOUBLE",
			`(0,"0"), (1,"1")`,
			[]interface{}{float64(0), float64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(1,0)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(15,0)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(30,0)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(65,0)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(1,1)",
			`(0,"0"), (1,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(15,1)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0.0", "1.0"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(30,1)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0.0", "1.0"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(65,1)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0.0", "1.0"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(15,15)",
			`(0,"0"), (1,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(30,15)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0.000000000000000", "1.000000000000000"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(65,15)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0.000000000000000", "1.000000000000000"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(30,30)",
			`(0,"0"), (1,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(65,30)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0.000000000000000000000000000000", "1.000000000000000000000000000000"},
			false,
		},
		{
			"DECIMAL(1)",
			"BIT(1)",
			`(0,"0"), (1,"1")`,
			[]interface{}{uint64(0), uint64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIT(8)",
			`(0,"0"), (1,"1")`,
			[]interface{}{uint64(0), uint64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIT(16)",
			`(0,"0"), (1,"1")`,
			[]interface{}{uint64(0), uint64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIT(24)",
			`(0,"0"), (1,"1")`,
			[]interface{}{uint64(0), uint64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIT(32)",
			`(0,"0"), (1,"1")`,
			[]interface{}{uint64(0), uint64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIT(48)",
			`(0,"0"), (1,"1")`,
			[]interface{}{uint64(0), uint64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIT(64)",
			`(0,"0"), (1,"1")`,
			[]interface{}{uint64(0), uint64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"TINYBLOB",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"BLOB",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"MEDIUMBLOB",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"LONGBLOB",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"TINYTEXT",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"TEXT",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"MEDIUMTEXT",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"LONGTEXT",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"CHAR(1)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"CHAR(10)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"CHAR(100)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"CHAR(255)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"BINARY(1)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"BINARY(10)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0\x00\x00\x00\x00\x00\x00\x00\x00\x00", "1\x00\x00\x00\x00\x00\x00\x00\x00\x00"},
			false,
		},
		{
			"DECIMAL(1)",
			"BINARY(100)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00", "1\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"},
			false,
		},
		{
			"DECIMAL(1)",
			"BINARY(255)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00", "1\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(1)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(10)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(100)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(255)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(1023)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(4095)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(1)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(10)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(100)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(255)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(1023)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(4095)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(16383)",
			`(0,"0"), (1,"1")`,
			[]interface{}{"0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"YEAR",
			`(0,"0"), (1,"1")`,
			[]interface{}{int64(0), int64(2001)},
			false,
		},
		{
			"DECIMAL(1)",
			"DATE",
			`(0,"0"), (1,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"TIME",
			`(0,"0"), (1,"1")`,
			[]interface{}{"00:00:00", "00:00:01"},
			false,
		},
		{
			"DECIMAL(1)",
			"TIMESTAMP",
			`(0,"0"), (1,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"DATETIME",
			`(0,"0"), (1,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"ENUM('A')",
			`(0,"0"), (1,"1")`,
			[]interface{}{"", "A"},
			false,
		},
		{
			"DECIMAL(1)",
			"ENUM('B')",
			`(0,"0"), (1,"1")`,
			[]interface{}{"", "B"},
			false,
		},
		{
			"DECIMAL(1)",
			"ENUM('C')",
			`(0,"0"), (1,"1")`,
			[]interface{}{"", "C"},
			false,
		},
		{
			"DECIMAL(1)",
			"ENUM('A','B')",
			`(0,"0"), (1,"1")`,
			[]interface{}{"", "A"},
			false,
		},
		{
			"DECIMAL(1)",
			"ENUM('A','C')",
			`(0,"0"), (1,"1")`,
			[]interface{}{"", "A"},
			false,
		},
		{
			"DECIMAL(1)",
			"ENUM('B','C')",
			`(0,"0"), (1,"1")`,
			[]interface{}{"", "B"},
			false,
		},
		{
			"DECIMAL(1)",
			"ENUM('A','B','C')",
			`(0,"0"), (1,"1")`,
			[]interface{}{"", "A"},
			false,
		},
		{
			"DECIMAL(1)",
			"ENUM('C','A','B')",
			`(0,"0"), (1,"1")`,
			[]interface{}{"", "C"},
			false,
		},
		{
			"DECIMAL(1)",
			"SET('A')",
			`(0,"0"), (1,"1")`,
			[]interface{}{"", "A"},
			false,
		},
		{
			"DECIMAL(1)",
			"SET('B')",
			`(0,"0"), (1,"1")`,
			[]interface{}{"", "B"},
			false,
		},
		{
			"DECIMAL(1)",
			"SET('C')",
			`(0,"0"), (1,"1")`,
			[]interface{}{"", "C"},
			false,
		},
		{
			"DECIMAL(1)",
			"SET('A','B')",
			`(0,"0"), (1,"1")`,
			[]interface{}{"", "A"},
			false,
		},
		{
			"DECIMAL(1)",
			"SET('A','C')",
			`(0,"0"), (1,"1")`,
			[]interface{}{"", "A"},
			false,
		},
		{
			"DECIMAL(1)",
			"SET('B','C')",
			`(0,"0"), (1,"1")`,
			[]interface{}{"", "B"},
			false,
		},
		{
			"DECIMAL(1)",
			"SET('A','B','C')",
			`(0,"0"), (1,"1")`,
			[]interface{}{"", "A"},
			false,
		},
		{
			"DECIMAL(1)",
			"SET('C','A','B')",
			`(0,"0"), (1,"1")`,
			[]interface{}{"", "C"},
			false,
		},
		{
			"DECIMAL(1)",
			"TINYINT",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{int64(-1), int64(0), int64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"SMALLINT",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{int64(-1), int64(0), int64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"MEDIUMINT",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{int64(-1), int64(0), int64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"INT",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{int64(-1), int64(0), int64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"BIGINT",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{int64(-1), int64(0), int64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"TINYINT UNSIGNED",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"SMALLINT UNSIGNED",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"MEDIUMINT UNSIGNED",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"INT UNSIGNED",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"BIGINT UNSIGNED",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"FLOAT",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{float64(-1), float64(0), float64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"DOUBLE",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{float64(-1), float64(0), float64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(1,0)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(15,0)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(30,0)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(65,0)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(1,1)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(15,1)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1.0", "0.0", "1.0"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(30,1)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1.0", "0.0", "1.0"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(65,1)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1.0", "0.0", "1.0"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(15,15)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(30,15)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1.000000000000000", "0.000000000000000", "1.000000000000000"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(65,15)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1.000000000000000", "0.000000000000000", "1.000000000000000"},
			false,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(30,30)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"DECIMAL(65,30)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1.000000000000000000000000000000", "0.000000000000000000000000000000", "1.000000000000000000000000000000"},
			false,
		},
		{
			"DECIMAL(1)",
			"BIT(1)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"BIT(8)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"BIT(16)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"BIT(24)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"BIT(32)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"BIT(48)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"BIT(64)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{uint64(18446744073709551615), uint64(0), uint64(1)},
			false,
		},
		{
			"DECIMAL(1)",
			"TINYBLOB",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"BLOB",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"MEDIUMBLOB",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"LONGBLOB",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"TINYTEXT",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"TEXT",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"MEDIUMTEXT",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"LONGTEXT",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"CHAR(1)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"CHAR(10)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"CHAR(100)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"CHAR(255)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"BINARY(1)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"BINARY(10)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1\x00\x00\x00\x00\x00\x00\x00\x00", "0\x00\x00\x00\x00\x00\x00\x00\x00\x00", "1\x00\x00\x00\x00\x00\x00\x00\x00\x00"},
			false,
		},
		{
			"DECIMAL(1)",
			"BINARY(100)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00", "0\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00", "1\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"},
			false,
		},
		{
			"DECIMAL(1)",
			"BINARY(255)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00", "0\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00", "1\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(1)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(10)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(100)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(255)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(1023)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARCHAR(4095)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(1)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(10)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(100)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(255)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(1023)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(4095)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"VARBINARY(16383)",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-1", "0", "1"},
			false,
		},
		{
			"DECIMAL(1)",
			"YEAR",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"DATE",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"TIME",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{"-00:00:01", "00:00:00", "00:00:01"},
			false,
		},
		{
			"DECIMAL(1)",
			"TIMESTAMP",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"DATETIME",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"ENUM('A')",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"ENUM('B')",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"ENUM('C')",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"ENUM('A','B')",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"ENUM('A','C')",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"ENUM('B','C')",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"ENUM('A','B','C')",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"ENUM('C','A','B')",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"SET('A')",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"SET('B')",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"SET('C')",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"SET('A','B')",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"SET('A','C')",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"SET('B','C')",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"SET('A','B','C')",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
		{
			"DECIMAL(1)",
			"SET('C','A','B')",
			`(0,"-1"), (1,"0"), (2,"1")`,
			[]interface{}{},
			true,
		},
	}

	RunModifyTypeTests(t, tests)
}
