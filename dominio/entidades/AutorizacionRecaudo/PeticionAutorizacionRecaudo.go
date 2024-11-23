package entidades_autorizacion

type PeticionAutorizacionRecaudo struct {
	TipoAutorizacion string `json:"tipo_autorizacion"`
	EdsId            int    `json:"eds_id"`
	Autorizacion     string `json:"autorizacion"`
}
