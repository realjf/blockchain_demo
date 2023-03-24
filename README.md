# blockchain_demo

blockchain demo

### Usage

```sh
make run CMD=createwallet

Hello, world.
wallets file is empty!
save wallets file!
New address is: 12E6bMd7VxJXv7dbeWKhKAQ6y2VZvbGzdz


make run CMD=createwallet

Hello, world.
save wallets file!
New address is: 18CsFYnLTdcuLcpwtVHbagseYTcDN7qZLC


make run CMD='createblockchain -address 12E6bMd7VxJXv7dbeWKhKAQ6y2VZvbGzdz'

Hello, world.
badger 2023/03/24 17:51:42 INFO: All 0 tables opened in 0s
000034a8a668ce86dcefd5fa303ae0ee6359aca46418beccf701c4c5cbfa34e8
Genesis created
badger 2023/03/24 17:51:46 DEBUG: Storing value log head: {Fid:0 Len:42 Offset:624}
badger 2023/03/24 17:51:46 INFO: Got compaction priority: {level:0 score:1.73 dropPrefixes:[]}
badger 2023/03/24 17:51:46 INFO: Running for level: 0
badger 2023/03/24 17:51:46 DEBUG: LOG Compact. Added 3 keys. Skipped 0 keys. Iteration took: 105.411µs
badger 2023/03/24 17:51:46 DEBUG: Discard stats: map[]
badger 2023/03/24 17:51:46 INFO: LOG Compact 0->1, del 1 tables, add 1 tables, took 12.037073ms
badger 2023/03/24 17:51:46 INFO: Compaction for level: 0 DONE
badger 2023/03/24 17:51:46 INFO: Force compaction on level 0 done
Finished!


make run CMD='getbalance -address 12E6bMd7VxJXv7dbeWKhKAQ6y2VZvbGzdz'

Hello, world.
badger 2023/03/24 17:53:37 INFO: All 1 tables opened in 0s
badger 2023/03/24 17:53:37 INFO: Replaying file id: 0 at offset: 666
badger 2023/03/24 17:53:37 INFO: Replay took: 11.97µs
badger 2023/03/24 17:53:37 DEBUG: Value log discard stats empty
Balance of 12E6bMd7VxJXv7dbeWKhKAQ6y2VZvbGzdz: 0
badger 2023/03/24 17:53:37 INFO: Got compaction priority: {level:0 score:1.73 dropPrefixes:[]}


make run CMD='reindexutxo'

Hello, world.
badger 2023/03/24 17:54:55 INFO: All 1 tables opened in 0s
badger 2023/03/24 17:54:55 INFO: Replaying file id: 0 at offset: 666
badger 2023/03/24 17:54:55 INFO: Replay took: 1.998µs
badger 2023/03/24 17:54:55 DEBUG: Value log discard stats empty
Done! There are 1 transactions in the UTXO set.
badger 2023/03/24 17:54:55 DEBUG: Storing value log head: {Fid:0 Len:42 Offset:887}
badger 2023/03/24 17:54:55 INFO: Got compaction priority: {level:0 score:1.73 dropPrefixes:[]}
badger 2023/03/24 17:54:55 INFO: Running for level: 0
badger 2023/03/24 17:54:55 DEBUG: LOG Compact. Added 5 keys. Skipped 0 keys. Iteration took: 401.992µs
badger 2023/03/24 17:54:55 DEBUG: Discard stats: map[]
badger 2023/03/24 17:54:55 INFO: LOG Compact 0->1, del 2 tables, add 1 tables, took 16.979811ms
badger 2023/03/24 17:54:55 INFO: Compaction for level: 0 DONE
badger 2023/03/24 17:54:55 INFO: Force compaction on level 0 done


make run CMD='getbalance -address 12E6bMd7VxJXv7dbeWKhKAQ6y2VZvbGzdz'

Hello, world.
badger 2023/03/24 17:55:26 INFO: All 1 tables opened in 1ms
badger 2023/03/24 17:55:26 INFO: Replaying file id: 0 at offset: 929
badger 2023/03/24 17:55:26 INFO: Replay took: 2.173µs
badger 2023/03/24 17:55:26 DEBUG: Value log discard stats empty
Balance of 12E6bMd7VxJXv7dbeWKhKAQ6y2VZvbGzdz: 100
badger 2023/03/24 17:55:26 INFO: Got compaction priority: {level:0 score:1.73 dropPrefixes:[]}


make run CMD='listaddresses'

Hello, world.
12E6bMd7VxJXv7dbeWKhKAQ6y2VZvbGzdz
18CsFYnLTdcuLcpwtVHbagseYTcDN7qZLC

```
