package domain

import "github.com/google/uuid"

type (
	AssetsGroup struct {
		Id                uuid.UUID
		Assets            []*Asset
		Label             string
		ContributionTotal float64
	}
	Asset struct {
		Id                  uuid.UUID
		Label               string
		Score               float32
		PreviousValue       float64
		CurrentValue        float64
		ValueVariation      float64
		PercentageFromTotal float64
		FinalContribution   float64
		Include             bool
	}
)

func (a *Asset) CalculateValueVariation() float64 {
	if a.PreviousValue == 0 {
		return 0
	}
	return (a.CurrentValue - a.PreviousValue) / a.PreviousValue
}
func (a *Asset) CalculatePercentageFromTotal(currentTotal float64) float64 {
	if currentTotal == 0 {
		return 0
	}
	return a.CurrentValue / currentTotal
}
func (a *Asset) CalculateFinalContribution(contributionTotal, currentTotal float64) float64 {
	return ((currentTotal + contributionTotal) * float64(a.Score) / 100) - a.CurrentValue
}

func (ag *AssetsGroup) CurrentTotal() float64 {
	result := 0.
	for _, v := range ag.Assets {
		if v.Include {
			result += v.CurrentValue
		}
	}

	return result
}

func NewAsset(label string, score float32, previousV, currentV, currentT, contributionT float64, include bool) *Asset {
	asset := &Asset{
		Label:         label,
		Score:         score,
		PreviousValue: previousV,
		CurrentValue:  currentV,
		Include:       include,
		Id:            uuid.New(),
	}
	if asset.Include {
		asset.ValueVariation = asset.CalculateValueVariation()
		asset.PercentageFromTotal = asset.CalculatePercentageFromTotal(currentT)
		asset.FinalContribution = asset.CalculateFinalContribution(contributionT, currentT)
	}

	return asset
}
func NewAssetGroup(label string, assets []*Asset, contributionT float64) *AssetsGroup {
	return &AssetsGroup{
		Id:                uuid.New(),
		Assets:            assets,
		Label:             label,
		ContributionTotal: contributionT,
	}
}
