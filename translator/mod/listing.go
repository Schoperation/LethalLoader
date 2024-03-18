package mod

import "schoperation/lethalloader/domain/mod"

type thunderstoreClient interface {
	GetModByNameAndAuthor(name, author string) (mod.ListingDto, error)
	GetModByNameAuthorVersion(name, author, version string) (mod.ListingDto, error)
}

type ListingTranslator struct {
	thunderstoreClient thunderstoreClient
}

func NewListingTranslator(
	thunderstoreClient thunderstoreClient,
) ListingTranslator {
	return ListingTranslator{
		thunderstoreClient: thunderstoreClient,
	}
}

func (translator ListingTranslator) GetByNameAndAuthor(name, author string) (mod.Listing, error) {
	dto, err := translator.thunderstoreClient.GetModByNameAndAuthor(name, author)
	if err != nil {
		return mod.Listing{}, err
	}

	return mod.ReformListing(dto), nil
}

func (translator ListingTranslator) GetBySlug(slug mod.Slug) (mod.Listing, error) {
	dto, err := translator.thunderstoreClient.GetModByNameAuthorVersion(slug.Name(), slug.Author(), slug.Version())
	if err != nil {
		return mod.Listing{}, err
	}

	return mod.ReformListing(dto), nil
}
