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


func (qr *QueryResolver) Aktor(args AktorQueryArgs) ([]*AktorResolver, error) {
	return NewAktorList(args)
}

func (qr *QueryResolver) SearchAktor(args AktorSearchArgs) ([]*AktorResolver, error) {
	return NewAktorResultList(args)
}
