# 🚦 Little Book of Semaphores

<div align="center">
<img width="150" height="150" src="https://i.pinimg.com/originals/e6/3f/af/e63fafc1600ddf61941ce34362704447.gif">
</div>

## ℹ️ Informações
- **Autores:** Arthur Both, Carolina Ferreira, Felipe Freitas, Gabriel Ferreira, Maria Eduarda Maia, Matheus Caçabuena.
- **Data:** 01/11/2023
- **Linguagem:** Go
- **Status:** Concluído
- **Descrição:** Implementação de uma seleção de problemas do livro "Little Book of Semaphores".
- **Repositório:** [GitHub](https://github.com/DudaWendelMaia/FPPD.git)

## 🎯 Problema
Este projeto aborda a implementação de três problemas discutidos em aula, escolhidos a partir do "Little Book of Semaphores".

### 1. Dining Philosophers (Filósofos Jantando)

- **Descrição:** Solução com um filósofo canhoto.
- **Questão:** Por que esta solução não sofre de inanição (starvation)?
- **Arquivo:** `aFilosofos.go`

### 2. Readers-Writers (Leitores e Escritores sem Starvation)

- **Descrição:** Implementação que evita starvation para leitores e escritores.
- **Capítulo:** 4.2.5.
- **Arquivo:** `dLeitoresEscritores.go`

### 3. Santa Claus (Papai Noel)

- **Descrição:** Implementação do problema de Santa Claus.
- **Capítulo:** 5.4.
- **Arquivo:** `gSantaClaus.go`

## ▶️ Como Executar
Para executar as soluções, siga os passos abaixo:

1. Certifique-se de ter o Go instalado em sua máquina.
2. Clone o repositório:
   ```sh
   git clone https://github.com/DudaWendelMaia/FPPD.git
   ```
3. Navegue até a pasta desse trabalho:
   ```sh
   cd Trabalho1
   ```
4. Execute o comando:
   ```sh
   go run .
   ```

Isso exibirá o resultado das implementações dos problemas selecionados.
