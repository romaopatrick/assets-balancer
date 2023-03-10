package adapters

import (
	"context"
	"errors"

	"github.com/romaopatrick/assets-balancer/internal/boundaries"
	"github.com/romaopatrick/assets-balancer/internal/domain"
	"github.com/romaopatrick/assets-balancer/internal/ports"

	"golang.org/x/exp/slices"
)

type (
	AssetsBalancerService struct {
		repository ports.Repository[*domain.AssetsGroup]
	}
)

func NewAssetsBalancerUseCase(repository ports.Repository[*domain.AssetsGroup]) ports.AssetBalancerUseCase {
	return &AssetsBalancerService{
		repository: repository,
	}
}

func (abs *AssetsBalancerService) GetAssetsGroups(ctx context.Context) []*domain.AssetsGroup {
	result := abs.repository.GetAll(ctx, nil)
	return result
}

func (abs *AssetsBalancerService) GetAssetsGroup(
	ctx context.Context, input *boundaries.GetAssetsGroupInput) *domain.AssetsGroup {
	return abs.repository.GetFirst(ctx, map[string]interface{}{
		"id": input.Id,
	})
}

func (abs *AssetsBalancerService) CreateAssetsGroup(
	ctx context.Context, input *boundaries.CreateAssetsGroupInput) (*domain.AssetsGroup, error) {

	assets := []*domain.Asset{}
	for _, v := range input.Assets {
		assets = append(assets, domain.NewAsset(
			v.Label, v.Score, v.PreviousValue, v.CurrentValue,
			input.CurrentTotal(), input.ContributionTotal, v.Include))
	}
	assetsGroup := domain.NewAssetGroup(input.Label, assets, input.ContributionTotal)

	abs.repository.Insert(ctx, assetsGroup)

	return assetsGroup, nil
}

func (abs *AssetsBalancerService) CreateAsset(
	ctx context.Context, input *boundaries.CreateAssetForGroupInput) (*domain.AssetsGroup, error) {

	assetsGroup := abs.repository.GetFirst(ctx, map[string]interface{}{
		"id": input.GroupId,
	})

	if assetsGroup == nil {
		return nil, errors.New(domain.ASSETS_GROUP_NOT_FOUND)
	}

	total := assetsGroup.CurrentTotal() + input.CurrentValue
	a := domain.NewAsset(input.Label,
		input.Score, input.PreviousValue, input.CurrentValue,
		total, assetsGroup.ContributionTotal, input.Include)

	assetsGroup.Assets = append(assetsGroup.Assets, a)
	balance(assetsGroup)

	abs.repository.Replace(ctx, map[string]interface{}{
		"id": input.GroupId,
	}, assetsGroup)

	return assetsGroup, nil
}

func (abs *AssetsBalancerService) UpdateAssetsGroup(
	ctx context.Context, input *boundaries.UpdateAssetsGroup) (*domain.AssetsGroup, error) {
	assetsGroup := abs.repository.GetFirst(ctx, map[string]interface{}{
		"id": input.Id,
	})

	if assetsGroup == nil {
		return nil, errors.New(domain.ASSETS_GROUP_NOT_FOUND)
	}

	if input.Label != "" {
		assetsGroup.Label = input.Label
	}
	if input.ContributionTotal != 0 {
		assetsGroup.ContributionTotal = input.ContributionTotal
	}
	balance(assetsGroup)

	abs.repository.Replace(ctx, map[string]interface{}{
		"id": input.Id,
	}, assetsGroup)

	return assetsGroup, nil

}
func (abs *AssetsBalancerService) UpdateAsset(
	ctx context.Context, input *boundaries.UpdateAssetInput) (*domain.AssetsGroup, error) {

	assetsGroup := abs.repository.GetFirst(ctx, map[string]interface{}{
		"id": input.GroupId, "assets": map[string]interface{}{
			"$elemMatch": map[string]interface{}{
				"id": input.Id,
			},
		},
	})

	if assetsGroup == nil {
		return nil, errors.New(domain.ASSETS_GROUP_NOT_FOUND)
	}

	idx := slices.IndexFunc(assetsGroup.Assets, func(a *domain.Asset) bool {
		return a.Id == input.Id
	})

	updateAsset(assetsGroup.Assets[idx], input)
	balance(assetsGroup)

	abs.repository.Replace(ctx, map[string]interface{}{
		"id": input.GroupId,
	}, assetsGroup)

	return assetsGroup, nil
}

func (abs *AssetsBalancerService) DeleteAsset(
	ctx context.Context, input *boundaries.DeleteAssetInput) (*domain.AssetsGroup, error) {
	assetsGroup := abs.repository.GetFirst(ctx, map[string]interface{}{
		"id": input.GroupId, "assets": map[string]interface{}{
			"$elemMatch": map[string]interface{}{
				"id": input.Id,
			},
		},
	})
	if assetsGroup == nil {
		return nil, errors.New(domain.ASSETS_GROUP_NOT_FOUND)
	}

	idx := slices.IndexFunc(assetsGroup.Assets, func(a *domain.Asset) bool {
		return a.Id == input.Id
	})

	assetsGroup.Assets = removeAsset(assetsGroup.Assets, idx)
	balance(assetsGroup)

	abs.repository.Replace(ctx, map[string]interface{}{
		"id": input.GroupId,
	}, assetsGroup)

	return assetsGroup, nil
}

func (abs *AssetsBalancerService) DeleteAssetsGroup(
	ctx context.Context, input *boundaries.DeleteAssetsGroupInput) error {
	assetsGroup := abs.repository.GetFirst(ctx,
		map[string]interface{}{
			"id": input.Id,
		})

	if assetsGroup == nil {
		return errors.New(domain.ASSETS_GROUP_NOT_FOUND)
	}

	abs.repository.DeleteAll(ctx, map[string]interface{}{
		"id": input.Id,
	})

	return nil
}
func removeAsset(assets []*domain.Asset, idx int) []*domain.Asset {
	return append(assets[:idx], assets[idx+1:]...)
}

func updateAsset(a *domain.Asset, input *boundaries.UpdateAssetInput) {
	a.CurrentValue = input.CurrentValue
	a.PreviousValue = input.PreviousValue
	a.Score = input.Score
	a.Include = input.Include
	if input.Label != "" {
		a.Label = input.Label
	}
}

func balance(group *domain.AssetsGroup) {
	for _, a := range group.Assets {
		if a.Include {
			a.PercentageFromTotal = a.CalculatePercentageFromTotal(group.CurrentTotal())
			a.ValueVariation = a.CalculateValueVariation()
			a.FinalContribution = a.CalculateFinalContribution(group.ContributionTotal, group.CurrentTotal())
		}
	}
}
