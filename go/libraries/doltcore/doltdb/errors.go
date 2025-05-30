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
	"bytes"
	"errors"
	"fmt"
)

var ErrInvBranchName = errors.New("not a valid user branch name")
var ErrInvWorkspaceName = errors.New("not a valid user workspace name")
var ErrInvTagName = errors.New("not a valid user tag name")
var ErrInvTableName = errors.New("not a valid table name")
var ErrInvHash = errors.New("not a valid hash")
var ErrInvalidAncestorSpec = errors.New("invalid ancestor spec")
var ErrInvalidBranchOrHash = errors.New("string is not a valid branch or hash")
var ErrInvalidHash = errors.New("string is not a valid hash")

var ErrFoundHashNotACommit = errors.New("the value retrieved for this hash is not a commit")
var ErrHashNotFound = errors.New("could not find a value for this hash")
var ErrBranchNotFound = errors.New("branch not found")
var ErrTagNotFound = errors.New("tag not found")
var ErrTupleNotFound = errors.New("tuple not found")
var ErrWorkingSetNotFound = errors.New("working set not found")
var ErrWorkspaceNotFound = errors.New("workspace not found")
var ErrTableNotFound = errors.New("table not found")
var ErrTableExists = errors.New("table already exists")
var ErrAlreadyOnBranch = errors.New("Already on branch")
var ErrAlreadyOnWorkspace = errors.New("Already on workspace")

var ErrNomsIO = errors.New("error reading from or writing to noms")

// ErrUpToDate is returned when a merge is up-to-date. Not actually an error, and we do use this message in non-error contexts.
var ErrUpToDate = errors.New("Everything up-to-date")
var ErrIsAhead = errors.New("cannot fast forward from a to b. a is ahead of b already")
var ErrIsBehind = errors.New("cannot reverse from b to a. b is a is behind a already")

var ErrUnresolvedConflictsOrViolations = errors.New("merge has unresolved conflicts or constraint violations")
var ErrMergeActive = errors.New("merging is not possible because you have not committed an active merge")

var ErrOperationNotSupportedInDetachedHead = errors.New("this operation is not supported while in a detached head state")

type ErrClientOutOfDate struct {
	RepoVer   FeatureVersion
	ClientVer FeatureVersion
}

func (e ErrClientOutOfDate) Error() string {
	return fmt.Sprintf(`client (version: %d) is out of date and must upgrade to read this repo (version: %d).
	visit https://github.com/dolthub/dolt/releases/latest/`, e.ClientVer, e.RepoVer)
}

func IsInvalidFormatErr(err error) bool {
	switch err {
	case ErrInvBranchName, ErrInvTableName, ErrInvHash, ErrInvalidAncestorSpec, ErrInvalidBranchOrHash:
		return true
	default:
		return false
	}
}

func IsNotFoundErr(err error) bool {
	return errors.Is(err, ErrHashNotFound) ||
		errors.Is(err, ErrBranchNotFound) ||
		errors.Is(err, ErrTableNotFound)
}

func IsNotACommit(err error) bool {
	return errors.Is(err, ErrHashNotFound) ||
		errors.Is(err, ErrBranchNotFound) ||
		errors.Is(err, ErrFoundHashNotACommit)
}

type RootType int

const (
	WorkingRoot RootType = iota
	StagedRoot
	CommitRoot
	HeadRoot
	InvalidRoot
)

func (rt RootType) String() string {
	switch rt {
	case WorkingRoot:
		return "working root"
	case StagedRoot:
		return "staged root"
	case CommitRoot:
		return "root value for commit"
	case HeadRoot:
		return "HEAD commit root value"
	}

	return "unknown"
}

type RootTypeSet map[RootType]struct{}

func NewRootTypeSet(rts ...RootType) RootTypeSet {
	mp := make(map[RootType]struct{})

	for _, rt := range rts {
		mp[rt] = struct{}{}
	}

	return RootTypeSet(mp)
}

func (rts RootTypeSet) Contains(rt RootType) bool {
	_, ok := rts[rt]
	return ok
}

func (rts RootTypeSet) First(rtList []RootType) RootType {
	for _, rt := range rtList {
		if _, ok := rts[rt]; ok {
			return rt
		}
	}

	return InvalidRoot
}

func (rts RootTypeSet) IsEmpty() bool {
	return len(rts) == 0
}

type RootValueUnreadable struct {
	RootType RootType
	Cause    error
}

func (rvu RootValueUnreadable) Error() string {
	rs, es := rvu.RootType.String(), rvu.Cause.Error()
	return fmt.Sprintf("unable to read %s: %s", rs, es)
}

func IsRootValUnreachable(err error) bool {
	_, ok := err.(RootValueUnreadable)
	return ok
}

func GetUnreachableRootType(err error) RootType {
	rvu, ok := err.(RootValueUnreadable)

	if !ok {
		panic("Must validate with IsRootValUnreachable before calling GetUnreachableRootType")
	}

	return rvu.RootType
}

func GetUnreachableRootCause(err error) error {
	rvu, ok := err.(RootValueUnreadable)

	if !ok {
		panic("Must validate with IsRootValUnreachable before calling GetUnreachableRootCause")
	}

	return rvu.Cause
}

// DoltIgnoreConflictError is an error that is returned when the user attempts to stage a table that matches conflicting dolt_ignore patterns
type DoltIgnoreConflictError struct {
	Table         TableName
	TruePatterns  []string
	FalsePatterns []string
}

func (dc DoltIgnoreConflictError) Error() string {
	var buffer bytes.Buffer
	buffer.WriteString("the table ")
	buffer.WriteString(dc.Table.Name)
	buffer.WriteString(" matches conflicting patterns in dolt_ignore:")

	for _, pattern := range dc.TruePatterns {
		buffer.WriteString("\nignored:     ")
		buffer.WriteString(pattern)
	}

	for _, pattern := range dc.FalsePatterns {
		buffer.WriteString("\nnot ignored: ")
		buffer.WriteString(pattern)
	}

	return buffer.String()
}

func AsDoltIgnoreInConflict(err error) *DoltIgnoreConflictError {
	di, ok := err.(DoltIgnoreConflictError)
	if ok {
		return &di
	}
	return nil
}
