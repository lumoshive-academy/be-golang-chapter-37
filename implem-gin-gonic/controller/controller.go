package controller

import "gorm.io/gorm"

type Controller struct {
	Product  ProductController
	Customer CustomerController
}

func NewController(db *gorm.DB) Controller {
	return Controller{
		Product:  *NewProductController(db),
		Customer: *NewCustomerController(db),
	}
}
