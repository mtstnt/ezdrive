package fs

import "strings"

// "Entity" is basically any file/folder. Idk what to name it...
type DriveEntity struct {
	ID    string
	Name  string
	IsDir bool
}

type DriveFolderList []DriveEntity

func (df DriveFolderList) GetFullPath() string {
	var b strings.Builder
	b.Grow(len(df))
	for idx, node := range df {
		b.WriteString(node.Name)
		if idx != len(df)-1 {
			b.WriteString("/")
		}
	}
	return b.String()
}

type ReplState struct {
	// The current account used. Defaults to "" (no account used)
	account string
	// The current workdir path.
	path DriveFolderList
}

func NewReplState() *ReplState {
	return &ReplState{
		account: "",
		path:    make([]DriveEntity, 0),
	}
}

func (r *ReplState) List() []DriveEntity {
	// Call to main drive service
	// GetDriveService().List(...)
	return []DriveEntity{}
}
