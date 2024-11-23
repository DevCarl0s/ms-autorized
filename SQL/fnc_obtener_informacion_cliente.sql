 DROP FUNCTION gnv.fnc_obtener_informacion_cliente(jsonb);

CREATE OR REPLACE FUNCTION gnv.fnc_obtener_informacion_cliente(i_datos jsonb)
 RETURNS jsonb
 LANGUAGE plpgsql
AS $function$
DECLARE

	i_autorizacion_valor 	TEXT;
	i_autorizacion_tipo 	TEXT;
	i_eds_id				INT;

    resultado JSONB;
BEGIN


	IF i_datos::text = '[]' OR i_datos::text = '{}' OR i_datos IS NULL OR i_datos::text = '' THEN
   		RETURN JSONB_BUILD_OBJECT(
			'success', FALSE, 
			'data', '[]'::JSONB,
			'message', 'No se encontraron datos entrantes'
		);
	END IF;

	i_autorizacion_valor 	:= 	( i_datos ->> 'autorizacion' ) :: text;
	i_autorizacion_tipo 	:=	( i_datos ->> 'tipo_autorizacion' ) :: text;
	i_eds_id				:=  ( i_datos ->> 'eds_id' ) :: int;

    SELECT INTO resultado 
        JSON_AGG(
            JSONB_BUILD_OBJECT(
                'vehiculo', to_jsonb(v.*),
				'cliente', to_jsonb(tc.*),
				'cliente_vehiculo', to_jsonb(cv.*),
				'cliente_destino', to_jsonb(tcd.*),
				'cliente_tipo_destino', to_jsonb(tctd.*),
                'recaudo_vehiculo', to_jsonb(rv.*),
                'recaudo_vehiculos_conceptos', to_jsonb(rvc.*),
				'vehiculo_identificador', to_jsonb(vi.*),
				'tipo_identificacion_vehiculo', to_jsonb(ttiv.*),
                'vehiculo_propiedades', to_jsonb(vp.*),
               'concepto_recaudo', to_jsonb(cr.*)
            )
        )
    FROM gnv.vehiculos v
	--INNER JOIN public.tbl_organizacion_empresa toe
		--ON toe.organizacion_id = v.id_organizacion
	--INNER JOIN public.ct_empresas ce 
	--	ON ce.id = toe.empresas_id 
    INNER JOIN gnv.clientes_vehiculos cv 
        ON cv.vehiculos_id_vehiculo = v.id_vehiculo
    INNER JOIN clientes_fidelizacion.tbl_clientes tc 
        ON tc.id = cv.tbl_clientes_id 
	INNER JOIN clientes_fidelizacion.tbl_clientes_destinos tcd 
		ON tcd.tbl_clientes_id = tc.id
	INNER JOIN clientes_fidelizacion.tbl_clientes_tipos_destinos tctd 
		ON tctd.id_destino = tcd.tbl_clientes_tipos_destinos_id_destino
    INNER JOIN gnv.recaudos_vehiculos rv 
        ON rv.vehiculos_id_vehiculo = v.id_vehiculo 
    INNER JOIN gnv.recaudos_vehiculos_conceptos rvc 
        ON rvc.recaudos_vehiculos_id_recaudo_vehiculo = rv.id_recaudo_vehiculo
	INNER JOIN gnv.conceptos_recaudos cr
		ON cr.id_concepto = rvc.conceptos_recaudos_id_concepto
    INNER JOIN gnv.vehiculos_identificadores vi
        ON vi.vehiculos_id_vehiculo = v.id_vehiculo
    INNER JOIN public.tbl_tipos_identificacion_vehiculo ttiv 
		ON ttiv.id_tipo_vehi = vi.tbl_tipos_identificacion_vehiculo_id_tipo_vehi 
    INNER JOIN gnv.vehiculos_propiedades vp 
        ON vp.vehiculos_id_vehiculo = v.id_vehiculo 
	WHERE 
		vi.valor = i_autorizacion_valor
		AND ttiv.descripcion = UPPER(i_autorizacion_tipo)
		--AND ce.id = i_eds_id
;

    IF resultado IS NULL THEN
        RETURN JSONB_BUILD_OBJECT(
			'success', FALSE, 
			'data', '[]'::JSONB,
			'message', 'No se encontro información relacionada'
		);
    END IF;

    -- Devolver el resultado con éxito
    RETURN JSONB_BUILD_OBJECT(
		'success', TRUE, 
		'data', resultado,
		'message', 'Información de cliente obtenida'
	);
EXCEPTION
    WHEN OTHERS THEN
        -- Manejar cualquier error devolviendo un array vacío y un indicador de éxito
        RETURN JSONB_BUILD_OBJECT(
			'success', FALSE, 
			'data', '[]'::JSONB,
			'message', 'Hubo un error en la función: ' || SQLERRM
		);
END;
$function$
;


