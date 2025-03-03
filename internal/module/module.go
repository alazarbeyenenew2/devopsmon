package module

import (
	"context"

	"github.com/alazarbeyenenew2/devopsmon/internal/constant/dto"
)

type Scanner interface {
	PortScanner(ctx context.Context, req dto.PortScannerReq) ([]dto.PortScannerRes, error)
}
