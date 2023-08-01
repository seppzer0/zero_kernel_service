package routing

import (
	"github.com/gin-gonic/gin"
)

type Bindable interface {
	Bind(router gin.IRouter)
}
