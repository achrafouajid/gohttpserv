package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API Service ...")

	fmt.Println(
		`
		      ___           ___                    ___           ___               
     /\__\         /\  \                  /\  \         /\  \              
    /:/ _/_       /::\  \                /::\  \       /::\  \     ___     
   /:/ /\  \     /:/\:\  \              /:/\:\  \     /:/\:\__\   /\__\    
  /:/ /::\  \   /:/  \:\  \            /:/ /::\  \   /:/ /:/  /  /:/__/    
 /:/__\/\:\__\ /:/__/ \:\__\          /:/_/:/\:\__\ /:/_/:/  /  /::\  \    
 \:\  \ /:/  / \:\  \ /:/  /          \:\/:/  \/__/ \:\/:/  /   \/\:\  \__ 
  \:\  /:/  /   \:\  /:/  /            \::/__/       \::/__/     ~~\:\/\__\
   \:\/:/  /     \:\/:/  /              \:\  \        \:\  \        \::/  /
    \::/  /       \::/  /                \:\__\        \:\__\       /:/  / 
     \/__/         \/__/                  \/__/         \/__/       \/__/  
		
	 from: https://asciified.thelicato.io/`)
err := http.ListenAndServe("localhost:8000",r)

if err != nil {
	log.Error(err)
}

}