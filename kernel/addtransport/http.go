package addtransport

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jinbanglin/moss/kernel/addendpoint"
	"github.com/jinbanglin/moss/kernel/payload"
	"github.com/jinbanglin/moss/log"
	"github.com/jinbanglin/moss/tracing"
	httptransport "github.com/jinbanglin/moss/transport/http"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	"github.com/jinbanglin/moss/auth/moss_jwt"
	"github.com/json-iterator/go"
	"github.com/opentracing/opentracing-go"
)

type MutilEndpoints struct {
	Endpoints map[string]addendpoint.Set
}

func MakeMutilHTTPHandler(r *mux.Router, endpoints MutilEndpoints, tracer opentracing.Tracer, serviceId string) http.Handler {
	if r == nil {
		r = mux.NewRouter()
	}
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorLogger(log.Logger{}),
	}
	for k, v := range endpoints.Endpoints {
		log.Info("HTTP SERVER |run at", k+serviceId)
		r.Methods("POST").Path(k + serviceId).Handler(httptransport.NewServer(
			v.InvokeEndpoint,
			decodeHTTPInvokeRequest,
			encodeHTTPGenericResponse,
			append(options, httptransport.ServerBefore(tracing.HTTPToContext(tracer, "handler")))...,
		))
	}
	return r
}

func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	if err != nil {
		if v := ctx.Value(http.StatusUnauthorized); v != nil {
			if v.(bool) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
		w.WriteHeader(http.StatusInternalServerError)
		jsoniter.NewEncoder(w).Encode(&payload.MossPacket{
			Message: &payload.Message{
				Code: 50001,
				Msg:  "server internal error ",
			},
		})
	}
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
		log.Errorf("decodeHTTPInvokeRequest |token=%v", token)
		ctx = context.WithValue(ctx, http.StatusUnauthorized, true)
		return response, err
	}
	vars := mux.Vars(r)
	c, ok := vars["protocol"]
	if !ok {
		return response, errors.New("no protocol")
	}
	serviceCode, err := strconv.Atoi(c)
	if err != nil {
		return response, err
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil || len(b) < 1 {
		log.Errorf("decodeHTTPInvokeRequest |err=%v |or EOF:%d", err, len(b))
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
