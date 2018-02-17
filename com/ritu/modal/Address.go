package modal

type Address struct {
	City         string `json:"city"`
	ZipCode      string `json:"zipcode"`
	State        string `json:"state"`
	AddressLine1 string `json:"line1"`
	AddressLine2 string `json:"line2"`
}
