package nad

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	ctlcniv1 "github.com/harvester/harvester/pkg/generated/controllers/k8s.cni.cncf.io/v1"
	nadv1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	"github.com/vishvananda/netlink"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/klog"

	"github.com/harvester/harvester-network-controller/pkg/config"
	"github.com/harvester/harvester-network-controller/pkg/network/vlan"
	"github.com/harvester/harvester-network-controller/pkg/utils"
)

const ControllerName = "harvester-network-nad-controller"

type Handler struct {
	nadCache ctlcniv1.NetworkAttachmentDefinitionCache
}

func Register(ctx context.Context, management *config.Management) error {
	nad := management.CniFactory.K8s().V1().NetworkAttachmentDefinition()

	handler := &Handler{
		nadCache: nad.Cache(),
	}

	nad.OnChange(ctx, ControllerName, handler.OnChange)
	nad.OnRemove(ctx, ControllerName, handler.OnRemove)

	return nil
}

func (h Handler) OnChange(_ string, nad *nadv1.NetworkAttachmentDefinition) (*nadv1.NetworkAttachmentDefinition, error) {
	if nad == nil || nad.DeletionTimestamp != nil {
		return nil, nil
	}
	if nad.Spec.Config == "" || nad.Labels == nil || nad.Labels[utils.KeyVlanLabel] == "" ||
		nad.Labels[utils.KeyClusterNetworkLabel] == "" {
		return nad, nil
	}

	klog.Infof("nad configuration %s has been changed: %s", nad.Name, nad.Spec.Config)

	if err := h.addLocalArea(nad); err != nil {
		return nil, err
	}

	return nad, nil
}

func (h Handler) OnRemove(_ string, nad *nadv1.NetworkAttachmentDefinition) (*nadv1.NetworkAttachmentDefinition, error) {
	if nad == nil {
		return nil, nil
	}
	if nad.Spec.Config == "" || nad.Labels == nil || nad.Labels[utils.KeyVlanLabel] == "" ||
		nad.Labels[utils.KeyClusterNetworkLabel] == "" {
		return nad, nil
	}
	klog.Infof("nad configuration %s has been deleted.", nad.Name)
	// Skip the case that there are nads with the same cluster network and VLAN id.
	if ok, err := h.existDuplicateNad(nad.Labels[utils.KeyVlanLabel], nad.Labels[utils.KeyClusterNetworkLabel]); err != nil {
		return nil, err
	} else if ok {
		return nad, nil
	}

	if err := h.removeLocalArea(nad); err != nil {
		return nil, err
	}

	return nad, nil
}

func (h Handler) addLocalArea(nad *nadv1.NetworkAttachmentDefinition) error {
	vlanID, err := strconv.Atoi(nad.Labels[utils.KeyVlanLabel])
	if err != nil {
		return fmt.Errorf("invalid vlanconfig %s", nad.Labels[utils.KeyVlanLabel])
	}

	v, err := vlan.GetVlan(nad.Labels[utils.KeyClusterNetworkLabel])
	if err != nil && !errors.As(err, &netlink.LinkNotFoundError{}) {
		return err
	} else if errors.As(err, &netlink.LinkNotFoundError{}) {
		return nil
	}

	layer3NetworkConf := &utils.Layer3NetworkConf{}
	if nad.Annotations != nil && nad.Annotations[utils.KeyNetworkConf] != "" {
		if layer3NetworkConf, err = utils.NewLayer3NetworkConf(nad.Annotations[utils.KeyNetworkConf]); err != nil {
			return err
		}
	}

	return v.AddLocalArea(uint16(vlanID), layer3NetworkConf.CIDR)
}

func (h Handler) existDuplicateNad(vlanIdStr, cn string) (bool, error) {
	nads, err := h.nadCache.List("", labels.Set(map[string]string{
		utils.KeyVlanLabel:           vlanIdStr,
		utils.KeyClusterNetworkLabel: cn,
	}).AsSelector())
	if err != nil {
		return false, err
	}

	return len(nads) > 1, nil
}

func (h Handler) removeLocalArea(nad *nadv1.NetworkAttachmentDefinition) error {
	vlanID, err := strconv.Atoi(nad.Labels[utils.KeyVlanLabel])
	if err != nil {
		return fmt.Errorf("invalid vlanconfig %s", nad.Labels[utils.KeyVlanLabel])
	}

	v, err := vlan.GetVlan(nad.Labels[utils.KeyClusterNetworkLabel])
	if err != nil && !errors.As(err, &netlink.LinkNotFoundError{}) {
		return err
	} else if errors.As(err, &netlink.LinkNotFoundError{}) {
		return nil
	}

	layer3NetworkConf := &utils.Layer3NetworkConf{}
	if nad.Annotations != nil && nad.Annotations[utils.KeyNetworkConf] != "" {
		layer3NetworkConf, err = utils.NewLayer3NetworkConf(nad.Annotations[utils.KeyNetworkConf])
		if err != nil {
			return err
		}
	}

	return v.RemoveLocalArea(uint16(vlanID), layer3NetworkConf.CIDR)
}
