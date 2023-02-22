package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Should_CalculateValueVariation(t *testing.T) {
	assert := assert.New(t)
	a := NewAsset("test", 100, 100, 70, 100, 100, true)

	assert.Equal(-.3, a.CalculateValueVariation())
}
func Test_Should_CalculatePercentageFromTotal(t *testing.T) {
	assert := assert.New(t)
	total := 100.0
	a := NewAsset("test", 100, 100, 70, total, 100, true)

	assert.Equal(.7, a.CalculatePercentageFromTotal(total))
}
func Test_Should_CalculateFinalContribution(t *testing.T) {
	assert := assert.New(t)
	total := 100.0
	contribution := 100.0
	a := NewAsset("test", 50, 100, 70, total, contribution, true)

	assert.EqualValues(30, a.CalculateFinalContribution(contribution, total))
}
