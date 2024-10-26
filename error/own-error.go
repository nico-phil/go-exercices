package main

import (
	"errors"
	"fmt"
)


type Status int

const (
	InvalidLogin Status = iota + 1
	Notfound
)

type StatusErr struct {
	Status Status
	Message string
}

func(se StatusErr) Error() string{
	return se.Message
}

func LoginAndGetData(uuid, pwd, file string)([]byte, error){
	token, err := Login(uuid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status: InvalidLogin,
			Message: fmt.Sprintf("invalid cresentials in login"),
		}
	}

	data, err := getData(token, file)
	if err != nil {
		return nil, StatusErr{
			Status: Notfound,
			Message: "file not found",
		}
	}

	return data, nil
}

func Login(uuid, pws string)(string, error){
	return "", errors.New("error loging")
	// return "1223", nil
}

func getData(token, file string)([]byte, error){
	return []byte("this is a bunch of data"), nil
}

func OwnError(){




}