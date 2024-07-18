
examples:

```
1. token

path:
/token-0.0.1/{token_addr}  

val:
{
    creator_addr,
    block_height,
    status,                 // normal , stopped , expiring , expired
    expire_block_height,    // blockheight to be expired
    total_minted,
    total_burned,
    mint_max_limit,         //    
}

```


```
2. token minter

/token-0.0.1/{token_addr}/minter/{minter_addr} 
val:
{
    mint_max_limit
}

```
 

3. token holder

/token-0.0.1/{token_addr}/holder/{holder_addr} => file

val:
{
    amount
}

```

5. token exchange

/token-0.0.1/{token_1_addr}/exchange/{token_2_addr}/{seller_addr} => file

val:
{
    sell:[
        ratio_1=>amount,
        ratio_2=>amount,
        ratio_3=>amount,
        ratio_4=>amount,
    ],
    buy:[
        ratio_1=>amount,
        ratio_2=>amount,
        ratio_3=>amount,
        ratio_4=>amount,
    ]
}

hash("/token-0.0.1/{token_1_addr}/exchange/{token_2_addr}/{seller_addr}") == {hash_x}

```


```
6. token spender signature

/token-0.0.1/{token_addr}/{allower_addr}/{spender_public_key}

path short hash needed
val:
{
    spent_amount
    spend_max_limit
    status : # stopped ,normal
}

```


```
10.file system

/file-0.0.1/root_domain/folder1
/file-0.0.1/root_domain/folder1/{bigfile_name}/0
/file-0.0.1/root_domain/folder1/{bigfile_name}/1
/file-0.0.1/root_domain/folder1/{bigfile_name}/2
/file-0.0.1/root_domain/folder1/{bigfile_name}/3


```




```

10.dns system

/file-0.0.1/root_domain/subdomain1
/file-0.0.1/root_domain/subdomain2
/file-0.0.1/root_domain/subdomain3

```


