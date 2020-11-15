package echo

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/coocn-cn/bc3/app/task/ent"
	"github.com/coocn-cn/bc3/app/task/ent/task"
	"github.com/coocn-cn/bc3/pkg/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewServer 创建一个服务器
func NewServer(client *ent.Client) *echo.Echo {
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			ec.Set("context.Context", context.Background())

			return h(ec)
		}
	})

	app.GET("/", func(ec echo.Context) error {
		ctx := ec.Get("context.Context").(context.Context)

		tasks, err := client.Task.Query().Where(task.Not(task.HasParent())).WithStatus().WithAssign(func(query *ent.TaskAssignQuery) {
			query.WithUser()
		}).WithChild(func(query *ent.TaskQuery) {
			query.WithStatus().WithAssign()
		}).All(ctx)
		if err != nil {
			return err
		}

		return ec.JSON(http.StatusOK, tasks)
	})

	app.POST("/", func(ec echo.Context) error {
		ctx := ec.Get("context.Context").(context.Context)

		cmd := ent.Task{}
		if err := ec.Bind(&cmd); err != nil {
			return err
		}

		save := &ent.Task{}
		err := db.WithTx(ctx, client, func(tx *ent.Tx) (err error) {
			builder := tx.Task.Create()

			builder = builder.
				SetOperatorID(1).
				SetTitle(cmd.Title).
				SetContent(cmd.Content).
				SetManHour(cmd.ManHour).
				SetStatusID(cmd.Edges.Status.ID).
				SetNillableEndTime(cmd.EndTime).
				SetNillableStartTime(cmd.StartTime)

			if cmd.Edges.Parent == nil {
				builder = builder.SetNillableParentID(nil)
			} else {
				builder = builder.SetParentID(cmd.Edges.Parent.ID)
			}

			save, err = builder.Save(ctx)
			if err != nil {
				return err
			}

			for _, assgin := range cmd.Edges.Assign {
				if assgin.Edges.User == nil {
					return fmt.Errorf("用户ID不能为空")
				}

				builder := tx.TaskAssign.Create()

				builder = builder.
					SetOperatorID(1).
					SetTaskID(save.ID).
					SetUserID(assgin.Edges.User.ID)

				if _, err := builder.Save(ctx); err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			log.Fatal(err)
		}

		return ec.JSON(http.StatusOK, save)
	})

	app.GET("/:id", func(ec echo.Context) error {
		ctx := ec.Get("context.Context").(context.Context)
		id, err := strconv.ParseInt(ec.Param("id"), 10, 32)
		if err != nil {
			return err
		}

		tasks, err := client.Task.Query().Where(task.ID(int(id))).WithStatus().WithAssign(func(query *ent.TaskAssignQuery) {
			query.WithUser()
		}).WithChild(func(query *ent.TaskQuery) {
			query.WithStatus()
		}).Only(ctx)
		if err != nil {
			return err
		}

		return ec.JSON(http.StatusOK, tasks)
	})

	app.PUT("/:id", func(ec echo.Context) error {
		ctx := ec.Get("context.Context").(context.Context)
		id, err := strconv.ParseInt(ec.Param("id"), 10, 32)
		if err != nil {
			return err
		}

		cmd := ent.Task{}
		if err := ec.Bind(&cmd); err != nil {
			return err
		}

		builder := client.Task.UpdateOneID(int(id))

		builder = builder.
			SetOperatorID(1).
			SetTitle(cmd.Title).
			SetContent(cmd.Content).
			SetManHour(cmd.ManHour).
			SetStatusID(cmd.Edges.Status.ID).
			SetNillableEndTime(cmd.EndTime).
			SetNillableStartTime(cmd.StartTime)

		if cmd.Edges.Parent == nil {
			builder = builder.SetNillableParentID(nil)
		} else {
			builder = builder.SetParentID(cmd.Edges.Parent.ID)
		}

		t, err := builder.Save(ctx)
		if err != nil {
			return err
		}

		return ec.JSON(http.StatusOK, t)
	})

	app.PATCH("/:id", func(ec echo.Context) error {
		ctx := ec.Get("context.Context").(context.Context)
		id, err := strconv.ParseInt(ec.Param("id"), 10, 32)
		if err != nil {
			return err
		}

		params := make(map[string]interface{})
		if err := ec.Bind(&params); err != nil {
			return err
		}

		builder := client.Task.UpdateOneID(int(id))

		for field, value := range params {
			switch field {
			case task.FieldTitle:
				builder = builder.SetTitle(value.(string))
			case task.FieldContent:
				builder = builder.SetContent(value.(string))
			case task.FieldManHour:
				builder = builder.SetManHour(int(value.(float64)))
			case task.FieldEndTime:
				t := time.Now()
				if err := t.UnmarshalText([]byte(value.(string))); err != nil {
					return err
				}
				builder = builder.SetEndTime(t)
			case task.FieldStartTime:
				t := time.Now()
				if err := t.UnmarshalText([]byte(value.(string))); err != nil {
					return err
				}
				builder = builder.SetStartTime(t)
			case task.StatusColumn:
				builder = builder.SetStatusID(int(value.(float64)))
			case task.OperatorColumn:
				builder = builder.SetParentID(int(value.(float64)))
			default:
				return fmt.Errorf("无法识别的字段 %v", field)
			}
		}

		t, err := builder.Save(ctx)
		if err != nil {
			return err
		}

		return ec.JSON(http.StatusOK, t)
	})

	return app
}
