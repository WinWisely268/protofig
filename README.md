# Protofig

An example of declaring textual config using protobuf.

Given a file it try to fulfill these rules:

```
[<ComponentName>]:
```<content_of_proto>```
:
[<ComponentName>]:
```<content_of_proto>```
:
```

### Keywords

- compstart: for the start of a config
- compend: for the end of config
