package ipfix

import (
	"fmt"
	"loxiflow/pkg/dpebpf"
	"time"

	"github.com/vmware/go-ipfix/pkg/entities"
	"github.com/vmware/go-ipfix/pkg/exporter"
	"github.com/vmware/go-ipfix/pkg/registry"
)

func Start(ctMap *dpebpf.Maps, colIP string, collPort, setSecond int, colProto string) {

	address := fmt.Sprintf("%s:%d", colIP, collPort) // "127.0.0.1:4739"
	SetSecond := setSecond
	// Define ExporterInput with required fields
	expInput := exporter.ExporterInput{
		CollectorAddress:    address,
		CollectorProtocol:   colProto,         // Can be "udp" or "tcp"
		ObservationDomainID: 12345,            // Example Observation Domain ID
		TempRefTimeout:      1800,             // Template Refresh Timeout (in seconds)
		IsIPv6:              false,            // Set to true if using IPv6
		SendJSONRecord:      false,            // Set to true if you want to send JSON format records
		JSONBufferLen:       0,                // Only used if SendJSONRecord is true
		CheckConnInterval:   10 * time.Second, // Interval to check connection (used for TCP mostly)
	}

	// Initialize the IPFIX exporter using ExporterInput
	expProcess, err := exporter.InitExportingProcess(expInput)
	if err != nil {
		panic(err)
	}
	defer expProcess.CloseConnToCollector()

	// Register elements to the registry
	registry.LoadRegistry()
	// Template Set
	templateID := expProcess.NewTemplateID()
	templateSet := entities.NewSet(false)
	err = templateSet.PrepareSet(entities.Template, templateID)
	// Make tempelte
	tmpelements := make([]entities.InfoElementWithValue, 0)
	flowTypeElement, _ := registry.GetInfoElement("flowType", registry.AntreaEnterpriseID)
	ie, _ := entities.DecodeAndCreateInfoElementWithValue(flowTypeElement, nil)
	tmpelements = append(tmpelements, ie)

	flowStartElement, _ := registry.GetInfoElement("flowStartSeconds", registry.IANAEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(flowStartElement, nil)
	tmpelements = append(tmpelements, ie)

	flowEndElement, _ := registry.GetInfoElement("flowEndSeconds", registry.IANAEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(flowEndElement, nil)
	tmpelements = append(tmpelements, ie)

	flowEndReason, _ := registry.GetInfoElement("flowEndReason", registry.IANAEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(flowEndReason, nil)
	tmpelements = append(tmpelements, ie)

	sourcePortElement, _ := registry.GetInfoElement("sourceTransportPort", registry.IANAEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(sourcePortElement, nil)
	tmpelements = append(tmpelements, ie)

	destPortElement, _ := registry.GetInfoElement("destinationTransportPort", registry.IANAEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(destPortElement, nil)
	tmpelements = append(tmpelements, ie)

	protocolElement, _ := registry.GetInfoElement("protocolIdentifier", registry.IANAEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(protocolElement, nil)
	tmpelements = append(tmpelements, ie)

	packetsElement, _ := registry.GetInfoElement("packetTotalCount", registry.IANAEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(packetsElement, nil)
	tmpelements = append(tmpelements, ie)

	bytesElement, _ := registry.GetInfoElement("octetTotalCount", registry.IANAEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(bytesElement, nil)
	tmpelements = append(tmpelements, ie)

	packetsDeltaElement, _ := registry.GetInfoElement("packetDeltaCount", registry.IANAEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(packetsDeltaElement, nil)
	tmpelements = append(tmpelements, ie)

	DeltabytesElement, _ := registry.GetInfoElement("octetDeltaCount", registry.IANAEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(DeltabytesElement, nil)
	tmpelements = append(tmpelements, ie)

	sourceIPElement, _ := registry.GetInfoElement("sourceIPv4Address", registry.IANAEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(sourceIPElement, nil)
	tmpelements = append(tmpelements, ie)

	destIPElement, _ := registry.GetInfoElement("destinationIPv4Address", registry.IANAEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(destIPElement, nil)
	tmpelements = append(tmpelements, ie)

	reverseOctetTotalCount, _ := registry.GetInfoElement("reverseOctetTotalCount", registry.IANAReversedEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(reverseOctetTotalCount, nil)
	tmpelements = append(tmpelements, ie)

	reverseOctetDeltaCount, _ := registry.GetInfoElement("reverseOctetDeltaCount", registry.IANAReversedEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(reverseOctetDeltaCount, nil)
	tmpelements = append(tmpelements, ie)

	reversePacketTotalCount, _ := registry.GetInfoElement("reversePacketTotalCount", registry.IANAReversedEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(reversePacketTotalCount, nil)
	tmpelements = append(tmpelements, ie)

	reversePacketDeltaCount, _ := registry.GetInfoElement("reversePacketDeltaCount", registry.IANAReversedEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(reversePacketDeltaCount, nil)
	tmpelements = append(tmpelements, ie)

	reverseOctetDeltaCountFromSourceNode, _ := registry.GetInfoElement("reverseOctetDeltaCountFromSourceNode", registry.AntreaEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(reverseOctetDeltaCountFromSourceNode, nil)
	tmpelements = append(tmpelements, ie)
	reverseOctetTotalCountFromSourceNode, _ := registry.GetInfoElement("reverseOctetTotalCountFromSourceNode", registry.AntreaEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(reverseOctetTotalCountFromSourceNode, nil)
	tmpelements = append(tmpelements, ie)
	reversePacketDeltaCountFromSourceNode, _ := registry.GetInfoElement("reversePacketDeltaCountFromSourceNode", registry.AntreaEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(reversePacketDeltaCount, nil)
	tmpelements = append(tmpelements, ie)
	reversePacketTotalCountFromSourceNode, _ := registry.GetInfoElement("reversePacketTotalCountFromSourceNode", registry.AntreaEnterpriseID)
	ie, _ = entities.DecodeAndCreateInfoElementWithValue(reversePacketTotalCountFromSourceNode, nil)
	tmpelements = append(tmpelements, ie)

	tcpStateElement, _ := registry.GetInfoElement("tcpState", registry.AntreaEnterpriseID)
	tcpStateIe, _ := entities.DecodeAndCreateInfoElementWithValue(tcpStateElement, nil)
	tmpelements = append(tmpelements, tcpStateIe)

	templateSet.AddRecord(tmpelements, templateID)
	expProcess.SendSet(templateSet)
	for {
		ctMap.Mutex.Lock()
		for _, v := range ctMap.CtMap {

			// Data Set
			dataSet := entities.NewSet(false)
			err = dataSet.PrepareSet(entities.Data, templateID)
			if err != nil {
				panic(err)
			}

			elements := make([]entities.InfoElementWithValue, 0)
			// Intra-Node Flow: 동일한 노드 내부에서 발생하는 흐름. 노드 내에서 Pod 간 통신이 이루어질 때 발생합니다.
			// Inter-Node Flow: 서로 다른 노드 간에 발생하는 흐름. 다른 노드에 있는 Pod 간 통신 시 기록됩니다.
			// Ingress Flow: 외부에서 클러스터 내부로 들어오는 흐름. 외부 서비스가 클러스터의 리소스에 접근할 때 발생합니다.
			// Egress Flow: 클러스터 내부에서 외부로 나가는 흐름. 클러스터 내부의 Pod가 외부 리소스와 통신할 때 발생합니다.
			// LoxiLB는 현재 무조건 외부이기 때문에, In, eg로 일단 스태틱하게 박음, 여기서는 in e 아무거나.
			flowTypeValue := entities.NewUnsigned8InfoElement(flowTypeElement, 4)
			elements = append(elements, flowTypeValue)

			// Flow가 시작된 시간
			flowStartValue := entities.NewUnsigned32InfoElement(flowStartElement, uint32(v.STs.Unix()))
			elements = append(elements, flowStartValue)

			flowEndValue := entities.NewUnsigned32InfoElement(flowEndElement, uint32(v.ETs.Unix()))
			elements = append(elements, flowEndValue)
			// 0은 아직 종료가 아님
			// Idle timeout: 흐름이 일정 시간 동안 활동이 없을 때 종료됨.
			// Active timeout: 활성 상태가 유지되지만 일정 시간이 경과하여 흐름이 종료됨.
			// End of flow: 세션이 정상적으로 완료되어 흐름이 종료됨.
			// Forced end: 비정상적인 종료, 예를 들어 네트워크 재설정 등에 의해 강제로 흐름이 종료됨.
			flowEndReasonValue := entities.NewUnsigned8InfoElement(flowEndReason, v.FlowEndReason)
			elements = append(elements, flowEndReasonValue)

			// Add Source Port
			sourcePortValue := entities.NewUnsigned16InfoElement(sourcePortElement, v.Sport)
			elements = append(elements, sourcePortValue)

			// Add Destination Port
			destPortValue := entities.NewUnsigned16InfoElement(destPortElement, v.Dport)
			elements = append(elements, destPortValue)

			// Add Protocol
			protocolValue := entities.NewUnsigned8InfoElement(protocolElement, protoToNumber(v.Proto)) // TCP protocol
			elements = append(elements, protocolValue)

			// Add Packets
			packetsValue := entities.NewUnsigned64InfoElement(packetsElement, v.Packets)
			elements = append(elements, packetsValue)

			// Add Bytes
			bytesValue := entities.NewUnsigned64InfoElement(bytesElement, v.Bytes)
			elements = append(elements, bytesValue)

			packetsDeltaValue := entities.NewUnsigned64InfoElement(reverseOctetTotalCount, v.DeltaPackets)
			elements = append(elements, packetsDeltaValue)

			DeltabytesElementValue := entities.NewUnsigned64InfoElement(DeltabytesElement, v.DeltaBytes)
			elements = append(elements, DeltabytesElementValue)

			// Add Source IPv4 Address
			sourceIPValue, _ := entities.DecodeAndCreateInfoElementWithValue(sourceIPElement, v.SIP)
			elements = append(elements, sourceIPValue)

			// Add Destination IPv4 Address
			destIPValue, _ := entities.DecodeAndCreateInfoElementWithValue(destIPElement, v.DIP)
			elements = append(elements, destIPValue)

			// 반대로 돌아오는 값..? 어찌 구하지.
			reverseOctetTotalCountValue := entities.NewUnsigned64InfoElement(reverseOctetTotalCount, 1)
			elements = append(elements, reverseOctetTotalCountValue)

			reverseOctetDeltaCountValue := entities.NewUnsigned64InfoElement(reverseOctetDeltaCount, 1)
			elements = append(elements, reverseOctetDeltaCountValue)

			reversePacketTotalCountValue := entities.NewUnsigned64InfoElement(reversePacketTotalCount, 1)
			elements = append(elements, reversePacketTotalCountValue)

			reversePacketDeltaCountValue := entities.NewUnsigned64InfoElement(reversePacketDeltaCount, 1)
			elements = append(elements, reversePacketDeltaCountValue)

			reverseOctetDeltaCountFromSourceNodeValue := entities.NewUnsigned64InfoElement(reverseOctetDeltaCountFromSourceNode, 1)
			elements = append(elements, reverseOctetDeltaCountFromSourceNodeValue)

			reverseOctetTotalCountFromSourceNodeValue := entities.NewUnsigned64InfoElement(reverseOctetTotalCountFromSourceNode, 1)
			elements = append(elements, reverseOctetTotalCountFromSourceNodeValue)

			reversePacketDeltaCountFromSourceNodeValue := entities.NewUnsigned64InfoElement(reversePacketDeltaCountFromSourceNode, 1)
			elements = append(elements, reversePacketDeltaCountFromSourceNodeValue)

			reversePacketTotalCountFromSourceNodeValue := entities.NewUnsigned64InfoElement(reversePacketTotalCountFromSourceNode, 1)
			elements = append(elements, reversePacketTotalCountFromSourceNodeValue)

			tcpStateIe.SetStringValue(stateToTCPState(v.CState))
			//tcpStateValue := entities.NewUnsigned64InfoElement(tcpStateElement, 0)
			elements = append(elements, tcpStateIe)

			dataSet.AddRecord(elements, templateID)
			expProcess.SendSet(dataSet)
			if v.DeleteFlag {
				delete(ctMap.CtMap, v.Key())
			}
		}

		ctMap.Mutex.Unlock()
		time.Sleep(time.Second * time.Duration(SetSecond))
	}
}

func stateToTCPState(CTState string) (TCPState string) {
	switch CTState {
	case "sync-sent":
		TCPState = "SYN-SENT"
	case "sync-ack":
		TCPState = "SYN-RECEIVED"
	case "est":
		TCPState = "ESTABLISHED"
	case "closed-wait":
		TCPState = "CLOSE-WAIT"
	case "closed":
		TCPState = "CLOSED"
	default:
		TCPState = "TIME-WAIT" // default TIME-WAIT,
	}
	return TCPState
}

func protoToNumber(proto string) (protoNum uint8) {
	switch proto {
	case "icmp":
		protoNum = 1
	case "tcp":
		protoNum = 6
	case "udp":
		protoNum = 17
	case "sctp":
		protoNum = 132
	default:
		protoNum = 6 // default TCP
	}
	return protoNum
}
