X
package main

import bencode "github.com/zeebo/bencode"

import (
        "os"
        "sort"
)



func ReadFile(filename string) map[string]interface{}{
        file, err := os.Open(filename)
        check(err)
        defer file.Close()

        decode := bencode.NewDecoder(file)
        var data map[string]interface{}
        if err := decode.Decode(&data); err != nil {
                panic(err)
        }
        return data
}

func SortBenmap(benmap map[string]interface{}) []string {
        table := make([]string, 0, len(benmap))
        for t := range benmap {
                table = append(table, t)
        }

        sort.Strings(table)
        return table
}
