package dto

import "github.com/google/uuid"

type BasicAnalyticsSalesData struct {
}

type OrderCreated struct {
	ID            uuid.UUID
	CategoryOrder CategoryOrder
	ProductsData  ProductsData
	SaleMediaData string
	Total         float64
}

type CategoryOrder struct {
	ID     uuid.UUID
	Gender string
	Age    int16
	Work   bool
}

type ProductsData struct {
	ID          uuid.UUID
	Product     string
	Category    string
	SubCategory string
	Material    string
	Quantity    int16
	Price       float64
}

type SalesMaketingData struct {
	ID               uuid.UUID
	ProductId        uuid.UUID
	SocialMedia      string
	Gender           string
	Age              string
	Work             bool
	TotalSocialMedia map[string]int
	TotalGender      map[string]int
	TotalAge         map[string]int
}
