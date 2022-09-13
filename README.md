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
go test -bench . -benchman
```
for the benchmarks

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
go test -bench . -benchman
```
for the benchmarks
# More information about benchmarks
![benchmark_results d02152ba](https://user-images.githubusercontent.com/56276570/190010901-b2dd2771-0e54-46e2-bbef-d5cf710b9d49.png)

