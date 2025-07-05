package weberr;

import(
	Fmt     "fmt"
	HTTP    "net/http"
	JSON    "encoding/json"
	Runtime "runtime"
);



type WebError struct {
	Status  int
	Err     error
	Trace   string
	Is_JSON bool
}



func IntServerErr(err error) *WebError {
	trace := make([]byte, 1024);
	n := Runtime.Stack(trace, true);
	return &WebError{
		Status:  HTTP.StatusInternalServerError,
		Err:     err,
		Trace:   string(trace[:n]),
		Is_JSON: false,
	};
}



func (weberr *WebError) Write(out HTTP.ResponseWriter) {
	if weberr.Is_JSON {
		json, err := JSON.Marshal(
			struct{
				Error string
			}{
				Error: weberr.Err.Error(),
			});
		if err != nil { panic(err); }
		HTTP.Error(out, string(json), weberr.Status);
	} else {
		HTTP.Error(out,
			Fmt.Sprintf("%s\n%s", weberr.Err.Error(), weberr.Trace),
			weberr.Status,
		);
	}
}



func (weberr *WebError) IsJSON() *WebError {
	weberr.Is_JSON = true;
	return weberr;
}
