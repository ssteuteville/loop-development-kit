[Javascript Loop Development Kit](../README.md) › [Globals](../globals.md) › [Filesystem](filesystem.md)

# Interface: Filesystem

The FileSystem interfaces provides access to the ability to read, write, delete files.

## Hierarchy

* **Filesystem**

## Index

### Methods

* [copy](filesystem.md#copy)
* [dir](filesystem.md#dir)
* [exists](filesystem.md#exists)
* [join](filesystem.md#join)
* [listenDir](filesystem.md#listendir)
* [listenFile](filesystem.md#listenfile)
* [makeDir](filesystem.md#makedir)
* [move](filesystem.md#move)
* [readFile](filesystem.md#readfile)
* [remove](filesystem.md#remove)
* [stat](filesystem.md#stat)
* [writeFile](filesystem.md#writefile)

## Methods

###  copy

▸ **copy**(`source`: string, `destination`: string): *Promise‹void›*

*Defined in [filesystem/index.ts:31](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/filesystem/index.ts#L31)*

Copies a file from one location to another.

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`source` | string | path to the file to copy from |
`destination` | string | path to the file to copy to  |

**Returns:** *Promise‹void›*

___

###  dir

▸ **dir**(`path`: string): *Promise‹[FileInfo](fileinfo.md)[]›*

*Defined in [filesystem/index.ts:37](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/filesystem/index.ts#L37)*

Returns all files in the specified directory

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`path` | string | path of the specified directory  |

**Returns:** *Promise‹[FileInfo](fileinfo.md)[]›*

___

###  exists

▸ **exists**(`path`: string): *Promise‹boolean›*

*Defined in [filesystem/index.ts:43](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/filesystem/index.ts#L43)*

Return true if a file or directory exists at the specified location.

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`path` | string | path to the specified file or directory  |

**Returns:** *Promise‹boolean›*

___

###  join

▸ **join**(`segments`: string[]): *Promise‹string›*

*Defined in [filesystem/index.ts:112](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/filesystem/index.ts#L112)*

Join joins an array of path elements into a single path, separating them with an OS specific Separator.
Empty elements are ignored. The result is Cleaned. However, if the argument list is empty or all its elements are empty,
Join returns an empty string. On Windows, the result will only be a UNC path if the first non-empty element is a UNC path.

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`segments` | string[] | an array of path segments to join |

**Returns:** *Promise‹string›*

- a single path separated with an OS specific Separator

___

###  listenDir

▸ **listenDir**(`path`: string, `callback`: function): *Promise‹[Cancellable](cancellable.md)›*

*Defined in [filesystem/index.ts:51](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/filesystem/index.ts#L51)*

Listen changes to the contents of the directory.

**Parameters:**

▪ **path**: *string*

path of directory to listen.

▪ **callback**: *function*

the callback function that's called when a file in the directory changes.

▸ (`fileEvent`: [FileEvent](fileevent.md) | [RenamedFileEvent](renamedfileevent.md) | [RemovedFileEvent](removedfileevent.md)): *void*

**Parameters:**

Name | Type |
------ | ------ |
`fileEvent` | [FileEvent](fileevent.md) &#124; [RenamedFileEvent](renamedfileevent.md) &#124; [RemovedFileEvent](removedfileevent.md) |

**Returns:** *Promise‹[Cancellable](cancellable.md)›*

___

###  listenFile

▸ **listenFile**(`path`: string, `callback`: function): *Promise‹[Cancellable](cancellable.md)›*

*Defined in [filesystem/index.ts:62](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/filesystem/index.ts#L62)*

Listen changes to a specific file.

**Parameters:**

▪ **path**: *string*

path of file to listen.

▪ **callback**: *function*

the callback function called when the file changes.

▸ (`fileEvent`: [FileEvent](fileevent.md) | [RenamedFileEvent](renamedfileevent.md) | [RemovedFileEvent](removedfileevent.md)): *void*

**Parameters:**

Name | Type |
------ | ------ |
`fileEvent` | [FileEvent](fileevent.md) &#124; [RenamedFileEvent](renamedfileevent.md) &#124; [RemovedFileEvent](removedfileevent.md) |

**Returns:** *Promise‹[Cancellable](cancellable.md)›*

___

###  makeDir

▸ **makeDir**(`destination`: string, `writeMode`: [WriteMode](../globals.md#writemode)): *Promise‹void›*

*Defined in [filesystem/index.ts:72](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/filesystem/index.ts#L72)*

Makes a directory at the specified location.

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`destination` | string | destination of where directory needs to be created at |
`writeMode` | [WriteMode](../globals.md#writemode) | file mode and permission bits  |

**Returns:** *Promise‹void›*

___

###  move

▸ **move**(`source`: string, `destination`: string): *Promise‹void›*

*Defined in [filesystem/index.ts:79](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/filesystem/index.ts#L79)*

Moves a file from one location to another.

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`source` | string | path of where to move the file from |
`destination` | string | path of where to move the file to  |

**Returns:** *Promise‹void›*

___

###  readFile

▸ **readFile**(`path`: string): *Promise‹Uint8Array›*

*Defined in [filesystem/index.ts:85](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/filesystem/index.ts#L85)*

Returns the contents of the specified file.

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`path` | string | path of the specified file  |

**Returns:** *Promise‹Uint8Array›*

___

###  remove

▸ **remove**(`source`: string): *Promise‹void›*

*Defined in [filesystem/index.ts:91](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/filesystem/index.ts#L91)*

Removes a file/directory at the specified path.

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`source` | string | path of the file or directory to remove  |

**Returns:** *Promise‹void›*

___

###  stat

▸ **stat**(`path`: string): *Promise‹[FileInfo](fileinfo.md)›*

*Defined in [filesystem/index.ts:97](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/filesystem/index.ts#L97)*

Returns info about a specified file/directory.

**Parameters:**

Name | Type | Description |
------ | ------ | ------ |
`path` | string | path to the specified file or directory  |

**Returns:** *Promise‹[FileInfo](fileinfo.md)›*

___

###  writeFile

▸ **writeFile**(`params`: [WriteFileParams](writefileparams.md)): *Promise‹void›*

*Defined in [filesystem/index.ts:102](https://github.com/open-olive/loop-development-kit/blob/ba5f0aac/ldk/javascript/src/filesystem/index.ts#L102)*

Writes (overwrites or appends) data to the specified file with specific permissions. New file will be created if file not exist

**Parameters:**

Name | Type |
------ | ------ |
`params` | [WriteFileParams](writefileparams.md) |

**Returns:** *Promise‹void›*
