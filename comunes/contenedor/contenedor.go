package contenedor

import (
	servicio_autorizacionrecaudo "genesis/ho/ms-autorized/aplicacion/servicios/AutorizacionRecaudo"
	casosuso_comunes "genesis/ho/ms-autorized/comunes/aplicacion/casosuso"
	infraestructura_db_cliente "genesis/ho/ms-autorized/infraestructura/db/cliente"
	repositorios_infraestructura "genesis/ho/ms-autorized/infraestructura/db/repositorios"
	infraestructura_http_cliente "genesis/ho/ms-autorized/infraestructura/http/cliente"
	"log"
	"os"
)

var ProcesarAutorizacion *servicio_autorizacionrecaudo.ProcesarAutorizacion

// INICIALIZA LAS INJECIONES DEL SERIVICIO
func InicializarContenedor() error {

	var err error

	//CLIENTES
	clienteDB, err := infraestructura_db_cliente.InicializarCliente(os.Getenv("DB_POSTGRES"))
	if err != nil {
		log.Panic(err)
		return err

	}

	/* clienteHttp, err := */
	infraestructura_http_cliente.InicializarCliente()
	// if err != nil {
	// 	log.Panic(err)
	// 	return err

	// }

	//REPOSITORIOS
	obtenerInformacionClienteRepo := &repositorios_infraestructura.ObtenerInformacionCliente{
		Cliente: clienteDB,
	}

	// CASO DE USOS
	obtenerInformacionCliente := &casosuso_comunes.ObtenerInformacionCliente{
		InformacionClienteRepo: obtenerInformacionClienteRepo,
	}

	// SERVICIOS

	ProcesarAutorizacion = &servicio_autorizacionrecaudo.ProcesarAutorizacion{
		InformacionCliente: obtenerInformacionCliente,
	}

	return nil

}
