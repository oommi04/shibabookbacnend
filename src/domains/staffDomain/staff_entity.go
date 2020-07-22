package staffDomain

type Staff struct {
	ID                   string `json:"_id,omitempty"`
	Name                 string `json:"name,omitempty"`
	LastName             string `json:"lastName,omitempty"`
	IdentificationNumber string `json:"identificationNumber,omitempty"`
	Role                 string `json:"role,omitempty"`
	Status               string `json:"status,omitempty"`
}
