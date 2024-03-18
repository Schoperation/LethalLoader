package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"schoperation/lethalloader/domain/mod"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ThunderstoreClient struct{}

func NewThunderstoreClient() ThunderstoreClient {
	return ThunderstoreClient{}
}

type getLatestPackageModel struct {
	Name   string      `json:"name"`
	Owner  string      `json:"owner"`
	Latest latestModel `json:"latest"`
}

type latestModel struct {
	Description   string    `json:"description"`
	Dependencies  []string  `json:"dependencies"`
	VersionNumber string    `json:"version_number"`
	DownloadUrl   string    `json:"download_url"`
	DateCreated   time.Time `json:"date_created"`
}

type getSpecificPackageModel struct {
	Namespace     string    `json:"namespace"`
	Name          string    `json:"name"`
	VersionNumber string    `json:"version_number"`
	Description   string    `json:"description"`
	Dependencies  []string  `json:"dependencies"`
	DownloadUrl   string    `json:"download_url"`
	DateCreated   time.Time `json:"date_created"`
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
	author = strings.ReplaceAll(author, " ", "_")
	name = strings.ReplaceAll(name, " ", "_")

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

	var model getLatestPackageModel
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
		DateCreated:  model.Latest.DateCreated,
		Dependencies: model.Latest.Dependencies,
	}, nil
}

func (client ThunderstoreClient) GetModByNameAuthorVersion(name, author, version string) (mod.ListingDto, error) {
	author = strings.ReplaceAll(author, " ", "_")
	name = strings.ReplaceAll(name, " ", "_")

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://thunderstore.io/api/experimental/package/%s/%s/%s", author, name, version), nil)
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

	var model getSpecificPackageModel
	err = json.Unmarshal(body, &model)
	if err != nil {
		return mod.ListingDto{}, err
	}

	return mod.ListingDto{
		Name:         model.Name,
		Version:      model.VersionNumber,
		Author:       model.Namespace,
		Description:  model.Description,
		DownloadUrl:  model.DownloadUrl,
		DateCreated:  model.DateCreated,
		Dependencies: model.Dependencies,
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
	queryParams.Add("q", term)
	queryParams.Add("ordering", "top-rated")
	queryParams.Add("section", "mods")
	request.URL.RawQuery = queryParams.Encode()

	response, err := client.doReq(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	var searchResultDtos []mod.SearchResultDto

	// Each result starts with a div with this class
	doc.Find(".col-6.col-md-4.col-lg-3.mb-2.p-1.d-flex.flex-column").Each(func(i int, s *goquery.Selection) {
		if i > 9 {
			return
		}

		title := s.Find(".mb-0.overflow-hidden.text-nowrap.w-100").Text()
		author := s.Find(".overflow-hidden.text-nowrap.w-100").Find("a").Text()
		description := s.Find(".bg-light.px-2.flex-grow-1").First().Text()

		description = strings.Trim(description, "\n")
		description = strings.Trim(description, " ")
		description = strings.Trim(description, "\n") // yes we have to do this twice...

		searchResultDtos = append(searchResultDtos, mod.SearchResultDto{
			Name:        title,
			Author:      author,
			Description: description,
		})
	})

	return searchResultDtos, nil
}
