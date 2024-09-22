package main

import (
	"context"
	"errors"
	"fmt"
)

/**
	2.1: Context with Values
		Write a function that creates a context.Context with a key-value pair (e.g., user ID or request ID).
		Pass this context down to multiple functions, and in each function, retrieve and print the stored value.
		Explain why context.WithValue should not be used for passing optional parameters.
**/


type User struct {
	ID int64
}

type userKey string

func createContext(id int64) context.Context{
	user := User{ID: id}
	key := userKey("user")
	userContext := context.WithValue(context.Background(), key, user)

	return userContext
}

func retrieveUser1(ctx context.Context) (User, error){
	key := userKey("user")
	user, ok:= ctx.Value(key).(User)
	if !ok{
		return User{}, errors.New("user not found")
	}

	if user.ID == 0 {
		return User{}, errors.New("id cannot be zero")
	}

	fmt.Println("func 1. user:", user.ID)

	return user, nil
}

func retrieveUser2(ctx context.Context) (User, error){
	key := userKey("user")
	user, ok:= ctx.Value(key).(User)
	if !ok {
		return User{}, errors.New("user not found")
	}

	if user.ID == 0 {
		return User{}, errors.New("id cannot be zero")
	}

	fmt.Println("func 2. user:", user.ID)

	return user, nil
}


func main(){
	userContext1 := createContext(0)
	userContext2 := createContext(2)

	_, err := retrieveUser1(userContext1)
	if err != nil {
		fmt.Println(err)
	}
	_, err = retrieveUser2(userContext2)
	if err != nil {
		fmt.Println(err)
	}
	
}