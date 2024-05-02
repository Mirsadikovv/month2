package model

type Products struct {
	Id               int
	Product_name     string
	Product_price    int
	Product_quantity int
}

type Categories struct {
	Id            int
	Category_name string
	Product_id    int
}

type Prod_and_cat struct {
	Id                int
	Product_name      string
	Product_price     int
	Product_quantity  int
	Product_categorie string
}
