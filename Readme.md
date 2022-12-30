## Description
This repo is just a test to understand a problem happening with Golang packages not being versionated correctly. 

The error is: 
```
go get github.com/arpesam/go-pkg-test@v3.0.0
go: github.com/arpesam/go-pkg-test@v3.0.0: invalid version: module contains a go.mod file, so module path must match major version ("github.com/arpesam/go-pkg-test/v3")
``` 
 

## Results of this test
In order to create new versions of packages in Golang, we have these options :
1. In the major branch:
    - Update the module with the new version at the end of the path, like `github.com/myuser/pkg/v2` 
    - Tag this version as `v2.0.0` 
    - Then you can download this package using `go get github.com/myuser/pkg/v2`
>NOTE: You don't need to do this for version v0 or v1

2. Create a version (v2) folder in the root directory. This is the suggested approach since clients that are already using the current version will not be impacted. They will be able to pick which version they want to use. A bennefit of this approach is the possibility to maintain two versions in the same repo and branch. Don't forget to tag the branch correctly.

3. Create a new branch with the updated module name, adding `v2` at the end, tag it with the version name and then make this the default branch. Using this approach is similar to creating a subfolder (option 2), the difference is that instead of two folders we now gonna have a branch per version.

## References
- https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher
- https://stackoverflow.com/questions/70058951/project-structure-with-v2-go-modules
- https://github.com/golang/go/issues/35732
- https://medium.com/rungo/anatomy-of-modules-in-go-c8274d215c16
- https://go.dev/doc/modules/managing-dependencies#getting_version