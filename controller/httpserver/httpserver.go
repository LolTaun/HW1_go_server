package httpserver

import (
	"HW1_http/gates/psg"
	"HW1_http/models/dto"
	"encoding/json"
	"errors"

	"HW1_http/pkg"
	"io"
	"net/http"
)

type HttpServer struct {
	srv http.Server
	db  *psg.Psg
}

func NewHttpServer(addr string, p *psg.Psg) (hs *HttpServer) {
	hs = new(HttpServer)
	hs.srv = http.Server{}
	mux := http.NewServeMux()
	mux.Handle("/create", http.HandlerFunc(hs.recordCreateHandler))
	mux.Handle("/get", http.HandlerFunc(hs.recordsGetHandler))
	mux.Handle("/update", http.HandlerFunc(hs.recordUpdateHandler))
	mux.Handle("/delete", http.HandlerFunc(hs.recordDeleteByPhone))
	hs.srv.Handler = mux
	hs.srv.Addr = addr
	hs.db = p
	return hs
}

func (hs *HttpServer) Start() (err error) {
	eW := pkg.NewEWrapper("(hs *HttpServer) Start()")

	err = hs.srv.ListenAndServe()
	if err != nil {
		err = eW.WrapError(err, "hs.srv.ListenAndServe()")
		return
	}
	return
}

func (hs *HttpServer) recordCreateHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	eW := pkg.NewEWrapper("(hs *HttpServer) recordCreateHandler()")

	record := dto.Record{}
	byteReq, err := io.ReadAll(req.Body)
	if err != nil {
		eW.LogError(err, "io.ReadAll(req.Body)")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	err = json.Unmarshal(byteReq, &record)
	if err != nil {
		eW.LogError(err, "json.Unmarshal(req)")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	if record.Name == "" || record.LastName == "" || record.Address == "" || record.Phone == "" {
		err = errors.New("required data is missing")
		eW.LogError(err, "json.Unmarshal")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	record.Phone, err = pkg.PhoneNormalize(record.Phone)
	if err != nil {
		eW.LogError(err, "pkg.PhoneNormalize(record.Phone)")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	err = hs.db.RecordSave(record)

	if err != nil {
		eW.LogError(err, "hs.db.RecordSave(record)")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (hs *HttpServer) recordsGetHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	eW := pkg.NewEWrapper("(hs *HttpServer) recordsGetHandler()")

	record := dto.Record{}
	byteReq, err := io.ReadAll(req.Body)
	if err != nil {
		eW.LogError(err, "io.ReadAll(req.Body)")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	err = json.Unmarshal(byteReq, &record)
	if err != nil {
		eW.LogError(err, "json.Marshal(req)")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	records, err := hs.db.RecordsGet(record)
	if err != nil {
		eW.LogError(err, "hs.db.RecordsGet(record)")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	err = json.NewEncoder(w).Encode(records)
	if err != nil {
		eW.LogError(err, "json.NewEncoder(w).Encode(records)")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (hs *HttpServer) recordUpdateHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	eW := pkg.NewEWrapper("(hs *HttpServer) recordUpdateHandler()")

	record := dto.Record{}
	byteReq, err := io.ReadAll(req.Body)
	if err != nil {
		eW.LogError(err, "io.ReadAll(req.Body)")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	err = json.Unmarshal(byteReq, &record)
	if err != nil {
		eW.LogError(err, "json.Unmarshal(byteReq, &record)")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}


	if (record.Name == "" && record.LastName == "" && record.MiddleName == "" && record.Address == "") || record.Phone == "" {
		err = errors.New("required data is missing")
		eW.LogError(err, "json.Unmarshal")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	err = hs.db.RecordUpdate(record)
	if err != nil {
		eW.LogError(err, "hs.db.RecordUpdate(record)")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (hs *HttpServer) recordDeleteByPhone(w http.ResponseWriter, r *http.Request) {
	var err error
	eW := pkg.NewEWrapper("(hs *HttpServer) recordDeleteByPhone()")

	record := dto.Record{}
	byteReq, err := io.ReadAll(r.Body)
	if err != nil {
		eW.LogError(err, "io.ReadAll(r.Body)")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	err = json.Unmarshal(byteReq, &record)
	if err != nil {
		eW.LogError(err, "json.Unmarshal(byteReq, &record)")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	if record.Phone == "" {
		err = errors.New("phone data is missing")
		eW.LogError(err, "json.Unmarshal")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	err = hs.db.RecordDeleteByPhone(record.Phone)
	if err != nil {
		eW.LogError(err, "hs.db.RecordDeleteByPhone(record.Phone)")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	w.WriteHeader(http.StatusOK)
}
