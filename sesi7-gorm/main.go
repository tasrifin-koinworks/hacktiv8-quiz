package main

import (
	"fmt"
	"sesi7-gorm/database"
	"sesi7-gorm/repository"
)

func main() {
	db := database.StartDB()

	userRepo := repository.NewUserRepo(db)

	//CREATE NEW USER
	// user := models.User{
	// 	Email: "coba3@gmail.com",
	// }

	// err := userRepo.CreateUser(&user)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("Create user success")

	//UPDATE USER
	// newEmail := "newEmailcoba@gmail.com"
	// _, errs := userRepo.UpdateUserByID(4, newEmail)
	// if errs != nil {
	// 	fmt.Println(errs.Error())
	// 	return
	// }

	//DELETE USER
	// deleteErr := userRepo.DeleteUser(4)
	// if deleteErr != nil {
	// 	fmt.Println(deleteErr.Error())
	// 	return
	// }
	// fmt.Println("Delete user success")

	employees, err := userRepo.GetAllUsers()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for v, emp := range *employees {
		fmt.Println("User ke \t:", v+1)
		emp.Print()
	}

	emp, err := userRepo.GetUserByID(3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("***********************************")
	fmt.Println("Search ID \t:", 3)
	emp.Print()
	fmt.Println("***********************************")

	// ProcessProduct()

}

func ProcessProduct() {
	db := database.StartDB()
	productRepo := repository.NewProductRepo(db)

	//CREATE NEW PRODUCT
	// product := models.Product{
	// 	Name:   "Brownies",
	// 	Brand:  "Amanda",
	// 	UserID: 3,
	// }

	// err := productRepo.CreateProduct(&product)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("Create Product success")

	//UPDATE PRODUCT BY ID
	newProductName := "Mentos"
	newProductBrand := "Mentos"
	_, err := productRepo.UpdateProductByID(3, newProductName, newProductBrand)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Update Product success")

	//DELETE PRODUCT
	deleteProductErr := productRepo.DeleteProduct(5)
	if deleteProductErr != nil {
		fmt.Println(deleteProductErr.Error())
		return
	}
	fmt.Println("Delete Product success")

	products, err := productRepo.GetAllProducts()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for v, p := range *products {
		fmt.Println("Product ke \t:", v+1)
		p.PrintProduct()
	}

	product, err := productRepo.GetProductByID(3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("*******************************")
	fmt.Println("Search ID \t:", 3)
	product.PrintProduct()
	fmt.Println("*******************************")

}
