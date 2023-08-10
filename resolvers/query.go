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

func (qr *QueryResolver) Aktør(args QueryArgs) ([]*AktørResolver, error) {
	return NewAktørList(args)
}