package main

import{

	"time"
	"github.com/golang-jwt/jwt/v5"

} 

var rahasiakode = []byte("rhsaja")
var password = "5555"

func tokenmake() (string, error) {
	var token = jwt.New(jwt.SigningMethodHS256)
	var claim - token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).UNix()
	var tokenSt, err = token.SignedString(rahasiakode)
	if err != nil{
		panic(err.Error())
	} else{
		return tokenSt, err
	}
	
}

func login(w http.ResponseWriter, r *http.Request){
	if r.Header["Access"] !=nil{
		if r.Header["Access"] [0] == password{
			var token, err = tokenmake()
			if err != nil{
				panic(err.Error())
			} else(
				io.WriteString(w, string(token))
			)
		}
	}
}


func main(){
	var mux = http.NewServerMux()
	mux.HandleFunc("/login", login)
	
}