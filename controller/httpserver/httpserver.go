package httpserver

import (
	"HW1_http/gates/psg"
	"HW1_http/models/dto"
	"encoding/json"
	"errors"
	"log"

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
	
	if err != nil {
		err = eW.WrapError(err, "pkg.NewEWrapper()")
		return
	}

	err = hs.srv.ListenAndServe()
	if err != nil {
		err = eW.WrapError(err, "hs.srv.ListenAndServe()")
		return
	}
	return
}

func (hs *HttpServer) recordCreateHandler(w http.ResponseWriter, req *http.Request) {
	eW, err := pkg.NewEWrapperWithFile("(hs *HttpServer) recordCreateHandler()")
	if err != nil {
		log.Println("(hs *HttpServer) recordCreateHandler: NewEWrapperWithFile()", err)
	}

	resp := &dto.Response{}
	defer responseReturn(w, eW, resp)

	if req.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

	record := dto.Record{}
	byteReq, err := io.ReadAll(req.Body)
	if err != nil {
		resp.Wrap("Error reading request", nil, err.Error())
		eW.LogError(err, "io.ReadAll(req.Body)")
		return
	}
	err = json.Unmarshal(byteReq, &record)
	if err != nil {
		resp.Wrap("Error JSON", nil, err.Error())
		eW.LogError(err, "json.Unmarshal(req)")
		return
	}

	if record.Name == "" || record.LastName == "" || record.Address == "" || record.Phone == "" {
		err = errors.New("required data is missing")
		resp.Wrap("Required data is missing", nil, err.Error())
		eW.LogError(err, "json.Unmarshal")
		return
	}

	record.Phone, err = pkg.PhoneNormalize(record.Phone)
	if err != nil {
		resp.Wrap("Error: wrong Phone", nil, err.Error())
		eW.LogError(err, "pkg.PhoneNormalize(record.Phone)")
		return
	}

	err = hs.db.RecordSave(record)

	if err != nil {
		resp.Wrap("Error in saving record", nil, err.Error())
		eW.LogError(err, "hs.db.RecordSave(record)")
		return
	}

	resp.Wrap("Successfully added", nil, "")
}

func (hs *HttpServer) recordsGetHandler(w http.ResponseWriter, req *http.Request) {
	eW, err := pkg.NewEWrapperWithFile("(hs *HttpServer) recordsGetHandler()")
	if err != nil {
		log.Println("(hs *HttpServer) recordCreateHandler: NewEWrapperWithFile()", err)
	}
	resp := &dto.Response{}
	defer responseReturn(w, eW, resp)

	if req.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

	record := dto.Record{}
	byteReq, err := io.ReadAll(req.Body)
	if err != nil {
		resp.Wrap("Error reading request", nil, err.Error())
		eW.LogError(err, "io.ReadAll(req.Body)")
		return
	}
	err = json.Unmarshal(byteReq, &record)
	if err != nil {
		resp.Wrap("Error JSON", nil, err.Error())
		eW.LogError(err, "json.Marshal(req)")
		return
	}

	if record.Phone != ""{
		record.Phone, err = pkg.PhoneNormalize(record.Phone)
		if err != nil {
			resp.Wrap("Error: wrong Phone", nil, err.Error())
			eW.LogError(err, "pkg.PhoneNormalize(record.Phone)")
			return
		}
	}

	records, err := hs.db.RecordsGet(record)
	if err != nil {
		resp.Wrap("Error in finding records", nil, err.Error())
		eW.LogError(err, "hs.db.RecordsGet(record)")
		return
	}

	recordsJSON, err := json.Marshal(records)
	if err != nil {
		resp.Wrap("Error JSON", nil, err.Error())
		eW.LogError(err, "json.Marshal(records)")
		return
	}

	resp.Wrap("Success", recordsJSON, "")
}

func (hs *HttpServer) recordUpdateHandler(w http.ResponseWriter, req *http.Request) {
	eW, err := pkg.NewEWrapperWithFile("(hs *HttpServer) recordUpdateHandler()")
	if err != nil {
		log.Println("(hs *HttpServer) recordCreateHandler: NewEWrapperWithFile()", err)
	}

	resp := &dto.Response{}
	defer responseReturn(w, eW, resp)

	if req.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

	record := dto.Record{}
	byteReq, err := io.ReadAll(req.Body)
	if err != nil {
		resp.Wrap("Error reading request", nil, err.Error())
		eW.LogError(err, "io.ReadAll(req.Body)")
		return
	}
	err = json.Unmarshal(byteReq, &record)
	if err != nil {
		resp.Wrap("Error JSON", nil, err.Error())
		eW.LogError(err, "json.Unmarshal(byteReq, &record)")
		return
	}

	if (record.Name == "" && record.LastName == "" && record.MiddleName == "" && record.Address == "") || record.Phone == "" {
		err = errors.New("required data is missing")
		resp.Wrap("Required data is missing", nil, err.Error())
		eW.LogError(err, "json.Unmarshal")
		return
	}

	record.Phone, err = pkg.PhoneNormalize(record.Phone)
	if err != nil {
		resp.Wrap("Error: wrong Phone", nil, err.Error())
		eW.LogError(err, "pkg.PhoneNormalize(record.Phone)")
		return
	}

	err = hs.db.RecordUpdate(record)
	if err != nil {
		resp.Wrap("Error in updating record", nil, err.Error())
		eW.LogError(err, "hs.db.RecordUpdate(record)")
		return
	}
	resp.Wrap("Success", nil, "")
}

func (hs *HttpServer) recordDeleteByPhone(w http.ResponseWriter, req *http.Request) {
	eW, err := pkg.NewEWrapperWithFile("(hs *HttpServer) recordDeleteByPhone()")
	if err != nil {
		log.Println("(hs *HttpServer) recordCreateHandler: NewEWrapperWithFile()", err)
	}

	resp := &dto.Response{}
	defer responseReturn(w, eW, resp)

	if req.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

	record := dto.Record{}
	byteReq, err := io.ReadAll(req.Body)
	if err != nil {
		resp.Wrap("Error reading request", nil, err.Error())
		eW.LogError(err, "io.ReadAll(r.Body)")
		return
	}
	err = json.Unmarshal(byteReq, &record)
	if err != nil {
		resp.Wrap("Error JSON", nil, err.Error())
		eW.LogError(err, "json.Unmarshal(byteReq, &record)")
		return
	}

	if record.Phone == "" {
		err = errors.New("phone data is missing")
		resp.Wrap("Phone data is missing", nil, err.Error())
		eW.LogError(err, "json.Unmarshal")
		return
	}

	record.Phone, err = pkg.PhoneNormalize(record.Phone)
	if err != nil {
		resp.Wrap("Error: wrong Phone", nil, err.Error())
		eW.LogError(err, "pkg.PhoneNormalize(record.Phone)")
		return
	}

	err = hs.db.RecordDeleteByPhone(record.Phone)
	if err != nil {
		resp.Wrap("Error in deleting record", nil, err.Error())
		eW.LogError(err, "hs.db.RecordDeleteByPhone(record.Phone)")
		return
	}
	resp.Wrap("Success", nil, "")
}

func responseReturn(w http.ResponseWriter, eW *pkg.EWrapper, resp *dto.Response){
	err_encode := json.NewEncoder(w).Encode(resp)
	if err_encode != nil {
		eW.LogError(err_encode, "json.NewEncoder(w).Encode(resp)")
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	eW.Close()
}