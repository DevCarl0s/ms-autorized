package casosuso_comunes

import (
	repositorios_comunes "genesis/ho/ms-autorized/comunes/dominio/repositorios"
	entidades_autorizacion "genesis/ho/ms-autorized/dominio/entidades/AutorizacionRecaudo"
)

type ObtenerInformacionCliente struct {
	InformacionClienteRepo repositorios_comunes.IObtenerInformacionCliente
}

func (RDE *ObtenerInformacionCliente) Ejecutar(body *entidades_autorizacion.PeticionAutorizacionRecaudo) (*entidades_autorizacion.InformacionCliente, error) {
	return RDE.InformacionClienteRepo.Consultar(body)
}
