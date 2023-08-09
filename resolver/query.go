package resolver 

type QueryResolver struct {}

func (qr *QueryResolver) Afstemning(args AfstemningQueryArgs) (*AfstemningResolver, error) {
	return NewAfstemning(args)
}