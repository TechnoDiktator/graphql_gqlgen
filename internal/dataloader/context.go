package loader

import "context"

type loadersKey struct{}


func WithLoaders(ctx context.Context, loaders *Loaders) context.Context {
	return context.WithValue(ctx, loadersKey{}, loaders)
}

func For(ctx context.Context) *Loaders {

	loaders, ok := ctx.Value(loadersKey{}).(*Loaders)

	if !ok {
		panic("loaders not found in context")
	}

	return loaders
}
