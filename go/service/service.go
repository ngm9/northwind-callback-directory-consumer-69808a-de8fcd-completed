package service

import (
	"northwind/callbackdirectory/client"
)

type DirectoryService struct {
	client client.CountriesClient
}

func NewDirectoryService(c client.CountriesClient) *DirectoryService {
	return &DirectoryService{client: c}
}

func (s *DirectoryService) GetPhoneNumbers(country string, phone string) (string, error) {
	payload, err := s.client.FetchCountry(country)
	if err != nil {
		return "", err
	}

	callingCode := "-1"
	if len(payload.Data) > 0 {
		record := payload.Data[0]
		if len(record.CallingCodes) > 0 {
			callingCode = record.CallingCodes[len(record.CallingCodes)-1]
		}
	}

	if callingCode == "-1" {
		return callingCode, nil
	}

	return "+" + callingCode + " " + phone, nil
}
