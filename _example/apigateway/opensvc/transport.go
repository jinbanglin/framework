package opensvc

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jinbanglin/moss/_example/pb"
	"github.com/jinbanglin/moss/auth/moss_jwt"
	"github.com/jinbanglin/moss/kernel/addtransport"
	"github.com/jinbanglin/moss/kernel/payload"
	mosshttp "github.com/jinbanglin/moss/transport/http"

	"github.com/json-iterator/go"
	"github.com/spf13/viper"

	"github.com/jinbanglin/moss/log"

	"github.com/gorilla/mux"
)

func MakeOpensvcHTTPHandler() *mux.Router {
	var r, e = mux.NewRouter(), MakeServerEndpoints(LoggingMiddleware()(NewEntryService()))
	options := []mosshttp.ServerOption{
		mosshttp.ServerErrorLogger(log.Logger{}),
		mosshttp.ServerErrorEncoder(encodeError),
	}
	log.Info("HTTP SERVER |json at", "/api/v1/open/login/{protocol}/"+viper.GetString("etcdv3.server_id"))
	r.Methods("POST").Path("/api/v1/open/login/{protocol}/" + viper.GetString("etcdv3.server_id")).Handler(mosshttp.NewServer(
		e.RegisterEndpoint,
		decodeRegisterRequest,
		encodeRegisterResponse,
		options...,
	))
	return r
}

func decodeRegisterRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	p, ok := mux.Vars(r)["protocol"]
	if !ok {
		return nil, errors.New("no protocol")
	}
	protocol, err := strconv.Atoi(p)
	if err != nil {
		return nil, errors.New("no protocol")
	}
	ctx = context.WithValue(ctx, "client_ip", r.RemoteAddr)
	req := &payload.MossPacket{
		ServiceCode: uint32(protocol),
		Payload:     nil,
		Message:     nil,
		UserId:      "",
		ClientIp:    r.RemoteAddr,
	}
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
	if err := addtransport.GetCodecerByServiceCode(res.ServiceCode).Unmarshal(res.Payload, ret); err != nil {
		return err
	}
	signedKey, err := moss_jwt.NewJwtToken(ret.UserName, ret.UserId, ret.Audience)
	if err != nil {
		return err
	}
	log.Info("encodeRegisterResponse |signedKey", signedKey)
	w.Header().Set("Authorization", "Bearer "+signedKey)
	return jsoniter.NewEncoder(w).Encode(res)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err != nil {
		panic(err)
	}
}
