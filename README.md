# Golang-online-store

## Description
Golang-online-store is a simple CRUD rest api, made for trying to implement regular thing for backend like database, redis, jwt auth 
you can clone this application and try ```localhost:port//ping``` to check is the instalation sukses or not.

this project contain some kind of application the first one is signup.
```signup``` . it has a body 
'''
{

    "username":"user",
    "email":"user@gmail.com",
    "password":"password"
}
'''

generate this signup enpoint will be automatically fill your account balance default. 

the next one is login. login is a gate before we can try other product endpoint. this enpoint will be generate jwt and save it into a cookie.
 and there are some enpint else that in here that is 
 ```
{	
	- Method : GET,
- path : "/logout"
}
{	
	- Method : GET,
- path : "/getproductbycategory"
- query-param: category
}
{	
	- Method : GET,
- path : "/addcart"
- query-param: cartid
}
{
	- Method : GET,
- path : "/getcart"
}
{
	- Method : GET,
- path : "/delete"
}
{
	- Method : GET,
- path : "/checkout"
}


    
