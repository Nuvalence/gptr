# GCP Permission-to-Roles Utility

Rather than sorting through [Google's documentation](https://cloud.google.com/iam/docs/permissions-reference)
to find a role that contains the permission you need, simply pass the permission as an argument to get 
the list of roles that contain that permissionl  

`gptr <some.google.iam-permission>`

**Example**

```
╰─ gptr storage.hmacKeys.list
Owner (roles/owner)
Editor (roles/editor)
Viewer (roles/viewer)
Security Admin (roles/iam.securityAdmin)
Security Reviewer (roles/iam.securityReviewer)
Storage HMAC Key Admin (roles/storage.hmacKeyAdmin)
```

## Build
Using the [correct operating system paramters](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63) for your environment, run this command.

`env GOOS=<your-os> GOARCH=<your-arch> go build -o bin/gptr-<your-os><version> cmd/gptr/main.go`


## Installation
The executable `gptr-darwin0.0.1` in [/bin](./bin) is a go binary built for 64-bit mac. 
If you're using mac, rename to `gptr` and put on your path. Otherwise build 
using the instructions above and use the resulting binary.


## Future Work
- Implement fuzzy search
- Tab completion
