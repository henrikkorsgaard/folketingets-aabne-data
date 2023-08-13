package resolvers

import (
	"context"
	"fmt"
	"henrikkorsgaard/folketingets-aabne-data/ftoda"
	"time"

	"github.com/graph-gophers/dataloader"
	graphql "github.com/graph-gophers/graphql-go"
)

var (
	loader *dataloader.Loader
)

func init() {
	loader = dataloader.NewBatchedLoader(StemmeBatchFunction)
}

type StemmeQueryArgs struct {
	QueryArgs
	AfstemningId *int32
}

type StemmeResolver struct {
	stemme ftoda.Stemme
}

// I need to initiate the loader for stemmer
// I need to be able to call it
// Where can I bind it for now?

func StemmeBatchFunction(ctx context.Context, keys dataloader.Keys) (results []*dataloader.Result) {

	// This is where I fetch all the Stemmer based on ids from the database

	return
}

func NewStemmeList(args StemmeQueryArgs) (resolvers []*StemmeResolver, err error) {

	repo := ftoda.NewRepository()

	if args.Id != nil {
		var stemme ftoda.Stemme
		stemme, err = repo.GetStemme(int(*args.Id))

		stemmeResolver := StemmeResolver{stemme}
		resolvers = append(resolvers, &stemmeResolver)

		return
	}

	// if the query does not supply an offset
	// we set it to 0
	if args.Offset == nil {
		var offset int32 = 0
		args.Offset = &offset
	}

	stemmer, err := repo.GetAllStemme(200, int(*args.Offset))

	for _, stemme := range stemmer {
		stemmeResolver := StemmeResolver{stemme}
		resolvers = append(resolvers, &stemmeResolver)
	}

	return
}

func NewStemme(ctx Context, args StemmeQueryArgs) (resolver *StemmeResolver, err error) {

	// we get context from the base query like here: https://github.com/tonyghita/graphql-go-example/blob/37cd51aae44b998ee3baa2b7e9c21c56e11a5fe3/resolver/query.go#L36
	// I eksemplet der er indgangen NewFilms som så får Resolvers fra NewFilm: https://github.com/tonyghita/graphql-go-example/blob/37cd51aae44b998ee3baa2b7e9c21c56e11a5fe3/resolver/film.go#L50
	// NewFilms: for hver film i søgningen appender den url'en til loaderen https://github.com/tonyghita/graphql-go-example/blob/37cd51aae44b998ee3baa2b7e9c21c56e11a5fe3/resolver/film.go#L56C25-L56C34
	// NewFilm: Loader med den medgivne URL: https://github.com/tonyghita/graphql-go-example/blob/37cd51aae44b998ee3baa2b7e9c21c56e11a5fe3/resolver/film.go#L30

	// Mit eksempel: når NewSgtemme kaldes, så sendes id'en med konteksten til loaderen
	// Den kalder loadBatch på et tidspunkt, hvor jeg så henter alle stemmer baseret på id'er
	// returnere Stemmer og så pakker dem ind i resolvers?

	// hvad returnere loadfilm?

	// Batch function returnere et map med objekter og errors

	thunk := loader.Load(ctx, dataloader.StringKey(*args.Id))
	result, err := thunk()
	fmt.Println(result)

	// This is the wrong way around -- whenever NewStemme is called, it should send the id to the loader

	resolvers, err := NewStemmeList(args)
	if err != nil {
		return
	}

	if len(resolvers) > 0 {
		resolver = resolvers[0]
	}

	return
}

func (s *StemmeResolver) Id() int32 {
	return int32(s.stemme.Id)
}

func (s *StemmeResolver) Type() *string {
	return &s.stemme.Type
}

func (s *StemmeResolver) Opdateringsdato() graphql.Time {
	t, err := time.Parse(time.DateTime, s.stemme.Opdateringsdato)
	if err != nil {
		panic(err) // This field is not null, I want to catch errors in development
	}
	return graphql.Time{t}
}

func (s *StemmeResolver) Aktør() (*AktørResolver, error) {
	id := int32(s.stemme.AktørId)
	args := AktørQueryArgs{QueryArgs: QueryArgs{Id: &id}}
	return NewAktør(args)
}

func (s *StemmeResolver) Afstemning() (*AfstemningResolver, error) {
	id := int32(s.stemme.AfstemningId)
	args := QueryArgs{Id: &id}
	return NewAfstemning(args)
}
