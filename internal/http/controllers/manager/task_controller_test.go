package manager

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Bruno07/tasks-api/internal/http/policies"
	"github.com/Bruno07/tasks-api/internal/http/requests"
	"github.com/Bruno07/tasks-api/internal/http/responses"
	"github.com/Bruno07/tasks-api/internal/repositories"
	"github.com/Bruno07/tasks-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTaskController_All(t *testing.T) {

    taskService := services.NewTaskService(
		&repositories.TaskMockRepository{},
		repositories.UserMockRepository{},
	)
	
	taskController := NewTaskController(
		*taskService,
		policies.TaskPolicy{},
	)

	t.Run("Should return a list of tasks", func(t *testing.T) {

		gin.SetMode(gin.TestMode)

		w := httptest.NewRecorder()
		
		ctx, _ := gin.CreateTestContext(w)

		_, err := http.NewRequest(http.MethodGet, "/api/manager/tasks", nil)
		if err != nil {
			panic(err.Error())
		}

		ctx.Set("permissions", []interface{}{"CREATE", "UPDATE", "VIEW", "DELETE"})

		taskController.All(ctx)

		var response []responses.TaskResponse
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, 3, len(response))

	})
}

func TestTaskController_Delete(t *testing.T) {

	t.Run("It should return without an error", func(t *testing.T) {
		taskService := services.NewTaskService(
			&repositories.TaskMockRepository{},
			&repositories.UserMockRepository{},
		)

		err := taskService.Delete(requests.TaskRequest{
			ID: 1,
		})
	
		assert.NoError(t, err)
	})

	t.Run("It should return with an error", func(t *testing.T) {
		taskService := services.NewTaskService(
			&repositories.TaskMockRepository{},
			&repositories.UserMockRepository{},
		)

		err := taskService.Delete(requests.TaskRequest{
			ID: 4,
		})
	
		assert.Error(t, err)
	})
}
