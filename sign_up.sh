curl -X POST -H 'content-type:application/json' http://localhost:8080/api/users/sign-up -d '
{
    "email":"vas.shopuk@gmail.com", 
    "password":"1111",
    "name":"vasyl"
}
' | json_pp