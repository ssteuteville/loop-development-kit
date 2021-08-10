[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [SocketConfiguration](socketconfiguration.md)

# Interface: SocketConfiguration

Configuration object to configure a websocket handshake

## Hierarchy

* **SocketConfiguration**

## Index

### Properties

* [headers](socketconfiguration.md#optional-headers)
* [subprotocols](socketconfiguration.md#optional-subprotocols)
* [url](socketconfiguration.md#url)
* [useCompression](socketconfiguration.md#optional-usecompression)

## Properties

### `Optional` headers

• **headers**? : *Record‹string, string[]›*

*Defined in [network/types.ts:60](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/types.ts#L60)*

Collection of the handshake headers

___

### `Optional` subprotocols

• **subprotocols**? : *Array‹string›*

*Defined in [network/types.ts:68](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/types.ts#L68)*

Specifies the client's requested subprotocols

___

###  url

• **url**: *string*

*Defined in [network/types.ts:56](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/types.ts#L56)*

Websocket server endpoint url: '{schema}://{host}:{port}' - schema should be wss

___

### `Optional` useCompression

• **useCompression**? : *undefined | false | true*

*Defined in [network/types.ts:64](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/types.ts#L64)*

Specifies if compression is used
