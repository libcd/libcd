package libcd

import (
	"fmt"

	"github.com/libcd/libcd/parse"
)

// Spec defines the pipeline configuration and exeuction.
type Spec struct {
	// Volumes defines a list of all container volumes.
	Volumes []*Volume `json:"volumes"`

	// Containers defines a list of all containers in the pipeline.
	Containers []*Container `json:"objects"`

	// Nodes defines the container execution tree.
	Nodes parse.Tree `json:"nodes"`
}

// lookupContainer is a helper funciton that returns the named container from
// the slice of containers.
func (s *Spec) lookupContainer(name string) (*Container, error) {
	for _, container := range s.Containers {
		if container.Name == name {
			return container, nil
		}
	}
	return nil, fmt.Errorf("runner: unknown container %s", name)
}
