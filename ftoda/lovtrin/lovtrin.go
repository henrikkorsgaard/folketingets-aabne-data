package lovtrin

import (
	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"github.com/henrikkorsgaard/folketingets-aabne-data/repository"
)

type LovtrinType int

const (
	Fremsættelse      LovtrinType = iota
	FørsteBehandling  LovtrinType
	Udvalgsbehandling LovtrinType
	AndenBehandling   LovtrinType
	TredjeBehandling  LovtrinType
	NoClue            LovtrinType //we use this until we understand the mapping
)

var lovtrinTypeName = map[LovtrinType]string{
	Fremsættelse:      "Fremsættelse af forslag",
	FørsteBehandling:  "1. Behandling",
	Udvalgsbehandling: "Udvalgsbehandling",
	AndenBehandling:   "2. Behandling",
	TredjeBehandling:  "3. Behandling",
	NoClue: "No Clue"
}

func (lt LovtrinType) String() string {
	return lovtrinTypeName[lt]
}

func SagstrintypeToLovtrinType(sagtrinstype ftoda.Sagstrinstype) LovtrinType {
	// see data here: https://oda.ft.dk/api/Sagstrinstype. Should properly try to handle this dynamically. This is a critical area for error 
	switch sagtrinstype.Id {
	//Fremsættelse
	case 6, 20, 31, 32, 39, 77: 
		return Fremsættelse
	// Første behandling
	case 12, 23, 87:
		return FørsteBehandling
	default:
		return NoClue
	}
}

/*
	TODO: Before I can continue on this, I need to make an analysis of the distribution of Sagstrin for Lovforslag (sag type eq 3)
*/

type Lovtrin struct {
	Type     string
	Sagstrin []ftoda.Sagstrin
}

func NewFromSagId(sagid int, repo *repository.FTODAService) (l []Lovtrin, err error) {

	// I need to fetch sagstrin by id

	// then I need the logic to determine in which lovtrin each go
	/*
		Options:
			- Fremsættelse
				- Fremsættelse
				- Lovforslag som fremsat
			- Første Behandling
	*/

	// then

	return
}
