package analysis

import (
	"fmt"
	"strconv"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"github.com/henrikkorsgaard/folketingets-aabne-data/repository"
)

// I don't know if I want to do the analysis in go and return it as an analysis
// or return the objects to the template and write the there with html+js+vis
func LovforslagSagstrinTypeDistribution(lovforslagCount int, repo *repository.FTODAService) (sagstrinsager []ftoda.SagSagstrin, err error) {

	// we also need sagstrin status to be able to match this.

	sagstrinsager, err = repo.GetSagerByTypeWithSagstrin(3, 400, 0)
	if err != nil {
		return sagstrinsager, err
	}

	//histogram based on number of sagstrin
	sagstrinCountHistogram := make(map[int]int)

	//historgram based on the type of sagstrin in sag
	sagstrinTypeHistorgram := make(map[int]int)

	//histogram based on the event series - that would be unique event series and then count them
	sagstrinEventSeriesHistorgram := make(map[string]int)

	for _, s := range sagstrinsager {
		sagstrinCount := len(s.Sagstrin)
		sagstrinCountHistogram[sagstrinCount] += 1
		sagstrinEventseries := ""
		for _, st := range s.Sagstrin {
			sagstrinTypeHistorgram[st.Typeid] += 1
			sagstrinEventseries += strconv.Itoa(st.Typeid)
		}
		sagstrinEventSeriesHistorgram[sagstrinEventseries] += 1
	}

	fmt.Printf("Analysing %d lovforslag:\n\tSagstring count: %+v\n\tSagstring type: %+v\n\tSagstring series: %+v\n\t", len(sagstrinsager), sagstrinCountHistogram, sagstrinTypeHistorgram, sagstrinEventSeriesHistorgram)

	return sagstrinsager, err
}
