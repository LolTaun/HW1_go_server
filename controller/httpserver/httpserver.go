package httpserver

import (
	"HW1_http/gates/psg"
	"HW1_http/models/dto"
	"encoding/json"
	"fmt"
	"io"
	"log"

	//"errors"
	"net/http"

	"github.com/pkg/errors"
)

type HttpServer struct {
	srv http.Server
	db  *psg.Psg
}

func NewHttpServer(addr string) (hs *HttpServer, err error) {
	hs = new(HttpServer)
	hs.srv = http.Server{}
	mux := http.NewServeMux()
	mux.Handle("/create", http.HandlerFunc(hs.recordCreateHandler))
	mux.Handle("/get", http.HandlerFunc(hs.recordsGetHandler))
	mux.Handle("/update", http.HandlerFunc(hs.RecordUpdateHandler))
	mux.Handle("/delete", http.HandlerFunc(hs.RecordDeleteByPhone))
	hs.srv.Handler = mux
	hs.srv.Addr = addr

	//\\ Postges default connection
	fmt.Println("Enter DB URL (leave blank for default):")
	dburl := ""
	fmt.Scanln(&dburl)
	if dburl == "" {
		dburl = "postgres://127.0.0.1:5432/web-programming"
	}

	fmt.Println("Enter DB login (leave blank for default):")
	login := ""
	fmt.Scanln(&login)
	if login == "" {
		login = "postgres"
	}

	fmt.Println("Enter DB password:")
	pass := ""
	fmt.Scanln(&pass)
	if pass == "" {
		err = errors.New("Password is empty")
		errors.Wrap(err, "NewHttpServer(): psg.NewPsg()")
		return
	}
	p, err := psg.NewPsg(dburl, login, pass)
	if err != nil {
		errors.Wrap(err, "NewHttpServer(): psg.NewPsg()")
		return
	}
	hs.db = p
	//\\

	return hs, err
}

func (hs *HttpServer) Start() (err error) {
	err = hs.srv.ListenAndServe()
	if err != nil {
		err = errors.Wrap(err, "hs.srv.ListenAndServe()")
		return err
	}
	return err
}

func (hs *HttpServer) recordCreateHandler(w http.ResponseWriter, req *http.Request) {
	record := dto.Record{}
	byteReq, err := io.ReadAll(req.Body)
	if err != nil {
		err = errors.Wrap(err, "httpserver: (hs *HttpServer) recordCreateHandler(): io.ReadAll(req.Body)")
		log.Println(err)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	err = json.Unmarshal(byteReq, &record)
	if err != nil {
		err = errors.Wrap(err, "httpserver: (hs *HttpServer) recordCreateHandler(): json.Marshal(req)")
		log.Println(err)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	if err = dto.RecordValidation(record); err != nil {
		err = errors.Wrap(err, "httpserver: (hs *HttpServer) recordCreateHandler(): dto.RecordValidation(record)")
		log.Println(err)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	err = hs.db.RecordSave(record)

	if err != nil {
		err = errors.Wrap(err, "httpserver: (hs *HttpServer) recordCreateHandler(): hs.db.RecordSave(record)")
		log.Println(err)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (hs *HttpServer) recordsGetHandler(w http.ResponseWriter, req *http.Request) {
	record := dto.Record{}
	byteReq, err := io.ReadAll(req.Body)
	if err != nil {
		err = errors.Wrap(err, "httpserver: (hs *HttpServer) recordGetHandler(): io.ReadAll(req.Body)")
		log.Println(err)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	err = json.Unmarshal(byteReq, &record)
	if err != nil {
		err = errors.Wrap(err, "httpserver: (hs *HttpServer) recordGetHandler(): json.Marshal(req)")
		log.Println(err)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	records, err := hs.db.RecordsGet(record)
	if err != nil {
		err = errors.Wrap(err, "httpserver: (hs *HttpServer) recordGetHandler(): hs.db.RecordsGet(record)")
		log.Println(err)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	err = json.NewEncoder(w).Encode(records)
	if err != nil {
		err = errors.Wrap(err, "httpserver: (hs *HttpServer) recordGetHandler(): json.NewEncoder(w).Encode(records)")
		log.Println(err)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (hs *HttpServer) RecordUpdateHandler(w http.ResponseWriter, req *http.Request) {
	record := dto.Record{}
	byteReq, err := io.ReadAll(req.Body)
	if err != nil {
		err = errors.Wrap(err, "httpserver: (hs *HttpServer) recordUpdateHandler(): io.ReadAll(req.Body)")
		log.Println(err)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	err = json.Unmarshal(byteReq, &record)
	if err != nil {
		err = errors.Wrap(err, "httpserver: (hs *HttpServer) recordUpdateHandler(): json.Unmarshal(req)")
		log.Println(err)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	if err = dto.RecordValidation(record); err != nil {
		err = errors.Wrap(err, "httpserver: (hs *HttpServer) recordCreateHandler(): dto.RecordValidation(record)")
		log.Println(err)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	err = hs.db.RecordUpdate(record)
	if err != nil {
		err = errors.Wrap(err, "httpserver: (hs *HttpServer) recordUpdateHandler(): hs.db.RecordUpdate(record)")
		log.Println(err)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (hs *HttpServer) RecordDeleteByPhone(w http.ResponseWriter, r *http.Request) {
	record := dto.Record{}
	byteReq, err := io.ReadAll(r.Body)
	if err != nil {
		err = errors.Wrap(err, "httpserver: (hs *HttpServer) recordDeleteByPhone(): io.ReadAll(req.Body)")
		log.Println(err)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	err = json.Unmarshal(byteReq, &record)
	if err != nil {
		err = errors.Wrap(err, "httpserver: (hs *HttpServer) recordDeleteByPhone(): json.Unmarshal(req)")
		log.Println(err)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	err = hs.db.RecordDeleteByPhone(record.Phone)
	if err != nil {
		err = errors.Wrap(err, "httpserver: (hs *HttpServer) recordDeleteByPhone(): hs.db.RecordDeleteByPhone(record.Phone)")
		log.Println(err)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	w.WriteHeader(http.StatusOK)
}
