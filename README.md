# CLI NameSpace (CNS)

CLI NameSpace (CNS) is a GoLang-based CLI application designed to manage command sessions for users. It allows users to store, retrieve, and execute commands within named sessions, enhancing the command-line experience by organizing commands in a user-friendly manner. CNS is perfect for users who frequently work with long or complex commands and want to save them for future use.

## Features

- **Session Management:** Users can create, list, and switch between sessions to manage different command sets.
- **Command Storage:** Commands are stored within sessions, allowing users to retrieve and execute them later easily.
- **Ease of Use:** CNS provides a straightforward CLI for managing commands and sessions without clutter.

## Getting Started

### Prerequisites

- GoLang installed on your machine.
- Basic knowledge of CLI operations.

### Installation

1. **Compile the Application:**

   Before using CNS, compile the application using the provided Makefile:

```shell
make compile
```

2. **Install the Application:**

After compilation, install CNS to make it available in your system:

```shell
cns install
```

### Usage

CNS operations are simple and designed to be intuitive for the user:

1. **Start a New Session or List Sessions:**

To start a new session, use:

```shell
cns start __session_name__
```

If the session name is not provided, CNS returns a list of existing sessions:

```shell
cns start
```

Alternatively, to create a new session explicitly:
```shell
cns create __session_name__
```

2. **List Commands:**

To list all pinned commands within the current session:

```shell
cns list
```

3. **Execute Commands:**

- To execute a specific command by its ID:

```shell
cns __id__
```

- To add and execute a new command, giving it a new ID automatically:

```shell
cns command args
```

- To run a command with specific args and assign a custom identifier:

```shell
cns -i command_name command args
```

4. **Delete Commands or Sessions:**

- To delete a command by its ID:

```shell
cns rm __id__
```

- To stop and delete a session:

```shell
cns stop
cns delete __session_name__
```

## Contributing

Contributions to CLI NameSpace are welcome! Please submit pull requests or create issues for bugs and feature requests.

## License

CLI NameSpace is open-source software licensed under the MIT license.
