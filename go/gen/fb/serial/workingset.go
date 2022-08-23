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

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package serial

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type WorkingSet struct {
	_tab flatbuffers.Table
}

func InitWorkingSetRoot(o *WorkingSet, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	o.Init(buf, n+offset)
	if WorkingSetNumFields < o.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func TryGetRootAsWorkingSet(buf []byte, offset flatbuffers.UOffsetT) (*WorkingSet, error) {
	x := &WorkingSet{}
	return x, InitWorkingSetRoot(x, buf, offset)
}

func GetRootAsWorkingSet(buf []byte, offset flatbuffers.UOffsetT) *WorkingSet {
	x := &WorkingSet{}
	InitWorkingSetRoot(x, buf, offset)
	return x
}

func TryGetSizePrefixedRootAsWorkingSet(buf []byte, offset flatbuffers.UOffsetT) (*WorkingSet, error) {
	x := &WorkingSet{}
	return x, InitWorkingSetRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func GetSizePrefixedRootAsWorkingSet(buf []byte, offset flatbuffers.UOffsetT) *WorkingSet {
	x := &WorkingSet{}
	InitWorkingSetRoot(x, buf, offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *WorkingSet) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *WorkingSet) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *WorkingSet) WorkingRootAddr(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *WorkingSet) WorkingRootAddrLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *WorkingSet) WorkingRootAddrBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *WorkingSet) MutateWorkingRootAddr(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *WorkingSet) StagedRootAddr(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *WorkingSet) StagedRootAddrLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *WorkingSet) StagedRootAddrBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *WorkingSet) MutateStagedRootAddr(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *WorkingSet) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *WorkingSet) Email() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *WorkingSet) Desc() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *WorkingSet) TimestampMillis() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *WorkingSet) MutateTimestampMillis(n uint64) bool {
	return rcv._tab.MutateUint64Slot(14, n)
}

func (rcv *WorkingSet) MergeState(obj *MergeState) *MergeState {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(MergeState)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *WorkingSet) TryMergeState(obj *MergeState) (*MergeState, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(MergeState)
		}
		obj.Init(rcv._tab.Bytes, x)
		if MergeStateNumFields < obj.Table().NumFields() {
			return nil, flatbuffers.ErrTableHasUnknownFields
		}
		return obj, nil
	}
	return nil, nil
}

const WorkingSetNumFields = 7

func WorkingSetStart(builder *flatbuffers.Builder) {
	builder.StartObject(WorkingSetNumFields)
}
func WorkingSetAddWorkingRootAddr(builder *flatbuffers.Builder, workingRootAddr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(workingRootAddr), 0)
}
func WorkingSetStartWorkingRootAddrVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func WorkingSetAddStagedRootAddr(builder *flatbuffers.Builder, stagedRootAddr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(stagedRootAddr), 0)
}
func WorkingSetStartStagedRootAddrVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func WorkingSetAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(name), 0)
}
func WorkingSetAddEmail(builder *flatbuffers.Builder, email flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(email), 0)
}
func WorkingSetAddDesc(builder *flatbuffers.Builder, desc flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(desc), 0)
}
func WorkingSetAddTimestampMillis(builder *flatbuffers.Builder, timestampMillis uint64) {
	builder.PrependUint64Slot(5, timestampMillis, 0)
}
func WorkingSetAddMergeState(builder *flatbuffers.Builder, mergeState flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(mergeState), 0)
}
func WorkingSetEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type MergeState struct {
	_tab flatbuffers.Table
}

func InitMergeStateRoot(o *MergeState, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	o.Init(buf, n+offset)
	if MergeStateNumFields < o.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func TryGetRootAsMergeState(buf []byte, offset flatbuffers.UOffsetT) (*MergeState, error) {
	x := &MergeState{}
	return x, InitMergeStateRoot(x, buf, offset)
}

func GetRootAsMergeState(buf []byte, offset flatbuffers.UOffsetT) *MergeState {
	x := &MergeState{}
	InitMergeStateRoot(x, buf, offset)
	return x
}

func TryGetSizePrefixedRootAsMergeState(buf []byte, offset flatbuffers.UOffsetT) (*MergeState, error) {
	x := &MergeState{}
	return x, InitMergeStateRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func GetSizePrefixedRootAsMergeState(buf []byte, offset flatbuffers.UOffsetT) *MergeState {
	x := &MergeState{}
	InitMergeStateRoot(x, buf, offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MergeState) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MergeState) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MergeState) PreWorkingRootAddr(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *MergeState) PreWorkingRootAddrLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MergeState) PreWorkingRootAddrBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *MergeState) MutatePreWorkingRootAddr(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *MergeState) FromCommitAddr(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *MergeState) FromCommitAddrLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MergeState) FromCommitAddrBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *MergeState) MutateFromCommitAddr(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

const MergeStateNumFields = 2

func MergeStateStart(builder *flatbuffers.Builder) {
	builder.StartObject(MergeStateNumFields)
}
func MergeStateAddPreWorkingRootAddr(builder *flatbuffers.Builder, preWorkingRootAddr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(preWorkingRootAddr), 0)
}
func MergeStateStartPreWorkingRootAddrVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func MergeStateAddFromCommitAddr(builder *flatbuffers.Builder, fromCommitAddr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(fromCommitAddr), 0)
}
func MergeStateStartFromCommitAddrVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func MergeStateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
