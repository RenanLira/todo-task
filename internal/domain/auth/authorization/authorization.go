package authorization

import "github.com/casbin/casbin/v2"

var (
	enforcer *casbin.Enforcer
)

func init() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		panic(err)
	}

	enforcer = e
}

func GetEnforcer() *casbin.Enforcer {
	return enforcer
}
