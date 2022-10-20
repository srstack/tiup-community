# tiup-community
tiup community 


## init tiup mirror 
1. create TIUP_KEY 
```
# creare rsa key peer
openssl genrsa -out rsa.key 2048 
```
2. SET Github Actions secrets
Store the resulting `rsa.key` in Github Actions secrets, named `TIUP_KEY`

3. RUN Init Mirror
Run Init-Mirror Github Action Workflow