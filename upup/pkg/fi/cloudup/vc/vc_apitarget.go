package vc

import "k8s.io/kops/upup/pkg/fi"

type VCAPITarget struct {
	Cloud *VCCloud
}

var _ fi.Target = &VCAPITarget{}

func NewVCAPITarget(cloud *VCCloud) *VCAPITarget {
	return &VCAPITarget{
		Cloud: cloud,
	}
}

func (t *VCAPITarget) Finish(taskMap map[string]fi.Task) error {
	return nil
}

func (t *VCAPITarget) ProcessDeletions() bool {
	return true
}
