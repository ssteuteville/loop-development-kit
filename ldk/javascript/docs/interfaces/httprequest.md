[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [HTTPRequest](httprequest.md)

# Interface: HTTPRequest

The HTTP Request configuration.

## Hierarchy

* **HTTPRequest**

## Index

### Properties

* [body](httprequest.md#optional-body)
* [headers](httprequest.md#optional-headers)
* [method](httprequest.md#method)
* [timeoutMs](httprequest.md#optional-timeoutms)
* [url](httprequest.md#url)

## Properties

### `Optional` body

• **body**? : *string | Uint8Array*

*Defined in [network/types.ts:10](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/types.ts#L10)*

Request Body to send to the provided url

___

### `Optional` headers

• **headers**? : *Record‹string, string[]›*

*Defined in [network/types.ts:14](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/types.ts#L14)*

Collection of request headers

___

###  method

• **method**: *string*

*Defined in [network/types.ts:18](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/types.ts#L18)*

Request method type

___

### `Optional` timeoutMs

• **timeoutMs**? : *undefined | number*

*Defined in [network/types.ts:26](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/types.ts#L26)*

HttpRequest timeout

___

###  url

• **url**: *string*

*Defined in [network/types.ts:22](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/types.ts#L22)*

Endpoint url
