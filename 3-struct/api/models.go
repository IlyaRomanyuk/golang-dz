package api

type SnippetMeta struct {
	Name string `json:"name"`
}
type BinItem struct {
	Private     bool        `json:"private"`
	SnippetMeta SnippetMeta `json:"snippetMeta"`
	Record      string      `json:"record"`
	CreatedAt   string      `json:"createdAt"`
}

type MetaData struct {
	Metadata MetaDataItem `json:"metadata"`
}
type MetaDataItem struct {
	Private   bool   `json:"private"`
	Name      string `json:"name"`
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}
