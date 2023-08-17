package ftoda

import (
	"fmt"
	"slices"
	"context"
	"sync"
	dataloader "github.com/graph-gophers/dataloader/v7"
)

var (
	aktørLoader *dataloader.Loader[int, *Aktør]
	aktørLoaderOnce sync.Once
) 

func aktørBatchFunction(ctx context.Context, keys[] int) (results []*dataloader.Result[*Aktør]) {
	repo := newRepository()
	aktører, err := repo.getAktørerByIds(keys)

	if err != nil {
		panic(err)
	}


	// Instead of iterating through the database result,
	// We iterate through the keys and then use the slice.IndexFunc to see if an Afstemning with the key is in the results. If not, we append an error to the results.
	for _, key := range keys {
		i := slices.IndexFunc(aktører, func(aktør Aktør) bool {
			return aktør.Id == key
		})

		if i == -1 {
			e := fmt.Errorf("record not found: Aktør with id %d not found", key)
			results = append(results, &dataloader.Result[*Aktør]{Data: &Aktør{},Error: e})
		} else {
			results = append(results, &dataloader.Result[*Aktør]{Data: &aktører[i]})
		}
	}

	return 
}

func newAktørLoader() *dataloader.Loader[int, *Aktør] {
	aktørLoaderOnce.Do(func() {
		cache := &dataloader.NoCache[int, *Aktør]{}
		aktørLoader = dataloader.NewBatchedLoader(aktørBatchFunction, dataloader.WithCache[int, *Aktør](cache))
	})

	return aktørLoader
}

type Aktør struct {
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
}

func (Aktør) TableName() string {
	return "Aktør"
}

func LoadAktørById(id int) (aktør Aktør, err error) {
	loader := newAktørLoader()
	thunk := loader.Load(context.Background(), id)

	result, err := thunk()

	aktør = *result

	return
}

func LoadAktørByName(name string) (aktør Aktør, err error) {
	repo := newRepository()
	return repo.getAktørByName(name)

	return
}

func LoadAktører(limit int, offset int) (aktører []Aktør, err error) {
	repo := newRepository()
	return repo.getAktører(limit, offset)
}

func LoadAktørerByType(limit int, offset int, aktørType string) (aktører []Aktør, err error) {
	repo := newRepository()
	return repo.getAktørerByType(limit, offset, aktørType)
}

func SearchAktørByName(limit int, aktørName string) (aktører []Aktør, err error) {
	repo := newRepository()
	return repo.searchAktørByName(limit, aktørName)
}

func (r *Repository) getAktører(limit int, offset int) (aktører []Aktør, err error) {
	result := r.db.Table("Aktør").Limit(limit).Offset(offset).Select("Aktør.*, Aktørtype.type").Joins("left join Aktørtype on Aktør.typeid = Aktørtype.id").Find(&aktører)
	err = result.Error
	return
}

func (r *Repository) getAktørerByType(limit int, offset int, aktørType string) (aktører []Aktør, err error) {
	result := r.db.Table("Aktør").Limit(limit).Offset(offset).Select("Aktør.*, Aktørtype.type").Joins("left join Aktørtype on Aktør.typeid = Aktørtype.id").Where("Aktørtype.type = ?", aktørType).Find(&aktører)
	err = result.Error
	return
}

func (r *Repository) getAktørerByIds(ids []int) (aktører []Aktør, err error) {
	result := r.db.Table("Aktør").Select("Aktør.*, Aktørtype.type").Joins("left join Aktørtype on Aktør.typeid = Aktørtype.id").Find(&aktører, ids)
	err = result.Error
	return 
}

func (r *Repository) getAktørByName(name string) (aktør Aktør, err error) {
	result := r.db.Where("navn = ?",name).Find(&aktør)
	err = result.Error
	return 
}

func (r *Repository) searchAktørByName(limit int, name string) (aktører []Aktør, err error) {
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
	
	result := r.db.Raw(search_conditions).Scan(&aktører)	
	err = result.Error
	return 
}
