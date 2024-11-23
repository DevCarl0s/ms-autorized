package comunes_entidades

import entidades_autorizacion "genesis/ho/ms-autorized/dominio/entidades/AutorizacionRecaudo"

type RespuestaAutorizacion struct {
	Estado     string                              `json:"state"`
	Mensaje    string                              `json:"message"`
	Data       *entidades_autorizacion.ClienteData `json:"data"`
	httpEstado int
}

func (RMP *RespuestaAutorizacion) SetHttpEstado(estado int) {
	RMP.httpEstado = estado
}

func (RMP *RespuestaAutorizacion) GetHttpEstado() (estado int) {
	return RMP.httpEstado
}
