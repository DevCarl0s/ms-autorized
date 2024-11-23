package comunes_db_clientes

type IClienteDB interface {
	Select(query string, argumentos []any) ([][]interface{}, error)
}
