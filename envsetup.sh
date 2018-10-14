
go get golang.org/x/crypto/bcrypt
go get github.com/gorilla/context
go get github.com/gorilla/mux
go get github.com/dgrijalva/jwt-go
go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson

# Start the mongo service 
sudo service mongod start
mongo localhost:27017/test mongo_setup.js