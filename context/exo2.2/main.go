package main

import (
	"context"
	"fmt"
)

/**
	2.2: Hierarchical Context with Values
		Create a parent context with a key-value pair and a child context that overrides one of the values.
		Demonstrate how values are looked up in a hierarchical manner, with the child context overriding the parent's value.
**/

type contextKey string

type User struct{
	ID int64
}

func printContextValue(ctx context.Context){
	user, ok := ctx.Value(contextKey("user")).(User)
	if !ok{
		return 
	}
	role, ok := ctx.Value(contextKey("role")).(string)
	if !ok {
		return
	}

	fmt.Printf("UserID: %v, Role: %v\n", user.ID, role)
}



func main(){
	parentCtx := context.WithValue(context.Background(), contextKey("user"), User{ID: 123})
	parentCtx = context.WithValue(parentCtx,  contextKey("role"), "admin")

	fmt.Println("Parent context:")
	printContextValue(parentCtx)

	childContext := context.WithValue(parentCtx, contextKey("role"), "user")
	
	fmt.Println("child context:")
	printContextValue(childContext)

	anOtherChildContext := context.WithValue(parentCtx, contextKey("user"), User{ID: 555})
	
	fmt.Println("another child context:")
	printContextValue(anOtherChildContext)

}