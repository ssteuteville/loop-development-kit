[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [Network](network.md)

# Interface: Network

The Network Aptitude provides access to network calls through Olive Helps

## Hierarchy

* **Network**

## Index

### Methods

* [decode](network.md#decode)
* [encode](network.md#encode)
* [httpRequest](network.md#httprequest)
* [webSocketConnect](network.md#websocketconnect)

## Methods

###  decode

▸ **decode**(`encodedValue`: Uint8Array, `encoding`: string): *Promise‹string›*

*Defined in [network/index.ts:35](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/index.ts#L35)*

Decoding provided value

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`encodedValue` | Uint8Array | Specified encoded value to decode |
`encoding` | string | - |

**Returns:** *Promise‹string›*

A promise resolving with the decoded text

___

###  encode

▸ **encode**(`text`: string, `encoding`: string): *Promise‹Uint8Array›*

*Defined in [network/index.ts:27](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/index.ts#L27)*

Encoding provided text

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`text` | string | Specified text to encode |
`encoding` | string | - |

**Returns:** *Promise‹Uint8Array›*

A promise resolving with the encoded Uint8Array

___

###  httpRequest

▸ **httpRequest**(`request`: [HTTPRequest](httprequest.md)): *Promise‹[HTTPResponse](httpresponse.md)›*

*Defined in [network/index.ts:19](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/index.ts#L19)*

Generates a HTTP request with the provided configuration.

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`request` | [HTTPRequest](httprequest.md) | The HTTP request to make. |

**Returns:** *Promise‹[HTTPResponse](httpresponse.md)›*

A Promise resolving with the response.

___

###  webSocketConnect

▸ **webSocketConnect**(`socketConfiguration`: [SocketConfiguration](socketconfiguration.md)): *Promise‹[Socket](socket.md)›*

*Defined in [network/index.ts:43](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/index.ts#L43)*

 Connects to a specified websocket

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`socketConfiguration` | [SocketConfiguration](socketconfiguration.md) | A configuration object defines websocket |

**Returns:** *Promise‹[Socket](socket.md)›*

A promise with Socket
