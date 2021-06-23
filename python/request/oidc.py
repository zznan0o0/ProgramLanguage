from ClientRequest import ClientRequest

clientRequest = ClientRequest()
submit_data = {
    "userName": "root",
    "password": "root",
    "returnUrl": "/connect/authorize/callback?client_id=UserCenterCode&redirect_uri=http%3A%2F%2F192.168.10.153%3A4003%2FIds4%2FIndexCodeToken&response_type=code&scope=UserWebApi%20openid%20offline_access",
}
result = clientRequest.post(
    "http://192.168.10.153:4000/Account/Login", submit_data)
print(result)
