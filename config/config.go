package config

import (
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

// ConfStruct is used to unmarshal the config.toml
type ConfStruct struct {
	Prefix            []string `toml:"Prefix"`
	BotToken          string   `toml:"Bot_Token"`
	MongoURL          string   `toml:"Mongo_URL"`
	BotStatus         string   `toml:"Bot_Status"`
	MaxCharacterRoll  uint64   `toml:"Max_Character_Roll"`
	MaxCharacterDrop  uint     `toml:"Max_Character_Drop"`
	DropsOnInteract   uint64   `toml:"Drops_On_Interact"`
	ListLen           int      `toml:"List_Len"`
	ListMaxUpdateTime duration `toml:"List_Max_Update_Time"`
	TimeBetweenRolls  duration `toml:"Time_Between_Rolls"`
	LogToFile         bool     `toml:"Log_To_File"`
}
type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) (err error) {
	d.Duration, err = time.ParseDuration(string(text))
	return
}

// Retrieve retrieves the config from the file
func Retrieve(filename string) (config ConfStruct) {
	if _, err := toml.DecodeFile(filename, &config); err != nil {
		log.Fatalln("Couldn't read configuration : ", err)
	}
	return
}
