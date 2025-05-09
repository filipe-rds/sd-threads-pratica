
# 📚 Serviço de Dicionário Distribuído com Threads em Go

Este projeto implementa um **Serviço de Dicionário Distribuído** como atividade prática da disciplina de **Programação Distribuída**, com ênfase na utilização de **threads (concorrência)**, comunicação **cliente-servidor via RPC**, e **sincronização de dados compartilhados** utilizando **mutex** em Golang.


## 🚀 Visão Geral

O sistema consiste em dois componentes principais:

- **Servidor (`server.go`)**: Responsável por armazenar e manipular um dicionário remoto (`RemoteMap`) de forma concorrente e segura.
- **Cliente (`client.go`)**: Interface de terminal que permite ao usuário interagir com o servidor, enviando operações remotas como adicionar, consultar ou remover pares chave-valor.


## 🧩 Tecnologias e Conceitos Utilizados

- **Go (Golang)** — Linguagem principal.
- **Threads / Concorrência** — Uso de `sync.Mutex` para garantir segurança em acessos simultâneos.
- **RPC (Remote Procedure Call)** — Comunicação remota usando a biblioteca [`hprose-go`](https://github.com/hprose/hprose-go).
- **Arquitetura Cliente-Servidor** — Separação clara entre as responsabilidades de cada componente.



## 🛠️ Funcionalidades

- ✅ Inserção/Atualização de valores (`Update`)
- ✅ Consulta de valores (`Get`)
- ✅ Remoção de valores (`Remove`)
- ✅ Interface CLI com menu interativo
- ✅ Limpeza de tela para melhor experiência do usuário

## ▶️ Como Executar

### Pré-requisitos

- [Go instalado](https://golang.org/doc/install)
- Conexão local via `localhost:8080`

### Passos:

1. **Instale as dependências**:

```bash
go mod tidy
````

2. **Inicie o servidor**:

```bash
go run server.go
```

3. **Em outro terminal, inicie o cliente**:

```bash
go run client.go
```


## 📋 Exemplo de Uso no Cliente

```bash
Escolha uma operação:
1. Adicionar/atualizar par no dicionário (Update)
2. Obter valor de uma chave do dicionário (Get)
3. Remover par no dicionário através da chave (Remove)
4. Sair

Digite sua opção: 1

Digite a chave: nome
Digite o valor: 42
Par inserido com sucesso!

Pressione qualquer tecla para continuar...
```


## 📚 Aprendizados e Objetivos da Atividade

Esta atividade prática aborda os seguintes conceitos-chave da disciplina:

* Arquitetura de sistemas distribuídos
* Comunicação entre processos remotos
* Gerenciamento de concorrência e sincronização
* Experiência de desenvolvimento com Go e RPC



