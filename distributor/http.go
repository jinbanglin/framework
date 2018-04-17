package distributor

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jinbanglin/moss/endpoint"
	"github.com/jinbanglin/moss/log"
	mosshttp "github.com/jinbanglin/moss/transport/http"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	"github.com/jinbanglin/moss/auth/moss_jwt"
	"github.com/jinbanglin/moss/payload"
	"github.com/json-iterator/go"
)

type MutilEndpoints struct {
	Endpoints map[string]endpoint.Endpoint
}

func MakeHTTPGateway(r *mux.Router, endpoints MutilEndpoints) http.Handler {
	for k, v := range endpoints.Endpoints {
		log.Infof("MOSS |gateway route at=%s", k)
		r.Methods("POST").Path(k).Handler(mosshttp.NewServer(
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
		MossMessage: payload.StatusText(payload.StatusUnauthorized),
	}
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(*jwtgo.Token) (interface{}, error) {
			return moss_jwt.JwtKey, nil
		})
	if err != nil || !token.Valid {
		log.Errorf("MOSS |token=%v", token)
		ctx = context.WithValue(ctx, payload.StatusUnauthorized, true)
		return response, err
	}
	vars := mux.Vars(r)
	c, ok := vars["service_code"]
	if !ok {
		return response, errors.New("no service_code")
	}
	serviceCode, err := strconv.Atoi(c)
	if err != nil {
		log.Error("MOSS |err=", err)
		return response, err
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("MOSS |err=%v ", err)
		return response, err
	}
	response.UserId = token.Claims.(jwtgo.MapClaims)["UserId"].(string)
	response.ServiceCode = uint32(serviceCode)
	response.Payload = b
	response.ClientIp = r.RemoteAddr
	response.MossMessage = payload.StatusText(payload.StatusOK)
	return response, nil
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Write(response.(*payload.MossPacket).Payload)
	return nil
}
