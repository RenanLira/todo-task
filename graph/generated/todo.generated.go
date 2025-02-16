// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"todo-tasks/graph/model"
	"todo-tasks/internal/domain/todos"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type MutationResolver interface {
	CreateTodo(ctx context.Context, input model.TodoInput) (*todos.Todo, error)
	UpdateTodo(ctx context.Context, id string, input model.TodoInput) (*todos.Todo, error)
	DeleteTodo(ctx context.Context, id string) (*todos.Todo, error)
}
type QueryResolver interface {
	Todos(ctx context.Context, page *model.PageInput) (*model.TodosResponse, error)
	TodoByID(ctx context.Context, id string) (*todos.Todo, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Mutation_createTodo_args(ctx context.Context, rawArgs map[string]any) (map[string]any, error) {
	var err error
	args := map[string]any{}
	arg0, err := ec.field_Mutation_createTodo_argsInput(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["input"] = arg0
	return args, nil
}
func (ec *executionContext) field_Mutation_createTodo_argsInput(
	ctx context.Context,
	rawArgs map[string]any,
) (model.TodoInput, error) {
	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
	if tmp, ok := rawArgs["input"]; ok {
		return ec.unmarshalNTodoInput2todoᚑtasksᚋgraphᚋmodelᚐTodoInput(ctx, tmp)
	}

	var zeroVal model.TodoInput
	return zeroVal, nil
}

func (ec *executionContext) field_Mutation_deleteTodo_args(ctx context.Context, rawArgs map[string]any) (map[string]any, error) {
	var err error
	args := map[string]any{}
	arg0, err := ec.field_Mutation_deleteTodo_argsID(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["id"] = arg0
	return args, nil
}
func (ec *executionContext) field_Mutation_deleteTodo_argsID(
	ctx context.Context,
	rawArgs map[string]any,
) (string, error) {
	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
	if tmp, ok := rawArgs["id"]; ok {
		return ec.unmarshalNID2string(ctx, tmp)
	}

	var zeroVal string
	return zeroVal, nil
}

func (ec *executionContext) field_Mutation_updateTodo_args(ctx context.Context, rawArgs map[string]any) (map[string]any, error) {
	var err error
	args := map[string]any{}
	arg0, err := ec.field_Mutation_updateTodo_argsID(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["id"] = arg0
	arg1, err := ec.field_Mutation_updateTodo_argsInput(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["input"] = arg1
	return args, nil
}
func (ec *executionContext) field_Mutation_updateTodo_argsID(
	ctx context.Context,
	rawArgs map[string]any,
) (string, error) {
	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
	if tmp, ok := rawArgs["id"]; ok {
		return ec.unmarshalNID2string(ctx, tmp)
	}

	var zeroVal string
	return zeroVal, nil
}

func (ec *executionContext) field_Mutation_updateTodo_argsInput(
	ctx context.Context,
	rawArgs map[string]any,
) (model.TodoInput, error) {
	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
	if tmp, ok := rawArgs["input"]; ok {
		return ec.unmarshalNTodoInput2todoᚑtasksᚋgraphᚋmodelᚐTodoInput(ctx, tmp)
	}

	var zeroVal model.TodoInput
	return zeroVal, nil
}

func (ec *executionContext) field_Query___type_args(ctx context.Context, rawArgs map[string]any) (map[string]any, error) {
	var err error
	args := map[string]any{}
	arg0, err := ec.field_Query___type_argsName(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["name"] = arg0
	return args, nil
}
func (ec *executionContext) field_Query___type_argsName(
	ctx context.Context,
	rawArgs map[string]any,
) (string, error) {
	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
	if tmp, ok := rawArgs["name"]; ok {
		return ec.unmarshalNString2string(ctx, tmp)
	}

	var zeroVal string
	return zeroVal, nil
}

func (ec *executionContext) field_Query_todoById_args(ctx context.Context, rawArgs map[string]any) (map[string]any, error) {
	var err error
	args := map[string]any{}
	arg0, err := ec.field_Query_todoById_argsID(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["id"] = arg0
	return args, nil
}
func (ec *executionContext) field_Query_todoById_argsID(
	ctx context.Context,
	rawArgs map[string]any,
) (string, error) {
	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
	if tmp, ok := rawArgs["id"]; ok {
		return ec.unmarshalNID2string(ctx, tmp)
	}

	var zeroVal string
	return zeroVal, nil
}

func (ec *executionContext) field_Query_todos_args(ctx context.Context, rawArgs map[string]any) (map[string]any, error) {
	var err error
	args := map[string]any{}
	arg0, err := ec.field_Query_todos_argsPage(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["page"] = arg0
	return args, nil
}
func (ec *executionContext) field_Query_todos_argsPage(
	ctx context.Context,
	rawArgs map[string]any,
) (*model.PageInput, error) {
	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("page"))
	if tmp, ok := rawArgs["page"]; ok {
		return ec.unmarshalOPageInput2ᚖtodoᚑtasksᚋgraphᚋmodelᚐPageInput(ctx, tmp)
	}

	var zeroVal *model.PageInput
	return zeroVal, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Mutation_createTodo(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_createTodo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().CreateTodo(rctx, fc.Args["input"].(model.TodoInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*todos.Todo)
	fc.Result = res
	return ec.marshalNTodo2ᚖtodoᚑtasksᚋinternalᚋdomainᚋtodosᚐTodo(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_createTodo(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Todo_id(ctx, field)
			case "text":
				return ec.fieldContext_Todo_text(ctx, field)
			case "done":
				return ec.fieldContext_Todo_done(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Todo", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_createTodo_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_updateTodo(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_updateTodo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().UpdateTodo(rctx, fc.Args["id"].(string), fc.Args["input"].(model.TodoInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*todos.Todo)
	fc.Result = res
	return ec.marshalNTodo2ᚖtodoᚑtasksᚋinternalᚋdomainᚋtodosᚐTodo(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_updateTodo(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Todo_id(ctx, field)
			case "text":
				return ec.fieldContext_Todo_text(ctx, field)
			case "done":
				return ec.fieldContext_Todo_done(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Todo", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_updateTodo_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_deleteTodo(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_deleteTodo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().DeleteTodo(rctx, fc.Args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*todos.Todo)
	fc.Result = res
	return ec.marshalNTodo2ᚖtodoᚑtasksᚋinternalᚋdomainᚋtodosᚐTodo(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_deleteTodo(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Todo_id(ctx, field)
			case "text":
				return ec.fieldContext_Todo_text(ctx, field)
			case "done":
				return ec.fieldContext_Todo_done(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Todo", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_deleteTodo_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Query_todos(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query_todos(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Todos(rctx, fc.Args["page"].(*model.PageInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.TodosResponse)
	fc.Result = res
	return ec.marshalNTodosResponse2ᚖtodoᚑtasksᚋgraphᚋmodelᚐTodosResponse(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query_todos(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "todos":
				return ec.fieldContext_TodosResponse_todos(ctx, field)
			case "pageInfo":
				return ec.fieldContext_TodosResponse_pageInfo(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type TodosResponse", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Query_todos_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Query_todoById(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query_todoById(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().TodoByID(rctx, fc.Args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*todos.Todo)
	fc.Result = res
	return ec.marshalOTodo2ᚖtodoᚑtasksᚋinternalᚋdomainᚋtodosᚐTodo(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query_todoById(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Todo_id(ctx, field)
			case "text":
				return ec.fieldContext_Todo_text(ctx, field)
			case "done":
				return ec.fieldContext_Todo_done(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Todo", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Query_todoById_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query___type(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectType(fc.Args["name"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query___type(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "kind":
				return ec.fieldContext___Type_kind(ctx, field)
			case "name":
				return ec.fieldContext___Type_name(ctx, field)
			case "description":
				return ec.fieldContext___Type_description(ctx, field)
			case "fields":
				return ec.fieldContext___Type_fields(ctx, field)
			case "interfaces":
				return ec.fieldContext___Type_interfaces(ctx, field)
			case "possibleTypes":
				return ec.fieldContext___Type_possibleTypes(ctx, field)
			case "enumValues":
				return ec.fieldContext___Type_enumValues(ctx, field)
			case "inputFields":
				return ec.fieldContext___Type_inputFields(ctx, field)
			case "ofType":
				return ec.fieldContext___Type_ofType(ctx, field)
			case "specifiedByURL":
				return ec.fieldContext___Type_specifiedByURL(ctx, field)
			case "isOneOf":
				return ec.fieldContext___Type_isOneOf(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Type", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Query___type_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query___schema(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectSchema()
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	fc.Result = res
	return ec.marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query___schema(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "description":
				return ec.fieldContext___Schema_description(ctx, field)
			case "types":
				return ec.fieldContext___Schema_types(ctx, field)
			case "queryType":
				return ec.fieldContext___Schema_queryType(ctx, field)
			case "mutationType":
				return ec.fieldContext___Schema_mutationType(ctx, field)
			case "subscriptionType":
				return ec.fieldContext___Schema_subscriptionType(ctx, field)
			case "directives":
				return ec.fieldContext___Schema_directives(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Schema", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Todo_id(ctx context.Context, field graphql.CollectedField, obj *todos.Todo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Todo_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Todo_id(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Todo",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Todo_text(ctx context.Context, field graphql.CollectedField, obj *todos.Todo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Todo_text(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Text, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Todo_text(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Todo",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Todo_done(ctx context.Context, field graphql.CollectedField, obj *todos.Todo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Todo_done(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Done, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Todo_done(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Todo",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _TodosResponse_todos(ctx context.Context, field graphql.CollectedField, obj *model.TodosResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TodosResponse_todos(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Todos, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*todos.Todo)
	fc.Result = res
	return ec.marshalNTodo2ᚕᚖtodoᚑtasksᚋinternalᚋdomainᚋtodosᚐTodoᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_TodosResponse_todos(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TodosResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Todo_id(ctx, field)
			case "text":
				return ec.fieldContext_Todo_text(ctx, field)
			case "done":
				return ec.fieldContext_Todo_done(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Todo", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _TodosResponse_pageInfo(ctx context.Context, field graphql.CollectedField, obj *model.TodosResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TodosResponse_pageInfo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PageInfo, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.PageInfo)
	fc.Result = res
	return ec.marshalOPageInfo2ᚖtodoᚑtasksᚋgraphᚋmodelᚐPageInfo(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_TodosResponse_pageInfo(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TodosResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "hasNextPage":
				return ec.fieldContext_PageInfo_hasNextPage(ctx, field)
			case "hasPreviousPage":
				return ec.fieldContext_PageInfo_hasPreviousPage(ctx, field)
			case "quantity":
				return ec.fieldContext_PageInfo_quantity(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type PageInfo", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputTodoInput(ctx context.Context, obj any) (model.TodoInput, error) {
	var it model.TodoInput
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"text"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "text":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("text"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Text = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var mutationImplementors = []string{"Mutation"}

func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mutationImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Mutation",
	})

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		innerCtx := graphql.WithRootFieldContext(ctx, &graphql.RootFieldContext{
			Object: field.Name,
			Field:  field,
		})

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "createTodo":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_createTodo(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "updateTodo":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_updateTodo(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "deleteTodo":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_deleteTodo(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var queryImplementors = []string{"Query"}

func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, queryImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Query",
	})

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		innerCtx := graphql.WithRootFieldContext(ctx, &graphql.RootFieldContext{
			Object: field.Name,
			Field:  field,
		})

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "todos":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_todos(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx,
					func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return rrm(innerCtx) })
		case "todoById":
			field := field

			innerFunc := func(ctx context.Context, _ *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_todoById(ctx, field)
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx,
					func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return rrm(innerCtx) })
		case "__type":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Query___type(ctx, field)
			})
		case "__schema":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Query___schema(ctx, field)
			})
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var todoImplementors = []string{"Todo"}

func (ec *executionContext) _Todo(ctx context.Context, sel ast.SelectionSet, obj *todos.Todo) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, todoImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Todo")
		case "id":
			out.Values[i] = ec._Todo_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "text":
			out.Values[i] = ec._Todo_text(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "done":
			out.Values[i] = ec._Todo_done(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var todosResponseImplementors = []string{"TodosResponse"}

func (ec *executionContext) _TodosResponse(ctx context.Context, sel ast.SelectionSet, obj *model.TodosResponse) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, todosResponseImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("TodosResponse")
		case "todos":
			out.Values[i] = ec._TodosResponse_todos(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "pageInfo":
			out.Values[i] = ec._TodosResponse_pageInfo(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNTodo2todoᚑtasksᚋinternalᚋdomainᚋtodosᚐTodo(ctx context.Context, sel ast.SelectionSet, v todos.Todo) graphql.Marshaler {
	return ec._Todo(ctx, sel, &v)
}

func (ec *executionContext) marshalNTodo2ᚕᚖtodoᚑtasksᚋinternalᚋdomainᚋtodosᚐTodoᚄ(ctx context.Context, sel ast.SelectionSet, v []*todos.Todo) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNTodo2ᚖtodoᚑtasksᚋinternalᚋdomainᚋtodosᚐTodo(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNTodo2ᚖtodoᚑtasksᚋinternalᚋdomainᚋtodosᚐTodo(ctx context.Context, sel ast.SelectionSet, v *todos.Todo) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Todo(ctx, sel, v)
}

func (ec *executionContext) unmarshalNTodoInput2todoᚑtasksᚋgraphᚋmodelᚐTodoInput(ctx context.Context, v any) (model.TodoInput, error) {
	res, err := ec.unmarshalInputTodoInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNTodosResponse2todoᚑtasksᚋgraphᚋmodelᚐTodosResponse(ctx context.Context, sel ast.SelectionSet, v model.TodosResponse) graphql.Marshaler {
	return ec._TodosResponse(ctx, sel, &v)
}

func (ec *executionContext) marshalNTodosResponse2ᚖtodoᚑtasksᚋgraphᚋmodelᚐTodosResponse(ctx context.Context, sel ast.SelectionSet, v *model.TodosResponse) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._TodosResponse(ctx, sel, v)
}

func (ec *executionContext) marshalOTodo2ᚖtodoᚑtasksᚋinternalᚋdomainᚋtodosᚐTodo(ctx context.Context, sel ast.SelectionSet, v *todos.Todo) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Todo(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
