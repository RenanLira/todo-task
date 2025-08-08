package customvalidators

import (
	"todo-tasks/internal/infra/database"
	"todo-tasks/internal/utils"

	"github.com/go-playground/validator/v10"
)

func IsUniqueField(field validator.FieldLevel) bool {
	var db = database.NewDB()

	var count int64

	t := field.Top().Type()
	tableName := utils.GetPackageName(t)

	db.Table(tableName).Where(field.FieldName()+" = ?", field.Field().String()).Count(&count)

	return count == 0
}
