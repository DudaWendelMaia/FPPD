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

ranscrevemos as Propriedades do DiMEX abaixo:

DMX1: (não-postergação e não bloqueio) se um processo solicita Entry,   decorrido algum tempo, o acesso será permitido,
ou seja ele entrega   resp*  

(na implementação, significa que se em um processo distribuído, a aplicação escreve "dmxReq[Entry]" para o módulo DIMEX,
então decorrido um tempo ele vai garantidamente escrever um "dmxResp" no canal de indicação do módulo DIMEX para a aplicação)

DMX2: (mutex) Se um processo p entregou dmxResp, nenhum outro processo entregará   dmxResp antes que p sinalize Exit;

(na implementação, significa que
se em um processo p o módulo  DIMEX entregou "dmxResp" no canal de indicação para a aplicação, então
em nenhum outro processo  q o módulo DIMEX de q  entregará "dmxResp" para sua aplicação antes que em p
a aplicação escreva "dmxReq[EXIT]"  no canal de requisição para DIMEX em p) 

* Assumimos que todo processo sinaliza Exit ao final do seu uso, e que seu uso termina.  
