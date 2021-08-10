[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [OliveHelps](../modules/olivehelps.md) › [Filesystem](olivehelps.filesystem.md)

# Interface: Filesystem

## Hierarchy

* **Filesystem**

## Index

### Properties

* [copy](olivehelps.filesystem.md#copy)
* [dir](olivehelps.filesystem.md#dir)
* [exists](olivehelps.filesystem.md#exists)
* [join](olivehelps.filesystem.md#join)
* [listenDir](olivehelps.filesystem.md#listendir)
* [listenFile](olivehelps.filesystem.md#listenfile)
* [makeDir](olivehelps.filesystem.md#makedir)
* [move](olivehelps.filesystem.md#move)
* [readFile](olivehelps.filesystem.md#readfile)
* [remove](olivehelps.filesystem.md#remove)
* [stat](olivehelps.filesystem.md#stat)
* [writeFile](olivehelps.filesystem.md#writefile)

## Properties

###  copy

• **copy**: *[ReadableWithTwoParams](../modules/olivehelps.md#readablewithtwoparams)‹string, string, void›*

*Defined in [types/oliveHelps/index.d.ts:273](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/types/oliveHelps/index.d.ts#L273)*

___

###  dir

• **dir**: *[ReadableWithParam](../modules/olivehelps.md#readablewithparam)‹string, [FileInfo](olivehelps.fileinfo.md)[]›*

*Defined in [types/oliveHelps/index.d.ts:275](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/types/oliveHelps/index.d.ts#L275)*

___

###  exists

• **exists**: *[ReadableWithParam](../modules/olivehelps.md#readablewithparam)‹string, boolean›*

*Defined in [types/oliveHelps/index.d.ts:277](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/types/oliveHelps/index.d.ts#L277)*

___

###  join

• **join**: *[ReadableWithParam](../modules/olivehelps.md#readablewithparam)‹string[], string›*

*Defined in [types/oliveHelps/index.d.ts:295](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/types/oliveHelps/index.d.ts#L295)*

___

###  listenDir

• **listenDir**: *[ListenableWithParam](../modules/olivehelps.md#listenablewithparam)‹string, [FileEvent](olivehelps.fileevent.md)›*

*Defined in [types/oliveHelps/index.d.ts:279](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/types/oliveHelps/index.d.ts#L279)*

___

###  listenFile

• **listenFile**: *[ListenableWithParam](../modules/olivehelps.md#listenablewithparam)‹string, [FileEvent](olivehelps.fileevent.md)›*

*Defined in [types/oliveHelps/index.d.ts:281](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/types/oliveHelps/index.d.ts#L281)*

___

###  makeDir

• **makeDir**: *[ReadableWithTwoParams](../modules/olivehelps.md#readablewithtwoparams)‹string, [WriteMode](../modules/olivehelps.md#writemode), void›*

*Defined in [types/oliveHelps/index.d.ts:283](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/types/oliveHelps/index.d.ts#L283)*

___

###  move

• **move**: *[ReadableWithTwoParams](../modules/olivehelps.md#readablewithtwoparams)‹string, string, void›*

*Defined in [types/oliveHelps/index.d.ts:285](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/types/oliveHelps/index.d.ts#L285)*

___

###  readFile

• **readFile**: *[ReadableWithParam](../modules/olivehelps.md#readablewithparam)‹string, ArrayBuffer›*

*Defined in [types/oliveHelps/index.d.ts:287](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/types/oliveHelps/index.d.ts#L287)*

___

###  remove

• **remove**: *[ReadableWithParam](../modules/olivehelps.md#readablewithparam)‹string, void›*

*Defined in [types/oliveHelps/index.d.ts:289](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/types/oliveHelps/index.d.ts#L289)*

___

###  stat

• **stat**: *[ReadableWithParam](../modules/olivehelps.md#readablewithparam)‹string, [FileInfo](olivehelps.fileinfo.md)›*

*Defined in [types/oliveHelps/index.d.ts:291](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/types/oliveHelps/index.d.ts#L291)*

___

###  writeFile

• **writeFile**: *[ReadableWithFourParams](../modules/olivehelps.md#readablewithfourparams)‹string, Array‹number›, [WriteOperation](../modules/olivehelps.md#writeoperation), [WriteMode](../modules/olivehelps.md#writemode), void›*

*Defined in [types/oliveHelps/index.d.ts:293](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/types/oliveHelps/index.d.ts#L293)*
