package utils

import "github.com/open-olive/loop-development-kit/ldk/go/v2/proto"

type FileAction int

func (f FileAction) String() string {
	switch f {
	case FileActionUnknown:
		return "Unknown"
	case FileActionCreate:
		return "Create"
	case FileActionWrite:
		return "Write"
	case FileActionRemove:
		return "Remove"
	case FileActionRename:
		return "Rename"
	case FileActionChmod:
		return "Chmod"
	default:
		return "Unknown"
	}
}

const (
	FileActionUnknown FileAction = 0
	FileActionCreate  FileAction = 1
	FileActionWrite   FileAction = 2
	FileActionRemove  FileAction = 3
	FileActionRename  FileAction = 4
	FileActionChmod   FileAction = 5
)

func (f FileAction) ToProto() proto.FileAction {
	switch f {
	case FileActionCreate:
		return proto.FileAction_FILE_ACTION_CREATE
	case FileActionWrite:
		return proto.FileAction_FILE_ACTION_WRITE
	case FileActionRemove:
		return proto.FileAction_FILE_ACTION_REMOVE
	case FileActionRename:
		return proto.FileAction_FILE_ACTION_RENAME
	case FileActionChmod:
		return proto.FileAction_FILE_ACTION_CHMOD
	case FileActionUnknown:
		return proto.FileAction_FILE_ACTION_UNKNOWN
	default:
		return proto.FileAction_FILE_ACTION_UNKNOWN
	}
}

func ProtoActionToAction(action proto.FileAction) FileAction {
	switch action {
	case proto.FileAction_FILE_ACTION_CREATE:
		return FileActionCreate
	case proto.FileAction_FILE_ACTION_WRITE:
		return FileActionWrite
	case proto.FileAction_FILE_ACTION_REMOVE:
		return FileActionRemove
	case proto.FileAction_FILE_ACTION_RENAME:
		return FileActionRename
	case proto.FileAction_FILE_ACTION_CHMOD:
		return FileActionChmod
	case proto.FileAction_FILE_ACTION_UNKNOWN:
		return FileActionUnknown
	default:
		return FileActionUnknown
	}
}
