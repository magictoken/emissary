package v2

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func convert(src, dst runtime.Object) error {
	s, err := SchemeBuilder.Build()
	if err != nil {
		return err
	}
	return s.Convert(src, dst, nil)
}
