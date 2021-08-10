[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [UI](ui.md)

# Interface: UI

## Hierarchy

* **UI**

## Index

### Methods

* [listenGlobalSearch](ui.md#listenglobalsearch)
* [listenSearchbar](ui.md#listensearchbar)

## Methods

###  listenGlobalSearch

▸ **listenGlobalSearch**(`cb`: function): *Promise‹[Cancellable](cancellable.md)›*

*Defined in [ui/index.ts:20](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/ui/index.ts#L20)*

Creates a stream receiving updates whenever the user enters a search in the Olive Helps Global Search.

**Parameters:**

▪ **cb**: *function*

The callback function called when an update to the global searchbar occurs.

▸ (`val`: string): *void*

**Parameters:**

Name | Type |
------ | ------ |
`val` | string |

**Returns:** *Promise‹[Cancellable](cancellable.md)›*

___

###  listenSearchbar

▸ **listenSearchbar**(`cb`: function): *Promise‹[Cancellable](cancellable.md)›*

*Defined in [ui/index.ts:13](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/ui/index.ts#L13)*

Creates a stream receiving updates whenever the user enters a search in the Olive Helps Searchbar.

**Parameters:**

▪ **cb**: *function*

The callback function called when an update to the searchbar occurs.

▸ (`val`: string): *void*

**Parameters:**

Name | Type |
------ | ------ |
`val` | string |

**Returns:** *Promise‹[Cancellable](cancellable.md)›*
