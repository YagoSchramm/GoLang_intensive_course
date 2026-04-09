package model

type CreateNotebookInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Notebook struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
