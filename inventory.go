package main

// Product  is an object for an inventory item in the storeItemList
// you can add an additional attribute below like the example, which
// is commented out.
type Product struct {
	ID      int
	SKU     string
	ImgPath string
	Name    string
	Desc    string
	Cost    int
	// Weight float32
}

// ProductInventory is a list of all of the products in your store
type ProductInventory []Product

// ProductOrder is a single SKU and a single quantity that is created
// when a user browses to an item, then adds 1 or more qty of that item
// to their shopping cart.
type ProductOrder struct {
	ProductID string
	Quantity  int
}

// ShoppingCart consists of those instances of a user adding those
// items to cart. This will make tallying the total for an order
// easy later.
type ShoppingCart []ProductOrder

// This is your inventory. The structure of the product objects should
// always match the structure of the Product struct.
var (
	StoreInventory = ProductInventory{
		Product{
			0,
			"CB-001",
			"https://picsum.photos/id/10/400/300",
			"Cheap Beer",
			"It will be cold and bubbly, but nothing more.",
			10000,
			// 10.00,
		},
		Product{
			1,
			"NB-001",
			"https://picsum.photos/id/200/400/300",
			"Nice Beer",
			"Quenching and refreshing, a pleasure to the tastebuds..",
			25000,
			// 10.00,
		},
		Product{
			2,
			"6P-001",
			"https://picsum.photos/id/30/400/300",
			"Sixer",
			"Cheap beer or good beer, six of them will get you where you need to be.",
			100000,
			// 60.00,
		},
		Product{
			3,
			"FL-001",
			"https://picsum.photos/id/400/400/300",
			"Footlong",
			"12 beers. Not enough and too much, all at the same time. ",
			200000,
			// 120.00,
		},
	}
)
