package apiservice

type APIService struct{}

func NewAPIService() *APIService {
	return &APIService{}
}

func (this *APIService) Request(endpoint string) string {
	return "response from " + endpoint
}
