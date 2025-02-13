package hypervsdk

import (
	"github.com/rokukoo/hypervctl/pkg/hypervsdk/errors"
	"github.com/rokukoo/hypervctl/pkg/wmiext"
)

const (
	HyperVNamespace = `root\virtualization\v2`
)

func NewLocalHyperVService() (*wmiext.Service, error) {
	var (
		s   *wmiext.Service
		err error
	)
	if s, err = wmiext.NewLocalService(HyperVNamespace); err != nil {
		return nil, errors.TranslateCommonHyperVWmiError(err)
	}
	return s, nil
}
