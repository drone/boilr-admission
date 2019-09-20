# boilr-admission

This is a [boilr template](https://github.com/tmrts/boilr) for creating an admission extension. Use an admission extension to control registration and authentication. Get started by cloning the project and installing the template:

```console
$ cd boilr-admission
$ boilr template save . drone-admission -f
```

create a project in directory my-admitter:

```console
$ boilr template use drone-admission my-admitter
```

enter the docker registry name for this project:

```text
[?] Please choose a value for "DockerRepository" [default: "foo/bar"]:
```

enter the go module name:

```text
[?] Please choose a value for "GoModule" [default: "github.com/foo/bar":
```
