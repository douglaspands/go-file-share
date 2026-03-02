# Go File Share 🚀

Um servidor de arquivos HTTP ultra-rápido, simples e cross-platform escrito em Go.

## 🎯 Motivação

A necessidade de compartilhar arquivos rapidamente em uma rede local sem a complexidade de configurar servidores FTP, SMB ou instalar dependências pesadas (como Python ou Node.js). Este projeto foca em:

- **Velocidade**: Utiliza as capacidades nativas do Go (`net/http`) para alta performance.
- **Simplicidade**: Um único binário executável sem dependências externas.
- **Cross-Platform**: Compile facilmente para Windows, Linux e macOS.
- **Funcionalidades Úteis**: Upload de arquivos, navegação recursiva e interface amigável.

## 📦 Releases

Link para download: [https://github.com/douglaspands/go-file-share/releases](https://github.com/douglaspands/go-file-share/releases)


## 🛠️ Como Compilar

O projeto inclui um `Makefile` para facilitar a compilação cruzada.

### Pré-requisitos
- [Go](https://go.dev/dl/) 1.25+ instalado.
- `make` (opcional, mas recomendado).

### Comandos de Build

| Plataforma | Comando | Arquivo Gerado |
|------------|---------|----------------|
| **Atual**  | `make build` | `file-share` (ou .exe no Windows) |
| **Linux**  | `make build/linux` | `file-share` |
| **Windows**| `make build/windows` | `file-share.exe` |
| **macOS**  | `make build/darwin` | `file-share` |

Para limpar os arquivos compilados:
```bash
make clean
```

## 🚀 Como Executar

Após compilar, você pode executar o binário diretamente.

### Uso Básico
```bash
./file-share
```
Isso iniciará o servidor na porta **8080** compartilhando o diretório atual.

### Opções e Flags

| Flag | Abrev. | Descrição | Padrão |
|------|--------|-----------|--------|
| `--port` | `-p` | Porta do servidor | `8080` |
| `--dir` | `-d` | Diretório a ser compartilhado | `~` (home) |
| `--recursive` | `-R` | Permite navegar em subpastas | `true` |

### Exemplos

Compartilhar a pasta de Downloads na porta 8080, permitindo navegação em subpastas:
```bash
./file-share server --dir ~/Downloads -p 8080 -R
```

Compartilhar apenas a pasta atual (sem acesso a subpastas):
```bash
./file-share
```

## 🌐 Acesso

Ao iniciar, o servidor exibirá os endereços de acesso:
- **Local**: `http://localhost:8080`
- **Rede**: `http://192.168.1.X:8080` (acessível por outros dispositivos na mesma rede)
