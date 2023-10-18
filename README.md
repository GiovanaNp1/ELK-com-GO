# ELK-com-GO

![Go presentation](https://gophersource.com/img/build.png)


## Sobre
Projeto feito para demonstração de Execução de APM com ELK(Elastic, Logstash e Kibana) na palestra sobre Go com observabilidade

Palestra realizada no Golang SP na Cloudwalk: Mais ativos do que nunca: [https://thedevconf.com/tdc/2022/business/trilha-devops](https://www.meetup.com/pt-BR/golangbr/events/296500714/)

## Versões utilizadas 💘

Versão do Go: 1.21.3

Versão do Docker Desktop Server: Docker Desktop 4.5.0 (74594)


## Passo a passo de como utilizar no SpringBoot 💡💎

1. Acesse o site https://cloud.mongodb.com/ e crie sua conta por lá
2. Acessando sua conta crie um novo projeto e deixe o nome do seu projeto(pode ser qualquer um) e em add membros deixe como está no caso ele vai deixar você seu usuario de cadastro como owner
3. depois de criar clique na aba Network Access e clicando nele clica em Add IP Address e seleciona ALLOW ACCESS FROM ANYWHERE e clica em confirm
4. depois de criar clique na aba clica em Add New Database User e em Password Authentication coloque um user e password que desejar
5. Clica na aba de DataBase clica em Create e seleciona o banco share, e clica em confirm, ele demora um pouco para carregar mas logo ja estará criado seu banco
6. Clica em Connect e logo em seguida seleciona connect your application e nele você  receberá na parte Add your connection string into your application code a url do nosso banco, copia e cola isso dentro do seu .env

## Passo a passo de como utilizar no Docker 💡🐳

1. Verifique que está com seu docker ligado no seu computador com o docker desktop
2. Para elevar o container basta colcocar o comando:  docker compose up
3. Para desativar os container: docker compose down

## Tecnologias usadas nesse projeto projeto 💻

- 🍃 [MongoDB](https://www.mongodb.com/pt-br)
- 🐹 [GO](https://go.dev/)
- 🐳 [Docker](https://www.docker.com/)
- 🦌 [ELK](https://www.elastic.co/pt/what-is/elk-stack)
- 💌 [Postman](https://www.postman.com/)

## Artigos recomendados

https://www.elastic.co/pt/blog/how-to-instrument-your-go-app-with-the-elastic-apm-go-agent
