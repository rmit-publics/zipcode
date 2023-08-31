# Imagem base
FROM golang

# diretoria um diretorio de trabalho
WORKDIR /app/src/address

# aponta a variavel gopath do go para o diretorio app
ENV GOPATH=/app

# copia os arquivos do projeto para o workdir do container
COPY . /app/src/address

# execulta o main.go e baixa as dependencias do projeto
RUN go build /app/src/address/cmd/address/main.go

# Comando para rodar o executavel
ENTRYPOINT ["./main"]

# expõe a pota 8080
EXPOSE 8080