package main

import (
	cliente "app3/internal"
	"errors"
	"fmt"
	"io"
	"os"
)

var clientes []cliente.Cliente = []cliente.Cliente{
	cliente.Cliente{
		File:  "customer.txt",
		Name:  "Juan",
		ID:    1,
		Phone: 123456789,
		Home:  "Calle 123",
	},
}

var counter int = 0

func customerExist(customer cliente.Cliente) bool {
	//check if customer exist in clientes array
	for _, c := range clientes {
		if c.ID == customer.ID {
			return true
		}
	}
	return false
}

func validateInfo(file string, name string, id int, phone int, home string) (bool, error) {
	if file == "" {
		return false, errors.New("Error: file is empty")
	} else if name == "" {
		return false, errors.New("Error: name is empty")
	} else if id == 0 {
		return false, errors.New("Error: id is empty")
	} else if phone == 0 {
		return false, errors.New("Error: phone is empty")
	} else if home == "" {
		return false, errors.New("Error: home is empty")
	}
	return true, nil
}

func createCustomer(file string, name string, id int, phone int, home string) cliente.Cliente {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			counter++
		}
	}()

	customer := cliente.Cliente{
		File:  file,
		Name:  name,
		ID:    id,
		Phone: phone,
		Home:  home,
	}

	ok, err := validateInfo(file, name, id, phone, home)

	if !ok {
		panic(err)
	}

	if customerExist(customer) {
		panic("Error: client already exists")
	}

	clientes = append(clientes, customer)
	return customer
}

func main() {

	// customer := cliente.Cliente{
	// 	File:  "customer.txt",
	// 	Name:  "Juan",
	// 	ID:    1,
	// 	Phone: 123456789,
	// 	Home:  "Calle 123",
	// }
	customer := cliente.Cliente{
		File:  "",
		Name:  "Juan",
		ID:    1,
		Phone: 123456789,
		Home:  "Calle 123",
	}

	createCustomer(customer.File, customer.Name, customer.ID, customer.Phone, customer.Home)
	// createCustomer(customer.File, customer.Name, customer.ID, customer.Phone, customer.Home)

	fmt.Println("End of execution")
	if counter > 1 {
		fmt.Println("Several errors were detected at runtime")
	}

}

func ReadFile(path string) string {
	file, err := os.Open(path)

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	// read bytes from file
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic("An error occurred while reading the file")
	}

	defer file.Close()
	return string(bytes)
}
