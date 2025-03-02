package dto

type PortScannerReq struct {
	StartPort       int    `json:"start_port"`
	EndPort         int    `json:"end_port"`
	IPAddress       string `json:"ip_address"`
	NumberOfThreads int    `json:"number_of_threads"`
	Type            string `json:"type"`
}

type PortScannerRes struct {
	Status string `json:"status"`
	Port   int    `json:"port"`
}
