# CESapi
### This is a Currency Exchange Service API that uses dependency injection design.

## RUN
* docker build -t 'cesapi' .
* docker run --rm -d -p 8080:8080 cesapi
* Open the website http://localhost:8080 
* http://localhost:8080/convert?source=TWD&target=JPY&amount=1000

![alt text](https://github.com/helgesander02/CESapi/blob/main/img/input.png)
![alt text](https://github.com/helgesander02/CESapi/blob/main/img/output.png)