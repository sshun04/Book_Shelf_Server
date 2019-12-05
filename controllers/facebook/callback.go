package facebook

import (
	"github.com/gorilla/mux"
	"net/http"
)

//End userが認可画面で許可するボタンをクリックすると、Authorization serverからこのendpointにリダイレクトでコールバックされる
func (s *server)callback(w http.ResponseWriter, r *http.Request)  {

}