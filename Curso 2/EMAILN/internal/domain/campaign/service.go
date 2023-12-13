package campaign

import (
	"emailn/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(campaignDTO contract.CampaignDTO) (string, error) {
	campaign, err := NewCampaign(campaignDTO.Name, campaignDTO.Content, campaignDTO.Emails)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", err
	}

	return campaign.Id, nil
}
