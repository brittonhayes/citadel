package risk

type Risk struct {
	ID              string `json:"id"`
	AssociatedAudit string `json:"associated_audit"`
	Description     string `json:"description"`
	Owner           string `json:"owner"`
	Severity        string `json:"severity"`
}
