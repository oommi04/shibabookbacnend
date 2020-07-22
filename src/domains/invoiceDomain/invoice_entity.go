package invoiceDomain

type Invoice struct {
	ID   string `json:"_id,omitempty"`
	Path string `json:"path,omitempty"`
	Mail string `json:"mail,omitempty"`
}
