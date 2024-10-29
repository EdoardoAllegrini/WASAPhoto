<a id="readme-top"></a>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/EdoardoAllegrini/WASAPhoto">
    <img src="https://i.postimg.cc/cHJKh58L/images-1.jpg" alt="Logo" width="80">
  </a>

  <h3 align="center">WASAPhoto</h3>

  <p align="center">
    The social media I developed
    <br />
    <br />
    ·
    <a href="https://github.com/EdoardoAllegrini/WASAPhoto/issues">Report Bug</a>
    ·
    <a href="https://github.com/EdoardoAllegrini/WASAPhoto/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details open>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#functional-design-specifications">Functional design specifications</a></li>
        <li><a href="#built-with">Built With</a></li>
        <li><a href="#some-snaps">Some snaps</a></li>
        <li><a href="#project-structure">Project structure</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    	<ul>
        	<li><a href="#how-to-build">How to build</a></li>
        	<li><a href="#how-to-run">How to run</a></li>
        	<li><a href="#container">Container</a></li>
		</ul>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

Keep in touch with your friends by sharing photos of special moments, thanks to WASAPhoto! You can upload your photos directly from your PC, and they will be visible to everyone following you.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Functional design specifications

Each user will be presented with a stream of photos (images) in reverse chronological order, with information about when each photo was uploaded (date and time) and how many likes and comments it has. The stream is composed by photos from “following” (other users that the user follows). Users can place (and later remove) a “like” to photos from other users. Also, users can add comments to any image (even those uploaded by themself). Only authors can remove their comments.
Users can ban other users. If user Alice bans user Eve, Eve won’t be able to see any information about Alice. Alice can decide to remove the ban at any moment.
Users will have their profiles. The personal profile page for the user shows: the user’s photos (in reverse chronological order), how many photos have been uploaded, and the user’s followers and following. Users can change their usernames, upload photos, remove photos, and follow/unfollow other users. Removal of an image will also remove likes and comments.
A user can search other user profiles via username.
A user can log in just by specifying the username. See the “Simplified login” section for details.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

<div align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white">
  <img src="https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white">
  <img src="https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white">
  <img src="https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white">
  <img src="https://img.shields.io/badge/JavaScript-323330?style=for-the-badge&logo=javascript&logoColor=F7DF1E">
  <img src="https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vue.js&logoColor=4FC08D">
  <img src="https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white">
  <img src="https://img.shields.io/badge/Node.js-43853D?style=for-the-badge&logo=node.js&logoColor=white">
  <img src="https://img.shields.io/badge/GIT-E44C30?style=for-the-badge&logo=git&logoColor=white">
</div>

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Some snaps
[![Screenshot-2023-06-27-at-18-55-38.png](https://i.postimg.cc/9MnQXmgG/Screenshot-2023-06-27-at-18-55-38.png)](https://postimg.cc/342TnTTR)
---
<img width="1440" alt="Screenshot-2023-06-27-at-18-56-26" src="https://github.com/user-attachments/assets/018977bc-4ee3-42e5-bf12-51fe81ed4550">




### Project structure

* `cmd/` contains all executables; Go programs here should only do "executable-stuff", like reading options from the CLI/env, etc.
	* `cmd/healthcheck` is an example of a daemon for checking the health of servers daemons; useful when the hypervisor is not providing HTTP readiness/liveness probes (e.g., Docker engine)
	* `cmd/webapi` contains an example of a web API server daemon
* `demo/` contains a demo config file
* `doc/` contains the documentation (usually, for APIs, this means an OpenAPI file)
* `service/` has all packages for implementing project-specific functionalities
	* `service/api` contains an example of an API server
	* `service/globaltime` contains a wrapper package for `time.Time` (useful in unit testing)
* `vendor/` is managed by Go, and contains a copy of all dependencies
* `webui/` is an example of a web frontend in Vue.js; it includes:
	* Bootstrap JavaScript framework
	* a customized version of "Bootstrap dashboard" template
	* feather icons as SVG
	* Go code for release embedding

Other project files include:
* `open-npm.sh` starts a new (temporary) container using `node:lts` image for safe web frontend development (you don't want to use `npm` in your system, do you?)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

* go
* node.js
* vue.js


## Usage

### How to build

If you're not using the WebUI, or if you don't want to embed the WebUI into the final executable, then:

```shell
go build ./cmd/webapi/
```

If you're using the WebUI and you want to embed it into the final executable:

```shell
./open-npm.sh
# (inside the NPM container)
npm run build-embed
exit
# (outside the NPM container)
go build -tags webui ./cmd/webapi/
```


### How to run
Execute backend
```shell
cd cmd/webapi
go run .
```

Execute frontend
```
./open-npm.sh
# (here you're inside the NPM container)
npm run dev
```

### Container
Build and run backend image
```
docker build -t wasa-photos-backend:latest -f Dockerfile.backend .
docker run -it --rm -p 3000:3000 wasa-photos-backend:latest
```
Build and run frontend image
```
docker build -t wasa-photos-frontend:latest -f Dockerfile.frontend .
docker run -it --rm -p 8081:80 wasa-photos-frontend:latest
```

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Edoardo Allegrini - [website](https://EdoardoAllegrini.github.io)

Project Link: [https://github.com/EdoardoAllegrini/WASAPhoto](https://github.com/EdoardoAllegrini/WASAPhoto)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
