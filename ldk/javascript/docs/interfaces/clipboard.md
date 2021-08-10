[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [Clipboard](clipboard.md)

# Interface: Clipboard

## Hierarchy

* **Clipboard**

## Index

### Methods

* [listen](clipboard.md#listen)
* [read](clipboard.md#read)
* [write](clipboard.md#write)

## Methods

###  listen

▸ **listen**(`includeOliveHelpsEvents`: boolean, `callback`: function): *Promise‹[Cancellable](cancellable.md)›*

*Defined in [clipboard/index.ts:26](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/clipboard/index.ts#L26)*

Starts listening to changes to the clipboard.

**Parameters:**

▪ **includeOliveHelpsEvents**: *boolean*

if passed in true, callback will be called while olive helps window is in focus

▪ **callback**: *function*

A function that's called whenever the clipboard's contents change.

▸ (`clipboardText`: string): *void*

**Parameters:**

Name | Type |
------ | ------ |
`clipboardText` | string |

**Returns:** *Promise‹[Cancellable](cancellable.md)›*

___

###  read

▸ **read**(): *Promise‹string›*

*Defined in [clipboard/index.ts:11](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/clipboard/index.ts#L11)*

**Returns:** *Promise‹string›*

A Promise resolving with the current contents of the clipboard.

___

###  write

▸ **write**(`text`: string): *Promise‹void›*

*Defined in [clipboard/index.ts:18](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/clipboard/index.ts#L18)*

Writes the provided text into the clipboard.

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`text` | string | A string to write to clipboard  |

**Returns:** *Promise‹void›*
