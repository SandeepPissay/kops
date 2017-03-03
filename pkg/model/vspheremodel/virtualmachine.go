package vspheremodel

import (
	"k8s.io/kops/upup/pkg/fi"
	"fmt"
	"k8s.io/kops/upup/pkg/fi/cloudup/vspheretasks"
)

// AutoscalingGroupModelBuilder configures AutoscalingGroup objects
type VirtualMachineModelBuilder struct {
	*VSphereModelContext
}

func (b *VirtualMachineModelBuilder) Build(c *fi.ModelBuilderContext) error {
	fmt.Print("In VirtualMachineModelBuilder.Build function!!")
	var masterVmTask = &vspheretasks.VirtualMachine{
		Name: fi.String("dummyMasterVmName"),
		VMTemplateName: fi.String("dummyVmTemplate"),
	}
	var nodeVmTask = &vspheretasks.VirtualMachine{
		Name: fi.String("dummyWorkerVmName"),
		VMTemplateName: fi.String("dummyVmTemplate"),
	}
	c.AddTask(masterVmTask)
	c.AddTask(nodeVmTask)
	return nil
}