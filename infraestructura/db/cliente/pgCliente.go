package infraestructura_db_cliente

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ClienteDB struct {
	Conn        *pgxpool.Pool
	UrlConecion string
	contexto    context.Context
}

func (CDB *ClienteDB) Select(query string, argumentos []any) ([][]interface{}, error) {
	connecion, err := CDB.Conn.Acquire(CDB.contexto)

	if err != nil {
		return nil, err
	}
	defer connecion.Release()
	err = connecion.Ping(CDB.contexto)
	if err != nil {
		return nil, err
	}
	rows, err := connecion.Query(CDB.contexto, query, argumentos...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var resultado [][]interface{}
	for rows.Next() {
		valores, err := rows.Values()

		if err != nil {
			return nil, err
		}
		// fmt.Printf("%+v", valores[0])
		resultado = append(resultado, valores)

	}

	return resultado, nil

}

const defaultMaxConns = int32(4)
const defaultMinConns = int32(0)
const defaultMaxConnLifetime = time.Hour
const defaultMaxConnIdleTime = time.Minute * 30
const defaultHealthCheckPeriod = time.Minute
const defaultConnectTimeout = time.Second * 5

func InicializarCliente(UrlConn string) (*ClienteDB, error) {
	cliente := &ClienteDB{
		contexto:    context.Background(),
		UrlConecion: UrlConn,
	}
	if dbConfig, err := pgxpool.ParseConfig(cliente.UrlConecion); err != nil {
		return nil, err
	} else {
		dbConfig.MaxConns = defaultMaxConns
		dbConfig.MinConns = defaultMinConns
		dbConfig.MaxConnLifetime = defaultMaxConnLifetime
		dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
		dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
		dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout
		cliente.Conn, err = pgxpool.NewWithConfig(cliente.contexto, dbConfig)

		if err != nil {
			return nil, err
		}
	}

	return cliente, nil
}
