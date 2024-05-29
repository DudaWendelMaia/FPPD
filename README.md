# Sistemas Distribuídos

<div align="center">
  <img width="200" height="250" src="https://meneguite.com/2017/10/01/golang-desbravando-uma-linguagem-de-programacao-parte-1/001.gif">
</div>

Esta pasta contém os elementos relacionados ao Trabalho Final de Sistemas Distribuídos, incluindo desenvolvimento, documentação e outros artefatos essenciais para a conclusão do projeto.

## ℹ️ Informações
- **Autores:** Arthur Both, Carolina Ferreira, Felipe Freitas, Gabriel Ferreira, Maria Eduarda Maia, Matheus Caçabuena.
- **Data:** 01/11/2023
- **Linguagem:** Go
- **Status:** Concluído
- **Descrição:** implementação DIMEX para sistemas distribuídos.
- **Repositório:** [GitHub](https://github.com/DudaWendelMaia/FPPD.git)

## 🎯 Problema
O trabalho aborda a implementação de um algoritmo de Exclusão Mútua Distribuída (DIMEX) para sistemas distribuídos. Os detalhes do problema e suas soluções estão descritos nos arquivos `useDIMEX.go` e `useDIMEX-f.go`. O objetivo é garantir que os processos distribuídos operem de forma segura e eficiente, respeitando as propriedades de exclusão mútua.

### Propriedades do DiMEX
Transcrevemos as principais propriedades do algoritmo DiMEX abaixo:
- **DMX1:** Garante não-postergação e não-bloqueio.
- **DMX2:** Assegura a exclusão mútua.

## ▶️ Como Executar
Para testar e executar as soluções, siga os passos abaixo:

1. **Instalação do Go:**
   Certifique-se de ter o Go instalado em sua máquina.

2. **Clonagem do Repositório:**
   ```sh
   git clone https://github.com/DudaWendelMaia/FPPD.git
   ```
   
3. **Navegação até a Pasta do Trabalho:**
   ```sh
   cd TrabalhoFinal
   ```
   
### Sistema Paralelo
Para rodar o sistema paralelo, execute o seguinte comando no terminal:
```sh
go run nomeArquivo.go
```
Substitua `nomeArquivo.go` pelo nome do arquivo que deseja executar.

- **Link para o Relatório:** [Google Docs](https://docs.google.com/document/d/1hZwiRl1R6qfaZnRajG81AuIG6c954upsHb9wEF3NbPY/edit?usp=sharing)

### Sistema Distribuído
Para executar o sistema distribuído, siga os passos abaixo:

1. **Configuração dos IPs:**
   Entre na pasta `usedimex-f` e ajuste os IPs necessários nos arquivos de configuração.

2. **Execução em Terminais Diferentes:**
   Abra três terminais e execute o comando para cada um deles, conforme configurado.

- **Link para o Relatório:** [Google Docs](https://docs.google.com/document/d/17DKEDn-sbCOc2t-X6sATWdOPyLLwz9Xp4i4_NPJq6hc/edit?usp=sharing)

Isso exibirá o resultado das implementações dos problemas selecionados, demonstrando a funcionalidade do algoritmo DiMEX em ambientes paralelos e distribuídos.
