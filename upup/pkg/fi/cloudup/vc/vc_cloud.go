package vc

import (
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kubernetes/federation/pkg/dnsprovider"
	"github.com/golang/glog"
	k8sroute53 "k8s.io/kubernetes/federation/pkg/dnsprovider/providers/aws/route53"
	"fmt"
)

type VCCloud struct {
	// dummy field
	name string
	Region  string
}

var _ fi.Cloud = &VCCloud{}

func (c *VCCloud) ProviderID() fi.CloudProviderID {
	return fi.CloudProviderVC
}

func (c *VCCloud) DNS() (dnsprovider.Interface, error) {
	glog.Warning("DNS() not implemented on VC")
	provider, err := dnsprovider.GetDnsProvider(k8sroute53.ProviderName, nil)
	if err != nil {
		return nil, fmt.Errorf("Error building (k8s) DNS provider: %v", err)
	}
	return provider, nil

}

func (c *VCCloud) FindVPCInfo(id string) (*fi.VPCInfo, error) {
	glog.Warningf("FindVPCInfo not (yet) implemented on GCE")
	return nil, nil
}
