package main
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)


func main() {


	const (
		PORT  = ":3000"
	)

	//logger := log.New(os.Stdout,"",1)

	r := httprouter.New()

	r.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Welcome!\n")
	})




	err := http.ListenAndServe(PORT, r)
	if err != nil {

	}


}
