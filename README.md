# hertz-session (dev-version)

> Refer to [here](https://github.com/cloudwego/hertz-examples/tree/main/bizdemo/hertz_session) for more information

## Introduce

A demo with `Hertz` and `Session`

- Use `thrift` IDL to define `HTTP` interface
- Use `hz` to generate code
- Use `hertz-contrib/sessions` to store sessions
- Use `hertz-contrib/csrf` to prevent Cross-Site Request Forgery attacks
- Use `Gorm` and `MySQL`
- Use `AdminLTE` as frontend page

## hertz-contrib/sessions

Use Hertz's sessions extension, refer to [hertz-contrib/sessions](https://github.com/hertz-contrib/sessions)

## hertz-contrib/csrf

Use Hertz's csrf extension, refer to [hertz-contrib/csrf](https://github.com/hertz-contrib/csrf)

## IDL

This demo use `thrift` IDL to define `HTTP` interface. The specific interface define in [user.thrift](idl/user.thrift)

## Code generation tool

This demo use `hz` to generate code. The use of `hz` refers to [hz](https://www.cloudwego.io/docs/hertz/tutorials/toolkit/toolkit/)

The `hz` commands used can be found in [Makefile](Makefile)

## Binding and Validate

The use of binding and Validate refers
to [Binding and Validate](https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/binding-and-validate/)

## Gorm

This demo use `Gorm` to operate `MySQL` and refers to [Gorm](https://gorm.io/)

## AdminLTE

This demo captures the [AdminLTE](https://github.com/ColorlibHQ/AdminLTE) login and registration page as the front-end page.

## How to run

### Run MySQL and Redis docker

```bash
docker-compose up
```

### Run demo

```bash
go build -o hertz-session && ./hertz-session
```