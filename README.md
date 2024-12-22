# AuthZEN Interop PDP Example
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)  
Experimental PDP implementation with OPA for [OpenID AuthZEN interop](https://authzen-interop.net/).

[OpenID AuthZEN Interop](https://authzen-interop.net/) is interoperability testing and results for AuthZEN implementations.  
This implementation works as a PDP (Policy Decision Point) as defined in the [NIST SP 800-162](https://csrc.nist.gov/pubs/sp/800/162/upd2/final). In addition, it also stores information that is expected to come from PIP (Policy Information Point) and PAP (Policy Administration Point). So, it can receive requests from PEP and evaluate access immediately after starting up.

## Usage

This PDP works with [AuthZEN Interop TODO Application](https://github.com/openid/authzen/tree/main/interop). If you are only testing AuthZEN Interop, you can send requests directly from PEP to PDP, so there is no need to configure the front end. 

![](https://github.com/kg0r0/authzen-interop-pdp-example/blob/assets/todoapp.drawio.png?raw=true)

### Try AUthZEN Interop with this PDP implementation

1) Execute the following command to start PDP (default port is `8001`):

```bash
$ git clone 
$ go run main.go
```

2) Clone [authzen-todo-backend](https://github.com/openid/authzen/tree/main/interop/authzen-todo-backend) and execute the following command to test the PDP as the target:

```bash
$ git clone https://github.com/openid/authzen.git
$ cd interop/authzen-todo-backend
$ yarn build
$ yarn test http://localhost:8001
```

You will be able to get results like below:

<details>
<summary> Result </summary>

```
$ yarn test http://localhost:8001
yarn run v1.22.22
$ node build/test/runner.js http://localhost:8001
PASS REQ: {"subject":{"type":"user","id":"CiRmZDA2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_read_user"},"resource":{"type":"user","id":"beth@the-smiths.com"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDA2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_read_user"},"resource":{"type":"user","id":"rick@the-citadel.com"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDA2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_read_todos"},"resource":{"type":"todo","id":"todo-1"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDA2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_create_todo"},"resource":{"type":"todo","id":"todo-1"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDA2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_update_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b92","properties":{"ownerID":"rick@the-citadel.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDA2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_update_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b91","properties":{"ownerID":"morty@the-citadel.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDA2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_delete_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b92","properties":{"ownerID":"rick@the-citadel.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDA2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_delete_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b91","properties":{"ownerID":"morty@the-citadel.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDE2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_read_user"},"resource":{"type":"user","id":"beth@the-smiths.com"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDE2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_read_user"},"resource":{"type":"user","id":"morty@the-citadel.com"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDE2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_read_todos"},"resource":{"type":"todo","id":"todo-1"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDE2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_create_todo"},"resource":{"type":"todo","id":"todo-1"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDE2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_update_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b92","properties":{"ownerID":"rick@the-citadel.com"}}}
PASS REQ: {"subject":{"type":"user","_note":"ID for Morty","id":"CiRmZDE2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_update_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b91","properties":{"ownerID":"morty@the-citadel.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDE2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_delete_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b92","properties":{"ownerID":"rick@the-citadel.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDE2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_delete_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b91","properties":{"ownerID":"morty@the-citadel.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDI2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_read_user"},"resource":{"type":"user","id":"beth@the-smiths.com"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDI2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_read_user"},"resource":{"type":"user","id":"summer@the-smiths.com"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDI2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_read_todos"},"resource":{"type":"todo","id":"todo-1"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDI2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_create_todo"},"resource":{"type":"todo","id":"todo-1"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDI2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_update_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b92","properties":{"ownerID":"rick@the-citadel.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDI2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_update_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b93","properties":{"ownerID":"summer@the-smiths.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDI2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_delete_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b92","properties":{"ownerID":"rick@the-citadel.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDI2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_delete_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b93","properties":{"ownerID":"summer@the-smiths.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDM2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_read_user"},"resource":{"type":"user","id":"beth@the-smiths.com"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDM2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_read_user"},"resource":{"type":"user","id":"beth@the-smiths.com"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDM2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_read_todos"},"resource":{"type":"todo","id":"todo-1"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDM2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_create_todo"},"resource":{"type":"todo","id":"todo-1"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDM2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_update_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b92","properties":{"ownerID":"rick@the-citadel.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDM2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_update_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b94","properties":{"ownerID":"beth@the-smiths.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDM2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_delete_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b92","properties":{"ownerID":"rick@the-citadel.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDM2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_delete_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b94","properties":{"ownerID":"beth@the-smiths.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDQ2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_read_user"},"resource":{"type":"user","id":"beth@the-smiths.com"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDQ2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_read_user"},"resource":{"type":"user","id":"jerry@the-smiths.com"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDQ2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_read_todos"},"resource":{"type":"todo","id":"todo-1"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDQ2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_create_todo"},"resource":{"type":"todo","id":"todo-1"}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDQ2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_update_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b92","properties":{"ownerID":"rick@the-citadel.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDQ2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_update_todo"},"resource":{"type":"todo","id":"240d0db-8ff0-41ec-98b2-34a096273b95","properties":{"ownerID":"jerry@the-smiths.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDQ2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_delete_todo"},"resource":{"type":"todo","id":"7240d0db-8ff0-41ec-98b2-34a096273b92","properties":{"ownerID":"rick@the-citadel.com"}}}
PASS REQ: {"subject":{"type":"user","id":"CiRmZDQ2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs"},"action":{"name":"can_delete_todo"},"resource":{"type":"todo","id":"240d0db-8ff0-41ec-98b2-34a096273b95","properties":{"ownerID":"jerry@the-smiths.com"}}}
✨  Done in 0.42s.
```

</details>

**Note:** You can see the request and response payloads for each of the API requests in the [Todo interop scenario.](https://authzen-interop.net/docs/scenarios/todo-1.1/#overview-of-the-scenario).  
I've confirmed that this PDP implementation passes the test in [Authorization API 1.1 – draft 01](https://openid.github.io/authzen/authorization-api-1_1_01). However, test results for other versions are unconfirmed.

## References
- https://github.com/openid/authzen
- https://openid.net/wg/authzen/specifications/
- https://csrc.nist.gov/pubs/sp/800/162/upd2/final