# logutils

`logutils` é uma biblioteca de utilitários para logging em Go, que utiliza o pacote [zerolog](https://github.com/rs/zerolog) para fornecer logs estruturados e configuráveis. Esta biblioteca fornece funções para logging em diferentes níveis (Info, Error, Debug, Warn, Fatal) e é facilmente configurável através de variáveis de ambiente.

## Funcionalidades

- Configuração padrão do logger com formatação JSON.
- Logs em diferentes níveis de severidade.
- Capacidade de definir o nível de log via variável de ambiente.
- Função de log crítico que encerra o programa.

## Instalação

Para instalar a biblioteca, você pode usar o comando `go get`:

```bash
go get github.com/Mona-bele/logutils-go
```

## Uso

### Inicializando o Logger

Antes de usar a biblioteca para registrar mensagens, você precisa inicializar o logger. Isso é feito chamando a função `InitLogger`:

```go
import "github.com/Mona-bele/logutils-go"

func main() {
    logutils.InitLogger()
    // Your application code here
}
```

### Logando Mensagens

Depois de inicializar o logger, você pode usar as funções `Info`, `Error`, `Debug`, `Warn` e `Fatal` para registrar mensagens em diferentes níveis de severidade.

#### Exemplo de Log de Informação

```go
import "github.com/Mona-bele/logutils-go"

func main() {
    logutils.InitLogger()

    fields := map[string]interface{}{
        "user_id": "12345",
        "action":  "login",
    }
    logutils.Info("User logged in", fields)
}
```

#### Exemplo de Log de Erro

```go
import "github.com/Mona-bele/logutils-go"

func main() {
    logutils.InitLogger()

    err := someFunctionThatMightFail()
    if err != nil {
        fields := map[string]interface{}{
            "context": "someFunction",
            "detail":  "failure reason",
        }
        logutils.Error("An error occurred", err, fields)
    }
}
```

#### Exemplo de Log de Debug

```go
import "github.com/Mona-bele/logutils-go"

func main() {
    logutils.InitLogger()

    fields := map[string]interface{}{
        "debug_info": "some internal state",
    }
    logutils.Debug("Debugging information", fields)
}
```

#### Exemplo de Log de Aviso

```go
import "github.com/Mona-bele/logutils-go"

func main() {
    logutils.InitLogger()

    fields := map[string]interface{}{
        "warning_context": "something unusual",
    }
    logutils.Warn("This is a warning", fields)
}
```

#### Exemplo de Log Crítico

```go
import "github.com/Mona-bele/logutils-go"

func main() {
    logutils.InitLogger()

    err := someCriticalFunctionThatMightFail()
    if err != nil {
        fields := map[string]interface{}{
            "critical_context": "critical failure",
        }
        logutils.Fatal("A critical error occurred", err, fields)
    }
}
```

### Variável de Ambiente

Você pode configurar o nível de log usando a variável de ambiente `LOG_LEVEL`. Os valores aceitos são `debug`, `info`, `warn`, `error`, e `fatal`.

```bash
export LOG_LEVEL=info
```

### Contribuição

Contribuições são bem-vindas! Siga estas etapas para contribuir:

1. Faça um fork deste repositório.
2. Crie uma branch para sua feature ou correção (`git checkout -b minha-feature`).
3. Faça suas alterações e commite-as (`git commit -am 'Adiciona nova feature'`).
4. Faça um push para a branch (`git push origin minha-feature`).
5. Abra um Pull Request.

Certifique-se de seguir o estilo de código existente e de incluir testes, se aplicável.

### Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.
