package any

type Cart struct {
	Id     int
	ItemId int
	UserId string
}

type Items struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Cost int    `json:"cost"`
	Stok bool   `json:"stok"`
}

type Order struct {
	Id     int    `json:"id"`
	CartId int    `json:"cart_id"`
	Adress string `json:"adr"`
}
