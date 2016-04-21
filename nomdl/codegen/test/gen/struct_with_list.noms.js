// This file was generated by nomdl/codegen.
// @flow
/* eslint-disable */

import {
  Field as _Field,
  Kind as _Kind,
  Package as _Package,
  boolType as _boolType,
  createStructClass as _createStructClass,
  int64Type as _int64Type,
  makeCompoundType as _makeCompoundType,
  makeListType as _makeListType,
  makeStructType as _makeStructType,
  makeType as _makeType,
  newList as _newList,
  registerPackage as _registerPackage,
  stringType as _stringType,
  uint8Type as _uint8Type,
} from '@attic/noms';
import type {
  NomsList as _NomsList,
  Struct as _Struct,
  int64 as _int64,
  uint8 as _uint8,
} from '@attic/noms';

const _pkg = new _Package([
  _makeStructType('StructWithList',
    [
      new _Field('l', _makeCompoundType(_Kind.List, _uint8Type), false),
      new _Field('b', _boolType, false),
      new _Field('s', _stringType, false),
      new _Field('i', _int64Type, false),
    ],
    [

    ]
  ),
], [
]);
_registerPackage(_pkg);
const StructWithList$type = _makeType(_pkg.ref, 0);
const StructWithList$typeDef = _pkg.types[0];


type StructWithList$Data = {
  l: _NomsList<_uint8>;
  b: boolean;
  s: string;
  i: _int64;
};

interface StructWithList$Interface extends _Struct {
  constructor(data: StructWithList$Data): void;
  l: _NomsList<_uint8>;  // readonly
  setL(value: _NomsList<_uint8>): StructWithList$Interface;
  b: boolean;  // readonly
  setB(value: boolean): StructWithList$Interface;
  s: string;  // readonly
  setS(value: string): StructWithList$Interface;
  i: _int64;  // readonly
  setI(value: _int64): StructWithList$Interface;
}

export const StructWithList: Class<StructWithList$Interface> = _createStructClass(StructWithList$type, StructWithList$typeDef);

export function newListOfUint8(values: Array<_uint8>): Promise<_NomsList<_uint8>> {
  return _newList(values, _makeListType(_uint8Type));
}
