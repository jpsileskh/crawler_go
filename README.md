# Documentação - Web Scraper em Go

## Visão Geral

Este código implementa um web scraper simples em Go utilizando a biblioteca Colly. O programa navega recursivamente por todas as páginas do domínio `eteccamargoaranha.cps.sp.gov.br`, extraindo e visitando todos os links encontrados.

## Dependências

```go
import (
    "fmt"
    "os"
    "github.com/gocolly/colly"
)
```

- **fmt**: Pacote padrão do Go para formatação e impressão de dados
- **os**: Para operações com arquivos
- **colly**: Framework para web scraping que facilita a coleta de dados de websites

## Estrutura do Código


### Criação do Arquivo de Saída

```go
file, err := os.Create("links_coletados.txt")
if err != nil {
	fmt.Println("Erro ao criar arquivo:", err)
	return
}
defer file.Close()
```

Cria um arquivo chamado `links_coletados.txt` no diretório atual e verifica se houve erro na criação, `defer file.Close()` garante que o arquivo será fechado ao final da execução
### Inicialização do Collector

```go
c := colly.NewCollector(
    colly.AllowedDomains("eteccamargoaranha.cps.sp.gov.br"),
)
```

Cria uma nova instância do Colly com restrição de domínio. O scraper só navegará por páginas dentro do domínio especificado, evitando que o programa saia para sites externos.

### Handler de Extração de Links

```go
c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    link := e.Attr("href")
    fmt.Printf("Link found: %q -> %s\n", e.Text, link)
    c.Visit(e.Request.AbsoluteURL(link))
})
```

**Função**: Captura todos os elementos `<a>` que possuem o atributo `href` (links).

**Comportamento**:
- Extrai o valor do atributo `href` do link
- Imprime o texto do link e sua URL
- Converte URLs relativas em absolutas usando `AbsoluteURL()`
- Visita recursivamente cada link encontrado através de `c.Visit()`

### Handler de Requisições

```go
c.OnRequest(func(r *colly.Request) {
    fmt.Println("Acessei o link", r.URL.String())
})
```

**Função**: Executado antes de cada requisição HTTP.

**Comportamento**: Registra no console cada URL que está sendo acessada, permitindo monitorar o progresso do scraper.

### Início da Navegação

```go
c.Visit("https://eteccamargoaranha.cps.sp.gov.br/")
```

Inicia o processo de scraping a partir da URL raiz do site.

## Fluxo de Execução

1. O programa cria um collector configurado para o domínio específico
2. Visita a página inicial do site
3. Para cada página visitada:
   - Registra o acesso (OnRequest)
   - Extrai todos os links da página (OnHTML)
   - Para cada link encontrado:
     - Imprime o texto e URL do link
     - Visita recursivamente o link
4. O processo continua até que não haja mais links novos para visitar

## Características

- **Navegação Recursiva**: Cada link encontrado é automaticamente visitado
- **Restrição de Domínio**: Apenas páginas do domínio especificado são acessadas
- **Logging**: Todas as requisições e links encontrados são registrados no console

## Casos de Uso

Este scraper é útil para:
- Mapear a estrutura de links de um site
- Descobrir todas as páginas de um domínio
- Criar índices de conteúdo
- Verificar links quebrados
- Análise de estrutura de navegação

## Observações Importantes

⚠️ **Atenção**: Este código não implementa limitação de taxa de requisições ou respeito ao arquivo `robots.txt`. Para uso em produção, considere adicionar:
- Delay entre requisições
- Respeito às diretivas do robots.txt
- Tratamento de erros
- Limite de profundidade de navegação
- Armazenamento dos dados coletados

## Como executar

Após instalar as dependências com o comando `go mod tidy` use o comando `./crawler-go` para executar no terminal.