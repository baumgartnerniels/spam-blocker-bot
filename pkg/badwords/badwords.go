package badwords

import (
	"encoding/csv"
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

type BadWords struct {
	BadWords []string `json:"strings"`

	cfg *viper.Viper
}

// New creates new BadWords, importing it from CSV
func New(cfg *viper.Viper) (list *BadWords) {
	var err error
	list = &BadWords{cfg: cfg}
	// defer list loading from filesystem if we had an error in the process

	csvbuf, err := ioutil.ReadFile(list.cfg.GetString("badwords.path"))
	if err != nil {
		log.WithError(err).Warn("Unable to import Badwords")
		return
	}
	csvstr := string(csvbuf)
	csvstrreader := strings.NewReader(csvstr)
	csvReader := csv.NewReader(csvstrreader)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.WithError(err).Warn("Unable to read CAS list")
		return
	}

	list.BadWords = make([]string, 0, len(records))
	for _, r := range records {
		if len(r) == 0 {
			continue
		}

		list.BadWords = append(list.BadWords, r[0])
	}

	return
}
