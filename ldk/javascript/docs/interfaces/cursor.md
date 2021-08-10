[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [Cursor](cursor.md)

# Interface: Cursor

The CursorService provides access to the cursor position.

## Hierarchy

* **Cursor**

## Index

### Methods

* [listenPosition](cursor.md#listenposition)
* [position](cursor.md#position)

## Methods

###  listenPosition

▸ **listenPosition**(`callback`: function): *Promise‹[Cancellable](cancellable.md)›*

*Defined in [cursor/index.ts:26](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/cursor/index.ts#L26)*

Starts listening to changes to the clipboard.

**Parameters:**

▪ **callback**: *function*

The callback function called when the function changes.

▸ (`pos`: [Position](position.md)): *void*

**Parameters:**

Name | Type |
------ | ------ |
`pos` | [Position](position.md) |

**Returns:** *Promise‹[Cancellable](cancellable.md)›*

___

###  position

▸ **position**(): *Promise‹[Position](position.md)›*

*Defined in [cursor/index.ts:19](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/cursor/index.ts#L19)*

**Returns:** *Promise‹[Position](position.md)›*

Promise resolving with the cursor position.
