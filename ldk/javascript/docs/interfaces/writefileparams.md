[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [WriteFileParams](writefileparams.md)

# Interface: WriteFileParams

## Hierarchy

* **WriteFileParams**

## Index

### Properties

* [data](writefileparams.md#data)
* [path](writefileparams.md#path)
* [writeMode](writefileparams.md#writemode)
* [writeOperation](writefileparams.md#writeoperation)

## Properties

###  data

• **data**: *string | Uint8Array*

*Defined in [filesystem/types.ts:72](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/filesystem/types.ts#L72)*

file data

___

###  path

• **path**: *string*

*Defined in [filesystem/types.ts:68](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/filesystem/types.ts#L68)*

path to the file location to be written

___

###  writeMode

• **writeMode**: *[WriteMode](../globals.md#writemode)*

*Defined in [filesystem/types.ts:80](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/filesystem/types.ts#L80)*

file mode and permission bits. Should be provided as octal value (ex. 0o755, 0o777, ...)

___

###  writeOperation

• **writeOperation**: *[WriteOperation](../enums/writeoperation.md)*

*Defined in [filesystem/types.ts:76](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/filesystem/types.ts#L76)*

indicates if file should be overwritten or appended with the provided data
