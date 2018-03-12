package scanners

import (
	"../scan_manager"
	"fmt"
	"strings"
)

type FirefoxScanner struct{}

func (scanner *FirefoxScanner) HandleFile(obj *scan_manager.FileScanObject) {
	if strings.Contains(obj.Path(), "key3.db") {
		// TODO: figure out if key3.db is a firefox database
		// TODO: figure out if key3.db contains the master key or w/e
		// TODO: figure out how to get logins from key3 and other files.
		// TODO: print out logins or the file contents if an option is enabled and if we're unable to get the password for the keystore
		fmt.Println(obj.Path())
	}
}

func (scanner *FirefoxScanner) ProcessScanObject(obj scan_manager.ScanObject) error {
	switch v := obj.(type) {
	case *scan_manager.FileScanObject:
		scanner.HandleFile(v)
	}
	return nil
}

func init() {
	scanner := &FirefoxScanner{}
	scan_manager.RegisterScanner(scanner)
}
