### 
# Exercício 2 - Multithreading

### Introdução
Exercício da Pós Graduação em Desenvolvimento Avançado com Golang.

### Para executar
Clone o repositório e acesse a pasta pelo terminal.

Execute o seguinte comando:
```bash
go run main.go {CEP}
```
Por exemplo:
```bash
go run main.go 01001000
```

### Instruções do exercício
Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

- https://cdn.apicep.com/file/apicep/" + cep + ".json
- http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:
- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.
- O resultado da request deverá ser exibido no command line, bem como qual API a enviou.
- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

