package kasper

import (
	"testing"

	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()
	assert.NotNil(t, config.MarkOffsetsHook)
	assert.Equal(t, sarama.WaitForAll, config.RequiredAcks)
	assert.Equal(t, 5000, config.MaxInFlightMessageGroups)
}