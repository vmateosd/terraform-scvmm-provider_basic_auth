package scvmm

import (
	"fmt"
	"log"
	"strings"

	"github.com/dpotapov/winrm-auth-krb5"
	"github.com/masterzen/winrm"
)

// Config ... SCVMM configuration details
type Config struct {
	ServerIP string
	Port     int
	Username string
	Password string
}

//Connection ... Create a new connection with winrm to Powershell.
func (c *Config) Connection() (*winrm.Client, error) {

	endpoint := winrm.NewEndpoint(c.ServerIP, c.Port, false, false, nil, nil, nil, 0)
log.Printf("[DEBUG] "endpoint creado")
	// Añadido de krb5
	winrm.DefaultParameters.TransportDecorator = func() winrm.Transporter {
		return &winrmkrb5.Transport{}
	}
	//Fin de añadido
	//Lo que habia: winrmConnection, err := winrm.NewClient(endpoint, c.Username, c.Password)
	//Añado linea
	winrmConnection, err := winrm.NewClientWithParameters(endpoint, c.Username, c.Password, winrm.DefaultParameters)
log.Printf("[DEBUG] "despues de la linea añadida de krb5")
	//fin añadido
	if err != nil {
		log.Printf("[ERROR] Failed to connect winrm: %v\n", err)
		return nil, err
	}

	shell, err := winrmConnection.CreateShell()
	if err != nil {
		log.Printf("[Error] While creating Shell %s", err)
		if strings.Contains(err.Error(), "http response error: 401") {
			return nil, fmt.Errorf("[Error] Please check whether username and password are correct.\n Error: %s", err.Error())
		} else if strings.Contains(err.Error(), "unknown error Post") {
			return nil, fmt.Errorf("[Error] Please check whether server ip and port number are correct.\n Error: %s", err.Error())
		} else {
			return nil, fmt.Errorf("[Error] While creating Shell %s", err)
		}
	}
	defer shell.Close()

	log.Printf("[DEBUG] Winrm connection successful")
	return winrmConnection, nil
}
