package controller

import "net/http"

func Home() http.Handler { return http.FileServer(http.Dir("./static")) }
