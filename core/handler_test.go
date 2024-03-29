package core

import (
	pb "github.com/conseweb/icebox/protos"
	. "github.com/smartystreets/goconvey/convey"

	_ "fmt"
	"github.com/conseweb/icebox/core/common"
	"testing"

	"context"
	"github.com/golang/protobuf/proto"
)

//go:generate mockgen -source=handler_test.go -destination=../mocks/mock_IceboxMessage.go -package=mocks github.com/conseweb/icebox/core IceboxMsgIntf

type IceboxMsgIntf interface {
	Reset()
	String() string
	ProtoMessage()
	Descriptor() ([]byte, []int)

	GetVersion() uint32
	GetType() pb.IceboxMessage_Command
	GetSessionId() uint32
	GetPayload() []byte
	GetSignature() []byte
}

func TestHelloWithErrorMsgType(t *testing.T) {
	Convey(`Hello request should received error for invalid msg type'.`, t, func() {

		//req := pb.NewHelloRequest(common.App_magic)
		payload, _ := pb.EncodeHelloRequest(common.App_magic)
		ct := pb.NewIceboxMessage(pb.IceboxMessage_HELLO+1, payload)
		handler := NewIcebergHandler()
		res, err := handler.Execute(context.Background(), ct)
		So(err, ShouldEqual, nil)
		So(res.GetHeader().GetCmd(), ShouldEqual, pb.IceboxMessage_ERROR)
	})
}

func TestHelloWithErrorMagicNumber(t *testing.T) {
	Convey(`Hello request should received error for invalid magic number'.`, t, func() {

		//req := pb.NewHelloRequest(common.App_magic+1)
		payload, _ := pb.EncodeHelloRequest(common.App_magic + 1)
		ct := pb.NewIceboxMessage(pb.IceboxMessage_HELLO, payload)
		handler := NewIcebergHandler()
		res, err := handler.Execute(context.Background(), ct)
		So(err, ShouldEqual, nil)
		So(res.GetHeader().GetCmd(), ShouldEqual, pb.IceboxMessage_ERROR)

		var result = &pb.PbError{}
		err = proto.Unmarshal(res.GetPayload(), result)
		So(result.GetCode(), ShouldEqual, 500)
		So(result.GetMessage(), ShouldEqual, "Unknown app!")
	})
}

func TestHelloSuccess(t *testing.T) {
	Convey(`Hello request should received device's magic number'.`, t, func() {
		//req := pb.NewHelloRequest(common.App_magic)
		payload, _ := pb.EncodeHelloRequest(common.App_magic)
		ct := pb.NewIceboxMessage(pb.IceboxMessage_HELLO, payload)
		handler := NewIcebergHandler()
		res, err := handler.Execute(context.Background(), ct)
		So(err, ShouldEqual, nil)
		So(res.GetHeader().GetCmd(), ShouldNotEqual, pb.IceboxMessage_ERROR)
		So(res.GetHeader().GetCmd(), ShouldEqual, pb.IceboxMessage_HELLO)

		var result = &pb.HelloReply{}
		err = proto.Unmarshal(res.GetPayload(), result)
		So(err, ShouldEqual, nil)
		So(result.GetMagicB(), ShouldEqual, common.Device_magic)
	})

}

func TestNegotiateSuccess(t *testing.T) {
	Convey(`Negotiate request should received a key'.`, t, func() {
		//req := pb.NewNegotiateRequest()
		//payload, _ := pb.EncodeNegotiateRequest(common.App_magic)
		//ct := pb.NewIceboxMessage(pb.IceboxMessage_HELLO, payload)
		//handler := NewIcebergHandler()
		//res, err := handler.Chat(context.Background(), ct)
		//So(err, ShouldEqual, nil)
		//So(res.GetType(), ShouldNotEqual, pb.IceboxMessage_ERROR)
		//So(res.GetType(), ShouldEqual, pb.IceboxMessage_HELLO)
		//
		//var result = &pb.HiReply{}
		//err = proto.Unmarshal(res.GetPayload(), result)
		//So(err, ShouldEqual, nil)
		//So(result.GetMagicB(), ShouldEqual, common.Device_magic)
	})

}
