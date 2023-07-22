# token-checker

This repos helps to verify the token address is valid or not.

### Show some :heart: and star the repo to support the project

### Getting Started
- The concept revolves around utilizing Trust Wallet as a reference to validate token addresses. The information is then stored in a MongoDB database, streamlining its accessibility for developers seeking better control and ease of use.
- MongoDB / Golang are required to run this tool. 
### How to Run

#### 1. Setup:
- Clone the repos
    ```shell
        git clone https://github.com/trongdth/token-checker
        cd token-checker
    ```
- Copy `config.yml`
    ```shell
        cp config/conf.yaml.default config/conf.yaml 
    ```
- Open `config/conf.yaml` and change `mongo_db_url`, `mongo_db_name` with your configuration.
    ```shell
        nano config/conf.yaml
    ```
- Update submodule
    ```shell
        git submodule update --remote --merge
    ```
#### 2. Sync data:
- Run application with `sync` command
    ```shell
        go run main.go sync
    ```
    Notes: 
    - Update submodule whenever you want to sync a newest data from trust wallet source.
#### 3. Verify token address:
- Run application with `check` command
    ```shell
         go run main.go check --address=0x84d7aeef42d38a5ffc3ccef853e1b82e4958659d16a7de736a29c55fbbeb0114::staked_aptos_coin::StakedAptosCoin
    ```
    Notes: 
    - To see data from `check` command, you need to run `sync` before.

### What's next

- [ ] Run `sync` command in parallel.
- [ ] Add more trusted sources.
### Created & Maintained By

[Trong Dinh](https://github.com/trongdth) ([@trongdth](https://www.twitter.com/trongdth))

> If you found this project helpful or you learned something from the source code and want to thank me, consider buying me a cup of :coffee:
>
> * [Ethereum address: 0x54aa977d8aAec77935DC4b31DEE40d47054fEB2a]
