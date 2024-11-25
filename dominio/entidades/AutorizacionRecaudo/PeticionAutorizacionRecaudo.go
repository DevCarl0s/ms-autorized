package entidades_autorizacion

type PeticionAutorizacionRecaudo struct {
	TipoAutorizacion string `json:"medio_autorizacion"`
	EdsId            int    `json:"eds_id"`
	Autorizacion     string `json:"data"`
	Cara             int    `json:"cara"`
}
