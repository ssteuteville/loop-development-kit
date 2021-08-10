[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [Window](window.md)

# Interface: Window

## Hierarchy

* **Window**

## Index

### Methods

* [activeWindow](window.md#activewindow)
* [all](window.md#all)
* [listenActiveWindow](window.md#listenactivewindow)
* [listenAll](window.md#listenall)

## Methods

###  activeWindow

▸ **activeWindow**(): *Promise‹[WindowInfo](windowinfo.md)›*

*Defined in [window/index.ts:42](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/window/index.ts#L42)*

Get the currently focused window and it's data.

**Returns:** *Promise‹[WindowInfo](windowinfo.md)›*

A promise containing active window info.

___

###  all

▸ **all**(): *Promise‹[WindowInfo](windowinfo.md)[]›*

*Defined in [window/index.ts:56](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/window/index.ts#L56)*

Get a list of all the windows and their information.

**Returns:** *Promise‹[WindowInfo](windowinfo.md)[]›*

A promise containing all window info.

___

###  listenActiveWindow

▸ **listenActiveWindow**(`callback`: function): *Promise‹[Cancellable](cancellable.md)›*

*Defined in [window/index.ts:49](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/window/index.ts#L49)*

Receive notifications whenever the currently focused window changes.

**Parameters:**

▪ **callback**: *function*

A function called when active window changes.

▸ (`windowInfo`: [WindowInfo](windowinfo.md)): *void*

**Parameters:**

Name | Type |
------ | ------ |
`windowInfo` | [WindowInfo](windowinfo.md) |

**Returns:** *Promise‹[Cancellable](cancellable.md)›*

___

###  listenAll

▸ **listenAll**(`callback`: function): *Promise‹[Cancellable](cancellable.md)›*

*Defined in [window/index.ts:63](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/window/index.ts#L63)*

Receive a notification whenever a window is opened, closed, focused, unfocused, moved, resized, or its title changes. A window that is opened with focus will generate an Opened event and a Focused event.

**Parameters:**

▪ **callback**: *function*

A function called when any window changes.

▸ (`windowEvent`: [WindowEvent](windowevent.md)): *void*

**Parameters:**

Name | Type |
------ | ------ |
`windowEvent` | [WindowEvent](windowevent.md) |

**Returns:** *Promise‹[Cancellable](cancellable.md)›*
