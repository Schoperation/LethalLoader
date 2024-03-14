package mod

import "schoperation/lethalloader/domain/mod"

type thunderstoreSearcher interface {
	Search(term string) ([]mod.SearchResultDto, error)
}

type SearchResultTranslator struct {
	thunderstoreSearcher thunderstoreSearcher
}

func NewSearchResultTranslator(
	thunderstoreSearcher thunderstoreSearcher,
) SearchResultTranslator {
	return SearchResultTranslator{
		thunderstoreSearcher: thunderstoreSearcher,
	}
}

func (translator SearchResultTranslator) Search(term string) ([]mod.SearchResult, error) {
	dtos, err := translator.thunderstoreSearcher.Search(term)
	if err != nil {
		return nil, err
	}

	results := make([]mod.SearchResult, len(dtos))
	for i, dto := range dtos {
		results[i] = mod.ReformSearchResult(dto)
	}

	return results, nil
}
