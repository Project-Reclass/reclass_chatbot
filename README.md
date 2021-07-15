# Chatbot
A chat application to eventually demonstrate the usefulness of Kubernetes
## About

This chatbot posts to a Chatback application API (see Cloning & Setup below) periodically. The username, message, and the posting intervals can be specified using CLI commands (see CLI Commands).

## Getting Started

### Dependencies

- [Go required version 1.16](https://golang.org/doc/install)
- [Git](https://git-scm.com/downloads)
- [golangci-lint (Optional)](https://golangci-lint.run/usage/install/#local-installation) - Used for running lint on your development machine.
- [Minikube](https://minikube.sigs.k8s.io/docs/start/)
- [Docker] (https://docs.docker.com/engine/install/) - recommended, can use other container or VM

### Cloning & Setup

```bash
git clone https://github.com/Project-Reclass/reclass_chatbot.git
minikube start
minikube kubectl -- apply -f https://raw.githubusercontent.com/scottjr632/chatback-k8s-example/main/chatback-remote.yml
minikube service frontend -n chatback
minikube service backend -n chatback
```

### Running

```bash
go run main.go
```

### Testing

```bash
go test ./...
```

### CLI Commands
When main.go is run, a new post is created with default parameters. The username is "Reclass Bot," the message is the current time, and the API is posted to every 3 seconds. Note that the time printed in the body of the post and the time written in the post details may be off by a fraction of a second.

The parameters, however, can be specified with CLI commands. The CLI commands are now as follows:
*   `-username="..."` // Sets the username of each post (string)
*   `-message="..."` // Sets the message of each post (string)
*   `-interval=5` // Sets the wait time for each post OR the upper range for wait times if 'random' is selected (int)
*   `-random` // Determines how interval is used (boolean)


## GitHub Actions

This project has 3 CI jobs that run on every push. This jobs can be found in the [.github/workflows/ci.yml](.github/workflows/ci.yml) file.

- **Test**: The test jobs run `go test ./...` to run every unit test in the codebase. 

- **Vet**: The Vet job runs the `go vet ./...` on the codebase. The `go vet` command is used to report suspicious constructs and smelly code. More about `vet` can be read on the [golang.org/cmd/vet](https://golang.org/cmd/vet/) webpage.

- **Lint**: The Lint job, like the `vet` command, also checks for [smelly code](https://en.wikipedia.org/wiki/Code_smell) but runs more checks (e.g. are all errors handled?). The Lint job will also provide [annotations on pull requests](https://github.com/Project-Reclass/pup/commit/8f62d4be715d369f95745eeba1df996f3e8afeea#diff-2873f79a86c0d8b3335cd7731b0ecf7dd4301eb19a82ef7a1cba7589b5252261R9) for issues related to new commits.

These CI jobs help to ensure that only good quality code makes it into the codebase. At Project Reclass, each of our code bases has some form of CI setup so being familiar with these is helpful.


