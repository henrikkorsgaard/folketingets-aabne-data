package resolvers

type QueryArgs struct {
	Id     *int32
	Offset *int32
}

type QueryResolver struct{}

func (qr *QueryResolver) Afstemning(args AfstemningQueryArgs) ([]*AfstemningResolver, error) {

	return NewAfstemningList(args)
}

func (qr *QueryResolver) Sag(args SagQueryArgs) ([]*SagResolver, error) {
	return NewSagList(args)
}


func (qr *QueryResolver) Aktør(args AktørQueryArgs) ([]*AktørResolver, error) {
	return NewAktørList(args)
}

func (qr *QueryResolver) SearchAktør(args AktørSearchArgs) ([]*AktørResolver, error) {
	return NewAktørResultList(args)
}
