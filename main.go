package main



/*
TODO
* Re-Implement the relation between args and folderData, make it so that folderData parses either the entire directory or paths provided by os.args
* Some other non-sense 

*/



import (
	"flag"
	"fmt"
)

type options struct {
	cliFlag *bool
	keyFlag *string
	valFlag *string
	patFlag *string
	dryFlag *bool
	args	[]string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func (o *options) construct() {
	o.cliFlag = flag.Bool("c", true, "Run bredit in CLI mode")
	o.keyFlag = flag.String("k", "", "Bencode key")
	o.valFlag = flag.String("v", "", "Expression to replace key's value")
	o.patFlag = flag.String("p", "", "Pattern used to replace key's value")
	o.dryFlag = flag.Bool("d", false, "Doesn't overwrite files")

	flag.Parse()
	o.args = flag.Args()
}

// Re-write this with proper function & switch case
/*
proposal:
// this can be a lambda 
func check_all_flags -> {
	if empty -> return num
	if not empty -> return num+1
	etc.
}

switch func
case num
print err
kill
case num+1
print err+1
kill
*/
func (o *options) check() {
	if *o.cliFlag {
		if *o.keyFlag == "" || *o.valFlag == "" && *o.patFlag == "" {
			err := "key & val or pat flag not set, quitting"
			panic(err)
		}
		if o.args[0] == "" {
			err := "no directory or file provided in CLI mode"
			panic(err)
		}
	}
}

func (o *options) checkValidKey(data bencodeData) {
	if *o.cliFlag {
		if err, ok := data.benmap[*o.keyFlag]; !ok {
			err = "The key provided does not exist!"
			panic(err)
		}
	}
}

func main() {
	var (
		opt options
		data bencodeData
		generic genericData
		folder folderData
	)
	ext := make(map[string]string)
	ext["rtorrent"] = ".rtorrent"

	opt.construct()
	opt.check()
	if *opt.cliFlag {
		generic.construct(*opt.valFlag, *opt.patFlag, *opt.keyFlag, opt.args[0])
	} else {
		// for GUI
	}

	pathType, err := CheckPathType(generic.path)
	check(err)

	switch pathType {
	case 1:
		folder.ParseFolder(generic.path, ext)
		arrData := make([]bencodeData, len(folder))
		arrData = DecodeSync(folder)
		EditDataSync(arrData, generic)
		if !*opt.dryFlag {
			EncodeSync(arrData, folder)
		}

	case 2:
		data.DecodeFile(generic.path)
		data.SortBenmap()
		data.EditData(generic)
		if !*opt.dryFlag {
			data.EncodeFile(generic.path)
		} else {
			fmt.Printf("Changed key:\n%s\nvalue to:\n%s\nin file:\n%s\n", generic.key, data.benmap[generic.key], generic.path)
		}


	}

}
