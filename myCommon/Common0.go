package myCommon
import (
    // "encoding/json"
    // "flag"
    // "fmt"
    // "html/template"
    // "log"
    // "math"
    "net/http"
    // "net/url"
    "os"
    // "strconv"
    // "time"
)
func common_main_port(portNN string)string{
    port := os.Getenv("PORT")
    if port == "" {
        port = portNN
    }
    return port
}
func common_main_static(staticFolder string){
	s = "/" + staticFolder + "/"
	mux := http.NewServeMux()
    mux.Handle(s, http.StripPrefix(s, http.FileServer(http.Dir(staticFolder))))
    return mux
}
func mySum(a, b int64)int64{
	return a+b
}