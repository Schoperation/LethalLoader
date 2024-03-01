package rest

type ThunderstoreClient struct{}

func NewThunderstoreClient() ThunderstoreClient {
	return ThunderstoreClient{}
}

func (client ThunderstoreClient) GetModByNameAndAuthor(name, author string)
