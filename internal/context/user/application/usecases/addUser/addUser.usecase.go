package adduser

import (
	"context"
	"fmt"
)

func (a *AddUserUsecase) AddUser() error {
	res, _ := a.userService.GetUserByID(context.Background(), "67e32c16e9dead498d27f593")
	fmt.Println(res)
	fmt.Println("AddUserUsecase.AddUser")
	return nil
}
