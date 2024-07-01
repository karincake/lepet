# Lepet - Simple Dictionary
A simple dictionary based language manager

# Usage
Getting the library
```
go get github.com:karincake/lepet
```

Using the code
```Go
package main

import (
	"fmt"

	"github.com/karincake/lepet"
)

func main() {
	myLepet := lepet.New()
	myLepet.AddMsg("hello-world", "Hello world and all people")
	fmt.Println(myLepet.Msg("hello-world"))
}
```

Using built-in instance and load a saved data
```Go
package main

import (
	"fmt"

	"github.com/karincake/lepet"
)

func main() {
	cfg := &lepet.LangCfg{
		Active:   "en",
		Path:     ".",
		FileName: "data.json",
	}
	lepet.Init(cfg)
	fmt.Println(lepet.I.Msg("request-bad"))
}
```

Example of data that can be used as source, using basic key-val structure
```json
{
	"request-ok": "request is done succesfully",
	"request-methodNotAllowed": "request has been rejected, method is not allowed",
	"request-toomany": "too many request",
	"payload-bad":   "invalid data structre"
}
```

