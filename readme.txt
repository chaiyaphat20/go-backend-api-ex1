1.ลง gin lib
go get -u github.com/gin-gonic/gin

2.วิธีรัน go run .

3.hot reload
go install github.com/cosmtrek/air@latest
3.1 nano ~/.zshrc
3.2 export PATH=$PATH:~/go/bin
3.3 source ~/.zshrc
3.4 พิมพ์ air ในโปรเจค


4.gorm
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

5.Declaring Models
https://gorm.io/docs/models.html#Fields-Tags

6.model binding
https://gin-gonic.com/en/docs/examples/binding-and-validation/

7.gorm orm query
https://gorm.io/docs/query.html

8.scope
https://gorm.io/docs/scopes.html

9.jwt
go get -u github.com/golang-jwt/jwt/v5

10. cors 
https://github.com/gin-contrib/cors