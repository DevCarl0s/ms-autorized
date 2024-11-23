package repositorios_infraestructura

import (
	"encoding/json"
	comunes_db_clientes "genesis/ho/ms-autorized/comunes/dominio/puertos/clientes/db"
	entidades_autorizacion "genesis/ho/ms-autorized/dominio/entidades/AutorizacionRecaudo"
	"log"
)

type ObtenerInformacionCliente struct {
	Cliente comunes_db_clientes.IClienteDB
}

func (OIC *ObtenerInformacionCliente) Consultar(datos *entidades_autorizacion.PeticionAutorizacionRecaudo) (*entidades_autorizacion.InformacionCliente, error) {
	jsonDatos, _ := json.Marshal(datos)
	args := []any{string(jsonDatos)}
	query := `SELECT * FROM gnv.fnc_obtener_informacion_cliente($1)`

	var errorStr string
	respuesta, err := OIC.Cliente.Select(query, args)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		errorStr = "Error al ejecutar la consulta en la base de datos"
		return &entidades_autorizacion.InformacionCliente{Success: false, Message: errorStr}, err
	}

	bytes, err := json.Marshal(respuesta[0][0])
	if err != nil {
		errorStr = "Error convertir respuesta a json"
		return &entidades_autorizacion.InformacionCliente{Success: false, Message: errorStr}, err
	}

	var cliente *entidades_autorizacion.InformacionCliente
	err = json.Unmarshal(bytes, &cliente)
	if err != nil {
		return &entidades_autorizacion.InformacionCliente{Success: false, Message: errorStr}, err
	}

	return cliente, nil
}
