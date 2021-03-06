package metrics

import (
	"github.com/singnet/snet-daemon/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
)

//Request stats that will be captured
type RequestStats struct {
	Type                       string `json:"type"`
	RegistryAddressKey         string `json:"registry_address_key"`
	EthereumJsonRpcEndpointKey string `json:"ethereum_json_rpc_endpoint"`
	RequestID                  string `json:"request_id"`
	InputDataSize              string `json:"input_data_size"`
	ServiceMethod              string `json:"service_method"`
	RequestReceivedTime        string `json:"request_received_time"`
	OrganizationID             string `json:"organization_id"`
	ServiceID                  string `json:"service_id"`
	GroupID                    string `json:"group_id"`
	DaemonEndPoint             string `json:"daemon_end_point"`
}

//Create a request Object and Publish this to a service end point
func PublishRequestStats(commonStat *CommonStats, inStream grpc.ServerStream) bool {
	request := createRequestStat(commonStat)
	if md, ok := metadata.FromIncomingContext(inStream.Context()); ok {
		request.setDataFromContext(md)
	}
	return Publish(request, config.GetString(config.MonitoringServiceEndpoint)+"/event")
}

func (request *RequestStats) setDataFromContext(md metadata.MD) {
	request.InputDataSize = strconv.FormatUint(GetSize(md), 10)

}

func createRequestStat(commonStat *CommonStats) *RequestStats {
	request := &RequestStats{
		Type:                       "request",
		RegistryAddressKey:         config.GetString(config.RegistryAddressKey),
		EthereumJsonRpcEndpointKey: config.GetString(config.EthereumJsonRpcEndpointKey),
		RequestID:                  commonStat.ID,
		GroupID:                    commonStat.GroupID,
		DaemonEndPoint:             commonStat.DaemonEndPoint,
		OrganizationID:             commonStat.OrganizationID,
		ServiceID:                  commonStat.ServiceID,
		RequestReceivedTime:        commonStat.RequestReceivedTime,
		ServiceMethod:              commonStat.ServiceMethod,
	}
	return request
}
