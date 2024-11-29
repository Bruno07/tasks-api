package policies

import (
	"github.com/gin-gonic/gin"
)

type TaskPolicy struct {}

func (tp TaskPolicy) Allow(role string, ctx *gin.Context) bool {
	
	for _, permission := range ctx.MustGet("permissions").([]interface{}) {
		if permission == role {
			return true
		}
	}

	return false

}
