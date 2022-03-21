<div id="top"></div>

<div>
    <!-- Project Logo -->
    <div align="center">
        <a href="images/shoes-service-station-cancel.png">
            <img src="images/shoes-service-station-cancel.png" alt="Shoes Service Station Logo" width="400">
        </a>
        <h3 align="center">
            Shoes Service Station
        </h3>
    </div>
</div>

# Technology Stack
![GitHub](https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white)
![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)
![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Go Echo](https://img.shields.io/badge/-Echo-4CE1FF?logo=go&logoColor=white&style=for-the-badge)
![GORM](https://img.shields.io/badge/-GORM-56A6EE?logo=go&logoColor=white&style=for-the-badge)
![Testify](https://img.shields.io/badge/Testify-blue?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/static/v1?style=for-the-badge&message=MySQL&color=4479A1&logo=MySQL&logoColor=FFFFFF&label=)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Kubernetes](https://img.shields.io/badge/kubernetes-%23326ce5.svg?style=for-the-badge&logo=kubernetes&logoColor=white)
![Amazon S3](https://img.shields.io/static/v1?style=for-the-badge&message=Amazon+S3&color=569A31&logo=Amazon+S3&logoColor=FFFFFF&label=)
![Amazon RDS](https://img.shields.io/badge/Amazon%20RDS-4053D6?style=for-the-badge&logo=Amazon%20DynamoDB&logoColor=white)
![Midtrans](https://img.shields.io/badge/-midtrans-0A2955?style=for-the-badge)
![Okteto](https://img.shields.io/badge/-Okteto-1E222B?style=for-the-badge)
![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)
![Trello](https://img.shields.io/badge/Trello-%23026AA7.svg?style=for-the-badge&logo=Trello&logoColor=white)
![Google Meet](https://img.shields.io/badge/Google%20Meet-00897B?style=for-the-badge&logo=google-meet&logoColor=white)
<p align="right">(<a href="#top">back to top</a>)</p>

# About the Project
<!-- Project Description -->
<div>
    <p style="text-align:left">
        Nowadays, some people are very busy and don't even have time to clean their shoes.
        Therefore we need a platform that is able to connect customers with shoes washing service provider.
    </p>
</div>
<p align="right">(<a href="#top">back to top</a>)</p>

# High Level Architecture
![High Level Architecture](images/HLA-updated.jpeg)
<p align="right">(<a href="#top">back to top</a>)</p>

# Project Structure
```
BE
├── configs
│   └── config.go
├── deliveries
│   ├── controllers
│   │   ├── auth
│   │   │   ├── auth_test.go
│   │   │   ├── auth.go
│   │   │   ├── request.go
│   │   │   └── response.go
│   │   ├── common
│   │   │   └── common.go
│   │   ├── order
│   │   │   ├── order_test.go
│   │   │   ├── order.go
│   │   │   ├── request.go
│   │   │   └── response.go
│   │   ├── payment-method
│   │   │   ├── payment-method_test.go
│   │   │   ├── payment-method.go
│   │   │   ├── request.go
│   │   │   └── response.go
│   │   ├── review
│   │   │   ├── request.go
│   │   │   ├── review_test.go
│   │   │   └── review.go
│   │   ├── service
│   │   │   ├── request.go
│   │   │   ├── response.go
│   │   │   ├── service_test.go
│   │   │   └── service.go
│   │   └── user
│   │       ├── request.go
│   │       ├── response.go
│   │       ├── user_test.go
│   │       └── user.go
│   ├── helpers
│   │   └── hash
│   │       └── hash.go
│   ├── middlewares
│   │   ├── jwtAuth.go
│   │   └── jwtMiddleware.go
│   ├── mocks
│   │   ├── auth
│   │   │   └── auth.go
│   │   ├── order
│   │   │   └── order.go
│   │   ├── payment-method
│   │   │   └── payment-method.go
│   │   ├── review
│   │   │   └── review.go
│   │   ├── service
│   │   │   └── service.go
│   │   └── user
│   │       └── user.go
│   ├── routes
│   │   └── route.go
│   └── validators
│       └── validator.go
├── entities
│   ├── order
│   │   └── order.go
│   ├── payment-method
│   │   └── payment-method.go
│   ├── review
│   │   └── review.go
│   ├── service
│   │   └── service.go
│   └── user
│       └── user.go
├── ERD
│   └── erd.drawio
├── external
│   ├── aws-s3
│   │   ├── aws-s3.go
│   │   └── interface.go
│   └── midtrans-pay
│       ├── interface.go
│       └── midtrans-pay.go
├── images
│   ├── HLA-updated.jpeg
│   └── shoes-service-station.png
├── OpenAPI
│   └── openapi.yaml
├── repositories
│   ├── auth
│   │   ├── auth_test.go
│   │   ├── auth.go
│   │   └── interface.go
│   ├── hash
│   │   └── hash.go
│   ├── mocks
│   │   ├── order
│   │   │   └── order.go
│   │   ├── payment-method
│   │   │   └── payment-method.go
│   │   ├── review
│   │   │   └── review.go
│   │   ├── service
│   │   │   └── service.go
│   │   └── user
│   │       └── user.go
│   ├── order
│   │   ├── formatter.go
│   │   ├── interface.go
│   │   ├── order_test.go
│   │   └── order.go
│   ├── payment-method
│   │   ├── interface.go
│   │   ├── payment-method_test.go
│   │   └── payment-method.go
│   ├── review
│   │   ├── formatter.go
│   │   ├── interface.go
│   │   ├── review_test.go
│   │   └── review.go
│   ├── service
│   │   ├── interface.go
│   │   ├── service_test.go
│   │   └── service.go
│   └── user
│       ├── interface.go
│       ├── user_test.go
│       └── user.go
├── utils
│   └── mysqldriver.go
├── .gitignore
├── app-pod.yaml
├── coverage.out
├── docker-compose.yaml
├── dockerfile
├── go.mod
├── go.sum
├── main.go
├── README.md
```
<p align="right">(<a href="#top">back to top</a>)</p>

# Unit Test
![Testing Coverage - 1](images/coverage-1.png)
![Testing Coverage - 1](images/coverage-2.png)
<p align="right">(<a href="#top">back to top</a>)</p>

# How to Contribute
<p align="right">(<a href="#top">back to top</a>)</p>

# Contacts
- [![GitHub](https://img.shields.io/badge/ynwahid-100000?style=for-the-badge&logo=github&logoColor=white)](https://github.com/ynwahid/) [![LinkedIn](https://img.shields.io/badge/ynwahid-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/ynwahid/)
- [![GitHub](https://img.shields.io/badge/fransihsan-100000?style=for-the-badge&logo=github&logoColor=white)](https://github.com/fransihsan/) [![LinkedIn](https://img.shields.io/badge/fransihsan-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/fransihsan/)
<p align="right">(<a href="#top">back to top</a>)</p>