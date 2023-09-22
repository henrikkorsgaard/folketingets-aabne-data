package ftoda

import (
	"fmt"
	"slices"
	"context"
	"sync"
	"encoding/xml"
	"html"
	dataloader "github.com/graph-gophers/dataloader/v7"
)

var (
	aktorLoader *dataloader.Loader[int, *Aktor]
	aktorLoaderOnce sync.Once
) 

func aktorBatchFunction(ctx context.Context, keys[] int) (results []*dataloader.Result[*Aktor]) {
	repo := newRepository()
	aktorer, err := repo.getAktorerByIds(keys)

	if err != nil {
		panic(err)
	}


	// Instead of iterating through the database result,
	// We iterate through the keys and then use the slice.IndexFunc to see if an Afstemning with the key is in the results. If not, we append an error to the results.
	for _, key := range keys {
		i := slices.IndexFunc(aktorer, func(aktor Aktor) bool {
			return aktor.Id == key
		})

		if i == -1 {
			e := fmt.Errorf("record not found: Aktor with id %d not found", key)
			results = append(results, &dataloader.Result[*Aktor]{Data: &Aktor{},Error: e})
		} else {
			results = append(results, &dataloader.Result[*Aktor]{Data: &aktorer[i]})
		}
	}

	return 
}

func newAktorLoader() *dataloader.Loader[int, *Aktor] {
	aktorLoaderOnce.Do(func() {
		cache := &dataloader.NoCache[int, *Aktor]{}
		aktorLoader = dataloader.NewBatchedLoader(aktorBatchFunction, dataloader.WithCache[int, *Aktor](cache))
	})

	return aktorLoader
}

type AktorBio struct {
	Kon string `gorm:"-" xml:"sex"`  
	Fodselsdato string `gorm:"-" xml:"born"` 
	Dodsdato string `gorm:"-" xml:"died"` 
	Parti string `gorm:"-" xml:"party"` 
	Billede string `gorm:"-" xml:"pictureMiRes"`
	Personligt string `gorm:"-" xml:"personalInformation>memberData>p"`
	Uddannelsesniveau string `gorm:"-" xml:"educationStatistic"`
	Uddannelse []string `gorm:"-" xml:"educations>education"` 
	Beskaftigelse []string `gorm:"-" xml:"occupations>occupation"`
	Ministerposter []string `gorm:"-" xml:"career>minister"`
	Nomineringer []string `gorm:"-" xml:"career>nominations"`
	Embede []string `gorm:"-" xml:"career>parliamentaryPositionsOfTrust>parliamentaryPositionOfTrust"`
	Positioner []string `gorm:"-" xml:"positionsOfTrust>positionOfTrust"`
	Valgkredse []string `gorm:"-" xml:"career>constituencies>constituency"`
	Titel string `gorm:"-" xml:"personalInformation>function>formattedTitles"`
}

type Aktor struct {
	Id int `gorm:"primaryKey"`
	Type string
	GruppeNavnKort string `gorm:"column:gruppenavnkort"`
	Navn string
	Fornavn string 
	Efternavn string
	Biografi string
	Periode int 
	Opdateringsdato string
	Startdato string
	Slutdato string
	// Extracted from biografi
	AktorBio
}

func (Aktor) TableName() string {
	return "Aktør"
}

func (a *Aktor) parseBio() {

	if a.Biografi == "" {
		return
	}

	err := xml.Unmarshal([]byte(a.Biografi), &a.AktorBio)
	if err != nil {
		return
	}
	a.Personligt = html.UnescapeString(a.Personligt)
}

func LoadAktorById(id int) (aktor Aktor, err error) {
	loader := newAktorLoader()
	thunk := loader.Load(context.Background(), id)

	result, err := thunk()

	aktor = *result

	aktor.parseBio()
	return
}

func LoadAktorByName(name string) (aktor Aktor, err error) {
	repo := newRepository()
	aktor, err = repo.getAktorByName(name)
	if err != nil {
		return aktor, err
	}
	aktor.parseBio()
	return aktor, err
}

func LoadAktorer(limit int, offset int) (aktorer []Aktor, err error) {
	repo := newRepository()
	aktorer, err = repo.getAktorer(limit, offset)
	if err != nil {
		return
	}
	for i := 0; i < len(aktorer); i++ {
		aktorer[i].parseBio()
	}

	return 
}

func LoadAktorerByType(limit int, offset int, aktorType string) (aktorer []Aktor, err error) {
	repo := newRepository()
	aktorer, err = repo.getAktorerByType(limit, offset, aktorType)
	if err != nil {
		return
	}
	
	for i := 0; i < len(aktorer); i++ {
		aktorer[i].parseBio()
	}

	return
}

func SearchAktorByName(limit int, aktorName string) (aktorer []Aktor, err error) {
	repo := newRepository()
	aktorer, err = repo.searchAktorByName(limit, aktorName)
	if err != nil {
		return
	}

	for i := 0; i < len(aktorer); i++ {
		aktorer[i].parseBio()
	}
	return 
}

func (r *Repository) getAktorer(limit int, offset int) (aktorer []Aktor, err error) {
	result := r.db.Table("Aktør").Limit(limit).Offset(offset).Select("Aktør.*, Aktørtype.type").Joins("left join Aktørtype on Aktør.typeid = Aktørtype.id").Find(&aktorer)
	err = result.Error
	return
}

func (r *Repository) getAktorerByType(limit int, offset int, aktorType string) (aktorer []Aktor, err error) {
	result := r.db.Table("Aktør").Limit(limit).Offset(offset).Select("Aktør.*, Aktørtype.type").Joins("left join Aktørtype on Aktør.typeid = Aktørtype.id").Where("Aktørtype.type = ?", aktorType).Find(&aktorer)
	err = result.Error
	return
}

func (r *Repository) getAktorerByIds(ids []int) (aktorer []Aktor, err error) {
	result := r.db.Table("Aktør").Select("Aktør.*, Aktørtype.type").Joins("left join Aktørtype on Aktør.typeid = Aktørtype.id").Find(&aktorer, ids)
	err = result.Error
	return 
}

func (r *Repository) getAktorByName(name string) (aktor Aktor, err error) {
	result := r.db.Where("navn = ?",name).Find(&aktor)
	err = result.Error
	return 
}

func (r *Repository) searchAktorByName(limit int, name string) (aktorer []Aktor, err error) {
	//name = name
	//result := r.db.Limit(limit).Where("navn LIKE ? OR fornavn LIKE ? OR efternavn LIKE ? OR gruppenavnkort LIKE ?",name, name, name, name).Find(&aktører)
	// Short search ranking algorithm
	// We prioritize matches with name and first name + members of parliament (type 5)
	// Then we prioritize last name + member of parliament
	// Then group name OR lastname and broader groups within the parliament
	// finally we broaden this up amtches with name and lastname
	// private individuals (type 12) need closer matches
	search_conditions := fmt.Sprintf(`SELECT *, Aktørtype.type, CASE 
	WHEN navn LIKE '%[1]s%%' AND fornavn LIKE '%[1]s%%' AND typeid = 5 THEN 4
	WHEN efternavn LIKE '%[1]s%%' AND typeid = 5 THEN 3
	WHEN (gruppenavnkort LIKE '%[1]s%%' OR efternavn) LIKE '%[1]s%%' AND typeid < 10 THEN 2
	WHEN (navn LIKE '%[1]s%%' OR efternavn LIKE '%[1]s%%') AND typeid != 12 THEN 1
	WHEN navn LIKE '%[1]s%%' AND typeid = 12 THEN 1
	END AS search_rank FROM Aktør LEFT JOIN Aktørtype on Aktør.typeid = Aktørtype.id WHERE search_rank > 0 ORDER BY search_rank DESC LIMIT %d`, name, limit)
	
	result := r.db.Raw(search_conditions).Scan(&aktorer)	
	err = result.Error
	return 
}
