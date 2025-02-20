package authorization

import (
	"errors"
)

func UserTodoAllow(userId string, action, owner string) error {
	e := GetEnforcer()
	allowed, err := e.Enforce(userId, "todo", action, owner)
	if err != nil {
		return err
	}
	if !allowed {
		return errors.New("forbidden")
	}

	return nil
}
