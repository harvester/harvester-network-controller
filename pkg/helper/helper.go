package helper

import (
	"net"

	ctlcni "github.com/harvester/harvester/pkg/generated/controllers/k8s.cni.cncf.io"
	ctlcniv1 "github.com/harvester/harvester/pkg/generated/controllers/k8s.cni.cncf.io/v1"
	nadv1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/harvester/harvester-network-controller/pkg/utils"
)

type NetHelper struct {
	nadClient ctlcniv1.NetworkAttachmentDefinitionClient
}

func New(cniFactory *ctlcni.Factory) *NetHelper {
	return &NetHelper{
		nadClient: cniFactory.K8s().V1().NetworkAttachmentDefinition(),
	}
}

func (n *NetHelper) GetVLANLayer3Network(selectedNetwork *nadv1.NetworkSelectionElement, serverIPAddr string) (*utils.Layer3NetworkConf, error) {
	cidr, gw, err := obtainCIDRAndGw(selectedNetwork.InterfaceRequest, net.ParseIP(serverIPAddr))
	if err != nil {
		return nil, err
	}

	return &utils.Layer3NetworkConf{
		Mode:         utils.Auto,
		ServerIPAddr: serverIPAddr,
		CIDR:         cidr.String(),
		Gateway:      gw.String(),
		Connectivity: utils.Unknown,
	}, nil
}

func (n *NetHelper) RecordToNad(selectedNetwork *nadv1.NetworkSelectionElement, networkConf *utils.Layer3NetworkConf) error {
	nad, err := n.nadClient.Get(selectedNetwork.Namespace, selectedNetwork.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	nadCopy := nad.DeepCopy()
	if nadCopy.Annotations == nil {
		nadCopy.Annotations = make(map[string]string)
	}

	confStr, err := networkConf.ToString()
	if err != nil {
		return err
	}

	nadCopy.Annotations[utils.KeyNetworkConf] = confStr

	_, err = n.nadClient.Update(nadCopy)
	return err
}
