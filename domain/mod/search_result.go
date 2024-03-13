package mod

type SearchResultDto struct {
	Name        string
	Author      string
	Description string
}

type SearchResult struct {
	name        string
	author      string
	description string
}

func ReformSearchResult(dto SearchResultDto) SearchResult {
	return SearchResult{
		name:        dto.Name,
		author:      dto.Author,
		description: dto.Description,
	}
}

func (result SearchResult) Name() string {
	return result.name
}

func (result SearchResult) Author() string {
	return result.author
}

func (result SearchResult) Description() string {
	return result.description
}
