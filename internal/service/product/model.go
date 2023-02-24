package product

var allProducts = []Product{
	{Id: 1, Title: "smartphone"},
	{Id: 2, Title: "tv"},
	{Id: 3, Title: "earphones"},
}

type Product struct {
	Id    int
	Title string
}
