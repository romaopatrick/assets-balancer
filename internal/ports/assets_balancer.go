package ports

import (
	"balancer/internal/boundaries"
	"balancer/internal/domain"
	"context"
)

type (
	AssetBalancerUseCase interface {
		CreateAssetsGroup(ctx context.Context, input *boundaries.CreateAssetsGroupInput) (*domain.AssetsGroup, error)
		CreateAsset(ctx context.Context, input *boundaries.CreateAssetForGroupInput) (*domain.AssetsGroup, error)
		UpdateAsset(ctx context.Context, input *boundaries.UpdateAssetInput) (*domain.AssetsGroup, error)
		UpdateAssetsGroup(ctx context.Context, input *boundaries.UpdateAssetsGroup) (*domain.AssetsGroup, error)
		DeleteAsset(ctx context.Context, input *boundaries.DeleteAssetInput) (*domain.AssetsGroup, error)
		DeleteAssetsGroup(ctx context.Context, input *boundaries.DeleteAssetsGroupInput) error

		GetAssetsGroups(ctx context.Context) []*domain.AssetsGroup
		GetAssetsGroup(ctx context.Context, input *boundaries.GetAssetsGroupInput) *domain.AssetsGroup
	}
)
