package naming

import "Middleware-IF711-2020-3/l5/clientproxy"

type NamingService struct {
	Table map[string]clientproxy.ClientProxy
}

// Register registra o clientProxy no serviço de nomes
func (naming *NamingService) Register(name string, proxy clientproxy.ClientProxy) bool {
	naming.Table[name] = clientproxy.ClientProxy{
		proxy.Host,
		proxy.Port,
		proxy.Id,
		proxy.TypeName,
	}
	_, found := naming.Table[name]
	return found
}

// Lookup retorna o clientproxy do registro de nomes
func (naming NamingService) Lookup(name string) clientproxy.ClientProxy {
	value, _ := naming.Table[name]
	return value
}

// List retorna o map do serviço de nomes
func (naming NamingService) List() map[string]clientproxy.ClientProxy {
	return naming.Table
}
