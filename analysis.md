
/   => triedb root
/account            =>0
/token-0.0.1        =>1
    action:
            0 => create


/account/{wallet_addr}/ =>              hash("/account/{wallet_addr}/") => 0x123456789
val
{
    pk:xxxxxxx
    balance:xxxxx
    nonce:xxxx
}



//
hash_index          hash("/account/adsfds/afdsf/0x512412341234")    = 0x1c21341234

hash_short_index    0x1c21341234 => 0x1c21

//

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
    mint_max_limit,         //    upper bound is uint64          
}

TX (1,0,{mint_max_limit}) , {0x1234,signature}  //  



```


```
2. token minter

/token-0.0.1/{token_addr}/minter/{minter_addr}      => hash("/token-0.0.1/{token_addr}/minter/{minter_addr}") => hash_index_1   , tx( hash_index_1=> hash_short_index_2)
val:
{
    mint_amount     //how much this user has minted
    mint_max_limit
}

Tx (update ,"/token-0.0.1/{token_addr}/minter/{minter_addr}"=>hash_short_index_2)

signature ("/token-0.0.1/{token_addr}/minter/{minter_addr},mint_max_limit")

Tx (,hash_short_index_2,signature) 

wallet => server (full_node)
full node => broadcast


block 
{
    TX : (,hash_short_index_2,signature) 
    Tx : (,hash_short_index_2,signature)  
}



/////////


```
 
```
3. token holder

/token-0.0.1/{token_addr}/holder/{holder_addr} => file

val:
{
    amount
    exchange lock amount
}


/account/token-0.0.1/{token_addr}/

```


TX(1,2,short_hash_index,amount,short_hash_index)  gas = op gas + store gas 

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

/token-0.0.1/{token_addr}/{allower_addr}/{spender_public_key}       [hash index required]

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


