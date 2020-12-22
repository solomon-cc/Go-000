// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"http-server/internal/dao"
	"http-server/internal/service"
	"http-server/internal/server/grpc"
	"http-server/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
