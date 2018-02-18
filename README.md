# goLang
goLang

Chapter 2:
Simple rest api (using the bolt Database)
EndPoints
1. GetUsers: http://localhost:8089/api/bolt/GetUsers
2. SaveOrUpdateUser: http://localhost:8089/api/bolt/saveUser (It wll insert or update based on ID -)
  {
"fname": "insert-testHorray5000-TrymeAgain",
"lname": "insert-testHorray5000-Tryme",
 "id" :5000 ,
"address": {
"city": "Chicago",
"zipcode": "60661",
"state": "IL",
"line1": "500 West Madison",
"line2": "addressLine2"
}
}
3. GetUser: http://localhost:8089/api/bolt/getUser/{id}
