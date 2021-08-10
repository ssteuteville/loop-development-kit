[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [Socket](socket.md)

# Interface: Socket

Object to communicate with the websocket

## Hierarchy

* **Socket**

## Index

### Properties

* [setMessageHandler](socket.md#setmessagehandler)

### Methods

* [close](socket.md#close)
* [setCloseHandler](socket.md#setclosehandler)
* [writeMessage](socket.md#writemessage)

## Properties

###  setMessageHandler

• **setMessageHandler**: *function*

*Defined in [network/types.ts:88](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/types.ts#L88)*

Allows to listen for a websocket message (there must be only one listener registered per socket for messages to be fully received)

**`param`** Receives text or data message from websocket and error if occurs (error could be returned if reading from the closed connection)

#### Type declaration:

▸ (`handler`: function): *Promise‹[Cancellable](cancellable.md)›*

**Parameters:**

▪ **handler**: *function*

▸ (`error`: Error | undefined, `message`: string | Uint8Array): *void*

**Parameters:**

Name | Type |
------ | ------ |
`error` | Error &#124; undefined |
`message` | string &#124; Uint8Array |

## Methods

###  close

▸ **close**(): *Promise‹void›*

*Defined in [network/types.ts:83](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/types.ts#L83)*

Closes websocket

**Returns:** *Promise‹void›*

___

###  setCloseHandler

▸ **setCloseHandler**(`handler`: function): *Promise‹void›*

*Defined in [network/types.ts:95](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/types.ts#L95)*

Allows to provide handler when websocket closing

**Parameters:**

▪ **handler**: *function*

Receives code status and text received from the peer

▸ (`error`: Error | undefined, `code`: number, `text`: string): *void*

**Parameters:**

Name | Type |
------ | ------ |
`error` | Error &#124; undefined |
`code` | number |
`text` | string |

**Returns:** *Promise‹void›*

___

###  writeMessage

▸ **writeMessage**(`message`: string | Uint8Array): *Promise‹void›*

*Defined in [network/types.ts:79](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/types.ts#L79)*

Writes message to a websocket

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`message` | string &#124; Uint8Array | Text or data message  |

**Returns:** *Promise‹void›*
