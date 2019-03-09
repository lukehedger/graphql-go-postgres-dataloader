package dataloaders

import (
	"time"

	"github.com/lukehedger/graphql-go-postgres-dataloader"
)

func NewPersonResolverLoader() *PersonLoader {
	return &PersonLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(ids []string) ([]*api.Person, []error) {
			persons := make([]*api.Person, len(ids))
			errors := make([]error, len(ids))

			for i := range ids {
				persons[i] = &api.Person{ID: "id", Name: "person"}
			}
			return persons, errors
		},
	}
}
