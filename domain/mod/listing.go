package mod

type ListingDto struct {
	Name         string
	Version      string
	Author       string
	Description  string
	DownloadUrl  string
	Dependencies []string
}

type Listing struct {
	name         string
	version      string
	author       string
	description  string
	downloadUrl  string
	dependencies []string
}

func ReformListing(dto ListingDto) Listing {
	return Listing{
		name:         dto.Name,
		version:      dto.Version,
		author:       dto.Author,
		description:  dto.Description,
		downloadUrl:  dto.DownloadUrl,
		dependencies: dto.Dependencies,
	}
}

func (listing Listing) Name() string {
	return listing.name
}

func (listing Listing) Version() string {
	return listing.version
}

func (listing Listing) Author() string {
	return listing.author
}

func (listing Listing) Description() string {
	return listing.description
}

func (listing Listing) DownloadUrl() string {
	return listing.downloadUrl
}

func (listing Listing) Dependencies() []string {
	return listing.dependencies
}
