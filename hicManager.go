package data

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nimezhu/indexed/hic"
	"github.com/nimezhu/netio"
)

/*HicManager implement DataManager Inteface */
type HicManager struct {
	uriMap  map[string]string
	dataMap map[string]*hic.HiC
	dbname  string
}

func (m *HicManager) Add(key string, reader io.ReadSeeker, uri string) error {
	vhic, err := hic.DataReader(reader)
	if err != nil {
		return err
	}
	m.uriMap[key] = uri
	m.dataMap[key] = vhic
	return nil
}
func (m *HicManager) AddURI(uri string, key string) error {
	m.uriMap[key] = uri
	m.dataMap[key] = readhic(uri)
	log.Println("add uri", uri, key)
	return nil
}
func (m *HicManager) Del(key string) error {
	delete(m.uriMap, key)
	delete(m.dataMap, key)
	return nil
}
func (m *HicManager) Get(key string) (string, bool) {
	v, ok := m.uriMap[key]
	return v, ok
}
func (m *HicManager) Move(key1 string, key2 string) bool {
	v, ok1 := m.uriMap[key1]
	d, ok2 := m.dataMap[key1]
	if ok1 && ok2 {
		m.uriMap[key2] = v
		m.dataMap[key2] = d
		delete(m.uriMap, key1)
		delete(m.dataMap, key1)
	}
	return ok1 && ok2
}
func (m *HicManager) List() []string {
	keys := []string{}
	for k, _ := range m.uriMap {
		keys = append(keys, k)
	}
	return keys
}
func (m *HicManager) ServeTo(router *mux.Router) {
	prefix := "/" + m.dbname
	router.HandleFunc(prefix+"/ls", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		jsonHic, _ := json.Marshal(m.uriMap)
		w.Write(jsonHic)
	})
	AddHicsHandle(router, m.dataMap, prefix)
}

func NewHicManager(uri string, dbname string) *HicManager {
	uriMap := LoadURI(uri)
	dataMap := make(map[string]*hic.HiC)
	dataList := []string{}
	for k, v := range uriMap {
		dataMap[k] = readhic(v)
		dataList = append(dataList, k)
	}
	/*
		if _, ok := dataMap["default"]; ok {
			//has default
		} else {
			dataMap["default"] = dataMap[dataList[0]]
		}
	*/
	m := HicManager{
		uriMap,
		dataMap,
		dbname,
	}
	//m.ServeTo(router)
	return &m
}
func InitHicManager(dbname string) *HicManager {
	uriMap := make(map[string]string)
	dataMap := make(map[string]*hic.HiC)
	m := HicManager{
		uriMap,
		dataMap,
		dbname,
	}
	return &m
}

func readhic(uri string) *hic.HiC {
	reader, err := netio.NewReadSeeker(uri)
	checkErr(err)
	vhic, err := hic.DataReader(reader)
	checkErr(err)
	return vhic
}
