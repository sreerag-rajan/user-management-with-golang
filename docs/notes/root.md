---
id: g3r7jbohqrgggifb3ov8ckv
title: Root
desc: ''
updated: 1693675062891
created: 1693673119808
---
# Welcome to User Management with Golang


## CONCEPT
<br>
Hi there!, This is an attempt by me to learn golang by building an user management application. The goal is achieve a basic User Management service with the following features

1. Create User
2. Login User
3. Forgot Password
4. Reset Password
5. Mail Verification through OTP

We will also implement key features of user management such as 
1. JWT Token
2. Password Hashing


## PROJECT STRUCTURE
So as I am new to GOLANG, I am not entirely sure how to structure my project. But taking guidance from youtube and some other online sources I have decided to structure it as follows:

- [[API|api]]
  > This is where the handler and middleware resides. The routes will connect to one of the handlers here and probable use one of the middlewares (If Needed). This will then connect to one of the services
- [[CONFIG|config]]
  > Config holds the major configurations, in this instance the DB
- [[MODEL|model]]
  > Models holds all the models
- [[REPOSITORY|repository]]
  > Repository holds the code that will enable to interact with the models. Its wierd that I need to add code to interact with each model here
- [[ROUTES|routes]]
  > As the name suggests this is where I have defined the routes to accessing the application
- [[SERVICE|service]]
  > The services hold the logic as well. API is the controller layer, whereas service is the logic resides. It will connect to the repositories to interact with the models
- [[UTILS|utils]]
  > Utils are for easy and reusable functions like hashing and creation of JWT tokens