package handler

import (
	"github.com/agustincaro/gRPC/internal/buyer"
	"github.com/agustincaro/gRPC/pkg/errs"
	"github.com/agustincaro/gRPC/pkg/web"
	"github.com/gin-gonic/gin"
)

type Buyer struct {
	buyerService buyer.Service
}

var errorsMsgBuyer = errs.ErrorsMsg{
	ErrNotFound: buyer.ErrNotFound.Error(),
	ErrDb:       buyer.ErrDb.Error(),
	ErrServer:   buyer.ErrServer.Error(),
	ErrCid:      buyer.ErrCid.Error()}

func NewBuyer(b buyer.Service) *Buyer {
	return &Buyer{
		buyerService: b,
	}
}

func (b *Buyer) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		buyers, err := b.buyerService.GetAll(c)
		if err != nil {

			errs.Err(c, err, errorsMsgBuyer)
			return
		}

		if len(buyers) == 0 {
			web.Success(c, 204, "No buyers in database")
			return
		}
		web.Success(c, 200, buyers)

	}
}
