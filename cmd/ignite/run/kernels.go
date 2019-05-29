package run

import (
	"fmt"
	"github.com/luxas/ignite/pkg/metadata/kernmd"
	"github.com/luxas/ignite/pkg/util"
)

type KernelsOptions struct {
	Kernels []*kernmd.KernelMetadata
}

func Kernels(ko *KernelsOptions) error {
	o := util.NewOutput()
	defer o.Flush()

	o.Write("KERNEL ID", "CREATED", "SIZE", "NAME")
	for _, md := range ko.Kernels {
		size, err := md.Size()
		if err != nil {
			return fmt.Errorf("failed to get size for %s %q: %v", md.Type, md.ID, err)
		}

		o.Write(md.ID, md.Created, util.ByteCountDecimal(size), md.Name)
	}

	return nil
}
