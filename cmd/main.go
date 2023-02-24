package main

import (
	"context"
	"os"

	"github.com/romaopatrick/assets-balancer/internal/adapters"
	"github.com/romaopatrick/assets-balancer/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

func main() {
	c := provideDependencies()
	if err := c.Invoke(startupApplication); err != nil {
		panic(err)
	}
}

func startupApplication(
	eng *gin.Engine,
	cl *mongo.Client,
	ph *adapters.AssetsBalancerHandler) {
	defer cl.Disconnect(context.Background())
	adapters.ConfigureRouter(eng, ph)
	eng.Run(":8081")
}

func provideDependencies() (c *dig.Container) {
	c = dig.New()
	c.Provide(initializeViper)
	c.Provide(gin.Default)
	provideMongo(c)
	provideRepositories(c)
	provideHandlers(c)
	provideUseCases(c)

	return
}

func provideMongo(c *dig.Container) {
	c.Provide(adapters.NewClientOptions)
	c.Provide(adapters.NewMongoClient)
	c.Provide(adapters.NewMongoDatabase)
}

func initializeViper() *viper.Viper {
	v := viper.New()
	v.AddConfigPath("../configs")
	v.SetConfigType("json")
	v.SetConfigName(os.Getenv("env"))
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	return v
}

func provideRepositories(c *dig.Container) {
	c.Provide(adapters.NewMongoDbRepository[*domain.AssetsGroup])
}
func provideHandlers(c *dig.Container) {
	c.Provide(adapters.NewAssetsBalancerHandler)
}
func provideUseCases(c *dig.Container) {
	c.Provide(adapters.NewAssetsBalancerUseCase)
}
