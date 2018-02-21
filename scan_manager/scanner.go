package scan_manager

type Scanner interface {
	ProcessScanObject(object ScanObject) error
}
