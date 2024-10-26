# How

`how` is a simple command-line tool that gives you direct answers to your questions by querying OpenAI. Just ask a question from the terminal, and `how` will provide the exact command or solution you need.

## Installation

```bash
go install github.com/rymdlego/how@latest
```

Export your OpenAI API key:

```bash
export OPENAI_API_KEY=your_openai_api_key
```

## Usage

Simply type `how` followed by your question in natural language:

```bash
how how to install brew on macos
```

... the response would typically be:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Or ask something like:

```bash
how how to create an ed25519 SSH key with a comment
```

... and the reponse would be:

```bash
ssh-keygen -t ed25519 -C "example@example.com"

```

`how` will respond with only the necessary command, keeping it concise and to the point.

## License

This project is licensed under the MIT License.
