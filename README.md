### Player
to launch the API you need to move to the folder player
```
cd player
```
then
```
go run .
```
>**in  order that the api work you need to launch the blockchain service to**



### Blockchain service
same as api you need to move to the folder blockchain service in order to launch it
>you need new terminal
```
cd player
```
then
```
go run .
```
>**in  order that the api work you need to launch the blockchain service to**

***The only method supported is POST***
the addres of the post
```
http://localhost:8090/ethereum/wallets/create/
```
Example of body to provide :
```json
{
    "Username" : "test",
    "Password" : "password_",
    "PinCode" : "123456"
}
```
This endpoint will successfully respond with a body providing the name and the wallet address 
```json
{
    "WalletAddress":"0xL82RLB8YCJHVSZSGX5HYB679K4A81O5DQFCQ6C2FV9URL",
    "CurrencyBalance":"0"
    "CurrencyCode":"ETH"
}
```
#Benchmarks and tests
**to test and benchemark the service you need to follow the commands**
```
cd heandlers
```
then 
```
go test -v
```
for tests

or 
```
go test -bench .
```
for the benchmarks

**to test and benchemark the api you need to follow the commands**
```
cd heandlers
```
then 
```
go test -v
```
for tests

or 
```
go test -bench .
```
for the benchmarks

# More information about benchmarks
![benchmark_results d02152ba](https://user-images.githubusercontent.com/56276570/190010901-b2dd2771-0e54-46e2-bbef-d5cf710b9d49.png)

# Docker
```
docker-compose up
```
This will provide : 
 - An api exposing its endpoint to create a player and a wallet
 - A mock blockchain api to simulate the creation of a wallet on a blockchain which will be called by the other api

- [x] docker not working for the moment :frowning_face:
![image](https://user-images.githubusercontent.com/56276570/190148060-87eacb8b-592e-4e47-a5a6-61b0ae74839a.png)



