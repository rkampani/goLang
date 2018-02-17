package modal

type Person struct {
	Fname   string   `json:"fname"`
	Lname   string   `json:"lname"`
	ID      int      `json:"id"`
	Address *Address `json:"address"`
}

// just to wire down the method
func (p *Person) ToString() string {
	return p.Fname + " " + p.Lname + "  " + p.Address.AddressLine1 + " " + p.Address.AddressLine2 + " " + p.Address.City + " " + p.Address.State + " " + p.Address.ZipCode + " "
}
