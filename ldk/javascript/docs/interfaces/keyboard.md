[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [Keyboard](keyboard.md)

# Interface: Keyboard

## Hierarchy

* **Keyboard**

## Index

### Methods

* [listenCharacter](keyboard.md#listencharacter)
* [listenHotkey](keyboard.md#listenhotkey)
* [listenText](keyboard.md#listentext)

## Methods

###  listenCharacter

▸ **listenCharacter**(`callback`: function): *Promise‹[Cancellable](cancellable.md)›*

*Defined in [keyboard/index.ts:41](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/keyboard/index.ts#L41)*

Calls callback function when a character is detected from the clipboard.

**Parameters:**

▪ **callback**: *function*

The callback function called when a character is detected from the clipboard.

▸ (`char`: string): *void*

**Parameters:**

Name | Type |
------ | ------ |
`char` | string |

**Returns:** *Promise‹[Cancellable](cancellable.md)›*

___

###  listenHotkey

▸ **listenHotkey**(`hotkey`: [Hotkey](hotkey.md), `callback`: function): *Promise‹[Cancellable](cancellable.md)›*

*Defined in [keyboard/index.ts:27](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/keyboard/index.ts#L27)*

Calls callback function when the specified hotkey is pressed or released.

**Parameters:**

▪ **hotkey**: *[Hotkey](hotkey.md)*

The hotkey to monitor to initiate callback.

▪ **callback**: *function*

The callback function called when the specified hotkey is pressed or released.

▸ (`pressed`: boolean): *void*

**Parameters:**

Name | Type |
------ | ------ |
`pressed` | boolean |

**Returns:** *Promise‹[Cancellable](cancellable.md)›*

___

###  listenText

▸ **listenText**(`callback`: function): *Promise‹[Cancellable](cancellable.md)›*

*Defined in [keyboard/index.ts:34](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/keyboard/index.ts#L34)*

Calls callback function when text is detected from the clipboard.

**Parameters:**

▪ **callback**: *function*

The callback function called when text is detected from the clipboard.

▸ (`text`: string): *void*

**Parameters:**

Name | Type |
------ | ------ |
`text` | string |

**Returns:** *Promise‹[Cancellable](cancellable.md)›*
