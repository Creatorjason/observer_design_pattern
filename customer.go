package main

import "fmt"

type customer struct{
	id string
}

func (c *customer) getID() string{
	return c.id
}

func (c *customer) update(name string){
	fmt.Printf("Sending email to %s about %s\n ", c.id, name)
}