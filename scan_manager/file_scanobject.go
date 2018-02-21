package scan_manager

import "os"

type FileScanObject struct {
	path      string
	file_info os.FileInfo
}

func (obj *FileScanObject) Path() string {
	return obj.path
}

func (obj *FileScanObject) FileInfo() os.FileInfo {
	return obj.file_info
}
