# tiup-community
tiup community 


## init tiup mirror 
1. create TIUP_KEY 
```
# creare rsa key peer
openssl rand -base64 32 > tiup.key
```
2. SET Github Actions secrets
Store the resulting `tiup.key` in Github Actions secrets, named `TIUP_KEY`

3. RUN Init Mirror
Run Init-Mirror Github Action Workflow