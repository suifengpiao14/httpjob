package httpjob

import (
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

type HttpMessage struct {
	MessageID string `json:"messageId"`
	Method    string `json:"method"`
	URL       string `json:"url"`
	Payload   string `json:"payload"`
}

func Run(msg HttpMessage) (err error) {
	r := resty.New().R()
	method := strings.ToUpper(msg.Method)
	switch method {
	case http.MethodGet:
		r = r.SetQueryString(msg.Payload)
	case http.MethodPost:
		r = r.SetBody(msg.Payload)
	default:
		err = errors.Errorf("not suport mehtod:%s", msg.Method)
		return err
	}
	response, err := r.Execute(method, msg.URL)
	if err != nil {
		return err
	}
	body := string(response.Body())
	httpCode := response.StatusCode()
	if httpCode != http.StatusOK {
		err := errors.Errorf("response error:httpCode:%d,body:%s", httpCode, body)
		return err
	}
	//todo 验证内容是否正常
	return nil
}
