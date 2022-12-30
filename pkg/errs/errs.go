package errs

import (
	"github.com/agustincaro/gRPC/pkg/web"
	"github.com/gin-gonic/gin"
)

type ErrorsMsg struct {
	ErrNotFound   string
	ErrDb         string
	ErrServer     string
	ErrCid        string
	ErrFields     string
	ErrIdNotFound string
	ErrUnique     string
	ErrProductId  string
	ErrZipCode    string
	ErrZipExists  string
}

func Err(ctx *gin.Context, err error, errorsMsg ErrorsMsg) {

	switch err.Error() {
	case errorsMsg.ErrNotFound:
		web.Error(ctx, 404, err.Error())
		return
	case errorsMsg.ErrDb:
		web.Success(ctx, 200, err.Error())
		return
	case errorsMsg.ErrCid:
		web.Error(ctx, 409, err.Error())
		return
	case errorsMsg.ErrServer:
		web.Error(ctx, 500, err.Error())
		return
	case errorsMsg.ErrFields:
		web.Error(ctx, 422, err.Error())
		return
	case errorsMsg.ErrIdNotFound:
		web.Error(ctx, 404, err.Error())
		return
	case errorsMsg.ErrUnique:
		web.Error(ctx, 409, err.Error())
		return
	case errorsMsg.ErrZipCode:
		web.Error(ctx, 409, err.Error())
		return
	case errorsMsg.ErrZipExists:
		web.Error(ctx, 404, err.Error())
		return
	default:
		web.Error(ctx, 500, err.Error())
	}

}
