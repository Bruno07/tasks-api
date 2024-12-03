package policies

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTaskPolicy_Allow(t *testing.T) {
	t.Run("Must allow to create", func(t *testing.T) {

		gin.SetMode(gin.TestMode)

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)

		ctx.Set("permissions", []interface{}{"CREATE", "UPDATE", "VIEW", "DELETE"})

		taskPolicy := TaskPolicy{}

		assert.True(t, taskPolicy.Allow("CREATE", ctx))

	})

	t.Run("Should not allow to create", func(t *testing.T) {

		gin.SetMode(gin.TestMode)

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)

		ctx.Set("permissions", []interface{}{"UPDATE", "VIEW", "DELETE"})

		taskPolicy := TaskPolicy{}

		assert.False(t, taskPolicy.Allow("CREATE", ctx))

	})
}