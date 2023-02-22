package adapters

import (
	"balancer/internal/boundaries"
	"balancer/internal/domain"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	mockedRepository[T interface{}] struct {
		mockedDatabase []T
		mockInsert     func(e T)
		mockGetAll     func(filter map[string]interface{}) []T
		mockGetFirst   func(filter map[string]interface{}) T
		mockReplace    func(filter map[string]interface{}, entity T)
		mockDeleteAll  func(filter map[string]interface{})
	}
)

func Test_Should_CreateAssetsGroup(t *testing.T) {
	assert := assert.New(t)

	r := newMockedRepository[*domain.AssetsGroup]()
	r.mockInsert = func(e *domain.AssetsGroup) {
		r.mockedDatabase = append(r.mockedDatabase, e)
	}
	s := NewAssetsBalancerUseCase(r)
	input := Input_Test_Should_CreateAssetsGroup()
	res, err := s.CreateAssetsGroup(context.Background(), input)
	if !assert.Nil(err) ||
		!assert.NotNil(res) ||
		!assert.NotEmpty(res.Id) ||
		!assert.NotEmpty(res.Assets) ||
		!assert.Len(r.mockedDatabase, 1) {
		t.FailNow()
	}
}
func Test_Should_UpdateAsset(t *testing.T) {
	assert := assert.New(t)
	targetAsset := domain.NewAsset("testTarget", 35, 300, 303, 394, 100, true)
	assets := []*domain.Asset{
		targetAsset,
		domain.NewAsset("test", 15, 300, 303, 394, 100, true),
	}
	assetsGroup := domain.NewAssetGroup("test", assets, 100)
	r := newMockedRepository[*domain.AssetsGroup]()
	r.mockGetFirst = func(filter map[string]interface{}) *domain.AssetsGroup {
		return assetsGroup
	}
	r.mockReplace = func(filter map[string]interface{}, entity *domain.AssetsGroup) {
		assetsGroup = entity
	}

	s := NewAssetsBalancerUseCase(r)
	input := &boundaries.UpdateAssetInput{
		Id:           targetAsset.Id,
		GroupId:      assetsGroup.Id,
		Label:        "testTargetUpdated",
		CurrentValue: 306,
		Include:      true,
	}
	res, err := s.UpdateAsset(context.Background(), input)
	if !assert.Nil(err) ||
		!assert.EqualValues(res, assetsGroup) ||
		!assert.EqualValues(targetAsset.CurrentValue, input.CurrentValue) {
		t.FailNow()
	}
}

func Test_Should_DeleteAsset(t *testing.T) {
	assert := assert.New(t)
	targetAsset := domain.NewAsset("testTarget", 35, 300, 303, 394, 100, true)
	assets := []*domain.Asset{
		targetAsset,
	}
	assetsGroup := domain.NewAssetGroup("test", assets, 100)
	r := newMockedRepository[*domain.AssetsGroup]()
	r.mockGetFirst = func(filter map[string]interface{}) *domain.AssetsGroup {
		return assetsGroup
	}
	r.mockReplace = func(filter map[string]interface{}, entity *domain.AssetsGroup) {
		assetsGroup = entity
	}

	s := NewAssetsBalancerUseCase(r)
	input := &boundaries.DeleteAssetInput{
		Id:      targetAsset.Id,
		GroupId: assetsGroup.Id,
	}
	res, err := s.DeleteAsset(context.Background(), input)
	if !assert.Nil(err) ||
		!assert.EqualValues(res, assetsGroup) ||
		!assert.Len(assetsGroup.Assets, 0) {
		t.FailNow()
	}
}

func Test_Should_DeleteAssetsGroup(t *testing.T) {
	assert := assert.New(t)
	assets := []*domain.Asset{}
	assetsGroup := domain.NewAssetGroup("test", assets, 100)
	r := newMockedRepository[*domain.AssetsGroup]()
	r.mockGetFirst = func(filter map[string]interface{}) *domain.AssetsGroup {
		return assetsGroup
	}
	r.mockDeleteAll = func(filter map[string]interface{}) {
	}

	s := NewAssetsBalancerUseCase(r)
	input := &boundaries.DeleteAssetsGroupInput{
		Id: assetsGroup.Id,
	}
	err := s.DeleteAssetsGroup(context.Background(), input)
	if !assert.Nil(err) {
		t.FailNow()
	}
}

func Input_Test_Should_Not_CreateAssetsGroupWithInvalidInput() *boundaries.CreateAssetsGroupInput {
	return &boundaries.CreateAssetsGroupInput{
		Assets: []boundaries.CreateAssetInput{
			{
				Label:         "RF",
				Score:         60,
				PreviousValue: 3732.87,
				CurrentValue:  3730.87,
				Include:       false,
			},
			{
				Label:         "Ações",
				Score:         10,
				PreviousValue: 619.57,
				CurrentValue:  519.32,
				Include:       true,
			},
			{
				Label:         "FIIs",
				Score:         30,
				PreviousValue: 1220.44,
				CurrentValue:  1500,
				Include:       true,
			},
		},
	}
}
func Input_Test_Should_CreateAssetsGroup() *boundaries.CreateAssetsGroupInput {
	return &boundaries.CreateAssetsGroupInput{
		Assets: []boundaries.CreateAssetInput{
			{
				Label:         "RF",
				Score:         60,
				PreviousValue: 3732.87,
				CurrentValue:  3730.87,
				Include:       true,
			},
			{
				Label:         "Ações",
				Score:         10,
				PreviousValue: 619.57,
				CurrentValue:  519.32,
				Include:       true,
			},
			{
				Label:         "FIIs",
				Score:         30,
				PreviousValue: 1220.44,
				CurrentValue:  1500,
				Include:       true,
			},
		},
	}
}
func newMockedRepository[T interface{}]() *mockedRepository[T] {
	return &mockedRepository[T]{
		mockedDatabase: []T{},
	}
}

func (mr *mockedRepository[T]) Insert(ctx context.Context, e T) {
	mr.mockInsert(e)
}
func (mr *mockedRepository[T]) GetAll(ctx context.Context, filter map[string]interface{}) []T {
	return mr.mockGetAll(filter)
}
func (mr *mockedRepository[T]) GetFirst(ctx context.Context, filter map[string]interface{}) T {
	return mr.mockGetFirst(filter)
}
func (mr *mockedRepository[T]) Replace(ctx context.Context, filter map[string]interface{}, entity T) {
	mr.mockReplace(filter, entity)
}
func (mr *mockedRepository[T]) DeleteAll(ctx context.Context, filter map[string]interface{}) {
	mr.mockDeleteAll(filter)
}
