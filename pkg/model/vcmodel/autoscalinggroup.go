package vcmodel

import (
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/pkg/model"
	"github.com/golang/glog"
)

// AutoscalingGroupModelBuilder configures AutoscalingGroup objects
type AutoscalingGroupModelBuilder struct {
	*VCModelContext

	BootstrapScript *model.BootstrapScript
}

var _ fi.ModelBuilder = &AutoscalingGroupModelBuilder{}

func (b *AutoscalingGroupModelBuilder) Build(c *fi.ModelBuilderContext) error {
	glog.Warning("AutoscalingGroupModelBuilder.Build not implemented for VC")
	return nil
}
