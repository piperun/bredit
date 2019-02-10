package main


import (
        "fmt"
        "flag"
)

type bencodeData struct {
        benmap map[string]interface{}
        table []string
}

type optFlags struct {
        guiFlag *bool
        keyFlag *string
        valFlag *string
        args    []string
}

func check(e error) {
        if e != nil {
                panic(e)
        }
}


func (o *optFlags) construct() {
        o.guiFlag = flag.Bool("g", false, "[BOOL] Run bredit in GUI mode")
        o.keyFlag = flag.String("k", "", "[STR] bencode key")
        o.valFlag = flag.String("val", "", "[STR] value to replace key's value")
        o.args = flag.Args()


}

func (o *optFlags) checkFlags() {
        if *o.guiFlag != false {
                // Start Qt
        } else {
                if *o.keyFlag == "" || *o.valFlag == "" {
                        err := "key & val flag not set, quitting"
                        panic(err)
                }
        }
}

func (o *optFlags) checkArgs() {
        if *o.guiFlag == false {
                if o.args[0] == "" {
                        err := "no directory or file provided in CLI mode"
                        panic(err)
                }
        }
}

func main() {
        var (
                opt optFlags
        )

        opt.construct()

        flag.Parse()
        opt.checkFlags()
}
