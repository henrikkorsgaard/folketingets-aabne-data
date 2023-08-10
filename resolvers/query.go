package resolvers

type QueryArgs struct {
	Id *int32
}

type QueryResolver struct {}

func (qr *QueryResolver) Afstemning(args QueryArgs) ([]*AfstemningResolver, error) {
	return NewAfstemningList(args)
}

func (qr *QueryResolver) Stemme(args QueryArgs) ([]*StemmeResolver, error) {
	return NewStemmeList(args)
}

// TODO: Look at schema enums to see if this makes query better!
func (qr *QueryResolver) Aktør(args AktørQueryArgs) ([]*AktørResolver, error) {
	return NewAktørList(args)
}