# Common microservice for InstaUpload.
---
This is the common module for InstaUpload upload mainly this module contain all the gRPC genrated code for InstaUpload, as well as common function to be used in any of the other microservice.

## TODO:
- ### Add env.go file.
    -   Add function to get value from env, if not present then return the fallback.
- ### Add make file.
    - Add make file with command used to genrated `.go` code file from `.proto` file.
