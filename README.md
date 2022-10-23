# tiup-community
tiup community 


## init tiup mirror 
1. fork 12cat/tiup-community

2. create TIUP_KEY 
```
# creare rsa key peer
openssl rand -base64 32 > tiup.key
```
3. SET Github Actions secrets
Store the resulting `tiup.key` in Github Actions secrets, named `TIUP_KEY`

4. Delete init.lock
```
rm init.lock
git commit -a -m "delete init lock"
git push
```

4. Init Mirror

Run Init-Mirror Github Action Workflow

## Config Github Page
1. Build and deployment

- 1.1 choose Source: `Deploy from a branch` 

- 1.2 choose Branch: `main`/`docs`

2. Save Config

3. View github page
- `{{ your_repo_name }}.github.io/tiup-community/root.json`

like: https://12cat-tidb.github.io/tiup-community/root.json
