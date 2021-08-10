[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [Process](process.md)

# Interface: Process

The ProcessService provides access to the list of running processes for all users.

## Hierarchy

* **Process**

## Index

### Methods

* [all](process.md#all)
* [listenAll](process.md#listenall)

## Methods

###  all

▸ **all**(): *Promise‹[ProcessInfo](processinfo.md)[]›*

*Defined in [process/index.ts:30](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/process/index.ts#L30)*

Gets a list of the current running processes for all users.

**Returns:** *Promise‹[ProcessInfo](processinfo.md)[]›*

a Promise resolving with a list of the current processes.

___

###  listenAll

▸ **listenAll**(`callback`: function): *Promise‹[Cancellable](cancellable.md)›*

*Defined in [process/index.ts:37](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/process/index.ts#L37)*

Starts listening for processes starting and stopping for all users.

**Parameters:**

▪ **callback**: *function*

callback function called every time a process is started or stopped.

▸ (`event`: [ProcessEvent](processevent.md)): *void*

**Parameters:**

Name | Type |
------ | ------ |
`event` | [ProcessEvent](processevent.md) |

**Returns:** *Promise‹[Cancellable](cancellable.md)›*
