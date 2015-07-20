package pegasus

import (
	"github.com/HearthSim/hs-proto/go"
	"github.com/golang/protobuf/proto"
)

// EncodeUtilResponse builds a buffer encoding the response packetId and body.
func EncodeUtilResponse(packetId int32, body []byte) ([]byte, error) {
	res := hsproto.BnetProtocolGameUtilities_ClientResponse{}
	res.Attribute = make([]*hsproto.BnetProtocolAttribute_Attribute, 2)
	res.Attribute[0] = &hsproto.BnetProtocolAttribute_Attribute{
		Name: proto.String("t"),
		Value: &hsproto.BnetProtocolAttribute_Variant{
			IntValue: proto.Int64(int64(packetId)),
		},
	}
	res.Attribute[1] = &hsproto.BnetProtocolAttribute_Attribute{
		Name: proto.String("p"),
		Value: &hsproto.BnetProtocolAttribute_Variant{
			BlobValue: body,
		},
	}
	return proto.Marshal(&res)
}

func encodeUtilPacketId(systemId, packetId int) int {
	return systemId<<16 | packetId
}