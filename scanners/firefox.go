package scanners

import (
	"../scan_manager"
	"fmt"
	"strings"
)

type FirefoxScanner struct{}

func (scanner *FirefoxScanner) HandleFile(obj *scan_manager.FileScanObject) {
	if strings.Contains(obj.Path(), "key3.db") {
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
