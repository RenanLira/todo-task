// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"bytes"
	"context"
	"errors"
	"sync/atomic"
	"todo-tasks/graph/model"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		schema:     cfg.Schema,
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Schema     *ast.Schema
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
	User() UserResolver
}

type DirectiveRoot struct {
	Authenticated func(ctx context.Context, obj any, next graphql.Resolver) (res any, err error)
}

type ComplexityRoot struct {
	Mutation struct {
		CreateTodo func(childComplexity int, input model.TodoInput) int
		CreateUser func(childComplexity int, username string, email string, password string) int
		DeleteTodo func(childComplexity int, id string) int
		LoginUser  func(childComplexity int, email string, password string) int
		UpdateTodo func(childComplexity int, id string, update model.UpdateTodo) int
	}

	PageInfo struct {
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		Quantity        func(childComplexity int) int
	}

	Query struct {
		MyTodos  func(childComplexity int, page *model.PageInput) int
		TodoByID func(childComplexity int, id string) int
		Todos    func(childComplexity int, page *model.PageInput) int
	}

	Todo struct {
		Done func(childComplexity int) int
		ID   func(childComplexity int) int
		Text func(childComplexity int) int
	}

	TodosResponse struct {
		PageInfo func(childComplexity int) int
		Todos    func(childComplexity int) int
	}

	User struct {
		Email func(childComplexity int) int
		ID    func(childComplexity int) int
		Name  func(childComplexity int) int
	}
}

type executableSchema struct {
	schema     *ast.Schema
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	if e.schema != nil {
		return e.schema
	}
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]any) (int, bool) {
	ec := executionContext{nil, e, 0, 0, nil}
	_ = ec
	switch typeName + "." + field {

	case "Mutation.createTodo":
		if e.complexity.Mutation.CreateTodo == nil {
			break
		}

		args, err := ec.field_Mutation_createTodo_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateTodo(childComplexity, args["input"].(model.TodoInput)), true

	case "Mutation.createUser":
		if e.complexity.Mutation.CreateUser == nil {
			break
		}

		args, err := ec.field_Mutation_createUser_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateUser(childComplexity, args["username"].(string), args["email"].(string), args["password"].(string)), true

	case "Mutation.deleteTodo":
		if e.complexity.Mutation.DeleteTodo == nil {
			break
		}

		args, err := ec.field_Mutation_deleteTodo_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteTodo(childComplexity, args["id"].(string)), true

	case "Mutation.loginUser":
		if e.complexity.Mutation.LoginUser == nil {
			break
		}

		args, err := ec.field_Mutation_loginUser_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.LoginUser(childComplexity, args["email"].(string), args["password"].(string)), true

	case "Mutation.updateTodo":
		if e.complexity.Mutation.UpdateTodo == nil {
			break
		}

		args, err := ec.field_Mutation_updateTodo_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpdateTodo(childComplexity, args["id"].(string), args["update"].(model.UpdateTodo)), true

	case "PageInfo.hasNextPage":
		if e.complexity.PageInfo.HasNextPage == nil {
			break
		}

		return e.complexity.PageInfo.HasNextPage(childComplexity), true

	case "PageInfo.hasPreviousPage":
		if e.complexity.PageInfo.HasPreviousPage == nil {
			break
		}

		return e.complexity.PageInfo.HasPreviousPage(childComplexity), true

	case "PageInfo.quantity":
		if e.complexity.PageInfo.Quantity == nil {
			break
		}

		return e.complexity.PageInfo.Quantity(childComplexity), true

	case "Query.myTodos":
		if e.complexity.Query.MyTodos == nil {
			break
		}

		args, err := ec.field_Query_myTodos_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.MyTodos(childComplexity, args["page"].(*model.PageInput)), true

	case "Query.todoById":
		if e.complexity.Query.TodoByID == nil {
			break
		}

		args, err := ec.field_Query_todoById_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.TodoByID(childComplexity, args["id"].(string)), true

	case "Query.todos":
		if e.complexity.Query.Todos == nil {
			break
		}

		args, err := ec.field_Query_todos_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Todos(childComplexity, args["page"].(*model.PageInput)), true

	case "Todo.done":
		if e.complexity.Todo.Done == nil {
			break
		}

		return e.complexity.Todo.Done(childComplexity), true

	case "Todo.id":
		if e.complexity.Todo.ID == nil {
			break
		}

		return e.complexity.Todo.ID(childComplexity), true

	case "Todo.text":
		if e.complexity.Todo.Text == nil {
			break
		}

		return e.complexity.Todo.Text(childComplexity), true

	case "TodosResponse.pageInfo":
		if e.complexity.TodosResponse.PageInfo == nil {
			break
		}

		return e.complexity.TodosResponse.PageInfo(childComplexity), true

	case "TodosResponse.todos":
		if e.complexity.TodosResponse.Todos == nil {
			break
		}

		return e.complexity.TodosResponse.Todos(childComplexity), true

	case "User.email":
		if e.complexity.User.Email == nil {
			break
		}

		return e.complexity.User.Email(childComplexity), true

	case "User.id":
		if e.complexity.User.ID == nil {
			break
		}

		return e.complexity.User.ID(childComplexity), true

	case "User.name":
		if e.complexity.User.Name == nil {
			break
		}

		return e.complexity.User.Name(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	opCtx := graphql.GetOperationContext(ctx)
	ec := executionContext{opCtx, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputPageInput,
		ec.unmarshalInputTodoInput,
		ec.unmarshalInputUpdateTodo,
	)
	first := true

	switch opCtx.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			var response graphql.Response
			var data graphql.Marshaler
			if first {
				first = false
				ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
				data = ec._queryMiddleware(ctx, opCtx.Operation, func(ctx context.Context) (any, error) {
					return ec._Query(ctx, opCtx.Operation.SelectionSet), nil
				})
			} else {
				if atomic.LoadInt32(&ec.pendingDeferred) > 0 {
					result := <-ec.deferredResults
					atomic.AddInt32(&ec.pendingDeferred, -1)
					data = result.Result
					response.Path = result.Path
					response.Label = result.Label
					response.Errors = result.Errors
				} else {
					return nil
				}
			}
			var buf bytes.Buffer
			data.MarshalGQL(&buf)
			response.Data = buf.Bytes()
			if atomic.LoadInt32(&ec.deferred) > 0 {
				hasNext := atomic.LoadInt32(&ec.pendingDeferred) > 0
				response.HasNext = &hasNext
			}

			return &response
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, opCtx.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
	deferred        int32
	pendingDeferred int32
	deferredResults chan graphql.DeferredResult
}

func (ec *executionContext) processDeferredGroup(dg graphql.DeferredGroup) {
	atomic.AddInt32(&ec.pendingDeferred, 1)
	go func() {
		ctx := graphql.WithFreshResponseContext(dg.Context)
		dg.FieldSet.Dispatch(ctx)
		ds := graphql.DeferredResult{
			Path:   dg.Path,
			Label:  dg.Label,
			Result: dg.FieldSet,
			Errors: graphql.GetErrors(ctx),
		}
		// null fields should bubble up
		if dg.FieldSet.Invalids > 0 {
			ds.Result = graphql.Null
		}
		ec.deferredResults <- ds
	}()
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(ec.Schema()), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(ec.Schema(), ec.Schema().Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../schemas/page.graphqls", Input: `input PageInput {
  limit: Int
  page: Int
  search: String
}

type PageInfo {
  hasNextPage: Boolean!
  hasPreviousPage: Boolean!
  quantity: Int!
}
`, BuiltIn: false},
	{Name: "../schemas/todo.graphqls", Input: `directive @authenticated on OBJECT | FIELD_DEFINITION | QUERY

type Todo {
  id: ID!
  text: String!
  done: Boolean!
}

input UpdateTodo {
  text: String
  done: Boolean
}

input TodoInput {
  text: String!
}

type TodosResponse {
  todos: [Todo!]!
  pageInfo: PageInfo
}

type Query {
  todos(page: PageInput): TodosResponse! @authenticated
  myTodos(page: PageInput): TodosResponse! @authenticated
  todoById(id: ID!): Todo @authenticated
}

type Mutation {
  createTodo(input: TodoInput!): Todo! @authenticated
  updateTodo(id: ID!, update: UpdateTodo!): Todo! @authenticated
  deleteTodo(id: ID!): Todo! @authenticated
}
`, BuiltIn: false},
	{Name: "../schemas/user.graphqls", Input: `type User {
    id: ID!
    name: String!
    email: String!
}


extend type Mutation {
    createUser(username: String!, email: String!, password: String!): User!
    loginUser(email: String!, password: String!): String!
}`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
