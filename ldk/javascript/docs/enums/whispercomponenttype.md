[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [WhisperComponentType](whispercomponenttype.md)

# Enumeration: WhisperComponentType

## Index

### Enumeration members

* [Box](whispercomponenttype.md#box)
* [Button](whispercomponenttype.md#button)
* [Checkbox](whispercomponenttype.md#checkbox)
* [CollapseBox](whispercomponenttype.md#collapsebox)
* [DateTimeInput](whispercomponenttype.md#datetimeinput)
* [Divider](whispercomponenttype.md#divider)
* [Email](whispercomponenttype.md#email)
* [Link](whispercomponenttype.md#link)
* [ListPair](whispercomponenttype.md#listpair)
* [Markdown](whispercomponenttype.md#markdown)
* [Message](whispercomponenttype.md#message)
* [Number](whispercomponenttype.md#number)
* [Password](whispercomponenttype.md#password)
* [RadioGroup](whispercomponenttype.md#radiogroup)
* [SectionTitle](whispercomponenttype.md#sectiontitle)
* [Select](whispercomponenttype.md#select)
* [Telephone](whispercomponenttype.md#telephone)
* [TextInput](whispercomponenttype.md#textinput)

## Enumeration members

###  Box

• **Box**: = "box"

*Defined in [whisper/types.ts:5](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L5)*

A container component for formatting other components.

___

###  Button

• **Button**: = "button"

*Defined in [whisper/types.ts:6](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L6)*

___

###  Checkbox

• **Checkbox**: = "checkbox"

*Defined in [whisper/types.ts:7](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L7)*

___

###  CollapseBox

• **CollapseBox**: = "collapseBox"

*Defined in [whisper/types.ts:11](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L11)*

A container component to allow content to be opened and closed with a button click.

___

###  DateTimeInput

• **DateTimeInput**: = "dateTimeInput"

*Defined in [whisper/types.ts:69](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L69)*

___

###  Divider

• **Divider**: = "divider"

*Defined in [whisper/types.ts:15](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L15)*

This component shows a horizontal divider to separate different kinds on content in a whisper. This component has no options.

___

###  Email

• **Email**: = "email"

*Defined in [whisper/types.ts:19](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L19)*

The text input field allows the user to provide an email address.

___

###  Link

• **Link**: = "link"

*Defined in [whisper/types.ts:23](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L23)*

This component shows a link that can either open a link in the user's default browser or function as an `onClick` to allow for loops to do things like send a new whisper.

___

###  ListPair

• **ListPair**: = "listPair"

*Defined in [whisper/types.ts:27](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L27)*

This component shows a two column view of information typically used for lists of information.

___

###  Markdown

• **Markdown**: = "markdown"

*Defined in [whisper/types.ts:28](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L28)*

___

###  Message

• **Message**: = "message"

*Defined in [whisper/types.ts:32](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L32)*

This component shows a banner in the whisper that functions as a call to action to the user.

___

###  Number

• **Number**: = "number"

*Defined in [whisper/types.ts:36](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L36)*

The text input field allows the user to provide a number within the parameters provided.

___

###  Password

• **Password**: = "password"

*Defined in [whisper/types.ts:40](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L40)*

The password input field allows the user to provide a password. This field protects the user by obscuring what they type. Showing each character as a solid black dot.

___

###  RadioGroup

• **RadioGroup**: = "radioGroup"

*Defined in [whisper/types.ts:46](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L46)*

The radio group allows a loop to provide the user with a collection of options in which they select a single result. The result is selected by clicking one of the radio elements in the radio group.

A selected value of -1 indicates that nothing is selected.

___

###  SectionTitle

• **SectionTitle**: = "sectionTitle"

*Defined in [whisper/types.ts:64](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L64)*

The section title field allows the user to provide section title information.

___

###  Select

• **Select**: = "select"

*Defined in [whisper/types.ts:50](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L50)*

A selected value of -1 indicates that nothing is selected.

___

###  Telephone

• **Telephone**: = "telephone"

*Defined in [whisper/types.ts:54](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L54)*

The text input field allows the user to provide a telephone number.

___

###  TextInput

• **TextInput**: = "textInput"

*Defined in [whisper/types.ts:60](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/whisper/types.ts#L60)*

The text input field allows the user to provide text information.

The text can be pre-populated by the loop.
