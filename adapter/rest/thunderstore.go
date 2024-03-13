package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"schoperation/lethalloader/domain/mod"
	"time"

	"golang.org/x/net/html"
)

type ThunderstoreClient struct{}

func NewThunderstoreClient() ThunderstoreClient {
	return ThunderstoreClient{}
}

type getPackageModel struct {
	Name   string      `json:"name"`
	Owner  string      `json:"owner"`
	Latest latestModel `json:"latest"`
}

type latestModel struct {
	Description   string   `json:"description"`
	Dependencies  []string `json:"dependencies"`
	VersionNumber string   `json:"version_number"`
	DownloadUrl   string   `json:"download_url"`
}

func (client ThunderstoreClient) doReq(req *http.Request) (*http.Response, error) {
	backOffSchedule := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		10 * time.Second,
	}

	var response *http.Response
	var err error
	for _, backoff := range backOffSchedule {
		response, err = http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		if response.StatusCode == http.StatusOK {
			break
		} else if response.StatusCode == http.StatusGatewayTimeout || response.StatusCode == http.StatusServiceUnavailable || response.StatusCode == http.StatusRequestTimeout {
			time.Sleep(backoff)
			continue
		} else {
			return nil, fmt.Errorf("failed to get response from thunderstore; status code %d", response.StatusCode)
		}
	}

	return response, nil
}

func (client ThunderstoreClient) GetModByNameAndAuthor(name, author string) (mod.ListingDto, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://thunderstore.io/api/experimental/package/%s/%s", author, name), nil)
	if err != nil {
		return mod.ListingDto{}, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", "LethalLoader")

	response, err := client.doReq(request)
	if err != nil {
		return mod.ListingDto{}, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return mod.ListingDto{}, err
	}

	var model getPackageModel
	err = json.Unmarshal(body, &model)
	if err != nil {
		return mod.ListingDto{}, err
	}

	return mod.ListingDto{
		Name:         model.Name,
		Version:      model.Latest.VersionNumber,
		Author:       model.Owner,
		Description:  model.Latest.Description,
		DownloadUrl:  model.Latest.DownloadUrl,
		Dependencies: model.Latest.Dependencies,
	}, nil
}

func (client ThunderstoreClient) Search(term string) ([]mod.SearchResultDto, error) {
	request, err := http.NewRequest(http.MethodGet, "https://thunderstore.io/c/lethal-company/", nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "text/html")
	request.Header.Set("User-Agent", "LethalLoader")

	queryParams := url.Values{}
	queryParams.Add("q", url.QueryEscape(term))
	queryParams.Add("ordering", "top-rated")
	queryParams.Add("section", "mods")
	request.URL.RawQuery = queryParams.Encode()

	response, err := client.doReq(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	doc, err := html.Parse(response.Body)
	if err != nil {
		return nil, err
	}
}

// TODO use goquery?
func (client ThunderstoreClient) processHtmlDoc(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "div" {
		client.processResult(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		client.processHtmlDoc(c)
	}
}

func (client ThunderstoreClient) processResult(n *html.Node) {

}
