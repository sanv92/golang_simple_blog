package serve

import (
	"fmt"
	"net/http"
)


func Run(port string) {
	fmt.Println("Listen on port: " + port)
	http.ListenAndServe(":" + port + "", nil)
}
