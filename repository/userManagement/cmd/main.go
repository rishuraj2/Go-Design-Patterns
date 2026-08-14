package main

import (
	"fmt"
	"usermanagement/database"
	"usermanagement/internal/domain"
)

func main() {
	repo := database.NewInMemoryStore()
	userService := domain.NewUserService(repo)

	userService.RegisterUser("abc", "abc@gmail.com")
	userService.RegisterUser("def", "def@gmail.com")

	fmt.Println(userService.GetAllUsers())
	fmt.Println(userService.FindUser("abcabc@gmail.com"))
}
