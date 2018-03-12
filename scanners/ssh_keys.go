package scanners

import (
	"../scan_manager"
	"os"
	"fmt"
	"path/filepath"
)

type SSHKeyScanner struct{}

func (scanner *SSHKeyScanner) HandleFile(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		if f.Name() == "known_hosts" {
			// TODO: Dump list of hosts this user has connected to
			fmt.Println("Known hosts: " + path)
		} else if f.Name() == "authorized_keys" {
			// TODO: Dump list of users and keys that log into this machine
			fmt.Println("Authorized keys: " + path)
		} else if f.Name() == "config" {
			// TODO: Dump list of config options that this user uses to connect (also get hosts that this user has configured to connect to)
			fmt.Println("Config: " + path)
		} else {
			// TODO: dump keys
			fmt.Println("Key: " + path)
		}
	}
	return nil
}

func (scanner *SSHKeyScanner) HandleDirectory(obj *scan_manager.DirScanObject) {
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
