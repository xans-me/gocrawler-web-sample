# Golang Web Crawler Example by [Teuku Mulia Ichsan](https://github.com/xans-me) 


## 1. Development Pattern

### **Google Wire Depedency Injection**
- Documentation of Wire DI
    - Check it out: [google/wire](https://github.com/google/wire)


### **Domain Driven Development**
- app
- env
- infrastucture
- domain
  - public
  - users 
  - monitoring
  - etc

### **Repository Service Pattern / Clean Code**

- on the domain of DDD pattern has like `public` or `users` :
    - route
    - delivery / controller
    - services
    - repository


## 2. Project Infrastucture

### **HTTP Certificates**
- For accepted criteria permission, with openSSL RSA 2048

### **Postgres**
- For storage of main data in service
