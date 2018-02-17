


### Chapter 1: 
	Simple rest api (using the json/mux libraries- accepting application/json requestBody &response where ever required)
##### EndPoints 
1. GetUsers: http://localhost:8089/api/GetUsers  
2. DeleteUsers: http://localhost:8089/api/updateUser/{id} 
3. GetUser: http://localhost:8089/api/getUser/{id} 
4. CreateUser: http://localhost:8089/api/createUser -(For UniqueId: using the math.randon api 
		Sample Post for create
		{
		"fname": "update-testHorray-TrymeAgain",
		"lname": "update-testHorray1212-Tryme",
		"address": {
		"city": "Chicago",
		"zipcode": "60661",
		"state": "IL",
		"line1": "500 West Madison",
		"line2": "addressLine2"
		}
		})
5. DeleteUser: http://localhost:8089/api/removeUser/{id}
