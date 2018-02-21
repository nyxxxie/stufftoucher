package scanners

import (
	"../scan_manager"
	"fmt"
)

type ScannerExample struct{}

func (scanner *ScannerExample) ProcessScanObject(obj scan_manager.ScanObject) error {
	switch obj.(type) {
	case *scan_manager.DirScanObject:
		fmt.Println("DirScanObject");
		break
	case *scan_manager.FileScanObject:
		fmt.Println("FileScanObject");
		break
	case *scan_manager.EnvVariableScanObject:
		fmt.Println("EnvVariableScanObject");
		break
	default:
		fmt.Println("No idea what kind of object we're dealing with");
	}
	return nil
}

func init() {
	// TODO: uncomment this so that the scanner will be registered
	//scanner := &ScannerExample{}
	//scan_manager.RegisterScanner(scanner)
}
