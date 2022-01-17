package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/input/gg", Pos).Methods("GET")
	r.HandleFunc("/input", getdata).Methods("GET") // выводит данные
	log.Fatal(http.ListenAndServe(":8000", r))
}

var Comand string
var x = 10
var y = 19

//var CurentPosition = mas[10][19]
var Win = mas[10][19]

var mas = [20][20]int{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //0
	/*                   */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //1
	/*                   */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0}, //2
	/*                   */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, //3
	/*                   */ {0, 0, 0, 0, 0, 0, 2, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, //4/**//**//**/
	/*                   */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, //5
	/*                   */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, //6
	/*                   */ {0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0}, //7
	/*                   */ {0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //8
	/*                   */ {0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //9
	/*                   */ {0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //10
	/*                   */ {0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //11
	/*                   */ {0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //12
	/*                   */ {0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //13
	/*                   */ {0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //14
	/*                   */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //15
	/*                   */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //16
	/*                   */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //17
	/*                   */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //18
	/*                   */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}} //20

func getdata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if Comand == "/cheak" {
		if mas[y][x-1] != 0 {
			//	CurentPosition = mas[y][x-1]
		} else if mas[y][x-1] == 2 {
			json.NewEncoder(w).Encode("Вы дошли до выхода")
		}
	}

	json.NewEncoder(w).Encode(&Comands{}) // вывод данных
}

type Com struct {
	Cp string `json:"comand"`
}

type Comands struct {
	Positionup    int `json:"comand"`
	Positiondown  int `json:"comand1"`
	Positionleft  int `json:"comand12"`
	Positionright int `json:"comand13"`
}

func Pos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var hh Com
	_ = json.NewDecoder(r.Body).Decode(&hh)
	if Comand == "/init" {
	}
	if hh.Cp == "/move_top" {
		if mas[y-1][x] != 0 {
			//CurentPosition = mas[y-1][x]
			y -= 1
		} else if mas[y-1][x] == 2 {
			json.NewEncoder(w).Encode("Вы дошли до выхода")
		}
		json.NewEncoder(w).Encode("Команда /move_top принята")
	}
	if hh.Cp == "/move_down" {
		if mas[y+1][x] != 0 {
			//CurentPosition = mas[y+1][x]
			y += 1
		} else if mas[y+1][x] == 2 {
			json.NewEncoder(w).Encode("Вы дошли до выхода")
		}
		json.NewEncoder(w).Encode("Команда /move_down принята")
	}
	if hh.Cp == "/move_left" {
		if mas[y][x-1] != 0 {
			//CurentPosition = mas[y][x-1]
			x -= 1
		} else if mas[y][x-1] == 2 {
			json.NewEncoder(w).Encode("Вы дошли до выхода")
		}
		json.NewEncoder(w).Encode("Команда /move_left принята")
	}
	if hh.Cp == "/move_right" {
		if mas[y][x+1] != 0 {
			//CurentPosition = mas[y][x+1]
			x += 1
		} else if mas[y][x+1] == 2 {
			json.NewEncoder(w).Encode("Вы дошли до выхода")
		}
		json.NewEncoder(w).Encode("Команда /move_right принята")
	}
	if "/check" == hh.Cp {

		json.NewEncoder(w).Encode(Comands{Positionleft: mas[y][x-1], Positionright: mas[y][x+1], Positionup: mas[y-1][x], Positiondown: mas[y+1][x]})

	}
	if "/position" == hh.Cp {

		json.NewEncoder(w).Encode(y)
		json.NewEncoder(w).Encode(x)
	}
	json.NewEncoder(w).Encode(hh)
}
