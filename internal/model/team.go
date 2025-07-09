package model

type Project struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type Team struct {
	Name     string    `json:"name"`
	Leader   bool      `json:"leader"`
	Projects []Project `json:"projects"`
}
