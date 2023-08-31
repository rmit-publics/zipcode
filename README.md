# Micro Serviço em GO para consulta de endereço por CEP

## Este micro serviço é open-source, é apenas um material de estudos para melhorar o conhecimento da linguagem GO

### Tecnologias
- Docker
- Go

### O que é necessário para rodar este micro serviço

- Ter docker instalado

### Como iniciar

utilizando o comando `docker compose up `, o sistema irá iniciar em `http:://localhost:3000/address/{cep}` apenas substitua o `{cep}` por um cep válido que deseja consultar.

### Retorno

```json
{
    "Zipcode": "00000-000",
    "Logradouro": "Nome da Rua",
    "Neighborhood": "Nome do Bairro",
    "State": "Estado (SP, MG, Rj, etc ...)",
    "City": "Cidade"
}
```

### Contatos
* Autor - [Rodrigo Mariano](www.linkedin.com/in/rodrigo-mariano-05368819)