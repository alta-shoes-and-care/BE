<div id="top"></div>

<div>
    <!-- Project Logo -->
    <div align="center">
        <a href="https://images.unsplash.com/photo-1580902215262-9b941bc6eab3">
            <img src="https://images.unsplash.com/photo-1580902215262-9b941bc6eab3" alt="logo" width="400">
        </a>
        <h3 align="center">
            Washing Shoes Platform
        </h3>
    </div>
</div>

# Shoes and Care
<!-- Project Description -->
<div>
    <p style="text-align:left">
        Nowadays, some people are very busy and don't even have time to clean their shoes.
        Therefore we need a platform that is able to connect customers with shoes washing service provider.
    </p>
</div>

# Table of Contents
- [Shoes and Care](#shoes-and-care)
- [Table of Contents](#table-of-contents)
- [High Level Architecture](#high-level-architecture)
- [Technology Stack](#technology-stack)
  - [Framework](#framework)
  - [Packages](#packages)
  - [Database](#database)
  - [Deployments](#deployments)
  - [Collaboration](#collaboration)
- [Project Structure](#project-structure)
- [Unit Test](#unit-test)
- [How to Contribute](#how-to-contribute)
- [Contacts](#contacts)

<p style="text-align:right">(<a href="#top">back to top</a>)</p>

# High Level Architecture
<p style="text-align:right">(<a href="#top">back to top</a>)</p>

# Technology Stack
## Framework
- [Echo (Go Web Framework)](https://echo.labstack.com/)

## Packages
- [GORM (Golang ORM Library)](https://gorm.io/)
- [Testify (Unit Test)](https://pkg.go.dev/github.com/stretchr/testify)
- [Midtrans (Payment Gateway)](https://midtrans.com/)
- [AWS SDK V1 (AWS Software Development Kit for Go)](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/welcome.html)

## Database
- [MySQL (Open Source Database)](https://www.mysql.com/)

## Deployments
- [AWS S3 (Cloud Storage)](https://aws.amazon.com/s3)
- [Docker (Image Container)](https://www.docker.com/)
- [Kubernetes (Container Orchestration)](https://kubernetes.io/)
- [Okteto (Kubernetes Platform)](https://www.okteto.com/)

## Collaboration
- [Trello (Work Management Tool)](https://trello.com/)
- [GitHub (Version Control System Platform)](https://github.com/)
<p style="text-align:right">(<a href="#top">back to top</a>)</p>

# Project Structure
```
BE
├── configs
│   └── config.go
├── deliveries
│   ├── controllers
│   │   ├── common
│   │   │   └── common.go
│   │   └── user
│   ├── middlewares
│   │   ├── jwtAuth.go
│   │   └── jwtMiddleware.go
│   └── routes
│       └── route.go
├── entities
│   └── user
├── ERD
├── OpenAPI
│   └── openapi.yaml
├── repositories
│   ├── auth
│   ├── hash
│   └── user
├── services
│   ├── aws-s3
│   │   └── aws-s3.go
│   └── midtrans-pay
│       └── midtrans-pay.go
├── utils
│   └── mysqldriver.go
├── .env
├── .gitignore
├── docker-compose.yaml
├── dockerfile
├── go.mod
├── go.sum
└── README.md
```
<p style="text-align:right">(<a href="#top">back to top</a>)</p>

# Unit Test
<p style="text-align:right">(<a href="#top">back to top</a>)</p>

# How to Contribute
<p style="text-align:right">(<a href="#top">back to top</a>)</p>

# Contacts
- Yusuf Nur Wahid | [GitHub](https://github.com/ynwahid) • [LinkedIn](https://www.linked.com/in/ynwahid)
- Frans Ihsan | [GitHub](https://github.com/fransihsan) • [LinkedIn](https://www.linkedin.com/in/fransihsan/)
<p style="text-align:right">(<a href="#top">back to top</a>)</p>