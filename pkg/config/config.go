package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// DatabaseConfigurations exported
type Configurations struct {
	APIPort    string `mapstructure:"API_PORT" validate:"required"`
	DBName     string `mapstructure:"DB_NAME" validate:"required"`
	DBUser     string `mapstructure:"DB_USER" validate:"required"`
	DBPassword string `mapstructure:"DB_PASSWORD" validate:"required"`
	DBSSLmode  string `mapstructure:"DB_SSLMODE" validate:"required"`
	DBMaxOpenConns  int `mapstructure:"DB_MAX_OPEN_CONNS" validate:"required"`
	DBMaxIdleConns  int `mapstructure:"DB_MAX_IDLE_CONNS"`
}

// function to get characters from strings using index
func getChar(str string, index int) rune {
	return []rune(str)[index]
}

func printMaskedSecret(secret string, shownChars int) string {
	passLen := len(secret)

	if shownChars > passLen {
		fmt.Printf("Error: Secret is too short\nExiting..")
		os.Exit(42)
	}

	strSlice := make([]string, passLen)
	for idx := 0; idx < passLen; idx++ {
		if idx <= shownChars {
			strSlice[idx] = string(getChar(secret, idx))
		} else {
			strSlice[idx] = "*"
		}
	}

	return strings.Join(strSlice,"")
}

func Readconfig(configFileNameRoot, configFileNameExt string) (configurations *Configurations, err error) {

	configFileName := fmt.Sprintf("./%s.%s", configFileNameRoot, configFileNameExt)
	fmt.Printf("configFileName: %s\n", configFileName)

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Set the file name of the configurations file
	viper.SetConfigName(configFileNameRoot)
	viper.SetConfigType(configFileNameExt)
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("unable to read configfile %v", err)
		os.Exit(42)
	}


	// Enable VIPER to read Environment Variables
	// Except viper is ridiculously annoying - never using it again
	viper.AutomaticEnv()

    if uerr := viper.UnmarshalExact(&configurations); uerr!=nil {
        log.Fatalf("unable to unmarshall configfile %v", uerr)
		os.Exit(42)
    }
    validate := validator.New(validator.WithRequiredStructEnabled())
    if verr := validate.Struct(configurations); verr!=nil{
        log.Fatalf("Missing required attributes %v\n", verr)
		for _, xerr := range verr.(validator.ValidationErrors) {
			fmt.Println(xerr.Field(), xerr.Tag())
		}
		os.Exit(42)
    }

	// Reading variables using the config struct
	fmt.Println("Reading variables using the config struct..")
	fmt.Println("APIPort is\t\t", configurations.APIPort)
	fmt.Println("DBName is\t\t", configurations.DBName)
	fmt.Println("DBUser is\t\t", configurations.DBUser)
	fmt.Println("DBPassword is\t\t", printMaskedSecret(configurations.DBPassword, 4))
	fmt.Println("DBSSLmode is\t\t", configurations.DBSSLmode)
	fmt.Println("DBMaxOpenConns is\t\t", configurations.DBMaxOpenConns)
	fmt.Println("DBMaxIdleConns is\t\t", configurations.DBMaxIdleConns)



	return
}

// https://medium.com/@bnprashanth256/reading-configuration-files-and-environment-variables-in-go-golang-c2607f912b63
// https://dev.to/techschoolguru/load-config-from-file-environment-variables-in-golang-with-viper-2j2d