package entidades_autorizacion

type Cliente struct {
	ID                    *int    `json:"id"`
	Nombres               *string `json:"nombres"`
	Apellidos             *string `json:"apellidos"`
	Telefono              *int64  `json:"telefono"`
	Direccion             *string `json:"direccion"`
	EstadoID              *int    `json:"estado_id"`
	GeneroID              *int    `json:"genero_id"`
	MunicipioID           *int    `json:"municipio_id"`
	RazonSocial           *string `json:"razon_social"`
	FechaCreacion         *string `json:"fecha_creacion"`
	Identificacion        *string `json:"identificacion"`
	EstadoCivilID         *int    `json:"estado_civil_id"`
	TipoClienteID         *int    `json:"tipo_cliente_id"`
	FechaNacimiento       *string `json:"fecha_nacimiento"`
	CorreoElectronico     *string `json:"correo_electronico"`
	CreacionOrigenID      *int    `json:"creacion_origen_id"`
	FechaModificacion     *string `json:"fecha_modificacion"`
	DigitoVerificacion    *int    `json:"digito_verificacion"`
	UsuarioCreacionID     *int    `json:"usuario_creacion_id"`
	TipoIdentificacionID  *int    `json:"tipo_identificacion_id"`
	UsuarioModificacionID *int    `json:"usuario_modificacion_id"`
}

// Vehiculo representa los datos de un vehículo.
type Vehiculo struct {
	Placa                 *string `json:"placa"`
	EstadoID              *int    `json:"estado_id"`
	IDFamilia             *int    `json:"id_familia"`
	IDVehiculo            *int    `json:"id_vehiculo"`
	FechaCreacion         *string `json:"fecha_creacion"`
	IDOrganizacion        *int    `json:"id_organizacion"`
	IDTipoServicio        *int    `json:"id_tipo_servicio"`
	IDTipoVehiculo        *int    `json:"id_tipo_vehiculo"`
	FechaModificacion     *string `json:"fecha_modificacion"`
	UsuarioCreacionID     *int    `json:"usuario_creacion_id"`
	UsuarioModificacionID *int    `json:"usuario_modificacion_id"`
}

type ClienteVehiculo struct {
	TblClientesID       *int `json:"tbl_clientes_id"`
	IDClienteVehiculo   *int `json:"id_cliente_vehiculo"`
	VehiculosIDVehiculo *int `json:"vehiculos_id_vehiculo"`
}

type ClienteDestino struct {
	IDEstado                          *int `json:"id_estado"`
	TblClientesID                     *int `json:"tbl_clientes_id"`
	IDClienteDestino                  *int `json:"id_cliente_destino"`
	TblClientesTiposDestinosIDDestino *int `json:"tbl_clientes_tipos_destinos_id_destino"`
}

type ClienteTipoDestino struct {
	IDEstado    *int    `json:"id_estado"`
	IDDestino   *int    `json:"id_destino"`
	Descripcion *string `json:"descripcion"`
}

type RecaudoVehiculo struct {
	IDEstado              *int    `json:"id_estado"`
	IDMedioPago           *int    `json:"id_medio_pago"`
	FechaCreacion         *string `json:"fecha_creacion"`
	FechaModificacion     *string `json:"fecha_modificacion"`
	IDRecaudoVehiculo     *int    `json:"id_recaudo_vehiculo"`
	UsuarioCreacionID     *int    `json:"usuario_creacion_id"`
	VehiculosIDVehiculo   *int    `json:"vehiculos_id_vehiculo"`
	UsuarioModificacionID *int    `json:"usuario_modificacion_id"`
}

type Concepto struct {
	Valor                       float64 `json:"valor"`
	IDConceptoVehiculo          int     `json:"id_concepto_vehiculo"`
	ConceptosRecaudosIDConcepto int     `json:"conceptos_recaudos_id_concepto"`
	RecaudosVehiculosIDRecaudo  int     `json:"recaudos_vehiculos_id_recaudo_vehiculo"`
}
type RecaudoVehiculosConcepto struct {
	Concepto        Concepto        `json:"concepto"`
	ConceptoRecaudo ConceptoRecaudo `json:"concepto_recaudo"`
}

type VehiculoIdentificador struct {
	Valor                                string `json:"valor"`
	IDEstado                             int    `json:"id_estado"`
	VehiculosIDVehiculo                  int    `json:"vehiculos_id_vehiculo"`
	IDVehiculoIdentificadores            int    `json:"id_vehiculo_identificadores"`
	TblTiposIdentificacionVehiculoIDTipo int    `json:"tbl_tipos_identificacion_vehiculo_id_tipo_vehi"`
}

type TipoIdentificacionVehiculo struct {
	IDEstado    int    `json:"id_estado"`
	Descripcion string `json:"descripcion"`
	IDTipoVehi  int    `json:"id_tipo_vehi"`
}

type VehiculoPropiedades struct {
	VIN                 string `json:"vin"`
	Color               string `json:"color"`
	Linea               string `json:"linea"`
	Marca               string `json:"marca"`
	Modelo              string `json:"modelo"`
	IDUnidad            int    `json:"id_unidad"`
	Cilindraje          int    `json:"cilindraje"`
	CapacidadMaxima     int    `json:"capacidad_maxima"`
	IDVehiculoPropiedad int    `json:"id_vehiculo_propiedad"`
	VehiculosIDVehiculo int    `json:"vehiculos_id_vehiculo"`
}

type ConceptoRecaudo struct {
	Concepto   string `json:"concepto"`
	IDEstado   int    `json:"id_estado"`
	IDConcepto int    `json:"id_concepto"`
}

type ClienteData struct {
	Vehiculo                   Vehiculo                   `json:"vehiculo"`
	Cliente                    Cliente                    `json:"cliente"`
	ClienteVehiculo            ClienteVehiculo            `json:"cliente_vehiculo"`
	ClienteDestino             ClienteDestino             `json:"cliente_destino"`
	ClienteTipoDestino         ClienteTipoDestino         `json:"cliente_tipo_destino"`
	RecaudoVehiculo            RecaudoVehiculo            `json:"recaudo_vehiculo"`
	RecaudoVehiculosConcepto   []RecaudoVehiculosConcepto `json:"recaudo_vehiculos_conceptos"`
	VehiculoIdentificador      VehiculoIdentificador      `json:"vehiculo_identificador"`
	TipoIdentificacionVehiculo TipoIdentificacionVehiculo `json:"tipo_identificacion_vehiculo"`
	VehiculoPropiedades        VehiculoPropiedades        `json:"vehiculo_propiedades"`
}

type InformacionCliente struct {
	Data    []ClienteData `json:"data"`
	Message string        `json:"message"`
	Success bool          `json:"success"`
}
