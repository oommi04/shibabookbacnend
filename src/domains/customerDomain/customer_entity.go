package customerDomain

type Customer struct {
	ID                   string `json:"_id,omitempty"`
	Name                 string `json:"name,omitempty"`
	LastName             string `json:"lastName,omitempty"`
	IdentificationNumber string `json:"identificationNumber,omitempty"`
	Mail                 string `json:"mail,omitempty"`
	Point                int    `json:"point,omitempty"`
}
