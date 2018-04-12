package distributor

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jinbanglin/moss/log"
	httptransport "github.com/jinbanglin/moss/transport/http"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	"github.com/jinbanglin/moss"
	"github.com/jinbanglin/moss/auth/moss_jwt"
	"github.com/jinbanglin/moss/payload"
	"github.com/json-iterator/go"
)

type MutilEndpoints struct {
	Endpoints map[string]moss.Endpoint
}

func MakeHTTPGateway(r *mux.Router, endpoints MutilEndpoints, serviceId string) http.Handler {
	for k, v := range endpoints.Endpoints {
		log.Info("run at", k+serviceId)
		r.Methods("POST").Path(k + serviceId).Handler(httptransport.NewServer(
			v,
			decodeHTTPInvokeRequest,
			encodeHTTPGenericResponse,
			errorEncoder,
		))
	}
	return r
}

func errorEncoder(ctx context.Context, response interface{}, w http.ResponseWriter) {
	jsoniter.NewEncoder(w).Encode(response)
}

func decodeHTTPInvokeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var response = &payload.MossPacket{
		Message: &payload.Message{
			Code: 40001,
			Msg:  "status unauthorized",
		},
	}
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(*jwtgo.Token) (interface{}, error) {
			return moss_jwt.JwtKey, nil
		})
	if err != nil || !token.Valid {
		log.Errorf("token=%v", token)
		ctx = context.WithValue(ctx, http.StatusUnauthorized, true)
		return response, err
	}
	vars := mux.Vars(r)
	c, ok := vars["service_code"]
	if !ok {
		return response, errors.New("no protocol")
	}
	serviceCode, err := strconv.Atoi(c)
	if err != nil {
		return response, err
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil || len(b) < 1 {
		log.Errorf("err=%v |or EOF:%d", err, len(b))
		return response, errors.New("client data error")
	}
	response.UserId = token.Claims.(jwtgo.MapClaims)["UserId"].(string)
	response.ServiceCode = uint32(serviceCode)
	response.Payload = b
	response.ClientIp = r.RemoteAddr
	response.Message = &payload.Message{
		Code: 20000,
		Msg:  "SUCCESS",
	}
	return response, nil
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Write(response.(*payload.MossPacket).Payload)
	return nil
}
