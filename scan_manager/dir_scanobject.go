package scan_manager

import "os"

type DirScanObject struct {
	path      string
	file_info os.FileInfo
}

func (obj *DirScanObject) Path() string {
	return obj.path
}

func (obj *DirScanObject) FileInfo() os.FileInfo {
	return obj.file_info
}
