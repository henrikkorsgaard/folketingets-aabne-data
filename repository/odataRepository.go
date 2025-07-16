package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
)

var (
	ErrOdataRequest    = errors.New("odata request error")
	ErrEncodingUrl     = errors.New("error encoding the odata url")
	ErrParsingBody     = errors.New("error parsing odata response body")
	ErrUnmarshallOdata = errors.New("odata: cannot parse odata result ")
	ErrUnmarshallType  = errors.New("odata: cannot parse datatype")
)

type odataRepository struct {
	host string
}

func newOdataRepo(host string) *odataRepository {
	return &odataRepository{
		host: host,
	}
}

func (repo *odataRepository) getSagById(id int) (sag ftoda.Sag, err error) {
	q := odataQuery{
		entity: "Sag",
		filter: "id eq " + strconv.Itoa(id),
		expand: "EmneordSag",
	}

	var sager []ftoda.Sag
	err = repo.getData(q, &sager)
	if err != nil {
		return sag, errors.Join(ErrRepoGettingSag, err)
	}
	return sager[0], nil
}

func (repo *odataRepository) getSagerByType(sagtype int) (sager []ftoda.Sag, err error) {
	q := odataQuery{
		entity: "Sag",
		filter: "typeid eq " + strconv.Itoa(sagtype),
	}

	err = repo.getData(q, &sager)
	if err != nil {
		return sager, errors.Join(ErrRepoGettingSag, err)
	}
	return sager, nil
}

func (repo *odataRepository) getSagerByTypeWithSagstrin(sagtype int) (sager []ftoda.SagSagstrin, err error) {
	q := odataQuery{
		entity: "Sag",
		filter: "typeid eq " + strconv.Itoa(sagtype),
		expand: "Sagstrin",
	}

	err = repo.getData(q, &sager)
	if err != nil {
		return sager, errors.Join(ErrRepoGettingSag, err)
	}
	fmt.Println(sager[0])
	return sager, err
}

func (repo *odataRepository) getSagstrinById(id int) (sagstrin ftoda.Sagstrin, err error) {
	q := odataQuery{
		entity: "Sagstrin",
		filter: "id eq " + strconv.Itoa(id),
	}
	var sagstrinArr []ftoda.Sagstrin
	err = repo.getData(q, &sagstrinArr)
	if err != nil {
		return sagstrin, errors.Join(ErrRepoGettingSagstrin, err)
	}

	return sagstrinArr[0], err
}

func (repo *odataRepository) getSagstrinBySagId(sagid int) (sagstrin []ftoda.Sagstrin, err error) {
	q := odataQuery{
		entity: "Sagstrin",
		filter: "sagid eq " + strconv.Itoa(sagid),
		order:  "asc",
	}

	err = repo.getData(q, &sagstrin)
	if err != nil {
		return sagstrin, errors.Join(ErrRepoGettingSagstrin, err)
	}
	return sagstrin, err
}

func (repo *odataRepository) getSagstrinstype() (sagstrintypes []ftoda.Sagstrinstype, err error) {
	q := odataQuery{
		entity: "Sagstrinstype",
		top:    200,
	}

	err = repo.getData(q, &sagstrintypes)
	if err != nil {
		return sagstrintypes, errors.Join(ErrRepoGettingSagstrin, err)
	}
	return sagstrintypes, err
}

/*
Helper type that allow us to fetch odata in one call
while still adhearing to the repository interface definitions

Details: The ftoda odata version does not allow $filter operations
on $expand. Instead we query Emneordsag with sagid filter and expand
to Emneord.
*/
type EmneordSagEmneord struct {
	Id        int `json:"id"`
	EmneordId int `json:"emneordid"`
	SagId     int `json:"sagid"`
	Emneord   ftoda.Emneord
}

func (repo *odataRepository) getEmneordBySagId(sagid int) (emneord []ftoda.Emneord, err error) {

	q := odataQuery{
		entity: "EmneordSag",
		filter: "sagid eq " + strconv.Itoa(sagid),
		expand: "Emneord",
	}

	var emneordsager []EmneordSagEmneord
	err = repo.getData(q, &emneordsager)
	if err != nil {
		return emneord, errors.Join(ErrGettingEmneord, err)
	}

	for _, ems := range emneordsager {
		emneord = append(emneord, ftoda.Emneord{
			Id:      ems.Emneord.Id,
			Emneord: ems.Emneord.Emneord,
			TypeId:  ems.Emneord.TypeId,
			EmneordSag: ftoda.EmneordSag{
				Id:        ems.Id,
				EmneordId: ems.EmneordId,
				SagId:     ems.SagId,
			},
		})
	}

	return emneord, err
}

func (repo *odataRepository) getData(q odataQuery, v any) error {

	queryUrl, err := q.getEncodedUrl(repo.host)
	if err != nil {
		return errors.Join(ErrEncodingUrl, err)
	}

	odata, err := queryOdata(queryUrl)
	if err != nil {
		return err // we join errors in the query function
	}

	err = json.Unmarshal(odata.Result, v)
	if err != nil {
		fmt.Println(string(odata.Result))
		return errors.Join(ErrUnmarshallType, err)
	}

	fmt.Println(odata.NextLink)
	fmt.Println(odata.Skip) //DO I Pull skip from next link or do I query it from an URL?
	return nil
}

func queryOdata(urlString string) (result odataResult, err error) {

	res, err := http.Get(urlString)
	if err != nil {
		return result, errors.Join(ErrOdataRequest, err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return result, errors.Join(ErrParsingBody, err)
	}

	var odata odataResult
	err = json.Unmarshal(body, &odata)
	if err != nil {
		return result, errors.Join(ErrUnmarshallOdata, err)
	}

	return result, err
}

/*
	Odata helper functions
*/

type odataResult struct {
	Metadata string          `json:"odata.metadata"`
	Result   json.RawMessage `json:"value"`
	NextLink string          `json:"odata.nextLink"`
	Skip     int
}

type odataQuery struct {
	entity  string
	filter  string
	expand  string
	orderBy string
	order   string
	top     int
	skip    int
}

// For debugging odata url oddities
func (q *odataQuery) prettyUrl(host string) string {
	var sb strings.Builder

	sb.WriteString("https://")
	sb.WriteString(host)
	sb.WriteString("/api/")
	sb.WriteString(q.entity)
	sb.WriteString("?$format=json")
	if q.expand != "" {
		sb.WriteString("&$expand=")
		sb.WriteString(q.expand)
	}
	if q.filter != "" {
		sb.WriteString("&$filter=")
		sb.WriteString(q.filter)
	}
	if q.top != 0 {
		sb.WriteString("&$top=")
		sb.WriteString(strconv.Itoa(q.top))
	}
	sb.WriteString("&$skip=")
	sb.WriteString(strconv.Itoa(q.skip))

	if q.order == "" {
		q.order = "desc"
	}

	if q.orderBy == "" {
		sb.WriteString("&orderby=id ")
		sb.WriteString(q.orderBy)
	} else {
		sb.WriteString("&orderby=")
		sb.WriteString(q.orderBy)
		sb.WriteString(" ")
		sb.WriteString(q.order)
	}

	return sb.String()
}

func (q *odataQuery) getEncodedUrl(host string) (string, error) {
	baseUrl, err := url.Parse("https://" + host + "/api/" + q.entity)
	if err != nil {
		err = fmt.Errorf("error with baseurl: %s", err)
		return "", err
	}

	params := url.Values{}
	if q.expand != "" {
		params.Add("$expand", q.expand)
	}

	if q.filter != "" {
		params.Add("$filter", q.filter)
	}

	if q.top != 0 {
		params.Add("$top", strconv.Itoa(q.top))
	}

	// Pagination should be handles elsewhere to be honest.
	params.Add("$skip", strconv.Itoa(q.skip)) //Defaults to zero

	if q.order == "" {
		q.order = "desc"
	}

	if q.orderBy == "" {
		params.Add("$orderby", "id "+q.order)
	} else {
		params.Add("$orderby", q.orderBy+" "+q.order)
	}
	params.Add("$format", "json")

	baseUrl.RawQuery = params.Encode()

	return baseUrl.String(), nil
}
