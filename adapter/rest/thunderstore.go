package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"schoperation/lethalloader/domain/mod"
	"time"
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

func (client ThunderstoreClient) GetModByNameAndAuthor(name, author string) (mod.ListingDto, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://thunderstore.io/api/experimental/package/%s/%s", author, name), nil)
	if err != nil {
		return mod.ListingDto{}, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", "LethalLoader")

	backOffSchedule := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		10 * time.Second,
	}

	var response *http.Response
	for _, backoff := range backOffSchedule {
		response, err = http.DefaultClient.Do(request)
		if err != nil {
			return mod.ListingDto{}, err
		}

		if response.StatusCode == http.StatusOK {
			break
		} else if response.StatusCode == http.StatusGatewayTimeout || response.StatusCode == http.StatusServiceUnavailable || response.StatusCode == http.StatusRequestTimeout {
			time.Sleep(backoff)
			continue
		} else {
			return mod.ListingDto{}, fmt.Errorf("failed to get response from thunderstore; status code %d", response.StatusCode)
		}
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
