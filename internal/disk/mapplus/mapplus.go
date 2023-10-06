package mapplus

import (
	"fmt"

	"github.com/fengdotdev/goerrorsplus/e"
	"github.com/google/uuid"
)


const (
	D	string= "data"
	G string = "garbage"
	N string = "none"
)


type keyAndLocation struct {
	key string
	location string
}


type move struct {
	id string
	from string
	to string
}


type Map struct {
	idKey map[string]keyAndLocation	
	data    map[string]interface{}
	garbage map[string]interface{}
	history []move
}

func NewMap() *Map {
	return &Map{
		idKey: make(map[string]keyAndLocation),
		data:    make(map[string]interface{}),
		garbage: make(map[string]interface{}),
		history: make([]move, 0),
	}
}
//WORKING HERE
// TESTME
func (m *Map) Add(key string, value interface{}) *e.ErrorPlus{


//generate id 
	id, err := generateID()
	if err != nil {
	return err
	}

//add to idKey
	m.idKey[id]=keyAndLocation{
		key: key,
		location: D,
	}

//add to data
	m.data[id]=value
	m.history=append(m.history, move{id: id, from: N, to: D})
	return nil
}

// TESTME
func (m *Map) Get(key string) interface{} {
	panic("implement me")
}

// TESTME
func (m *Map) Remove(key string) {
	panic("implement me")
}

// TESTME
func (m *Map) Undo() {
	panic("implement me")
}

// TESTME
func (m *Map) Redo() {
	panic("implement me")
}

// TESTME
func (m *Map) Clear() {
	panic("implement me")
}

// TESTME
func (m *Map) DropGarbage() {
	panic("implement me")
}

func (m *Map) SelfMantain() {
	panic("implement me")
}


func generateID() (s string,err *e.ErrorPlus){
	defer func() {
		if r := recover(); r != nil {
			err=e.E(fmt.Errorf("generateID panic"), "recover from panic using UUID of google", []string{"mapplus", "generateID"}, generateID, r)
			s=""
		}
	}()
	i:=uuid.New()
	return i.String(), nil
}