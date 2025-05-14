# Sprint 4 - Dynamic Programming

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


## Indice

1. [Instalar Gin-Gonic](#instalar-gin-gonic)
2. [Iniciar o Projeto](#iniciar-o-projeto)
3. [Pagina de Testes](#pagina-de-testes)
4. [Requisitos](#requisitos-1)


## Instalar Gin-Gonic

1. **Garanta que seu console esta na pasta Dynamic**

```bash
    cd Dynamic-Sprint4
```

2. **Instale o Pacote e suas dependencias**
```bash
    go get github.com/gin-gonic/gin@latest
```

- **Ambos comandos concatenados**
```bash
    cd Dynamic-Sprint4 && go get github.com/gin-gonic/gin@latest
```

## Iniciar o Projeto

1. **Garanta que seu console esta na pasta Dynamic**

```bash
    cd Dynamic-Sprint4
```

2. **Pre-Compile ou Rode de Forma Interpretada**

-  **Interpretada**

    ```bash
        go run .
    ```

    ---

- **Pre Compilado**

    ```bash
        go build .
    ```

    ```bash
        <nome>.exe
    ```

    ---

## Pagina de Testes

> **Para acessar a pagina de exibição e testes utilize o navegador de preferencia e acesse ao seguinte link**

1. **View [http://127.0.0.1/](http://127.0.0.1/)**

2. **API [http://127.0.0.1/rota?from=Armazém%20A&to=Loja%20E](http://127.0.0.1/rota?from=Armazém%20A&to=Loja%20E)**

> Caso tenha modificado a porta lembre-se de modifica no link tambem

3. **[http://127.0.0.1:port/](http://127.0.0.1:8080/rota?from=Armazém%20A&to=Loja%20E)**

## Requisitos

> 1. **[Dijkstra Path Finder](./dijkstra.go) Code**

## Prova de Conceito

![image](https://github.com/user-attachments/assets/d03aa329-8892-4317-a728-168c994cdbbc)

- Mapa de Endereços e "Edges" [main file](./main.go#L15)
- Quando gerei as rotas as regras foram: que nenhum ponto fique sem conexão e que nenhum ponto tenha mais de tres vertices. (nao foram levadas em consideração proximidade ou ruas nos mapas, os dados e locais foram gerados atravez de IA para facilitação, entao sao dados sem significado real e tem o unico proposito de demonstrar o funcionamento)

