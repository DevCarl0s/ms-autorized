package repositorios_comunes

import entidades_autorizacion "genesis/ho/ms-autorized/dominio/entidades/AutorizacionRecaudo"

type IObtenerInformacionCliente interface {
	Consultar(datos *entidades_autorizacion.PeticionAutorizacionRecaudo) (*entidades_autorizacion.InformacionCliente, error)
}
