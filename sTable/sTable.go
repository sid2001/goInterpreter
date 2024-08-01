package sTable

import (
	"github.com/goInterpreter/token"
)

type VariableAttributes struct {
	Addr	int
	Type	token.DataType
}

type STableData struct {
	Name	string
	Type	string
	Size	int
	Addr	int
}

type STable []STableData

var Variables  = map[string]VariableAttributes{}


