[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [Vault](vault.md)

# Interface: Vault

## Hierarchy

* **Vault**

## Index

### Methods

* [exists](vault.md#exists)
* [read](vault.md#read)
* [remove](vault.md#remove)
* [write](vault.md#write)

## Methods

###  exists

▸ **exists**(`key`: string): *Promise‹boolean›*

*Defined in [vault/index.ts:21](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/vault/index.ts#L21)*

Returns true if the specified vault entry exists.

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`key` | string | The key for the item to check existence. |

**Returns:** *Promise‹boolean›*

True if the specified vault entry exists.

___

###  read

▸ **read**(`key`: string): *Promise‹string›*

*Defined in [vault/index.ts:29](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/vault/index.ts#L29)*

Returns the stored value for the specified vault entry.

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`key` | string | The key for the item to read. |

**Returns:** *Promise‹string›*

Stored value for the specified vault entry.

___

###  remove

▸ **remove**(`key`: string): *Promise‹void›*

*Defined in [vault/index.ts:13](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/vault/index.ts#L13)*

Removes an entry from the vault.

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`key` | string | The key for item to remove. |

**Returns:** *Promise‹void›*

Status of the remove operation.

___

###  write

▸ **write**(`key`: string, `value`: string): *Promise‹void›*

*Defined in [vault/index.ts:38](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/vault/index.ts#L38)*

Writes the provided value to the specified vault entry.

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`key` | string | The key for the item to write. |
`value` | string | The value to write. |

**Returns:** *Promise‹void›*

The status of the write operation.
