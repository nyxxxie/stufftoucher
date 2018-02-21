package scan_manager

import (
	"os"
	"strings"
	"path/filepath"
)

var scanner_objects []Scanner

func RegisterScanner(scanner Scanner) {
	scanner_objects = append(scanner_objects, scanner)
}

func ScanDirectory(path string) error {
	var err error = nil

	err = filepath.Walk(path, HandlePath)

	return err
}

func ScanEnvironmentVariables() error {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		obj := EnvVariableScanObject{pair[0], pair[1]}
		return ProcessScanObject(obj)
	}

	return nil
}

func ProcessScanObject(obj ScanObject) error {
	for _, scanner := range scanner_objects {
		scanner.ProcessScanObject(obj)
	}

	return nil
}

func HandlePath(path string, f os.FileInfo, err error) error {
	if f != nil {
		if f.IsDir() {
			obj := DirScanObject{path, f}
			return ProcessScanObject(&obj)
		} else {
			obj := FileScanObject{path, f}
			return ProcessScanObject(&obj)
		}
	}

	return nil
}
