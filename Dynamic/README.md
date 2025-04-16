# Sprint 3 - Dynamic Programming

### Grupo
|Nome|RM|
|:-:|:-:|
|Diogo Julio|553837 |
|Vinicius Silva|553240|
|Victor Didoff|552965|
|Matheus Zottis|94119|
|Jonata Rafael|552939|



## Requisitos

1. [**[Golang 1.23.2]**](https://dl.google.com/go/go1.23.2.windows-amd64.msi) or [download other](https://go.dev/dl/)
2. [**Gin-Gonic**](#instalar-gin-gonic)
3. [**Config File**](./config.json)

## Instalar Gin-Gonic

1. **Garanta que seu console esta na pasta Dynamic**

```bash
    cd Dynamic
```

2. **Instale o Pacote e suas dependencias**
```bash
    go get github.com/gin-gonic/gin@latest
```

- **Ambos comandos concatenados**
```bash
    cd Dynamic && go get github.com/gin-gonic/gin@latest
```

## Modelo de arquivo de configuração

1. **Exemplo do modelo do arquivo**

```json
{
  "version": "1.0.0",
  "ambient": {
    "name" : "Development",
    "description" : "Development environment",
    "host" : "localhost",
    "port" : 80,
    "protocol" : "http"
  },
  "database": {
    "provider" : "",
    "host" : "",
    "port" : 0,
    "service_name" : "",
    "username" : "",
    "password" : "",
    "schema" : "",
    "connection_string" : ""
  },
  "logging": {
    "level" : "*",
    "file" : "logs/",
    "mimetype" : "text/plain",
    "extension" : ".log"
  },
  "template" : "./website/html/",
  "static" : "./website/static/"
}
```

## Iniciar o Projeto

1. **Garanta que seu console esta na pasta Dynamic**

```bash
    cd Dynamic
```

2. **Pre-Compile ou Rode de Forma Interpretada**

-  **Interpretada**

    ```bash
        go run . --config "caminho/para/o/arquivo/de/configuração"
    ```

    ---

- **Interpretada Simplificado**

    > OBS : "O arquivo `config.json` tem que estar na mesma pasta!" 

    ```bash
        go run .
    ```

    ---

- **Pre Compilado**

    ```bash
        go build .
    ```

    > OBS : "O caminho pode ser full-qualified-path or dynamic-path baseado na pasta atual do cmd!" 

    ```bash
        <nome>.exe --config "caminho/para/o/arquivo/de/configuração"
    ```

    ---

- **Pre Compilado Simplificado**


    ```bash
        go build .
    ```

    > OBS : "O arquivo `config.json` tem que estar na mesma pasta que o executavel!" 

    ```bash
        <nome>.exe
    ```

## Pagina de Testes

> **Para acessar a pagina de exibição e testes utilize o navegador de preferencia e acesse ao seguinte link**

1. **[http://127.0.0.1/](http://127.0.0.1/)**

> Caso tenha modificado a porta lembre-se de modifica no link tambem

2. **[http://127.0.0.1:port/](http://127.0.0.1:8080/)**