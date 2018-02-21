package scanners

import (
	"../scan_manager"
	"fmt"
)

type SSHKeyScanner struct{}

func HandleDirectory(obj *scan_manager.DirScanObject) {
	info := obj.FileInfo();
	if info.Name() == ".ssh" {
		fmt.Printf("%s    %s    %s\n", info.ModTime(), info.Mode(), obj.Path());
	}
}

func (scanner *SSHKeyScanner) ProcessScanObject(obj scan_manager.ScanObject) error {
	switch v := obj.(type) {
	case *scan_manager.DirScanObject:
		HandleDirectory(v)
		break
	}
	return nil
}

func init() {
	scanner := &SSHKeyScanner{}
	scan_manager.RegisterScanner(scanner)
}
