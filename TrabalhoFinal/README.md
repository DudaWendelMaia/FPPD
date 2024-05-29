#  Sistemas Distribuídos
<div align="center">
<img width="250" height="150" src="https://media.tenor.com/iOc4hqMDGuEAAAAj/bubu-dudu-love.gif">
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




# Trabalho Final
Nesta pasta, você encontrará os elementos relacionados ao Trabalho Final. Isso engloba o desenvolvimento, documentação e outros artefatos essenciais para a conclusão do projeto final.

## Sistema Paralelo
- Link relatório: https://docs.google.com/document/d/1hZwiRl1R6qfaZnRajG81AuIG6c954upsHb9wEF3NbPY/edit?usp=sharing
- Para rodar a pasta com o sistema paralelo tem que abrir o terminal e digitar "go run nomeArquivo.go"

## Sistema Distribuido
- Link relatório: https://docs.google.com/document/d/17DKEDn-sbCOc2t-X6sATWdOPyLLwz9Xp4i4_NPJq6hc/edit?usp=sharing
- Para rodar a parte do sistema distribuído tem que entar na pasta usedimex-f, copiar os IPs abrir, 3 terminais, colar nos três e rodar, isso criará um novo arquivo que mostra a entrada e saída dos processos (sendo as "|" as entradas e os "." as saídas)


Trabalho SD
  Encontre na pasta de programas abaixo os fontes do template para o exercício de Sistemas Distribuídos.
  Conforme visto em aula:
        DIMEX é um template para implementar o algoritmo de Exclusão Mútua Distribuída
        useDIMEX.go e useDIMEX-f.go são aplicações que representam os processos distribuído.
  Voce deve iniciar, em diferentes Shell (terminais) na sua máquina, os processos useDimex.go. (ou -f)
  Voce pode rodar em diferentes maquinas bastando acertar os enderecos IP utilizados.
  Leia nos comentários iniciais do arquivo como fazer isto.

  O useDIMEX.go implementa processos que ficam em loop eterno 
        pedindo ao DIMEX a exclusao mutua,  (ENTRY)
        aguarda ate o DIMEX permitir,       (aguarda indicação do DIMEX)
        usa o recurso (cédigo vazio)  
        informa ao DIMEX que liberou        (EXIT)

  O useDIMEX-f.go implementa processos que ficam em loop eterno 
        pedindo ao DIMEX a exclusao mutua,  (ENTRY)
        aguarda ate o DIMEX permitir,       (aguarda indicação do DIMEX)
               abre um mesmo arquivo , 
               esceve nele,  
               fecha
        informa ao DIMEX que liberou        (EXIT)
** Em grupos de até 6, vocês devem:
    A) implementar o núcleo do algoritmo DIMEX para respeitar as propriedades
        da exclusão mútua vistas em aula, que estão nos slides;
        fazer rodar.  testar.
    B) entregar o codigo e um relatorio sucinto em pdf na sala moodle
       o relatório deve conter a argumentação de porque o sistema funciona.
       ou seja, usando as propriedades de P2PPLink e dado o algoritmo DIMEX,
       por que as propriedades (veja nos slides) do DiMEX são mantidas.

Transcrevemos as Propriedades do DiMEX abaixo:

DMX1: (não-postergação e não bloqueio) se um processo solicita Entry,   decorrido algum tempo, o acesso será permitido,
ou seja ele entrega   resp*  

(na implementação, significa que se em um processo distribuído, a aplicação escreve "dmxReq[Entry]" para o módulo DIMEX,
então decorrido um tempo ele vai garantidamente escrever um "dmxResp" no canal de indicação do módulo DIMEX para a aplicação)

DMX2: (mutex) Se um processo p entregou dmxResp, nenhum outro processo entregará   dmxResp antes que p sinalize Exit;


3. **Trabalho Final:**
   - Acesse a pasta "Trabalho Final" para explorar os componentes do projeto final.
   - Encontre a documentação abrangente e os recursos essenciais para compreender e reproduzir o trabalho desenvolvido.
   - Para rodar a pasta com o sistema paralelo tem que abrir o terminal e digitar "go run nomeArquivo.go" 
   - Para rodar a parte do sistema distribuído tem que entar na pasta usedimex-f, copiar os IPs abrir, 3 terminais, colar nos três e rodar, isso criará um novo arquivo que mostra a entrada e saída dos processos (sendo as "|" as entradas e os "." as saídas)

(na implementação, significa que
se em um processo p o módulo  DIMEX entregou "dmxResp" no canal de indicação para a aplicação, então
em nenhum outro processo  q o módulo DIMEX de q  entregará "dmxResp" para sua aplicação antes que em p
a aplicação escreva "dmxReq[EXIT]"  no canal de requisição para DIMEX em p) 

* Assumimos que todo processo sinaliza Exit ao final do seu uso, e que seu uso termina.  
