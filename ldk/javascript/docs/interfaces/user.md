[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [User](user.md)

# Interface: User

## Hierarchy

* **User**

## Index

### Methods

* [jwt](user.md#jwt)

## Methods

###  jwt

▸ **jwt**(`jwtConfig?`: [JWTConfig](jwtconfig.md)): *Promise‹string›*

*Defined in [user/index.ts:16](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/user/index.ts#L16)*

Returns a JWT identifying the current OliveHelps user.

**Parameters:**

Name | Type |
------ | ------ |
`jwtConfig?` | [JWTConfig](jwtconfig.md) |

**Returns:** *Promise‹string›*

JWT with the current username in the subject field.
