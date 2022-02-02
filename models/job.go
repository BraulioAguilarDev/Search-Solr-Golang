package models

type ResponseHeader struct {
	Status int `json:"status"`
	QTime  int `json:"QTime"`
	Params struct {
		Q      string `json:"q"`
		Indent string `json:"indent"`
		QOp    string `json:"q.op"`
	} `json:"params"`
}

type Docs struct {
	ID       string   `json:"id"`
	Title    []string `json:"title"`
	Text     []string `json:"text"`
	Salary   []int    `json:"salary"`
	Locality []string `json:"locality"`
	Category []string `json:"category"`
	Version  int64    `json:"_version_"`
}

type Response struct {
	NumFound      int    `json:"numFound"`
	Start         int    `json:"start"`
	NumFoundExact bool   `json:"numFoundExact"`
	Docs          []Docs `json:"docs"`
}

// Represents main Model that we receive from Solr storage
type SolrResponse struct {
	ResponseHeader ResponseHeader `json:"responseHeader"`
	Response       Response       `json:"response"`
}
