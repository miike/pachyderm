package fuse

import (
	"github.com/hanwen/go-fuse/v2/fs"

	"github.com/pachyderm/pachyderm/src/client/pkg/errors"
	"github.com/pachyderm/pachyderm/src/server/pkg/uuid"
)

// Options is for configuring fuse mounts. Any of the fields may be left nil
// and `nil` itself is a valid set of Options which uses the default for
// everything.
type Options struct {
	Fuse *fs.Options

	// Write indicates that the pfs mount should allow writes.
	// Writes will be written back to the filesystem.
	Write bool

	// Branches is a map from repos to branches, if a repo is unspecified then
	// the "master" will be used.
	Branches map[string]string

	// Unmount is a channel that will be closed when the filesystem has been
	// unmounted. It can be nil in which case it's ignored.
	Unmount chan struct{}
}

func (o *Options) getFuse() *fs.Options {
	if o == nil || o.Fuse == nil {
		// We always return a struct here because otherwise the defaults that
		// fuse sets make files inaccessible.
		return &fs.Options{}
	}
	return o.Fuse
}

func (o *Options) getBranches() map[string]string {
	if o == nil || o.Branches == nil {
		return make(map[string]string)
	}
	return o.Branches
}

func (o *Options) getWrite() bool {
	if o == nil {
		return false
	}
	return o.Write
}

func (o *Options) getUnmount() chan struct{} {
	if o == nil {
		return nil
	}
	return o.Unmount
}

func (o *Options) validate() error {
	if o == nil {
		return nil
	}
	if o.Write {
		for _, commit := range o.Branches {
			if uuid.IsUUIDWithoutDashes(commit) {
				return errors.Errorf("can't mount commits %s in Write mode (mount a branch instead)", commit)
			}
		}
	}
	return nil
}
