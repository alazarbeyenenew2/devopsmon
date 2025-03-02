package scanner

import (
	"context"
	"fmt"
	"net"

	"github.com/alazarbeyenenew2/devopsmon/internal/constant"
	"github.com/alazarbeyenenew2/devopsmon/internal/constant/dto"
	"go.uber.org/zap"
)

func workers(ports, results chan int, endPoint string) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", endPoint, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		results <- p
		conn.Close()

	}

}

func (s *scanner) PortScanner(ctx context.Context, req dto.PortScannerReq) ([]dto.PortScannerRes, error) {
	var resp []dto.PortScannerRes
	results := make(chan int)

	if req.NumberOfThreads <= 0 {
		//use default number of threads
		req.NumberOfThreads = 10
	}

	if req.StartPort == 0 && req.Type != constant.PORT_SCANNER_TYPE_SINGLE {
		s.log.Error("please provide port number", zap.Any("req", req))
		return []dto.PortScannerRes{}, fmt.Errorf("please provide port number")
	}

	if req.Type == constant.PORT_SCANNER_TYPE_RANGE && (req.StartPort > req.EndPort || req.EndPort <= 0 || req.StartPort < 0) {
		s.log.Error("please provide port number", zap.Any("req", req))
		return []dto.PortScannerRes{}, fmt.Errorf("please provide valid start and end ip address")
	}

	ports := make(chan int, req.NumberOfThreads)
	for i := 0; i < cap(ports); i++ {
		go workers(ports, results, req.IPAddress)
	}

	go func() {
		for i := req.StartPort; i <= req.EndPort; i++ {
			ports <- i
		}
	}()

	for i := req.StartPort; i <= req.EndPort; i++ {
		port := <-results
		if port != 0 {
			resp = append(resp, dto.PortScannerRes{
				Status: "OPEN",
				Port:   port,
			})
		}
	}

	return resp, nil
}
