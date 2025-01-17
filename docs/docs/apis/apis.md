---
title: ZITADEL APIs
---

All of our APIs are generated by proto defintions. You can find all the proto definitions in the [Proto API Definitions](proto/auth).

> More about [Protocol Buffer](https://developers.google.com/protocol-buffers)

## Swagger Documentation

We provide some json files for the swagger documentation of our APIs with the following link: [https://api.zitadel.ch/openapi/v2/swagger/](https://api.zitadel.ch/openapi/v2/swagger/)
The easiest way to have a look at them is, to import them in the [Swagger Editor](https://editor.swagger.io/)

## Authentication API aka Auth

The authentication API (aka Auth API) is used for all operations on the currently logged in user.
The user id is taken from the sub claim in the token.

| Service | URI                                                                                                                         |
|:--------|:----------------------------------------------------------------------------------------------------------------------------|
| REST    | [https://api.zitadel.ch/auth/v1/](https://api.zitadel.ch/auth/v1/)                                                          |
| GRPC    | [https://api.zitadel.ch/caos.zitadel.auth.api.v1.AuthService/](https://api.zitadel.ch/caos.zitadel.auth.api.v1.AuthService) |

> At a later date we might expose functions to build your own login GUI
> You can build your own user Register GUI already by utilizing the [Management API](#management)

[Latest API Version](https://github.com/caos/zitadel/blob/main/proto/zitadel/auth.proto)


## Management API

The management API is as the name states the interface where systems can mutate IAM objects like, organisations, projects, clients, users and so on if they have the necessary access rights.
To identify the current organisation you can send a header `x-zitadel-orgid` or if no header is set, the organisation of the authenticated user is set.

| Service | URI                                                                                                                                                 |
|:--------|:----------------------------------------------------------------------------------------------------------------------------------------------------|
| REST    | [https://api.zitadel.ch/management/v1/](https://api.zitadel.ch/management/v1/)                                                                      |
| GRPC    | [https://api.zitadel.ch/caos.zitadel.management.api.v1.ManagementService/](https://api.zitadel.ch/caos.zitadel.management.api.v1.ManagementService) |

[Latest API Version](https://github.com/caos/zitadel/blob/main/proto/zitadel/management.proto)


## Administration API aka Admin

This API is intended to configure and manage the IAM itself.

| Service | URI                                                                                                                             |
|:--------|:--------------------------------------------------------------------------------------------------------------------------------|
| REST    | [https://api.zitadel.ch/admin/v1/](https://api.zitadel.ch/admin/v1/)                                                            |
| GRPC    | [https://api.zitadel.ch/caos.zitadel.admin.api.v1.AdminService/](https://api.zitadel.ch/caos.zitadel.admin.api.v1.AdminService) |

[Latest API Version](https://github.com/caos/zitadel/blob/main/proto/zitadel/admin.proto)

## Assets API


The Assets API allows you to up- and download all kinds of assets. This can be files such as logos, fonts or user avatar.


| Service | URI                                                                                                                             |
|:--------|:--------------------------------------------------------------------------------------------------------------------------------|
| REST    | [https://api.zitadel.ch/assets/v1/](https://api.zitadel.ch/assets/v1/)                                                            |
