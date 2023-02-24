package boundaries

import "github.com/google/uuid"

type (
	CreateAssetsGroupInput struct {
		Assets            []CreateAssetInput
		Label             string
		ContributionTotal float64
	}
	CreateAssetInput struct {
		Label         string
		Score         float32
		PreviousValue float64
		CurrentValue  float64
		Include       bool
	}
	CreateAssetForGroupInput struct {
		GroupId       uuid.UUID
		Label         string
		Score         float32
		PreviousValue float64
		CurrentValue  float64
		Include       bool
	}
	UpdateAssetInput struct {
		Id            uuid.UUID
		GroupId       uuid.UUID
		Label         string
		Score         float32
		PreviousValue float64
		CurrentValue  float64
		Include       bool
	}
	UpdateAssetsGroup struct {
		Id                uuid.UUID
		ContributionTotal float64
		Label             string
	}
	DeleteAssetInput struct {
		Id      uuid.UUID
		GroupId uuid.UUID
	}
	DeleteAssetsGroupInput struct {
		Id uuid.UUID
	}

	GetAssetsGroupInput struct {
		Id uuid.UUID
	}
)

func (cagi *CreateAssetsGroupInput) CurrentTotal() (result float64) {
	for _, v := range cagi.Assets {
		if v.Include {
			result += v.CurrentValue
		}
	}

	return result
}
