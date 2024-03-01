# go-solid
golang solid study

SOLID em go lang

S (Princípio da Responsabilidade Única):

Cada tipo (Circulo, Retangulo) tem uma única responsabilidade: representar uma forma geométrica e calcular sua área.

O (Princípio da Abertura/Fechamento):

O código é aberto para extensão (novas formas podem ser adicionadas facilmente) e fechado para modificação (as implementações existentes não são alteradas ao adicionar novas formas).

L (Princípio da Substituição de Liskov):

As instâncias de Circulo e Retangulo podem ser substituídas por instâncias de Shape sem afetar a corretude do programa. A função printShapeInfo aceita qualquer tipo que implemente a interface Shape.

I (Princípio da Segregação de Interface):

A interface Shape é minimalista e específica para o que é necessário pelas funções que a utilizam. Cada forma implementa apenas os métodos necessários para sua funcionalidade.