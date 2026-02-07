package models

type MenuItem struct {
	ID string
	Name string
	Description string
	Price float32
}

type Vendor struct {
	ID string
	Name string
	Description string
	Menu []MenuItem
	OwnerIDs []string
}

var vendors = []Vendor{}

func (v Vendor) Save() {
	vendors = append(vendors, v)
}

func GetVendors() []Vendor {
	return vendors
}