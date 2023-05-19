## State diagram

![](/home/drag/.var/app/com.github.marktext.marktext/config/marktext/images/2023-05-20-02-31-41-image.png)



---

### Test with single worker



These tasks are added

```json
[
    {
        "name": "Tuhin1",
        "gender": "male",
        "country": "india"
    },
    {
        "name": "Tuhin2",
        "gender": "male",
        "country": "india"
    },
    {
        "name": "Tuhin3",
        "gender": "male",
        "country": "india"
    },
    {
        "name": "Tuhin4",
        "gender": "male",
        "country": "india"
    },
    {
        "name": "Tuhin5",
        "gender": "male",
        "country": "india"
    },
    {
        "name": "Tuhin6",
        "gender": "male",
        "country": "india"
    }, 
    {
        "name": "Tuhin7",
        "gender": "male",
        "country": "india"
    }
]
```



curl

```bash
curl --location 'http://localhost:5000/add' \
--header 'Content-Type: application/json' \
--data '[
    {
        "name": "Tuhin1",
        "gender": "male",
        "country": "india"
    },
    {
        "name": "Tuhin2",
        "gender": "male",
        "country": "india"
    },
    {
        "name": "Tuhin3",
        "gender": "male",
        "country": "india"
    },
    {
        "name": "Tuhin4",
        "gender": "male",
        "country": "india"
    },
    {
        "name": "Tuhin5",
        "gender": "male",
        "country": "india"
    },
    {
        "name": "Tuhin6",
        "gender": "male",
        "country": "india"
    }, 
    {
        "name": "Tuhin7",
        "gender": "male",
        "country": "india"
    }
]
'
```





Set the worker config and start immidiately0

```shell
curl --location 'http://localhost:5000/set-worker' \
--header 'Content-Type: application/json' \
--data '{
    "worker_count":1,
    "waiting_period": 10,
    "resting_period": 5,
    "auto_start": true
}   '
```



Verbose log

```shell
1684527645  |  166e436b21c5a07b0764e279a6cd907c  |  i'm the first one. doing start formalities
1684527645  |  166e436b21c5a07b0764e279a6cd907c  |  example entry formality task
1684527645  |  166e436b21c5a07b0764e279a6cd907c  |  waiting for 10 seconds before starting
1684527655  |  166e436b21c5a07b0764e279a6cd907c  |  sleeping for resting period of 5 seconds
1684527660  |  166e436b21c5a07b0764e279a6cd907c  |  sleeping for resting period of 5 seconds
1684527665  |  166e436b21c5a07b0764e279a6cd907c  |  retry limit reached, pushing the task to failed queue
1684527670  |  166e436b21c5a07b0764e279a6cd907c  |  task b16007a3a6ed406269f8c9051efddd13 failed, retrying at 1684527675 epoch
1684527670  |  166e436b21c5a07b0764e279a6cd907c  |  Qtask b16007a3a6ed406269f8c9051efddd13 added to pending queue for retry, waiting for 1s
1684527671  |  166e436b21c5a07b0764e279a6cd907c  |  sleeping for resting period of 5 seconds
1684527676  |  166e436b21c5a07b0764e279a6cd907c  |  task d5b355bf05d527e87ad7e2759ef96f5f failed, retrying at 1684527681 epoch
1684527676  |  166e436b21c5a07b0764e279a6cd907c  |  Qtask d5b355bf05d527e87ad7e2759ef96f5f added to pending queue for retry, waiting for 1s
1684527677  |  166e436b21c5a07b0764e279a6cd907c  |  sleeping for resting period of 5 seconds
1684527682  |  166e436b21c5a07b0764e279a6cd907c  |  task b16007a3a6ed406269f8c9051efddd13 failed, retrying at 1684527690 epoch
1684527682  |  166e436b21c5a07b0764e279a6cd907c  |  Qtask b16007a3a6ed406269f8c9051efddd13 added to pending queue for retry, waiting for 1s
1684527683  |  166e436b21c5a07b0764e279a6cd907c  |  task d5b355bf05d527e87ad7e2759ef96f5f failed, retrying at 1684527691 epoch
1684527683  |  166e436b21c5a07b0764e279a6cd907c  |  Qtask d5b355bf05d527e87ad7e2759ef96f5f added to pending queue for retry, waiting for 1s
1684527684  |  166e436b21c5a07b0764e279a6cd907c  |  retry task b16007a3a6ed406269f8c9051efddd13 has a future timestamp: 1684527690 re-adding to pending queue
1684527690  |  166e436b21c5a07b0764e279a6cd907c  |  retry task d5b355bf05d527e87ad7e2759ef96f5f has a future timestamp: 1684527691 re-adding to pending queue
1684527695  |  166e436b21c5a07b0764e279a6cd907c  |  task b16007a3a6ed406269f8c9051efddd13 failed, retrying at 1684527706 epoch
1684527695  |  166e436b21c5a07b0764e279a6cd907c  |  Qtask b16007a3a6ed406269f8c9051efddd13 added to pending queue for retry, waiting for 1s
1684527696  |  166e436b21c5a07b0764e279a6cd907c  |  task d5b355bf05d527e87ad7e2759ef96f5f failed, retrying at 1684527707 epoch
1684527696  |  166e436b21c5a07b0764e279a6cd907c  |  Qtask d5b355bf05d527e87ad7e2759ef96f5f added to pending queue for retry, waiting for 1s
1684527697  |  166e436b21c5a07b0764e279a6cd907c  |  retry task b16007a3a6ed406269f8c9051efddd13 has a future timestamp: 1684527706 re-adding to pending queue
1684527702  |  166e436b21c5a07b0764e279a6cd907c  |  retry task d5b355bf05d527e87ad7e2759ef96f5f has a future timestamp: 1684527707 re-adding to pending queue
1684527707  |  166e436b21c5a07b0764e279a6cd907c  |  sleeping for resting period of 5 seconds
1684527712  |  166e436b21c5a07b0764e279a6cd907c  |  task d5b355bf05d527e87ad7e2759ef96f5f failed, retrying at 1684527726 epoch
1684527712  |  166e436b21c5a07b0764e279a6cd907c  |  Qtask d5b355bf05d527e87ad7e2759ef96f5f added to pending queue for retry, waiting for 1s
1684527713  |  166e436b21c5a07b0764e279a6cd907c  |  retry task d5b355bf05d527e87ad7e2759ef96f5f has a future timestamp: 1684527726 re-adding to pending queue
1684527718  |  166e436b21c5a07b0764e279a6cd907c  |  popped the same Qtask d5b355bf05d527e87ad7e2759ef96f5f again. delaying  seconds loopbackoff
1684527718  |  166e436b21c5a07b0764e279a6cd907c  |  retry task d5b355bf05d527e87ad7e2759ef96f5f has a future timestamp: 1684527726 re-adding to pending queue
1684527723  |  166e436b21c5a07b0764e279a6cd907c  |  popped the same Qtask d5b355bf05d527e87ad7e2759ef96f5f again. delaying  seconds loopbackoff
1684527723  |  166e436b21c5a07b0764e279a6cd907c  |  retry task d5b355bf05d527e87ad7e2759ef96f5f has a future timestamp: 1684527726 re-adding to pending queue
1684527728  |  166e436b21c5a07b0764e279a6cd907c  |  task d5b355bf05d527e87ad7e2759ef96f5f failed, retrying at 1684527745 epoch
1684527728  |  166e436b21c5a07b0764e279a6cd907c  |  Qtask d5b355bf05d527e87ad7e2759ef96f5f added to pending queue for retry, waiting for 1s
1684527729  |  166e436b21c5a07b0764e279a6cd907c  |  popped the same Qtask d5b355bf05d527e87ad7e2759ef96f5f again. delaying  seconds loopbackoff
1684527729  |  166e436b21c5a07b0764e279a6cd907c  |  retry task d5b355bf05d527e87ad7e2759ef96f5f has a future timestamp: 1684527745 re-adding to pending queue
1684527734  |  166e436b21c5a07b0764e279a6cd907c  |  retry task d5b355bf05d527e87ad7e2759ef96f5f has a future timestamp: 1684527745 re-adding to pending queue
1684527739  |  166e436b21c5a07b0764e279a6cd907c  |  popped the same Qtask d5b355bf05d527e87ad7e2759ef96f5f again. delaying  seconds loopbackoff
1684527739  |  166e436b21c5a07b0764e279a6cd907c  |  retry task d5b355bf05d527e87ad7e2759ef96f5f has a future timestamp: 1684527745 re-adding to pending queue
1684527744  |  166e436b21c5a07b0764e279a6cd907c  |  popped the same Qtask d5b355bf05d527e87ad7e2759ef96f5f again. delaying  seconds loopbackoff
1684527744  |  166e436b21c5a07b0764e279a6cd907c  |  retry task d5b355bf05d527e87ad7e2759ef96f5f has a future timestamp: 1684527745 re-adding to pending queue
1684527749  |  166e436b21c5a07b0764e279a6cd907c  |  retry limit reached, pushing the task to failed queue
1684527754  |  166e436b21c5a07b0764e279a6cd907c  |  no Qtask pending
1684527754  |  166e436b21c5a07b0764e279a6cd907c  |  i am the last one. doing exit formalities
1684527754  |  166e436b21c5a07b0764e279a6cd907c  |  example exit formality task
```



The final batch report

```json
{
  "batch_duration": 113,
  "batch_id": "617ba0885dea15c41a54ae06f8665ffe",
  "end_time": 1684527754,
  "failed_tasks": [
    {
      "batch_id": "617ba0885dea15c41a54ae06f8665ffe",
      "created_at": 1684527641,
      "last_retry_timestamp": 1684527665,
      "remark": "NO_RETRY | user do not have whatsapp number",
      "serviceq_id": "ea1d64c967fb1adba1855474d0d86db0",
      "task": {
        "country": "india",
        "gender": "male",
        "name": "Tuhin3"
      },
      "task_id": "291f1e4b1813dc309efd6668da6976fc",
      "total_retry_attempts": 0
    },
    {
      "batch_id": "617ba0885dea15c41a54ae06f8665ffe",
      "created_at": 1684527641,
      "last_retry_timestamp": 1684527749,
      "remark": "RETRY | send for retry",
      "serviceq_id": "ea1d64c967fb1adba1855474d0d86db0",
      "task": {
        "country": "india",
        "gender": "male",
        "name": "Tuhin6"
      },
      "task_id": "d5b355bf05d527e87ad7e2759ef96f5f",
      "total_retry_attempts": 5
    }
  ],
  "passed_tasks": [
    {
      "batch_id": "617ba0885dea15c41a54ae06f8665ffe",
      "created_at": 1684527641,
      "last_retry_timestamp": 1684527655,
      "remark": "SUCCESS | the task was successful",
      "serviceq_id": "ea1d64c967fb1adba1855474d0d86db0",
      "task": {
        "country": "india",
        "gender": "male",
        "name": "Tuhin1"
      },
      "task_id": "114d6a2e9744a212c23b66aae8c82c33",
      "total_retry_attempts": 0
    },
    {
      "batch_id": "617ba0885dea15c41a54ae06f8665ffe",
      "created_at": 1684527641,
      "last_retry_timestamp": 1684527660,
      "remark": "SUCCESS | the task was successful",
      "serviceq_id": "ea1d64c967fb1adba1855474d0d86db0",
      "task": {
        "country": "india",
        "gender": "male",
        "name": "Tuhin2"
      },
      "task_id": "397abff60761c2d40a6550b8fb5503d6",
      "total_retry_attempts": 0
    },
    {
      "batch_id": "617ba0885dea15c41a54ae06f8665ffe",
      "created_at": 1684527641,
      "last_retry_timestamp": 1684527671,
      "remark": "SUCCESS | the task was successful",
      "serviceq_id": "ea1d64c967fb1adba1855474d0d86db0",
      "task": {
        "country": "india",
        "gender": "male",
        "name": "Tuhin5"
      },
      "task_id": "1763dbb3ea326917289a85ee5db99c96",
      "total_retry_attempts": 0
    },
    {
      "batch_id": "617ba0885dea15c41a54ae06f8665ffe",
      "created_at": 1684527641,
      "last_retry_timestamp": 1684527677,
      "remark": "SUCCESS | the task was successful",
      "serviceq_id": "ea1d64c967fb1adba1855474d0d86db0",
      "task": {
        "country": "india",
        "gender": "male",
        "name": "Tuhin7"
      },
      "task_id": "7903038831d7a3239e371004d0f0cf96",
      "total_retry_attempts": 0
    },
    {
      "batch_id": "617ba0885dea15c41a54ae06f8665ffe",
      "created_at": 1684527641,
      "last_retry_timestamp": 1684527707,
      "remark": "SUCCESS | the task was finally successful",
      "serviceq_id": "ea1d64c967fb1adba1855474d0d86db0",
      "task": {
        "country": "india",
        "gender": "male",
        "name": "Tuhin4"
      },
      "task_id": "b16007a3a6ed406269f8c9051efddd13",
      "total_retry_attempts": 3
    }
  ],
  "pending_tasks": [],
  "serviceq_id": "ea1d64c967fb1adba1855474d0d86db0",
  "serviceq_name": "drag",
  "start_time": 1684527641,
  "status": "FINISHED",
  "total_failed": 2,
  "total_pending": 0,
  "total_submitted": 7,
  "total_success": 5
}
```





---

Now the same ecperiment but this time with three workers



cut the workes 

```shell
curl --location 'http://localhost:5000/set-worker' \
--header 'Content-Type: application/json' \
--data '{
    "worker_count":3,
    "waiting_period": 10,
    "resting_period": 5,
    "auto_start": true
}   '
```



Verbose log

```shell
1684528204  |  7b4aae0f2e60770b61e782229f25884b  |  i'm the first one. doing start formalities
1684528204  |  7b4aae0f2e60770b61e782229f25884b  |  example entry formality task
1684528204  |  7b4aae0f2e60770b61e782229f25884b  |  waiting for 10 seconds before starting
1684528208  |  121468493e914a9d2f9e78059a79813d  |  waiting for 10 seconds before starting
1684528208  |  dc43b48da8d4e819b269bb00655ea2f0  |  waiting for 10 seconds before starting
1684528214  |  7b4aae0f2e60770b61e782229f25884b  |  sleeping for resting period of 5 seconds
1684528218  |  121468493e914a9d2f9e78059a79813d  |  retry limit reached, pushing the task to failed queue
1684528218  |  dc43b48da8d4e819b269bb00655ea2f0  |  sleeping for resting period of 5 seconds
1684528219  |  7b4aae0f2e60770b61e782229f25884b  |  sleeping for resting period of 5 seconds
1684528223  |  121468493e914a9d2f9e78059a79813d  |  sleeping for resting period of 5 seconds
1684528223  |  dc43b48da8d4e819b269bb00655ea2f0  |  task 525cccd9e70469b894c73995c83e04d8 failed, retrying at 1684528228 epoch
1684528223  |  dc43b48da8d4e819b269bb00655ea2f0  |  Qtask 525cccd9e70469b894c73995c83e04d8 added to pending queue for retry, waiting for 1s
1684528224  |  7b4aae0f2e60770b61e782229f25884b  |  sleeping for resting period of 5 seconds
1684528224  |  dc43b48da8d4e819b269bb00655ea2f0  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528228 re-adding to pending queue
1684528228  |  121468493e914a9d2f9e78059a79813d  |  task 525cccd9e70469b894c73995c83e04d8 failed, retrying at 1684528236 epoch
1684528228  |  121468493e914a9d2f9e78059a79813d  |  Qtask 525cccd9e70469b894c73995c83e04d8 added to pending queue for retry, waiting for 1s
1684528229  |  7b4aae0f2e60770b61e782229f25884b  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528236 re-adding to pending queue
1684528229  |  121468493e914a9d2f9e78059a79813d  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528229  |  121468493e914a9d2f9e78059a79813d  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528236 re-adding to pending queue
1684528229  |  dc43b48da8d4e819b269bb00655ea2f0  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528236 re-adding to pending queue
1684528234  |  7b4aae0f2e60770b61e782229f25884b  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528234  |  7b4aae0f2e60770b61e782229f25884b  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528236 re-adding to pending queue
1684528234  |  121468493e914a9d2f9e78059a79813d  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528234  |  121468493e914a9d2f9e78059a79813d  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528236 re-adding to pending queue
1684528234  |  dc43b48da8d4e819b269bb00655ea2f0  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528234  |  dc43b48da8d4e819b269bb00655ea2f0  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528236 re-adding to pending queue
1684528239  |  7b4aae0f2e60770b61e782229f25884b  |  task 525cccd9e70469b894c73995c83e04d8 failed, retrying at 1684528250 epoch
1684528239  |  7b4aae0f2e60770b61e782229f25884b  |  Qtask 525cccd9e70469b894c73995c83e04d8 added to pending queue for retry, waiting for 1s
1684528239  |  121468493e914a9d2f9e78059a79813d  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528239  |  121468493e914a9d2f9e78059a79813d  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528250 re-adding to pending queue
1684528239  |  dc43b48da8d4e819b269bb00655ea2f0  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528239  |  dc43b48da8d4e819b269bb00655ea2f0  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528250 re-adding to pending queue
1684528240  |  7b4aae0f2e60770b61e782229f25884b  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528250 re-adding to pending queue
1684528244  |  121468493e914a9d2f9e78059a79813d  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528244  |  121468493e914a9d2f9e78059a79813d  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528250 re-adding to pending queue
1684528244  |  dc43b48da8d4e819b269bb00655ea2f0  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528244  |  dc43b48da8d4e819b269bb00655ea2f0  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528250 re-adding to pending queue
1684528245  |  7b4aae0f2e60770b61e782229f25884b  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528245  |  7b4aae0f2e60770b61e782229f25884b  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528250 re-adding to pending queue
1684528249  |  121468493e914a9d2f9e78059a79813d  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528249  |  121468493e914a9d2f9e78059a79813d  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528250 re-adding to pending queue
1684528249  |  dc43b48da8d4e819b269bb00655ea2f0  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528249  |  dc43b48da8d4e819b269bb00655ea2f0  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528250 re-adding to pending queue
1684528250  |  7b4aae0f2e60770b61e782229f25884b  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528250  |  7b4aae0f2e60770b61e782229f25884b  |  task 525cccd9e70469b894c73995c83e04d8 failed, retrying at 1684528264 epoch
1684528250  |  7b4aae0f2e60770b61e782229f25884b  |  Qtask 525cccd9e70469b894c73995c83e04d8 added to pending queue for retry, waiting for 1s
1684528251  |  7b4aae0f2e60770b61e782229f25884b  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528251  |  7b4aae0f2e60770b61e782229f25884b  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528264 re-adding to pending queue
1684528254  |  121468493e914a9d2f9e78059a79813d  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528254  |  121468493e914a9d2f9e78059a79813d  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528264 re-adding to pending queue
1684528254  |  dc43b48da8d4e819b269bb00655ea2f0  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528254  |  dc43b48da8d4e819b269bb00655ea2f0  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528264 re-adding to pending queue
1684528256  |  7b4aae0f2e60770b61e782229f25884b  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528256  |  7b4aae0f2e60770b61e782229f25884b  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528264 re-adding to pending queue
1684528259  |  121468493e914a9d2f9e78059a79813d  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528259  |  121468493e914a9d2f9e78059a79813d  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528264 re-adding to pending queue
1684528259  |  dc43b48da8d4e819b269bb00655ea2f0  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528259  |  dc43b48da8d4e819b269bb00655ea2f0  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528264 re-adding to pending queue
1684528261  |  7b4aae0f2e60770b61e782229f25884b  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528261  |  7b4aae0f2e60770b61e782229f25884b  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528264 re-adding to pending queue
1684528264  |  121468493e914a9d2f9e78059a79813d  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528264  |  121468493e914a9d2f9e78059a79813d  |  task 525cccd9e70469b894c73995c83e04d8 failed, retrying at 1684528281 epoch
1684528264  |  121468493e914a9d2f9e78059a79813d  |  Qtask 525cccd9e70469b894c73995c83e04d8 added to pending queue for retry, waiting for 1s
1684528264  |  dc43b48da8d4e819b269bb00655ea2f0  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528264  |  dc43b48da8d4e819b269bb00655ea2f0  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528281 re-adding to pending queue
1684528265  |  121468493e914a9d2f9e78059a79813d  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528281 re-adding to pending queue
1684528266  |  7b4aae0f2e60770b61e782229f25884b  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528281 re-adding to pending queue
1684528269  |  dc43b48da8d4e819b269bb00655ea2f0  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528269  |  dc43b48da8d4e819b269bb00655ea2f0  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528281 re-adding to pending queue
1684528270  |  121468493e914a9d2f9e78059a79813d  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528281 re-adding to pending queue
1684528271  |  7b4aae0f2e60770b61e782229f25884b  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528281 re-adding to pending queue
1684528274  |  dc43b48da8d4e819b269bb00655ea2f0  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528274  |  dc43b48da8d4e819b269bb00655ea2f0  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528281 re-adding to pending queue
1684528275  |  121468493e914a9d2f9e78059a79813d  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528275  |  121468493e914a9d2f9e78059a79813d  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528281 re-adding to pending queue
1684528276  |  7b4aae0f2e60770b61e782229f25884b  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528276  |  7b4aae0f2e60770b61e782229f25884b  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528281 re-adding to pending queue
1684528279  |  dc43b48da8d4e819b269bb00655ea2f0  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528281 re-adding to pending queue
1684528280  |  121468493e914a9d2f9e78059a79813d  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528280  |  121468493e914a9d2f9e78059a79813d  |  retry task 525cccd9e70469b894c73995c83e04d8 has a future timestamp: 1684528281 re-adding to pending queue
1684528281  |  7b4aae0f2e60770b61e782229f25884b  |  popped the same Qtask 525cccd9e70469b894c73995c83e04d8 again. delaying  seconds loopbackoff
1684528281  |  7b4aae0f2e60770b61e782229f25884b  |  retry limit reached, pushing the task to failed queue
1684528284  |  dc43b48da8d4e819b269bb00655ea2f0  |  no Qtask pending
1684528284  |  dc43b48da8d4e819b269bb00655ea2f0  |  worker terminated
1684528285  |  121468493e914a9d2f9e78059a79813d  |  no Qtask pending
1684528285  |  121468493e914a9d2f9e78059a79813d  |  worker terminated
1684528286  |  7b4aae0f2e60770b61e782229f25884b  |  no Qtask pending
1684528286  |  7b4aae0f2e60770b61e782229f25884b  |  i am the last one. doing exit formalities
1684528286  |  7b4aae0f2e60770b61e782229f25884b  |  example exit formality task
1684528286  |  7b4aae0f2e60770b61e782229f25884b  |  worker terminated
```



Final batch report

```json
{
  "batch_duration": 82,
  "batch_id": "bcc815a75c0cc75324956cb6227a2052",
  "end_time": 1684528286,
  "failed_tasks": [
    {
      "batch_id": "bcc815a75c0cc75324956cb6227a2052",
      "created_at": 1684528204,
      "last_retry_timestamp": 1684528218,
      "remark": "NO_RETRY | user do not have whatsapp number",
      "serviceq_id": "ea1d64c967fb1adba1855474d0d86db0",
      "task": {
        "country": "india",
        "gender": "male",
        "name": "Tuhin3"
      },
      "task_id": "4446b719250f7ade338a168454444a76",
      "total_retry_attempts": 0
    },
    {
      "batch_id": "bcc815a75c0cc75324956cb6227a2052",
      "created_at": 1684528204,
      "last_retry_timestamp": 1684528281,
      "remark": "RETRY | send for retry",
      "serviceq_id": "ea1d64c967fb1adba1855474d0d86db0",
      "task": {
        "country": "india",
        "gender": "male",
        "name": "Tuhin6"
      },
      "task_id": "525cccd9e70469b894c73995c83e04d8",
      "total_retry_attempts": 5
    }
  ],
  "passed_tasks": [
    {
      "batch_id": "bcc815a75c0cc75324956cb6227a2052",
      "created_at": 1684528204,
      "last_retry_timestamp": 1684528214,
      "remark": "SUCCESS | the task was successful",
      "serviceq_id": "ea1d64c967fb1adba1855474d0d86db0",
      "task": {
        "country": "india",
        "gender": "male",
        "name": "Tuhin1"
      },
      "task_id": "b2adcdb1f38f168d4a96ed9a933dd5ec",
      "total_retry_attempts": 0
    },
    {
      "batch_id": "bcc815a75c0cc75324956cb6227a2052",
      "created_at": 1684528204,
      "last_retry_timestamp": 1684528218,
      "remark": "SUCCESS | the task was successful",
      "serviceq_id": "ea1d64c967fb1adba1855474d0d86db0",
      "task": {
        "country": "india",
        "gender": "male",
        "name": "Tuhin2"
      },
      "task_id": "4e33f47d0a92eab1bccf6974cdf8acdd",
      "total_retry_attempts": 0
    },
    {
      "batch_id": "bcc815a75c0cc75324956cb6227a2052",
      "created_at": 1684528204,
      "last_retry_timestamp": 1684528219,
      "remark": "SUCCESS | the task was finally successful",
      "serviceq_id": "ea1d64c967fb1adba1855474d0d86db0",
      "task": {
        "country": "india",
        "gender": "male",
        "name": "Tuhin4"
      },
      "task_id": "8f1d1ebc2955f091f36529af46ba40e5",
      "total_retry_attempts": 0
    },
    {
      "batch_id": "bcc815a75c0cc75324956cb6227a2052",
      "created_at": 1684528204,
      "last_retry_timestamp": 1684528223,
      "remark": "SUCCESS | the task was successful",
      "serviceq_id": "ea1d64c967fb1adba1855474d0d86db0",
      "task": {
        "country": "india",
        "gender": "male",
        "name": "Tuhin5"
      },
      "task_id": "71515a70c595476055bece6d4a2145ea",
      "total_retry_attempts": 0
    },
    {
      "batch_id": "bcc815a75c0cc75324956cb6227a2052",
      "created_at": 1684528204,
      "last_retry_timestamp": 1684528224,
      "remark": "SUCCESS | the task was successful",
      "serviceq_id": "ea1d64c967fb1adba1855474d0d86db0",
      "task": {
        "country": "india",
        "gender": "male",
        "name": "Tuhin7"
      },
      "task_id": "229b806cfea90c7ddef2cd45f7c2e65c",
      "total_retry_attempts": 0
    }
  ],
  "pending_tasks": [],
  "serviceq_id": "ea1d64c967fb1adba1855474d0d86db0",
  "serviceq_name": "drag",
  "start_time": 1684528204,
  "status": "FINISHED",
  "total_failed": 2,
  "total_pending": 0,
  "total_submitted": 7,
  "total_success": 5
}
```

Note: This time the batch completed faster 82 seconds instead of 113 seconds






















