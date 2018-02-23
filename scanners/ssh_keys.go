package scanners

import (
	"../scan_manager"
	"os"
	"fmt"
	"path/filepath"
)

type SSHKeyScanner struct{}

func HandleFile(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		if f.Name() == "known_hosts" {
			fmt.Println("Known hosts: " + path)
		} else if f.Name() == "authorized_keys" {
			fmt.Println("Authorized keys: " + path)
		} else if f.Name() == "config" {
			fmt.Println("Config: " + path)
		} else {
			fmt.Println("Key: " + path)
		}
	}
	return nil
}

func HandleDirectory(obj *scan_manager.DirScanObject) {
	info := obj.FileInfo();
	if info.Name() == ".ssh" {
		filepath.Walk(obj.Path(), HandleFile)
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
