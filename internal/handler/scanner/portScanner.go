package scanner

import (
	"net/http"

	"github.com/alazarbeyenenew2/devopsmon/internal/constant/dto"
	"github.com/alazarbeyenenew2/devopsmon/internal/constant/errors"
	"github.com/alazarbeyenenew2/devopsmon/internal/constant/model/response"
	"github.com/gin-gonic/gin"
)

func (s *scanner) PortScanner(c *gin.Context) {
	var portScannerReq dto.PortScannerReq
	if err := c.ShouldBind(&portScannerReq); err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, err.Error())
		_ = c.Error(err)
		return
	}

	resp, err := s.scannerModule.PortScanner(c, portScannerReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.SendSuccessResponse(c, http.StatusOK, resp)

}
