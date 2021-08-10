[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [InputComponent](inputcomponent.md)

# Interface: InputComponent ‹**T1, T2, T3**›

## Type parameters

▪ **T1**: *[WhisperComponentType](../enums/whispercomponenttype.md)*

▪ **T2**

▪ **T3**

## Hierarchy

* [WhisperComponent](whispercomponent.md)‹T1›

  ↳ **InputComponent**

## Index

### Properties

* [id](inputcomponent.md#optional-id)
* [key](inputcomponent.md#optional-key)
* [label](inputcomponent.md#label)
* [layout](inputcomponent.md#optional-layout)
* [onBlur](inputcomponent.md#optional-onblur)
* [onChange](inputcomponent.md#onchange)
* [onFocus](inputcomponent.md#optional-onfocus)
* [tooltip](inputcomponent.md#optional-tooltip)
* [type](inputcomponent.md#type)
* [validationError](inputcomponent.md#optional-validationerror)
* [value](inputcomponent.md#optional-value)

## Properties

### `Optional` id

• **id**? : *undefined | string*

*Inherited from [WhisperComponent](whispercomponent.md).[id](whispercomponent.md#optional-id)*

*Defined in [whisper/types.ts:163](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L163)*

___

### `Optional` key

• **key**? : *undefined | string*

*Inherited from [WhisperComponent](whispercomponent.md).[key](whispercomponent.md#optional-key)*

*Defined in [whisper/types.ts:169](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L169)*

The key is used to maintain the object state.
The component's key must be unique among its sibling components.

___

###  label

• **label**: *string*

*Defined in [whisper/types.ts:175](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L175)*

___

### `Optional` layout

• **layout**? : *[LayoutOptions](layoutoptions.md)*

*Inherited from [WhisperComponent](whispercomponent.md).[layout](whispercomponent.md#optional-layout)*

*Defined in [whisper/types.ts:170](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L170)*

___

### `Optional` onBlur

• **onBlur**? : *undefined | function*

*Defined in [whisper/types.ts:179](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L179)*

___

###  onChange

• **onChange**: *[WhisperHandlerWithParam](../globals.md#whisperhandlerwithparam)‹T3›*

*Defined in [whisper/types.ts:181](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L181)*

___

### `Optional` onFocus

• **onFocus**? : *undefined | function*

*Defined in [whisper/types.ts:180](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L180)*

___

### `Optional` tooltip

• **tooltip**? : *undefined | string*

*Defined in [whisper/types.ts:176](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L176)*

___

###  type

• **type**: *T1*

*Inherited from [WhisperComponent](whispercomponent.md).[type](whispercomponent.md#type)*

*Defined in [whisper/types.ts:164](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L164)*

___

### `Optional` validationError

• **validationError**? : *undefined | string*

*Defined in [whisper/types.ts:177](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L177)*

___

### `Optional` value

• **value**? : *T2*

*Defined in [whisper/types.ts:178](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L178)*
