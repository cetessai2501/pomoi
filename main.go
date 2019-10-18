package main

import (
  "log"
  "encoding/json"
  "net/http"
  "os"
  "fmt"  
  "github.com/gorilla/mux"
)


type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Employees []Employee


var employees []Employee



type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"getEmployees",
		"GET",
		"/employees",
		getEmployees,
	},
	Route{
		"addEmployee",
		"POST",
		"/employee/add",
		addEmployee,
	},
        
}

func init() {
        
	employees = Employees{
		Employee{Id: "1", FirstName: "Foo", LastName: "Bar"},
		Employee{Id: "2", FirstName: "Baz", LastName: "Qux"},
	}
        
}


func getEmployees(w http.ResponseWriter, r *http.Request) {
    
    w.Header().Set("Content-Type", "application/json") 
    //log.Printf("eeee", employees);
    //log.WithFields(log.Fields{"employees": employees, }).Info("message");  
    json.NewEncoder(w).Encode(employees)
	
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Print("error occurred while decoding employee data :: ", err)
		return
	}
	
	employees = append(employees, Employee{Id: employee.Id, FirstName: employee.FirstName, LastName: employee.LastName})
	json.NewEncoder(w).Encode(employees)
}


func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}


func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}



func main() {
  addr , err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }  
  muxRouter := mux.NewRouter().StrictSlash(true)
  
  router := AddRoutes(muxRouter)
  router.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))
  
  log.Printf("Listening on %s...\n", addr)
  err = http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
  }  
  
}
