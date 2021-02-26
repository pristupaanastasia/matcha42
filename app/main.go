package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

/*type indexHandler struct {
	mu sync.Mutex // guards n
	n  int
}
*/
func loginHandler(next http.Handler) http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}
func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
/*func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "Welcome!")
}*/

func registrHandler(w http.ResponseWriter, r *http.Request){
	c := http.Cookie{
		Name:   "ithinkidroppedacookie",
		Value:  "thedroppedcookiehasgoldinit",
		MaxAge: 3600}
	http.SetCookie(w, &c)

	w.Write([]byte("new cookie created!\n"))
}

func indexHandler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Welcome!")
}
func main() {

	connStr := "host=db port=5432 user=postgres password=1805 dbname=db_matcha sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	http.Handle("/", loginHandler(recoverHandler(http.HandlerFunc(indexHandler))))
	fmt.Println("Server is listening...")
	http.ListenAndServe(":9000", nil)
}