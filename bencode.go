package main

import bencode "github.com/zeebo/bencode"

import (
	"os"
	"sort"
	"fmt"
	"strings"
	"syscall"
	"sync"
)


type genericData struct {
	val string
	pat string
	key string
	path string
}

type bencodeData struct {
	benmap map[string]interface{}
	encmap map[string]interface{}
	table []string
}

func (gd *genericData) construct(v string, p string, k string, d string) {
	gd.val = v
	gd.pat = p
	gd.key = k
	gd.path = d
}

/*
TODO
For now we let the goroutine spawn as many gophers it want, but it might be needed to limit them to the OS rlimit size
*/

func EditDataSync(data []bencodeData, gd genericData) {

	var (
		wg sync.WaitGroup
	)

	wg.Add(len(data))

	for i := 0; i < len(data); i += 1 {
		go func(i int) {
			defer wg.Done()
			data[i].EditData(gd)
		}(i)
	}
	wg.Wait()

}


func DecodeSync(fd folderData) []bencodeData{
	var (
		wg sync.WaitGroup
	)
	var rlimit syscall.Rlimit
	syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit)

	arrData := make([]bencodeData, len(fd))
	wg.Add(len(fd))

	for i := 0; i < len(fd); i += 1 {
		go func(i int) {
			defer wg.Done()
			arrData[i].DecodeFile(fd[i])
			arrData[i].SortBenmap()
		}(i)
	}
	wg.Wait()
	return arrData
}




// TODO #3
func (data *bencodeData) EncodeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0664)
	check(err)
	defer file.Close()

	encode := bencode.NewEncoder(file)
	if err := encode.Encode(data.benmap); err != nil {
		panic(err)
	}
}

// TODO #3
func (data *bencodeData) DecodeFile(filename string) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	decode := bencode.NewDecoder(file)
	var tmp map[string]interface{}
	if err := decode.Decode(&tmp); err != nil {
		panic(err)
	}
	data.benmap = tmp
}

func (data *bencodeData) EditData(gd genericData) {
	tmp := data.benmap[gd.key]


	switch tmp := tmp.(type) {
	case int64:
		data.benmap[gd.key] = gd.val
	case string:
		if gd.pat != "" {
			tmp = strings.Replace(tmp, gd.pat, gd.val, 1)
			data.benmap[gd.key] = strings.Replace(tmp, "//", "/", -1)
		} else {
			data.benmap[gd.key] = strings.Replace(tmp, tmp, gd.val, 1)
		}
	case []interface{}:
		fmt.Printf("interface type: %T\n", tmp)
		// TODO
		// Not Implemented Yet
	case map[string]interface{}:
		fmt.Printf("interface type: %T\n", tmp)
		// TODO
		// Not Implemented Yet
	default:
		fmt.Printf("default type: %T\n", tmp)


	}
}
func (data *bencodeData) SortBenmap() {
	for t := range data.benmap {
		data.table = append(data.table, t)
	}
	sort.Strings(data.table)
}
