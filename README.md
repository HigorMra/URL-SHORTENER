Go URL Shortener (Serviço de Encurtamento de URL)
Este projeto é um serviço de encurtamento de URL minimalista construído em Go, que demonstra a capacidade da linguagem de criar APIs e microsserviços rápidos e de alta concorrência.

Destaques do Projeto
API REST Simples: Implementa dois endpoints HTTP essenciais: um para criar URLs curtas e outro para realizar redirecionamento.

Performance (Go): Utiliza o pacote padrão net/http do Go, o que resulta em um servidor web leve e rápido, ideal para lidar com muitas requisições de redirecionamento simultâneas.

Estrutura de Módulos: Segue as melhores práticas de Go, utilizando módulos (go.mod) e separando a lógica de rotas/servidor (em main.go) do core do negócio (no pacote handler).

Concorrência Implícita: O net/http do Go lida com cada requisição em sua própria goroutine (thread leve), garantindo que o serviço possa escalar eficientemente.
