package resolvers

type QueryResolver struct {}

func (qr *QueryResolver) Afstemning(args AfstemningQueryArgs) ([]*AfstemningResolver, error) {
	return NewAfstemningList(args)
}