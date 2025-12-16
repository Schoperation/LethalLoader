package mod

import "fmt"

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

func (result SearchResult) PrettyPrint(nameLimit, authorLimit int) string {
	modName := result.Name()
	if len(modName) > nameLimit {
		modName = modName[:nameLimit-3] + "..."
	} else {
		for range nameLimit - len(modName) {
			modName += " "
		}
	}

	modAuthor := result.Author()
	if len(modAuthor) > authorLimit {
		modAuthor = modAuthor[:authorLimit-3] + "..."
	} else {
		for range authorLimit - len(modAuthor) {
			modAuthor += " "
		}
	}

	return fmt.Sprintf("%s ~ by %s", modName, modAuthor)
}
