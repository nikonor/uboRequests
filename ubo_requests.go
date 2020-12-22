package uboRequests

import (
	"errors"
	"net/http"
	"time"

	"github.com/nikonor/uroboros"
)

var CantDoRequest = errors.New("can't do request now")

type UboReq struct {
	u   *uroboros.Uroboros
	cli http.Client
}

// NewRule - создаем новое "окно" для запросов
//      cli - http-клиент для запроса
//      period - время окна
//      maxRequests - максимальное кол-во запросов за время
func NewRule(cli http.Client, period time.Duration, maxRequests int) *UboReq {
	return &UboReq{
		cli: cli,
		u:   uroboros.New(maxRequests, period),
	}
}

// ReadyForRequest - функция проверки возможности сделать запрос
func (u UboReq) ReadyForRequest() bool {
	return u.u.Can(time.Now())
}

// DoRequest - выполняем запрос, если это возможно
func (u UboReq) DoRequest(req *http.Request) (*http.Response, error) {
	if !u.ReadyForRequest() {
		return nil, CantDoRequest
	}

	return u.cli.Do(req)
}
