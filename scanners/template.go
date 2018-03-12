package scanners

import (
	"../scan_manager"
	"fmt"
)

type ScannerExample struct{}

func (scanner *ScannerExample) HandleFile(obj *scan_manager.FileScanObject) {
	fmt.Println("FileScanObject");
}

func (scanner *ScannerExample) HandleDir(obj *scan_manager.DirScanObject) {
	fmt.Println("DirScanObject");
}

func (scanner *ScannerExample) HandleEnvVariable(obj *scan_manager.EnvVariableScanObject) {
	fmt.Println("EnvVariableScanObject");
}

func (scanner *ScannerExample) ProcessScanObject(obj scan_manager.ScanObject) error {
	switch v := obj.(type) {
	case *scan_manager.DirScanObject:
		scanner.HandleDir(v)
		break
	case *scan_manager.FileScanObject:
		scanner.HandleFile(v)
		break
	case *scan_manager.EnvVariableScanObject:
		scanner.HandleEnvVariable(v)
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
