# Digiutilsapi
>
> A simple RESTFUL API that provides many handy tools to make your life much easier.

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

This simple app is the implementation of my [digiutils](https://github.com/Digisata/digiutilsapi) library, and also one form of [#100DaysOfCode](https://www.100daysofcode.com/) movement 🏃‍♂️. There is a [docker image](https://hub.docker.com/r/digisata/digiutilsapi) 🐋 version as well and if you find it useful please give it a star ⭐ or even do your [contribution](#contributing). Cheers 🥂

## Live Demo

You can check the live demo [here](https://empty-impala-87.a276.dcdg.xyz/swagger/index.html).

## Prerequisites

- `Go 1.19.3` or higher.
- [Swaggo](https://github.com/swaggo/swag).

## Installation

- Clone this repository.
- Rename [.env.example](.env.example) file to `.env` and change the `PORT` value with your desired port.
- Then run this command to synchronize all the dependencies.
```shell
go mod tidy
```
- After that run the command below to generate the required files for swaggo.
```shell
swag init
```
- Finally, run the project.
```shell
go run .
```

## Usage

Open your browser and go to `http://localhost:3000/swagger/index.html`.
> **_NOTE:_**  Replace the 3000 with your defined `PORT` value in your `.env` file.

## Meta

Hanif Naufal – [@Digisata](https://twitter.com/Digisata) – [hnaufal123@gmail.com](mailto:hnaufal123@gmail.com)

Distributed under the MIT license. See [LICENSE](LICENSE.md) for more information.

## Contributing

1. Fork this repository.
2. Create your own branch (`git checkout -b fooBar`).
3. Commit your changes (`git commit -am 'Add some fooBar'`).
4. Push to the branch (`git push origin fooBar`).
5. Create your awesome Pull Request.
