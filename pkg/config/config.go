package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// DatabaseConfigurations exported
type Configurations struct {
	APIPort    string `mapstructure:"API_PORT"`
	DBName     string `mapstructure:"DBNAME_SCAPS"`
	DBUser     string `mapstructure:"DBUSER_SCAPS"`
	DBPassword string `mapstructure:"DBPASSWORD_SCAPS"`
	DBSSLmode  string `mapstructure:"DBSSL_MODE_SCAPS"`
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

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Set the file name of the configurations file
	viper.SetConfigName(configFileNameRoot)
	viper.SetConfigType(configFileNameExt)

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
    if err != nil {
        return
    }

    err = viper.Unmarshal(&configurations)

	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("Database is\t", configurations.DBName)
	fmt.Println("APIPort is\t\t", configurations.APIPort)
	fmt.Println("DBUser is\t\t", configurations.DBUser)
	fmt.Println("DBPassword is\t\t", printMaskedSecret(configurations.DBPassword, 4))
	fmt.Println("DBSSLmode is\t\t", configurations.DBSSLmode)

	return
}

// https://medium.com/@bnprashanth256/reading-configuration-files-and-environment-variables-in-go-golang-c2607f912b63
// https://dev.to/techschoolguru/load-config-from-file-environment-variables-in-golang-with-viper-2j2d