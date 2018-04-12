package opensvc

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jinbanglin/moss/_example/pb"
	"github.com/jinbanglin/moss/auth/moss_jwt"
	mosshttp "github.com/jinbanglin/moss/transport/http"

	"github.com/json-iterator/go"
	"github.com/spf13/viper"

	"github.com/jinbanglin/moss/log"

	"github.com/gorilla/mux"
	"github.com/jinbanglin/moss"
	"github.com/jinbanglin/moss/payload"
)

func MakeOpensvcHTTPHandler() *mux.Router {
	var r, e = mux.NewRouter(), MakeServerEndpoints()
	log.Info("HTTP SERVER |json at", "/api/v1/open/sns/{service_code}/"+viper.GetString("etcdv3.server_id"))
	r.Methods("POST").Path("/api/v1/open/sns/{service_code}/" + viper.GetString("etcdv3.server_id")).Handler(mosshttp.NewServer(
		e.SnsEndpoint,
		decodeRegisterRequest,
		encodeRegisterResponse,
		encodeError,
	))
	return r
}

func decodeRegisterRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	req := &payload.MossPacket{
		ServiceCode: 0,
		Payload:     nil,
		Message: &payload.Message{
			Code: 40000,
			Msg:  "request data error",
		},
		UserId:   "",
		ClientIp: r.RemoteAddr,
	}
	p, ok := mux.Vars(r)["service_code"]
	if !ok {
		return req, errors.New("no service_code")
	}
	serviceCode, err := strconv.Atoi(p)
	if err != nil {
		return req, errors.New("no service code")
	}
	req.ServiceCode = uint32(serviceCode)
	if b, err := ioutil.ReadAll(r.Body); err != nil || len(b) < 1 {
		return req, errors.New("client data error")
	} else {
		req.Payload = b
	}
	return req, nil
}

func encodeRegisterResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*payload.MossPacket)
	ret := &pb.RegisterRes{}
	if err := moss.GetCodec(res.ServiceCode).Unmarshal(res.Payload, ret); err != nil {
		return err
	}
	signedKey, err := moss_jwt.NewJwtToken(ret.UserName, ret.UserId, ret.Audience)
	if err != nil {
		return err
	}
	log.Info("encodeRegisterResponse |signedKey", signedKey)
	w.Header().Set("Authorization", "Bearer "+signedKey)
	return jsoniter.NewEncoder(w).Encode(ret)
}

func encodeError(_ context.Context, response interface{}, w http.ResponseWriter) {
	jsoniter.NewEncoder(w).Encode(response)
}
