package mod

import (
	"fmt"
	"strings"
)

type Slug struct {
	name    string
	author  string
	version string
}

func NewSlug(name, author, version string) (Slug, error) {
	if strings.TrimSpace(name) == "" {
		return Slug{}, fmt.Errorf("name must not be blank")
	}

	if strings.TrimSpace(author) == "" {
		return Slug{}, fmt.Errorf("author must not be blank")
	}

	if strings.TrimSpace(version) == "" {
		return Slug{}, fmt.Errorf("version must not be blank")
	}

	return Slug{
		name:    name,
		author:  author,
		version: version,
	}, nil
}

func ReformSlug(name, author, version string) Slug {
	return Slug{
		name:    name,
		author:  author,
		version: version,
	}
}

func NewSlugFromString(slug string) (Slug, error) {
	if strings.TrimSpace(slug) == "" {
		return Slug{}, fmt.Errorf("slug must not be blank")
	}

	if strings.Count(slug, "-") != 2 {
		return Slug{}, fmt.Errorf("slug must have 2 dashes to denote separators")
	}

	splitSlug := strings.Split(slug, "-")

	return Slug{
		name:    splitSlug[1],
		author:  splitSlug[0],
		version: splitSlug[2],
	}, nil
}

func ReformSlugFromString(slug string) Slug {
	splitSlug := strings.Split(slug, "-")

	return Slug{
		name:    splitSlug[1],
		author:  splitSlug[0],
		version: splitSlug[2],
	}
}

func (slug Slug) Name() string {
	return slug.name
}

func (slug Slug) Author() string {
	return slug.author
}

func (slug Slug) Version() string {
	return slug.version
}

func (slug Slug) String() string {
	return slug.author + "-" + slug.name + "-" + slug.version
}
