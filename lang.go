package lepet

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"golang.org/x/exp/maps"
)

// Configuration type that is used by the Init function. Sets the Active language and loads key-val json structured data into it.
type LangCfg struct {
	Active   string
	Path     string
	FileName string `yaml:"fileName"`
}

// Main types
type LangItem map[string]string
type LangData struct {
	Active string
	List   map[string]LangItem
}

// Instance of the language
var I *LangData // instance

// Generate LangData instance with `Actice` set to `en`, and `List` having `en` as an index of empty map.
func New() LangData {
	return LangData{
		Active: "en",
		List: map[string]LangItem{
			"en": map[string]string{},
		},
	}
}

// Initilize built-in instance based on the given LangCfg value as configuration.
func Init(conf *LangCfg) error {
	I = &LangData{}
	if conf.Active != "" {
		I.Active = conf.Active
	} else {
		I.Active = "en"
	}
	I.List = map[string]LangItem{I.Active: defaultList}

	jsonFile, err := os.Open(fmt.Sprintf("%v/%v", conf.Path, conf.FileName))
	if err != nil {
		panic("failed to load source file. " + err.Error())
	}
	defer jsonFile.Close()

	var myLI LangItem = LangItem{}
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &myLI)
	maps.Copy(I.List[I.Active], myLI)
	return nil
}

// Adds langauge
func (a *LangData) Add(name string) {
	_, ok := a.List[name]
	if !ok {
		a.List[name] = LangItem{}
	}
}

// Adds a message to the active language
func (a *LangData) AddMsg(code string, message string, opt ...string) {
	lang := a.Active
	if len(opt) > 0 {
		a.Add(opt[1])
		lang = opt[1]
	}
	a.List[lang][code] = message
}

// Add messages to the active language
func (a *LangData) AddMsgList(list LangItem, opt ...string) {
	lang := a.Active
	if len(opt) > 0 {
		a.Add(opt[1])
		lang = opt[1]
	}

	maps.Copy(a.List[lang], list)
}

// Get message from the active language
func (a *LangData) Msg(k string, opt ...string) string {
	lang := a.Active
	if len(opt) > 0 {
		a.Add(opt[1])
		lang = opt[1]
	}

	if msg, ok := a.List[lang][k]; !ok {
		return "** warning: usage of unlisted code **"
	} else {
		return msg
	}
}
