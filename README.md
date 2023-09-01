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
go run main.go 01001-000
```

### Observações
Durante os testes notou-se que ViaCEP responde mais rápido na maioria dos casos. Para testar o retorno da APICEP foi necessário adicionar um `time.Sleep()` na chamada para ViaCEP, que está comentado.

Além disso, notou-se que a APICEP não funciona com todos os CEPs enviados, porém foi possível testar utilizando o CEP indicado na documentação da própria API:
```bash
go run main.go 06233-030
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

