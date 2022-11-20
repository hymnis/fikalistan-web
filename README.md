## Fikalistan

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GoT](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev)

<p align="center"><img alt="Logo" src="https://user-images.githubusercontent.com/xxx.png" height="200px"/></p>

## Table of Contents

- [Gettings started](#gettings-started)
    - [Dependencies](#dependencies)
    - [Make commands](#make-commands)
- [Credits](#credits)


## Getting started

### Dependencies

- [Go](https://go.dev)
- [Docker](https://docker.com)
- [Docker Compose](https://docs.docker.com/compose/install)

### Make commands

- `make up`: Setup service containers
- `make run`: Run application container
- `make reset`: Re-create containers
- `make test`: Run all tests, in correct order
- `make db`: Connects to the primary database
- `make db-test`: Connects to the test database
- `make cache`: Connects to the primary cache
- `make cache-test`: Connects to the test cache


## Credits

- [Pagoda](https://github.com/mikestefanello/pagoda): A great starter blueprint for your full-stack web development projects
