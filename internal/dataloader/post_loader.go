package loader

import (
	"github.com/graph-gophers/dataloader/v7"
	entity "github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
)

type PostLoader struct {
	loader *dataloader.Loader[int64, *entity.User]
}
