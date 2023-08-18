package entity

type Operation struct {
	Operator string `json:"operator"`
	IP       string `json:"ip"`
	Agent    string `json:"agent"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	Query    string `json:"query"`
	Body     string `json:"body"`
	Status   int    `json:"status"`
	Resp     string `json:"resp"`
}
