package resolvers

import (
	

	graphql "github.com/graph-gophers/graphql-go"
)

type AktørQueryArgs struct {
	QueryArgs
	Type *string
}

type AktørResolver struct {
	aktør Aktør
}

type Aktør struct {
	Id int32
	Type string
	GruppeNavnKort string
	Navn string
	Fornavn string
	Efternavn string
	Biografi string 
	PeriodeID int32 
	StartDato graphql.Time
	SlutDato graphql.Time
	Opdateringsdato graphql.Time
}

func NewAktørList(args AktørQueryArgs) (resolvers []*AktørResolver, err error){
	/*
	repo := newSqlite()

	query := "SELECT Aktør.id, Aktørtype.type, Aktør.gruppenavnkort, Aktør.navn, Aktør.fornavn, Aktør.efternavn, Aktør.biografi, Aktør.periodeid, Aktør.startdato, Aktør.slutdato, Aktør.opdateringsdato FROM Aktør JOIN Aktørtype ON Aktør.typeid = Aktørtype.id"

	if args.Id != nil {
		query +=  " WHERE Aktør.id=" + fmt.Sprintf("%d", *args.Id)
	}

	if args.Type != nil {
		if args.Id != nil {
			query += " AND "
		} else {
			query +=  " WHERE "
		}

		query += "Aktørtype.type='" + fmt.Sprintf("%s", strings.Title(*args.Type)) + "'"
	}

	if args.Offset == nil {
		var offset int32 = 0
		args.Offset = &offset
	}

	query+= " LIMIT " + os.Getenv("GRAPHQL_QUERY_LIMIT") + " OFFSET " + fmt.Sprintf("%d", *args.Offset) + ";"

	rows, err := repo.db.Query(query)
	if err != nil {
		return
	}

	for rows.Next(){
		aktør := Aktør{}

		var gruppenavnkort sql.NullString 
		var navn sql.NullString 
		var fornavn sql.NullString 
		var efternavn sql.NullString 
		var biografi sql.NullString 
		var periodeid sql.NullInt32
		var startdato sql.NullString
		var slutdato sql.NullString
		var opdateringsdato string

		err = rows.Scan(&aktør.Id, &aktør.Type, &gruppenavnkort, &navn,&fornavn, &efternavn, &biografi, &periodeid, &startdato, &slutdato, &opdateringsdato)

		if err != nil {
			break
		}

		if navn.Valid {
			aktør.Navn = navn.String
		}

		if fornavn.Valid {
			aktør.Fornavn = fornavn.String
		}

		if efternavn.Valid {
			aktør.Efternavn = efternavn.String
		}

		if biografi.Valid {
			aktør.Biografi = biografi.String
		}

		if periodeid.Valid {
			aktør.PeriodeID = periodeid.Int32
		}

		if startdato.Valid {
			t_start, err := time.Parse(time.DateTime, startdato.String)
			if err != nil {
				break
			}
			aktør.StartDato.Time = t_start
		}

		if slutdato.Valid {
			t_slut, err := time.Parse(time.DateTime, slutdato.String)
			if err != nil {
				break
			}
			aktør.SlutDato.Time = t_slut
		}

		t_op, err := time.Parse(time.DateTime, opdateringsdato)
		if err != nil {
			break
		}

		
		aktør.Opdateringsdato.Time = t_op

		aktørResolver := AktørResolver{aktør}
		resolvers = append(resolvers, &aktørResolver)
	}

	if rows.Err() != nil {
		err = rows.Err()
	}

	if args.Id != nil && len(resolvers) == 0 {
		err = fmt.Errorf("Unable to resolve Aktør: Id %d does not exist", *args.Id)
	}
	*/

	return
}

// we need the single return if other objects has a single entity in their schema
func NewAktør(args AktørQueryArgs) (resolver *AktørResolver,err error) {
	resolvers, err := NewAktørList(args)
	if err != nil {
		return
	}

	if len(resolvers) > 0 {
		resolver = resolvers[0]
	}

	return 
}

func (a *AktørResolver) Id() int32 {
	return a.aktør.Id
}

func (a *AktørResolver) Type() string {
	return a.aktør.Type
}

func (a *AktørResolver) Gruppenavnkort() *string {
	return &a.aktør.GruppeNavnKort
}

func (a *AktørResolver) Navn() *string {
	return &a.aktør.Navn
}

func (a *AktørResolver) Fornavn() *string {
	return &a.aktør.Navn
}

func (a *AktørResolver) Efternavn() *string {
	return &a.aktør.Navn
}

func (a *AktørResolver) Biografi() *string {
	return &a.aktør.Navn
}

//TODO: this should return a periodeResolver
func (a *AktørResolver) PeriodeID() *int32 {
	return &a.aktør.PeriodeID
}

func (a *AktørResolver) Startdato() *graphql.Time {
	return &a.aktør.StartDato
}

func (a *AktørResolver) Slutdato() *graphql.Time {
	return &a.aktør.SlutDato
}


func (a *AktørResolver) Opdateringsdato() graphql.Time {
	return a.aktør.Opdateringsdato
}
