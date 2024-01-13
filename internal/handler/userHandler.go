package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-gin-mall/internal/domain/user"
)

func RegisterUserHandler(r *gin.Engine, handler *UserHandler) {
	g := r.Group("/user")
	g.POST("/register", handler.Register)
}

func NewUserHandler(logger *logrus.Logger, service *user.UserDomainService) *UserHandler {
	return &UserHandler{
		Logger:    logger,
		DomainSvc: service,
	}
}

type UserHandler struct {
	Logger    *logrus.Logger
	DomainSvc *user.UserDomainService
}

func (h *UserHandler) Register(c *gin.Context) {
	//var in HelloRequest
	//if err := ctx.BindQuery(&in); err != nil {
	//	return err
	//}
	//if err := ctx.BindVars(&in); err != nil {
	//	return err
	//}
	//http.SetOperation(ctx, OperationDemoSayHello)
	//h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
	//	return srv.SayHello(ctx, req.(*HelloRequest))
	//})
	//out, err := h(ctx, &in)
	//if err != nil {
	//	return err
	//}
	//reply := out.(*HelloReply)
	//return ctx.Result(200, reply)
}
