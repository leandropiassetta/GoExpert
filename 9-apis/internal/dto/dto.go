package dto

// dto = data transfer object
// it is a struct that we use to transfer data between layers
// binding data from the request to a struct is a good practice

// bind -> is the process of getting the data from the request and putting it into a struct
type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
