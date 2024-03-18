package mod

import "time"

type ListingDto struct {
	Name         string
	Version      string
	Author       string
	Description  string
	DownloadUrl  string
	DateCreated  time.Time
	Dependencies []string
}

type Listing struct {
	name         string
	version      string
	author       string
	description  string
	downloadUrl  string
	dateCreated  time.Time
	dependencies []Slug
}

func ReformListing(dto ListingDto) Listing {
	var deps []Slug
	for _, dep := range dto.Dependencies {
		depSlug := ReformSlugFromString(dep)
		deps = append(deps, depSlug)
	}

	return Listing{
		name:         dto.Name,
		version:      dto.Version,
		author:       dto.Author,
		description:  dto.Description,
		downloadUrl:  dto.DownloadUrl,
		dateCreated:  dto.DateCreated,
		dependencies: deps,
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

func (listing Listing) DateCreated() time.Time {
	return listing.dateCreated
}

func (listing Listing) Dependencies() []Slug {
	return listing.dependencies
}

func (listing Listing) Slug() Slug {
	return ReformSlug(listing.name, listing.author, listing.version)
}
