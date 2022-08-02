package main

import (
"fmt"
)
/*
Problem: How can you implement such a license mechanism so that no change to either in License package or client 
machines would be necessary. Design and implement your solution using proxy design pattern. Implement your own error mechanism to represent 
the case where no license is available.
*/

type Proxy interface {
	Proxy()
}

type Product struct{}

func (p *Product) Proxy() {
	fmt.Println("Keeps working.Haven't reached the limit.")
}

type User struct {
	Num0fUsers int
}

type productProxy struct {
	prod Product
	user *User
}

func main() {
	prod := ProductProxy(&User{45})
	prod.Proxy()
}

func (p *productProxy) Proxy()  {
	
	for p.user.Num0fUsers = 45 ; p.user.Num0fUsers <=50 ;p.user.Num0fUsers ++ {
		if p.user.Num0fUsers < 50 {
			p.prod.Proxy()
		} else {
			fmt.Println("Reached the limit.")
		}
	}
	
}

func ProductProxy(user *User) *productProxy {
	return &productProxy{Product{}, user}
}

