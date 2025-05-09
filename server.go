package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/hprose/hprose-go/hprose"
)

type RemoteMap struct {
	mutex sync.Mutex
	data  map[string]int
}

func NewRemoteMap() *RemoteMap {
	return &RemoteMap{
		data: make(map[string]int),
	}
}

func (m *RemoteMap) Update(key string, value int) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if lastValue, ok := m.data[key]; ok {
		m.data[key] = value
		fmt.Printf("Par atualizado no dicionário: (%s:%d) -> (%s:%d)\n", key, lastValue, key, value)
		return false
	}else{
		m.data[key] = value
		fmt.Printf("Novo par cadastrado no dicionário: (%s:%d)\n", key, value)
		return true
	}
}

func (m *RemoteMap) Get(key string) int {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	value, ok := m.data[key]
	if ok {
		fmt.Printf("Consulta no dicionário a partir da chave '%s': %d\n", key, value)
		return value
	}
	fmt.Printf("Consulta no dicionário a partir da chave '%s': não existe!\n", key)
	return -1
}

func (m *RemoteMap) Remove(key string) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if value, ok := m.data[key]; ok {
		fmt.Printf("Par (%s:%d) foi removido do dicionário!\n", key, value)
		delete(m.data, key)
		return true
	}	
	return false
}

func main() {
	rm := NewRemoteMap()

	service := hprose.NewHttpService()
	service.AddMethods(rm)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", service)
}
