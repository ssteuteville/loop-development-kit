[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [OliveHelps](../modules/olivehelps.md) › [Socket](olivehelps.socket.md)

# Interface: Socket

## Hierarchy

* **Socket**

## Index

### Properties

* [listenMessage](olivehelps.socket.md#listenmessage)

### Methods

* [close](olivehelps.socket.md#close)
* [onCloseHandler](olivehelps.socket.md#onclosehandler)
* [writeMessage](olivehelps.socket.md#writemessage)

## Properties

###  listenMessage

• **listenMessage**: *function*

*Defined in [types/oliveHelps/index.d.ts:183](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/types/oliveHelps/index.d.ts#L183)*

#### Type declaration:

▸ (`callback`: function, `returnCb`: [ReturnCallback](../modules/olivehelps.md#returncallback)): *void*

**Parameters:**

▪ **callback**: *function*

▸ (`error`: Error | undefined, `messageType`: [MessageType](../modules/olivehelps.md#messagetype), `data`: ArrayBuffer): *void*

**Parameters:**

Name | Type |
------ | ------ |
`error` | Error &#124; undefined |
`messageType` | [MessageType](../modules/olivehelps.md#messagetype) |
`data` | ArrayBuffer |

▪ **returnCb**: *[ReturnCallback](../modules/olivehelps.md#returncallback)*

## Methods

###  close

▸ **close**(`callback`: function): *void*

*Defined in [types/oliveHelps/index.d.ts:182](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/types/oliveHelps/index.d.ts#L182)*

**Parameters:**

▪ **callback**: *function*

▸ (`error`: Error | undefined): *void*

**Parameters:**

Name | Type |
------ | ------ |
`error` | Error &#124; undefined |

**Returns:** *void*

___

###  onCloseHandler

▸ **onCloseHandler**(`callback`: function): *void*

*Defined in [types/oliveHelps/index.d.ts:187](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/types/oliveHelps/index.d.ts#L187)*

**Parameters:**

▪ **callback**: *function*

▸ (`error`: Error | undefined, `code`: number, `text`: string): *void*

**Parameters:**

Name | Type |
------ | ------ |
`error` | Error &#124; undefined |
`code` | number |
`text` | string |

**Returns:** *void*

___

###  writeMessage

▸ **writeMessage**(`messageType`: [MessageType](../modules/olivehelps.md#messagetype), `data`: Array‹number›, `callback`: function): *void*

*Defined in [types/oliveHelps/index.d.ts:177](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/types/oliveHelps/index.d.ts#L177)*

**Parameters:**

▪ **messageType**: *[MessageType](../modules/olivehelps.md#messagetype)*

▪ **data**: *Array‹number›*

▪ **callback**: *function*

▸ (`error`: Error | undefined): *void*

**Parameters:**

Name | Type |
------ | ------ |
`error` | Error &#124; undefined |

**Returns:** *void*
