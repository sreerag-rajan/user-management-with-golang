---
id: d36vkkc2sew64s5tqxltw25
title: User
desc: ''
updated: 1693673811256
created: 1693673607907
---

## SCHEMA

## FIELDS

| Field Name    | Type          | Join                            | Comments        |
| ------------- | ------------- | -------------                   |   ------------- |
| ObjectId      | objectId      | -                               | Mongo Id        |
| firstName     | string        | -                               | -               |
| middleName    | string        | -                               | -               |
| lastName      | string        | -                               | -               |
| email         | string        | -                               | -               |
| password      | string        | -                               | -               |
| role          | objectId      | [[model.user_role]].objectId    | MongoID         |
