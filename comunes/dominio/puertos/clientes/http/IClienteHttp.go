package comunes_http_clientes

import comunes_entidades "genesis/ho/ms-autorized/comunes/dominio/entidades"

type IClienteHttp interface {
	Enviar(mensaje *comunes_entidades.HttpRequest) (*comunes_entidades.HttpResponse, error)
}
