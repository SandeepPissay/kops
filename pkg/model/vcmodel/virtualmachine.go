package vcmodel

import (
	"k8s.io/kops/upup/pkg/fi"
	"fmt"
	"k8s.io/kops/upup/pkg/fi/cloudup/vctasks"
)

// AutoscalingGroupModelBuilder configures AutoscalingGroup objects
type VirtualMachineModelBuilder struct {
	*VCModelContext
}

func (b *VirtualMachineModelBuilder) Build(c *fi.ModelBuilderContext) error {
	fmt.Print("In VirtualMachineModelBuilder.Build function!!")
	var masterVmTask = &vctasks.VirtualMachine{
		Name: fi.String("dummyMasterVmName"),
		VMTemplateName: fi.String("dummyVmTemplate"),
	}
	var nodeVmTask = &vctasks.VirtualMachine{
		Name: fi.String("dummyWorkerVmName"),
		VMTemplateName: fi.String("dummyVmTemplate"),
	}
	c.AddTask(masterVmTask)
	c.AddTask(nodeVmTask)
	return nil
}