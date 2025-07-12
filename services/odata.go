package ftoda

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var (
	ErrOdataRequest    = errors.New("odata request error")
	ErrEncodingUrl     = errors.New("error encoding the odata url")
	ErrParsingBody     = errors.New("error parsing odata response body")
	ErrUnmarshallOdata = errors.New("error unmarshalling json to odata ")
	ErrUnmarshallType  = errors.New("error unmashalling json to type")
)

type apiRepository struct {
	host string
}

func newAPIRepository(host string) *apiRepository {
	return &apiRepository{
		host: host,
	}
}

func (repo *apiRepository) getData(q odataQuery, v any) error {

	queryUrl, err := q.GetEncodedUrl(repo.host)
	if err != nil {
		return errors.Join(ErrEncodingUrl, err)
	}

	res, err := http.Get(queryUrl)
	if err != nil {
		return errors.Join(ErrOdataRequest, err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.Join(ErrParsingBody, err)
	}

	var odata odataResult
	err = json.Unmarshal(body, &odata)
	if err != nil {
		return errors.Join(ErrUnmarshallOdata, err)
	}

	err = json.Unmarshal(odata.Result, v)
	if err != nil {
		return errors.Join(ErrUnmarshallType, err)
	}
	return nil
}

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
func (q *odataQuery) PrettyUrl(host string) string {
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

func (q *odataQuery) GetEncodedUrl(host string) (string, error) {
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
