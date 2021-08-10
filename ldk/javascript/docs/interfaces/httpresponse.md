[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [HTTPResponse](httpresponse.md)

# Interface: HTTPResponse

The HTTP Response data.

## Hierarchy

* **HTTPResponse**

## Index

### Properties

* [body](httpresponse.md#body)
* [headers](httpresponse.md#headers)
* [statusCode](httpresponse.md#statuscode)

## Properties

###  body

• **body**: *Uint8Array*

*Defined in [network/types.ts:40](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/types.ts#L40)*

The HTTP response as a byte array. To decode into a UTF-8 string you can:
```
let decodedText = network.decode(data);
```

___

###  headers

• **headers**: *Record‹string, string[]›*

*Defined in [network/types.ts:41](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/types.ts#L41)*

___

###  statusCode

• **statusCode**: *number*

*Defined in [network/types.ts:33](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/network/types.ts#L33)*
