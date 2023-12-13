package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewcampaign(t *testing.T) {
	assert := assert.New(t)

	name := "Campaing X"
	content := "Body"
	contacts := []string{"email@gmail.com", "email2@gmail.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
}
