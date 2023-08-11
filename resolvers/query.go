package resolvers

type QueryArgs struct {
	Id *int32
	Offset *int32
}

type QueryResolver struct {}

func (qr *QueryResolver) Afstemning(args QueryArgs) ([]*AfstemningResolver, error) {
	return NewAfstemningList(args)
}

func (qr *QueryResolver) Stemme(args QueryArgs) ([]*StemmeResolver, error) {

	sargs := StemmeQueryArgs{QueryArgs:args}
	return NewStemmeList(sargs)
}

// TODO: Look at  enums to see if this makes query better!
// https://github.com/tonyghita/graphql-go-example/blob/main/schema/type/starship.graphql
// https://github.com/tonyghita/graphql-go-example/blob/main/resolver/mass_unit.go
func (qr *QueryResolver) Aktør(args AktørQueryArgs) ([]*AktørResolver, error) {

	return NewAktørList(args)
}