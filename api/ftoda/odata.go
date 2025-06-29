package ftoda

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type apiRepository struct {
	host string
}

// This should be internal
// And the API should be Public with a private db and api
func newAPIRepository(host string) *apiRepository {

	return &apiRepository{
		host: host,
	}
}

// TODO: add select as option. This allow us to filter have minimal and full service
// This will happen when we start consuming the bff with the frontend elements.
type odataQuery struct {
	entity string
	filter string
	expand string
	order  string
	top    int
	skip   int
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
		sb.WriteString("&orderby=id desc")

	} else {
		sb.WriteString("&orderby=")
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
		params.Add("$orderby", "id desc")
	} else {
		params.Add("$orderby", q.order)
	}
	params.Add("$format", "json")

	baseUrl.RawQuery = params.Encode()

	return baseUrl.String(), nil
}

type odataResult struct {
	Metadata string          `json:"odata.metadata"`
	Result   json.RawMessage `json:"value"`
	NextLink string          `json:"odata.nextLink"`
	Skip     int
}

func (repo *apiRepository) getData(q odataQuery) (odata odataResult, err error) {

	queryUrl, err := q.GetEncodedUrl(repo.host)
	if err != nil {
		return odata, err
	}

	res, err := http.Get(queryUrl)
	if err != nil {
		err = fmt.Errorf("error making http request: %s", err)
		return odata, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("error reading the body of the request: %s", err)
		return odata, err
	}

	err = json.Unmarshal(body, &odata)
	if err != nil {
		err = fmt.Errorf("unable to marshal JSON: %s", err)
		return odata, err
	}

	return odata, err
}
