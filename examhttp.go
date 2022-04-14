package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)



type Task struct {
	Description    string 
	Done bool
}


//j'ai pas reuissi a remplir la slice  tasks = append(tasks, &Task{{"Faire les courses" false}})
var tasks = make([]Task, 10)


func listt(rw http.ResponseWriter, _ *http.Request){
	var task1,task2 Task
	task1.Description="Faire les courses"
	task1.Done=false
	task2.Description="Faire les courses"
	task2.Done=false
	tasks = append(tasks,task1)
	tasks = append(tasks,task2) 
	for i := 1; i < len(tasks); i++ {
	if tasks[i].Done==false{
	 machaine:=tasks[i].Description
	 rw.WriteHeader(http.StatusOK)
     rw.Header().Write([]byte(machaine)) }
  }
}
  func done(rw http.ResponseWriter,r *http.Request){
	  switch r.Method{
	  case "GET" :
		for i := 1; i < len(tasks); i++ {
		if tasks[i].Done==true{
			machaine:=tasks[i].Description
			rw.WriteHeader(http.StatusOK)
			rw.Header().Write([]byte(machaine)) 
			 }
		}
	  case "POST":
		body, err :=ioutil.ReadAll(r.Body) 
	    if err !=nil {
	       fmt.Printf("Error reading body: %v", err)
	       http.Error(rw,"can't read body", http.StatusBadRequest,)
	       intbody:=int(body)
		   for i := 1; i < len(tasks); i++ {
		    tasks[i].Done=false	
		}
	  }
	  default :
      rw.WriteHeader(http.StatusBadRequest)
  }
  func add(rw http.ResponseWriter,r *http.Request){
  methode:=r.Method
  if methode!= "POST"{
	rw.WriteHeader(http.StatusBadRequest)
  }else{
	body, err :=ioutil.ReadAll(r.Body) 
	if err !=
	nil {
	fmt.Printf("Error reading body: %v", err)
	http.Error(rw,"can't read body", http.StatusBadRequest,)
	var slicebody  Task
	 slicebody.Description=string(body)
	 tasks = append(tasks,slicebody)
     rw.WriteHeader(http.StatusOK)
	return
	}
  }
  }

func main() {
 
	http.Handle("/", list())
	http.Handle("/done", done())
	http.Handle("/add", add())

	http.ListenAndServe("localhost:8000", nil)
}



